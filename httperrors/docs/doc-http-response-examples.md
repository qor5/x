# httperrors HTTP Response Examples

本文档详细说明 `httperrors` 包在各种场景下的 HTTP 响应格式，明确区分 **Response Header** 和 **Response Body** 中的信息。

---

## 设计原则

- HTTP status code **只在 Response Header** 中，Body 中**不重复**
- Body 使用 **camelCase** 字段命名
- 空值字段通过 `omitempty` **自动省略**，不会出现在 Body 中
- `Content-Type` 固定为 `application/json`

---

## 错误传播模型

在正常的 Go 程序中，错误通过 `return error` 在调用链中逐层传递。
只有在 HTTP handler 的最终入口处，才通过 `panic` 将错误交给 `ErrorMiddleware` 捕获并写入 HTTP 响应。

```go
// ===== 业务/服务层：始终 return error =====

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")
    }
    return user, nil
}

// ===== HTTP handler 层：调用业务逻辑，panic 传播错误给 middleware =====

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.Context(), r.PathValue("id"))
    if err != nil {
        panic(err) // ErrorMiddleware 会捕获并写入 JSON 响应
    }
    // 正常响应...
    json.NewEncoder(w).Encode(user)
}
```

> **关键点**: `panic` 只出现在 handler 的边界处，是将 `error` 交给 middleware 的手段。
> 业务代码中一律使用 `return error`，与标准 Go 惯例一致。

以下各场景的 Go 代码示例分为 **业务层**（return）和 **handler 层**（panic）两部分。

---

## 1. 简单错误（无额外信息）

**场景**: 资源未找到

```go
// 业务层
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    // ...
    return nil, httperrors.Error(http.StatusNotFound, "NOT_FOUND", "user not found")
}

// handler 层
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

> `metadata`、`fieldViolations`、`localizedMessage` 均为空，不出现在 Body 中。

---

## 2. 简单错误 + i18n 翻译

**场景**: 同上，但请求方指定了中文

### Request Header

```
x-selected-language: zh
```

或者:

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
  "message": "未找到"
}
```

> `message` 被翻译为对应语言（翻译后替换原始 message）。`code` 始终是原始 reason 常量，不翻译。

---

## 3. 带 Metadata 的错误

**场景**: 权限不足，附带额外上下文

```go
// 业务层
func (s *ProjectService) Delete(ctx context.Context, id string) error {
    if !hasPermission(ctx, "project", "delete") {
        return httperrors.New(http.StatusForbidden, "PERMISSION_DENIED", "permission denied").
            WithMetadata(map[string]string{
                "resource": "project",
                "action":   "delete",
            }).Err()
    }
    // ...
    return nil
}

// handler 层
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

## 4. 字段校验错误（无翻译）

**场景**: 表单提交，多个字段不合法

```go
// 业务层 / 校验逻辑
func (req *CreateUserRequest) Validate() error {
    var fvs httperrors.FieldViolations
    if req.Email == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("email", "REQUIRED", "email is required"))
    }
    if len(req.Password) < 8 {
        fvs = append(fvs, httperrors.NewFieldViolation("password", "TOO_SHORT", "password must be at least 8 characters"))
    }
    return httperrors.ValidationError(fvs...).Err() // 无违规时返回 nil
}

// handler 层
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    json.NewDecoder(r.Body).Decode(&req)
    if err := req.Validate(); err != nil {
        panic(err)
    }
    // ...
}
```

### Response Header

```
HTTP/1.1 422 Unprocessable Entity
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

> 每个 `fieldViolation` 中的 `localizedMessage` 为空时不出现。

---

## 5. 字段校验错误 + i18n 翻译

**场景**: 同上，请求方指定中文

### Request Header

```
x-selected-language: zh
```

### Response Header

```
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
```

### Response Body

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "参数无效",
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

> - 顶层 `message` 被翻译（替换原始 message）
> - 每个 fieldViolation 的 `message` 保持原始英文描述不变
> - 翻译结果放在 `localizedMessage` 字段中

---

## 6. 嵌套字段校验错误

**场景**: 嵌套对象校验，字段路径用 `.` 分隔

```go
// 子对象校验
func (a *Address) Validate() error {
    var fvs httperrors.FieldViolations
    if a.Street == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("street", "REQUIRED", "street is required"))
    }
    if !isValidZip(a.ZipCode) {
        fvs = append(fvs, httperrors.NewFieldViolation("zipCode", "INVALID_FORMAT", "zip code format is invalid"))
    }
    return httperrors.ValidationError(fvs...).Err()
}

// 父对象校验：用 ToFieldViolations 组合子对象的错误并 prepend 字段前缀
func (req *CreateOrderRequest) Validate() error {
    var all httperrors.FieldViolations
    if err := req.Address.Validate(); err != nil {
        all = append(all, httperrors.ToFieldViolations(err, "address")...)
    }
    return httperrors.ValidationError(all...).Err()
}
```

### Response Header

```
HTTP/1.1 422 Unprocessable Entity
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

## 7. 数组元素校验错误

**场景**: 数组中某个元素校验失败

```go
// 业务层
func (req *CreateOrderRequest) Validate() error {
    var fvs httperrors.FieldViolations
    for i, item := range req.Items {
        if item.Quantity <= 0 {
            fvs = append(fvs, httperrors.NewFieldViolationf(
                fmt.Sprintf("items[%d].quantity", i), "OUT_OF_RANGE",
                "quantity must be positive",
            ))
        }
    }
    return httperrors.ValidationError(fvs...).Err()
}
```

### Response Header

```
HTTP/1.1 422 Unprocessable Entity
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

