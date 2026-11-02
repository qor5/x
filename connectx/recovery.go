package connectx

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
)

func RecoveryUnaryConnectInterceptor(handler recovery.RecoveryHandlerFuncContext) connect.UnaryInterceptorFunc {
	if handler == nil {
		panic("recovery.RecoveryHandlerFuncContext is required")
	}
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (_ connect.AnyResponse, retErr error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}
			defer func() {
				if r := recover(); r != nil {
					// net/http checks for ErrAbortHandler with ==, so we should too.
					if r == http.ErrAbortHandler { //nolint:errorlint
						panic(r) //nolint:forbidigo
					}
					retErr = handler(ctx, r)
				}
			}()
			res, err := next(ctx, req)
			return res, err
		})
	}
}
