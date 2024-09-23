# vx-send-variables 输入框

## 基本用法

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const seo = ref()
const value = ref('1eiuh1ihaubfiad')
</script>

<template>
  <vx-send-variables ref="seo">
    <v-chip @click="seo.addTags('Test')" :label="true" :outlined="true" variant="outlined">
      <v-icon icon="mdi-plus"></v-icon>
      Test
    </v-chip>
    <vx-field
      counter
      v-model="value"
      label="Title"
      @focus="seo.tagInputsFocus($refs.setting_title)"
      ref="setting_title"
    >
    </vx-field>
    <vx-field
      type="textarea"
      counter="200"
      v-model="value"
      label="textarea"
      @focus="seo.tagInputsFocus($refs.textarea)"
      ref="textarea"
    >
    </vx-field>
    <div>Value: {{value}} </div>
  </vx-send-variables>
</template>

```
:::