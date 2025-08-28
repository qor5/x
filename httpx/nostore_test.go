package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoStoreMiddleware(t *testing.T) {
	tests := []struct {
		name          string
		protoMajor    int
		protoMinor    int
		wantCacheCtrl string
		wantPragma    string
		wantExpires   string
	}{
		{
			name:          "HTTP/1.1",
			protoMajor:    1,
			protoMinor:    1,
			wantCacheCtrl: "no-store, no-cache, must-revalidate, max-age=0",
			wantPragma:    "",
			wantExpires:   "",
		},
		{
			name:          "HTTP/1.0",
			protoMajor:    1,
			protoMinor:    0,
			wantCacheCtrl: "no-store, no-cache, must-revalidate, max-age=0",
			wantPragma:    "no-cache",
			wantExpires:   "0",
		},
		{
			name:          "HTTP/2.0",
			protoMajor:    2,
			protoMinor:    0,
			wantCacheCtrl: "no-store, no-cache, must-revalidate, max-age=0",
			wantPragma:    "",
			wantExpires:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			nextHandler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {})
			handler := NoStore(nextHandler)

			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req.ProtoMajor = tt.protoMajor
			req.ProtoMinor = tt.protoMinor
			rec := httptest.NewRecorder()

			// Execute
			handler.ServeHTTP(rec, req)

			// Verify headers
			headers := rec.Header()
			assert.Equal(t, tt.wantCacheCtrl, headers.Get("Cache-Control"), "Cache-Control header")
			assert.Equal(t, tt.wantPragma, headers.Get("Pragma"), "Pragma header")
			assert.Equal(t, tt.wantExpires, headers.Get("Expires"), "Expires header")
		})
	}
}

func TestNoStoreNextHandler(t *testing.T) {
	// Setup
	nextCalled := false
	nextHandler := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		nextCalled = true
	})

	handler := NoStore(nextHandler)
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	rec := httptest.NewRecorder()

	// Execute
	handler.ServeHTTP(rec, req)

	// Verify next handler was called
	assert.True(t, nextCalled, "NoStore middleware should call the next handler")
}
