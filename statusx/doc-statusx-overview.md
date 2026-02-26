# statusx 包概览

## 1. 包的目的

`statusx` 是一个基于 gRPC Status 协议构建的**结构化错误处理包**。它为项目提供了统一的错误创建、传播、翻译和序列化机制，核心目标是：

- **统一错误模型**：基于 gRPC `codes.Code` + `reason` + `message` 三元组，提供一致的错误表示
- **字段级验证错误（Field Violations）**：支持表单/请求级别的字段校验错误，可以携带多个字段的多个违规信息
- **国际化（i18n）翻译**：错误消息支持多语言翻译，通过 `Localized` 模板 + `i18nx` 实现
- **跨协议兼容**：同一个 `Status` 对象可以序列化为 gRPC Status、Connect Error、VProto HTTP Error 等多种协议格式

## 2. 核心数据结构

### 2.1 `Status` 结构体（`status.go`）

```go
type Status struct {
    code       codes.Code                  // gRPC 状态码
    message    string                      // 调试用的人类可读消息
    errorInfo  *errdetails.ErrorInfo       // 包含 reason 和 metadata
    localized  *statusv1.Localized         // i18n 翻译模板（key + args）
    badRequest *statusv1.BadRequest        // 字段级验证违规
    details    []proto.Message             // 其他附加详情
    cause      error                       // 原始错误（带 stacktrace）
}
```

这是整个包的核心，采用**不可变设计**（immutable）——所有 `With*` 方法返回新的 `Status` 副本（通过 `Clone`）。

### 2.2 `StatusError`（`status.go`）

实现了 `error`、`GRPCStatus()`、`Unwrap()`、`Is()`、`Format()` 等接口，使得 `Status` 可以作为标准 Go error 在调用链中传播。

### 2.3 `FieldViolation` / `FieldViolations`（`field.go`）

字段级验证错误，每个 `FieldViolation` 包含：
- `field`：字段路径（支持嵌套如 `user.addresses[0].street`）
- `reason`：错误原因码
- `description`：人类可读描述
- `localized`：i18n 翻译模板
- `localizedMessage`：已翻译的消息（优先级高于 `localized`）

### 2.4 `Localized`（`translate.go`）

本地化模板，包含 `key`（i18n 键）和 `args`（模板参数）。支持 Go 原生类型与 protobuf `Any` 之间的双向转换。

## 3. 核心功能模块

### 3.1 错误创建（`status.go`、`code.go`）

| 函数 | 说明 |
|------|------|
| `New(code, reason, message)` | 创建 Status，自动捕获 stacktrace |
| `Newf(code, reason, format, args...)` | 格式化版本 |
| `NewCode(code, message)` | 自动从 code 推导 reason |
| `Error(code, reason, message)` | 直接返回 error |
| `Wrap(err, code, reason, message)` | 包装已有错误 |
| `WrapCode(err, code, message)` | 包装已有错误，自动推导 reason |
| `BadRequest(inputs...)` | 快速创建字段验证错误 |

### 3.2 错误元数据（`status.go`）

通过链式调用丰富错误信息：
- `WithCode()` / `WithMessage()` / `WithReason()` — 修改基本属性
- `WithCause()` — 设置原始错误
- `WithMetadata()` — 附加键值对元数据
- `WithLocalized(key, args...)` — 设置 i18n 翻译 key 和参数
- `WithLocalizedArgs(args...)` — 只设置翻译参数，保留已有 key
- `WithFieldViolations(...)` — 添加字段级验证违规
- `WithFlattenFieldViolations(...)` — 扁平化多种输入类型的字段违规
- `WithDetails(...)` — 附加自定义 protobuf 详情

### 3.3 错误解析（`status.go`）

| 函数 | 说明 |
|------|------|
| `FromError(err)` | 从 error 还原为 `*Status`，支持 `StatusError`、gRPC Status、`context.DeadlineExceeded/Canceled` |
| `Convert(err)` | `FromError` 的简化版，忽略 ok 标志 |
| `Code(err)` / `Reason(err)` | 快速提取错误码/原因 |

### 3.4 翻译系统（`translate.go`）

翻译优先级：
1. **已存在 `LocalizedMessage`** → 跳过翻译（最高优先级）
2. **存在 `Localized` 模板** → 用 i18n 翻译模板（中优先级）
3. **使用 error reason** → 用 reason 作为 i18n key 翻译（兜底）

核心函数：
- `TranslateError(err, i18n, lang)` — 翻译错误消息和字段违规
- `TranslateStatusErrorOnly(err, i18n, lang)` — 只翻译 `StatusError` 类型
- `Status.Translated(i18n, lang)` — 返回翻译后的新 Status

支持 Go 原生类型 ↔ protobuf `Any` 的双向转换（`toProtoMessage` / `extractValueFromAny`），用于在 protobuf 传输中携带 i18n 模板参数。

