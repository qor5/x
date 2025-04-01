# 时间选择器

`vx-time-picker` 组件用于选择时间，提供灵活的时间选择功能。

**目录**：

- [Props](#props)
- [Events](#events)
- [值类型](#值类型)
- [日期格式化 format](#日期格式化-format)
- [配合 datePickerProps 控制](#配合-datepickerprops-控制)
- [各种状态](#各种状态)
- [blur 事件示例](#blur-事件示例)

### Props

#### 公共props

> 该组件继承了 [vx-field的props](/Components/VXField/#props)，所以是通用的，以下仅罗列 vx-time-picker 特有的Props

| Name            | Introduction                                                                          | Type                     | Default Value |
| --------------- | ------------------------------------------------------------------------------------- | ------------------------ | ------------- |
| format          | 见[日期格式化](./#日期格式化-format)                                                  | `String`                 | 'HH:mm:ss'    |
| clearable       | 显示可清除样式                                                                        | `Boolean`                | `false`       |
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
      <vx-time-picker v-model="value1" label="传入值是 Date" placeholder="选择时间" />
      <div class="text-caption">v-model: {{ value1 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-time-picker v-model="value2" label="传入值是 timestamp" placeholder="选择时间" />
      <div class="text-caption">v-model: {{ value2 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-time-picker v-model="value3" label="传入值是 String" placeholder="选择时间" />
      <div class="text-caption">v-model: {{ value3 }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value1 = ref(new Date())
const value2 = ref(1733390010891)
const value3 = ref('14:30:00')
</script>
```

:::

### 日期格式化 format

- 默认格式是 HH:mm:ss
- **format 不仅影响展示也影响组件的可选项，以及绑定的 modelValue**

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker
        v-model="value1"
        format="HH:mm"
        label="不显示秒 (HH:mm)"
        placeholder="选择时间"
      />
      <div>selected value: {{ value1 || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker v-model="value2" label="默认格式 (HH:mm:ss)" placeholder="选择时间" />
      <div>selected value: {{ value2 || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="value3"
        format="HH时mm分ss秒"
        label="自定义格式 (HH时mm分ss秒)"
        placeholder="选择时间"
      />
      <div>selected value: {{ value3 || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="value4"
        format="h:mm a"
        label="12小时制 (h:mm a)"
        placeholder="选择时间"
      />
      <div>selected value: {{ value4 || 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value1 = ref(Date.now())
const value2 = ref(Date.now())
const value3 = ref('')
const value4 = ref('')
</script>
```

:::

### 配合 datePickerProps 控制

可以通过 datePickerProps 控制时间选择器的行为，例如设置最小值、最大值等。

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker
        v-model="value"
        label="限制可选时间范围"
        placeholder="选择时间"
        :date-picker-props="{
          min: '2024-10-22 09:00:00',
          max: '2024-10-22 18:00:00',
          disableSecond: true
        }"
      />
      <div class="text-caption">只能选择9:00到18:00之间的时间，不显示秒</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="value2"
        label="禁用小时/分钟/秒选择"
        placeholder="选择时间"
        :date-picker-props="{
          disableHour: false,
          disableMinute: false,
          disableSecond: true
        }"
      />
      <div class="text-caption">可以选择小时和分钟，但禁用了秒选择</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref('')
const value2 = ref('')
</script>
```

:::

### 各种状态

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
const value1 = ref(Date.now())
const value2 = ref('')
const value3 = ref('')
const value4 = ref('')
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker v-model="value1" label="可清除 (clearable)" clearable tips="示例提示" />
    </v-col>

    <v-col cols="6">
      <vx-time-picker v-model="value2" label="禁用 (disabled)" placeholder="选择时间" disabled />
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="value3"
        required
        label="必填 (required)"
        placeholder="选择时间"
        :rules="[(value) => !!value || '必须选择时间']"
      />
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="value4"
        label="错误状态 (error-message)"
        placeholder="选择时间"
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
      <vx-time-picker
        v-model="timeValue"
        label="blur 事件示例"
        placeholder="选择时间"
        clearable
        @blur="onTimePickerBlur"
      />
      <div class="text-caption">blur 事件触发次数: {{ blurCount }}</div>
      <div class="text-caption">最后一次 blur 事件值: {{ lastBlurValue }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const timeValue = ref(new Date())
const blurCount = ref(0)
const lastBlurValue = ref('')

function onTimePickerBlur(value) {
  blurCount.value++
  lastBlurValue.value = value
  console.log('Time picker blur event:', value)
}
</script>
```

:::
