package httperrors

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/qor5/x/v3/hook"
	"github.com/qor5/x/v3/i18nx"
	"golang.org/x/text/language"
)

type (
	HTTPWriteErrorInput struct {
		Conf *HTTPErrorMiddlewareConfig
		W    http.ResponseWriter
		R    *http.Request
		Err  error
	}
	HTTPWriteErrorOutput struct {
		Written bool
	}
	HTTPWriteErrorFunc func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error)
)

type HTTPErrorMiddlewareConfig struct {
	I18N           *i18nx.I18N
	writeErrorHook func(next HTTPWriteErrorFunc) HTTPWriteErrorFunc
}

func (c *HTTPErrorMiddlewareConfig) WithHTTPWriteErrorHook(hooks ...hook.Hook[HTTPWriteErrorFunc]) *HTTPErrorMiddlewareConfig {
	c.writeErrorHook = hook.Prepend(c.writeErrorHook, hooks...)
	return c
}

// ErrorResponse is the standard JSON error response body.
// Fields use camelCase for frontend compatibility.
type ErrorResponse struct {
	Code             string                         `json:"code"`
	Message          string                         `json:"message"`
	LocalizedMessage string                         `json:"localizedMessage,omitempty"`
	Metadata         map[string]string              `json:"metadata,omitempty"`
	FieldViolations  []*ErrorResponseFieldViolation `json:"fieldViolations,omitempty"`
}

// ErrorResponseFieldViolation represents a single field violation in the JSON error response.
type ErrorResponseFieldViolation struct {
	Field            string `json:"field"`
	Code             string `json:"code"`
	Message          string `json:"message"`
	LocalizedMessage string `json:"localizedMessage,omitempty"`
}

// ErrorMiddleware creates an HTTP middleware that:
// 1. Injects i18n context into the request
// 2. Recovers from panics containing errors
// 3. Translates and writes structured JSON error responses
//
// Currently uses panic-based error propagation to keep handler signatures as standard http.HandlerFunc.
// NOTE: In the future, this may be changed to support a handler signature that returns error directly
// (e.g., type HandlerFunc func(w http.ResponseWriter, r *http.Request) error).
// The panic-based approach is chosen for now to maintain compatibility with standard http.HandlerFunc.
func ErrorMiddleware(conf *HTTPErrorMiddlewareConfig) func(http.Handler) http.Handler {
	if conf == nil || conf.I18N == nil {
		panic("HTTPErrorMiddlewareConfig.I18N is required")
	}
	return func(next http.Handler) http.Handler {
		defWriteError := func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
			lang := languageFromRequest(input.Conf.I18N, input.R)
			err := TranslateError(input.Err, input.Conf.I18N, lang)
			werr := WriteJSONError(err, input.W)
			if werr != nil {
				return &HTTPWriteErrorOutput{Written: false}, werr
			}
			return &HTTPWriteErrorOutput{Written: true}, nil
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
							slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err)
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

// NewErrorMiddleware is a convenience function that creates an ErrorMiddleware with default configuration.
func NewErrorMiddleware(ib *i18nx.I18N) func(http.Handler) http.Handler {
	return ErrorMiddleware(&HTTPErrorMiddlewareConfig{I18N: ib})
}

// languageFromRequest extracts the preferred language from HTTP request headers.
// It checks x-selected-language first, then falls back to accept-language.
func languageFromRequest(ib *i18nx.I18N, r *http.Request) language.Tag {
	selected := r.Header.Get(i18nx.HeaderSelectedLanguage)
	accept := r.Header.Get(i18nx.HeaderAcceptLanguage)
	return ib.MatchStrings(selected, accept)
}

// WriteJSONError writes a structured JSON error response from an error.
// The HTTP status code is set from the Status object.
// The response body follows the ErrorResponse format with camelCase fields.
func WriteJSONError(err error, w http.ResponseWriter) error {
	st := Convert(err)

	resp := &ErrorResponse{
		Code:     st.Reason(),
		Message:  st.Message(),
		Metadata: st.Metadata(),
	}
	if st.localizedMessage != nil {
		resp.LocalizedMessage = st.localizedMessage.Message
	}

	for _, fv := range st.fieldViolations {
		efv := &ErrorResponseFieldViolation{
			Field:   fv.Field(),
			Code:    fv.Reason(),
			Message: fv.Description(),
		}
		if fv.localizedMessage != nil {
			efv.LocalizedMessage = fv.localizedMessage.Message
		}
		resp.FieldViolations = append(resp.FieldViolations, efv)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(st.StatusCode())

	data, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		return marshalErr
	}

	_, writeErr := w.Write(data)
	return writeErr
}

// HandleError translates and writes a structured JSON error response for the given error.
// This is intended for use inside individual handlers that want to explicitly handle errors
// without relying on panic-based middleware.
//
// Usage:
//
//	func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
//	    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
//	    if err != nil {
//	        httperrors.HandleError(h.conf, w, r, err)
//	        return
//	    }
//	    json.NewEncoder(w).Encode(user)
//	}
func HandleError(conf *HTTPErrorMiddlewareConfig, w http.ResponseWriter, r *http.Request, err error) {
	if conf == nil || conf.I18N == nil {
		panic("HTTPErrorMiddlewareConfig.I18N is required")
	}

	writeError := func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
		lang := languageFromRequest(input.Conf.I18N, input.R)
		translated := TranslateError(input.Err, input.Conf.I18N, lang)
		werr := WriteJSONError(translated, input.W)
		if werr != nil {
			return &HTTPWriteErrorOutput{Written: false}, werr
		}
		return &HTTPWriteErrorOutput{Written: true}, nil
	}
	if conf.writeErrorHook != nil {
		writeError = conf.writeErrorHook(writeError)
	}

	_, werr := writeError(r.Context(), &HTTPWriteErrorInput{
		Conf: conf,
		W:    w, R: r, Err: err,
	})
	if werr != nil {
		slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err)
	}
}

// WrapHandlerFunc wraps a single http.HandlerFunc with httperrors panic recovery and i18n translation.
// This is useful when only some handlers in a mux use httperrors, and a global middleware is not appropriate.
//
// Usage:
//
//	mux.HandleFunc("/api/users/{id}", httperrors.WrapHandlerFunc(conf, h.GetUser))
//	mux.HandleFunc("/legacy/other", legacyHandler) // not wrapped
func WrapHandlerFunc(conf *HTTPErrorMiddlewareConfig, handler http.HandlerFunc) http.HandlerFunc {
	wrapped := ErrorMiddleware(conf)(handler)
	return wrapped.ServeHTTP
}
