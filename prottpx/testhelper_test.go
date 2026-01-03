package prottpx

import (
	"context"
	"net/http/httptest"
	"testing"

	testdatav1 "github.com/qor5/x/v3/prottpx/gen/testdata/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestClient_Call(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	client := WrapClient(server.Client(), server.URL)

	t.Run("successful call", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "hello"}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: hello", resp.Message)
	})

	t.Run("empty message", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: ""}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: ", resp.Message)
	})
}

func TestClient_CallWithError(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	client := WrapClient(server.Client(), server.URL)

	t.Run("service returns error", func(t *testing.T) {
		req := &testdatav1.EchoWithErrorRequest{
			Message:   "test",
			ErrorType: testdatav1.ErrorType_ERROR_TYPE_GRPC_STATUS,
		}
		resp := &testdatav1.EchoWithErrorResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/EchoWithError", req, resp)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "HTTP error: status=400")
	})

	t.Run("not found", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "hello"}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/nonexistent.Service/Method", req, resp)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "HTTP error: status=404")
	})
}

func TestClient_WithAuthorization(t *testing.T) {
	var capturedAuthHeader string

	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			vals := md.Get("authorization")
			if len(vals) > 0 {
				capturedAuthHeader = vals[0]
			}
		}
		return handler(ctx, req)
	}

	hdr := NewHandler(ChainUnaryInterceptor(interceptor))
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	t.Run("with authorization token", func(t *testing.T) {
		capturedAuthHeader = ""

		client := WrapClient(server.Client(), server.URL).WithAuthorization("Bearer test-token")

		req := &testdatav1.EchoRequest{Message: "hello"}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
		require.NoError(t, err)
		assert.Equal(t, "Bearer test-token", capturedAuthHeader)
	})

	t.Run("without authorization token", func(t *testing.T) {
		capturedAuthHeader = ""

		client := WrapClient(server.Client(), server.URL)

		req := &testdatav1.EchoRequest{Message: "hello"}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
		require.NoError(t, err)
		assert.Empty(t, capturedAuthHeader)
	})
}

func TestClient_WithInterceptor(t *testing.T) {
	var callOrder []string

	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		callOrder = append(callOrder, "interceptor")
		return handler(ctx, req)
	}

	hdr := NewHandler(ChainUnaryInterceptor(interceptor))
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	client := WrapClient(server.Client(), server.URL)

	callOrder = nil
	req := &testdatav1.EchoRequest{Message: "hello"}
	resp := &testdatav1.EchoResponse{}

	err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
	require.NoError(t, err)
	assert.Equal(t, []string{"interceptor"}, callOrder)
}

func TestClient_MultipleServices(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})
	testdatav1.RegisterGreetServiceServer(hdr, &greetServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	client := WrapClient(server.Client(), server.URL)

	t.Run("call EchoService", func(t *testing.T) {
		req := &testdatav1.EchoRequest{Message: "test"}
		resp := &testdatav1.EchoResponse{}

		err := client.Call(context.Background(), "/testdata.v1.EchoService/Echo", req, resp)
		require.NoError(t, err)
		assert.Equal(t, "Echo: test", resp.Message)
	})

	t.Run("call GreetService", func(t *testing.T) {
		req := &testdatav1.GreetRequest{Name: "World"}
		resp := &testdatav1.GreetResponse{}

		err := client.Call(context.Background(), "/testdata.v1.GreetService/Greet", req, resp)
		require.NoError(t, err)
		assert.Equal(t, "Hello, World!", resp.Greeting)
	})
}

func TestClient_ContextCancellation(t *testing.T) {
	hdr := NewHandler()
	testdatav1.RegisterEchoServiceServer(hdr, &echoServer{})

	server := httptest.NewServer(hdr)
	defer server.Close()

	client := WrapClient(server.Client(), server.URL)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req := &testdatav1.EchoRequest{Message: "hello"}
	resp := &testdatav1.EchoResponse{}

	err := client.Call(ctx, "/testdata.v1.EchoService/Echo", req, resp)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "context canceled")
}
