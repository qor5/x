package httpx

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"net"
	"net/http"

	"github.com/pkg/errors"
	kitlog "github.com/theplant/appkit/log"
	"github.com/theplant/inject/lifecycle"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func SetupFactory(handler http.Handler) []any {
	return []any{
		SetupListener,
		SetupServer(handler),
	}
}

type Listener net.Listener

func SetupListener(lc *lifecycle.Lifecycle, conf *Config) (Listener, error) {
	listener, err := net.Listen("tcp", conf.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen on %s", conf.Address)
	}
	lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
		return errors.Wrap(listener.Close(), "failed to close HTTP listener")
	}).WithName("http-listener"))
	return Listener(listener), nil
}

func SetupServer(handler http.Handler) func(lc *lifecycle.Lifecycle, conf *Config, listener Listener, logger *kitlog.Logger) (*http.Server, error) {
	return func(lc *lifecycle.Lifecycle, conf *Config, listener Listener, logger *kitlog.Logger) (*http.Server, error) {
		srv, err := NewServer(conf, handler)
		if err != nil {
			return nil, err
		}
		lc.Add(lifecycle.NewFuncService(func(ctx context.Context) error {
			if conf.TLS.Enabled {
				logger.Info().Log("msg", "HTTPS server listening", "address", listener.Addr())
				if err := srv.ServeTLS(listener, "", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "failed to start HTTPS server")
				}
			} else {
				logger.Info().Log("msg", "HTTP server listening", "address", listener.Addr())
				if err := srv.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
					return errors.Wrap(err, "failed to start HTTP server")
				}
			}
			return nil
		}).WithStop(func(ctx context.Context) error {
			err := srv.Shutdown(ctx)
			if err != nil {
				return errors.Wrap(err, "failed to shutdown HTTP server")
			}
			return errors.Wrap(srv.Close(), "failed to close HTTP server")
		}).WithName("http-server"))
		return srv, nil
	}
}

func NewServer(conf *Config, handler http.Handler) (*http.Server, error) {
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
