package prottpx

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"connectrpc.com/connect"
	"github.com/qor5/x/v3/normalize"
	testdatav1 "github.com/qor5/x/v3/prottpx/gen/testdata/v1"
	"github.com/qor5/x/v3/statusx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// echoServer implements testdatav1.EchoServiceServer for testing.
type echoServer struct {
	testdatav1.UnimplementedEchoServiceServer
}

func (s *echoServer) Echo(ctx context.Context, req *testdatav1.EchoRequest) (*testdatav1.EchoResponse, error) {
	return &testdatav1.EchoResponse{Message: "Echo: " + req.Message}, nil
}

func (s *echoServer) EchoWithError(ctx context.Context, req *testdatav1.EchoWithErrorRequest) (*testdatav1.EchoWithErrorResponse, error) {
	switch req.ErrorType {
	case testdatav1.ErrorType_ERROR_TYPE_NONE, testdatav1.ErrorType_ERROR_TYPE_UNSPECIFIED:
		return &testdatav1.EchoWithErrorResponse{Message: "EchoWithError: " + req.Message}, nil
	case testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS:
		return nil, status.Error(codes.InvalidArgument, "grpc status error")
	case testdatav1.ErrorType_ERROR_TYPE_STATUSX:
		return nil, statusx.Error(codes.InvalidArgument, "STATUSX_ERROR", "statusx error")
	case testdatav1.ErrorType_ERROR_TYPE_CONNECT:
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}
	return nil, errors.New("unknown error type")
}

func TestServeMux_ContentType(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	tests := []struct {
		name                string
		reqContentType      string
		reqAccept           string
		expectedRespCT      string
		useJSON             bool
		expectedRespUseJSON bool
	}{
		{
			name:                "JSON request, no Accept -> JSON response",
			reqContentType:      JSONContentType,
			reqAccept:           "",
			expectedRespCT:      JSONContentType,
			useJSON:             true,
			expectedRespUseJSON: true,
		},
		{
			name:                "Proto request, no Accept -> Proto response",
			reqContentType:      ProtoContentType,
			reqAccept:           "",
			expectedRespCT:      ProtoContentType,
			useJSON:             false,
			expectedRespUseJSON: false,
		},
		{
			name:                "JSON request, Accept JSON -> JSON response",
			reqContentType:      JSONContentType,
			reqAccept:           JSONContentType,
			expectedRespCT:      JSONContentType,
			useJSON:             true,
			expectedRespUseJSON: true,
		},
		{
			name:                "JSON request, Accept Proto -> Proto response",
			reqContentType:      JSONContentType,
			reqAccept:           ProtoContentType,
			expectedRespCT:      ProtoContentType,
			useJSON:             true,
			expectedRespUseJSON: false,
		},
		{
			name:                "Proto request, Accept JSON -> JSON response",
			reqContentType:      ProtoContentType,
			reqAccept:           JSONContentType,
			expectedRespCT:      JSONContentType,
			useJSON:             false,
			expectedRespUseJSON: true,
		},
		{
			name:                "Proto request, Accept Proto -> Proto response",
			reqContentType:      ProtoContentType,
			reqAccept:           ProtoContentType,
			expectedRespCT:      ProtoContentType,
			useJSON:             false,
			expectedRespUseJSON: false,
		},
		{
			name:                "No Content-Type, no Accept -> Proto response (default)",
			reqContentType:      "",
			reqAccept:           "",
			expectedRespCT:      ProtoContentType,
			useJSON:             false,
			expectedRespUseJSON: false,
		},
		{
			name:                "No Content-Type, Accept JSON -> JSON response",
			reqContentType:      "",
			reqAccept:           JSONContentType,
			expectedRespCT:      JSONContentType,
			useJSON:             false,
			expectedRespUseJSON: true,
		},
		{
			name:                "Content-Type with charset, no Accept -> JSON response",
			reqContentType:      JSONContentType + "; charset=utf-8",
			reqAccept:           "",
			expectedRespCT:      JSONContentType,
			useJSON:             true,
			expectedRespUseJSON: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &testdatav1.EchoRequest{Message: "hello"}
			var body []byte
			var err error
			if tt.useJSON {
				body, err = JSONMarshalOptions.Marshal(req)
			} else {
				body, err = proto.Marshal(req)
			}
			require.NoError(t, err)

			httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
			if tt.reqContentType != "" {
				httpReq.Header.Set(HeaderContentType, tt.reqContentType)
			}
			if tt.reqAccept != "" {
				httpReq.Header.Set(HeaderAccept, tt.reqAccept)
			}

			rec := httptest.NewRecorder()
			hdr.ServeHTTP(rec, httpReq)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, tt.expectedRespCT, rec.Header().Get(HeaderContentType))

			// Verify Content-Length header matches actual body length
			contentLength := rec.Header().Get(HeaderContentLength)
			assert.NotEmpty(t, contentLength)
			expectedLen, err := strconv.Atoi(contentLength)
			require.NoError(t, err)
			assert.Equal(t, expectedLen, len(rec.Body.Bytes()))

			// Verify response can be unmarshaled
			resp := &testdatav1.EchoResponse{}
			if tt.expectedRespUseJSON {
				err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
			} else {
				err = proto.Unmarshal(rec.Body.Bytes(), resp)
			}
			require.NoError(t, err)
			assert.Equal(t, "Echo: hello", resp.Message)
		})
	}
}

