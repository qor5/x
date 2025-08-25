# jsonx

一个专门的 JSON 处理包，增强了对 Protocol Buffers 的支持。

## 特点

- **Marshal 方法**：当处理`proto.Message`或其指针（无论嵌套多深）时，内部使用`protojson`加上`EmitUnpopulated(true)`选项进行序列化，以确保所有数据消费者之间的一致性。

- **Unmarshal 方法**：当处理`proto.Message`或其多层指针时，内部使用`protojson`加上`DiscardUnknown(true)`选项进行反序列化，以确保向前兼容性。

- **Patch 方法**：`Unmarshal`的增强版，能够保留原始值。

- **Copy 方法**：基于`Patch`实现，用于对象间的数据复制。

## 处理 Nil / Null

序列化 nil 值时：

- `protojson.Marshal`对 nil 的根 Protocol Buffer 消息返回空对象（`"{}"`）
- 这与标准的`json.Marshal`行为不同，后者通常会返回`"null"`
- jsonx 包有意遵循`json.Marshal`的行为，对 nil 的 Protocol Buffer 消息返回`"null"`而非空对象，提供更直观和一致的结果

反序列化整个 JSON 字符串为 `"null"` 时：

- `json.Unmarshal`在遇到整个 JSON 为`"null"`时不会产生错误，它只是保持消息不变
- `protojson.Unmarshal`在遇到整个 JSON 为`"null"`时会返回错误，错误信息类似"unexpected token null"
- jsonx 包遵循`json.Unmarshal`的行为，确保在将整个 JSON 为`"null"`反序列化到目标时不会发生错误

## Patch vs. Unmarshal

`Patch`与标准反序列化方法的关键区别：

1. `json.Unmarshal`可以将补丁字符串反序列化到目标结构体，同时保留补丁中未指定字段的原始值。但是，当目标是嵌套的`map[string]any`时，它会丢失原始值。

2. `protojson.Unmarshal`与标准 JSON 反序列化器不同，不保留原始值。

3. `Patch`通过在复杂嵌套结构中也保留原始值来解决这些限制。

## 与 RFC7386 的差异

`Patch`未完全遵循 RFC7386 JSON Merge Patch 标准，原因如下：

1. 在 RFC7386 中，补丁中的`null`值表示删除目标 JSON 中的原始字段。

2. 在 Go 的典型反序列化场景中应用于结构体时，`null`通常用于覆盖原始值。

3. `Patch`首先使用`Marshal(dest)`创建一个可以恢复原始值的补丁，然后与用户提供的补丁合并，创建满足所需用例的最终补丁。

## 重要说明

对于数组字段，`Patch`、`Copy`和 RFC7386 都使用全量替换。如果您的应用程序需要对数组进行增量修改，您应该：

1. 在调用`Patch`/`Copy`之前克隆原始信息
2. 调用后手动处理增量变化
