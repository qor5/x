package normalize

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/qor5/x/v3/grpcx"
)

type HTTPMeta struct {
	ClientKind ClientKind
	R          *http.Request
	W          http.ResponseWriter
}

func MustHTTPMetaFromContext(ctx context.Context) *HTTPMeta {
	return mustFromContext(ctx).httpMeta
}

func HTTPMetaFromContext(ctx context.Context) (*HTTPMeta, error) {
	nlz, err := fromContext(ctx)
	if err != nil {
		return nil, err
	}
	return nlz.httpMeta, nil
}

func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqMD := grpcx.MetadataFromHeader(r.Header, r.RemoteAddr)
		ctx := metadata.NewIncomingContext(r.Context(), reqMD)

		nlz := &normalization{
			httpMeta: &HTTPMeta{ClientKind: ClientKindPublic, W: w, R: r},
			setHeader: func(key, value string) {
				w.Header().Set(key, value)
			},
		}
		ctx = context.WithValue(ctx, ctxKeyNormalization{}, nlz)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
