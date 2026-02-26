# 新 HTTP 错误处理包设计提案

## 1. 背景与定位

**这不是 `statusx` 的迁移或重构，而是一个全新的包。** 它吸取 `statusx` 在结构化错误处理方面的优秀设计，抛弃因 gRPC 协议而存在的部分，为纯 HTTP 服务场景重新设计。

`statusx` 的设计围绕 gRPC Status 协议：`codes.Code` 作为状态码、`errdetails.ErrorInfo` 作为错误信息容器、`proto.Message` 作为详情载体、`protobuf Any` 作为 i18n 参数传输格式。这些选择在 gRPC 场景下是合理的——它们是 gRPC 规范规定的字段和结构。但在纯 HTTP 场景中，这些 gRPC 规范结构不再有存在的必要。

未来项目统一使用 HTTP 通信，服务会直接从 gRPC 切换到 HTTP，不存在新旧包共存的情况。

## 2. statusx 结构逐项审查：哪些源于 gRPC 规范，哪些值得保留

### 2.1 `Status` 核心字段

| statusx 字段                      | 来源                                                                                                                                                 | HTTP 包是否需要                   | 分析                                                                                                                                                                    |
| --------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `code codes.Code`                 | **gRPC 规范**。gRPC 定义了 16 个状态码（OK, CANCELED, UNKNOWN...），是 gRPC Status 的必要字段                                                        | **替换** → `httpStatus int`       | HTTP 有自己的状态码体系（400, 404, 500...），直接使用 `int` 或 `net/http` 常量即可，无需中间映射层                                                                      |
| `message string`                  | 协议无关的通用设计                                                                                                                                   | **保留**                          | 调试用的人类可读消息，与协议无关                                                                                                                                        |
| `errorInfo *errdetails.ErrorInfo` | **gRPC 规范**。`google.rpc.ErrorInfo` 是 Google 定义的标准错误详情 protobuf message，包含 `reason`、`domain`、`metadata` 三个字段                    | **拆解**                          | `reason` 和 `metadata` 是有价值的概念，但不需要通过 `errdetails.ErrorInfo` 这个 protobuf 容器来承载。直接用 Go 原生字段：`reason string` + `metadata map[string]string` |
| `localized *statusv1.Localized`   | **自定义 protobuf**。项目自定义的 protobuf message，包含 `key` 和 `args []*anypb.Any`，设计为 protobuf 是为了能作为 gRPC Status detail 传输          | **简化**                          | i18n 模板（key + args）的概念值得保留，但不需要 protobuf 序列化。直接用 Go struct：`Localized{key string, args []any}`                                                  |
| `badRequest *statusv1.BadRequest` | **自定义 protobuf**。扩展了 Google 标准 `google.rpc.BadRequest`，增加了 `Localized` 和 `LocalizedMessage` 字段，设计为 protobuf 同样是为了 gRPC 传输 | **简化**                          | 字段验证违规的概念值得保留，但容器从 protobuf message 变为 Go struct 即可                                                                                               |
| `details []proto.Message`         | **gRPC 规范**。gRPC Status 的 details 字段，类型为 `repeated google.protobuf.Any`，要求所有附加信息都是 protobuf message                             | **替换** → `details []any` 或去除 | HTTP JSON 响应没有这个约束。需要评估是否还需要一个通用的"附加详情"机制，如果需要可以用 `[]any`（JSON 可序列化即可）                                                     |
| `cause error`                     | 协议无关的通用设计                                                                                                                                   | **保留**                          | 错误链和 stacktrace，与协议无关                                                                                                                                         |

### 2.2 `StatusError` 接口实现

| 方法/接口                        | 来源                                                 | HTTP 包是否需要          | 分析                                                                                |
| -------------------------------- | ---------------------------------------------------- | ------------------------ | ----------------------------------------------------------------------------------- |
| `Error() string`                 | Go 标准                                              | **保留**                 |                                                                                     |
| `Unwrap() error`                 | Go 标准                                              | **保留**                 |                                                                                     |
| `Cause() error`                  | `pkg/errors` 兼容                                    | **保留**                 |                                                                                     |
| `Is(target error) bool`          | Go 标准                                              | **保留**，但实现需要变化 | statusx 中使用 `proto.Equal` 比较两个 gRPC Status proto，新包需要自行实现结构化比较 |
| `Format(s fmt.State, verb rune)` | `pkg/errors` 兼容                                    | **保留**                 | 支持 `%+v` 打印 stacktrace                                                          |
| `GRPCStatus() *status.Status`    | **gRPC 规范**。gRPC 框架通过此接口识别自定义错误类型 | **去除**                 | HTTP 场景不需要                                                                     |