// TestServeMux_ErrorResponse verifies error response behavior.
// Error responses are always returned as JSON regardless of Accept or Content-Type headers.
// This is the default behavior of connect.ErrorWriter and we maintain consistency with it.
func TestServeMux_ErrorResponse(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	tests := []struct {
		name               string
		reqContentType     string
		reqAccept          string
		errorType          testdatav1.ErrorType
		expectedStatusCode int
	}{
		{
			name:               "grpc status error with JSON request",
			reqContentType:     JSONContentType,
			reqAccept:          JSONContentType,
			errorType:          testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "grpc status error with Proto request",
			reqContentType:     ProtoContentType,
			reqAccept:          ProtoContentType,
			errorType:          testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "statusx error with JSON request",
			reqContentType:     "",
			reqAccept:          JSONContentType,
			errorType:          testdatav1.ErrorType_ERROR_TYPE_STATUSX,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "statusx error with Proto request",
			reqContentType:     ProtoContentType,
			reqAccept:          ProtoContentType,
			errorType:          testdatav1.ErrorType_ERROR_TYPE_STATUSX,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "connect error with JSON request",
			reqContentType:     "",
			reqAccept:          "",
			errorType:          testdatav1.ErrorType_ERROR_TYPE_CONNECT,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "connect error with Proto request",
			reqContentType:     "",
			reqAccept:          ProtoContentType,
			errorType:          testdatav1.ErrorType_ERROR_TYPE_CONNECT,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "unknown error type returns internal server error",
			reqContentType:     JSONContentType,
			reqAccept:          JSONContentType,
			errorType:          testdatav1.ErrorType(999),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &testdatav1.EchoWithErrorRequest{Message: "error test", ErrorType: tt.errorType}
			var body []byte
			var err error
			if tt.reqContentType == JSONContentType {
				body, err = JSONMarshalOptions.Marshal(req)
			} else {
				body, err = proto.Marshal(req)
			}
			require.NoError(t, err)

			httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/EchoWithError", bytes.NewReader(body))
			httpReq.Header.Set(HeaderContentType, tt.reqContentType)
			if tt.reqAccept != "" {
				httpReq.Header.Set(HeaderAccept, tt.reqAccept)
			}

			rec := httptest.NewRecorder()
			hdr.ServeHTTP(rec, httpReq)

			// Verify expected status code
			assert.Equal(t, tt.expectedStatusCode, rec.Code)

			// Error responses are always JSON (connect.ErrorWriter default behavior)
			respContentType := rec.Header().Get(HeaderContentType)
			assert.True(t, strings.Contains(respContentType, JSONContentType),
				"error response Content-Type should be JSON, got: %s", respContentType)

			// Verify response body is valid JSON
			respBody := rec.Body.String()
			t.Logf("Error response body: %s", respBody)
			assert.True(t, strings.HasPrefix(strings.TrimSpace(respBody), "{"),
				"error response body should be JSON, got: %s", respBody)
		})
	}
}

