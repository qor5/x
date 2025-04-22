# vx-tabs

带有预设样式的 v-tabs 组件。

## API

### Props

| Name            | Introduction           | Type                  | Default Value |
| --------------- | ---------------------- | --------------------- | ------------- |
| underlineBorder | tabs底部border样式控制 | `"contain"` ` "full"` | -             |
| pill            | 胶囊样式tabs           | `Boolean`             | `false`       |

> 除此之外所有的 [v-tabs](https://vuetifyjs.com/en/api/v-tabs/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Events

| Name              | Payload  | Introduction       |
| ----------------- | -------- | ------------------ |
| update:modelValue | `number` | model 值改变时触发 |

## 示例

### underlineBorder

:::demo

```vue
<template>
  <v-row>
    <v-col cols="3" style="align-content:center">no underline-border</v-col>
    <v-col cols="9">
      <vx-tabs v-model="tab1">
        <v-tab :value="1">Landscape</v-tab>
        <v-tab :value="2">City</v-tab>
        <v-tab :value="3">Abstract</v-tab>
      </vx-tabs>
      <div class="text-caption mt-2">current tab value:{{ tab1 }}</div>
    </v-col>

    <v-col cols="3" style="align-content:center">underline-border: contain </v-col>
    <v-col cols="9">
      <vx-tabs underline-border="contain" v-model="tab2">
        <v-tab :value="1">Landscape</v-tab>
        <v-tab :value="2">City</v-tab>
        <v-tab :value="3">Abstract</v-tab>
      </vx-tabs>
      <div class="text-caption mt-2">current tab value:{{ tab2 }}</div>
    </v-col>

    <v-col cols="3" style="align-content:center">underline-border: full</v-col>
    <v-col cols="9">
      <vx-tabs underline-border="full" v-model="tab3">
        <v-tab :value="1">Landscape</v-tab>
        <v-tab :value="2">City</v-tab>
        <v-tab :value="3">Abstract</v-tab>
      </vx-tabs>
      <div class="text-caption mt-2">current tab value:{{ tab3 }}</div>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'
const tab1 = ref(1)
const tab2 = ref(2)
const tab3 = ref(3)
</script>

<style scoped lang="css"></style>
```

:::

### 胶囊样式

:::demo

```vue
<template>
  <v-row>
    <v-col cols="12">
      <vx-tabs pill v-model="pillTab">
        <v-tab :value="1">Once</v-tab>
        <v-tab :value="2">Daily</v-tab>
        <v-tab :value="3">Weekly</v-tab>
        <v-tab :value="4">Monthly</v-tab>
      </vx-tabs>
      <div class="text-caption mt-2">当前选中: {{ pillTab }}</div>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'
const pillTab = ref(1)
</script>

<style scoped lang="css"></style>
```

:::

### 与 v-tabs-window 组合使用

:::demo

```vue
<template>
  <v-row>
    <div class="text-caption mt-2">current tab value:{{ tab3 }}</div>
    <v-col cols="12">
      <vx-tabs underline-border="full" v-model="tab3">
        <v-tab value="tab-1">Landscape</v-tab>
        <v-tab value="tab-2">City</v-tab>
        <v-tab value="tab-3">Abstract</v-tab>
      </vx-tabs>

      <v-tabs-window v-model="tab3">
        <v-tabs-window-item v-for="i in 3" :key="i" :value="'tab-' + i">
          <v-card elevation="0">
            <v-card-text>
              <div
                class="border border-dashed text-primary font-weight-bold border-primary text-center border-opacity-100 pa-4"
              >
                {{ tab3 }}
              </div></v-card-text
            >
          </v-card>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const tab3 = ref(3)
</script>

<style scoped lang="css"></style>
```

:::
