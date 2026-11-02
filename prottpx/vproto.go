package prottpx

import (
	"context"
	"log/slog"

	"github.com/qor5/x/v3/hook"
	"github.com/qor5/x/v3/i18nx"
	"github.com/qor5/x/v3/statusx"
)

// NewVProtoWriteErrorHook returns a prottpx.WriteErrorFunc hook that writes VProto-compatible error responses.
// It translates statusx errors and writes them in VProto format for compatibility with frontend clients.
func NewVProtoWriteErrorHook(ib *i18nx.I18N) hook.Hook[WriteErrorFunc] {
	return func(next WriteErrorFunc) WriteErrorFunc {
		return func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
			lang := ib.LanguageFromContext(ctx)

			// Translate statusx errors
			err, translated := statusx.TranslateStatusErrorOnly(input.Error, ib, lang)

			// If x-ensure-connect-error header is set, use connect error writer
			if statusx.EnsureConnectError(ctx) {
				written := statusx.WriteConnectErrorOnly(input.ConnectErrorWriter, input.W, input.R, err)
				return &WriteErrorOutput{Written: written}, nil
			}

			// For non-StatusError types, delegate to the default handler
			if !translated {
				return next(ctx, input)
			}

			// Write VProto error response
			werr := statusx.WriteVProtoHTTPError(err, input.W, input.R)
			if werr != nil {
				slog.Error("Failed to write vproto http error", "error", werr)
			}

			return &WriteErrorOutput{Written: true}, nil
		}
	}
}
