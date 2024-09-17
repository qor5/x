# vx-label

## 基本用法

```html
  <vx-label>formLabel</vx-label>
```



:::demo

```vue
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

  <div>
     <vx-label toggleLabel>Title1 toggle to switch content
      <template #toggle-content>
         <div class="timeline-block">asdfsdf</div>
      </template>
     </vx-label>

     <vx-label toggleLabel>Title2 toggle to switch content
      <template #toggle-content>
         <div class="timeline-block">asdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdf</div>
      </template>
     </vx-label>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue"
const formatTooltip = ref(`{
      "defaultValue": {
              "value": "3426"
      },
      "valueType": "STRING"
}`)
</script>
```

<style scoped></style>
:::

