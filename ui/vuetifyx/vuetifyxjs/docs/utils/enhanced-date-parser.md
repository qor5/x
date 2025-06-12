# 增强日期解析器 (EnhancedDateParser)

增强的日期解析器，专门为解决dayjs对中文年月日格式和各种非标准日期格式支持不佳的问题而设计。

## 功能特性

### 1. 支持更多日期格式

- **标准格式**: `2023-12-25`, `2023/12/25`, `2023.12.25`
- **中文格式**: `2023年12月25日`, `23年12月25日`
- **中文数字**: `二零二三年十二月二十五日`, `二零二三年一月一日`
- **短格式**: `12月25日`, `12-25`, `12/25`
- **带时间**: `2023年12月25日 14:30:00`

### 2. 智能相对日期

- **固定词汇**: `今天`, `明天`, `昨天`, `后天`, `前天`, `大后天`, `大前天`
- **相对表达**: `3天后`, `2周前`, `1个月后`, `2年前`

### 3. 智能年份处理

- 自动补全世纪: `23年12月25日` → `2023年12月25日`
- 自动补全年份: `12月25日` → `2024年12月25日` (当前年份)

## 使用方法

### 基本使用

```typescript
import { EnhancedDateParser } from '@/lib/utils/dateParser'

// 解析日期
const parsed = EnhancedDateParser.parseDate('2023年12月25日')
console.log(parsed?.format('YYYY-MM-DD')) // '2023-12-25'

// 检查日期是否有效
const isValid = EnhancedDateParser.isValidDate('2023年12月25日')
console.log(isValid) // true

// 格式化日期
const formatted = EnhancedDateParser.formatDate(new Date(), 'YYYY年MM月DD日')
console.log(formatted) // '2024年01月15日'
```

### 在DatePicker中的应用

DatePicker组件已经自动集成了增强日期解析器，无需额外配置：

```vue
<template>
  <vx-date-picker v-model="date" placeholder="支持输入: 2023年12月25日, 明天, 3天后等" />
</template>
```

### 支持的输入示例

```typescript
// 所有这些输入都能被正确解析
const examples = [
  '2023-12-25',
  '2023年12月25日',
  '二零二三年十二月二十五日',
  '23年12月25日',
  '12月25日',
  '今天',
  '明天',
  '3天后',
  '2周前',
  '1个月后'
]

examples.forEach((input) => {
  const result = EnhancedDateParser.parseDate(input)
  console.log(`${input} → ${result?.format('YYYY-MM-DD')}`)
})
```

## API 参考

### parseDate(dateInput)

解析日期字符串、数字或Date对象。

- **参数**: `dateInput: string | number | Date`
- **返回**: `dayjs.Dayjs | null`

### parseRelativeDate(dateStr)

专门解析相对日期字符串。

- **参数**: `dateStr: string`
- **返回**: `dayjs.Dayjs | null`

### formatDate(date, format)

格式化日期。

- **参数**:
  - `date: dayjs.Dayjs | string | number | Date`
  - `format: string = 'YYYY-MM-DD'`
- **返回**: `string`

### isValidDate(dateInput)

检查日期是否有效。

- **参数**: `dateInput: string | number | Date`
- **返回**: `boolean`

### getSupportedFormats()

获取支持的日期格式列表。

- **返回**: `string[]`

## 中文数字支持

解析器支持以下中文数字：

- 基础数字: 一、二、三、四、五、六、七、八、九、十
- 月份: 一月到十二月
- 日期: 一日到三十一日

## 测试

可以运行测试来验证解析器的功能：

```typescript
import { testEnhancedDateParser } from '@/lib/utils/dateParser.test'

// 在开发环境下运行测试
testEnhancedDateParser()
```

## 错误处理

当所有解析方法都失败时，解析器会：

1. 返回 `null` (对于 `parseDate` 方法)
2. 记录警告信息到控制台
3. 在DatePicker中会回退到原始dayjs解析作为最后尝试

## 性能考虑

- 解析器按照常用度排序格式，优先尝试最常见的格式
- 使用惰性匹配，找到匹配格式后立即返回
- 正则表达式经过优化，避免回溯问题
