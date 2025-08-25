package connectx

import (
	"net/http"
	"slices"

	"connectrpc.com/connect"

	"github.com/qor5/x/v3/normalize"
)

func NewHandler[T any](
	fn func(svc T, opts ...connect.HandlerOption) (string, http.Handler),
	svc T, opts ...connect.HandlerOption,
) (string, http.Handler) {
	opts = slices.Concat(
		[]connect.HandlerOption{
			connect.WithInterceptors(
				normalize.UnaryConnectInterceptor(svc),
			),
		},
		opts,
	)
	return fn(svc, opts...)
}

func Response[T any](res *T, err error) (*connect.Response[T], error) {
	if err != nil {
		return nil, err //nolint:errhandle
	}
	return connect.NewResponse(res), nil
}