func TestServeMux_MethodNotAllowed(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	methods := []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch}
	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			httpReq := httptest.NewRequest(method, "/testdata.v1.EchoService/Echo", nil)
			rec := httptest.NewRecorder()
			hdr.ServeHTTP(rec, httpReq)

			assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
		})
	}
}

func TestServeMux_WithPathPrefix(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	// Create a parent hdr with a path prefix
	parentMux := http.NewServeMux()
	parentMux.Handle("/api/", http.StripPrefix("/api", hdr))

	req := &testdatav1.EchoRequest{Message: "prefixed"}
	body, err := JSONMarshalOptions.Marshal(req)
	require.NoError(t, err)

	// Request with prefix
	httpReq := httptest.NewRequest(http.MethodPost, "/api/testdata.v1.EchoService/Echo", bytes.NewReader(body))
	httpReq.Header.Set(HeaderContentType, JSONContentType)

	rec := httptest.NewRecorder()
	parentMux.ServeHTTP(rec, httpReq)

	assert.Equal(t, http.StatusOK, rec.Code)

	resp := &testdatav1.EchoResponse{}
	err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
	require.NoError(t, err)
	assert.Equal(t, "Echo: prefixed", resp.Message)
}

func TestServeMux_ChainUnaryInterceptor(t *testing.T) {
	var callOrder []string

	interceptor1 := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		callOrder = append(callOrder, "interceptor1-before")
		resp, err := handler(ctx, req)
		callOrder = append(callOrder, "interceptor1-after")
		return resp, err
	}

	interceptor2 := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		callOrder = append(callOrder, "interceptor2-before")
		resp, err := handler(ctx, req)
		callOrder = append(callOrder, "interceptor2-after")
		return resp, err
	}

	hdr := NewHandler(ChainUnaryInterceptor(interceptor1, interceptor2))
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	req := &testdatav1.EchoRequest{Message: "interceptor test"}
	body, err := JSONMarshalOptions.Marshal(req)
	require.NoError(t, err)

	httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
	httpReq.Header.Set(HeaderContentType, JSONContentType)

	rec := httptest.NewRecorder()
	hdr.ServeHTTP(rec, httpReq)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, []string{
		"interceptor1-before",
		"interceptor2-before",
		"interceptor2-after",
		"interceptor1-after",
	}, callOrder)
}

func TestServeMux_InterceptorReturnsError(t *testing.T) {
	errorInterceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		return nil, status.Error(codes.PermissionDenied, "interceptor denied")
	}

	hdr := NewHandler(ChainUnaryInterceptor(errorInterceptor))
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	req := &testdatav1.EchoRequest{Message: "should fail"}
	body, err := JSONMarshalOptions.Marshal(req)
	require.NoError(t, err)

	httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
	httpReq.Header.Set(HeaderContentType, JSONContentType)

	rec := httptest.NewRecorder()
	hdr.ServeHTTP(rec, httpReq)

	// Should return error status
	assert.Equal(t, http.StatusForbidden, rec.Code)

	// Verify error response body
	respBody := rec.Body.String()
	t.Logf("Interceptor error response body: %s", respBody)
	assert.Contains(t, respBody, "interceptor denied")
}

func TestServeMux_InvalidRequestBody(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	tests := []struct {
		name        string
		contentType string
		body        []byte
	}{
		{
			name:        "Invalid JSON",
			contentType: JSONContentType,
			body:        []byte(`{invalid json`),
		},
		{
			name:        "Invalid Proto",
			contentType: ProtoContentType,
			body:        []byte{0xff, 0xff, 0xff}, // Invalid protobuf
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(tt.body))
			httpReq.Header.Set(HeaderContentType, tt.contentType)

			rec := httptest.NewRecorder()
			hdr.ServeHTTP(rec, httpReq)

			// Should return error for invalid body
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

			// Print response body for debugging
			respBody := rec.Body.String()
			t.Logf("Invalid request body error response: %s", respBody)
		})
	}
}

