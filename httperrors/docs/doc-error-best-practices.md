# httperrors 错误场景最佳实践

本文档列举各种常见错误场景的最佳实现方式，包括 Go 代码和对应的 JSON 返回。

> **约定**：`reason` 一律使用 `httperrors.Reason*` 常量（如 `ReasonNotFound`、`ReasonBadRequest`），不硬编码字符串。  
> FieldViolation 的 `reason` 是业务级校验码（如 `REQUIRED`、`TOO_SHORT`），同时作为 i18n key。

---

## 目录

1. [字段校验错误（422 Unprocessable Entity）](#1-字段校验错误422-unprocessable-entity)
2. [参数无效（400 Bad Request）](#2-参数无效400-bad-request)
3. [资源不存在（404 Not Found）](#3-资源不存在404-not-found)
4. [未认证（401 Unauthorized）](#4-未认证401-unauthorized)
5. [权限不足（403 Forbidden）](#5-权限不足403-forbidden)
6. [资源冲突（409 Conflict）](#6-资源冲突409-conflict)
7. [内部错误（500 Internal Server Error）](#7-内部错误500-internal-server-error)
8. [嵌套对象校验](#8-嵌套对象校验)
9. [数组元素校验](#9-数组元素校验)
10. [混合嵌套 + 数组校验](#10-混合嵌套--数组校验)
11. [带 i18n 模板参数的校验](#11-带-i18n-模板参数的校验)
12. [子服务错误聚合](#12-子服务错误聚合)
13. [带 Metadata 的业务错误](#13-带-metadata-的业务错误)
14. [请求体解析失败](#14-请求体解析失败)
15. [限流（429 Too Many Requests）](#15-限流429-too-many-requests)
16. [超时（504 Gateway Timeout）](#16-超时504-gateway-timeout)

---

## 错误传播模型

```
┌─────────────┐     return error      ┌──────────────┐     panic(err)     ┌────────────────┐
│  Service 层  │ ──────────────────→  │  Handler 层   │ ────────────────→ │   Middleware    │
│  (业务逻辑)  │                      │  (HTTP 边界)  │                   │ (翻译+写JSON)  │
└─────────────┘                       └──────────────┘                    └────────────────┘
```

- **Service 层**：用 `return error` 传播，不 panic
- **Handler 层**：用 `panic(err)` 将错误交给 middleware 处理
- **Middleware**：自动翻译（i18n）+ 写 JSON 响应

---

## 1. 字段校验错误（422 Unprocessable Entity）

**场景**：表单提交，多个字段校验失败

### Service 层

```go
type CreateUserRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Age      int    `json:"age"`
}

// 推荐：在请求对象上实现 Validator 接口
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
    // 无错误时返回 nil（ValidationError 传入空切片时 Err() 返回 nil）
    return httperrors.ValidationError(fvs).Err()
}

func (s *UserService) Create(ctx context.Context, req *CreateUserRequest) (*User, error) {
    if err := httperrors.Validate(ctx, req); err != nil {
        return nil, err
    }
    return s.repo.Create(ctx, req)
}
```

### Handler 层

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        panic(httperrors.Error(http.StatusBadRequest, httperrors.ReasonBadRequest, "invalid request body"))
    }
    user, err := h.service.Create(r.Context(), &req)
    if err != nil {
        panic(err)
    }
    json.NewEncoder(w).Encode(user)
}
```

### JSON 返回

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    { "field": "email", "code": "REQUIRED", "message": "email is required" },
    {
      "field": "password",
      "code": "TOO_SHORT",
      "message": "password must be at least 8 characters"
    }
  ]
}
```

### JSON 返回（中文 i18n）

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
      "localizedMessage": "密码太短"
    }
  ]
}
```

> **要点**：一次性收集所有字段错误，不要遇到第一个就 return。`reason` 默认作为 i18n key。

---

## 2. 参数无效（400 Bad Request）

**场景**：URL 参数、查询参数等非字段级别的参数错误

### Service 层

```go
func (s *ConsentService) UpdateConsent(ctx context.Context, consentType string) error {
    if consentType != "CONSENT_TYPE_GRANT" && consentType != "CONSENT_TYPE_REVOKE" {
        return httperrors.Error(
            http.StatusBadRequest,
            httperrors.ReasonBadRequest,
            "invalid consent type, must be CONSENT_TYPE_GRANT or CONSENT_TYPE_REVOKE",
        )
    }
    // ... 业务逻辑
    return nil
}
```

### JSON 返回

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid consent type, must be CONSENT_TYPE_GRANT or CONSENT_TYPE_REVOKE"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid consent type, must be CONSENT_TYPE_GRANT or CONSENT_TYPE_REVOKE",
  "localizedMessage": "无效的请求参数"
}
```

### 使用 `NewStatus` 简写

如果不需要指定 reason，可以用 `NewStatus` 根据 HTTP 状态码自动推导：

```go
return httperrors.NewStatus(http.StatusBadRequest, "invalid page number").Err()
```

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid page number"
}
```

> **要点**：`reason` 使用 `httperrors.Reason*` 常量（如 `ReasonBadRequest`），不要硬编码字符串。`reason` 同时作为 i18n key，翻译自动填入 `localizedMessage`。

---

## 3. 资源不存在（404 Not Found）

### Service 层

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

### JSON 返回

```json
{
  "code": "NOT_FOUND",
  "message": "user abc123 not found"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "NOT_FOUND",
  "message": "user abc123 not found",
  "localizedMessage": "未找到"
}
```

> **要点**：`message` 中可以包含动态信息（如 ID）帮助调试，`reason` 使用 `Reason*` 常量。

---

## 4. 未认证（401 Unauthorized）

**场景**：Token 缺失或过期

### 中间件 / Service 层

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

### JSON 返回

```json
{
  "code": "UNAUTHENTICATED",
  "message": "token has expired"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "UNAUTHENTICATED",
  "message": "token has expired",
  "localizedMessage": "未认证"
}
```

> **要点**：使用 `ReasonUnauthenticated` 常量。前端通过 `message` 区分"未登录"和"Token 过期"，或通过不同的 `localizedMessage` 翻译展示给用户。

---

## 5. 权限不足（403 Forbidden）

**场景**：用户已认证但无权操作，附加 metadata 提供上下文

### Service 层

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

### JSON 返回

```json
{
  "code": "PERMISSION_DENIED",
  "message": "admin role required to delete project",
  "metadata": {
    "requiredRole": "admin",
    "currentRole": "editor",
    "projectId": "proj-001"
  }
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "PERMISSION_DENIED",
  "message": "admin role required to delete project",
  "localizedMessage": "权限不足",
  "metadata": {
    "requiredRole": "admin",
    "currentRole": "editor",
    "projectId": "proj-001"
  }
}
```

> **要点**：`metadata` 提供结构化上下文，便于前端展示详细信息或日志追踪。

---

## 6. 资源冲突（409 Conflict）

**场景**：唯一约束冲突（如邮箱已注册）

### Service 层

```go
func (s *UserService) Register(ctx context.Context, email, password string) (*User, error) {
    existing, _ := s.repo.FindByEmail(ctx, email)
    if existing != nil {
        return nil, httperrors.New(http.StatusConflict, httperrors.ReasonAlreadyExists, "email already registered").
            WithMetadata(map[string]string{"field": "email"}).Err()
    }
    // ...
    return s.repo.Create(ctx, email, password)
}
```

### JSON 返回

```json
{
  "code": "ALREADY_EXISTS",
  "message": "email already registered",
  "metadata": {
    "field": "email"
  }
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "ALREADY_EXISTS",
  "message": "email already registered",
  "localizedMessage": "已存在",
  "metadata": {
    "field": "email"
  }
}
```

---

## 7. 内部错误（500 Internal Server Error）

**场景**：数据库、外部服务等不可预期的错误

### Service 层

```go
func (s *OrderService) Create(ctx context.Context, req *CreateOrderRequest) (*Order, error) {
    order, err := s.repo.Insert(ctx, req)
    if err != nil {
        // WrapStatus 保留原始错误的堆栈信息，方便日志追踪
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "failed to create order")
    }

    if err := s.paymentClient.Charge(ctx, order.ID, req.Amount); err != nil {
        // 外部服务调用失败
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "payment service failed")
    }

    return order, nil
}
```

### JSON 返回

```json
{
  "code": "INTERNAL",
  "message": "failed to create order"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INTERNAL",
  "message": "failed to create order",
  "localizedMessage": "内部错误"
}
```

> **要点**：
>
> - 用 `WrapStatus` 而非 `Error`，保留原始错误链（堆栈、cause），方便后端日志定位
> - `message` 不要暴露内部细节（如 SQL 语句、连接字符串）
> - 原始错误可以通过 Hook 记录到日志/监控

---

## 8. 嵌套对象校验

**场景**：请求体含嵌套对象，字段路径用 `.` 分隔

### Service 层

```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

func (a *Address) Validate() error {
    var fvs []*httperrors.FieldViolation
    if a.Street == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("street", "REQUIRED", "street is required"))
    }
    if a.City == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("city", "REQUIRED", "city is required"))
    }
    return httperrors.ValidationError(fvs).Err()
}

type CreateOrderRequest struct {
    ProductID string   `json:"productId"`
    Address   *Address `json:"address"`
}

func (r *CreateOrderRequest) Validate() error {
    var fvs httperrors.FieldViolations
    if r.ProductID == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("productId", "REQUIRED", "product ID is required"))
    }
    if r.Address != nil {
        // ToFieldViolations 自动给子对象的字段加上 "address." 前缀
        if err := r.Address.Validate(); err != nil {
            fvs = append(fvs, httperrors.ToFieldViolations(err, "address")...)
        }
    } else {
        fvs = append(fvs, httperrors.NewFieldViolation("address", "REQUIRED", "address is required"))
    }
    return httperrors.ValidationError(fvs).Err()
}
```

### JSON 返回

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "productId",
      "code": "REQUIRED",
      "message": "product ID is required"
    },
    {
      "field": "address.street",
      "code": "REQUIRED",
      "message": "street is required"
    },
    {
      "field": "address.city",
      "code": "REQUIRED",
      "message": "city is required"
    }
  ]
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "productId",
      "code": "REQUIRED",
      "message": "product ID is required",
      "localizedMessage": "必填"
    },
    {
      "field": "address.street",
      "code": "REQUIRED",
      "message": "street is required",
      "localizedMessage": "必填"
    },
    {
      "field": "address.city",
      "code": "REQUIRED",
      "message": "city is required",
      "localizedMessage": "必填"
    }
  ]
}
```

---

## 9. 数组元素校验

**场景**：批量操作，数组中某些元素校验失败

### Service 层

```go
type Item struct {
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func (item *Item) Validate() error {
    var fvs []*httperrors.FieldViolation
    if item.Name == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("name", "REQUIRED", "name is required"))
    }
    if item.Price <= 0 {
        fvs = append(fvs, httperrors.NewFieldViolation("price", "OUT_OF_RANGE", "price must be positive"))
    }
    return httperrors.ValidationError(fvs).Err()
}

type BatchCreateRequest struct {
    Items []*Item `json:"items"`
}

func (r *BatchCreateRequest) Validate() error {
    var fvs httperrors.FieldViolations
    if len(r.Items) == 0 {
        fvs = append(fvs, httperrors.NewFieldViolation("items", "REQUIRED", "at least one item is required"))
    }
    for i, item := range r.Items {
        if err := item.Validate(); err != nil {
            // 用 items[0]、items[1] 等作为字段前缀
            fvs = append(fvs, httperrors.ToFieldViolations(err, fmt.Sprintf("items[%d]", i))...)
        }
    }
    return httperrors.ValidationError(fvs).Err()
}
```

### JSON 返回

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "items[0].name",
      "code": "REQUIRED",
      "message": "name is required"
    },
    {
      "field": "items[2].price",
      "code": "OUT_OF_RANGE",
      "message": "price must be positive"
    }
  ]
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "items[0].name",
      "code": "REQUIRED",
      "message": "name is required",
      "localizedMessage": "必填"
    },
    {
      "field": "items[2].price",
      "code": "OUT_OF_RANGE",
      "message": "price must be positive",
      "localizedMessage": "超出范围"
    }
  ]
}
```

---

## 10. 混合嵌套 + 数组校验

**场景**：复杂表单含嵌套对象内部有数组

```go
type Team struct {
    Name    string    `json:"name"`
    Members []*Member `json:"members"`
}

