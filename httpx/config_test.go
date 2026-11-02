package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qor5/confx"
)

func TestCORSConfig_Validation(t *testing.T) {
	suite := confx.NewValidationSuite(t)

	suite.RunTests([]confx.ExpectedValidation{
		{
			Name: "valid config",
			Config: &CORSConfig{
				Debug:          false,
				AllowedOrigins: []string{"http://localhost:3000", "https://example.com"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
				AllowedHeaders: []string{"Content-Type", "Authorization"},
				ExposedHeaders: []string{"X-Request-ID"},
				MaxAge:         24 * time.Hour,
			},
		},
		{
			Name: "invalid origins",
			Config: &CORSConfig{
				AllowedOrigins: []string{"invalid-url", "localhost:3000"},
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "CORSConfig.AllowedOrigins[0]", Tag: "http_url"},
				{Path: "CORSConfig.AllowedOrigins[1]", Tag: "http_url"},
			},
		},
		{
			Name: "invalid methods",
			Config: &CORSConfig{
				AllowedMethods: []string{"INVALID", "POST", "UNKNOWN"},
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "CORSConfig.AllowedMethods[0]", Tag: "oneof"},
				{Path: "CORSConfig.AllowedMethods[2]", Tag: "oneof"},
			},
		},
	})
}

func TestSecurityConfig_Validation(t *testing.T) {
	suite := confx.NewValidationSuite(t)

	suite.RunTests([]confx.ExpectedValidation{
		{
			Name: "valid config with all security features enabled",
			Config: &SecurityConfig{
				CORS: CORSConfig{
					AllowedOrigins: []string{"http://localhost:3000"},
					AllowedMethods: []string{"GET", "POST"},
				},
				DenyMIMETypeSniffing: true,
				DenyClickjacking:     true,
				EnableHSTS:           true,
			},
		},
		{
			Name: "valid config with all security features disabled",
			Config: &SecurityConfig{
				CORS: CORSConfig{
					AllowedOrigins: []string{"http://localhost:3000"},
					AllowedMethods: []string{"GET", "POST"},
				},
				DenyMIMETypeSniffing: false,
				DenyClickjacking:     false,
				EnableHSTS:           false,
			},
		},
		{
			Name: "invalid CORS config",
			Config: &SecurityConfig{
				CORS: CORSConfig{
					AllowedOrigins: []string{"invalid-url"},
					AllowedMethods: []string{"INVALID"},
				},
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "SecurityConfig.CORS.AllowedOrigins[0]", Tag: "http_url"},
				{Path: "SecurityConfig.CORS.AllowedMethods[0]", Tag: "oneof"},
			},
		},
	})
}

func TestServerConfig_PathPrefix(t *testing.T) {
	// Create a test handler that returns the requested path
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(r.URL.Path))
	})

	tests := []struct {
		name           string
		pathPrefix     string
		requestPath    string
		expectedPath   string
		expectedStatus int
	}{
		{
			name:           "no path prefix",
			pathPrefix:     "",
			requestPath:    "/api/users",
			expectedPath:   "/api/users",
			expectedStatus: 200,
		},
		{
			name:           "with path prefix",
			pathPrefix:     "/api/v1",
			requestPath:    "/api/v1/users",
			expectedPath:   "/users",
			expectedStatus: 200,
		},
		{
			name:           "path prefix not matched",
			pathPrefix:     "/api/v1",
			requestPath:    "/api/users",
			expectedPath:   "404 page not found\n", // StripPrefix returns 404 when prefix doesn't match
			expectedStatus: 404,
		},
		{
			name:           "prefix normalization test",
			pathPrefix:     "api/v1/", // Will be normalized to "/api/v1"
			requestPath:    "/api/v1/users",
			expectedPath:   "/users",
			expectedStatus: 200,
		},
		{
			name:           "root path prefix should not strip",
			pathPrefix:     "/", // Root path should not strip anything
			requestPath:    "/users",
			expectedPath:   "/users",
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ServerConfig{
				Address:    ":8080",
				PathPrefix: tt.pathPrefix,
			}

			// Create server with the test handler
			srv, err := NewServer(conf, testHandler)
			if err != nil {
				t.Fatalf("Failed to create server: %v", err)
			}

			// Create a request to test
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			w := httptest.NewRecorder()

			// Serve the request
			srv.Handler.ServeHTTP(w, req)

			// Check the response
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			responseBody := w.Body.String()
			if responseBody != tt.expectedPath {
				t.Errorf("Expected path %q, got %q", tt.expectedPath, responseBody)
			}
		})
	}
}
