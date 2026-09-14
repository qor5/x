# 从 statusx 迁移到 httperrors

## Q: 每个服务都要单独写 middleware 吗？

**不需要。** `httperrors` 包已经提供了通用的 `ErrorMiddleware` 和便捷函数 `NewErrorMiddleware`，开箱即用。
服务只需要在启动时注册一次，不需要自己实现 panic recovery、错误翻译、JSON 序列化等逻辑。

```go
// 一行搞定，不需要每个服务单独写 middleware
mux.Use(httperrors.NewErrorMiddleware(ib))
```

如果需要自定义错误写入逻辑（如日志、监控），可以通过 Hook 扩展：

```go
conf := &httperrors.HTTPErrorMiddlewareConfig{I18N: ib}
conf = conf.WithHTTPWriteErrorHook(func(next httperrors.HTTPWriteErrorFunc) httperrors.HTTPWriteErrorFunc {
    return func(ctx context.Context, input *httperrors.HTTPWriteErrorInput) (*httperrors.HTTPWriteErrorOutput, error) {
        slog.ErrorContext(ctx, "API error", "error", input.Err, "path", input.R.URL.Path)
        return next(ctx, input)
    }
})
mux.Use(httperrors.ErrorMiddleware(conf))
```

---

## 旧 statusx vs 新 httperrors：服务接入对比

### 旧方式：statusx（Connect RPC / gRPC）

旧的 `statusx` 服务通常走 Connect RPC 协议，接入链路涉及多个组件：

```go
package main

import (
    "net/http"
    "strings"

    "connectrpc.com/connect"
    "github.com/qor5/x/v3/connectx"
    "github.com/qor5/x/v3/i18nx"
)

//go:embed i18n/messages.csv
var messagesCSV string

func main() {
    // 1. 初始化 i18n
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))

    // 2. 创建 Connect mux（内部自动接入 statusx 的 error 处理链）
    //    内部做了很多事：
    //    - statusx.UnaryConnectInterceptor (错误翻译 + Connect 错误转换)
    //    - statusx.NewVProtoHTTPErrorWriter (VProto 格式兼容)
    //    - EnsureConnectError 逻辑 (双路径：Connect 格式 vs VProto 格式)
    //    - rejectGRPCProtocol 中间件
    mux := connectx.NewMux(ib)

    // 3. 注册 Connect RPC handler
    mux.Handle(
        userv1connect.NewUserServiceHandler,
        orderv1connect.NewOrderServiceHandler,
    )

    // 4. 启动
    http.ListenAndServe(":8080", mux)
}
```

**业务层：**

```go
// 服务实现 Connect RPC 接口
func (s *UserServiceServer) GetUser(
    ctx context.Context,
    req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
    user, err := s.repo.FindByID(ctx, req.Msg.Id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        // 使用 gRPC codes + protobuf ErrorReason
        return nil, statusx.New(codes.NotFound, "NOT_FOUND", "user not found").Err()
    }
    return connect.NewResponse(&userv1.GetUserResponse{User: toProto(user)}), nil
}
```

**涉及的依赖和组件：**

| 组件                                        | 说明                                                 |
| ------------------------------------------- | ---------------------------------------------------- |
| `connectrpc.com/connect`                    | Connect RPC 框架                                     |
| `google.golang.org/grpc/codes`              | gRPC status codes                                    |
| `google.golang.org/grpc/status`             | gRPC status 转换                                     |
| `google.golang.org/protobuf`                | Protobuf 序列化                                      |
| `google.golang.org/genproto/.../errdetails` | gRPC error details (ErrorInfo, BadRequest)           |
| `statusx/proto/status/v1`                   | 自定义 protobuf (ErrorReason, Localized, BadRequest) |
| `statusx.UnaryConnectInterceptor`           | Connect 拦截器                                       |
| `statusx.ConvertToConnectError`             | StatusError → connect.Error 转换                     |
| `statusx.WriteConnectErrorOnly`             | Connect 错误写入                                     |
| `statusx.NewVProtoHTTPErrorWriter`          | VProto 兼容写入                                      |
| `statusx.EnsureConnectError`                | 双路径错误格式选择                                   |

