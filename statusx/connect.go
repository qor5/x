package statusx

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/i18nx"
	"google.golang.org/grpc/status"
)

func UnaryConnectInterceptor(ib *i18nx.I18N, shouldConvert func(ctx context.Context, req connect.AnyRequest) bool) connect.UnaryInterceptorFunc {
	if ib == nil {
		panic("i18n.I18N is required")
	}
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			ctx = i18nx.NewContext(ctx, ib)
			lang := ib.LanguageFromContext(ctx)
			res, err := next(ctx, req)
			err = TranslateError(ib, lang, err)
			if err != nil {
				convert := true
				if shouldConvert != nil && !shouldConvert(ctx, req) {
					convert = false
				}
				if convert {
					err = ConvertToConnectError(err)
				}
			}
			return res, err
		})
	}
}

func ConvertToConnectError(err error) *connect.Error {
	st := status.Convert(err).Proto()
	cerr := connect.NewError(connect.Code(st.Code), errors.New(st.Message))
	for _, v := range st.Details {
		ed, _ := connect.NewErrorDetail(v)
		cerr.AddDetail(ed)
	}
	return cerr //nolint:errhandle
}

func WriteConnectErrorOnly(errWriter *connect.ErrorWriter, w http.ResponseWriter, r *http.Request, err error) (written bool) {
	if se := new(StatusError); errors.As(err, &se) {
		err = ConvertToConnectError(err)
	}
	if ce := new(connect.Error); errors.As(err, &ce) {
		_ = errWriter.Write(w, r, err)
		return true
	}
	return false
}
