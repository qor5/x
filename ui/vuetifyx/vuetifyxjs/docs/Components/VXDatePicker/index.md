# vx-datepicker 时间选择器

包含了 datepicker 和 rangepicker

## API

| Name       | Introduction | Type     | Default Value |
| ---------- | ------------ | -------- | ------------- |
| type       |              | `String` | -             |
| modelValue | 绑定的值     | `String` | -             |

### Props

## 示例

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'
function onSelectDate(value) {
  console.log('outputvalue', new Date(value))
}

const value = ref('')
</script>

<template>
  <v-row>
    <v-col cols="3">
      <vx-date-picker
        v-model="value"
        label="选择日期"
        placeholder="Start at"
        :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }"
        @update:modelValue="onSelectDate"
      />
    </v-col>

    <v-col cols="3">
      <vx-date-picker
        v-model="value"
        label="选择时间日期"
        type="datetimepicker"
        placeholder="Start at"
        :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }"
        @update:modelValue="onSelectDate"
      />
    </v-col>

    <v-col cols="3">
      <vx-range-picker
        v-model="value"
        label="时间区间选择"
        :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }"
      />
    </v-col>
  </v-row>
</template>
```

:::

## legacy component

### vx-date-picker 日期选择器

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'

const value = ref('2023-10-01')
</script>

<template>
  <p>{{ value }}</p>
  <vx-date-picker v-model="value" :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }" />
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