**错误响应格式** (Connect 协议)：

```json
{
  "code": "not_found",
  "message": "user not found",
  "details": [
    {
      "type": "google.rpc.ErrorInfo",
      "value": "...(base64 protobuf)...",
      "debug": { "reason": "NOT_FOUND", "domain": "", "metadata": {} }
    }
  ]
}
```

---

### 新方式：httperrors（纯 HTTP）

```go
package main

import (
    "net/http"
    "strings"

    "github.com/qor5/x/v3/httperrors"
    "github.com/qor5/x/v3/i18nx"
)

//go:embed i18n/messages.csv
var messagesCSV string

func main() {
    // 1. 初始化 i18n（完全一样）
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))

    // 2. 创建路由 + 注册通用 error middleware（一行）
    mux := http.NewServeMux()
    handler := httperrors.NewErrorMiddleware(ib)(mux)

    // 3. 注册普通 HTTP handler
    mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
    mux.HandleFunc("POST /orders", orderHandler.Create)

    // 4. 启动
    http.ListenAndServe(":8080", handler)
}
```

**业务层：**

```go
// 服务层：标准 Go error 返回
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to query user")
    }
    if user == nil {
        // 直接用 HTTP status code + reason 字符串，无 gRPC 依赖
        return nil, httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")
    }
    return user, nil
}

// HTTP handler 层：panic 将 error 交给 middleware
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

**涉及的依赖和组件：**

| 组件                         | 说明                                       |
| ---------------------------- | ------------------------------------------ |
| `net/http`                   | Go 标准库                                  |
| `httperrors.ErrorMiddleware` | 通用中间件（panic recovery + i18n + JSON） |
| `i18nx`                      | i18n 翻译（与旧版相同）                    |

**没有任何 gRPC / protobuf / Connect 依赖。**

**错误响应格式** (纯 JSON)：

```json
{
  "code": "NOT_FOUND",
  "message": "user not found"
}
```

---

## API 返回格式对比

statusx 有**两条错误响应路径**（Connect 协议 和 VProto 协议），httperrors 只有一条纯 JSON 路径。

### 场景 1：简单错误（无 i18n）

**Go 代码：**

```go
// statusx
statusx.New(codes.NotFound, "NOT_FOUND", "user not found")

// httperrors
httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")
```

**statusx — Connect 路径** (`x-ensure-connect-error: true`)：

```
HTTP/1.1 404 Not Found
Content-Type: application/json
```

```json
{
  "code": "not_found",
  "message": "user not found",
  "details": [
    {
      "type": "google.rpc.ErrorInfo",
      "value": "CglOT1RfRk9VTkQ=",
      "debug": { "reason": "NOT_FOUND", "domain": "", "metadata": {} }
    }
  ]
}
```

> `code` 是 gRPC code 的小写字符串（`not_found`），`reason` 藏在 `details[].debug.reason` 里，`value` 是 base64 编码的 protobuf。

**statusx — VProto 路径**（默认前端请求）：

```
HTTP/1.1 404 Not Found
Content-Type: application/json
```

```json
{
  "code": "NOT_FOUND",
  "msg": "user not found",
  "defaultViewMsg": "user not found",
  "fieldViolations": []
}
```

> VProto 使用 `msg` 而非 `message`，使用 `defaultViewMsg` 作为展示消息。无翻译时 `defaultViewMsg == msg`。

**httperrors**：

```
HTTP/1.1 404 Not Found
Content-Type: application/json
```

```json
{
  "code": "NOT_FOUND",
  "message": "user not found"
}
```

> 无 `details`、无 protobuf、无冗余字段。`metadata`、`fieldViolations`、`localizedMessage` 为空时不出现。

---

### 场景 2：简单错误 + i18n（中文）

**Go 代码（业务层相同，翻译由 middleware 自动完成）：**

```go
// 请求 Header: x-selected-language: zh 或 Accept-Language: zh

