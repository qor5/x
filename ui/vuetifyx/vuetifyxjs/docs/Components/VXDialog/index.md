# vx-dialog 弹窗

##

:::demo

```vue
<template>
  <vx-dialog title="Confirm" text="Are you sure to close the dialog?" v-model="dialogVisible" />
  <v-btn color="primary" @click="dialogVisible = true">Open Dialog</v-btn>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const dialogVisible = ref(false)
</script>
```

:::

### 4. 为组件撰写必要说明和参数

目前先随意，后期会有规范
