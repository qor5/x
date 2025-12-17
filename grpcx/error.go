package grpcx

import (
	"context"
	"fmt"

	"github.com/qor5/x/v3/statusx"
	"github.com/theplant/appkit/errornotifier"
	"github.com/theplant/appkit/logtracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ErrorUnaryServerInterceptor creates a gRPC unary server interceptor that deduplicates errors
// within the scope of a single interceptor instance while allowing different instances to handle
// the same error independently.
func ErrorUnaryServerInterceptor(errHandler func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error) grpc.UnaryServerInterceptor {
	key := new(int) // Each interceptor instance gets its own unique key (use int to avoid zero-size optimization)
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		handledErrors, ok := ctx.Value(key).(map[error]bool)
		if !ok {
			handledErrors = map[error]bool{}
			ctx = context.WithValue(ctx, key, handledErrors)
		}

		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		if handledErrors[err] {
			return resp, err
		}
		handledErrors[err] = true

		err = errHandler(ctx, req, info, err)

		return resp, err
	}
}

// DefaultErrorUnaryServerInterceptor creates a gRPC unary server interceptor that notifies errors using the provided notifier.
// It deduplicates errors within the scope of a single interceptor instance.
func DefaultErrorUnaryServerInterceptor(notifier errornotifier.Notifier) grpc.UnaryServerInterceptor {
	return ErrorUnaryServerInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, err error) error {
		st := statusx.Convert(err)

		span := logtracing.SpanFromContext(ctx)
		if span != nil {
			span.AppendKVs(
				"err.stacktrace", fmt.Sprintf("%+v", err),
				"err.reason", st.Reason(),
			)
		}

		switch st.Code() {
		case codes.Unknown, codes.Internal:
			notifier.Notify(err, nil, map[string]any{
				"full_method":    info.FullMethod,
				"err.stacktrace": fmt.Sprintf("%+v", err),
				"err.reason":     st.Reason(),
			})
		}

		return err
	})
}
