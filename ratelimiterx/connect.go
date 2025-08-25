package ratelimiterx

import (
	"context"

	"connectrpc.com/connect"
	"github.com/theplant/ratelimiter"
)

// UnaryConnectInterceptor creates a Connect unary interceptor for rate limiting
func UnaryConnectInterceptor(limiter ratelimiter.RateLimiter, evaluator Evaluator) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if err := allow(ctx, limiter, evaluator); err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}
