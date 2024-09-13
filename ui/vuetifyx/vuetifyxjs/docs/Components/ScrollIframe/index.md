# vx-scroll-iframe 输入框

## 基本用法

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const scroll = ref()
const srcdoc = ref(`
<!DOCTYPE html>
<body>
   <div id="app">
   <div data-container-id="test1"> test 1 </div>
<div data-container-id="test2"> test 2 </div>
    </div>
</body>
`)
const add = (s: string) => {
  scroll.value.addVirtualElement(s)
}
const remove = () => {
  scroll.value.removeVirtualElement()
}
const append = () => {
  scroll.value.appendVirtualElement()
}
</script>

<template>
  <vx-scroll-iframe
    ref="scroll"
    iframe-height="1000px"
    iframe-height-name="_iframeHeight"
    :srcdoc="srcdoc"
  ></vx-scroll-iframe>
  <v-btn @click="add('test1')">add after test1</v-btn>
  <v-btn @click="add('test2')">add after test2</v-btn>
  <v-btn @click="remove()">removeVirtual</v-btn>
  <v-btn @click="append()">append</v-btn>
</template>

<style scoped></style>


```
:::