### 2.3 `FieldViolation` 结构

| 字段                                            | 来源                                                                           | HTTP 包是否需要                                 | 分析                               |
| ----------------------------------------------- | ------------------------------------------------------------------------------ | ----------------------------------------------- | ---------------------------------- |
| `field string`                                  | 通用设计（Google `BadRequest.FieldViolation` 中也有）                          | **保留**                                        |                                    |
| `reason string`                                 | 通用设计                                                                       | **保留**                                        |                                    |
| `description string`                            | 通用设计                                                                       | **保留**                                        |                                    |
| `localized *Localized`                          | 项目设计，承载 i18n 模板                                                       | **保留**，简化为 Go struct                      |                                    |
| `localizedMessage *errdetails.LocalizedMessage` | **gRPC 规范**。`google.rpc.LocalizedMessage` 是 Google 定义的 protobuf message | **替换** → Go struct `{locale, message string}` | 概念保留（已翻译的消息），容器简化 |
| `Proto()` 方法                                  | gRPC 序列化需要                                                                | **去除**                                        | 不再需要 protobuf 序列化           |

### 2.4 翻译系统

| 组件                            | 来源              | HTTP 包是否需要            | 分析                                                                                                                   |
| ------------------------------- | ----------------- | -------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `Localized{key, args}`          | 项目设计          | **保留**，直接用 Go struct | `args` 从 `[]*anypb.Any` 变为 `[]any`，删除全部 `toProtoMessage` / `extractValueFromAny` / `convertArgsToAny` 转换代码 |
| `TranslateError` / `Translated` | 项目设计          | **保留**                   | 翻译优先级逻辑与协议无关                                                                                               |
| `TranslateStatusErrorOnly`      | 项目设计          | **保留**                   |                                                                                                                        |
| `LocalizedFromProto`            | protobuf 反序列化 | **去除**                   | 不再有 proto 格式                                                                                                      |

### 2.5 协议适配层

| 文件/函数                                                                                 | 来源                         | HTTP 包是否需要 | 分析                                                                                                                                |
| ----------------------------------------------------------------------------------------- | ---------------------------- | --------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `grpc.go`：`UnaryServerInterceptor`                                                       | gRPC 拦截器                  | **去除**        |                                                                                                                                     |
| `connect.go`：`UnaryConnectInterceptor`、`ConvertToConnectError`、`WriteConnectErrorOnly` | Connect 协议适配             | **去除**        |                                                                                                                                     |
| `vproto.go`：`WriteVProtoHTTPError`、`VProtoHTTPWriteErrorHook`、`HTTPStatusFromCode`     | VProto 兼容 + gRPC code 映射 | **去除**        | VProto 是为兼容旧 `theplant/validator` 协议，新包不需要。`HTTPStatusFromCode` 是 gRPC→HTTP 映射，新包直接用 HTTP 状态码，不需要映射 |
| `http.go`：`HTTPErrorWriter`（panic 恢复 + 错误写入中间件）                               | 项目设计                     | **重新设计**    | panic 恢复 + Hook 扩展的思路可以保留，但输出格式从 VProto/Connect 变为标准 JSON                                                     |
| `code.go`：`ReasonFromCode`、`NewCode`、`WrapCode`                                        | gRPC code → reason 映射      | **替换**        | 新包可以提供 HTTP status → 默认 reason 的映射                                                                                       |

### 2.6 验证集成

| 组件                                      | 来源               | HTTP 包是否需要 | 分析                                                                                             |
| ----------------------------------------- | ------------------ | --------------- | ------------------------------------------------------------------------------------------------ |
| `Validator` / `ContextValidator` 接口     | 项目设计           | **保留**        |                                                                                                  |
| proto-gen-validate 集成                   | gRPC/protobuf 生态 | **去除**        | HTTP 场景不再使用 protobuf message 作为请求体，`pgvErr` 接口和 `convertProtoGenErrToFV` 不再需要 |
| `FormatField`（camelCase + 数组索引保留） | 项目设计           | **保留**        | 字段路径格式化与协议无关                                                                         |

