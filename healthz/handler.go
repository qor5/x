package healthz

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/httpx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

const Path = "/healthz"

type Response struct {
	Status string `json:"status"`
}

// newHandler creates a new health check handler
func newHandler(opts ...Option) http.HandlerFunc {
	o := new(options)
	for _, opt := range opts {
		opt(o)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		for _, checker := range o.checkers {
			if err = checker(w, r); err != nil {
				break
			}
		}
		resp := &Response{}
		var statusCode int
		if err != nil {
			resp.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING.String()
			statusCode = http.StatusServiceUnavailable
		} else {
			resp.Status = grpc_health_v1.HealthCheckResponse_SERVING.String()
			statusCode = http.StatusOK
		}
		w.Header().Set(httpx.HeaderContentType, "application/json")
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// HTTPMiddleware creates a middleware that intercepts health check requests at the top level
func HTTPMiddleware(opts ...Option) func(next http.Handler) http.Handler {
	handler := newHandler(opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.TrimSpace(r.URL.Path) == Path {
				handler(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

type Checker func(w http.ResponseWriter, r *http.Request) error

type options struct {
	checkers []Checker
}

type Option func(*options)

func WithChecker(checker Checker) Option {
	return func(o *options) {
		o.checkers = append(o.checkers, checker)
	}
}

func WithGRPCHealthChecker(conn grpc.ClientConnInterface) Option {
	client := grpc_health_v1.NewHealthClient(conn)
	return WithChecker(func(_ http.ResponseWriter, r *http.Request) error {
		_, err := client.Check(r.Context(), &grpc_health_v1.HealthCheckRequest{Service: ""}, grpc.WaitForReady(true))
		return errors.Wrap(err, "grpc server is not ready")
	})
}
