package httperrors

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/qor5/x/v3/i18nx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func newTestI18NForHTTP(t *testing.T) *i18nx.I18N {
	t.Helper()
	csv := `key,en,zh
NOT_FOUND,Not Found,未找到
INTERNAL,Internal Error,内部错误
REQUIRED,Required,必填
INVALID_ARGUMENT,Invalid Argument,参数无效
`
	ib, err := i18nx.New(strings.NewReader(csv))
	require.NoError(t, err)
	return ib
}

func TestErrorMiddleware_PanicOnNilConfig(t *testing.T) {
	assert.Panics(t, func() {
		ErrorMiddleware(nil)
	})
}

func TestErrorMiddleware_PanicOnNilI18N(t *testing.T) {
	assert.Panics(t, func() {
		ErrorMiddleware(&HTTPErrorMiddlewareConfig{})
	})
}

func TestErrorMiddleware_PanicRecovery(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(Error(http.StatusNotFound, "NOT_FOUND", "user not found"))
	})

	middleware := NewErrorMiddleware(ib)
	server := middleware(handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Accept-Language", "en")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, "NOT_FOUND", resp.Code)
}

func TestErrorMiddleware_PanicRecovery_Chinese(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(Error(http.StatusNotFound, "NOT_FOUND", "user not found"))
	})

	middleware := NewErrorMiddleware(ib)
	server := middleware(handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("x-selected-language", "zh")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var resp ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	assert.Equal(t, "NOT_FOUND", resp.Code)
	// Original message preserved, translation in localizedMessage
	assert.Equal(t, "user not found", resp.Message)
	assert.Equal(t, "未找到", resp.LocalizedMessage)
}

func TestErrorMiddleware_NonErrorPanic(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("not an error")
	})

	middleware := NewErrorMiddleware(ib)
	server := middleware(handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	assert.Panics(t, func() {
		server.ServeHTTP(w, req)
	})
}

func TestErrorMiddleware_NoPanic(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	middleware := NewErrorMiddleware(ib)
	server := middleware(handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestErrorMiddleware_WithFieldViolations(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		panic(ValidationError(fv).Err())
	})

	middleware := NewErrorMiddleware(ib)
	server := middleware(handler)

	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	req.Header.Set("Accept-Language", "en")
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var resp ErrorResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)

	assert.Equal(t, "INVALID_ARGUMENT", resp.Code)
	require.Len(t, resp.FieldViolations, 1)
	assert.Equal(t, "email", resp.FieldViolations[0].Field)
	assert.Equal(t, "REQUIRED", resp.FieldViolations[0].Code)
	// Original description preserved, translation in localizedMessage
	assert.Equal(t, "email is required", resp.FieldViolations[0].Message)
	assert.Equal(t, "Required", resp.FieldViolations[0].LocalizedMessage)
}

