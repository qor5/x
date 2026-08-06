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
	// 连接数上限只防 fd 耗尽，不是并发闸门：HTTP/2 一条连接可承载多个 stream，
	// 真正的并发上限应由上游（网关的 circuit breaker）或 in-flight middleware 控制。
	// 超出限制时 Accept 阻塞（连接停在内核 accept queue），不是拒绝。
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

	// 包在最外层，让 body 上限先于路由与业务 handler 生效。
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

	// HTTP/2 在两种模式下都启用：TLS 经 ALPN 协商，明文经 h2c。
	//
	// 这取代了此前 `else` 分支里的 h2c.NewHandler —— 它已被 x/net 标记
	// Deprecated（"Set the http.Server Protocols field to use unencrypted
	// HTTP/2 instead"），且为支持 HTTP/1.1 Upgrade 模式会把 h2c 连接的**首个
	// 请求整体读入内存**（其文档要求用 MaxBytesHandler 包裹，此前并没有）。
	// 标准库实现只 Peek 24 字节比对 PRI 前导，仅支持 prior-knowledge 模式，
	// 没有这个内存放大面。
	//
	// 行为差异：依赖 `Upgrade: h2c` 头升级的客户端将静默退回 HTTP/1.1（不报错）。
	// Envoy / gRPC 客户端用的都是 prior-knowledge，不受影响。
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetHTTP2(true)
	protocols.SetUnencryptedHTTP2(true)
	srv.Protocols = protocols

	// 注意：Go 1.25 的 http.Server.HTTP2 字段注释仍写着 "This field does not yet
	// have any effect"，但那句已经过时——h2_bundle.go 的 configFromServer 会经
	// fillNetHTTPServerConfig 消费它。实测（1.25.6，读服务端 SETTINGS 帧）设 42
	// 即通告 42，不设则为 250。Go 1.26 已删掉那句注释。
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
