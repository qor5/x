# httperrors HTTP Response Examples

This document explains the HTTP response format produced by the `httperrors` package in different scenarios and clearly separates what belongs in the **Response Header** from what belongs in the **Response Body**.

---

## Design Principles

- The HTTP status code appears **only in the Response Header** and is **not duplicated** in the body
- The body uses **camelCase** field names
- Empty fields are automatically omitted through `omitempty`
- `Content-Type` is always `application/json`

---

## Error Propagation Model

In normal Go code, errors are propagated layer by layer using `return error`.
Only at the final HTTP handler boundary is an error handed to `ErrorMiddleware` via `panic` so it can be converted into an HTTP response.

```go
// ===== Service layer: always return error =====

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

// ===== HTTP handler layer: panic to hand the error to middleware =====

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

> **Key point**: `panic` only appears at the handler boundary. It is the mechanism used to hand an `error` to middleware.
> Business logic should still use `return error`, following standard Go conventions.

The code examples below are split into two parts where needed: the **service layer** (`return`) and the **handler layer** (`panic`).

---

## 1. Simple Error Without Extra Information

**Scenario**: resource not found.

```go
// Service layer
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    return nil, httperrors.Error(http.StatusNotFound, httperrors.ReasonNotFound, "user not found")
}

// Handler layer
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

### Response Header

```
HTTP/1.1 404 Not Found
Content-Type: application/json
```

### Response Body

```json
{
  "code": "NOT_FOUND",
  "message": "user not found"
}
```

> `metadata`, `fieldViolations`, and `localizedMessage` are omitted when empty.

---

## 2. Simple Error with i18n Translation

**Scenario**: same as above, but the caller requests Chinese.

### Request Header

```
x-selected-language: zh
```

or:

```
Accept-Language: zh
```

### Response Header

```
HTTP/1.1 404 Not Found
Content-Type: application/json
```

### Response Body

```json
{
  "code": "NOT_FOUND",
  "message": "user not found",
  "localizedMessage": "未找到"
}
```

> `message` keeps the original English text. The translation is placed in `localizedMessage`. `code` is never translated.

---

## 3. Error with Metadata

**Scenario**: permission denied with extra structured context.

```go
// Service layer
func (s *ProjectService) Delete(ctx context.Context, id string) error {
    if !hasPermission(ctx, "project", "delete") {
        return httperrors.New(http.StatusForbidden, httperrors.ReasonPermissionDenied, "permission denied").
            WithMetadata(map[string]string{
                "resource": "project",
                "action":   "delete",
            }).Err()
    }
    return nil
}

// Handler layer
func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
    if err := h.projectService.Delete(r.Context(), r.PathValue("id")); err != nil {
        panic(err)
    }
    w.WriteHeader(http.StatusNoContent)
}
```

### Response Header

```
HTTP/1.1 403 Forbidden
Content-Type: application/json
```

### Response Body

```json
{
  "code": "PERMISSION_DENIED",
  "message": "permission denied",
  "metadata": {
    "resource": "project",
    "action": "delete"
  }
}
```

---

## 4. Field Validation Error Without Translation

**Scenario**: a form submission contains multiple invalid fields.

```go
func (req *CreateUserRequest) Validate() error {
    var fvs httperrors.FieldViolations
    if req.Email == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("email", "REQUIRED", "email is required"))
    }
    if len(req.Password) < 8 {
        fvs = append(fvs, httperrors.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters"))
    }
    return httperrors.BadRequest(fvs...).Err()
}
```

### Response Header

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
```

### Response Body

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

---

## 5. Field Validation Error with i18n Translation

**Scenario**: same as above, but the caller requests Chinese.

### Request Header

```
x-selected-language: zh
```

### Response Header

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
```

### Response Body

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
      "localizedMessage": "邮箱不能为空"
    },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "message": "password must be at least 8 characters",
      "localizedMessage": "密码至少需要8个字符"
    }
  ]
}
```

> The top-level `message` remains unchanged and the translated text goes to `localizedMessage`.
> The same rule applies to each field violation.

---

## 6. Nested Field Validation Error

**Scenario**: nested object validation using dot-separated field paths.

```go
func (a *Address) Validate() error {
    var fvs httperrors.FieldViolations
    if a.Street == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("street", "REQUIRED", "street is required"))
    }
    if !isValidZip(a.ZipCode) {
        fvs = append(fvs, httperrors.NewFieldViolation("zipCode", "INVALID_FORMAT", "zip code format is invalid"))
    }
    return httperrors.BadRequest(fvs...).Err()
}

func (req *CreateOrderRequest) Validate() error {
    var all httperrors.FieldViolations
    if err := req.Address.Validate(); err != nil {
        all = append(all, httperrors.ToFieldViolations(err, "address")...)
    }
    return httperrors.BadRequest(all...).Err()
}
```

### Response Header

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
```

