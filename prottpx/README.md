# prottpx

`prottpx` is a lightweight adapter for exposing gRPC unary services over HTTP.

It is a variant of [`github.com/theplant/prottp`](https://github.com/theplant/prottp),
with additional extensions and conventions to better integrate into the `qor5/x` ecosystem.

## Features

- **Handler model**

  - Provides a `Handler` type that implements both `http.Handler` and `grpc.ServiceRegistrar`.
  - You can register gRPC service descriptors directly on the handler and mount it under any HTTP route.

- **Content negotiation**

  - Supports both `application/json` and `application/proto` encodings.
  - The request body is decoded according to `Content-Type` header. If no Content-Type is specified,
    it uses the default content type configured via `WithDefaultContentType` (defaults to `application/proto`).
  - The response format is decided by `Accept` header. If no Accept header is present,
    it follows the request's Content-Type format.
  - Use `WithDefaultContentType` to configure the default content type for requests
    without a Content-Type header.

- **`normalize` integration**

  - Internally and adaptively applies `normalize.HTTPMiddleware` so that gRPC handlers can easily
    access `HTTPMeta` and use functions from the `normalize` package.
  - Works together with `normalize.UnaryServerInterceptor` so that gRPC handlers can easily access `CallMeta`.

- **Extensible error handling**

  - By default, uses `connect.ErrorWriter` to produce JSON error responses (designed to work with `connect-es`).
  - `WithWriteErrorHook` allows you to plug in custom error-writing hooks:
    - Hooks receive a `WriteErrorInput` which exposes the original `error`,
      content-negotiation information, `ConnectErrorWriter`, etc.
    - Hooks may choose to call or bypass the default logic, enabling fully customized
      error payloads and status codes.
  - `WriteErrorIface` allows error types to implement `WriteError` themselves and
    take full control over the HTTP error response.

## Quick example

```go
// Create a Handler and register a gRPC service
h := prottpx.NewHandler(
    prottpx.ChainUnaryInterceptor(
        normalize.GRPCUnaryServerInterceptor(),
    ),
)

testdatav1.RegisterEchoServiceServer(h, echoServer)

// Use it as an http.Handler
http.Handle("/api/", http.StripPrefix("/api", h))
log.Fatal(http.ListenAndServe(":8080", nil))
```

For more details, please refer to the source code and tests of this package.