func TestErrorMiddleware_WithHook(t *testing.T) {
	ib := newTestI18NForHTTP(t)

	hookCalled := false
	conf := &HTTPErrorMiddlewareConfig{I18N: ib}
	conf = conf.WithHTTPWriteErrorHook(func(next HTTPWriteErrorFunc) HTTPWriteErrorFunc {
		return func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
			hookCalled = true
			return next(ctx, input)
		}
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(Error(http.StatusBadRequest, "TEST", "test error"))
	})

	server := ErrorMiddleware(conf)(handler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	assert.True(t, hookCalled)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWriteJSONError(t *testing.T) {
	t.Run("simple error", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		w := httptest.NewRecorder()

		werr := WriteJSONError(err, w)
		require.NoError(t, werr)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "NOT_FOUND", resp.Code)
		assert.Equal(t, "user not found", resp.Message)
		assert.Empty(t, resp.FieldViolations)
	})

	t.Run("with metadata", func(t *testing.T) {
		err := New(http.StatusBadRequest, "TEST", "test").
			WithMetadata(map[string]string{"key": "value"}).Err()

		w := httptest.NewRecorder()
		werr := WriteJSONError(err, w)
		require.NoError(t, werr)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "value", resp.Metadata["key"])
	})

	t.Run("with field violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")
		err := ValidationError(fv1, fv2).Err()

		w := httptest.NewRecorder()
		werr := WriteJSONError(err, w)
		require.NoError(t, werr)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		require.Len(t, resp.FieldViolations, 2)
		assert.Equal(t, "email", resp.FieldViolations[0].Field)
		assert.Equal(t, "REQUIRED", resp.FieldViolations[0].Code)
		assert.Equal(t, "name", resp.FieldViolations[1].Field)
		assert.Equal(t, "TOO_SHORT", resp.FieldViolations[1].Code)
	})

	t.Run("with translated field violation", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		// Translate to Chinese
		err = TranslateError(err, ib, language.Chinese)

		w := httptest.NewRecorder()
		werr := WriteJSONError(err, w)
		require.NoError(t, werr)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		require.Len(t, resp.FieldViolations, 1)
		assert.Equal(t, "email is required", resp.FieldViolations[0].Message)
		assert.Equal(t, "必填", resp.FieldViolations[0].LocalizedMessage)
	})

	t.Run("no redundant status code in body", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
		w := httptest.NewRecorder()
		WriteJSONError(err, w)

		// Verify raw JSON does not contain a "status" or "statusCode" field
		body := w.Body.String()
		assert.NotContains(t, body, `"status"`)
		assert.NotContains(t, body, `"statusCode"`)
		assert.NotContains(t, body, `"httpStatus"`)
	})

	t.Run("camelCase field names", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "required")
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithMetadata(map[string]string{"k": "v"}).
			WithFieldViolations(fv).Err()

		w := httptest.NewRecorder()
		WriteJSONError(err, w)

		body := w.Body.String()
		// Verify camelCase
		assert.Contains(t, body, `"code"`)
		assert.Contains(t, body, `"message"`)
		assert.Contains(t, body, `"fieldViolations"`)
		assert.Contains(t, body, `"metadata"`)
		// Verify no snake_case
		assert.NotContains(t, body, `"field_violations"`)
		// localizedMessage should not appear when untranslated (omitempty)
		assert.NotContains(t, body, `"localizedMessage"`)
	})
}

func TestHandleError(t *testing.T) {
	t.Run("simple error", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Accept-Language", "en")
		w := httptest.NewRecorder()

		err := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		HandleError(conf, w, req, err)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "NOT_FOUND", resp.Code)
		assert.Equal(t, "user not found", resp.Message)
		assert.Equal(t, "Not Found", resp.LocalizedMessage)
	})

	t.Run("Chinese translation", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("x-selected-language", "zh")
		w := httptest.NewRecorder()

		err := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		HandleError(conf, w, req, err)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "user not found", resp.Message)
		assert.Equal(t, "未找到", resp.LocalizedMessage)
	})

	t.Run("with field violations", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		req := httptest.NewRequest(http.MethodPost, "/test", nil)
		req.Header.Set("Accept-Language", "en")
		w := httptest.NewRecorder()

		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		err := ValidationError(fv).Err()
		HandleError(conf, w, req, err)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		require.Len(t, resp.FieldViolations, 1)
		assert.Equal(t, "email is required", resp.FieldViolations[0].Message)
		assert.Equal(t, "Required", resp.FieldViolations[0].LocalizedMessage)
	})

	t.Run("with hook", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		hookCalled := false
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}
		conf = conf.WithHTTPWriteErrorHook(func(next HTTPWriteErrorFunc) HTTPWriteErrorFunc {
			return func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
				hookCalled = true
				return next(ctx, input)
			}
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		HandleError(conf, w, req, Error(http.StatusBadRequest, "TEST", "test"))

		assert.True(t, hookCalled)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("panic on nil config", func(t *testing.T) {
		assert.Panics(t, func() {
			HandleError(nil, httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil), Error(http.StatusBadRequest, "T", "t"))
		})
	})

	t.Run("panic on nil I18N", func(t *testing.T) {
		assert.Panics(t, func() {
			HandleError(&HTTPErrorMiddlewareConfig{}, httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/", nil), Error(http.StatusBadRequest, "T", "t"))
		})
	})
}

