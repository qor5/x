package healthz_test

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/test/bufconn"

	"github.com/qor5/x/v3/healthz"
	"github.com/qor5/x/v3/healthz/testdata/gen"
)

// testServiceImpl implements the TestService
type testServiceImpl struct {
	gen.UnimplementedTestServiceServer
}

func (s *testServiceImpl) Echo(ctx context.Context, req *gen.EchoRequest) (*gen.EchoResponse, error) {
	return &gen.EchoResponse{Message: req.Message}, nil
}

func TestUnaryServerInterceptor(t *testing.T) {
	// Track middleware calls
	var mu sync.Mutex
	middlewareCalls := make([]string, 0)

	// Business middleware that should be bypassed for health checks
	businessMiddleware := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		mu.Lock()
		middlewareCalls = append(middlewareCalls, info.FullMethod)
		mu.Unlock()
		time.Sleep(1 * time.Millisecond) // Simulate processing
		return handler(ctx, req)
	}

	// Create health server
	healthServer := health.NewServer()
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// Create gRPC server with health interceptor first to bypass other middleware
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			healthz.UnaryServerInterceptor(healthServer), // Health interceptor first
			businessMiddleware,                           // Business middleware
		),
	)
	defer server.GracefulStop()

	// Register services
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	gen.RegisterTestServiceServer(server, &testServiceImpl{})

	// Setup server with bufconn
	listener := bufconn.Listen(1024 * 1024)
	go func() {
		_ = server.Serve(listener)
	}()

	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer conn.Close()

	healthClient := grpc_health_v1.NewHealthClient(conn)
	testClient := gen.NewTestServiceClient(conn)
	ctx := context.Background()

	// Test 1: Health check bypasses middleware
	mu.Lock()
	middlewareCalls = middlewareCalls[:0] // Reset
	mu.Unlock()

	_, err = healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
	require.NoError(t, err)

	mu.Lock()
	healthCalls := len(middlewareCalls)
	mu.Unlock()
	assert.Equal(t, 0, healthCalls, "Health check should bypass business middleware")

	// Test 2: Business request goes through middleware
	_, err = testClient.Echo(ctx, &gen.EchoRequest{Message: "test"})
	require.NoError(t, err)

	mu.Lock()
	businessCalls := len(middlewareCalls)
	mu.Unlock()
	assert.Equal(t, 1, businessCalls, "Business request should go through middleware")
	assert.Equal(t, gen.TestService_Echo_FullMethodName, middlewareCalls[0])
}
