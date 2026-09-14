# Best Practices for httperrors Error Scenarios

This document lists recommended implementations for common error scenarios, including Go examples and the resulting JSON responses.

> **Convention**: always use currently supported `httperrors.Reason*` constants such as `ReasonNotFound` and `ReasonInvalidArgument` instead of hard-coded strings.  
> The `reason` inside `FieldViolation` is a business-level validation code such as `REQUIRED` or `TOO_SHORT`, and also serves as the i18n key.

---

## Table of Contents

1. Field validation error (`400 Bad Request`)
2. Invalid parameter (`400 Bad Request`)
3. Resource not found (`404 Not Found`)
4. Unauthenticated (`401 Unauthorized`)
5. Permission denied (`403 Forbidden`)
6. Resource conflict (`409 Conflict`)
7. Internal error (`500 Internal Server Error`)
8. Nested object validation
9. Array element validation
10. Mixed nested and array validation
11. Validation with i18n template arguments
12. Aggregating child-service validation errors
13. Business error with metadata
14. Request body parse failure
15. Rate limiting (`429 Too Many Requests`)
16. Timeout (`504 Gateway Timeout`)

---

## Error Propagation Model

```text
Service layer --return error--> Handler layer --panic(err)--> Middleware
```

- The **service layer** returns errors and does not panic
- The **handler layer** may use `panic(err)` to hand errors to middleware
- The **middleware** performs translation and writes JSON responses

---

## 1. Field Validation Error (`400 Bad Request`)

**Scenario**: a form submission fails validation on multiple fields.

```go
type CreateUserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Age      int    `json:"age"`
}

func (r *CreateUserRequest) Validate() error {
    var fvs []*httperrors.FieldViolation
    if r.Email == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("email", "REQUIRED", "email is required"))
    } else if !strings.Contains(r.Email, "@") {
        fvs = append(fvs, httperrors.NewFieldViolation("email", "INVALID_FORMAT", "invalid email format"))
    }
    if len(r.Password) < 8 {
        fvs = append(fvs, httperrors.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters"))
    }
    if r.Age < 0 || r.Age > 150 {
        fvs = append(fvs, httperrors.NewFieldViolation("age", "OUT_OF_RANGE", "age must be between 0 and 150"))
    }
    return httperrors.BadRequest(fvs).Err()
}
```

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    { "field": "email", "code": "REQUIRED", "message": "email is required" },
    { "field": "password", "code": "TOO_SHORT", "message": "password must be at least 8 characters" }
  ]
}
```

---

## 2. Invalid Parameter (`400 Bad Request`)

**Scenario**: a non-field parameter such as a URL or query parameter is invalid.

```go
func (s *ConsentService) UpdateConsent(ctx context.Context, consentType string) error {
    if consentType != "CONSENT_TYPE_GRANT" && consentType != "CONSENT_TYPE_REVOKE" {
        return httperrors.Error(
            http.StatusBadRequest,
            httperrors.ReasonInvalidArgument,
            "invalid consent type, must be CONSENT_TYPE_GRANT or CONSENT_TYPE_REVOKE",
        )
    }
    return nil
}
```

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid consent type, must be CONSENT_TYPE_GRANT or CONSENT_TYPE_REVOKE"
}
```

---

## 3. Resource Not Found (`404 Not Found`)

```go
func (s *UserService) GetByID(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to query user")
    }
    if user == nil {
        return nil, httperrors.Errorf(http.StatusNotFound, httperrors.ReasonNotFound, "user %s not found", id)
    }
    return user, nil
}
```

---

## 4. Unauthenticated (`401 Unauthorized`)

```go
func (s *AuthService) Authenticate(ctx context.Context, token string) (*Claims, error) {
    if token == "" {
        return nil, httperrors.Error(http.StatusUnauthorized, httperrors.ReasonUnauthenticated, "authentication required")
    }
    claims, err := s.verifyToken(ctx, token)
    if err != nil {
        return nil, httperrors.Error(http.StatusUnauthorized, httperrors.ReasonUnauthenticated, "token has expired")
    }
    return claims, nil
}
```

---

## 5. Permission Denied (`403 Forbidden`)

