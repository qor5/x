package normalize

import (
	"context"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

	"github.com/qor5/x/v3/grpcx"
)

func UnaryConnectInterceptor[T any](svc T) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			reqMD := grpcx.MetadataFromHeader(req.Header(), req.Peer().Addr)
			ctx = metadata.NewIncomingContext(ctx, reqMD)

			callMeta := &CallMeta{
				ClientKind: ClientKindPublic,
				Service:    svc,
				FullMethod: req.Spec().Procedure,
				Req:        req.Any(),
			}
			resMD := metadata.Pairs()

			nlz, _ := fromContext(ctx)
			if nlz == nil {
				nlz = &normalization{}
			}
			if nlz.callMeta == nil {
				nlz.callMeta = callMeta
			}
			if nlz.setHeader == nil {
				nlz.setHeader = func(key, value string) {
					resMD.Set(key, value)
				}
			}
			ctx = context.WithValue(ctx, ctxKeyNormalization{}, nlz)

			res, err := next(ctx, req)
			if err != nil {
				if len(resMD) > 0 {
					var ce *connect.Error
					if errors.As(err, &ce) {
						for key, vv := range resMD {
							for _, v := range vv {
								ce.Meta().Add(key, v)
							}
						}
					}
				}
				return nil, err
			}

			for key, vv := range resMD {
				for _, v := range vv {
					res.Header().Add(key, v)
				}
			}
			return res, nil
		}
	}
}