// statusx
statusx.New(codes.NotFound, "NOT_FOUND", "user not found")

// httperrors
httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")
```

**statusx — Connect 路径**：

```json
{
  "code": "not_found",
  "message": "user not found",
  "details": [
    {
      "type": "google.rpc.ErrorInfo",
      "value": "CglOT1RfRk9VTkQ=",
      "debug": { "reason": "NOT_FOUND", "domain": "", "metadata": {} }
    },
    {
      "type": "google.rpc.LocalizedMessage",
      "value": "CgJ6aBLJiaa+vuWIsA==",
      "debug": { "locale": "zh", "message": "未找到" }
    }
  ]
}
```

> 翻译结果作为 `LocalizedMessage` detail 追加到 `details` 数组中，`message` **保持原始不变**。前端需要遍历 `details` 数组并找到 `type == "google.rpc.LocalizedMessage"` 的条目。

**statusx — VProto 路径**：

```json
{
  "code": "NOT_FOUND",
  "msg": "user not found",
  "defaultViewMsg": "未找到",
  "fieldViolations": []
}
```

> `msg` 保持原始英文，`defaultViewMsg` 为翻译后的中文。前端直接取 `defaultViewMsg` 展示。

**httperrors**：

```json
{
  "code": "NOT_FOUND",
  "message": "user not found",
  "localizedMessage": "未找到"
}
```

> `message` 保持原始不变，翻译结果放在 `localizedMessage` 中。前端用 `localizedMessage || message` 展示。

---

### 场景 3：字段校验错误（无 i18n）

**Go 代码：**

```go
// statusx
fv1 := statusx.NewFieldViolation("email", "REQUIRED", "email is required")
fv2 := statusx.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters")
statusx.BadRequest(fv1, fv2)

 // httperrors
 fv1 := httperrors.NewFieldViolation("email", "REQUIRED", "email is required")
 fv2 := httperrors.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters")
 httperrors.BadRequest(fv1, fv2)
```

**statusx — Connect 路径**：

```json
{
  "code": "invalid_argument",
  "message": "invalid argument",
  "details": [
    {
      "type": "google.rpc.ErrorInfo",
      "value": "...",
      "debug": { "reason": "INVALID_ARGUMENT", "domain": "", "metadata": {} }
    },
    {
      "type": "google.rpc.BadRequest",
      "value": "...(base64 protobuf)...",
      "debug": {
        "fieldViolations": [
          {
            "field": "email",
            "description": "email is required",
            "reason": "REQUIRED"
          },
          {
            "field": "password",
            "description": "password must be at least 8 characters",
            "reason": "TOO_SHORT"
          }
        ]
      }
    }
  ]
}
```

> 字段违规信息藏在 `details[]` 中的 `BadRequest` detail 里，需要解析 protobuf 或读 `debug` 字段。

**statusx — VProto 路径**：

```json
{
  "code": "INVALID_ARGUMENT",
  "msg": "invalid argument",
  "defaultViewMsg": "invalid argument",
  "fieldViolations": [
    {
      "field": "email",
      "code": "REQUIRED",
      "msg": "email is required",
      "defaultViewMsg": "email is required"
    },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "msg": "password must be at least 8 characters",
      "defaultViewMsg": "password must be at least 8 characters"
    }
  ]
}
```

> VProto 路径结构比 Connect 路径更直观，但使用 `msg` / `defaultViewMsg` 命名。

**httperrors**：

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "email",
      "code": "REQUIRED",
      "message": "email is required"
    },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "message": "password must be at least 8 characters"
    }
  ]
}
```

> 结构与 VProto 类似但使用标准 `message` 命名。`localizedMessage` 为空时不出现。

---

### 场景 4：字段校验错误 + i18n（中文）

**statusx — Connect 路径**：

