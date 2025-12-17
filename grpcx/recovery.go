package grpcx

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// RecoveryUnaryServerInterceptor returns a gRPC unary server interceptor that recovers from panics.
func RecoveryUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, xerr error) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					xerr = errors.Wrap(err, "panic")
				} else {
					xerr = errors.Errorf("recovered from panic: %v", r)
				}
				slog.ErrorContext(ctx, "recovered from panic", "error", xerr, "method", info.FullMethod)
				return
			}
		}()
		return handler(ctx, req)
	}
}