### 2.7 其他工具

| 组件                                                   | 来源                | HTTP 包是否需要                                 | 分析                                       |
| ------------------------------------------------------ | ------------------- | ----------------------------------------------- | ------------------------------------------ |
| `Clone` 函数                                           | 不可变设计需要      | **保留**，但实现从 `proto.Clone` 改为手动深拷贝 | `maps.Clone` + 切片拷贝即可                |
| `ExtractDetail[T]`                                     | 通用泛型工具        | **保留**（如果保留 details 机制）               |                                            |
| `AssertFieldViolations`                                | 测试辅助            | **保留**，简化实现                              | 不再依赖 `proto.Equal`，直接结构体比较     |
| `error.proto`（`ErrorReason` 枚举）                    | gRPC 规范对齐       | **去除**                                        | 新包不需要 protobuf 枚举，直接用字符串常量 |
| `localized.proto`（`Localized`、`BadRequest` message） | protobuf 序列化需要 | **去除**                                        | 全部用 Go struct 替代                      |

## 3. 新包需要实现的功能

### 3.1 核心错误模型

```go
type Status struct {
    httpStatus int               // HTTP 状态码（200, 400, 404, 500...）
    reason     string            // 错误原因码（如 "NOT_FOUND"、"RATE_LIMITED"）
    message    string            // 调试用消息
    metadata   map[string]string // 键值对元数据
    localized  *Localized        // i18n 翻译模板
    fieldViolations []*FieldViolation // 字段级验证违规
    cause      error             // 原始错误（带 stacktrace）
}
```

**与 statusx 对比变化**：

- `codes.Code` → `int`（HTTP 状态码）
- `*errdetails.ErrorInfo` → 拆解为 `reason string` + `metadata map[string]string`（去掉了 `errdetails.ErrorInfo` 中的 `domain` 字段，statusx 实际也未使用它）
- `*statusv1.Localized` → `*Localized`（纯 Go struct）
- `*statusv1.BadRequest` → `[]*FieldViolation`（纯 Go struct 切片）
- `[]proto.Message` → 去除或改为 `[]any`
- 去除 `GRPCStatus()` 方法

**需要的构造函数与链式方法**：

- `New(httpStatus int, reason, message string) *Status`
- `Newf`、`Error`、`Errorf`、`Wrap`、`Wrapf` 等便捷函数
- `NewStatus(httpStatus int, message string) *Status` — 自动从 HTTP 状态码推导 reason
- `With*` 系列链式方法（不可变设计，返回新副本）
- `StatusError` 实现 `error`、`Unwrap()`、`Is()`、`Format()` 接口

### 3.2 字段验证错误

```go
type FieldViolation struct {
    field            string
    reason           string
    description      string
    localized        *Localized        // i18n 模板（翻译前）
    localizedMessage *LocalizedMessage // 已翻译消息（翻译后）
}

type LocalizedMessage struct {
    Locale  string // 语言标签，如 "zh-CN"
    Message string // 已翻译的文本
}
```

**与 statusx 对比变化**：

- `*errdetails.LocalizedMessage`（protobuf）→ `*LocalizedMessage`（纯 Go struct）
- 去除 `Proto()` 方法
- 保留 `FormatField`、`PrependField`、`FlattenFieldViolations` 等工具函数
- `BadRequest(inputs...)` → 可更名为 `ValidationError(inputs...)` 或保持

### 3.3 国际化翻译

```go
type Localized struct {
    key  string
    args []any  // 直接使用 Go 原生类型，不需要 protobuf Any
}
```

**与 statusx 对比的简化**：

- 删除全部 protobuf Any 转换代码（`toProtoMessage`、`extractValueFromAny`、`convertArgsToAny`、`LocalizedFromProto`）
- `args` 直接存储 Go 值，无需序列化/反序列化
- 翻译优先级逻辑保持不变

### 3.4 HTTP 错误响应

**JSON 响应格式**：

```json
{
  "code": "NOT_FOUND",
  "message": "user not found",
  "localizedMessage": "用户不存在",
  "metadata": { "key": "value" },
  "fieldViolations": [
    {
      "field": "email",
      "code": "REQUIRED",
      "message": "email is required",
      "localizedMessage": "邮箱不能为空"
    }
  ]
}
```

**需要实现**：