```json
{
  "code": "invalid_argument",
  "message": "invalid argument",
  "details": [
    {
      "type": "google.rpc.ErrorInfo",
      "value": "...",
      "debug": { "reason": "INVALID_ARGUMENT" }
    },
    {
      "type": "google.rpc.LocalizedMessage",
      "value": "...",
      "debug": { "locale": "zh", "message": "参数无效" }
    },
    {
      "type": "google.rpc.BadRequest",
      "value": "...",
      "debug": {
        "fieldViolations": [
          {
            "field": "email",
            "description": "email is required",
            "reason": "REQUIRED",
            "localizedMessage": { "locale": "zh", "message": "必填" }
          },
          {
            "field": "password",
            "description": "password must be at least 8 characters",
            "reason": "TOO_SHORT",
            "localizedMessage": { "locale": "zh", "message": "太短" }
          }
        ]
      }
    }
  ]
}
```

> Connect 路径的翻译信息分散在多个 detail 中，前端解析成本高。

**statusx — VProto 路径**：

```json
{
  "code": "INVALID_ARGUMENT",
  "msg": "invalid argument",
  "defaultViewMsg": "参数无效",
  "fieldViolations": [
    {
      "field": "email",
      "code": "REQUIRED",
      "msg": "email is required",
      "defaultViewMsg": "必填"
    },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "msg": "password must be at least 8 characters",
      "defaultViewMsg": "太短"
    }
  ]
}
```

> `msg` 保持原始，`defaultViewMsg` 为翻译后的文本。

**httperrors**：

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "email",
      "code": "REQUIRED",
      "message": "email is required",
      "localizedMessage": "必填"
    },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "message": "password must be at least 8 characters",
      "localizedMessage": "太短"
    }
  ]
}
```

> `message` 保持原始，`localizedMessage` 为翻译后的文本。

---

### 字段命名映射对照

| 语义         | statusx Connect 路径                     | statusx VProto 路径 | httperrors         |
| ------------ | ---------------------------------------- | ------------------- | ------------------ |
| 错误码       | `code`（gRPC 小写，如 `not_found`）      | `code`（大写）      | `code`（大写）     |
| 原始消息     | `message`                                | `msg`               | `message`          |
| 翻译消息     | `details[LocalizedMessage].message`      | `defaultViewMsg`    | `localizedMessage` |
| Reason       | `details[ErrorInfo].debug.reason`        | `code`              | `code`             |
| 字段路径     | `details[BadRequest]...field`            | `field`             | `field`            |
| 字段错误码   | `details[BadRequest]...reason`           | `code`              | `code`             |
| 字段原始描述 | `details[BadRequest]...description`      | `msg`               | `message`          |
| 字段翻译     | `details[BadRequest]...localizedMessage` | `defaultViewMsg`    | `localizedMessage` |
| 附加信息     | `details[ErrorInfo].debug.metadata`      | ❌ 不支持           | `metadata`         |

### 前端解析复杂度对比

| 操作                | statusx Connect 路径                                                    | statusx VProto 路径   | httperrors                              |
| ------------------- | ----------------------------------------------------------------------- | --------------------- | --------------------------------------- |
| 获取错误码          | `err.details.find(d => d.type === "google.rpc.ErrorInfo").debug.reason` | `err.code`            | `err.code`                              |
| 获取展示消息        | 遍历 `details` 找 `LocalizedMessage`，fallback `message`                | `err.defaultViewMsg`  | `err.localizedMessage \|\| err.message` |
| 获取字段违规        | 遍历 `details` 找 `BadRequest`，解析 `fieldViolations`                  | `err.fieldViolations` | `err.fieldViolations`                   |
| 获取字段翻译        | `fv.localizedMessage.message`（嵌套在 details 内）                      | `fv.defaultViewMsg`   | `fv.localizedMessage \|\| fv.message`   |
| 需要理解 protobuf？ | **是**（detail type、base64 value）                                     | 否                    | 否                                      |

---

## 对比总结

| 维度                    | statusx (旧)                             | httperrors (新)                                  |
| ----------------------- | ---------------------------------------- | ------------------------------------------------ |
| **协议**                | Connect RPC / gRPC                       | 纯 HTTP                                          |
| **服务接入**            | `connectx.NewMux(ib)`                    | `httperrors.NewErrorMiddleware(ib)(mux)`         |
| **Handler 签名**        | Connect RPC interface                    | 标准 `http.HandlerFunc`                          |
| **错误创建**            | `statusx.New(codes.NotFound, ...)`       | `httperrors.Error(http.StatusNotFound, ...)`     |
| **状态码类型**          | `codes.Code` (gRPC)                      | `int` (HTTP status code)                         |
| **错误传播**            | `return error` (Connect 拦截器处理)      | `return error` (业务层) + `panic` (handler 边界) |
| **错误序列化**          | Protobuf → Connect JSON (带 details)     | 直接 JSON                                        |
| **Reason**              | Protobuf enum `ErrorReason`              | 字符串常量                                       |
| **i18n 参数**           | `anypb.Any` (protobuf)                   | `[]any` (Go 原生)                                |
| **Field violations**    | `errdetails.BadRequest` (protobuf)       | `FieldViolation` (Go struct)                     |
| **Is() 比较**           | `proto.Equal(grpcStatus.Proto())`        | `httpStatus + reason`                            |
| **gRPC 依赖**           | 是 (codes, status, errdetails, protobuf) | 无                                               |
| **Connect 依赖**        | 是 (connect.Error, ErrorWriter)          | 无                                               |
| **Protobuf 依赖**       | 是 (proto, anypb, errdetails)            | 无                                               |
| **需要自写 middleware** | 否 (connectx.NewMux 内置)                | 否 (NewErrorMiddleware 内置)                     |
| **Hook 扩展**           | `WithHTTPWriteErrorHook`                 | `WithHTTPWriteErrorHook` (同名同模式)            |

### 业务代码改动量

```diff
 // 错误创建
