package connectx

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/qor5/x/v3/httpx"
	"github.com/qor5/x/v3/i18nx"
	"github.com/qor5/x/v3/statusx"
)

type Mux struct {
	mux         *http.ServeMux
	connectOpts []connect.HandlerOption
	handler     http.Handler
}

func (h *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func (h *Mux) Handle(fns ...func(opts ...connect.HandlerOption) (string, http.Handler)) *Mux {
	for _, fn := range fns {
		h.mux.Handle(fn(h.connectOpts...))
	}
	return h
}

// ensureConnectErrorIncoming clones incoming metadata and sets the ensure-connect-error flag.
func ensureConnectErrorIncoming(ctx context.Context) context.Context {
	md := metadata.ExtractIncoming(ctx).Clone()
	md.Set(statusx.HeaderEnsureConnectError, "true")
	return md.ToIncoming(ctx)
}

// rejectGRPCProtocol returns a middleware that rejects gRPC protocol requests.
func rejectGRPCProtocol(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaType, _, err := httpx.ParseContentType(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Invalid Content-Type header"}`))
			return
		}

		// Only allow Connect protocol content types (whitelist approach)
		switch mediaType {
		case "application/json", "application/proto":
			// Connect protocol content types are allowed
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"error": "Only Connect protocol is allowed. Supported Content-Type: application/json, application/proto"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// NewVProtoMux builds a Connect mux that returns VProto errors by default.
func NewVProtoMux(ib *i18nx.I18N, opts ...connect.HandlerOption) *Mux {
	connectOpts := []connect.HandlerOption{
		connect.WithCompressMinBytes(1024),
		connect.WithInterceptors(
			connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
				return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
					res, err := next(ctx, request)
					if err != nil {
						if !statusx.EnsureConnectError(ctx) {
							// Panic here to propagate the error to a top-level handler that will translate it into a VProto error response.
							// This approach ensures that errors not already marked for Connect error handling are consistently converted at a single point,
							// rather than being handled piecemeal throughout the codebase. While using panic for control flow is unusual in Go,
							// it is intentional here to centralize error translation and maintain consistent error responses for clients.
							panic(err)
						}
					}
					return res, err
				}
			}),
			statusx.UnaryConnectInterceptor(ib, func(ctx context.Context, _ connect.AnyRequest) bool {
				// Use request metadata to decide whether to ensure connect error
				return statusx.EnsureConnectError(ctx)
			}),
		),
	}
	connectOpts = append(connectOpts, opts...)

	mux := http.NewServeMux()
	h := statusx.NewVProtoHTTPErrorWriter(ib)(mux)

	// Apply gRPC protocol rejection middleware
	h = rejectGRPCProtocol(h)

	return &Mux{
		mux:         mux,
		connectOpts: connectOpts,
		handler:     h,
	}
}

// NewMux builds a Connect mux that returns Connect errors by default.
func NewMux(ib *i18nx.I18N, opts ...connect.HandlerOption) *Mux {
	mux := NewVProtoMux(ib, opts...)
	handler := mux.handler
	mux.handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r.WithContext(ensureConnectErrorIncoming(r.Context())))
	})
	return mux
}
