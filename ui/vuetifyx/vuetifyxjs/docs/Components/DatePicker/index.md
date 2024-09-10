# 时间选择器

## vx-datepicker 日期选择器

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

## vx-datetimepicker 日期选择器

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

## vx-textdatepicker

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const value = ref('2023-10-01')
</script>

<template>
  <p>{{ value }}</p>
  <vx-textdatepicker v-model="value" />
</template>
```
:::