-statusx.New(codes.NotFound, "NOT_FOUND", "user not found")
+httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")

 // 错误包装
-statusx.Wrap(err, codes.Internal, "INTERNAL", "database error")
+httperrors.WrapStatus(err, http.StatusInternalServerError, "database error")

 // 字段校验
 -statusx.BadRequest(fv1, fv2)
 +httperrors.BadRequest(fv1, fv2)

 // 字段违规创建
-statusx.NewFieldViolation("email", "REQUIRED", "email is required")
+httperrors.NewFieldViolation("email", "REQUIRED", "email is required")  // 完全一样

 // i18n
-s.WithLocalized("key", args...)
+s.WithLocalized("key", args...)  // 完全一样

 // 错误解析
-statusx.FromError(err)
+httperrors.FromError(err)
```

**核心变化**：把 `codes.Code` 换成 `http.StatusXxx`，其余 API 形状基本一致。

---

## 实际服务迁移示例（依赖注入场景）

以下基于真实服务代码，说明从 `statusx` + `prottpx` 迁移到 `httperrors` 需要做哪些工作。

### 旧代码：statusx + prottpx + gRPC

```go
var SetupHTTPServing = []any{
    SetupHTTPServeMux,
    SetupHTTPListener,
    SetupProttpHandler,   // ← prottpx：处理 gRPC-over-HTTP，内置 statusx 错误链
    SetupHTTPServer,
    SetupGRPCConn,        // ← 需要 gRPC 连接（用于 health check 等）
}

func SetupHTTPServeMux() *http.ServeMux {
    return http.NewServeMux()
}

func SetupHTTPListener(lc *lifecycle.Lifecycle, conf *Config) (httpx.Listener, error) {
    return httpx.SetupListener(lc, &conf.HTTP)
}

// SetupProttpHandler 创建 prottpx Handler，内部通过 gRPC interceptor + statusx 处理错误
// 通常还会接入 VProto 错误格式兼容（NewVProtoWriteErrorHook）
func SetupProttpHandler(ib *i18nx.I18N) *prottpx.Handler {
    return prottpx.NewHandler(
        prottpx.WithWriteErrorHook(
            prottpx.NewVProtoWriteErrorHook(ib),  // statusx 翻译 + VProto/Connect 双路径
        ),
        // gRPC interceptors...
    )
}

