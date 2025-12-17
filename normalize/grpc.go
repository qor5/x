package normalize

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	stdmetadata "google.golang.org/grpc/metadata"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, xerr error) {
		callMeta := &CallMeta{
			Service:    info.Server,
			FullMethod: info.FullMethod,
			Req:        req,
		}
		resMD := stdmetadata.Pairs()
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

		defer func() {
			if len(resMD) > 0 {
				if serr := grpc.SetHeader(ctx, resMD); serr != nil {
					serr = errors.Wrap(serr, "failed to set grpc header")
					slog.ErrorContext(ctx, "Failed to set grpc header", "error", serr)
					if xerr == nil {
						xerr = serr
					}
				}
			}
		}()
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
