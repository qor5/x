package ratelimiterx

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theplant/ratelimiter"
	"github.com/theplant/ratelimiter/sqlrl"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"gorm.io/gorm"

	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/healthz/testdata/gen"
	"github.com/qor5/x/v3/normalize"
)

// Test gRPC service implementation using healthz testdata
type testGRPCService struct {
	gen.UnimplementedTestServiceServer
}

func (s *testGRPCService) Echo(ctx context.Context, req *gen.EchoRequest) (*gen.EchoResponse, error) {
	return &gen.EchoResponse{Message: req.GetMessage()}, nil
}

var db *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()
	suite := gormx.MustStartTestSuite(ctx)
	defer func() { _ = suite.Stop(ctx) }()

	db = suite.DB()

	// Setup rate limiter tables
	if err := sqlrl.Migrate(context.Background(), db, "rate_limits"); err != nil {
		panic(err)
	}
	if err := sqlrl.Migrate(context.Background(), db, "rate_limits_connect"); err != nil {
		panic(err)
	}

	m.Run()
}

func TestUnaryServerInterceptor(t *testing.T) {
	// Create database-backed rate limiter
	limiter, err := sqlrl.New(db, "rate_limits")
	require.NoError(t, err)

	t.Run("allow request with no rate limit", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			// No rate limiting
			return nil, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		resp, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "test"})
		require.NoError(t, err)
		assert.Equal(t, "test", resp.GetMessage())
	})

	t.Run("handle empty reserve requests array", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			// Return empty array - should pass through without rate limiting
			return []*ratelimiter.ReserveRequest{}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		resp, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "empty"})
		require.NoError(t, err)
		assert.Equal(t, "empty", resp.GetMessage())
	})

	t.Run("allow request within rate limit", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return []*ratelimiter.ReserveRequest{
				{
					Key:              "test-key-allowed",
					Tokens:           1,
					DurationPerToken: time.Millisecond * 100, // 10 tokens per second
					Burst:            10,                     // High burst
				},
			}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		resp, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "allowed"})
		require.NoError(t, err)
		assert.Equal(t, "allowed", resp.GetMessage())
	})

	t.Run("block request when rate limited", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return []*ratelimiter.ReserveRequest{
				{
					Key:              "test-key-blocked-simple",
					Tokens:           2, // Request more tokens than burst allows
					DurationPerToken: time.Second,
					Burst:            1, // Only 1 token available
					MaxFutureReserve: 0, // Don't allow future reservations
				},
			}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		_, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "blocked"})
		require.Error(t, err)
		st := status.Convert(err)
		t.Logf("Error code: %v, Message: %s", st.Code(), st.Message())
		// When requesting more tokens than burst allows, sqlrl returns an error
		// which gets wrapped as codes.Internal by the allow() function
		assert.Equal(t, codes.Internal, st.Code())
	})

	t.Run("return error when evaluator fails", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return nil, errors.New("evaluator failed")
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		_, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "error"})
		require.Error(t, err)
		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
	})

	t.Run("return error when max future reserve is not zero", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return []*ratelimiter.ReserveRequest{
				{
					Key:              "test-key-future",
					Tokens:           1,
					DurationPerToken: time.Millisecond * 100,
					Burst:            10,
					MaxFutureReserve: time.Second, // Non-zero value should trigger error
				},
			}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		_, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "future"})
		require.Error(t, err)
		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Contains(t, st.Message(), "expect max future reserve is equal to 0")
	})

	t.Run("handle multiple rate limit requests with mixed results", func(t *testing.T) {
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return []*ratelimiter.ReserveRequest{
				{
					Key:              "test-key-multi-pass",
					Tokens:           1,
					DurationPerToken: time.Millisecond * 10, // Fast replenishment
					Burst:            5,                     // Allow multiple tokens
				},
				{
					Key:              "test-key-multi-fail",
					Tokens:           100, // Request way more than burst
					DurationPerToken: time.Hour,
					Burst:            1,
				},
			}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		_, err := client.Echo(context.Background(), &gen.EchoRequest{Message: "multi"})
		// Should fail because one of the requests exceeds limit
		require.Error(t, err)
		st := status.Convert(err)
		t.Logf("Multi request error: Code=%v, Message=%s", st.Code(), st.Message())
		// When requesting way more tokens than available, the rate limiter returns an error
		// which gets wrapped as codes.Internal
		assert.Equal(t, codes.Internal, st.Code())
	})
}