func TestServeMux_EmptyRequestBody(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	// Empty body should be valid (empty proto message)
	httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader([]byte{}))
	httpReq.Header.Set(HeaderContentType, ProtoContentType)

	rec := httptest.NewRecorder()
	hdr.ServeHTTP(rec, httpReq)

	assert.Equal(t, http.StatusOK, rec.Code)

	resp := &testdatav1.EchoResponse{}
	err := proto.Unmarshal(rec.Body.Bytes(), resp)
	require.NoError(t, err)
	assert.Equal(t, "Echo: ", resp.Message)
}

func TestServeMux_LargeRequestBody(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	// Create a large message
	largeMessage := make([]byte, 1024*1024) // 1MB
	for i := range largeMessage {
		largeMessage[i] = 'a'
	}

	req := &testdatav1.EchoRequest{Message: string(largeMessage)}
	body, err := proto.Marshal(req)
	require.NoError(t, err)

	httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
	httpReq.Header.Set(HeaderContentType, ProtoContentType)

	rec := httptest.NewRecorder()
	hdr.ServeHTTP(rec, httpReq)

	assert.Equal(t, http.StatusOK, rec.Code)

	resp := &testdatav1.EchoResponse{}
	err = proto.Unmarshal(rec.Body.Bytes(), resp)
	require.NoError(t, err)
	assert.True(t, len(resp.Message) > 1024*1024)
}

func TestServeMux_ConcurrentRequests(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	const numRequests = 100
	done := make(chan bool, numRequests)

	for i := 0; i < numRequests; i++ {
		go func(idx int) {
			req := &testdatav1.EchoRequest{Message: "concurrent"}
			body, _ := JSONMarshalOptions.Marshal(req)

			resp, err := http.Post(
				server.URL+"/testdata.v1.EchoService/Echo",
				JSONContentType,
				bytes.NewReader(body),
			)
			if err != nil {
				done <- false
				return
			}
			defer resp.Body.Close()

			respBody, _ := io.ReadAll(resp.Body)
			echoResp := &testdatav1.EchoResponse{}
			if err := JSONUnmarshalOptions.Unmarshal(respBody, echoResp); err != nil {
				done <- false
				return
			}

			done <- resp.StatusCode == http.StatusOK && echoResp.Message == "Echo: concurrent"
		}(i)
	}

	successCount := 0
	for i := 0; i < numRequests; i++ {
		if <-done {
			successCount++
		}
	}

	assert.Equal(t, numRequests, successCount)
}

func TestServeMux_NotFound(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	httpReq := httptest.NewRequest(http.MethodPost, "/nonexistent.Service/Method", nil)
	httpReq.Header.Set(HeaderContentType, JSONContentType)

	rec := httptest.NewRecorder()
	hdr.ServeHTTP(rec, httpReq)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestServeMux_DuplicateRegistration(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	// Registering the same service again should panic
	assert.Panics(t, func() {
		testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})
	})
}

