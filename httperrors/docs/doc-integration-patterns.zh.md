# httperrors 接入模式指南

本文档说明在不同场景下如何接入 `httperrors` 包，以及各种模式的最佳实践。

---

## 三种接入方式

| 方式 | 函数 | 适用场景 | 错误传播 |
| --- | --- | --- | --- |
| **全局中间件** | `ErrorMiddleware` / `NewErrorMiddleware` | 所有 handler 都使用 httperrors | panic |
| **单 handler 包裹** | `WrapHandlerFunc` | mux 中部分 handler 使用 httperrors | panic |
| **显式调用** | `WriteError` / `HandleError` | handler 内部自行处理错误，不依赖 panic | return error / direct handling |

---

## 1. 全局中间件模式

**适用场景**: 整个服务的所有 handler 都使用 `httperrors` 返回错误。

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
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))

    mux := http.NewServeMux()
    mux.HandleFunc("GET /api/users/{id}", userHandler.GetUser)
    mux.HandleFunc("POST /api/orders", orderHandler.Create)

    // 全局中间件：所有 handler 的 panic(error) 都会被捕获并转为 JSON 响应
    handler := httperrors.NewErrorMiddleware(ib)(mux)
    http.ListenAndServe(":8080", handler)
}
```

**业务层**（始终 return error）:

```go
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, httperrors.Error(http.StatusNotFound, httperrors.ReasonNotFound, "user not found")
    }
    return user, nil
}
```

**Handler 层**（panic 传播给中间件）:

```go
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

**优点**: 简洁，handler 代码最少  
**缺点**: 所有 handler 必须遵循 `panic(error)` 约定

---

## 2. 单 Handler 包裹模式

**适用场景**: 一个 mux 上混合了多种 handler，只有部分 handler 使用 `httperrors`。

```go
func main() {
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))
    conf := &httperrors.HTTPErrorMiddlewareConfig{I18N: ib}

    mux := http.NewServeMux()

    mux.HandleFunc("GET /api/users/{id}", httperrors.WrapHandlerFunc(conf, userHandler.GetUser))
    mux.HandleFunc("POST /api/orders", httperrors.WrapHandlerFunc(conf, orderHandler.Create))

    mux.HandleFunc("GET /health", healthHandler)
    mux.HandleFunc("GET /legacy/report", legacyReportHandler)

    http.ListenAndServe(":8080", mux)
}
```

被包裹的 handler 内部仍使用 panic 模式：

```go
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

**优点**: 灵活，逐个 handler 选择是否接入  
**缺点**: 每个 handler 注册时需要额外包裹一层

---

## 3. 显式调用模式

**适用场景**:
- handler 内部想用 `return` 而非 `panic` 处理错误
- 需要在写错误响应前后做额外逻辑（如日志、metrics）
- 与现有框架集成，无法使用 panic 模式

```go
func main() {
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))
    conf := &httperrors.HTTPErrorMiddlewareConfig{I18N: ib}

    mux := http.NewServeMux()
    h := &UserHandler{conf: conf, userService: userService}
    mux.HandleFunc("GET /api/users/{id}", h.GetUser)

    http.ListenAndServe(":8080", mux)
}
```

Handler 内部显式调用 `WriteError`：

```go
type UserHandler struct {
    conf        *httperrors.HTTPErrorMiddlewareConfig
    userService *UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        if werr := httperrors.WriteError(h.conf, w, r, err); werr != nil {
            slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err)
        }
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

如果你只需要兼容旧调用方式，也可以继续使用 `HandleError`：

```go
func (h *UserHandler) GetUserLegacy(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        httperrors.HandleError(h.conf, w, r, err)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

**优点**: 不依赖 panic，完全显式控制，更符合 Go 惯例  
**缺点**: 每个错误处理点都需要显式写出响应并 `return`

---

## 4. 混合模式（推荐的渐进式迁移方案）

**适用场景**: 大型项目中逐步从旧错误处理迁移到 `httperrors`。

```go
func main() {
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))
    conf := &httperrors.HTTPErrorMiddlewareConfig{I18N: ib}

    mux := http.NewServeMux()

    mux.HandleFunc("GET /api/users/{id}", httperrors.WrapHandlerFunc(conf, userHandler.GetUser))
    mux.HandleFunc("POST /api/users", httperrors.WrapHandlerFunc(conf, userHandler.Create))

    mux.HandleFunc("GET /api/orders/{id}", orderHandler.GetOrder)

    mux.HandleFunc("GET /legacy/report", legacyReportHandler)
    mux.HandleFunc("POST /legacy/upload", legacyUploadHandler)

    http.ListenAndServe(":8080", mux)
}
```

---

## 5. 带自定义 Hook 的接入

所有三种方式都支持通过 `HTTPErrorMiddlewareConfig` 配置 hook，用于在错误写入前后做自定义逻辑（如日志、metrics、修改响应）。

```go
conf := &httperrors.HTTPErrorMiddlewareConfig{I18N: ib}
conf = conf.WithHTTPWriteErrorHook(func(next httperrors.HTTPWriteErrorFunc) httperrors.HTTPWriteErrorFunc {
    return func(ctx context.Context, input *httperrors.HTTPWriteErrorInput) (*httperrors.HTTPWriteErrorOutput, error) {
        slog.ErrorContext(ctx, "API error",
            "method", input.R.Method,
            "path", input.R.URL.Path,
            "error", input.Err,
        )
        return next(ctx, input)
    }
})

handler := httperrors.ErrorMiddleware(conf)(mux)
wrapped := httperrors.WrapHandlerFunc(conf, h.GetUser)
if err := httperrors.WriteError(conf, w, r, err); err != nil {
    slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err)
}
httperrors.HandleError(conf, w, r, err)
```

---

## 模式选择决策树

```
你的 mux 上所有 handler 都用 httperrors 吗？
├── 是 → 全局中间件模式 (ErrorMiddleware)
└── 否 → 部分 handler 用 httperrors
         ├── handler 内部可以用 panic 模式吗？
         │   ├── 是 → 单 handler 包裹模式 (WrapHandlerFunc)
         │   └── 否 → 显式调用模式 (WriteError / HandleError)
         └── 正在渐进式迁移？
             └── 是 → 混合模式（WrapHandlerFunc + WriteError/HandleError + 旧代码共存）
```

---

## API 参考

### `ErrorMiddleware(conf) func(http.Handler) http.Handler`

全局中间件，包裹整个 `http.Handler`。捕获 `panic(error)` 并写入 JSON 错误响应。

### `NewErrorMiddleware(ib) func(http.Handler) http.Handler`

`ErrorMiddleware` 的便捷版本，使用默认配置。

### `WrapHandlerFunc(conf, handler) http.HandlerFunc`

包裹单个 `http.HandlerFunc`，行为与 `ErrorMiddleware` 一致（panic 捕获 + 翻译 + JSON 响应）。

### `WriteError(conf, w, r, err) error`

在 handler 内部显式调用，翻译错误并写入 JSON 响应。调用方可以显式处理写出失败。

### `HandleError(conf, w, r, err)`

兼容入口。内部调用 `WriteError` 并记录写出失败日志，不返回错误。

### `WriteJSONError(err, w) error`

最底层的写入函数，仅将 error 转为 JSON 写入 `ResponseWriter`。**不做翻译**。适合已经手动翻译过的场景。
