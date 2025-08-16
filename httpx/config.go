package httpx

import "time"

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
