package httpx

import (
	"time"
)

type ServerConfig struct {
	Address           string         `confx:"address" usage:"HTTP server address" validate:"required"`
	PathPrefix        string         `confx:"pathPrefix" usage:"Path prefix for all handlers. Will be normalized to start with '/' and not end with '/' (except for root path '/'). Root path '/' is treated as no prefix. Example: 'api/v1' or '/api/v1/' both become '/api/v1'"`
	ReadTimeout       time.Duration  `confx:"readTimeout" usage:"maximum duration before timing out read of the request"`
	ReadHeaderTimeout time.Duration  `confx:"readHeaderTimeout" usage:"maximum duration before timing out read of the request headers" validate:"ltefield=ReadTimeout"`
	WriteTimeout      time.Duration  `confx:"writeTimeout" usage:"maximum duration before timing out write of the response"`
	IdleTimeout       time.Duration  `confx:"idleTimeout" usage:"maximum amount of time to wait for the next request when keep-alives are enabled"`
	TLS               TLSConfig      `confx:"tls"`
	Security          SecurityConfig `confx:",squash"`
}

type TLSConfig struct {
	Enabled    bool   `confx:"enabled" usage:"Enable TLS"`
	CertBase64 string `confx:"certBase64" usage:"TLS certificate base64 encoded" validate:"required_if=Enabled true"`
	KeyBase64  string `confx:"keyBase64" usage:"TLS key base64 encoded" validate:"required_if=Enabled true"`
}

type CORSConfig struct {
	Debug              bool          `confx:"debug" usage:"CORS debug"`
	AllowedOrigins     []string      `confx:"allowedOrigins" usage:"CORS allowed origins" validate:"dive,http_url"`
	AllowedMethods     []string      `confx:"allowedMethods" usage:"CORS allowed methods, POST is always allowed" validate:"dive,oneof=GET HEAD POST PUT PATCH DELETE CONNECT OPTIONS TRACE"`
	AllowedHeaders     []string      `confx:"allowedHeaders" usage:"CORS allowed headers, Content-Type is always allowed"`
	ExposedHeaders     []string      `confx:"exposedHeaders" usage:"CORS exposed headers"`
	MaxAge             time.Duration `confx:"maxAge" usage:"CORS max age"`
	DenySimpleRequests bool          `confx:"denySimpleRequests" usage:"CORS Deny simple requests"`
}

type SecurityConfig struct {
	CORS                 CORSConfig `confx:"cors"`
	DenyMIMETypeSniffing bool       `confx:"denyMIMETypeSniffing" usage:"Deny MIME type sniffing"`
	DenyClickjacking     bool       `confx:"denyClickjacking" usage:"Deny clickjacking"`
	EnableHSTS           bool       `confx:"enableHSTS" usage:"Enable HSTS"`
}