func TestHandler_WithWriteResponseHook(t *testing.T) {
	var hookCalled bool
	var capturedResp proto.Message

	hdr := NewHandler(
		WithWriteResponseHook(func(next WriteResponseFunc) WriteResponseFunc {
			return func(ctx context.Context, input *WriteResponseInput) (*WriteResponseOutput, error) {
				hookCalled = true
				capturedResp = input.Response
				input.W.Header().Set("X-Response-Hook", "called")
				return next(ctx, input)
			}
		}),
	)
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	t.Run("hook is called on success", func(t *testing.T) {
		hookCalled = false
		capturedResp = nil

		req := &testdatav1.EchoRequest{Message: "hello"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, hookCalled, "hook should be called")
		assert.NotNil(t, capturedResp, "response message should be captured")
		assert.Equal(t, "called", rec.Header().Get("X-Response-Hook"))
	})

	t.Run("hook can override response", func(t *testing.T) {
		customHdr := NewHandler(
			WithWriteResponseHook(func(next WriteResponseFunc) WriteResponseFunc {
				return func(ctx context.Context, input *WriteResponseInput) (*WriteResponseOutput, error) {
					input.W.Header().Set("Content-Type", "text/plain")
					input.W.WriteHeader(http.StatusCreated)
					_, _ = input.W.Write([]byte("custom response"))
					return &WriteResponseOutput{Written: true}, nil
				}
			}),
		)
		testdatav1.RegisterEchoServiceServer(customHdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "hello"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		customHdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.Equal(t, "custom response", rec.Body.String())
	})

	t.Run("hook receives ContentTypeJSON and AcceptJSON", func(t *testing.T) {
		var receivedContentTypeJSON, receivedAcceptJSON bool

		jsonCheckHdr := NewHandler(
			WithWriteResponseHook(func(next WriteResponseFunc) WriteResponseFunc {
				return func(ctx context.Context, input *WriteResponseInput) (*WriteResponseOutput, error) {
					receivedContentTypeJSON = input.ContentTypeJSON
					receivedAcceptJSON = input.AcceptJSON
					return next(ctx, input)
				}
			}),
		)
		testdatav1.RegisterEchoServiceServer(jsonCheckHdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "hello"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)
		httpReq.Header.Set(HeaderAccept, ProtoContentType)

		rec := httptest.NewRecorder()
		jsonCheckHdr.ServeHTTP(rec, httpReq)

		assert.True(t, receivedContentTypeJSON, "ContentTypeJSON should be true for JSON request")
		assert.False(t, receivedAcceptJSON, "AcceptJSON should be false when Accept is proto")
	})
}

// greetServer implements testdatav1.GreetServiceServer for testing.
type greetServer struct {
	testdatav1.UnimplementedGreetServiceServer
}

func (s *greetServer) Greet(ctx context.Context, req *testdatav1.GreetRequest) (*testdatav1.GreetResponse, error) {
	return &testdatav1.GreetResponse{Greeting: "Hello, " + req.Name + "!"}, nil
}

func TestServeMux_MultipleServices(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})
	testdatav1.RegisterGreetServiceServer(hdr, &greetServer{})

	t.Run("EchoService works", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		resp := &testdatav1.EchoResponse{}
		err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("GreetService works", func(t *testing.T) {
		req := &testdatav1.GreetRequest{Name: "World"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.GreetService/Greet", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		resp := &testdatav1.GreetResponse{}
		err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Hello, World!", resp.Greeting)
	})
}

// normalizeEchoServer uses normalize functions in its handler.
type normalizeEchoServer struct {
	testdatav1.UnimplementedEchoServiceServer
}

func (s *normalizeEchoServer) Echo(ctx context.Context, req *testdatav1.EchoRequest) (*testdatav1.EchoResponse, error) {
	// Test MustCallMetaFromContext
	callMeta := normalize.MustCallMetaFromContext(ctx)
	if callMeta == nil {
		return nil, errors.New("callMeta is nil")
	}

	// Test MustHTTPMetaFromContext to access request headers
	httpMeta := normalize.MustHTTPMetaFromContext(ctx)
	var reqHeaderValue string
	if httpMeta != nil && httpMeta.R != nil {
		reqHeaderValue = httpMeta.R.Header.Get("X-Request-Header")
	}

	// Test MustSetHeader: set a fixed header and echo request header back
	normalize.MustSetHeader(ctx, "X-Custom-Header", "custom-value")
	if reqHeaderValue != "" {
		normalize.MustSetHeader(ctx, "X-Request-Header-Echo", reqHeaderValue)
	}

	// Test metadata.FromIncomingContext to access headers from gRPC metadata
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get("x-request-header")
		if len(vals) > 0 && vals[0] != "" {
			normalize.MustSetHeader(ctx, "X-Request-Header-FromMetadata", vals[0])
		}
	}

	return &testdatav1.EchoResponse{Message: "Echo: " + req.Message + " (method: " + callMeta.FullMethod + ")"}, nil
}

