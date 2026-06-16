package statusx

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/qor5/x/v3/jsonx"
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

func TestWriteVProtoHTTPError_MetadataHeader(t *testing.T) {
	md := map[string]string{"order_id": "123", "retry_after": "30s"}
	tests := []struct {
		name         string
		err          error
		wantMetadata map[string]string // nil means no metadata header expected
	}{
		{
			name:         "with metadata sets header",
			err:          New(codes.InvalidArgument, "INVALID_ARGUMENT", "invalid").WithMetadata(md).Err(),
			wantMetadata: md,
		},
		{
			name:         "without metadata omits header",
			err:          BadRequest(NewFieldViolation("email", "REQUIRED", "Email is required")).Err(),
			wantMetadata: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", strings.NewReader("req"))
			req.Header.Set("Accept", "application/json")
			w := httptest.NewRecorder()

			werr := WriteVProtoHTTPError(tt.err, w, req)
			require.NoError(t, werr)

			encoded := w.Header().Get(HeaderStatusMetadata)
			if tt.wantMetadata == nil {
				assert.Empty(t, encoded)
				return
			}

			require.NotEmpty(t, encoded)

			// Reverse the encoding done by WriteVProtoHTTPError: URL-unescape
			// then JSON-unmarshal the header value back into a metadata map.
			data, derr := url.QueryUnescape(encoded)
			require.NoError(t, derr)
			var decoded map[string]string
			require.NoError(t, jsonx.Unmarshal([]byte(data), &decoded))
			assert.Equal(t, tt.wantMetadata, decoded)
		})
	}
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