### 3.5 验证集成（`validate.go`）

`Validate(ctx, input)` 函数按优先级尝试多种验证接口：
1. `ContextValidator`（带 context 的验证）
2. `ValidateAll()`（proto-gen-validate，支持批量错误）
3. `Validator`（简单验证）

自动将 `proto-gen-validate` 的错误转换为 `FieldViolation`，并通过 `FormatField` 处理字段路径的 camelCase 转换和数组索引保留。

### 3.6 协议适配层

#### gRPC 拦截器（`grpc.go`）
- `UnaryServerInterceptor(i18n)` — gRPC 一元拦截器，注入 i18n context 并翻译响应错误

#### Connect 适配（`connect.go`）
- `UnaryConnectInterceptor(i18n, shouldConvert)` — Connect 协议拦截器
- `ConvertToConnectError(err)` — 将 StatusError 转换为 `connect.Error`
- `WriteConnectErrorOnly(errWriter, w, r, err)` — 仅写入 Connect 错误格式

#### HTTP/VProto 适配（`http.go`、`vproto.go`）
- `HTTPErrorWriter(conf)` — HTTP 中间件，捕获 panic 中的错误并写入响应
- `WriteVProtoHTTPError(err, w, r)` — 将错误写为 VProto 格式（兼容 `theplant/validator`）
- `HTTPStatusFromCode(code)` — gRPC code → HTTP status code 映射
- 支持 JSON / Proto / x.prottp 三种内容协商格式

### 3.7 自定义 Protobuf 定义（`proto/`）

- **`error.proto`**：`ErrorReason` 枚举，与 gRPC codes 一一对应
- **`localized.proto`**：
  - `Localized` message：i18n key + args（`google.protobuf.Any` 类型）
  - `BadRequest` message：扩展了 Google 标准 `BadRequest`，增加了 `Localized` 和 `LocalizedMessage` 字段

## 4. 外部依赖

### 4.1 gRPC / Protobuf 生态（核心依赖）

| 包 | 用途 |
|----|------|
| `google.golang.org/grpc` | gRPC 核心：`codes.Code`、`status.Status`、拦截器 |
| `google.golang.org/grpc/codes` | gRPC 状态码定义 |
| `google.golang.org/grpc/status` | gRPC Status 对象 |
| `google.golang.org/genproto/googleapis/rpc/errdetails` | Google 标准错误详情（`ErrorInfo`、`BadRequest`、`LocalizedMessage`） |
| `google.golang.org/protobuf` | Protobuf 核心：`proto.Message`、`proto.Clone`、`anypb`、`wrapperspb` 等 |
| `google.golang.org/protobuf/protoadapt` | Protobuf V1/V2 适配 |

### 4.2 Connect 生态

| 包 | 用途 |
|----|------|
| `connectrpc.com/connect` | Connect 协议支持：`connect.Error`、`connect.ErrorWriter`、拦截器 |

### 4.3 项目内部依赖

| 包 | 用途 |
|----|------|
| `github.com/qor5/x/v3/i18nx` | 国际化支持 |
| `github.com/qor5/x/v3/httpx` | HTTP 工具（Header 常量等） |
| `github.com/qor5/x/v3/jsonx` | JSON 序列化工具 |
| `github.com/qor5/x/v3/hook` | Hook 机制（用于 HTTP error writer 扩展） |
| `github.com/qor5/x/v3/statusx/gen/status/v1` | 自动生成的 protobuf Go 代码 |

### 4.4 第三方库

| 包 | 用途 |
|----|------|
| `github.com/pkg/errors` | 错误包装和 stacktrace 捕获 |
| `github.com/samber/lo` | 泛型工具函数（`lo.Map`、`lo.Find`、`lo.CamelCase`） |
| `github.com/theplant/validator/proto` | VProto 格式的 `ValidationError`（兼容旧系统） |
| `github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata` | gRPC metadata 提取 |
| `golang.org/x/text/language` | 语言标签处理 |

### 4.5 Protobuf 生成工具

| 工具 | 用途 |
|------|------|
| `buf` | Protobuf 构建管理工具 |
| `protoc-gen-validate` | Protobuf 字段验证代码生成 |

## 5. 设计特点总结

1. **以 gRPC Status 为核心模型**：所有错误最终都能序列化为 `grpc/status.Status`，通过 protobuf 详情（details）携带丰富的结构化信息
2. **不可变设计**：所有修改操作返回新副本，避免并发问题
3. **翻译与错误分离**：错误创建时只携带翻译模板（key + args），翻译在拦截器层统一执行
4. **多协议输出**：同一个 Status 可以输出为 gRPC、Connect、VProto HTTP 等不同格式
5. **字段验证一等支持**：`FieldViolation` 支持嵌套路径、数组索引、per-field 翻译
6. **兼容性层**：通过 `vproto.go` 兼容旧的 `theplant/validator` 协议