func TestServeMux_WithNormalizeInterceptor(t *testing.T) {
	// Use normalize.UnaryServerInterceptor for gRPC-style interceptor
	// and wrap with normalize.HTTPMiddleware for HTTP header support
	hdr := NewHandler(ChainUnaryInterceptor(normalize.GRPCUnaryServerInterceptor()))
	testdatav1.RegisterEchoServiceServer(hdr, &normalizeEchoServer{})

	t.Run("MustCallMetaFromContext works", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)

		resp := &testdatav1.EchoResponse{}
		err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)

		// Verify that callMeta.FullMethod was correctly set
		assert.Equal(t, "Echo: test (method: /testdata.v1.EchoService/Echo)", resp.Message)
	})

	t.Run("MustSetHeader works", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)

		// Verify that MustSetHeader correctly set the HTTP response header
		assert.Equal(t, "custom-value", rec.Header().Get("X-Custom-Header"))
	})

	t.Run("Request header is accessible in Echo via HTTPMeta", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)
		httpReq.Header.Set("X-Request-Header", "request-value")

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		// Echoed header should be set by MustSetHeader inside Echo using value from request header
		assert.Equal(t, "request-value", rec.Header().Get("X-Request-Header-Echo"))
	})

	t.Run("Request header is accessible via metadata.FromIncomingContext", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)
		httpReq.Header.Set("X-Request-Header", "request-meta-value")

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		// Header set via metadata.FromIncomingContext should be present
		assert.Equal(t, "request-meta-value", rec.Header().Get("X-Request-Header-FromMetadata"))
	})
}

func TestHandler_WithWriteErrorHook(t *testing.T) {
	var hookCalled bool
	var capturedErr error

	hdr := NewHandler(
		WithWriteErrorHook(func(next WriteErrorFunc) WriteErrorFunc {
			return func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
				hookCalled = true
				capturedErr = input.Error

				// Add custom header before calling next
				input.W.Header().Set("X-Error-Hook", "called")

				return next(ctx, input)
			}
		}),
	)
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	t.Run("hook is called on error", func(t *testing.T) {
		hookCalled = false
		capturedErr = nil

		req := &testdatav1.EchoWithErrorRequest{
			Message:   "test",
			ErrorType: testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
		}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/EchoWithError", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.True(t, hookCalled, "hook should be called")
		assert.NotNil(t, capturedErr, "error should be captured")
		assert.Equal(t, "called", rec.Header().Get("X-Error-Hook"))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("hook can modify error response", func(t *testing.T) {
		customHdr := NewHandler(
			WithWriteErrorHook(func(next WriteErrorFunc) WriteErrorFunc {
				return func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
					// Custom error response instead of default
					input.W.Header().Set("Content-Type", "text/plain")
					input.W.WriteHeader(http.StatusTeapot)
					_, _ = input.W.Write([]byte("custom error: " + input.Error.Error()))
					return &WriteErrorOutput{Written: true}, nil
				}
			}),
		)
		testdatav1.RegisterEchoServiceServer(customHdr, &echoServer{})

		req := &testdatav1.EchoWithErrorRequest{
			Message:   "test",
			ErrorType: testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
		}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/EchoWithError", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		customHdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusTeapot, rec.Code)
		assert.Equal(t, "text/plain", rec.Header().Get("Content-Type"))
		assert.Contains(t, rec.Body.String(), "custom error:")
	})

	t.Run("hook receives ContentTypeJSON and AcceptJSON", func(t *testing.T) {
		var receivedContentTypeJSON, receivedAcceptJSON bool

		jsonCheckHdr := NewHandler(
			WithWriteErrorHook(func(next WriteErrorFunc) WriteErrorFunc {
				return func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
					receivedContentTypeJSON = input.ContentTypeJSON
					receivedAcceptJSON = input.AcceptJSON
					return next(ctx, input)
				}
			}),
		)
		testdatav1.RegisterEchoServiceServer(jsonCheckHdr, &echoServer{})

		req := &testdatav1.EchoWithErrorRequest{
			Message:   "test",
			ErrorType: testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
		}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/EchoWithError", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)
		httpReq.Header.Set(HeaderAccept, ProtoContentType)

		rec := httptest.NewRecorder()
		jsonCheckHdr.ServeHTTP(rec, httpReq)

		assert.True(t, receivedContentTypeJSON, "ContentTypeJSON should be true for JSON request")
		assert.False(t, receivedAcceptJSON, "AcceptJSON should be false when Accept is proto")
	})
}

// customError implements WriteErrorIface for testing.
type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func (e *customError) WriteError(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
	input.W.Header().Set("Content-Type", "application/x-custom-error")
	input.W.Header().Set("X-Custom-Error", "true")
	input.W.WriteHeader(http.StatusUnprocessableEntity)
	_, err := input.W.Write([]byte(`{"custom_error":"` + e.message + `"}`))
	return &WriteErrorOutput{Written: true}, err
}

