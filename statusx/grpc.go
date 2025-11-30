package statusx

import (
	"context"

	"google.golang.org/grpc"

	"github.com/qor5/x/v3/i18nx"
	"github.com/theplant/appkit/logtracing"
)

func UnaryServerInterceptor(ib *i18nx.I18N) grpc.UnaryServerInterceptor {
	if ib == nil {
		panic("i18n.I18N is required")
	}
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = i18nx.NewContext(ctx, ib)
		lang := ib.LanguageFromContext(ctx)
		resp, err := handler(ctx, req)
		if err != nil {
			TracingReason(ctx, err)
		}
		err = TranslateError(err, ib, lang)
		return resp, err
	}
}

// TracingReason extracts the reason field from statusx.Status and adds it to tracing span
func TracingReason(ctx context.Context, err error) {
	span := logtracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	if st, ok := FromError(err); ok {
		if reason := st.Reason(); reason != "" {
			span.AppendKVs("reason", reason)
		}
	}
}
