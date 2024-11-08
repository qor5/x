# vx-pagination

带有样式修复的 v-pagination 组件。

## API

### Props

| Name          | Introduction                                  | Type     | Default Value |
| ------------- | --------------------------------------------- | -------- | ------------- |
| length        | The number of pages.                          | `number` | 总页数        |
| total-visible | Specify the total visible pagination numbers. | `number` | 5             |
| model-value   | The v-model value of the component            | `number` | -             |

> 除此之外所有的 [v-pagination](https://vuetifyjs.com/en/api/v-pagination/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Events

| Name              | Payload  | Introduction       |
| ----------------- | -------- | ------------------ |
| update:modelValue | `number` | model 值改变时触发 |

## 示例

- 基本用法如果因为数值过大而导致的数字溢出，应当考虑增加宽度或者减小 `total-visible` 的值, 或者调整size 到 `small`

:::demo

```vue
<template>
  <v-row>
    <v-col cols="12">
      <div class="text-primary">size: small</div>
      <vx-pagination size="small" :length="999999" v-model="currentPage1" />
    </v-col>
    <v-col cols="12">
      <div class="text-primary">size: default</div>
      <vx-pagination :length="999999" v-model="currentPage2" />
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'
const currentPage1 = ref(1)
const currentPage2 = ref(2)
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::
