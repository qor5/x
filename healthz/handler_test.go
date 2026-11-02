package healthz

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestHandler(t *testing.T) {
	t.Run("additional grpc health checker", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, Path, nil)
		w := httptest.NewRecorder()

		server := grpc.NewServer()
		healthServer := health.NewServer()
		grpc_health_v1.RegisterHealthServer(server, healthServer)
		listener := bufconn.Listen(1024 * 1024)
		serverErr := make(chan error)
		go func() {
			serverErr <- server.Serve(listener)
		}()
		conn, err := grpc.NewClient("passthrough://bufnet",
			grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
				return listener.DialContext(ctx)
			}),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		require.NoError(t, err)

		// execute the handler
		newHandler(WithGRPCHealthChecker(conn))(w, req)

		res := w.Result()
		t.Cleanup(func() {
			require.NoError(t, res.Body.Close())
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

		var response Response
		err = json.NewDecoder(res.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING.String(), response.Status)

		server.GracefulStop()
		require.NoError(t, conn.Close())
		require.NoError(t, <-serverErr)
	})

	t.Run("additional checker returns error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, Path, nil)
		w := httptest.NewRecorder()

		// execute the handler
		newHandler(
			WithChecker(func(_ http.ResponseWriter, _ *http.Request) error {
				return errors.New("some error")
			}),
		)(w, req)

		res := w.Result()
		t.Cleanup(func() {
			require.NoError(t, res.Body.Close())
		})

		assert.Equal(t, http.StatusServiceUnavailable, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

		var response Response
		err := json.NewDecoder(res.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, grpc_health_v1.HealthCheckResponse_NOT_SERVING.String(), response.Status)
	})
}

func TestHTTPMiddleware(t *testing.T) {
	t.Run("intercepts health check path", func(t *testing.T) {
		// Create a dummy next handler
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		})

		// Create middleware
		middleware := HTTPMiddleware()
		handler := middleware(nextHandler)

		// Test health check path
		req := httptest.NewRequest(http.MethodGet, Path, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		res := w.Result()
		t.Cleanup(func() {
			require.NoError(t, res.Body.Close())
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

		var response Response
		err := json.NewDecoder(res.Body).Decode(&response)
		require.NoError(t, err)
		assert.Equal(t, grpc_health_v1.HealthCheckResponse_SERVING.String(), response.Status)
	})

	t.Run("passes through non-health requests", func(t *testing.T) {
		// Create a dummy next handler
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("normal response"))
		})

		// Create middleware
		middleware := HTTPMiddleware()
		handler := middleware(nextHandler)

		// Test non-health check path
		req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		res := w.Result()
		t.Cleanup(func() {
			require.NoError(t, res.Body.Close())
		})

		assert.Equal(t, http.StatusOK, res.StatusCode)
		body := w.Body.String()
		assert.Contains(t, body, "normal response")
	})
}