// customErrorServer returns customError for testing WriteErrorIface.
type customErrorServer struct {
	testdatav1.UnimplementedEchoServiceServer
}

func (s *customErrorServer) Echo(ctx context.Context, req *testdatav1.EchoRequest) (*testdatav1.EchoResponse, error) {
	return nil, &customError{message: "custom error message"}
}

func TestHandler_WriteErrorIface(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &customErrorServer{})

	t.Run("error implementing WriteErrorIface handles its own response", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, "application/x-custom-error", rec.Header().Get("Content-Type"))
		assert.Equal(t, "true", rec.Header().Get("X-Custom-Error"))
		assert.Contains(t, rec.Body.String(), "custom error message")
	})

	t.Run("hook wraps default handler which delegates to WriteErrorIface", func(t *testing.T) {
		var hookCalled bool

		hdrWithHook := NewHandler(
			WithWriteErrorHook(func(next WriteErrorFunc) WriteErrorFunc {
				return func(ctx context.Context, input *WriteErrorInput) (*WriteErrorOutput, error) {
					hookCalled = true
					return next(ctx, input)
				}
			}),
		)
		testdatav1.RegisterEchoServiceServer(hdrWithHook, &customErrorServer{})

		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, JSONContentType)

		rec := httptest.NewRecorder()
		hdrWithHook.ServeHTTP(rec, httpReq)

		// Hook should still be called (it wraps the default handler)
		assert.True(t, hookCalled, "hook should be called")
		// But WriteErrorIface should handle the response
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Equal(t, "application/x-custom-error", rec.Header().Get("Content-Type"))
	})
}

func TestHandler_WithDefaultContentType(t *testing.T) {
	t.Run("default JSON content type affects request parsing", func(t *testing.T) {
		hdr := NewHandler(WithDefaultContentType(JSONContentType))
		testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := JSONMarshalOptions.Marshal(req)
		require.NoError(t, err)

		// No Content-Type header, should use default JSON for parsing
		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		// No Accept header, response should follow request format (JSON)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, JSONContentType, rec.Header().Get(HeaderContentType))

		resp := &testdatav1.EchoResponse{}
		err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("default proto content type affects request parsing", func(t *testing.T) {
		hdr := NewHandler(WithDefaultContentType(ProtoContentType))
		testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := proto.Marshal(req)
		require.NoError(t, err)

		// No Content-Type header, should use default proto for parsing
		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		// No Accept header, response should follow request format (proto)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ProtoContentType, rec.Header().Get(HeaderContentType))

		resp := &testdatav1.EchoResponse{}
		err = proto.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("explicit Content-Type overrides default", func(t *testing.T) {
		hdr := NewHandler(WithDefaultContentType(JSONContentType))
		testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := proto.Marshal(req)
		require.NoError(t, err)

		// Explicit Content-Type proto overrides default JSON
		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderContentType, ProtoContentType)
		// No Accept header, response should follow request format (proto)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, ProtoContentType, rec.Header().Get(HeaderContentType))

		resp := &testdatav1.EchoResponse{}
		err = proto.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("Accept header overrides Content-Type", func(t *testing.T) {
		hdr := NewHandler(WithDefaultContentType(ProtoContentType))
		testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

		req := &testdatav1.EchoRequest{Message: "test"}
		body, err := proto.Marshal(req)
		require.NoError(t, err)

		// Default proto, no Content-Type (so uses proto), but Accept JSON
		httpReq := httptest.NewRequest(http.MethodPost, "/testdata.v1.EchoService/Echo", bytes.NewReader(body))
		httpReq.Header.Set(HeaderAccept, JSONContentType)

		rec := httptest.NewRecorder()
		hdr.ServeHTTP(rec, httpReq)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, JSONContentType, rec.Header().Get(HeaderContentType))

		resp := &testdatav1.EchoResponse{}
		err = JSONUnmarshalOptions.Unmarshal(rec.Body.Bytes(), resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("invalid content type panics", func(t *testing.T) {
		assert.Panics(t, func() {
			NewHandler(WithDefaultContentType("text/plain"))
		})
	})
}
