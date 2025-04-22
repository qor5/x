# 时间选择器

`vx-time-picker` 组件用于选择时间。

**目录**：

- [Props](#props)
- [Events](#events)
- [值类型](#值类型)
- [时间格式化 format](#时间格式化-format)
- [配合 timePickerProps 控制](#配合-timepickerprops-控制)
- [各种状态](#各种状态)
- [需要确认的时间选择器](#需要确认的时间选择器)
- [blur 事件示例](#blur-事件示例)

### Props

#### 公共props

> 该组件继承了 [vx-field的props](/Components/VXField/#props)，所以是通用的，以下仅罗列 vx-time-picker 特有的Props

| Name      | Introduction                         | Type      | Default Value |
| --------- | ------------------------------------ | --------- | ------------- |
| format    | 见[时间格式化](./#时间格式化-format) | `String`  | 'HH:mm:ss'    |
| clearable | 显示可清除样式                       | `Boolean` | `false`       |

#### vx-time-picker props

| Name            | Introduction                    | Type              | Default Value |
| --------------- | ------------------------------- | ----------------- | ------------- |
| placeholder     | 占位文本                        | `String`          | -             |
| timePickerProps | 时间选择组件的配置参数          | `Object`          | -             |
| modelValue      | 绑定的值，见[值类型](./#值类型) | `String` `Number` | -             |
| hideAppendInner | 隐藏内部的附加图标              | `Boolean`         | `false`       |
| needConfirm     | 是否需要确认按钮                | `Boolean`         | `false`       |

### Events

| Name          | Payload  | Introduction                                                                                |
| ------------- | -------- | ------------------------------------------------------------------------------------------- |
| blur          | `string` | 某些情况如果期望失焦或者关闭了下拉后获取值时用这个事件，查看[blur 事件示例](#blur-事件示例) |
| click:confirm | `object` | 当点击确认按钮时触发                                                                        |

### 值类型

- 组件可以接受字符串或数字类型的时间值，最终会根据设定的format格式化输出

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker v-model="value1" label="传入值是 String" placeholder="选择时间" />
      <div class="text-caption">v-model: {{ value1 }}</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker v-model="value2" label="默认值" placeholder="选择时间" />
      <div class="text-caption">v-model: {{ value2 }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value1 = ref('14:30:00')
const value2 = ref('')
</script>
```

:::

### 时间格式化 format

- 默认格式是 HH:mm:ss
- 可以自定义格式，如 HH:mm
- **format 不仅影响展示也影响组件的可选项**，例如使用 HH:mm 格式时，秒选择器将不会显示

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker v-model="value" label="默认格式 (HH:mm:ss)" placeholder="选择时间" />
      <div>selected value: {{ value || 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="valueCustom"
        format="HH:mm"
        label="自定义格式 (HH:mm)"
        placeholder="选择时间"
      />
      <div>selected value: {{ valueCustom || 'unselected' }}</div>
      <div class="text-caption">使用 HH:mm 格式，秒选择器不会显示</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref('12:30:45')
const valueCustom = ref('14:30')
</script>
```

:::

### 配合 timePickerProps 控制

可以通过 timePickerProps 控制时间选择器的行为。

**注意**：禁用（disableXXX）和隐藏（hideXXX）是两种不同的状态：

- 禁用（disableXXX）：时间单位依然可见，但不可选择（显示为灰色）
- 隐藏（hideXXX）：时间单位完全不显示

:::demo

```vue
<template>
  <v-row>
    <v-col cols="4">
      <vx-time-picker
        v-model="value1"
        label="禁用秒选择"
        placeholder="选择时间"
        :time-picker-props="{ disableSecond: true }"
      />
      <div class="text-caption">禁用秒：秒选择器显示但不可选</div>
    </v-col>

    <v-col cols="4">
      <vx-time-picker v-model="value2" label="隐藏秒选择" format="HH:mm" placeholder="选择时间" />
      <div class="text-caption">隐藏秒：秒选择器完全不显示</div>
    </v-col>

    <v-col cols="4">
      <vx-time-picker
        v-model="value3"
        label="自定义时间范围"
        placeholder="选择时间"
        :time-picker-props="{ hourRange: [9, 18] }"
      />
      <div class="text-caption">限制可选小时范围：9-18</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value1 = ref('12:30:00')
const value2 = ref('12:30')
const value3 = ref('12:30:00')
</script>
```

:::

### 各种状态

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
const value = ref('12:30:00')
const valueEmpty = ref('')
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker v-model="value" label="可清除 (clearable)" clearable tips="示例提示" />
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="valueEmpty"
        label="禁用 (disabled)"
        placeholder="选择时间"
        disabled
      />
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="valueEmpty"
        required
        label="必填 (required)"
        placeholder="选择时间"
        :rules="[(value) => !!value || '必须选择时间']"
      />
    </v-col>

    <v-col cols="6">
      <vx-time-picker
        v-model="valueEmpty"
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

### 需要确认的时间选择器

使用 needConfirm 属性可以添加确认按钮，只有点击确认后才会更新值。

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-time-picker
        v-model="value"
        label="需要确认的时间选择器"
        placeholder="选择时间"
        need-confirm
        @click:confirm="onConfirm"
      />
      <div class="text-caption">选择时间后需要点击确认按钮才会更新值</div>
      <div class="text-caption">确认的值: {{ confirmedValue }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const value = ref('12:30:00')
const confirmedValue = ref('')

function onConfirm(event) {
  confirmedValue.value = event.value[0]
  console.log('Confirmed time:', event)
}
</script>
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

const timeValue = ref('14:30:00')
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