// setupGRPCServer creates a test gRPC server with the interceptor
func setupGRPCServer(t *testing.T, limiter ratelimiter.RateLimiter, evaluator Evaluator) (*grpc.Server, gen.TestServiceClient) {
	// Create gRPC server with interceptors
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			normalize.GRPCUnaryServerInterceptor(),
			UnaryServerInterceptor(limiter, evaluator),
		),
	)

	// Register test service
	testService := &testGRPCService{}
	gen.RegisterTestServiceServer(server, testService)

	// Setup server with bufconn
	listener := bufconn.Listen(1024 * 1024)
	go func() {
		_ = server.Serve(listener)
	}()

	// Create client
	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	client := gen.NewTestServiceClient(conn)
	return server, client
}

func TestUnaryServerInterceptor_CallMeta(t *testing.T) {
	limiter, err := sqlrl.New(db, "rate_limits")
	require.NoError(t, err)

	var capturedCallMeta *normalize.CallMeta
	evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
		capturedCallMeta = callMeta
		return nil, nil
	}

	server, client := setupGRPCServer(t, limiter, evaluator)
	defer server.GracefulStop()

	// Make request
	testReq := &gen.EchoRequest{Message: "test"}
	_, err = client.Echo(context.Background(), testReq)
	require.NoError(t, err)

	// Verify CallMeta was populated correctly
	require.NotNil(t, capturedCallMeta)
	assert.Equal(t, "/healthz.testdata.TestService/Echo", capturedCallMeta.FullMethod)
	assert.NotNil(t, capturedCallMeta.Service)
	assert.Equal(t, testReq.GetMessage(), capturedCallMeta.Req.(*gen.EchoRequest).GetMessage())
}

func TestUnaryServerInterceptor_ErrorMetadata(t *testing.T) {
	limiter, err := sqlrl.New(db, "rate_limits")
	require.NoError(t, err)

	t.Run("verify error contains rate limit metadata", func(t *testing.T) {
		// Test true rate limiting: first consume the token, then get rate limited
		evaluator := func(ctx context.Context, callMeta *normalize.CallMeta) ([]*ratelimiter.ReserveRequest, error) {
			return []*ratelimiter.ReserveRequest{
				{
					Key:              "test-key-metadata",
					Tokens:           1,
					DurationPerToken: time.Hour, // Very slow replenishment
					Burst:            1,         // Only 1 token
				},
			}, nil
		}

		server, client := setupGRPCServer(t, limiter, evaluator)
		defer server.GracefulStop()

		// First request should succeed and consume the only available token
		resp1, err1 := client.Echo(context.Background(), &gen.EchoRequest{Message: "first"})
		require.NoError(t, err1)
		assert.Equal(t, "first", resp1.GetMessage())

		// Second request should be rate limited (bucket is now empty)
		_, err2 := client.Echo(context.Background(), &gen.EchoRequest{Message: "second"})
		require.Error(t, err2)

		st := status.Convert(err2)
		t.Logf("Second request error: Code=%v, Message=%s", st.Code(), st.Message())

		// This is true rate limiting (reservation.OK = false), so codes.ResourceExhausted
		assert.Equal(t, codes.ResourceExhausted, st.Code())
		assert.Equal(t, "ratelimit exceeded", st.Message())

		// Verify error details contain rate limiting metadata
		details := st.Details()
		assert.NotEmpty(t, details, "Error should contain metadata details")

		// Extract ErrorInfo from details
		var errorInfo *errdetails.ErrorInfo
		for _, detail := range details {
			if ei, ok := detail.(*errdetails.ErrorInfo); ok {
				errorInfo = ei
				break
			}
		}
		require.NotNil(t, errorInfo, "Error should contain ErrorInfo")
		require.NotEmpty(t, errorInfo.Metadata, "ErrorInfo should contain metadata")

		// Verify the custom reason is set correctly
		assert.Equal(t, "RATE_LIMITED", errorInfo.Reason, "Error reason should be RATE_LIMITED")

		// Verify metadata contains required fields
		metadata := errorInfo.Metadata
		assert.Contains(t, metadata, "timeToAct", "Metadata should contain timeToAct field")
		assert.Contains(t, metadata, "reservedAt", "Metadata should contain reservedAt field")

		// Parse and validate timeToAct
		timeToActStr, exists := metadata["timeToAct"]
		assert.True(t, exists, "timeToAct should exist in metadata")
		assert.NotEmpty(t, timeToActStr, "timeToAct should not be empty")

		// Parse and validate reservedAt
		reservedAtStr, exists := metadata["reservedAt"]
		assert.True(t, exists, "reservedAt should exist in metadata")
		assert.NotEmpty(t, reservedAtStr, "reservedAt should not be empty")
	})
}
