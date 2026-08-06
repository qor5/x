package httpx_test

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"

	"github.com/qor5/x/v3/httpx"
)

// serve 起一个监听在随机端口上的 server，返回其地址。
func serve(t *testing.T, conf *httpx.ServerConfig, handler http.Handler) string {
	t.Helper()

	srv, err := httpx.NewServer(conf, handler)
	require.NoError(t, err)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	go func() { _ = srv.Serve(ln) }()
	t.Cleanup(func() { _ = srv.Close() })

	return ln.Addr().String()
}

// h2cClient 用 prior-knowledge 模式（直接发 HTTP/2 前导）连明文端口，
// 这正是 Envoy / gRPC 客户端在 appProtocol=h2c 下的行为。
func h2cClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
				return (&net.Dialer{}).DialContext(ctx, network, addr)
			},
		},
	}
}

// 迁移到 http.Server.Protocols 之后，明文 HTTP/2 必须仍然可用——
// 这是替换掉已废弃的 h2c.NewHandler 时最需要守住的行为。
func TestNewServer_H2C(t *testing.T) {
	addr := serve(t, &httpx.ServerConfig{Address: ":0"},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, r.Proto)
		}))

	resp, err := h2cClient().Get("http://" + addr)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "HTTP/2.0", string(body))
}

// 同一个 server 必须同时还能服务 HTTP/1.1。
func TestNewServer_HTTP1StillWorks(t *testing.T) {
	addr := serve(t, &httpx.ServerConfig{Address: ":0"},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, r.Proto)
		}))

	resp, err := http.Get("http://" + addr)
	require.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "HTTP/1.1", string(body))
}

// 读服务端在 SETTINGS 帧里通告的 MAX_CONCURRENT_STREAMS。
// 这是唯一能证明 http.Server.HTTP2 真的生效的方式——Go 1.25 的字段注释还写着
// "does not yet have any effect"，那句已经过时，但只能实测来确认。
func advertisedMaxStreams(t *testing.T, addr string) uint32 {
	t.Helper()

	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer func() { _ = conn.Close() }()
	require.NoError(t, conn.SetDeadline(time.Now().Add(5*time.Second)))

	_, err = io.WriteString(conn, http2.ClientPreface)
	require.NoError(t, err)

	fr := http2.NewFramer(conn, conn)
	require.NoError(t, fr.WriteSettings())

	for range 5 {
		f, err := fr.ReadFrame()
		require.NoError(t, err)
		sf, ok := f.(*http2.SettingsFrame)
		if !ok {
			continue
		}
		if v, ok := sf.Value(http2.SettingMaxConcurrentStreams); ok {
			return v
		}
	}
	t.Fatal("server never advertised MAX_CONCURRENT_STREAMS")
	return 0
}

func TestNewServer_MaxConcurrentStreams(t *testing.T) {
	noop := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

	t.Run("configured value is advertised", func(t *testing.T) {
		addr := serve(t, &httpx.ServerConfig{Address: ":0", MaxConcurrentStreams: 42}, noop)
		require.Equal(t, uint32(42), advertisedMaxStreams(t, addr))
	})

	t.Run("zero falls back to the Go default", func(t *testing.T) {
		addr := serve(t, &httpx.ServerConfig{Address: ":0"}, noop)
		require.Equal(t, uint32(250), advertisedMaxStreams(t, addr))
	})
}

func TestNewServer_MaxRequestBodySize(t *testing.T) {
	const limit = 16

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.ReadAll(r.Body); err != nil {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	t.Run("under the limit passes", func(t *testing.T) {
		addr := serve(t, &httpx.ServerConfig{Address: ":0", MaxRequestBodySize: limit}, handler)

		resp, err := http.Post("http://"+addr, "text/plain", strings.NewReader("short"))
		require.NoError(t, err)
		defer func() { _ = resp.Body.Close() }()
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("over the limit is rejected", func(t *testing.T) {
		addr := serve(t, &httpx.ServerConfig{Address: ":0", MaxRequestBodySize: limit}, handler)

		resp, err := http.Post("http://"+addr, "text/plain", strings.NewReader(strings.Repeat("x", limit*4)))
		require.NoError(t, err)
		defer func() { _ = resp.Body.Close() }()
		require.Equal(t, http.StatusRequestEntityTooLarge, resp.StatusCode)
	})

	t.Run("zero means unlimited", func(t *testing.T) {
		addr := serve(t, &httpx.ServerConfig{Address: ":0"}, handler)

		resp, err := http.Post("http://"+addr, "text/plain", strings.NewReader(strings.Repeat("x", limit*4)))
		require.NoError(t, err)
		defer func() { _ = resp.Body.Close() }()
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
