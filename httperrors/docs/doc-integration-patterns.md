# httperrors Integration Patterns Guide

This document explains how to integrate the `httperrors` package in different scenarios and outlines the recommended practices for each pattern.

---

## Three Integration Patterns

| Pattern                    | Functions                                | Suitable Scenario                                   | Error Propagation              |
| -------------------------- | ---------------------------------------- | --------------------------------------------------- | ------------------------------ |
| **Global middleware**      | `ErrorMiddleware` / `NewErrorMiddleware` | All handlers use httperrors                         | panic                          |
| **Wrapped single handler** | `WrapHandlerFunc`                        | Only part of a mux uses httperrors                  | panic                          |
| **Explicit handling**      | `WriteError` / `HandleError`             | The handler handles errors explicitly without panic | return error / direct handling |

---

## 1. Global Middleware Pattern

**Suitable scenario**: every handler in the service uses `httperrors` for error responses.

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

    handler := httperrors.NewErrorMiddleware(ib)(mux)
    http.ListenAndServe(":8080", handler)
}
```

**Service layer** (always returns `error`):

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

**Handler layer** (panics to hand the error to middleware):

```go
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

**Pros**: concise and requires the least handler code  
**Cons**: every handler must follow the `panic(error)` convention

---

## 2. Wrapped Single Handler Pattern

**Suitable scenario**: a mux mixes multiple styles and only some handlers use `httperrors`.

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

The wrapped handlers still use panic internally:

```go
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

**Pros**: flexible and can be adopted handler by handler  
**Cons**: each registration needs an extra wrapper

---

## 3. Explicit Handling Pattern

**Suitable scenario**:

- the handler prefers `return` over `panic`
- extra logic is needed before or after writing the error response, such as logging or metrics
- the integration environment cannot use panic-based handling

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

Use `WriteError` explicitly inside the handler:

```go
type UserHandler struct {
    conf        *httperrors.HTTPErrorMiddlewareConfig
    userService *UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        if werr := httperrors.WriteError(h.conf, w, r, err); werr != nil {
            slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err, "writeError", werr)
        }
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

If only backward compatibility is needed, `HandleError` can still be used:

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

**Pros**: does not rely on panic and is fully explicit, which is closer to standard Go style  
**Cons**: every error branch must explicitly write the response and return

---

## 4. Mixed Pattern (Recommended for Gradual Migration)

**Suitable scenario**: a large project is migrating gradually from legacy error handling to `httperrors`.

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

## 5. Integration with Custom Hooks

All three patterns support custom hooks through `HTTPErrorMiddlewareConfig`, which can be used for logging, metrics, or response customization before and after error writing.

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
if werr := httperrors.WriteError(conf, w, r, err); werr != nil {
    slog.ErrorContext(r.Context(), "Failed to write http response error", "error", err, "writeError", werr)
}
httperrors.HandleError(conf, w, r, err)
```

---

## Pattern Selection Decision Tree

```
Do all handlers on your mux use httperrors?
├── Yes → Global middleware pattern (ErrorMiddleware)
└── No → Only part of the handlers use httperrors
         ├── Can those handlers use panic-based propagation?
         │   ├── Yes → Wrapped single handler pattern (WrapHandlerFunc)
         │   └── No → Explicit handling pattern (WriteError / HandleError)
         └── Are you doing a gradual migration?
             └── Yes → Mixed pattern (WrapHandlerFunc + WriteError/HandleError + legacy code)
```

---

## API Reference

### `ErrorMiddleware(conf) func(http.Handler) http.Handler`

Global middleware that wraps an entire `http.Handler`. It catches `panic(error)` and writes a JSON error response.

### `NewErrorMiddleware(ib) func(http.Handler) http.Handler`

Convenience version of `ErrorMiddleware` using the default configuration.

### `WrapHandlerFunc(conf, handler) http.HandlerFunc`

Wraps a single `http.HandlerFunc`. Its behavior matches `ErrorMiddleware` (panic capture + translation + JSON response).

### `WriteError(conf, w, r, err) error`

Explicitly used inside a handler to translate an error and write a JSON response. The caller can explicitly handle write failures.

### `HandleError(conf, w, r, err)`

Backward-compatible wrapper. It internally calls `WriteError` and logs write failures without returning them.

### `WriteJSONError(err, w) error`

The lowest-level writer. It only converts an error into JSON and writes it to the `ResponseWriter`. It **does not translate** the error.
