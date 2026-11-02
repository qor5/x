package grpcx

import (
	"context"
	"errors"
	"net/http"
	"sync/atomic"
	"testing"

	"github.com/qor5/x/v3/statusx"
	"github.com/stretchr/testify/require"
	"github.com/theplant/appkit/logtracing"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestErrorUnaryServerInterceptor_DedupWithinSameInstance(t *testing.T) {
	var callCount atomic.Int32
	testErr := errors.New("test error")

	interceptor := ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		callCount.Add(1)
		return err
	})

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}

	// Simulate the interceptor being called twice in the same chain (via ChainUnaryServerInterceptor)
	// The inner interceptor calls the outer one's handler, which eventually returns the same error
	chained := ChainUnaryServerInterceptor(interceptor, interceptor)

	handler := func(ctx context.Context, req any) (any, error) {
		return nil, testErr
	}

	_, err := chained(context.Background(), nil, info, handler)
	require.ErrorIs(t, err, testErr)

	// Same interceptor instance appears twice, but same error should only be handled once per instance
	// Since it's the same instance used twice, the error is marked as handled after the first call
	// The second call in the chain sees the same error already handled
	require.Equal(t, int32(1), callCount.Load(), "errHandler should be called 1 time for same instance")
}

func TestErrorUnaryServerInterceptor_IndependentScopes(t *testing.T) {
	var callCount1, callCount2 atomic.Int32
	testErr := errors.New("test error")

	// Create two different interceptor instances
	interceptor1 := ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		callCount1.Add(1)
		return err
	})

	interceptor2 := ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		callCount2.Add(1)
		return err
	})

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}

	// Chain two different interceptor instances
	chained := ChainUnaryServerInterceptor(interceptor1, interceptor2)

	handler := func(ctx context.Context, req any) (any, error) {
		return nil, testErr
	}

	_, err := chained(context.Background(), nil, info, handler)
	require.ErrorIs(t, err, testErr)

	// Each interceptor instance should handle the error independently
	require.Equal(t, int32(1), callCount1.Load(), "interceptor1 errHandler should be called 1 time")
	require.Equal(t, int32(1), callCount2.Load(), "interceptor2 errHandler should be called 1 time")
}

func TestErrorUnaryServerInterceptor_NoErrorNoCall(t *testing.T) {
	var callCount atomic.Int32

	interceptor := ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		callCount.Add(1)
		return err
	})

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}

	handler := func(ctx context.Context, req any) (any, error) {
		return "success", nil
	}

	resp, err := interceptor(context.Background(), nil, info, handler)
	require.NoError(t, err)
	require.Equal(t, "success", resp)
	require.Equal(t, int32(0), callCount.Load(), "errHandler should not be called when no error")
}

func TestErrorUnaryServerInterceptor_ErrorCanBeTransformed(t *testing.T) {
	originalErr := errors.New("original error")
	transformedErr := errors.New("transformed error")

	interceptor := ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		return transformedErr
	})

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}

	handler := func(ctx context.Context, req any) (any, error) {
		return nil, originalErr
	}

	_, err := interceptor(context.Background(), nil, info, handler)
	require.ErrorIs(t, err, transformedErr)
}

// fakeNotifier records the arguments of the most recent Notify call so tests
// can assert on what DefaultErrorUnaryServerInterceptor reports.
type fakeNotifier struct {
	calls int
	err   any
	extra map[string]any
}

func (n *fakeNotifier) Notify(err any, _ *http.Request, extra map[string]any) {
	n.calls++
	n.err = err
	n.extra = extra
}

// captureExporter records every span exported while it is registered.
type captureExporter struct {
	spans []*logtracing.SpanData
}

func (e *captureExporter) ExportSpan(sd *logtracing.SpanData) {
	e.spans = append(e.spans, sd)
}

// findErrorInfo returns the ErrorInfo carried in a status' details, mirroring
// how statusx.FromError reads it back.
func findErrorInfo(details []any) *errdetails.ErrorInfo {
	for _, d := range details {
		if ei, ok := d.(*errdetails.ErrorInfo); ok {
			return ei
		}
	}
	return nil
}

