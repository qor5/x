# 时间区间选择器

`vx-range-picker` 组件用于选择时间区间，提供灵活的日期范围选择功能。

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

> 该组件继承了 [vx-field的props](/Components/VXField/#props)，所以是通用的，以下仅罗列 vx-range-picker 特有的Props

| Name      | Introduction                         | Type      | Default Value |
| --------- | ------------------------------------ | --------- | ------------- |
| type      | 见[组件类型](./#组件类型)            | `String`  | `datepicker`  |
| format    | 见[日期格式化](./#日期格式化-format) | `String`  | 'YYYY-MM-DD'  |
| clearable | 显示可清除样式                       | `Boolean` | `false`       |

#### vx-range-picker props

| Name            | Introduction                                                                          | Type                           | Default Value |
| --------------- | ------------------------------------------------------------------------------------- | ------------------------------ | ------------- |
| placeholder     | 占位文本                                                                              | `string[]`                     | `['', '']`    |
| needConfirm     | 选中的值需要点击确认才生效                                                            | `boolean`                      | `false`       |
| datePickerProps | vuetify [原生参数](https://vuetifyjs.com/en/api/v-date-picker/)，用于控制时间选择组件 | `datePickerProps[]`            | `[]`          |
| modelValue      | 绑定的值，见[值类型](./#支持的值类型)                                                 | `string[]` `number[]` `Date[]` | `['', '']`    |

### Events

| Name          | Payload                                   | Introduction                                                                                |
| ------------- | ----------------------------------------- | ------------------------------------------------------------------------------------------- |
| click:confirm | `{ value: Ref<number[]>, next: Promise }` | 配置了 needconfirm 可以结合 click:confirm [实现对值的校验](./#配合-datepickerprops-控制)    |
| blur          | `string`                                  | 某些情况如果期望失焦或者关闭了下拉后获取值时用这个事件，查看[blur 事件示例](#blur-事件示例) |

### 值类型

- 组件拥有很强的传入值适应性，可以传入各种类型的值并格式化字符串、时间戳、日期类型都可，最终都会被格式化成默认格式或者传入的format格式
- vx-range-picker 可以切换值选中模式，当 `needConfirm` 为 `true` 时，需要点击确认按钮值才生效

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-range-picker
        v-model="value1"
        label="范围选择器 (字符串)"
        :placeholder="['开始时间', '结束时间']"
      />
      <div class="text-caption">v-model: {{ value1 }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="value2"
        label="范围选择器 (时间戳)"
        :placeholder="['开始时间', '结束时间']"
      />
      <div class="text-caption">v-model: {{ value2 }}</div>
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <vx-range-picker v-model="value3" label="无需确认" :placeholder="['开始时间', '结束时间']" />
      <div class="text-caption">v-model: {{ value3 }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="value4"
        label="需要确认 (needConfirm)"
        :placeholder="['开始时间', '结束时间']"
        needConfirm
      />
      <div class="text-caption">v-model: {{ value4 }}</div>
    </v-col>
  </v-row>

  <v-row class="ml-1 mb-2 ">默认选择的时间部分以选择日的 00:00:00 为基准，格式是时间戳</v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value1 = ref(['2024/11/1', '2024/12/1'])
const value2 = ref([1733390010891, 1735981610891])
const value3 = ref(['2024/10/1', '2024/10/15'])
const value4 = ref(['', ''])
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
      <vx-range-picker
        v-model="rangeValueDate"
        label="type: datepicker"
        :placeholder="['开始时间', '结束时间']"
      />
      <div>selected value: {{ rangeValueDate ? rangeValueDate : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDateTime"
        label="type: datetimepicker"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
      />
      <div>selected value: {{ rangeValueDateTime ? rangeValueDateTime : 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const rangeValueDate = ref(['', ''])
const rangeValueDateTime = ref(['', ''])
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
      <vx-range-picker
        v-model="rangeValueDate"
        format="YYYY/MM/DD"
        label="type: datepicker (YYYY/MM/DD)"
        :placeholder="['开始时间', '结束时间']"
      />
      <div>selected value: {{ rangeValueDate ? rangeValueDate : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDateTime"
        label="type: datetimepicker (default format)"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
      />
      <div>selected value: {{ rangeValueDateTime ? rangeValueDateTime : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDateTimeCustom"
        label="datetimepicker (YYYY/MM/DD HH:mm)"
        format="YYYY/MM/DD HH:mm"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
        needConfirm
      />
      <div>
        selected value: {{ rangeValueDateTimeCustom ? rangeValueDateTimeCustom : 'unselected' }}
      </div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const rangeValueDate = ref(['', ''])
const rangeValueDateTime = ref(['', ''])
const rangeValueCustom = ref(['', ''])
const rangeValueDateTimeCustom = ref(['', ''])
</script>
```

:::

### 配合 datePickerProps 控制

可以通过 datePickerProps 对两个日期选择器分别进行控制，例如设置可选日期范围。

:::demo

```vue
<template>
  <v-row>
    <v-col cols="12">
      <vx-range-picker
        v-model="valueRangePicker"
        label="自定义可选日期范围"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
        :date-picker-props="[
          { min: new Date(), max: new Date(Date.now() + 30 * 24 * 60 * 60 * 1000) },
          { min: new Date(), max: new Date(Date.now() + 60 * 24 * 60 * 60 * 1000) }
        ]"
        needConfirm
        clearable
        @click:confirm="onClickConfirm"
      />
      <div class="text-caption">当选择了开始值大于结束值，不予保存</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const valueRangePicker = ref(['', ''])

function onClickConfirm({ value, next }) {
  console.log(value[0], value[1])
  if (value[0] > value[1]) {
    alert('开始日期必须早于结束日期！')
  } else if (!value[0] || !value[1]) {
    alert('请选择完整的日期范围')
  } else {
    next()
  }
}
</script>
```

:::

### 各种状态

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
const rangeValue = ref([Date.now(), Date.now() + 86400000])
const rangeValue1 = ref(['', ''])
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        label="可清除 (clearable)"
        :placeholder="['开始时间', '结束时间']"
        tips="可清除的范围选择器"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
        label="禁用 (disabled)"
        disabled
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue1"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
        label="必填 (required)"
        required
        clearable
        :rules="[(value) => value.every((item) => !!item) || '必须选择日期范围']"
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        :placeholder="['开始时间', '结束时间']"
        label="错误状态 (error)"
        error-messages="这是一条错误消息"
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
    <v-col cols="12">
      <vx-range-picker
        v-model="rangeValue"
        label="blur 事件示例"
        :placeholder="['开始日期', '结束日期']"
        clearable
        @blur="onRangePickerBlur"
      />
      <div class="text-caption">blur 事件触发次数: {{ rangeBlurCount }}</div>
      <div class="text-caption">最后一次 blur 事件值: {{ lastRangeBlurValue }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const rangeValue = ref(['2024/11/1', '2024/12/1'])
const rangeBlurCount = ref(0)
const lastRangeBlurValue = ref('')

function onRangePickerBlur(value) {
  rangeBlurCount.value++
  lastRangeBlurValue.value = value
  console.log('Range picker blur event:', value)
}
</script>
```

:::