func TestWrapHandlerFunc(t *testing.T) {
	t.Run("panic recovery", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		handler := func(w http.ResponseWriter, r *http.Request) {
			panic(Error(http.StatusNotFound, "NOT_FOUND", "not found"))
		}

		wrapped := WrapHandlerFunc(conf, handler)

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("x-selected-language", "zh")
		w := httptest.NewRecorder()

		wrapped.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		assert.Equal(t, "NOT_FOUND", resp.Code)
		assert.Equal(t, "not found", resp.Message)
		assert.Equal(t, "未找到", resp.LocalizedMessage)
	})

	t.Run("no panic passes through", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		handler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}

		wrapped := WrapHandlerFunc(conf, handler)

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		wrapped.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "ok", w.Body.String())
	})

	t.Run("non-error panic re-panics", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		handler := func(w http.ResponseWriter, r *http.Request) {
			panic("not an error")
		}

		wrapped := WrapHandlerFunc(conf, handler)

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()

		assert.Panics(t, func() {
			wrapped.ServeHTTP(w, req)
		})
	})

	t.Run("mixed mux - only wrapped handlers use httperrors", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}

		mux := http.NewServeMux()

		// Wrapped handler uses httperrors
		mux.HandleFunc("GET /api/users", WrapHandlerFunc(conf, func(w http.ResponseWriter, r *http.Request) {
			panic(Error(http.StatusNotFound, "NOT_FOUND", "user not found"))
		}))

		// Legacy handler does NOT use httperrors
		mux.HandleFunc("GET /legacy", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("legacy ok"))
		})

		// Test wrapped handler
		req1 := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		req1.Header.Set("Accept-Language", "en")
		w1 := httptest.NewRecorder()
		mux.ServeHTTP(w1, req1)
		assert.Equal(t, http.StatusNotFound, w1.Code)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w1.Body.Bytes(), &resp))
		assert.Equal(t, "NOT_FOUND", resp.Code)

		// Test legacy handler
		req2 := httptest.NewRequest(http.MethodGet, "/legacy", nil)
		w2 := httptest.NewRecorder()
		mux.ServeHTTP(w2, req2)
		assert.Equal(t, http.StatusOK, w2.Code)
		assert.Equal(t, "legacy ok", w2.Body.String())
	})

	t.Run("with hook", func(t *testing.T) {
		ib := newTestI18NForHTTP(t)
		hookCalled := false
		conf := &HTTPErrorMiddlewareConfig{I18N: ib}
		conf = conf.WithHTTPWriteErrorHook(func(next HTTPWriteErrorFunc) HTTPWriteErrorFunc {
			return func(ctx context.Context, input *HTTPWriteErrorInput) (*HTTPWriteErrorOutput, error) {
				hookCalled = true
				return next(ctx, input)
			}
		})

		wrapped := WrapHandlerFunc(conf, func(w http.ResponseWriter, r *http.Request) {
			panic(Error(http.StatusBadRequest, "TEST", "test"))
		})

		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		wrapped.ServeHTTP(w, req)

		assert.True(t, hookCalled)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestErrorResponse_OmitEmpty(t *testing.T) {
	err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
	w := httptest.NewRecorder()
	WriteJSONError(err, w)

	var raw map[string]any
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &raw))

	// These should be omitted when empty
	_, hasLocalizedMessage := raw["localizedMessage"]
	_, hasMetadata := raw["metadata"]
	_, hasFieldViolations := raw["fieldViolations"]
	assert.False(t, hasLocalizedMessage)
	assert.False(t, hasMetadata)
	assert.False(t, hasFieldViolations)
}
