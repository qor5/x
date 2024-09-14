# vx-label

## 基本用法

```html
  <vx-label>formLabel</vx-label>
```



:::demo

```vue
<script setup lang="ts">
import { ref } from "vue"
const formatTooltip = ref(`{
      "defaultValue": {
              "value": "3426"
      },
      "valueType": "STRING"
}`)
</script>
<template>
  <div>
      <vx-label>simple label</vx-label>
  </div>

  <div>
     <vx-label icon="mdi-help-circle-outline" icon-size="default" tooltip="I'm a tip" tooltip-icon-color="primary">label with tooltip</vx-label>
  </div>

  <div>
     <vx-label icon="mdi-help-circle-outline" icon-size="default"><span style="font-size:20px; color: red;">custom label</span></vx-label>
  </div>

  <div>
     <vx-label icon="mdi-help-circle-outline" :tooltip="formatTooltip" >label with format tooltip</vx-label>
  </div>
</template>
```

<style scoped></style>
:::

