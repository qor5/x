package statusx

import (
	"context"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"

	"github.com/qor5/x/v3/hook"
	"github.com/qor5/x/v3/i18nx"
)

type (
	HTTPWriteErrorInput struct {
		Conf *HTTPErrorWriterConfig
		W    http.ResponseWriter
		R    *http.Request
		Err  error
	}
	HTTPWriteErrorOutput struct {
		Written bool
	}
	HTTPWriteErrorFunc func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error)
)

type HTTPErrorWriterConfig struct {
	I18N           *i18nx.I18N
	writeErrorHook func(next HTTPWriteErrorFunc) HTTPWriteErrorFunc
}

func (c *HTTPErrorWriterConfig) WithHTTPWriteErrorHook(hooks ...hook.Hook[HTTPWriteErrorFunc]) *HTTPErrorWriterConfig {
	c.writeErrorHook = hook.Prepend(c.writeErrorHook, hooks...)
	return c
}

func HTTPErrorWriter(conf *HTTPErrorWriterConfig) func(http.Handler) http.Handler {
	if conf == nil || conf.I18N == nil {
		panic("ErrorWriterConfig.I18N is required")
	}
	return func(next http.Handler) http.Handler {
		errWriter := connect.NewErrorWriter()
		defWriteError := func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
			lang := input.Conf.I18N.LanguageFromContext(ctx)
			err := TranslateError(input.Err, input.Conf.I18N, lang)
			written := WriteConnectErrorOnly(errWriter, input.W, input.R, err)
			return &HTTPWriteErrorOutput{Written: written}, nil
		}
		writeError := defWriteError
		if conf.writeErrorHook != nil {
			writeError = conf.writeErrorHook(writeError)
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			newCtx := i18nx.NewContext(r.Context(), conf.I18N)
			defer func() {
				if v := recover(); v != nil {
					var written bool
					if err, ok := v.(error); ok && err != nil {
						output, werr := writeError(newCtx, &HTTPWriteErrorInput{
							Conf: conf,
							W:    w, R: r, Err: err,
						})
						if werr != nil {
							slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err) // This should not happen generally
							return
						}
						written = output.Written
					}
					if !written {
						panic(v)
					}
				}
			}()
			next.ServeHTTP(w, r.WithContext(newCtx))
		})
	}
}
