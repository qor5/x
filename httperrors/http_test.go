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
	// Message should be translated to Chinese
	assert.Equal(t, "未找到", resp.Message)
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
		fv := NewFieldViolation("email", "REQUIRED", "required")
		fv.localizedMessage = &LocalizedMessage{Locale: "zh", Message: "必填"}
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		w := httptest.NewRecorder()
		werr := WriteJSONError(err, w)
		require.NoError(t, werr)

		var resp ErrorResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
		require.Len(t, resp.FieldViolations, 1)
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
		fv.localizedMessage = &LocalizedMessage{Locale: "en", Message: "Required"}
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
		assert.Contains(t, body, `"localizedMessage"`)
		assert.Contains(t, body, `"metadata"`)
		// Verify no snake_case
		assert.NotContains(t, body, `"field_violations"`)
		assert.NotContains(t, body, `"localized_message"`)
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