func (t *Team) Validate() error {
    var fvs httperrors.FieldViolations
    if t.Name == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("name", "REQUIRED", "team name is required"))
    }
    for i, m := range t.Members {
        if err := m.Validate(); err != nil {
            fvs = append(fvs, httperrors.ToFieldViolations(err, fmt.Sprintf("members[%d]", i))...)
        }
    }
    return httperrors.ValidationError(fvs).Err()
}

type CreateProjectRequest struct {
    Title string `json:"title"`
    Team  *Team  `json:"team"`
}

func (r *CreateProjectRequest) Validate() error {
    var fvs httperrors.FieldViolations
    if r.Title == "" {
        fvs = append(fvs, httperrors.NewFieldViolation("title", "REQUIRED", "title is required"))
    }
    if r.Team != nil {
        if err := r.Team.Validate(); err != nil {
            fvs = append(fvs, httperrors.ToFieldViolations(err, "team")...)
        }
    }
    return httperrors.ValidationError(fvs).Err()
}
```

### JSON 返回

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    { "field": "title", "code": "REQUIRED", "message": "title is required" },
    {
      "field": "team.name",
      "code": "REQUIRED",
      "message": "team name is required"
    },
    {
      "field": "team.members[1].email",
      "code": "INVALID_FORMAT",
      "message": "invalid email format"
    }
  ]
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "title",
      "code": "REQUIRED",
      "message": "title is required",
      "localizedMessage": "必填"
    },
    {
      "field": "team.name",
      "code": "REQUIRED",
      "message": "team name is required",
      "localizedMessage": "必填"
    },
    {
      "field": "team.members[1].email",
      "code": "INVALID_FORMAT",
      "message": "invalid email format",
      "localizedMessage": "格式无效"
    }
  ]
}
```

