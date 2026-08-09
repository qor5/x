package httpx

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"log/slog"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/netx"
	"github.com/theplant/inject/lifecycle"
	"golang.org/x/net/netutil"
)

type Listener net.Listener

func SetupListener(lc *lifecycle.Lifecycle, conf *ServerConfig) (Listener, error) {
	listener, err := netx.SetupListenerFactory("http-listener", conf.Address)(lc)
	if err != nil {
		return nil, err
	}
	// A connection cap guards against fd exhaustion; it is NOT a concurrency
	// limit, since one HTTP/2 connection carries many streams. Bound concurrency
	// upstream (the gateway's circuit breaker) or with an in-flight middleware.
	// Past the limit Accept blocks — connections queue in the kernel backlog
	// rather than being rejected.
	if conf.MaxConnections > 0 {
		listener = netutil.LimitListener(listener, conf.MaxConnections)
	}
	return listener, nil
}

func SetupServerFactory(name string, handler http.Handler) func(ctx context.Context, lc *lifecycle.Lifecycle, conf *ServerConfig, listener Listener) (*http.Server, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle, conf *ServerConfig, listener Listener) (*http.Server, error) {
		srv, err := NewServer(conf, handler)
		if err != nil {
			return nil, err
		}
		lc.Add(lifecycle.NewFuncService(func(ctx context.Context) error {
			if srv.TLSConfig != nil {
				slog.InfoContext(ctx, "HTTPS server listening", "address", listener.Addr().String())
				if err := srv.ServeTLS(listener, "", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "failed to start HTTPS server")
				}
			} else {
				slog.InfoContext(ctx, "HTTP server listening", "address", listener.Addr().String())
				if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "failed to start HTTP server")
				}
			}
			return nil
		}).WithStop(func(ctx context.Context) error {
			// Attempt graceful shutdown first - waits for active connections to finish
			err := srv.Shutdown(ctx)
			if err != nil {
				// If graceful shutdown fails (timeout, context cancelled, etc.),
				// force immediate shutdown to ensure resources are released
				if closeErr := srv.Close(); closeErr != nil {
					return errors.Wrap(closeErr, "failed to force close HTTP server after shutdown failure")
				}
				return errors.Wrap(err, "graceful shutdown failed, forced close completed")
			}
			// Graceful shutdown succeeded - no need to call Close()
			return nil
		}).WithName(name))
		return srv, nil
	}
}

func NewServer(conf *ServerConfig, handler http.Handler) (*http.Server, error) {
	// Cannot be a `ltefield` struct tag: ReadTimeout == 0 means no read deadline
	// at all, not "zero seconds", so it is not an upper bound. That meaning is
	// defined by net/http, which the tag layer cannot see — tagged, a config
	// that sets only a header timeout would fail validation.
	if conf.ReadTimeout > 0 && conf.ReadHeaderTimeout > conf.ReadTimeout {
		return nil, errors.Errorf(
			"readHeaderTimeout (%s) must not exceed readTimeout (%s)",
			conf.ReadHeaderTimeout, conf.ReadTimeout)
	}

	// Normalize PathPrefix to ensure predictable behavior:
	// - Always starts with "/" (add if missing)
	// - Never ends with "/" unless it's the root path "/"
	// - Root path "/" is treated as no prefix (skips StripPrefix)
	// This prevents common configuration errors and makes http.StripPrefix behavior consistent
	if conf.PathPrefix != "" && conf.PathPrefix != "/" {
		pathPrefix := conf.PathPrefix

		// Ensure prefix starts with "/"
		if !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}
		// Remove trailing slash unless it's the root path "/"
		if len(pathPrefix) > 1 && strings.HasSuffix(pathPrefix, "/") {
			pathPrefix = strings.TrimSuffix(pathPrefix, "/")
		}

		handler = http.StripPrefix(pathPrefix, handler)
	}

	// Outermost, so the body cap applies before routing and before any
	// business handler gets to read.
	if conf.MaxRequestBodySize > 0 {
		handler = http.MaxBytesHandler(handler, conf.MaxRequestBodySize)
	}

	srv := &http.Server{
		ReadTimeout:       conf.ReadTimeout,
		ReadHeaderTimeout: conf.ReadHeaderTimeout,
		WriteTimeout:      conf.WriteTimeout,
		IdleTimeout:       conf.IdleTimeout,
		Handler:           handler,
	}

	// HTTP/2 on both paths: negotiated via ALPN under TLS, h2c in cleartext.
	//
	// This replaces the h2c.NewHandler that used to live in the `else` branch.
	// x/net marks it Deprecated ("Set the http.Server Protocols field to use
	// unencrypted HTTP/2 instead"), and to support HTTP/1.1 Upgrade it reads
	// the FIRST request on an h2c connection entirely into memory (its own doc
	// asks callers to wrap it in MaxBytesHandler; we never did). The stdlib
	// implementation only peeks 24 bytes for the PRI preface — prior-knowledge
	// mode only — so that memory amplification does not exist here.
	//
	// Behaviour change: a client relying on the `Upgrade: h2c` header no longer
	// upgrades. It falls back to HTTP/1.1 silently — the request is still served
	// normally (verified: 200 OK rather than 101 Switching Protocols), so this
	// is a downgrade in protocol, not a failure. Envoy (with appProtocol h2c)
	// and gRPC clients both use prior-knowledge and are unaffected.
	//
	// Also verified unchanged: http.Server.IdleTimeout still governs h2c
	// connections. The old code forwarded it explicitly via
	// &http2.Server{IdleTimeout: ...}; the stdlib path inherits it, and both
	// close an idle connection at the same moment.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(true)
	protocols.SetUnencryptedHTTP2(true)
	srv.Protocols = protocols

	// Heads-up for anyone auditing this: through Go 1.25 the doc comment on
	// http.Server.HTTP2 still reads "This field does not yet have any effect"
	// (go.dev/issue/67813). That comment is wrong, and has been since the field
	// landed — h2_bundle.go's configFromServer has always fed it through
	// fillNetHTTPConfig. Measured by reading the server's SETTINGS frame,
	// MaxConcurrentStreams: 42 is advertised as 42 on go1.24.1, 1.24.11, 1.25.1,
	// 1.25.6, 1.25.12 and 1.26.3 alike (250 when unset). Go 1.26 finally dropped
	// the stale comment. server_test.go pins this so it cannot silently regress.
	if conf.MaxConcurrentStreams > 0 {
		srv.HTTP2 = &http.HTTP2Config{MaxConcurrentStreams: conf.MaxConcurrentStreams}
	}

	if conf.TLS.Enabled {
		cert, err := loadTLSCertificate(conf.TLS.CertBase64, conf.TLS.KeyBase64)
		if err != nil {
			return nil, err
		}
		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}
	return srv, nil
}

func loadTLSCertificate(certBase64, keyBase64 string) (tls.Certificate, error) {
	certBytes, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		return tls.Certificate{}, errors.Wrap(err, "failed to decode certificate")
	}
	keyBytes, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return tls.Certificate{}, errors.Wrap(err, "failed to decode private key")
	}

	cert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return tls.Certificate{}, errors.Wrap(err, "failed to load key pair")
	}
	return cert, nil
}