func SetupHTTPServer(
    ctx context.Context,
    lc *lifecycle.Lifecycle,
    listener httpx.Listener,
    conf *Config,
    conn *grpc.ClientConn,         // ← gRPC 连接依赖
    mux *http.ServeMux,
    prottpHandler *prottpx.Handler, // ← prottpx Handler 依赖
) (http.Handler, *http.Server, error) {
    mux.Handle("/", prottpHandler)  // ← 所有请求走 prottpx

    conf.HTTP.Security.CORS.AllowedHeaders = lo.Uniq(
        slices.Concat(
            conf.HTTP.Security.CORS.AllowedHeaders,
            i18nx.AllowHeaders,
            statusx.AllowHeaders,   // ← statusx 专用 header（如 x-ensure-connect-error）
            auth.AllowHeaders,
            challenge.AllowHeaders,
        ),
    )

    handler := hook.Chain(
        server.DefaultMiddleware(kitlog.Default()),
        healthz.HTTPMiddleware(healthz.WithGRPCHealthChecker(conn)), // ← 依赖 gRPC conn
        httpx.NoStore,
        httpx.Security(conf.HTTP.Security),
        auth.Cookieize(conf.Auth.Cookie),
        normalize.HTTPMiddleware,
    )(mux)

    server, err := httpx.SetupServerFactory("http-server", handler)(ctx, lc, &conf.HTTP, listener)
    if err != nil {
        return nil, nil, err
    }
    return handler, server, nil
}
```

**这条链路的关键依赖关系：**

```
SetupHTTPServing
├── SetupHTTPServeMux         → *http.ServeMux
├── SetupHTTPListener         → httpx.Listener
├── SetupProttpHandler        → *prottpx.Handler      ← 依赖 gRPC proto 定义 + statusx
├── SetupGRPCConn             → *grpc.ClientConn      ← gRPC 连接
└── SetupHTTPServer           → http.Handler
    ├── prottpHandler.Handle("/")                     ← gRPC-over-HTTP
    ├── statusx.AllowHeaders                          ← statusx CORS header
    └── healthz.WithGRPCHealthChecker(conn)           ← 依赖 gRPC health check
```

---

### 新代码：httperrors（纯 HTTP）

```go
var SetupHTTPServing = []any{
    SetupHTTPServeMux,
    SetupHTTPListener,
    // SetupProttpHandler,   // ✅ 删除：不再需要 prottpx
    SetupHTTPServer,
    // SetupGRPCConn,        // ✅ 删除：不再需要 gRPC 连接（除非其他地方还用）
}

func SetupHTTPServeMux() *http.ServeMux {
    return http.NewServeMux()
}

func SetupHTTPListener(lc *lifecycle.Lifecycle, conf *Config) (httpx.Listener, error) {
    return httpx.SetupListener(lc, &conf.HTTP)
}

func SetupHTTPServer(
    ctx context.Context,
    lc *lifecycle.Lifecycle,
    listener httpx.Listener,
    conf *Config,
    ib *i18nx.I18N,                // ✅ 替代 gpc.ClientConn + prottpx.Handler，只需 i18n
    mux *http.ServeMux,
) (http.Handler, *http.Server, error) {
    // ✅ 注册 HTTP 路由（替代 prottpHandler.Handle("/")）
    registerRoutes(mux)

    conf.HTTP.Security.CORS.AllowedHeaders = lo.Uniq(
        slices.Concat(
            conf.HTTP.Security.CORS.AllowedHeaders,
            i18nx.AllowHeaders,
            // statusx.AllowHeaders,  // ✅ 删除：不再需要 statusx 专用 header
            auth.AllowHeaders,
            challenge.AllowHeaders,
        ),
    )

    handler := hook.Chain(
        server.DefaultMiddleware(kitlog.Default()),
        // healthz.HTTPMiddleware(healthz.WithGRPCHealthChecker(conn)),  // ✅ 改为 HTTP health check
        healthz.HTTPMiddleware(),
        httperrors.NewErrorMiddleware(ib),   // ✅ 新增：httperrors 通用错误中间件
        httpx.NoStore,
        httpx.Security(conf.HTTP.Security),
        auth.Cookieize(conf.Auth.Cookie),
        normalize.HTTPMiddleware,
    )(mux)

    server, err := httpx.SetupServerFactory("http-server", handler)(ctx, lc, &conf.HTTP, listener)
    if err != nil {
        return nil, nil, err
    }
    return handler, server, nil
}