---

## 11. 带 i18n 模板参数的校验

**场景**：翻译模板中包含动态参数（如最小长度、范围值）

### i18n CSV

```csv
key,en,zh
REQUIRED,Required,必填
TOO_SHORT,Must be at least %d characters,至少需要%d个字符
OUT_OF_RANGE,Must be between %d and %d,必须在%d到%d之间
```

### Service 层

```go
func (r *CreateUserRequest) Validate() error {
    var fvs []*httperrors.FieldViolation
    if len(r.Password) < 8 {
        fvs = append(fvs, httperrors.NewFieldViolation("password", "TOO_SHORT", "must be at least 8 characters").
            WithLocalizedArgs(8))   // 翻译时替换 %d → 8
    }
    if r.Age < 18 || r.Age > 120 {
        fvs = append(fvs, httperrors.NewFieldViolation("age", "OUT_OF_RANGE", "must be between 18 and 120").
            WithLocalizedArgs(18, 120))  // 翻译时替换两个 %d
    }
    return httperrors.ValidationError(fvs).Err()
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "password",
      "code": "TOO_SHORT",
      "message": "must be at least 8 characters",
      "localizedMessage": "至少需要8个字符"
    },
    {
      "field": "age",
      "code": "OUT_OF_RANGE",
      "message": "must be between 18 and 120",
      "localizedMessage": "必须在18到120之间"
    }
  ]
}
```

