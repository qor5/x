package httpx

import (
	"net/http"
	"time"
)

type ServerConfig struct {
	Address     string        `confx:"address" usage:"HTTP server address" validate:"required"`
	PathPrefix  string        `confx:"pathPrefix" usage:"Path prefix for all handlers. Will be normalized to start with '/' and not end with '/' (except for root path '/'). Root path '/' is treated as no prefix. Example: 'api/v1' or '/api/v1/' both become '/api/v1'"`
	ReadTimeout time.Duration `confx:"readTimeout" usage:"maximum duration before timing out read of the request"`
	// stop_if guards the ltefield: ReadTimeout == 0 means no read deadline at
	// all, so it is not an upper bound, and a plain `ltefield` would reject a
	// config that sets only a header timeout.
	ReadHeaderTimeout time.Duration `confx:"readHeaderTimeout" usage:"maximum duration before timing out read of the request headers" validate:"stop_if=ReadTimeout 0,ltefield=ReadTimeout"`
	WriteTimeout      time.Duration `confx:"writeTimeout" usage:"maximum duration before timing out write of the response"`
	IdleTimeout       time.Duration `confx:"idleTimeout" usage:"maximum amount of time to wait for the next request when keep-alives are enabled"`
	// MaxRequestBodySize caps the request body via http.MaxBytesHandler. 0 means unlimited.
	// Without it a single oversized body can be read entirely into memory.
	MaxRequestBodySize int64 `confx:"maxRequestBodySize" usage:"maximum request body size in bytes, 0 for unlimited" validate:"gte=0"`
	// MaxConcurrentStreams caps HTTP/2 streams per connection. 0 uses Go's default (250).
	//
	// This is PER CONNECTION, not global. Together with MaxConnections it gives a hard
	// upper bound on in-flight requests: MaxConnections × MaxConcurrentStreams. On its own
	// it bounds nothing — a client can just open more connections.
	//
	// Lowering it buys little: the same request volume just opens more connections.
	// Leave it at 0 unless you specifically need the in-flight bound to be
	// arithmetically knowable.
	MaxConcurrentStreams int `confx:"maxConcurrentStreams" usage:"max HTTP/2 streams per connection (per-connection, not global; multiply by maxConnections for the in-flight ceiling), 0 for Go default (250)" validate:"gte=0"`
	// MaxConnections caps concurrent TCP connections via netutil.LimitListener. 0 means unlimited.
	//
	// It counts CONNECTIONS, not requests. Under HTTP/1.1 a connection carries one request
	// at a time so the two roughly coincide, but under HTTP/2 a single connection multiplexes
	// many concurrent streams — so this is NOT a concurrency limit.
	//
	// Past the limit netutil.LimitListener stops calling Accept, so further connections
	// wait in the kernel backlog until the client gives up: nothing is logged and nothing
	// is rejected.
	MaxConnections int            `confx:"maxConnections" usage:"max concurrent TCP connections (connections, NOT requests: HTTP/2 multiplexes many requests per connection; guards fd exhaustion only), 0 for unlimited" validate:"gte=0"`
	TLS            TLSConfig      `confx:"tls"`
	Security       SecurityConfig `confx:",squash"`
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
	// SkipDenySimpleRequests allows selective exemption from X-Requested-By header requirement.
	// When this function returns true for a request, the header check is skipped.
	// Common use cases: health checks, webhooks, or specific API endpoints that need exemption.
	SkipDenySimpleRequests func(r *http.Request) bool `confx:"-" json:"-"`
}

type SecurityConfig struct {
	CORS                 CORSConfig `confx:"cors"`
	DenyMIMETypeSniffing bool       `confx:"denyMIMETypeSniffing" usage:"Deny MIME type sniffing"`
	DenyClickjacking     bool       `confx:"denyClickjacking" usage:"Deny clickjacking"`
	EnableHSTS           bool       `confx:"enableHSTS" usage:"Enable HSTS"`
}
