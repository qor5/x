package httpx

import (
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