### Response Body

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "address.street",
      "code": "REQUIRED",
      "message": "street is required"
    },
    {
      "field": "address.zipCode",
      "code": "INVALID_FORMAT",
      "message": "zip code format is invalid"
    }
  ]
}
```

---

## 7. Array Element Validation Error

**Scenario**: one or more elements inside an array are invalid.

```go
func (req *CreateOrderRequest) Validate() error {
    var fvs httperrors.FieldViolations
    for i, item := range req.Items {
        if item.Quantity <= 0 {
            fvs = append(fvs, httperrors.NewFieldViolation(
                fmt.Sprintf("items[%d].quantity", i), "OUT_OF_RANGE",
                "quantity must be positive",
            ))
        }
    }
    return httperrors.BadRequest(fvs...).Err()
}
```

### Response Header

```
HTTP/1.1 400 Bad Request
Content-Type: application/json
```

### Response Body

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "items[2].quantity",
      "code": "OUT_OF_RANGE",
      "message": "quantity must be positive"
    }
  ]
}
```

---

## 8. Internal Server Error

**Scenario**: unexpected backend failure.

```go
func (s *OrderService) Create(ctx context.Context, req *CreateOrderRequest) error {
    if err := s.repo.Insert(ctx, req); err != nil {
        return httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to create order")
    }
    return nil
}
```

### Response Header

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
```

### Response Body

```json
{
  "code": "INTERNAL",
  "message": "failed to create order"
}
```

---

## 9. Authentication and Authorization Errors

### 9a. Unauthenticated (401)

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !isAuthenticated(r) {
            panic(httperrors.Error(http.StatusUnauthorized, httperrors.ReasonUnauthenticated, "authentication required"))
        }
        next.ServeHTTP(w, r)
    })
}
```

#### Response Header

```
HTTP/1.1 401 Unauthorized
Content-Type: application/json
```

#### Response Body

```json
{
  "code": "UNAUTHENTICATED",
  "message": "authentication required"
}
```

### 9b. Permission Denied (403)

```go
func (s *ResourceService) Get(ctx context.Context, id string) (*Resource, error) {
    if !hasAccess(ctx, id) {
        return nil, httperrors.Error(http.StatusForbidden, httperrors.ReasonPermissionDenied, "you do not have access to this resource")
    }
    return resource, nil
}
```

#### Response Header

```
HTTP/1.1 403 Forbidden
Content-Type: application/json
```

#### Response Body

```json
{
  "code": "PERMISSION_DENIED",
  "message": "you do not have access to this resource"
}
```

---

## 10. Conflict Error

**Scenario**: optimistic locking or uniqueness conflict.

```go
func (s *UserService) Register(ctx context.Context, email string) error {
    exists, _ := s.repo.ExistsByEmail(ctx, email)
    if exists {
        return httperrors.New(http.StatusConflict, httperrors.ReasonAlreadyExists, "email already registered").
            WithMetadata(map[string]string{"field": "email"}).Err()
    }
    return nil
}
```

### Response Header

```
HTTP/1.1 409 Conflict
Content-Type: application/json
```

### Response Body

```json
{
  "code": "ALREADY_EXISTS",
  "message": "email already registered",
  "metadata": {
    "field": "email"
  }
}
```

---

## 11. Rate Limit Error

```go
func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if isRateLimited(r) {
            panic(httperrors.New(http.StatusTooManyRequests, httperrors.ReasonResourceExhausted, "rate limit exceeded").
                WithMetadata(map[string]string{"retryAfter": "30"}).Err())
        }
        next.ServeHTTP(w, r)
    })
}
```

### Response Header

```
HTTP/1.1 429 Too Many Requests
Content-Type: application/json
```

### Response Body

```json
{
  "code": "RESOURCE_EXHAUSTED",
  "message": "rate limit exceeded",
  "metadata": {
    "retryAfter": "30"
  }
}
```

---

## 12. Timeout Error

**Scenario**: upstream service timeout.

```go
func (s *PaymentService) Charge(ctx context.Context, req *ChargeRequest) error {
    _, err := s.gateway.Charge(ctx, req)
    if err != nil {
        return httperrors.WrapStatus(err, http.StatusGatewayTimeout, "upstream service timed out")
    }
    return nil
}
```

### Response Header

```
HTTP/1.1 504 Gateway Timeout
Content-Type: application/json
```

### Response Body

```json
{
  "code": "DEADLINE_EXCEEDED",
  "message": "upstream service timed out"
}
```

---

## Field Overview

### Information in the Response Header

| Header | Meaning | Example |
| --- | --- | --- |
| HTTP Status Line | HTTP status code and standard phrase | `HTTP/1.1 400 Bad Request` |
| Content-Type | Fixed value | `application/json` |

### Information in the Response Body

| Field | Type | Required | Meaning |
| --- | --- | --- | --- |
| `code` | `string` | Yes | Error reason constant such as `NOT_FOUND` or `INVALID_ARGUMENT` |
| `message` | `string` | Yes | Original human-readable message |
| `localizedMessage` | `string` | No | Translated message |
| `metadata` | `object` | No | Additional key-value data |
| `fieldViolations` | `array` | No | Field-level validation failures |

### Elements inside `fieldViolations`

| Field | Type | Required | Meaning |
| --- | --- | --- | --- |
| `field` | `string` | Yes | Field path, including nested `.` and array indices like `[n]` |
| `code` | `string` | Yes | Field-level validation code |
| `message` | `string` | Yes | Original field-level message |
| `localizedMessage` | `string` | No | Translated field-level message |