// ✅ 新增：注册纯 HTTP 路由（替代 prottpx 的 gRPC service 注册）
func registerRoutes(mux *http.ServeMux) {
    userHandler := &UserHandler{...}
    orderHandler := &OrderHandler{...}

    mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
    mux.HandleFunc("POST /users", userHandler.Create)
    mux.HandleFunc("POST /orders", orderHandler.Create)
    // ...
}
```

---

### 迁移清单

#### 1. DI 层（SetupHTTPServing）

| 改动                       | 说明                                                              |
| -------------------------- | ----------------------------------------------------------------- |
| 删除 `SetupProttpHandler`  | 不再需要 prottpx（gRPC-over-HTTP）                                |
| 删除 `SetupGRPCConn`       | 不再需要 gRPC 连接（除非其他地方还用）                            |
| `SetupHTTPServer` 参数变更 | 去掉 `*grpc.ClientConn` 和 `*prottpx.Handler`，加入 `*i18nx.I18N` |

#### 2. 中间件链（hook.Chain）

| 改动                                     | 说明                                                     |
| ---------------------------------------- | -------------------------------------------------------- |
| 新增 `httperrors.NewErrorMiddleware(ib)` | 插入 middleware 链中，处理 panic recovery + i18n + JSON  |
| 删除 `statusx.AllowHeaders`              | 不再需要 `x-ensure-connect-error` 等 statusx 专用 header |
| 调整 health check                        | 如果之前依赖 gRPC health check，改为 HTTP health check   |

#### 3. 路由注册

| 改动                                      | 说明                                                  |
| ----------------------------------------- | ----------------------------------------------------- |
| 从 `prottpHandler.RegisterXxxServer(svc)` | 改为 `mux.HandleFunc("METHOD /path", handler.Method)` |
| protobuf service 定义 → HTTP 路由         | 需要把 gRPC service method 映射为 HTTP endpoint       |

#### 4. 业务代码（Service 层）

| 改动                                                                | 说明                              |
| ------------------------------------------------------------------- | --------------------------------- |
| `statusx.New(codes.X, ...)` → `httperrors.Error(http.StatusX, ...)` | 状态码类型替换                    |
| `statusx.Wrap(...)` → `httperrors.WrapStatus(...)`                  | 错误包装                          |
| `statusx.BadRequest(...)` → `httperrors.BadRequest(...)`            | 字段校验                          |
| `statusx.FromError(...)` → `httperrors.FromError(...)`              | 错误解析                          |
| 返回值类型不变                                                      | 业务层仍然 `return error`，无改动 |

#### 5. Handler 层

| 改动                                              | 说明                                          |
| ------------------------------------------------- | --------------------------------------------- |
| gRPC service interface → `http.HandlerFunc`       | handler 签名变化                              |
| `return nil, err` → `panic(err)`                  | 在 handler 边界用 panic 传播错误给 middleware |
| protobuf request/response → JSON request/response | 序列化格式变化                                |

#### 6. 可删除的依赖

```diff
 go.mod:
-connectrpc.com/connect
-google.golang.org/grpc
-google.golang.org/protobuf
-google.golang.org/genproto
-github.com/qor5/x/v3/statusx
-github.com/qor5/x/v3/prottpx
-github.com/qor5/x/v3/connectx
+github.com/qor5/x/v3/httperrors  (新增)
```

> **注意**：如果服务中其他模块（如内部 gRPC 通信）仍然使用 gRPC/protobuf，则相关依赖不能删除。
> 以上清单仅针对 HTTP API 错误处理部分的迁移。
