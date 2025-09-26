package statusx

import (
	"context"

	"google.golang.org/grpc"

	"github.com/qor5/x/v3/i18nx"
)

func UnaryServerInterceptor(ib *i18nx.I18N) grpc.UnaryServerInterceptor {
	if ib == nil {
		panic("i18n.I18N is required")
	}
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx = i18nx.NewContext(ctx, ib)
		lang := ib.LanguageFromContext(ctx)
		resp, err := handler(ctx, req)
		err = TranslateError(err, ib, lang)
		return resp, err
	}
}
