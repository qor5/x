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
			name: "deny_simple_requests_with_invalid_content_type",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Set("Content-Type", "text/plain")
				return req
			}(),
			expectedStatus: http.StatusUnsupportedMediaType,
		},
		{
			name: "deny_simple_requests_with_missing_content_type",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request:        httptest.NewRequest(http.MethodPost, "/test", nil),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "deny_simple_requests_with_valid_json_content_type",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			expectedStatus: http.StatusOK,
		},
		{
			name: "deny_simple_requests_with_valid_proto_content_type",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Set("Content-Type", "application/proto")
				return req
			}(),
			expectedStatus: http.StatusOK,
		},
		{
			name: "multiple_content_type_headers",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Content-Type", "text/plain")
				return req
			}(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid_content_type_format",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/test", nil)
				req.Header.Set("Content-Type", "invalid content type")
				return req
			}(),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "deny_simple_requests_with_head_method",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodHead, "/test", nil)
				req.Header.Set("Content-Type", "text/plain")
				return req
			}(),
			expectedStatus: http.StatusUnsupportedMediaType,
		},
		{
			name: "deny_simple_requests_with_valid_head_content_type",
			config: SecurityConfig{
				CORS: CORSConfig{
					DenySimpleRequests: true,
				},
			},
			request: func() *http.Request {
				req := httptest.NewRequest(http.MethodHead, "/test", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
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
