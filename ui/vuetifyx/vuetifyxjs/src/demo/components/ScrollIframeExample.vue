<script setup lang="ts">
import ScrollIframe from '@/lib/ScrollIframe.vue'
import { ref } from 'vue'
const srcdoc = ref(`
<!DOCTYPE html>
<body>
<div >hello world</div>
</body>
`)
</script>

<template>
  <scroll-iframe
    iframe-height="1000px"
    iframe-height-name="_iframeHeight"
    :srcdoc="srcdoc"
  ></scroll-iframe>
</template>

<style scoped></style>