### 使用自定义 i18n key

当 `reason` 不适合作为 i18n key 时，用 `WithLocalized` 指定自定义 key：

```go
httperrors.NewFieldViolation("name", "TOO_LONG", "name exceeds maximum length").
    WithLocalized("custom.name_too_long", maxLen)
```

---

## 12. 子服务错误聚合

**场景**：一个主服务调用多个子服务，将子服务的字段校验错误聚合到统一响应中

### Service 层

```go
func (s *RegistrationService) Register(ctx context.Context, req *RegisterRequest) error {
    var fvs httperrors.FieldViolations

    // 子服务 1：用户校验
    if err := s.userService.ValidateProfile(ctx, req.Profile); err != nil {
        fvs = append(fvs, httperrors.ToFieldViolations(err, "profile")...)
    }

    // 子服务 2：地址校验
    if err := s.addressService.ValidateAddress(ctx, req.Address); err != nil {
        fvs = append(fvs, httperrors.ToFieldViolations(err, "address")...)
    }

    // 子服务 3：支付信息校验
    if err := s.paymentService.ValidatePayment(ctx, req.Payment); err != nil {
        fvs = append(fvs, httperrors.ToFieldViolations(err, "payment")...)
    }

    if len(fvs) > 0 {
        return httperrors.ValidationError(fvs).Err()
    }

    // 所有校验通过，执行业务逻辑
    return s.repo.Register(ctx, req)
}
```

