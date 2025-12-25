package grpcx

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
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
