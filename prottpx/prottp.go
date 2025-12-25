// Package prottpx provides a lightweight HTTP wrapper for gRPC unary services.
// It allows gRPC services to be served over HTTP without the overhead of a full gRPC server,
// while still supporting gRPC interceptors and protobuf serialization.
package prottpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/grpcx"
	"github.com/qor5/x/v3/hook"
	"github.com/qor5/x/v3/normalize"
	"github.com/qor5/x/v3/statusx"
	"google.golang.org/grpc"
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

type (
	// WriteErrorInput contains the input parameters for the error writing hook.
	WriteErrorInput struct {
		W                  http.ResponseWriter
		R                  *http.Request
		Error              error
		ContentTypeJSON    bool
		AcceptJSON         bool
		ConnectErrorWriter *connect.ErrorWriter
	}
	// WriteErrorOutput contains the output of the error writing hook.
	WriteErrorOutput struct {
		Written bool
	}

	// WriteResponseInput contains the input parameters for the response writing hook.
	WriteResponseInput struct {
		W               http.ResponseWriter
		R               *http.Request
		Response        proto.Message
		ContentTypeJSON bool
		AcceptJSON      bool
	}
	// WriteResponseOutput contains the output of the response writing hook.
	WriteResponseOutput struct {
		Written bool
	}

	// WriteErrorFunc is the function signature for writing errors.
	WriteErrorFunc func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error)

	// WriteResponseFunc is the function signature for writing responses.
	WriteResponseFunc func(ctx context.Context, input *WriteResponseInput) (*WriteResponseOutput, error)

	// WriteErrorIface is an interface that errors can implement to customize their own
	// HTTP response writing behavior. When an error implements this interface, the Handler
	// will delegate the response writing to the error itself instead of using the default
	// connect.ErrorWriter. This allows for custom error response formats and status codes.
	WriteErrorIface interface {
		WriteError(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error)
	}
)

var _ grpc.ServiceRegistrar = (*Handler)(nil)

// Handler is an HTTP handler that wraps gRPC unary services.
// It implements grpc.ServiceRegistrar for compatibility with generated gRPC code.
type Handler struct {
	mux               *http.ServeMux
	normalizeHandler  http.Handler
	registered        map[string]bool
	interceptors      []grpc.UnaryServerInterceptor
	connectErrWriter  *connect.ErrorWriter
	writeErrorHook    hook.Hook[WriteErrorFunc]
	writeResponseHook hook.Hook[WriteResponseFunc]
	defaultContentType string // default content type for requests when Content-Type header is not specified
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

// WithWriteErrorHook returns a HandlerOption that adds hooks for customizing error writing.
// Hooks are chained in the order they are provided.
func WithWriteErrorHook(hooks ...hook.Hook[WriteErrorFunc]) HandlerOption {
	return func(m *Handler) {
		m.writeErrorHook = hook.Prepend(m.writeErrorHook, hooks...)
	}
}

// WithWriteResponseHook returns a HandlerOption that adds hooks for customizing response writing.
// Hooks are chained in the order they are provided.
func WithWriteResponseHook(hooks ...hook.Hook[WriteResponseFunc]) HandlerOption {
	return func(m *Handler) {
		m.writeResponseHook = hook.Prepend(m.writeResponseHook, hooks...)
	}
}

// WithDefaultContentType returns a HandlerOption that sets the default content type for requests
// when the Content-Type header is not specified.
// contentType must be either "application/json" or "application/proto".
// Defaults to "application/proto" if not specified.
func WithDefaultContentType(contentType string) HandlerOption {
	contentType = strings.ToLower(strings.TrimSpace(contentType))
	if contentType != JSONContentType && contentType != ProtoContentType {
		panic(fmt.Sprintf("invalid content type: %s, must be %s or %s", contentType, JSONContentType, ProtoContentType))
	}
	return func(m *Handler) {
		m.defaultContentType = contentType
	}
}

// NewHandler creates a new Handler with the given options.
func NewHandler(opts ...HandlerOption) *Handler {
	m := &Handler{
		mux:              http.NewServeMux(),
		connectErrWriter: connect.NewErrorWriter(),
		registered:       make(map[string]bool),
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
	interceptor := grpcx.ChainUnaryServerInterceptor(m.interceptors...)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentTypeJSON := isContentTypeJSON(r, m.defaultContentType)
		acceptJSON := isAcceptJSON(r, contentTypeJSON)

		dec := func(msg any) error {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return errors.Wrap(err, "failed to read request body")
			}
			if contentTypeJSON {
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
			if writeErr := m.writeError(r.Context(), w, r, err, contentTypeJSON, acceptJSON); writeErr != nil {
				panic(writeErr)
			}
			return
		}

		if writeErr := m.writeResponse(r.Context(), w, r, resp.(proto.Message), contentTypeJSON, acceptJSON); writeErr != nil {
			panic(writeErr)
		}
	})
}

