package ratelimiterx

import (
	"context"

	"github.com/theplant/ratelimiter"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor creates a gRPC unary server interceptor for rate limiting
func UnaryServerInterceptor(limiter ratelimiter.RateLimiter, evaluator Evaluator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if err := allow(ctx, limiter, evaluator); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
