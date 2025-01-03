<template>
  <iframe
    :src="src"
    ref="iframe"
    frameborder="0"
    :style="{
      height: 'calc(100vh - 88px)',
      width: '100%'
    }"
    @load="onIframeLoad"
  >
  </iframe>
</template>

<script setup lang="ts">
import { ref, onUnmounted, defineExpose } from 'vue'
const emit = defineEmits(['load'])
const iframeLoaded = ref(false)
const iframe = ref()
const props = defineProps({
  src: { type: String, required: true }
})

onUnmounted(() => {
  window.removeEventListener('message', handleMessageFromIframe)
})

function onIframeLoad() {
  iframeLoaded.value = true
  window.addEventListener('message', handleMessageFromIframe)
}

function handleMessageFromIframe(event: MessageEvent) {
  // 确保是来自正确的 iframe 源
  if (event.origin !== new URL(props.src).origin) return
  console.log('Message from iframe:', event.data)
}

function sendMessageToIframe(message: unknown) {
  if (iframe && iframe.value.contentWindow) {
    iframe.value.contentWindow.postMessage(message, new URL(props.src).origin)
  }
}

async function getData() {
  const requestId = Date.now()
  const message = { type: 'getData', requestId }

  return new Promise((resolve, reject) => {
    const handleResponse = (event: MessageEvent) => {
      if (event.origin !== new URL(props.src).origin) return
      if (event.data.requestId === requestId) {
        window.removeEventListener('message', handleResponse)
        if (event.data.error) {
          reject(event.data.error)
        } else {
          resolve(event.data.value)
        }
      }
    }

    window.addEventListener('message', handleResponse)
    sendMessageToIframe(message)
  })
}

defineExpose({
  getData
})
</script>
<style scoped></style>