## 8. 服务端内部错误

**场景**: 未预期的异常

```go
// 业务层：包装底层错误
func (s *OrderService) Create(ctx context.Context, req *CreateOrderRequest) error {
    if err := s.repo.Insert(ctx, req); err != nil {
        return httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to create order")
    }
    return nil
}

// handler 层
func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    // ...
    if err := h.orderService.Create(r.Context(), &req); err != nil {
        panic(err)
    }
    // ...
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
  "message": "internal server error"
}
```

> 内部错误的 cause（堆栈、原始错误信息）**不会暴露给前端**，只在服务端日志中可见。

---

## 9. 认证/鉴权错误

### 9a. 未认证（401）

```go
// 中间件 / 业务层
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !isAuthenticated(r) {
            panic(httperrors.Error(http.StatusUnauthorized, "UNAUTHENTICATED", "authentication required"))
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

### 9b. 无权限（403）

```go
// 业务层
func (s *ResourceService) Get(ctx context.Context, id string) (*Resource, error) {
    if !hasAccess(ctx, id) {
        return nil, httperrors.Error(http.StatusForbidden, "PERMISSION_DENIED", "you do not have access to this resource")
    }
    // ...
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

## 10. 冲突错误

**场景**: 乐观锁冲突、唯一约束冲突

```go
// 业务层
func (s *UserService) Register(ctx context.Context, email string) error {
    exists, _ := s.repo.ExistsByEmail(ctx, email)
    if exists {
        return httperrors.New(http.StatusConflict, "ALREADY_EXISTS", "email already registered").
            WithMetadata(map[string]string{"field": "email"}).Err()
    }
    // ...
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

## 11. 限流错误

```go
// 中间件
func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if isRateLimited(r) {
            panic(httperrors.New(http.StatusTooManyRequests, "RESOURCE_EXHAUSTED", "rate limit exceeded").
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

## 12. 超时错误

**场景**: 上游服务超时

```go
// 业务层：context.DeadlineExceeded 会被 FromError 自动识别为 504
func (s *PaymentService) Charge(ctx context.Context, req *ChargeRequest) error {
    resp, err := s.gateway.Charge(ctx, req)
    if err != nil {
        // 如果 err 是 context.DeadlineExceeded，FromError 会自动转为 504 + DEADLINE_EXCEEDED
        // 也可以显式包装：
        return httperrors.WrapStatus(err, http.StatusGatewayTimeout, "upstream service timed out")
    }
    // ...
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

## 字段总览

### Response Header 中的信息

| Header           | 说明                   | 示例                                |
| ---------------- | ---------------------- | ----------------------------------- |
| HTTP Status Line | HTTP 状态码 + 标准短语 | `HTTP/1.1 422 Unprocessable Entity` |
| Content-Type     | 固定值                 | `application/json`                  |

### Response Body 中的信息

| 字段               | 类型     | 是否必有 | 说明                                               |
| ------------------ | -------- | -------- | -------------------------------------------------- |
| `code`             | `string` | **是**   | 错误原因常量（如 `NOT_FOUND`、`REQUIRED`），不翻译 |
| `message`          | `string` | **是**   | 人类可读消息，会被 i18n 翻译替换                   |
| `localizedMessage` | `string` | 否       | 顶层翻译消息（仅在翻译后出现）                     |
| `metadata`         | `object` | 否       | 键值对附加信息，空时省略                           |
| `fieldViolations`  | `array`  | 否       | 字段级校验错误列表，空时省略                       |

### fieldViolations 数组元素

| 字段               | 类型     | 是否必有 | 说明                                     |
| ------------------ | -------- | -------- | ---------------------------------------- |
| `field`            | `string` | **是**   | 字段路径，支持 `.` 嵌套和 `[n]` 数组索引 |
| `code`             | `string` | **是**   | 该字段的错误原因常量                     |
| `message`          | `string` | **是**   | 原始描述，始终保持不翻译                 |
| `localizedMessage` | `string` | 否       | 该字段的翻译消息，空时省略               |

---

## 前端处理建议

```typescript
interface ErrorResponse {
  code: string;
  message: string;
  localizedMessage?: string;
  metadata?: Record<string, string>;
  fieldViolations?: FieldViolation[];
}

interface FieldViolation {
  field: string;
  code: string;
  message: string;
  localizedMessage?: string;
}

// 使用示例
async function handleResponse(resp: Response) {
  if (!resp.ok) {
    const err: ErrorResponse = await resp.json();

    // 展示给用户的消息优先使用 localizedMessage（如果有），否则 fallback 到 message
    const displayMessage = err.localizedMessage || err.message;

    // 处理字段级错误
    if (err.fieldViolations) {
      for (const fv of err.fieldViolations) {
        const fieldMsg = fv.localizedMessage || fv.message;
        setFieldError(fv.field, fieldMsg);
      }
    }

    // 根据 code 做特定逻辑
    if (err.code === "UNAUTHENTICATED") {
      redirectToLogin();
    }
  }
}
```
