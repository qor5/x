package httpx_test

import (
	"bufio"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/theplant/inject/lifecycle"
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

// Migrating off h2c.NewHandler drops HTTP/1.1 Upgrade-based h2c (the stdlib
// only speaks prior-knowledge). The contract we must keep is that such a
// request still gets served — a protocol downgrade, never an error.
//
// Old behaviour: 101 Switching Protocols. New: 200 OK over HTTP/1.1.
func TestNewServer_H2CUpgradeFallsBackToHTTP1(t *testing.T) {
	addr := serve(t, &httpx.ServerConfig{Address: ":0"},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, r.Proto)
		}))

	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer func() { _ = conn.Close() }()
	require.NoError(t, conn.SetDeadline(time.Now().Add(5*time.Second)))

	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: x\r\n"+
		"Connection: Upgrade, HTTP2-Settings\r\nUpgrade: h2c\r\n"+
		"HTTP2-Settings: AAMAAABkAARAAAAAAAIAAAAA\r\n\r\n")
	require.NoError(t, err)

	status, err := bufio.NewReader(conn).ReadString('\n')
	require.NoError(t, err)
	require.Contains(t, status, "200 OK",
		"an Upgrade: h2c request must still be served, just without upgrading")
	require.NotContains(t, status, "101")
}

// The old code forwarded IdleTimeout explicitly (&http2.Server{IdleTimeout: …}).
// The stdlib path has to inherit it from http.Server, or h2c connections would
// silently start living forever.
func TestNewServer_IdleTimeoutAppliesToH2C(t *testing.T) {
	const idle = 500 * time.Millisecond

	addr := serve(t, &httpx.ServerConfig{Address: ":0", IdleTimeout: idle},
		http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))

	conn, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer func() { _ = conn.Close() }()
	require.NoError(t, conn.SetDeadline(time.Now().Add(5*time.Second)))

	_, err = io.WriteString(conn, http2.ClientPreface)
	require.NoError(t, err)
	fr := http2.NewFramer(conn, conn)
	require.NoError(t, fr.WriteSettings())

	// An idle h2c connection must be shut down; the server signals that with
	// GOAWAY (and then closes), so any read eventually stops succeeding.
	deadline := time.Now().Add(4 * time.Second)
	for time.Now().Before(deadline) {
		f, err := fr.ReadFrame()
		if err != nil {
			return // connection closed — IdleTimeout did its job
		}
		if _, ok := f.(*http2.GoAwayFrame); ok {
			return
		}
	}
	t.Fatal("idle h2c connection was never closed — IdleTimeout is not reaching HTTP/2")
}

// MaxConnections caps concurrent TCP connections. It takes effect in
// SetupListener (netutil.LimitListener), not in NewServer, so this test goes
// through the real wiring rather than the bare net.Listen used elsewhere.
//
// LimitListener enforces the cap by not Accept-ing past it — connections sit
// in the kernel backlog rather than being refused — so the observable effect
// is that a second connection gets no response while the first is held open.
func TestNewServer_MaxConnections(t *testing.T) {
	conf := &httpx.ServerConfig{Address: "127.0.0.1:0", MaxConnections: 1}

	lc := lifecycle.New()
	listener, err := httpx.SetupListener(lc, conf)
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })

	srv, err := httpx.NewServer(conf, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = io.WriteString(w, "ok")
		}))
	require.NoError(t, err)
	go func() { _ = srv.Serve(listener) }()
	t.Cleanup(func() { _ = srv.Close() })

	addr := listener.Addr().String()

	// First connection: served normally, then held open via keep-alive so it
	// keeps occupying the single slot.
	held, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer func() { _ = held.Close() }()
	require.NoError(t, held.SetDeadline(time.Now().Add(5*time.Second)))
	_, err = io.WriteString(held, "GET / HTTP/1.1\r\nHost: x\r\n\r\n")
	require.NoError(t, err)
	status, err := bufio.NewReader(held).ReadString('\n')
	require.NoError(t, err)
	require.Contains(t, status, "200 OK", "the first connection must be served")

	// Second connection: the TCP handshake still completes (kernel backlog),
	// but the server never Accepts it, so no response arrives.
	blocked, err := net.Dial("tcp", addr)
	require.NoError(t, err)
	defer func() { _ = blocked.Close() }()
	require.NoError(t, blocked.SetDeadline(time.Now().Add(700*time.Millisecond)))
	_, err = io.WriteString(blocked, "GET / HTTP/1.1\r\nHost: x\r\n\r\n")
	require.NoError(t, err)
	_, err = bufio.NewReader(blocked).ReadString('\n')
	require.Error(t, err, "a second connection must not be served while the cap is taken")

	// Releasing the slot lets the queued connection through — the cap blocks,
	// it does not permanently reject. (It is the already-queued one that gets
	// Accept-ed next, so re-use `blocked` rather than dialling afresh.)
	require.NoError(t, held.Close())
	require.NoError(t, blocked.SetDeadline(time.Now().Add(5*time.Second)))
	status, err = bufio.NewReader(blocked).ReadString('\n')
	require.NoError(t, err)
	require.Contains(t, status, "200 OK", "the slot must be reusable once freed")
}
