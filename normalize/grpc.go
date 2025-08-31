package normalize

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	stdmetadata "google.golang.org/grpc/metadata"

	"github.com/qor5/x/v3/statusx"
)

// HeaderEnsureClientKind is just used for reverse proxy such as grpc-gateway
const HeaderEnsureClientKind = "x-ensure-client-kind"

func UnaryServerInterceptor(defClientKind ClientKind) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, xerr error) {
		clientKind := defClientKind
		headerClientKind := metadata.ExtractIncoming(ctx).Get(HeaderEnsureClientKind)
		if headerClientKind != "" {
			if headerClientKind != string(ClientKindPublic) {
				return nil, statusx.NewCodef(codes.Internal, "only %q is supported for %q", ClientKindPublic, HeaderEnsureClientKind).Err()
			}
			clientKind = ClientKind(headerClientKind)
		}

		callMeta := &CallMeta{
			ClientKind: clientKind,
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

		if decorator, ok := nlz.callMeta.Service.(ContextDecorator); ok {
			ctx = decorator.DecorateContext(ctx)
		}

		defer func() {
			if len(resMD) > 0 {
				if serr := grpc.SetHeader(ctx, resMD); serr != nil {
					serr = statusx.WrapCode(serr, codes.Internal, "failed to set header").Err()
					slog.ErrorContext(ctx, "failed to set header", "error", serr)
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
