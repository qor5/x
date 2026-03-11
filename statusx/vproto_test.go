package statusx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestWriteVProtoHTTPError_JSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", strings.NewReader("req"))
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()
	err := BadRequest(NewFieldViolation("email", "REQUIRED", "Email is required")).Err()

	werr := WriteVProtoHTTPError(err, w, req)
	require.NoError(t, werr)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	assert.Contains(t, w.Body.String(), "INVALID_ARGUMENT")
}

func TestWriteVProtoHTTPError_MarshalFailureFallbackJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", strings.NewReader("req"))
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	st := New(codes.InvalidArgument, "INVALID_ARGUMENT", "invalid")
	st.message = string([]byte{0xff})

	werr := WriteVProtoHTTPError(st.Err(), w, req)
	require.Error(t, werr)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Empty(t, w.Body.String())
}
