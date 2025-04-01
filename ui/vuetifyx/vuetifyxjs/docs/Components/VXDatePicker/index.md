# 日期选择器

`vx-date-picker` 组件用于选择日期。

**目录**：

- [Props](#props)
- [Events](#events)
- [值类型](#值类型)
- [组件类型](#组件类型)
- [日期格式化 format](#日期格式化-format)
- [配合 datePickerProps 控制](#配合-datepickerprops-控制)
- [各种状态](#各种状态)
- [blur 事件示例](#blur-事件示例)

### Props

#### 公共props

> 该组件继承了 [vx-field的props](/Components/VXField/#props)，所以是通用的，以下仅罗列 vx-date-picker 特有的Props

| Name      | Introduction                         | Type      | Default Value |
| --------- | ------------------------------------ | --------- | ------------- |
| type      | 见[组件类型](./#组件类型)            | `String`  | `datepicker`  |
| format    | 见[日期格式化](./#日期格式化-format) | `String`  | 'YYYY-MM-DD'  |
| clearable | 显示可清除样式                       | `Boolean` | `false`       |

#### vx-date-picker props

| Name            | Introduction                                                                          | Type                     | Default Value |
| --------------- | ------------------------------------------------------------------------------------- | ------------------------ | ------------- |
| placeholder     | 占位文本                                                                              | `String`                 | -             |
| datePickerProps | vuetify [原生参数](https://vuetifyjs.com/en/api/v-date-picker/)，用于控制时间选择组件 | `Object`                 | -             |
| modelValue      | 绑定的值，见[值类型](./#支持的值类型)                                                 | `String` `Number` `Date` | -             |
| hideAppendInner | 隐藏内部的附加图标                                                                    | `Boolean`                | `false`       |

### Events

| Name | Payload  | Introduction                                                                                |
| ---- | -------- | ------------------------------------------------------------------------------------------- |
| blur | `string` | 某些情况如果期望失焦或者关闭了下拉后获取值时用这个事件，查看[blur 事件示例](#blur-事件示例) |

### 值类型

- 组件拥有很强的传入值适应性，可以传入各种类型的值并格式化字符串、时间戳、日期类型都可，最终都会被格式化成默认格式或者传入的format格式

:::demo

```vue
<template>
  <v-row>
    <v-col cols="4">
      <vx-date-picker v-model="value1" label="传入值是 Date" placeholder="选择日期" />
      <div class="text-caption">v-model: {{ value1 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-date-picker v-model="value2" label="传入值是 timestamp" placeholder="选择日期" />
      <div class="text-caption">v-model: {{ value2 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-date-picker v-model="value3" label="传入值是 String" placeholder="选择日期" />
      <div class="text-caption">v-model: {{ value3 }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value1 = ref(new Date())
const value2 = ref(1733390010891)
const value3 = ref('2024/11/1')
</script>
```

:::

### 组件类型

- type: `datepicker` `datetimepicker`

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker v-model="value" label="类型: datepicker" placeholder="选择日期" />
      <div>selected value: {{ value || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="类型: datetimepicker"
        format="YYYY-MM-DD HH:mm"
        type="datetimepicker"
        placeholder="选择日期和时间"
      />
      <div>selected value: {{ valueDateTime || 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref()
const valueDateTime = ref('')
</script>
```

:::

### 日期格式化 format

- datepicker 默认格式是 YYYY-MM-DD
- datetimepicker 默认格式是 YYYY-MM-DD HH:mm:ss
- **format 不仅影响展示也影响组件的可选项，以及绑定的 modelValue**

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="value"
        format="YYYY/MM/DD"
        label="datepicker (YYYY/MM/DD)"
        placeholder="选择日期"
      />
      <div>selected value: {{ value || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="datetimepicker (默认格式)"
        type="datetimepicker"
        placeholder="选择日期和时间"
      />
      <div>selected value: {{ valueDateTime || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueCustom"
        format="YYYY年MM月DD日"
        label="自定义格式 (YYYY年MM月DD日)"
        placeholder="选择日期"
      />
      <div>selected value: {{ valueCustom || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTimeCustom"
        label="datetimepicker (YYYY/MM/DD HH:mm)"
        format="YYYY/MM/DD HH:mm"
        type="datetimepicker"
        placeholder="选择日期和时间"
      />
      <div>selected value: {{ valueDateTimeCustom || 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref(Date.now())
const valueDateTime = ref(Date.now())
const valueCustom = ref('')
const valueDateTimeCustom = ref('')
</script>
```

:::

### 配合 datePickerProps 控制

可以通过 datePickerProps 控制日期选择器的行为，例如设置可选日期范围。

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="value"
        label="自定义可选日期范围"
        placeholder="选择日期"
        :date-picker-props="{ min: '2024/10/22', max: '2025/1/1' }"
      />
      <div class="text-caption">只能选择2024/10/22到2025/1/1之间的日期</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueWithTime"
        label="带时间的日期选择"
        type="datetimepicker"
        placeholder="选择日期和时间"
        :date-picker-props="{
          min: new Date(),
          max: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000),
          disableSecond: true
        }"
      />
      <div class="text-caption">只能选择未来7天内的日期和时间，不显示秒</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref(Date.now())
const valueWithTime = ref('')
</script>
```

:::

### 各种状态

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
const value = ref(Date.now())
const valueDateTime = ref('')
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker v-model="value" label="可清除 (clearable)" clearable tips="示例提示" />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="禁用 (disabled)"
        type="datetimepicker"
        placeholder="选择日期和时间"
        disabled
      />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        required
        label="必填 (required)"
        type="datetimepicker"
        placeholder="选择日期"
        :rules="[(value) => !!value || '必须选择日期']"
      />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="错误状态 (error-message)"
        type="datetimepicker"
        format="YYYY-MM-DD HH:mm:ss"
        placeholder="选择日期"
        error-messages="这是一条错误消息"
        clearable
      />
    </v-col>
  </v-row>
</template>
```

:::

### blur 事件示例

blur 事件在组件失焦或关闭下拉菜单时触发，可用于获取当前值或执行其他操作。

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="dateValue"
        label="blur 事件示例"
        placeholder="选择日期"
        clearable
        @blur="onDatePickerBlur"
      />
      <div class="text-caption">blur 事件触发次数: {{ blurCount }}</div>
      <div class="text-caption">最后一次 blur 事件值: {{ lastBlurValue }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const dateValue = ref(new Date())
const blurCount = ref(0)
const lastBlurValue = ref('')

function onDatePickerBlur(value) {
  blurCount.value++
  lastBlurValue.value = value
  console.log('Date picker blur event:', value)
}
</script>
```

:::
