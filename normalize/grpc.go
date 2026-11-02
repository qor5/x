package normalize

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	stdmetadata "google.golang.org/grpc/metadata"
)

// GRPCUnaryServerInterceptor is a unary server interceptor for gRPC
func GRPCUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return UnaryServerInterceptor(ClientKindUndefined)
}

// HeaderEnsureClientKind is just used for reverse proxy such as grpc-gateway
const HeaderEnsureClientKind = "x-ensure-client-kind"

// Deprecated: use GRPCUnaryServerInterceptor instead
func UnaryServerInterceptor(defClientKind ClientKind) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, xerr error) {
		clientKind := defClientKind
		headerClientKind := metadata.ExtractIncoming(ctx).Get(HeaderEnsureClientKind)
		if headerClientKind != "" {
			if headerClientKind != string(ClientKindPublic) {
				return nil, errors.Errorf("only %q is supported for %q", ClientKindPublic, HeaderEnsureClientKind)
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
