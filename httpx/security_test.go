package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSecurity verifies that the Security middleware correctly applies
// security headers and content type validation based on configuration.
func TestSecurity(t *testing.T) {
	tests := []struct {
		name               string
		config             SecurityConfig
		request            *http.Request
		expectedStatus     int
		expectedHeaders    map[string]string
		expectedHeadersSet []string
		notExpectedHeaders []string
	}{
		{
			name: "default_security_config",
			config: SecurityConfig{
				DenyMIMETypeSniffing: true,
				DenyClickjacking:     true,
				EnableHSTS:           true,
				CORS:                 CORSConfig{},
			},
			request:        httptest.NewRequest(http.MethodGet, "/test", nil),
			expectedStatus: http.StatusOK,
			expectedHeaders: map[string]string{
				"X-Content-Type-Options":    "nosniff",
				"Content-Security-Policy":   "frame-ancestors 'self'",
				"X-Frame-Options":           "SAMEORIGIN",
				"Strict-Transport-Security": "max-age=31536000; includeSubDomains; preload",
			},
		},
		{
			name: "without_security_measures",
			config: SecurityConfig{
				DenyMIMETypeSniffing: false,
				DenyClickjacking:     false,
				EnableHSTS:           false,
				CORS:                 CORSConfig{},
			},
			request:        httptest.NewRequest(http.MethodGet, "/test", nil),
			expectedStatus: http.StatusOK,
			notExpectedHeaders: []string{
				"X-Content-Type-Options",
				"Content-Security-Policy",
				"X-Frame-Options",
				"Strict-Transport-Security",
			},
		},
		{
			name: "with_allowed_origins",
			config: SecurityConfig{
				CORS: CORSConfig{
					AllowedOrigins: []string{"https://example.com", "https://test.com"},
				},
				DenyClickjacking: true,
			},
			request:        httptest.NewRequest(http.MethodGet, "/test", nil),
			expectedStatus: http.StatusOK,
			expectedHeaders: map[string]string{
				"Content-Security-Policy": "frame-ancestors 'self' https://example.com https://test.com;",
				"X-Frame-Options":         "SAMEORIGIN",
			},
		},
		{
			name: "deny_simple_requests_enabled",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Set("X-Requested-By", "fetch")
				return req
			}(),
			expectedStatus: http.StatusOK,
		},
		{
			name: "deny_simple_requests_with_skip_check",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
					SkipDenySimpleRequests: func(r *http.Request) bool {
						return r.URL.Path == "/healthz"
					},
				},
			},
			request:        httptest.NewRequest(http.MethodGet, "/healthz", nil),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			middleware := Security(tt.config)
			secureHandler := middleware(handler)

			w := httptest.NewRecorder()
			secureHandler.ServeHTTP(w, tt.request)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode, "Status code should match expected")

			// Check expected headers
			for key, value := range tt.expectedHeaders {
				assert.Equal(t, value, resp.Header.Get(key), "Header %s should match expected", key)
			}

			// Check headers presence
			for _, key := range tt.expectedHeadersSet {
				assert.NotEmpty(t, resp.Header.Get(key), "Header %s should be set", key)
			}

			// Check headers absence
			for _, key := range tt.notExpectedHeaders {
				assert.Empty(t, resp.Header.Get(key), "Header %s should not be set", key)
			}
		})
	}
}

