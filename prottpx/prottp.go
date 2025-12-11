// Package prottpx provides a lightweight HTTP wrapper for gRPC unary services.
// It allows gRPC services to be served over HTTP without the overhead of a full gRPC server,
// while still supporting gRPC interceptors and protobuf serialization.
package prottpx

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/grpcx"
	"github.com/qor5/x/v3/normalize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	HeaderContentType    = http.CanonicalHeaderKey("Content-Type")
	HeaderAccept         = http.CanonicalHeaderKey("Accept")
	HeaderContentLength  = http.CanonicalHeaderKey("Content-Length")
	JSONContentType      = "application/json"
	ProtoContentType     = "application/proto"
	JSONUnmarshalOptions = protojson.UnmarshalOptions{DiscardUnknown: true}
	JSONMarshalOptions   = protojson.MarshalOptions{EmitUnpopulated: true}
)

var _ grpc.ServiceRegistrar = (*Handler)(nil)

// Handler is an HTTP handler that wraps gRPC unary services.
// It implements grpc.ServiceRegistrar for compatibility with generated gRPC code.
type Handler struct {
	mux          *http.ServeMux
	interceptors []grpc.UnaryServerInterceptor
	errWriter    *connect.ErrorWriter
	registered   map[string]bool

	normalizeHandler http.Handler
}

// HandlerOption configures the Handler.
type HandlerOption func(*Handler)

// ChainUnaryInterceptor returns a HandlerOption that adds unary server interceptors.
// Interceptors are chained in the order they are provided.
// The first interceptor will be the outermost wrapper.
func ChainUnaryInterceptor(interceptors ...grpc.UnaryServerInterceptor) HandlerOption {
	return func(m *Handler) {
		m.interceptors = append(m.interceptors, interceptors...)
	}
}

// NewHandler creates a new Handler with the given options.
func NewHandler(opts ...HandlerOption) *Handler {
	m := &Handler{
		mux:        http.NewServeMux(),
		errWriter:  connect.NewErrorWriter(),
		registered: make(map[string]bool),
	}
	m.normalizeHandler = normalize.HTTPMiddleware(m.mux)
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// ServeHTTP implements http.Handler.
func (m *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler := http.Handler(m.mux)

	// If current request context has no HTTPMeta, it means no outer normalize middleware
	// is applied, so we wrap with HTTPMiddleware here.
	if httpMeta, _ := normalize.HTTPMetaFromContext(r.Context()); httpMeta == nil {
		handler = m.normalizeHandler
	}

	handler.ServeHTTP(w, r)
}

// RegisterService implements grpc.ServiceRegistrar.
// It registers a gRPC service with the Handler, making it accessible over HTTP.
// Panics if the same service is registered twice.
func (m *Handler) RegisterService(desc *grpc.ServiceDesc, impl any) {
	if m.registered[desc.ServiceName] {
		panic(errors.Errorf("service %s is already registered", desc.ServiceName))
	}
	m.registered[desc.ServiceName] = true

	for _, method := range desc.Methods {
		pattern := "/" + desc.ServiceName + "/" + method.MethodName
		m.mux.Handle(pattern, m.handleMethod(impl, method))
	}
}

// handleMethod wraps a gRPC method as an HTTP handler.
func (m *Handler) handleMethod(service any, method grpc.MethodDesc) http.Handler {
	interceptor := grpcx.ChainUnaryServerInterceptors(m.interceptors...)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqIsJSON := strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), JSONContentType)

		dec := func(msg any) error {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return errors.Wrap(err, "failed to read request body")
			}
			if reqIsJSON {
				if err := JSONUnmarshalOptions.Unmarshal(body, msg.(proto.Message)); err != nil {
					return errors.Wrap(err, "failed to unmarshal request body via json")
				}
				return nil
			}
			if err := proto.Unmarshal(body, msg.(proto.Message)); err != nil {
				return errors.Wrap(err, "failed to unmarshal request body via proto")
			}
			return nil
		}

		resp, err := method.Handler(service, r.Context(), dec, interceptor)
		if err != nil {
			// Error responses are always returned as JSON regardless of Accept or Content-Type headers.
			// This is the default behavior of connect.ErrorWriter and we maintain consistency with it.
			if werr := m.errWriter.Write(w, r, toConnectError(err)); werr != nil {
				panic(werr)
			}
			return
		}

		writeMessage(resp.(proto.Message), w, r, reqIsJSON)
	})
}

// writeMessage writes a protobuf message to the HTTP response.
func writeMessage(msg proto.Message, w http.ResponseWriter, r *http.Request, reqIsJSON bool) {
	acceptJSON := isAcceptJSON(r, reqIsJSON)
	if acceptJSON {
		w.Header().Set(HeaderContentType, JSONContentType)
	} else {
		w.Header().Set(HeaderContentType, ProtoContentType)
	}

	var data []byte
	var err error
	if acceptJSON {
		data, err = JSONMarshalOptions.Marshal(msg)
	} else {
		data, err = proto.Marshal(msg)
	}
	if err != nil {
		panic(err)
	}

	w.Header().Set(HeaderContentLength, strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(data); err != nil {
		panic(err)
	}
}

// isAcceptJSON determines if the response should be JSON.
// Priority: Accept header > Content-Type header > default (proto).
func isAcceptJSON(r *http.Request, reqIsJSON bool) bool {
	accept := strings.ToLower(r.Header.Get(HeaderAccept))
	if accept == "" {
		return reqIsJSON
	}
	return strings.Contains(accept, JSONContentType)
}

// toConnectError converts any error to a connect.Error.
func toConnectError(err error) *connect.Error {
	if ce := new(connect.Error); errors.As(err, &ce) {
		return ce
	}

	st := status.Convert(err).Proto()
	cerr := connect.NewError(connect.Code(st.Code), errors.New(st.Message))
	for _, d := range st.Details {
		if ed, e := connect.NewErrorDetail(d); e == nil {
			cerr.AddDetail(ed)
		}
	}
	return cerr //nolint:errhandle
}
