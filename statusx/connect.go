package statusx

import (
	"context"
	"log/slog"
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
			err = TranslateError(err, ib, lang)
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

// ConvertToConnectError converts a statusx.StatusError to a connect.Error.
func ConvertToConnectError(err error) *connect.Error {
	var ce *connect.Error
	if errors.As(err, &ce) {
		return ce
	}

	st := status.Convert(err).Proto()
	cerr := connect.NewError(connect.Code(st.Code), errors.New(st.Message))
	for _, d := range st.Details {
		if ed, e := connect.NewErrorDetail(d); e == nil {
			cerr.AddDetail(ed)
		} else {
			slog.Error("failed to convert connect error detail", "error", e, "detail", d)
		}
	}
	return cerr //nolint:errhandle
}

func WriteConnectErrorOnly(errWriter *connect.ErrorWriter, w http.ResponseWriter, r *http.Request, err error) (written bool) {
	var se *StatusError
	if errors.As(err, &se) {
		err = ConvertToConnectError(err)
	}
	var ce *connect.Error
	if errors.As(err, &ce) {
		_ = errWriter.Write(w, r, err)
		return true
	}
	return false
}