func defaultWriteError(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
	if errWriter, ok := input.Error.(WriteErrorIface); ok {
		return errWriter.WriteError(ctx, input)
	}

	// Error responses are always returned as JSON regardless of Accept or Content-Type headers.
	// This is the default behavior of connect.ErrorWriter and we maintain consistency with it.
	if werr := input.ConnectErrorWriter.Write(input.W, input.R, statusx.ConvertToConnectError(input.Error)); werr != nil {
		return nil, werr
	}
	return &WriteErrorOutput{Written: true}, nil
}

// writeError writes an error response, applying any configured hooks.
func (m *Handler) writeError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error, contentTypeJSON, acceptJSON bool) error {
	writeError := defaultWriteError
	if m.writeErrorHook != nil {
		writeError = m.writeErrorHook(writeError)
	}

	if _, err := writeError(ctx, &WriteErrorInput{
		W: w, R: r, Error: err,
		ContentTypeJSON:    contentTypeJSON,
		AcceptJSON:         acceptJSON,
		ConnectErrorWriter: m.connectErrWriter,
	}); err != nil {
		return err
	}
	return nil
}

func defaultWriteResponse(ctx context.Context, input *WriteResponseInput) (*WriteResponseOutput, error) {
	var data []byte
	var err error

	if input.AcceptJSON {
		input.W.Header().Set(HeaderContentType, JSONContentType)
		data, err = JSONMarshalOptions.Marshal(input.Response)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to marshal response to json")
		}
	} else {
		input.W.Header().Set(HeaderContentType, ProtoContentType)
		data, err = proto.Marshal(input.Response)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to marshal response to proto")
		}
	}

	input.W.Header().Set(HeaderContentLength, strconv.Itoa(len(data)))
	input.W.WriteHeader(http.StatusOK)
	if _, err := input.W.Write(data); err != nil {
		return nil, errors.Wrap(err, "failed to write response")
	}
	return &WriteResponseOutput{Written: true}, nil
}

// writeResponse writes a success response, applying any configured hooks.
func (m *Handler) writeResponse(ctx context.Context, w http.ResponseWriter, r *http.Request, resp proto.Message, contentTypeJSON, acceptJSON bool) error {
	writeResponse := defaultWriteResponse
	if m.writeResponseHook != nil {
		writeResponse = m.writeResponseHook(writeResponse)
	}

	if _, err := writeResponse(ctx, &WriteResponseInput{
		W: w, R: r,
		Response:        resp,
		ContentTypeJSON: contentTypeJSON,
		AcceptJSON:      acceptJSON,
	}); err != nil {
		return err
	}
	return nil
}

// isContentTypeJSON determines if the request body should be parsed as JSON.
// Uses defaultContentType if the Content-Type header is not specified.
func isContentTypeJSON(r *http.Request, defaultContentType string) bool {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		return defaultContentType == JSONContentType
	}
	return strings.Contains(strings.ToLower(contentType), JSONContentType)
}

// isAcceptJSON determines if the response should be JSON.
// Priority: Accept header > Content-Type header.
func isAcceptJSON(r *http.Request, contentTypeJSON bool) bool {
	accept := strings.ToLower(r.Header.Get(HeaderAccept))
	if accept == "" {
		return contentTypeJSON
	}
	return strings.Contains(accept, JSONContentType)
}
