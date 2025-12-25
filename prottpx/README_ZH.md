# prottpx

`prottpx` 是一个用于将 gRPC unary 服务以 HTTP 形式暴露的轻量级适配层。

它基于 [`github.com/theplant/prottp`](https://github.com/theplant/prottp) 的设计思路，
在此基础上做了一些扩展和约定，以便更好地集成到 `qor5/x` 生态中。

## 特性

- **Handler 模式**：

  - 提供 `Handler` 类型，同时实现 `http.Handler` 与 `grpc.ServiceRegistrar` 接口。
  - 可直接注册 gRPC service descriptor，然后挂到任意 HTTP 路由上。

- **内容协商**：

  - 支持 `application/json` 与 `application/proto` 两种编码。
  - 请求体根据 `Content-Type` 头解码。如果未指定 Content-Type，
    将使用通过 `WithDefaultContentType` 配置的默认内容类型（默认为 `application/proto`）。
  - 响应格式由 `Accept` 头决定。如果没有 Accept 头，
    将跟随请求的 Content-Type 格式。
  - 使用 `WithDefaultContentType` 可以为没有 Content-Type 头的请求
    配置默认内容类型。

- **normalize 集成**：

  - 内部自适应挂载 `normalize.HTTPMiddleware` 实现在 gRPC handler 中方便地访问 `HTTPMeta` 以及使用 normalize 的方法。
  - 配合 `normalize.UnaryServerInterceptor` 使用，可在 gRPC handler 中方便地访问 `CallMeta`。

- **错误处理扩展**：
  - 默认使用 `connect.ErrorWriter` 输出 JSON 错误响应（与 `connect-es` 结合使用）。
  - 通过 `WithWriteErrorHook` 支持自定义错误写入 hook：
    - hook 接收 `WriteErrorInput`，可访问原始 `error`、内容协商信息、`ConnectErrorWriter` 等。
    - 可选择调用或不调用默认逻辑，实现自定义错误格式/状态码。
  - 通过 `WriteErrorIface`，允许错误类型自行实现 `WriteError`，完全控制 HTTP 响应格式。

## 快速示例

```go
// 创建 Handler 并注册 gRPC 服务
h := prottpx.NewHandler(
    prottpx.ChainUnaryInterceptor(
        normalize.GRPCUnaryServerInterceptor(),
    ),
)

testdatav1.RegisterEchoServiceServer(h, echoServer)

// 作为 http.Handler 使用
http.Handle("/api/", http.StripPrefix("/api", h))
log.Fatal(http.ListenAndServe(":8080", nil))
```

如需了解更多用法细节，可直接查阅本包源码与测试用例。
