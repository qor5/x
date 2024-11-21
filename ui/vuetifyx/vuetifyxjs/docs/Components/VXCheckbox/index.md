# vx-checkbox 复选框

带更多扩展API的复选框

## API

### Props

| Name       | Introduction                   | Type      | Default Value |
| ---------- | ------------------------------ | --------- | ------------- |
| title      | 顶部标题（通常作为表单项使用） | `String`  | -             |
| label      | 复选框描述                     | `String`  | -             |
| readonly   | 是否只读（只读时禁用涟漪效果） | `Boolean` | `false`       |
| disabled   | 是否禁用                       | `Boolean` | `false`       |
| trueLabel  | 复选框描述（勾选时）           | `String`  | -             |
| falseLabel | 复选框描述（不勾选时）         | `String`  | -             |
| trueColor  | 复选框颜色（勾选时）           | `String`  | -             |
| falseColor | 复选框颜色（不勾选时）         | `String`  | -             |

> 除此之外所有的 [v-checkbox](https://vuetifyjs.com/en/api/v-checkbox/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

## 示例

:::demo

```vue
<template>
  <v-row>
    <v-col cols="4">
      <div class="mb-2">1.使用 v-model</div>
      <vx-checkbox v-model="checked" label="Checkbox" />
      <div class="text-caption">bind value: {{ checked }}</div>
    </v-col>

    <v-col cols="4">
      <div class="mb-2">2. readonly</div>

      <vx-checkbox :model-value="false" label="Checkbox" readonly />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">3. disabled</div>

      <vx-checkbox :model-value="true" label="Checkbox" disabled />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">4. true label</div>

      <vx-checkbox :model-value="true" label="Checkbox" true-label="Yes" false-label="NO" />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">5. false label</div>

      <vx-checkbox :model-value="false" label="Checkbox" true-label="Yes" false-label="NO" />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">5. true icon</div>

      <vx-checkbox
        :model-value="true"
        label="Checkbox"
        true-label="Yes"
        false-label="NO"
        true-icon="mdi-circle-outline"
      />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">6.false icon</div>

      <vx-checkbox
        :model-value="true"
        label="Checkbox"
        true-label="Yes"
        false-label="NO"
        false-icon="mdi-window-close"
      />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">7. readonly + true icon color</div>

      <vx-checkbox
        :model-value="true"
        label="Checkbox"
        readonly
        true-label="Yes"
        false-label="NO"
        true-icon="mdi-circle-outline"
        true-icon-color="primary"
      />
    </v-col>

    <v-col cols="4">
      <div class="mb-2">8. readonly + false icon color</div>

      <vx-checkbox
        :model-value="false"
        label="Checkbox"
        readonly
        true-label="Yes"
        false-label="NO"
        false-icon-color="error"
        false-icon="mdi-window-close"
      />
    </v-col>
  </v-row>

  <v-row>
    <v-col>
      <vx-checkbox :model-value="false" title="带标题 title" label="Checkbox" />
    </v-col>

    <v-col style="background-color: #f5f5f5;">
      <vx-checkbox
        :model-value="false"
        icon="mdi-circle-outline"
        title="复选框背景色为白色"
        label="Checkbox"
      />
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const checked = ref(false)
const readonlyTrue = ref(false)
</script>
```
