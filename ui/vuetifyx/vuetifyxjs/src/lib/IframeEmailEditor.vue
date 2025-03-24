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
import { ref, onUnmounted, defineExpose, defineEmits } from 'vue'
const iframeLoaded = ref(false)
const iframe = ref()
const props = defineProps({
  src: { type: String, required: true }
})
const emits = defineEmits(['load'])

onUnmounted(() => {
  window.removeEventListener('message', handleMessageFromIframe)
})

function onIframeLoad() {
  iframeLoaded.value = true
  window.addEventListener('message', handleMessageFromIframe)
  emits('load', null)
}

function handleMessageFromIframe(event: MessageEvent) {
  if (event.origin !== getPropsOrigin(props.src)) return
  // console.log('Message from iframe:', event.data)
}

function getPropsOrigin(src: string) {
  return src.startsWith('http') ? new URL(src).origin : window.location.origin
}

function sendMessageToIframe(message: unknown) {
  if (iframe && iframe.value.contentWindow) {
    const origin = getPropsOrigin(props.src)
    iframe.value.contentWindow.postMessage(message, origin)
  }
}

async function emit(eventName: string, data: unknown) {
  const requestId = Date.now()
  const message = { type: eventName, data, requestId }

  return new Promise((resolve, reject) => {
    const handleResponse = (event: MessageEvent) => {
      if (event.origin !== getPropsOrigin(props.src)) return
      if (event.data.requestId === requestId) {
        window.removeEventListener('message', handleResponse)
        if (event.data.error) {
          reject(event.data)
        } else {
          resolve(event.data)
        }
      }
    }

    window.addEventListener('message', handleResponse)
    sendMessageToIframe(message)
  })
}

defineExpose({
  emit
})
</script>
<style scoped></style>