### JSON 返回

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "fieldViolations": [
    {
      "field": "profile.email",
      "code": "INVALID_FORMAT",
      "message": "invalid email format"
    },
    {
      "field": "address.zipCode",
      "code": "INVALID_FORMAT",
      "message": "invalid zip code"
    },
    {
      "field": "payment.cardNumber",
      "code": "REQUIRED",
      "message": "card number is required"
    }
  ]
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "INVALID_ARGUMENT",
  "message": "invalid argument",
  "localizedMessage": "参数无效",
  "fieldViolations": [
    {
      "field": "profile.email",
      "code": "INVALID_FORMAT",
      "message": "invalid email format",
      "localizedMessage": "格式无效"
    },
    {
      "field": "address.zipCode",
      "code": "INVALID_FORMAT",
      "message": "invalid zip code",
      "localizedMessage": "格式无效"
    },
    {
      "field": "payment.cardNumber",
      "code": "REQUIRED",
      "message": "card number is required",
      "localizedMessage": "必填"
    }
  ]
}
```

---

## 13. 带 Metadata 的业务错误

**场景**：错误需要携带结构化的附加信息

### Service 层

```go
func (s *TransferService) Transfer(ctx context.Context, fromID, toID string, amount int64) error {
    balance, _ := s.accountService.GetBalance(ctx, fromID)
    if balance < amount {
        return httperrors.New(http.StatusBadRequest, httperrors.ReasonBadRequest, "insufficient balance for transfer").
            WithMetadata(map[string]string{
                "currentBalance":  fmt.Sprintf("%d", balance),
                "requestedAmount": fmt.Sprintf("%d", amount),
                "currency":        "CNY",
            }).Err()
    }
    return s.repo.Transfer(ctx, fromID, toID, amount)
}
```

### JSON 返回

```json
{
  "code": "BAD_REQUEST",
  "message": "insufficient balance for transfer",
  "metadata": {
    "currentBalance": "5000",
    "requestedAmount": "10000",
    "currency": "CNY"
  }
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "BAD_REQUEST",
  "message": "insufficient balance for transfer",
  "localizedMessage": "无效的请求参数",
  "metadata": {
    "currentBalance": "5000",
    "requestedAmount": "10000",
    "currency": "CNY"
  }
}
```

---

## 14. 请求体解析失败

**场景**：JSON 格式错误或 Content-Type 不匹配

### Handler 层

```go
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        panic(httperrors.New(http.StatusBadRequest, httperrors.ReasonBadRequest, "invalid JSON request body").
            WithCause(err).Err())
    }
    // ...
}
```

### JSON 返回

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid JSON request body"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "BAD_REQUEST",
  "message": "invalid JSON request body",
  "localizedMessage": "无效的请求参数"
}
```

