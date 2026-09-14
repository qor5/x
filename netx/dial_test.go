package netx

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theplant/inject/lifecycle"
)

func TestSetupListenerFactory(t *testing.T) {
	t.Run("listen on random port", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(SetupListenerFactory("test-listener", "127.0.0.1:0")))

		var listener net.Listener
		require.NoError(t, lc.Resolve(&listener))

		tcpAddr, ok := listener.Addr().(*net.TCPAddr)
		require.True(t, ok)
		assert.True(t, tcpAddr.IP.Equal(net.IPv4(127, 0, 0, 1)))
		assert.NotZero(t, tcpAddr.Port)
	})

	t.Run("addr is connectable when listening on unspecified", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(SetupListenerFactory("test-listener", ":0")))

		var listener net.Listener
		require.NoError(t, lc.Resolve(&listener))

		tcpAddr, ok := listener.Addr().(*net.TCPAddr)
		require.True(t, ok)
		assert.True(t, tcpAddr.IP.IsLoopback(),
			"expected loopback but got %s", tcpAddr.IP)
	})

	t.Run("addr is connectable when listening on [::]", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(SetupListenerFactory("test-listener", "[::]:0")))

		var listener net.Listener
		require.NoError(t, lc.Resolve(&listener))

		tcpAddr, ok := listener.Addr().(*net.TCPAddr)
		require.True(t, ok)
		assert.True(t, tcpAddr.IP.Equal(net.IPv6loopback),
			"expected ::1 but got %s", tcpAddr.IP)
	})

	t.Run("listener accepts connections", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(SetupListenerFactory("test-listener", "127.0.0.1:0")))

		var listener net.Listener
		require.NoError(t, lc.Resolve(&listener))

		done := make(chan struct{})
		go func() {
			defer close(done)
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conn.Close()
		}()

		conn, err := net.Dial("tcp", listener.Addr().String())
		require.NoError(t, err)
		conn.Close()
		<-done
	})

	t.Run("lifecycle stop closes listener", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		var listener net.Listener
		done := make(chan error, 1)
		go func() {
			done <- lifecycle.New().Serve(ctx,
				SetupListenerFactory("test-listener", "127.0.0.1:0"),
				func(lc *lifecycle.Lifecycle, l net.Listener) {
					listener = l
					lc.Add(lifecycle.NewFuncService(func(ctx context.Context) error {
						<-ctx.Done()
						return nil
					}))
				},
			)
		}()

		require.Eventually(t, func() bool { return listener != nil }, 5*time.Second, 10*time.Millisecond)

		cancel()
		require.NoError(t, <-done)

		_, err := listener.Accept()
		assert.Error(t, err)
	})

	t.Run("invalid address returns error", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(SetupListenerFactory("test-listener", "invalid-address")))

		var listener net.Listener
		err := lc.Resolve(&listener)
		require.Error(t, err)
	})
}
