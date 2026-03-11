# Migrating from statusx to httperrors

## Q: Does every service need to implement its own middleware?

**No.** The `httperrors` package already provides reusable `ErrorMiddleware` and `NewErrorMiddleware` helpers out of the box.
A service only needs to register the middleware once during startup. It does not need to implement panic recovery, error translation, or JSON serialization on its own.

```go
mux.Use(httperrors.NewErrorMiddleware(ib))
```

If you need custom error-writing logic such as logging or metrics, you can extend it through hooks:

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

## Old statusx vs New httperrors: Service Integration Comparison

### Old Approach: statusx (Connect RPC / gRPC)

Legacy `statusx` services typically run over the Connect RPC protocol and involve multiple components in the integration path.

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
    ib, _ := i18nx.New(strings.NewReader(messagesCSV))

    mux := connectx.NewMux(ib)
    mux.Handle(
        userv1connect.NewUserServiceHandler,
        orderv1connect.NewOrderServiceHandler,
    )

    http.ListenAndServe(":8080", mux)
}
```

**Business layer:**

```go
func (s *UserServiceServer) GetUser(
    ctx context.Context,
    req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
    user, err := s.repo.FindByID(ctx, req.Msg.Id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, statusx.New(codes.NotFound, "NOT_FOUND", "user not found").Err()
    }
    return connect.NewResponse(&userv1.GetUserResponse{User: toProto(user)}), nil
}
```

### New Approach: httperrors (Pure HTTP)

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
    handler := httperrors.NewErrorMiddleware(ib)(mux)

    mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
    mux.HandleFunc("POST /orders", orderHandler.Create)

    http.ListenAndServe(":8080", handler)
}
```

**Business layer:**

```go
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to query user")
    }
    if user == nil {
        return nil, httperrors.Error(http.StatusNotFound, httperrors.ReasonNotFound, "user not found")
    }
    return user, nil
}
```

---

## Response Format Comparison

`statusx` has two response paths: Connect protocol and VProto protocol.
`httperrors` only has a single plain JSON path.

### Scenario 1: Simple Error Without i18n

```go
// statusx
statusx.New(codes.NotFound, "NOT_FOUND", "user not found")

// httperrors
httperrors.Error(http.StatusNotFound, httperrors.ReasonNotFound, "user not found")
```

**httperrors**:

```json
{
  "code": "NOT_FOUND",
  "message": "user not found"
}
```

### Scenario 2: Simple Error with i18n

```json
{
  "code": "NOT_FOUND",
  "message": "user not found",
  "localizedMessage": "未找到"
}
```

### Scenario 3: Field Validation Error Without i18n

```go
fv1 := httperrors.NewFieldViolation("email", "REQUIRED", "email is required")
fv2 := httperrors.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters")
httperrors.BadRequest(fv1, fv2)
```

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

### Scenario 4: Field Validation Error with i18n

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

---

## Summary of Changes in Business Code

```diff
 // Error creation
-statusx.New(codes.NotFound, "NOT_FOUND", "user not found")
+httperrors.Error(http.StatusNotFound, httperrors.ReasonNotFound, "user not found")

 // Error wrapping
-statusx.Wrap(err, codes.Internal, "INTERNAL", "database error")
+httperrors.WrapStatus(err, http.StatusInternalServerError, "database error")

 // Field validation
-statusx.BadRequest(fv1, fv2)
+httperrors.BadRequest(fv1, fv2)

 // Field violation creation
-statusx.NewFieldViolation("email", "REQUIRED", "email is required")
+httperrors.NewFieldViolation("email", "REQUIRED", "email is required")

 // i18n
-s.WithLocalized("key", args...)
+s.WithLocalized("key", args...)

 // Error parsing
-statusx.FromError(err)
+httperrors.FromError(err)
```

**Core change**: replace `codes.Code` with `http.StatusXxx`. Most other API shapes remain similar.

---

## Practical Service Migration Checklist

### 1. DI Layer

| Change | Meaning |
| --- | --- |
| Remove `SetupProttpHandler` | No longer need prottpx |
| Remove `SetupGRPCConn` | No longer need gRPC connections unless something else still uses them |
| Update `SetupHTTPServer` parameters | Remove gRPC/prottpx dependencies and add `*i18nx.I18N` |

### 2. Middleware Chain

| Change | Meaning |
| --- | --- |
| Add `httperrors.NewErrorMiddleware(ib)` | Standard panic recovery + i18n + JSON output |
| Remove `statusx.AllowHeaders` | No longer need statusx-specific headers |
| Adjust health checks | Replace gRPC health checks with HTTP health checks if appropriate |

### 3. Routing

| Change | Meaning |
| --- | --- |
| Replace protobuf service registration | Register plain HTTP routes with `mux.HandleFunc` |
| Map RPC methods to HTTP endpoints | Convert service methods into HTTP handlers |

### 4. Service Layer

| Change | Meaning |
| --- | --- |
| `statusx.New(codes.X, ...)` → `httperrors.Error(http.StatusX, ...)` | Replace status code type |
| `statusx.Wrap(...)` → `httperrors.WrapStatus(...)` | Error wrapping |
| `statusx.BadRequest(...)` → `httperrors.BadRequest(...)` | Field validation |
| `statusx.FromError(...)` → `httperrors.FromError(...)` | Error parsing |

### 5. Handler Layer

| Change | Meaning |
| --- | --- |
| gRPC service interface → `http.HandlerFunc` | Handler signature changes |
| `return nil, err` → `panic(err)` | Use panic at the handler boundary so middleware can handle the error |
| protobuf request/response → JSON request/response | Serialization format changes |