// kvMap folds a logtracing keyvals slice into a map for easy assertions.
func kvMap(keyvals []any) map[string]any {
	m := map[string]any{}
	for i := 0; i+1 < len(keyvals); i += 2 {
		if k, ok := keyvals[i].(string); ok {
			m[k] = keyvals[i+1]
		}
	}
	return m
}

func TestDefaultErrorUnaryServerInterceptor_NotifiesInternalWithReasonAndDetails(t *testing.T) {
	notifier := &fakeNotifier{}
	interceptor := DefaultErrorUnaryServerInterceptor(notifier)

	md := map[string]string{"tenant": "acme", "trace_id": "abc123"}
	statusErr := statusx.New(codes.Internal, "INTERNAL_BOOM", "boom").WithMetadata(md).Err()

	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}
	handler := func(ctx context.Context, req any) (any, error) {
		return nil, statusErr
	}

	_, err := interceptor(context.Background(), nil, info, handler)
	require.ErrorIs(t, err, statusErr)

	require.Equal(t, 1, notifier.calls, "notifier should be called once for Internal errors")
	require.Equal(t, "INTERNAL_BOOM", notifier.extra["err.reason"])

	details, ok := notifier.extra["err.details"].([]any)
	require.True(t, ok, "err.details should be reported as []any")
	require.NotEmpty(t, details, "an Internal error should carry status details")

	errorInfo := findErrorInfo(details)
	require.NotNil(t, errorInfo, "status details should carry an ErrorInfo")
	require.Equal(t, "INTERNAL_BOOM", errorInfo.GetReason())
	require.Equal(t, md, errorInfo.GetMetadata())
}

func TestDefaultErrorUnaryServerInterceptor_DoesNotNotifyNonInternal(t *testing.T) {
	notifier := &fakeNotifier{}
	interceptor := DefaultErrorUnaryServerInterceptor(notifier)

	statusErr := statusx.New(codes.InvalidArgument, "INVALID_ARGUMENT", "bad").Err()
	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}
	handler := func(ctx context.Context, req any) (any, error) {
		return nil, statusErr
	}

	_, err := interceptor(context.Background(), nil, info, handler)
	require.ErrorIs(t, err, statusErr)
	require.Equal(t, 0, notifier.calls, "notifier should not be called for non-Internal/Unknown errors")
}

func TestDefaultErrorUnaryServerInterceptor_AppendsReasonAndDetailsToSpan(t *testing.T) {
	exporter := &captureExporter{}
	logtracing.RegisterExporter(exporter)
	defer logtracing.UnregisterExporter(exporter)

	// The span KVs are appended regardless of the gRPC code, so use a
	// non-Internal error to keep this test independent of the notifier path.
	interceptor := DefaultErrorUnaryServerInterceptor(&fakeNotifier{})

	md := map[string]string{"trace_id": "abc123"}
	statusErr := statusx.New(codes.InvalidArgument, "INVALID_ARGUMENT", "bad").WithMetadata(md).Err()
	info := &grpc.UnaryServerInfo{FullMethod: "/test.Service/Method"}
	handler := func(ctx context.Context, req any) (any, error) {
		return nil, statusErr
	}

	ctx, _ := logtracing.StartSpan(context.Background(), "test", logtracing.WithSampler(logtracing.AlwaysSample()))
	_, err := interceptor(ctx, nil, info, handler)
	require.ErrorIs(t, err, statusErr)
	logtracing.EndSpan(ctx, err)

	require.Len(t, exporter.spans, 1)
	kvs := kvMap(exporter.spans[0].Keyvals)

	require.Equal(t, "INVALID_ARGUMENT", kvs["err.reason"])
	details, ok := kvs["err.details"].([]any)
	require.True(t, ok, "err.details should be appended to the span as []any")

	errorInfo := findErrorInfo(details)
	require.NotNil(t, errorInfo, "span details should carry an ErrorInfo")
	require.Equal(t, md, errorInfo.GetMetadata())
}
