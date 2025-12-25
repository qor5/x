package grpcx

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryServerInterceptor chains multiple unary server interceptors into one.
// The first interceptor will be the outermost wrapper.
// The implementation mirrors gRPC's internal chaining strategy.
func ChainUnaryServerInterceptor(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	switch len(interceptors) {
	case 0:
		return nil
	case 1:
		return interceptors[0]
	default:
		return chainUnaryServerInterceptor(interceptors)
	}
}

// ChainUnaryClientInterceptor chains multiple unary client interceptors into one.
// The first interceptor will be the outermost wrapper.
// The implementation mirrors gRPC's internal chaining strategy.
func ChainUnaryClientInterceptor(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	switch len(interceptors) {
	case 0:
		return nil
	case 1:
		return interceptors[0]
	default:
		return chainUnaryClientInterceptor(interceptors)
	}
}

// chainUnaryServerInterceptor is adapted from gRPC's internal server interceptor chaining.
func chainUnaryServerInterceptor(interceptors []grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return interceptors[0](ctx, req, info, getChainUnaryServerHandler(interceptors, 0, info, handler))
	}
}

func getChainUnaryServerHandler(interceptors []grpc.UnaryServerInterceptor, curr int, info *grpc.UnaryServerInfo, finalHandler grpc.UnaryHandler) grpc.UnaryHandler {
	if curr == len(interceptors)-1 {
		return finalHandler
	}
	return func(ctx context.Context, req any) (any, error) {
		return interceptors[curr+1](ctx, req, info, getChainUnaryServerHandler(interceptors, curr+1, info, finalHandler))
	}
}

// chainUnaryClientInterceptor is adapted from gRPC's internal client interceptor chaining.
func chainUnaryClientInterceptor(interceptors []grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return interceptors[0](ctx, method, req, reply, cc, getChainUnaryClientInvoker(interceptors, 0, cc, invoker), opts...)
	}
}

func getChainUnaryClientInvoker(interceptors []grpc.UnaryClientInterceptor, curr int, cc *grpc.ClientConn, finalInvoker grpc.UnaryInvoker) grpc.UnaryInvoker {
	if curr == len(interceptors)-1 {
		return finalInvoker
	}
	return func(ctx context.Context, method string, req, reply any, _ *grpc.ClientConn, opts ...grpc.CallOption) error {
		return interceptors[curr+1](ctx, method, req, reply, cc, getChainUnaryClientInvoker(interceptors, curr+1, cc, finalInvoker), opts...)
	}
}