- JSON 序列化方法（`MarshalJSON` 或独立的响应 struct）
- HTTP 中间件（panic 恢复 + 翻译 + JSON 写入）
- Hook 扩展机制
- HTTP 状态码写入（直接从 `Status.httpStatus` 取，无需映射）

### 3.5 验证集成

- 保留 `Validator` / `ContextValidator` 接口
- 保留 `Validate(ctx, input)` 统一入口
- **去除** proto-gen-validate 集成（`pgvErr` 接口、`convertProtoGenErrToFV`）
- 可考虑集成 `go-playground/validator` 等 HTTP 场景常用验证库

### 3.6 错误解析

- `FromError(err)` — 从 error 还原为 `*Status`
- `Convert(err)` — 简化版
- `StatusCode(err)` / `Reason(err)` — 快速提取
- 保留 `context.DeadlineExceeded` → 504、`context.Canceled` → 499 的自动映射

### 3.7 测试辅助

- `AssertFieldViolations(t, err, ...)` — 实现从 `proto.Equal` 改为结构体比较
- `ExtractDetail[T]`（如果保留 details 机制）

## 4. 需要考虑的设计要点

### 4.1 HTTP 状态码 → 默认 reason 映射

gRPC 的 `codes.Code` 只有 16 个枚举值，所以 statusx 的 `ReasonFromCode` / `NewCode` 可以做到一一映射。HTTP 状态码数量更多，需要决定：

**建议方案**：为常用状态码提供默认 reason 映射函数 `ReasonFromStatus(httpStatus int) string`：

| HTTP Status | 默认 reason                                               |
| ----------- | --------------------------------------------------------- |
| 400         | `BAD_REQUEST`                                             |
| 401         | `UNAUTHENTICATED`                                         |
| 403         | `PERMISSION_DENIED`                                       |
| 404         | `NOT_FOUND`                                               |
| 409         | `CONFLICT`                                                |
| 422         | `INVALID_ARGUMENT`                                        |
| 429         | `RESOURCE_EXHAUSTED`                                      |
| 500         | `INTERNAL`                                                |
| 502         | `BAD_GATEWAY`                                             |
| 503         | `UNAVAILABLE`                                             |
| 504         | `DEADLINE_EXCEEDED`                                       |
| 其他        | `"UNKNOWN"` 或 `http.StatusText` 的 UPPER_SNAKE_CASE 转换 |

### 4.2 `details []any` 是否保留

> **决策：暂不提供 `details`，保持简洁。**

statusx 中 `details []proto.Message` 的设计来源于 gRPC Status 的 `repeated google.protobuf.Any details` 字段。在新包中翻译结果直接存在对应的 Go struct 字段中，不再需要 details 作为中间产物。如果未来有需求再添加。

### 4.3 `StatusError.Is()` 的比较策略

> **决策：方案 A — 只比较 `httpStatus` + `reason`，语义清晰。**

### 4.4 中间件 panic 恢复模式

> **决策：保留 panic 模式，保持 handler 签名优雅（标准 `http.HandlerFunc`）。代码中留好注释，未来可能切换到返回 error 模式。**

statusx 的 `HTTPErrorWriter` 通过 `defer recover()` 捕获 handler 中 panic 出来的 error。新包沿用此模式。

### 4.5 Stacktrace 策略

> **决策：短期继续使用 `pkg/errors`，将来作为独立议题处理。**

### 4.6 JSON 字段命名与响应结构

> **决策：camelCase，不在 body 中包含冗余的 HTTP status code（HTTP 响应头已有）。**

## 5. 代码切换指南

服务直接从 gRPC 切换到 HTTP，以下是 statusx 代码切换到新包的对照参考。

### 5.1 错误创建

```go
// statusx（旧）
statusx.New(codes.NotFound, "NOT_FOUND", "user not found")
statusx.New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "internal error")
statusx.NewCode(codes.NotFound, "user not found")
statusx.Wrap(err, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "failed")
statusx.WrapCode(err, codes.Internal, "failed")

// 新包
httperrors.New(http.StatusNotFound, "NOT_FOUND", "user not found")
httperrors.New(http.StatusInternalServerError, "INTERNAL", "internal error")
httperrors.NewStatus(http.StatusNotFound, "user not found")  // 自动推导 reason
httperrors.Wrap(err, http.StatusInternalServerError, "INTERNAL", "failed")
httperrors.WrapStatus(err, http.StatusInternalServerError, "failed")
```

