# vx-datepicker 时间选择器

包含了 datepicker 和 rangepicker

## API

| Name       | Introduction | Type     | Default Value |
| ---------- | ------------ | -------- | ------------- |
| type       | 组件类型     | `String` | `datepicker`  |
| modelValue | 绑定的值     | `String` | -             |

### Props

## 示例

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'
const value = ref()
const valueDateTime = ref('')
const rangeValue = ref([])
</script>

<template>
  <v-row>
    <v-col cols="6">
      <vx-date-picker
        v-model="value"
        label="选择日期(type: datepicker)"
        clearable
        placeholder="Start at"
        :date-picker-props="{ min: '2024-10-21', max: '2024-10-21' }"
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="选择时间日期(type: datetimepicker)"
        type="datetimepicker"
        placeholder="Start at"
        :date-picker-props="{ min: '2024-10-21', max: '2024-10-21', disableSecond: true }"
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="选择时间日期(type: datetimepicker)"
        type="datetimepicker"
        format="YYYY-MM-DD HH:mm:ss"
        placeholder="Start at"
        disabled
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-date-picker
        v-model="valueDateTime"
        label="选择时间日期(type: datetimepicker)"
        type="datetimepicker"
        format="YYYY-MM-DD HH:mm:ss"
        placeholder="Start at"
        :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }"
        error-messages="sfsfasf"
        clearable
      />
      <div>selected value: {{ valueDateTime ? new Date(valueDateTime) : 'unselected' }}</div>
    </v-col>

    <v-col cols="6">
      <vx-range-picker v-model="rangeValue" label="时间区间选择(日期)" />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        label="时间区间选择（日期时间）"
        format="YYYY-MM-DD HH:mm:ss"
        :date-picker-props="[{ min: '2024-06-15', max: '2024-06-20' }]"
      />
    </v-col>

    <v-col cols="6">
      <vx-range-picker
        v-model="rangeValue"
        type="datetimepicker"
        :placeholder="['start at', 'end at']"
        label="时间区间选择（日期时间）"
        needConfirm
        clearable
      />
    </v-col>
  </v-row>
</template>
```

:::

## legacy component

### vx-datepicker 日期选择器

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'

const value = ref('2023-10-01')
</script>

<template>
  <p>{{ value }}</p>
  <vx-datepicker v-model="value" :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }" />
</template>

<style scoped></style>
```

:::

### vx-date-timepicker 日期选择器

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import Datetimepicker from '@/lib/Datetimepicker.vue'

const value = ref('2023-10-01 22:33')
</script>

<template>
  <p>{{ value }}</p>
  <datetimepicker v-model="value" :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }" />
</template>

<style scoped></style>
```

:::