> **要点**：用 `WithCause(err)` 保留原始解析错误，但 `message` 不暴露内部细节。原始错误可通过 Hook 记录。

---

## 15. 限流（429 Too Many Requests）

### 中间件 / Service 层

```go
func (s *RateLimiter) Check(ctx context.Context, clientIP string) error {
    if s.isRateLimited(clientIP) {
        return httperrors.New(http.StatusTooManyRequests, httperrors.ReasonResourceExhausted, "too many requests").
            WithMetadata(map[string]string{
                "retryAfter": "60",
            }).Err()
    }
    return nil
}
```

### JSON 返回

```json
{
  "code": "RESOURCE_EXHAUSTED",
  "message": "too many requests",
  "metadata": {
    "retryAfter": "60"
  }
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "RESOURCE_EXHAUSTED",
  "message": "too many requests",
  "localizedMessage": "资源耗尽",
  "metadata": {
    "retryAfter": "60"
  }
}
```

---

## 16. 超时（504 Gateway Timeout）

### Service 层

```go
func (s *ReportService) Generate(ctx context.Context, req *ReportRequest) (*Report, error) {
    report, err := s.externalAPI.FetchReport(ctx, req)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            // context.DeadlineExceeded 会被 httperrors.Convert 自动识别为 504
            return nil, err
        }
        return nil, httperrors.WrapStatus(err, http.StatusInternalServerError, "report generation failed")
    }
    return report, nil
}
```

### JSON 返回

```json
{
  "code": "DEADLINE_EXCEEDED",
  "message": "context deadline exceeded"
}
```

### JSON 返回（中文 i18n）

```json
{
  "code": "DEADLINE_EXCEEDED",
  "message": "context deadline exceeded",
  "localizedMessage": "请求超时"
}
```

---

## 速查表

| 场景          | HTTP Status | Reason 常量               | 推荐 API                                                   |
| ------------- | ----------- | ------------------------- | ---------------------------------------------------------- |
| 字段校验错误  | 422         | `ReasonInvalidArgument`   | `ValidationError(fvs)`                                     |
| 参数无效      | 400         | `ReasonBadRequest`        | `Error(400, ReasonBadRequest, msg)`                        |
| 资源不存在    | 404         | `ReasonNotFound`          | `Errorf(404, ReasonNotFound, "user %s not found", id)`     |
| 未认证        | 401         | `ReasonUnauthenticated`   | `Error(401, ReasonUnauthenticated, msg)`                   |
| 权限不足      | 403         | `ReasonPermissionDenied`  | `New(403, ReasonPermissionDenied, msg).WithMetadata(...)`  |
| 资源冲突      | 409         | `ReasonAlreadyExists`     | `New(409, ReasonAlreadyExists, msg).WithMetadata(...)`     |
| 内部错误      | 500         | `ReasonInternal`          | `WrapStatus(err, 500, msg)`                                |
| 限流          | 429         | `ReasonResourceExhausted` | `New(429, ReasonResourceExhausted, msg).WithMetadata(...)` |
| 超时          | 504         | `ReasonDeadlineExceeded`  | 自动（`context.DeadlineExceeded`）                         |
| JSON 解析失败 | 400         | `ReasonBadRequest`        | `New(400, ReasonBadRequest, msg).WithCause(err)`           |

## 核心原则

1. **reason 用 `Reason*` 常量**：不硬编码字符串，如 `ReasonNotFound`、`ReasonBadRequest`
2. **收集而非快速返回**：字段校验时一次性收集所有错误，不要遇到第一个就 return
3. **Service 层 return，Handler 层 panic**：保持清晰的错误传播边界
4. **reason 既是错误码也是 i18n key**：默认行为，除非用 `WithLocalized` 覆盖
5. **message 给开发者，localizedMessage 给用户**：`message` 保持英文原始值，翻译自动写入 `localizedMessage`
6. **WrapStatus 保留错误链**：内部错误用 `WrapStatus` 而非 `Error`，保留堆栈 + cause
7. **metadata 传结构化上下文**：需要前端使用的附加信息放 `metadata`，不要塞进 `message`
