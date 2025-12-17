package grpcx

import (
	"context"

	"google.golang.org/grpc"
)

// RoutedUnaryServerInterceptor returns a gRPC unary server interceptor that routes requests to different interceptors.
func RoutedUnaryServerInterceptor(router func(ctx context.Context, req any, info *grpc.UnaryServerInfo) grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	if router == nil {
		panic("router is nil")
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		interceptor := router(ctx, req, info)
		if interceptor == nil {
			return handler(ctx, req)
		}
		return interceptor(ctx, req, info, handler)
	}
}
