package grpcx

import (
	"context"
	"log/slog"
	"slices"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/theplant/appkit/logtracing"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func withTraceIDFields(ctx context.Context) logging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return logging.Fields{"traceID", span.TraceID().String()}
	}
	span := logtracing.SpanFromContext(ctx)
	if span != nil {
		return logging.Fields{"traceID", span.TraceID().String()}
	}
	return nil
}

var defaultLoggingOptions = []logging.Option{
	logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	logging.WithFieldsFromContext(withTraceIDFields),
}

func LoggingUnaryServerInterceptor(opts ...logging.Option) grpc.UnaryServerInterceptor {
	return LoggingUnaryServerInterceptorFactory(slog.Default())(opts...)
}

func LoggingUnaryServerInterceptorFactory(logger *slog.Logger) func(opts ...logging.Option) grpc.UnaryServerInterceptor {
	if logger == nil {
		panic("logger is required")
	}
	return func(opts ...logging.Option) grpc.UnaryServerInterceptor {
		return logging.UnaryServerInterceptor(interceptorLogger(logger), slices.Concat(defaultLoggingOptions, opts)...)
	}
}

func LoggingUnaryClientInterceptor(opts ...logging.Option) grpc.UnaryClientInterceptor {
	return LoggingUnaryClientInterceptorFactory(slog.Default())(opts...)
}

func LoggingUnaryClientInterceptorFactory(logger *slog.Logger) func(opts ...logging.Option) grpc.UnaryClientInterceptor {
	if logger == nil {
		panic("logger is required")
	}
	return func(opts ...logging.Option) grpc.UnaryClientInterceptor {
		return logging.UnaryClientInterceptor(interceptorLogger(logger), slices.Concat(defaultLoggingOptions, opts)...)
	}
}
