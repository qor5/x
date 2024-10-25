# vx-label 标签

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
    <vx-label
      icon="mdi-help-circle-outline"
      icon-size="default"
      tooltip="I'm a tip"
      tooltip-icon-color="primary"
      >label with tooltip</vx-label
    >
  </div>

  <div>
    <vx-label icon="mdi-help-circle-outline" icon-size="default"
      ><span style="font-size:20px; color: red;">custom label</span></vx-label
    >
  </div>

  <div>
    <vx-label icon="mdi-help-circle-outline" :tooltip="formatTooltip"
      >label with format tooltip</vx-label
    >
  </div>

  <div>
    <vx-label toggleLabel class="sortable-label">
      <template #prepend>
        <div class="section-sortable-area">
          <v-btn-group
            border
            density="comfortable"
            class="d-flex flex-column h-auto section-sortable-group"
            elevation="5"
          >
            <v-btn
              icon="mdi-arrow-up"
              size="x-small"
              class="rounded-0 border-b"
              style="width:22px;height:30px;font-size:10px"
              :disabled="current == 0"
            />
            <v-btn
              icon="mdi-arrow-down"
              size="x-small"
              class="rounded-0"
              style="width:22px;height:30px;font-size:10px"
              :disabled="current == totalCount"
            />
          </v-btn-group>
        </div>
      </template>
      Title1 toggle to switch content
      <template #toggle-content>
        <div class="timeline-block">asdfsdf</div>
      </template>
    </vx-label>

    <vx-label toggleLabel class="sortable-label">
      <template #prepend>
        <div class="section-sortable-area">
          <v-btn-group
            border
            density="comfortable"
            class="d-flex flex-column h-auto section-sortable-group"
            elevation="5"
          >
            <v-btn
              icon="mdi-arrow-up"
              size="x-small"
              class="rounded-0 border-b"
              style="width:22px;height:30px;font-size:10px"
              :disabled="current == 0"
            />
            <v-btn
              icon="mdi-arrow-down"
              size="x-small"
              class="rounded-0"
              style="width:22px;height:30px;font-size:10px"
              :disabled="current == totalCount"
            />
          </v-btn-group>
        </div>
      </template>

      Title2 toggle to switch content
      <template #toggle-content>
        <div class="timeline-block">
          asdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdfasdfsdf
        </div>
      </template>
    </vx-label>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
const formatTooltip = ref(`{
      "defaultValue": {
              "value": "3426"
      },
      "valueType": "STRING"
}`)

const totalCount = ref(3)
const current = ref(0)
</script>
```

<style scoped></style>

:::