func TestDenySimpleRequests(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		contentType    string
		headerValue    string
		expectedStatus int
	}{
		{
			name:           "post_no_content_type_no_header",
			method:         http.MethodPost,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "post_no_content_type_with_header",
			method:         http.MethodPost,
			contentType:    "",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "post_simple_content_type_no_header",
			method:         http.MethodPost,
			contentType:    "application/x-www-form-urlencoded",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "post_simple_content_type_with_header",
			method:         http.MethodPost,
			contentType:    "application/x-www-form-urlencoded",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "post_multipart_no_header",
			method:         http.MethodPost,
			contentType:    "multipart/form-data",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "post_multipart_with_header",
			method:         http.MethodPost,
			contentType:    "multipart/form-data",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "post_text_plain_no_header",
			method:         http.MethodPost,
			contentType:    "text/plain; charset=utf-8",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "post_text_plain_with_header",
			method:         http.MethodPost,
			contentType:    "text/plain; charset=utf-8",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "post_json_no_header_allowed",
			method:         http.MethodPost,
			contentType:    "application/json",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "post_proto_no_header_allowed",
			method:         http.MethodPost,
			contentType:    "application/proto",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "get_no_header",
			method:         http.MethodGet,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "get_with_header",
			method:         http.MethodGet,
			contentType:    "",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "put_no_header_allowed",
			method:         http.MethodPut,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "delete_no_header_allowed",
			method:         http.MethodDelete,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "patch_no_header_allowed",
			method:         http.MethodPatch,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "options_no_header_allowed",
			method:         http.MethodOptions,
			contentType:    "",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			middleware := DenySimpleRequests(handler)

			req := httptest.NewRequest(tt.method, "/test", nil)
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			if tt.headerValue != "" {
				req.Header.Set("X-Requested-By", tt.headerValue)
			}

			w := httptest.NewRecorder()
			middleware.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

func TestDenySimpleRequestsFactory(t *testing.T) {
	tests := []struct {
		name           string
		skipCheck      func(r *http.Request) bool
		requestPath    string
		headerValue    string
		expectedStatus int
	}{
		{
			name:           "nil_skipCheck_missing_header",
			skipCheck:      nil,
			requestPath:    "/api/test",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "nil_skipCheck_with_header",
			skipCheck:      nil,
			requestPath:    "/api/test",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name: "skipCheck_returns_true",
			skipCheck: func(r *http.Request) bool {
				return r.URL.Path == "/healthz"
			},
			requestPath:    "/healthz",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
		{
			name: "skipCheck_returns_false_missing_header",
			skipCheck: func(r *http.Request) bool {
				return r.URL.Path == "/healthz"
			},
			requestPath:    "/api/test",
			headerValue:    "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "skipCheck_returns_false_with_header",
			skipCheck: func(r *http.Request) bool {
				return r.URL.Path == "/healthz"
			},
			requestPath:    "/api/test",
			headerValue:    "fetch",
			expectedStatus: http.StatusOK,
		},
		{
			name: "skipCheck_multiple_paths",
			skipCheck: func(r *http.Request) bool {
				return r.URL.Path == "/healthz" || r.URL.Path == "/metrics"
			},
			requestPath:    "/metrics",
			headerValue:    "",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			middleware := DenySimpleRequestsFactory(tt.skipCheck)(handler)

			req := httptest.NewRequest(http.MethodPost, tt.requestPath, nil)
			if tt.headerValue != "" {
				req.Header.Set("X-Requested-By", tt.headerValue)
			}

			w := httptest.NewRecorder()
			middleware.ServeHTTP(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}

// TestBuildFrameAncestors verifies that the frame ancestors header value
// is correctly constructed based on the provided origins.
func TestBuildFrameAncestors(t *testing.T) {
	tests := []struct {
		name     string
		origins  []string
		expected string
	}{
		{
			name:     "empty_origins",
			origins:  []string{},
			expected: "frame-ancestors 'self'",
		},
		{
			name:     "single_origin",
			origins:  []string{"https://example.com"},
			expected: "frame-ancestors 'self' https://example.com;",
		},
		{
			name:     "multiple_origins",
			origins:  []string{"https://example.com", "https://test.com"},
			expected: "frame-ancestors 'self' https://example.com https://test.com;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildFrameAncestors(tt.origins)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestParseContentType verifies that content type parsing correctly
// handles various scenarios including missing, multiple, and malformed headers.
func TestParseContentType(t *testing.T) {
	tests := []struct {
		name           string
		contentType    string
		expected       string
		expectedParams map[string]string
		expectedErrMsg string
	}{
		{
			name:           "missing_content_type",
			contentType:    "",
			expectedErrMsg: "Content-Type header not found",
		},
		{
			name:           "multiple_content_types",
			contentType:    "multiple",
			expectedErrMsg: "multiple Content-Type headers found",
		},
		{
			name:           "invalid_content_type",
			contentType:    "invalid",
			expectedErrMsg: "failed to parse Content-Type",
		},
		{
			name:        "valid_content_type",
			contentType: "application/json",
			expected:    "application/json",
		},
		{
			name:        "content_type_with_charset",
			contentType: "application/json; charset=utf-8",
			expected:    "application/json",
			expectedParams: map[string]string{
				"charset": "utf-8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tt.contentType != "" {
				switch tt.contentType {
				case "multiple":
					req.Header.Add("Content-Type", "application/json")
					req.Header.Add("Content-Type", "text/plain")
				case "invalid":
					req.Header.Set("Content-Type", "invalid content type")
				default:
					req.Header.Set("Content-Type", tt.contentType)
				}
			}

			result, params, err := ParseContentType(req)

			if tt.expectedErrMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
				if tt.expectedParams != nil {
					assert.Equal(t, tt.expectedParams, params)
				}
			}
		})
	}
}