```go
func (s *ProjectService) Delete(ctx context.Context, projectID string) error {
    role := auth.RoleFromContext(ctx)
    if role != "admin" {
        return httperrors.New(http.StatusForbidden, httperrors.ReasonPermissionDenied, "admin role required to delete project").
            WithMetadata(map[string]string{
                "requiredRole": "admin",
                "currentRole":  role,
                "projectId":    projectID,
            }).Err()
    }
    return s.repo.Delete(ctx, projectID)
}
```

---

## 6. Resource Conflict (`409 Conflict`)

```go
func (s *UserService) Register(ctx context.Context, email, password string) (*User, error) {
    existing, _ := s.repo.FindByEmail(ctx, email)
    if existing != nil {
        return nil, httperrors.New(http.StatusConflict, httperrors.ReasonAlreadyExists, "email already registered").
            WithMetadata(map[string]string{"field": "email"}).Err()
    }
    return s.repo.Create(ctx, email, password)
}
```

---

## 7. Internal Error (`500 Internal Server Error`)

```go
func (s *OrderService) Create(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    order, err := s.repo.Insert(ctx, req)
    if err != nil {
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to create order")
    }
    if err := s.paymentClient.Charge(ctx, order.ID, req.Amount); err != nil {
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "payment service failed")
    }
    return order, nil
}
```

---

## 8. Nested Object Validation

Use `ToFieldViolations` to prepend nested field paths and return `httperrors.BadRequest(...)` after aggregation.

## 9. Array Element Validation

Use array indices such as `items[0]` and return `httperrors.BadRequest(...)` after collecting all field failures.

## 10. Mixed Nested and Array Validation

Compose nested paths and array indices through `ToFieldViolations`, then return `httperrors.BadRequest(...)`.

## 11. Validation with i18n Template Arguments

Use `WithLocalizedArgs(...)` when a localized message contains placeholders.

## 12. Aggregating Child-Service Validation Errors

Collect all child-service validation errors into a single `FieldViolations` slice and return `httperrors.BadRequest(...)` once at the end.

## 13. Business Error with Metadata

For business errors that need structured context, attach extra data through `WithMetadata(...)` instead of embedding it into `message`.

## 14. Request Body Parse Failure

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        panic(httperrors.New(http.StatusBadRequest, httperrors.ReasonInvalidArgument, "invalid JSON request body").
            WithCause(err).Err())
    }
}
```

## 15. Rate Limiting (`429 Too Many Requests`)

Use `httperrors.New(http.StatusTooManyRequests, httperrors.ReasonResourceExhausted, ...)`.

## 16. Timeout (`504 Gateway Timeout`)

Use `context.DeadlineExceeded` directly or wrap with `httperrors.WrapStatus(err, http.StatusGatewayTimeout, ...)`.

---

## Quick Reference

| Scenario | HTTP Status | Reason | Recommended API |
| --- | --- | --- | --- |
| Field validation | 400 | `ReasonInvalidArgument` | `BadRequest(fvs)` |
| Invalid parameter | 400 | `ReasonInvalidArgument` | `Error(400, ReasonInvalidArgument, msg)` |
| Not found | 404 | `ReasonNotFound` | `Errorf(404, ReasonNotFound, ...)` |
| Unauthenticated | 401 | `ReasonUnauthenticated` | `Error(401, ReasonUnauthenticated, msg)` |
| Permission denied | 403 | `ReasonPermissionDenied` | `New(403, ReasonPermissionDenied, msg).WithMetadata(...)` |
| Conflict | 409 | `ReasonAlreadyExists` | `New(409, ReasonAlreadyExists, msg).WithMetadata(...)` |
| Internal | 500 | `ReasonInternal` | `WrapStatus(err, 500, msg)` |
| Rate limit | 429 | `ReasonResourceExhausted` | `New(429, ReasonResourceExhausted, msg).WithMetadata(...)` |
| Timeout | 504 | `ReasonDeadlineExceeded` | automatic via `context.DeadlineExceeded` |
| JSON parse failure | 400 | `ReasonInvalidArgument` | `New(400, ReasonInvalidArgument, msg).WithCause(err)` |

---

## Core Principles

1. Use `Reason*` constants instead of hard-coded strings
2. Collect all field validation errors instead of returning on the first one
3. Return errors in services and use panic only at the HTTP boundary when middleware owns error handling
4. `reason` is both the error code and the default i18n key
5. `message` is for developers, `localizedMessage` is for user-facing text
6. Use `WrapStatus` to preserve the error chain
7. Put structured extra context into `metadata`, not into `message`