**注意**：statusx 中大量使用 `statusv1.ErrorReason_XXX.String()` 作为 reason，这是 protobuf 枚举的 String() 方法。新包直接使用字符串常量即可，可以提供常量包：

```go
const (
    ReasonNotFound         = "NOT_FOUND"
    ReasonInternal         = "INTERNAL"
    ReasonInvalidArgument  = "INVALID_ARGUMENT"
    ReasonPermissionDenied = "PERMISSION_DENIED"
    // ...
)
```

### 5.2 链式方法

```go
// statusx（旧）
statusx.New(...).WithMetadata(md).WithLocalized("key", args...).WithFieldViolations(fvs...).Err()

// 新包（完全一致的链式 API）
httperrors.New(...).WithMetadata(md).WithLocalized("key", args...).WithFieldViolations(fvs...).Err()
```

链式方法 API 保持一致，只是入参类型变化（`codes.Code` → `int`）。

### 5.3 错误解析

```go
// statusx（旧）
st := statusx.Convert(err)
st.Code()   // codes.Code 类型
st.Reason()

// 新包
st := httperrors.Convert(err)
st.StatusCode()  // int 类型
st.Reason()
```

### 5.4 字段验证

```go
// statusx（旧）
statusx.BadRequest(
    statusx.NewFieldViolation("email", "REQUIRED", "email is required"),
)

// 新包（API 一致）
httperrors.ValidationError(  // 或 httperrors.BadRequest
    httperrors.NewFieldViolation("email", "REQUIRED", "email is required"),
)
```

### 5.5 拦截器/中间件

```go
// statusx（旧）— gRPC 拦截器
grpc.UnaryInterceptor(statusx.UnaryServerInterceptor(i18n))

// statusx（旧）— Connect 拦截器
statusx.UnaryConnectInterceptor(i18n, shouldConvert)

// statusx（旧）— HTTP 中间件（VProto）
statusx.NewVProtoHTTPErrorWriter(i18n)

// 新包 — 只有 HTTP 中间件
httperrors.ErrorMiddleware(i18n)
```

### 5.6 需要特别注意的切换点

1. **`statusv1.ErrorReason_XXX.String()`** → 字符串常量。全项目搜索 `statusv1.ErrorReason_` 替换
2. **`codes.XXX`** → `http.StatusXXX` 或 `net/http` 常量。全项目搜索 `codes.` 替换
3. **`st.Code()`** 返回值类型从 `codes.Code` 变为 `int`，涉及 switch 语句需要调整
4. **`st.Details()`** 如果新包去除 details 机制，依赖 `ExtractDetail` 的代码需要重写
5. **`FromError` 中对 gRPC `status.FromError` 的调用** → 新包只需处理自己的 `StatusError` 类型和标准 error

## 6. 预期依赖

### 不再需要的依赖

| 包                                             | 原因                                                  |
| ---------------------------------------------- | ----------------------------------------------------- |
| `google.golang.org/grpc`                       | gRPC codes、status、拦截器                            |
| `google.golang.org/protobuf`                   | proto.Message、proto.Clone、Any、wrapperspb 等        |
| `google.golang.org/genproto`                   | errdetails（ErrorInfo、BadRequest、LocalizedMessage） |
| `connectrpc.com/connect`                       | Connect 协议适配                                      |
| `github.com/theplant/validator`                | VProto 兼容                                           |
| `github.com/grpc-ecosystem/go-grpc-middleware` | gRPC metadata 提取                                    |
| `buf` / `protoc-gen-validate`                  | protobuf 代码生成工具链                               |

### 继续使用的依赖

| 包                           | 原因                        |
| ---------------------------- | --------------------------- |
| `github.com/qor5/x/v3/i18nx` | 国际化核心，与协议无关      |
| `github.com/qor5/x/v3/hook`  | Hook 扩展机制               |
| `github.com/qor5/x/v3/httpx` | HTTP 工具（Header 常量等）  |
| `github.com/pkg/errors`      | stacktrace 捕获（短期保留） |
| `github.com/samber/lo`       | 泛型工具函数                |
| `golang.org/x/text/language` | 语言标签处理                |
| `encoding/json`（标准库）    | JSON 序列化                 |

## 7. 建议的包名与路径

候选方案：

> **决策：`github.com/qor5/x/v3/httperrors`**
