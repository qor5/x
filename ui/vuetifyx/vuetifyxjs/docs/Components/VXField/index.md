# vx-field 输入框

## 基本用法

```html
<vx-field model-value="Hello World" label="field1" />

```

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-field :model-value="dataNoBinding" label="field1" />
    </v-col>
    <v-col cols="6" class="pt-12"><b> no binding:</b> {{ dataNoBinding }} </v-col>

    <v-col cols="6" sm="6">
      <vx-field v-model="dataWithBinding" label="field2" />
    </v-col>
    <v-col cols="6" class="pt-12"><b> with binding:</b> {{ dataWithBinding }} </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const dataNoBinding = ref('hello world')
const dataWithBinding = ref('hello world')
</script>

<style scoped lang="css">
* {
  white-space: break-word;
}
</style>
```

:::
