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
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Listener net.Listener

func SetupListener(lc *lifecycle.Lifecycle, conf *ServerConfig) (Listener, error) {
	return netx.SetupListenerFactory("http-listener", conf.Address)(lc)
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

	srv := &http.Server{
		ReadTimeout:       conf.ReadTimeout,
		ReadHeaderTimeout: conf.ReadHeaderTimeout,
		WriteTimeout:      conf.WriteTimeout,
		IdleTimeout:       conf.IdleTimeout,
		Handler:           handler,
	}
	if conf.TLS.Enabled {
		cert, err := loadTLSCertificate(conf.TLS.CertBase64, conf.TLS.KeyBase64)
		if err != nil {
			return nil, err
		}
		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	} else {
		srv.Handler = h2c.NewHandler(srv.Handler, &http2.Server{
			IdleTimeout: srv.IdleTimeout,
		})
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
