package grpcx

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	kitlog "github.com/theplant/appkit/log"
	"google.golang.org/grpc"
)

// LoggingUnaryServerInterceptor returns a new unary server interceptor that logs gRPC calls at various levels.
func LoggingUnaryServerInterceptor(l kitlog.Logger, opts ...logging.Option) grpc.UnaryServerInterceptor {
	interceptor := logging.UnaryServerInterceptor(logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = l.Debug().Log(largs...)
		case logging.LevelInfo:
			_ = l.Info().Log(largs...)
		case logging.LevelWarn:
			_ = l.Warn().Log(largs...)
		case logging.LevelError:
			_ = l.Error().Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	}), opts...)

	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctx = kitlog.Context(ctx, l)
		ctx = logging.InjectLogField(ctx, "req_id", uuid.New())
		return interceptor(ctx, req, info, handler)
	}
}
