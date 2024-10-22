# 时间选择器

包含了时间选择 `vx-date-picker` 和时间区间选择 `vx-range-picker`

### Props

#### 公共props

> 该组件继承了 [vx-field的props](/x/Components/VXField/#props)，所以是通用的，以下仅罗列 vx-date-picker 和 vx-range-picker 特有的Props

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

#### vx-range-picker props

> 和 vx-date-picker 最大的不同在于，其大多数接收的都是数组

| Name            | Introduction                                                                          | Type                           | Default Value |
| --------------- | ------------------------------------------------------------------------------------- | ------------------------------ | ------------- |
| placeholder     | 占位文本                                                                              | `string[]`                     | -             |
| needConfirm     | 选中的值需要点击确认才生效                                                            | `boolean`                      | `false`       |
| datePickerProps | vuetify [原生参数](https://vuetifyjs.com/en/api/v-date-picker/)，用于控制时间选择组件 | `datePickerProps[]`            | -             |
| modelValue      | 绑定的值，见[值类型](./#支持的值类型)                                                 | `string[]` `number[]` `Date[]` | -             |

#### vx-range-picker events

| Name          | Payload                                   | Introduction                                                                              |
| ------------- | ----------------------------------------- | ----------------------------------------------------------------------------------------- |
| click:confirm | `{ value: Ref<number[]>, next: Promise }` | 配置了 needconfirm 可以结合 click:confirm [实现对值的校验](.md#配合-datepickerprops-控制) |

### 值类型

- 组件拥有很强的传入值适应性，可以传入各种类型的值并格式化字符串、时间戳、日期类型都可
- vx-range-picker 可以切换值选中模式，当 `needConfirm` 为 `true` 时，需要点击确认按钮值才生效

:::demo

```vue
<template>
  <v-row class="pl-3 text-primary mt-2"><b>vx-date-picker </b></v-row>
  <v-row>
    <v-col cols="4">
      <vx-date-picker v-model="value1" label="传入值是 Date" placeholder="Start at" />
      <div class="text-caption">v-model: {{ value1 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-date-picker v-model="value2" label="传入值是 timestamp" placeholder="Start at" />
      <div class="text-caption">v-model: {{ value2 }}</div>
    </v-col>

    <v-col cols="4">
      <vx-date-picker v-model="value3" label="传入值是 String" placeholder="Start at" />
      <div class="text-caption">v-model: {{ value3 }}</div>
    </v-col>
  </v-row>

  <v-row class="pl-3 text-primary mt-5"><b>vx-range-picker</b></v-row>
  <v-row>
    <v-col cols="6">
      <vx-range-picker v-model="value4" label="range-picker" placeholder="Start at" />
      <div class="text-caption">v-model: {{ value4 }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="value5"
        label="range-picker (needConfirm)"
        :placeholder="['Start at', 'End at']"
        needConfirm
      />
      <div class="text-caption">v-model: {{ value5 }}</div>
    </v-col>
  </v-row>

  <v-row class="ml-1 mb-2 ">默认选择的时间部分以选择日的 00:00:00 为基准，格式是时间戳</v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'
const value1 = ref(new Date())
const value2 = ref(Date.now())
const value3 = ref('2024/11/1')
const value4 = ref(['2024/11/1', '2024/12/1 12:21'])
const value5 = ref(['', ''])
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
      <vx-date-picker v-model="value" label="vx-date-picker (datepicker)" placeholder="Start at" />
      <div>selected value: {{ value ? new Date(value) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="vx-date-picker (datetimepicker)"
        type="datetimepicker"
        placeholder="Choose Datetime"
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDate"
        label="vx-range-picker (datepicker)"
        :placeholder="['Start at', 'End at']"
      />
      <div>selected value: {{ rangeValueDate ? rangeValueDate : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDateTime"
        label="vx-range-picker (datetimepicker)"
        type="datetimepicker"
        :placeholder="['Start at', 'End at']"
      />
      <div>selected value: {{ rangeValueDateTime ? rangeValueDateTime : 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'
const value = ref()
const valueDateTime = ref('')
const rangeValueDate = ref(['', ''])
const rangeValueDateTime = ref(['', ''])
</script>
```

:::

### 日期格式化 format

- datepicker 默认格式是 YYYY-MM-DD
- datetimepicker 默认格式是 YYYY-MM-DD HH:mm
- format 不仅影响展示，也影响选择以后格式化的值，原则是格式化未覆盖到的值会被重置成0（time部分）

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="value"
        format="YYYY/MM-DD"
        label="vx-date-picker (datepicker)"
        placeholder="Start at"
      />
      <div>selected value: {{ value ? new Date(value) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="vx-date-picker (datetimepicker)"
        type="datetimepicker"
        format="YYYY/MM/DD/HH/mm/ss"
        placeholder="Choose Datetime"
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDate"
        label="vx-range-picker (datepicker)"
        :placeholder="['Start at', 'End at']"
      />
      <div>selected value: {{ rangeValueDate ? rangeValueDate : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValueDateTime"
        label="vx-range-picker (datetimepicker)"
        format="YYYY/MM/DD HH:mm"
        type="datetimepicker"
        :placeholder="['Start at', 'End at']"
        needConfirm
      />
      <div>selected value: {{ rangeValueDateTime ? rangeValueDateTime : 'unselected' }}</div>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'
const value = ref(Date.now())
const valueDateTime = ref(Date.now())
const rangeValueDate = ref(['', ''])
const rangeValueDateTime = ref(['', ''])
</script>
```

:::

### 配合 datePickerProps 控制

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="value"
        label="vx-date-picker custom select range"
        placeholder="Start at"
        :date-picker-props="{ min: '2024/10/22', max: '2025/1/1' }"
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="valueRangePicker"
        label="vx-range-picker custom select range"
        type="datetimepicker"
        :placeholder="['Start at', 'End at']"
        :date-picker-props="[
          { min: new Date(), max: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000) },
          { min: new Date(), max: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000) }
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
import Datepicker from '@/lib/Datepicker.vue'
const value = ref(Date.now())
const valueRangePicker = ref(['', ''])

function onClickConfirm({ value, next }) {
  console.log(value[0], value[1])
  if (value[0] > value[1]) {
    alert('start date should ahead of the end date !')
  } else if (!value[0] || !value[1]) {
    alert('please select a range date')
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
import Datepicker from '@/lib/Datepicker.vue'
const value = ref(Date.now())
const valueDateTime = ref('')
const rangeValue = ref([Date.now(), Date.now()])
const rangeValue1 = ref([])
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker v-model="value" label="vx-date-picker(clearable)" clearable tips="example" />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="vx-date-picker(disabled)"
        type="datetimepicker"
        placeholder="Choose Datetime"
        disabled
      />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        required
        label="vx-date-picker(required)"
        type="datetimepicker"
        placeholder="Select a date"
        :rules="[(value) => !!value || 'You must select a date']"
      />
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="vx-date-picker(error-message)"
        type="datetimepicker"
        format="YYYY-MM-DD HH:mm:ss"
        placeholder="Start at"
        error-messages="This is an error message"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        label="vx-range-picker （clearable）"
        :placeholder="['start at', 'end at']"
        tips="clearable rangepicker"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        :placeholder="['start at', 'end at']"
        label="vx-range-picker (disabled)"
        disabled
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue1"
        type="datetimepicker"
        :placeholder="['start at', 'end at']"
        label="vx-range-picker (required)"
        required
        clearable
        :rules="[(value) => value.every((item) => !!item) || 'You must select a date']"
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        :placeholder="['start at', 'end at']"
        label="vx-range-picker (disabled)"
        error-messages="This is a error message"
      />
    </v-col>
  </v-row>
</template>
```

:::
