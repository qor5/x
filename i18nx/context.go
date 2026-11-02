package i18nx

import (
	"context"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ctxKeyI18N struct{}

func NewContext(ctx context.Context, ib *I18N) context.Context {
	return context.WithValue(ctx, ctxKeyI18N{}, ib)
}

func FromContext(ctx context.Context) (*I18N, bool) {
	if ib, ok := ctx.Value(ctxKeyI18N{}).(*I18N); ok {
		return ib, true
	}
	return nil, false
}

func MustFromContext(ctx context.Context) *I18N {
	ib, ok := FromContext(ctx)
	if !ok {
		panic("i18n not found in context")
	}
	return ib
}

func LanguageFromContext(ctx context.Context) language.Tag {
	ib, ok := FromContext(ctx)
	if !ok {
		return FallbackTag
	}
	return ib.LanguageFromContext(ctx)
}

func MustSprintf(ctx context.Context, key message.Reference, args ...any) string {
	return MustFromContext(ctx).Sprintf(LanguageFromContext(ctx), key, args...)
}
