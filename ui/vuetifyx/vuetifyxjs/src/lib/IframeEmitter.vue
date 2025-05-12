<template>
  <iframe
    class="iframe-emitter-wrapper"
    :src="src"
    ref="iframe"
    frameborder="0"
    @load="onIframeLoad"
  >
  </iframe>
</template>

<script setup lang="ts">
import { ref, onUnmounted, defineExpose, defineEmits, onMounted } from 'vue'
const iframeLoaded = ref(false)
const iframe = ref()
const props = defineProps({
  src: { type: String, required: true }
})
const emits = defineEmits(['load', 'data'])

onUnmounted(() => {
  window.removeEventListener('message', sendMessageToParent)
  for (const [, pending] of pendingMap.entries()) {
    window.removeEventListener('message', pending.handler)
  }
  pendingMap.clear()
})

function onIframeLoad() {
  iframeLoaded.value = true
  emits('load', null)
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

type PendingHandler = {
  resolve: (value: any) => void
  reject: (reason?: any) => void
  handler: (event: MessageEvent) => void
  createdAt: number
}

// 用于追踪所有活跃的事件绑定
const pendingMap = new Map<string, PendingHandler>()

// 清理超过 5 分钟未响应的事件
function cleanupStalePendingEvents() {
  const now = Date.now()
  const threshold = 5 * 60 * 1000 // 5 minutes

  for (const [requestId, pending] of pendingMap.entries()) {
    if (now - pending.createdAt > threshold) {
      console.warn(`[emit] 清理超时事件: ${requestId}`)
      window.removeEventListener('message', pending.handler)
      pendingMap.delete(requestId)
    }
  }
}

setInterval(cleanupStalePendingEvents, 60 * 1000)

async function emit(eventName: string, data: any, id?: string): Promise<any> {
  const requestId = id || `${eventName}-${Date.now()}-${Math.random()}`

  // 移除旧的同类型事件监听器
  for (const [key, pending] of pendingMap.entries()) {
    if (key.startsWith(`${eventName}-`)) {
      window.removeEventListener('message', pending.handler)
      pendingMap.delete(key)
    }
  }

  const message = {
    type: eventName,
    data,
    requestId,
    source: 'from-outside'
  }

  return new Promise((resolve, reject) => {
    const handleResponse = (event: MessageEvent) => {
      const originMatches = event.origin === getPropsOrigin(props.src)
      const idMatches = event.data?.requestId === requestId

      if (!originMatches || !idMatches) {
        window.removeEventListener('message', handleResponse)
        pendingMap.delete(requestId)
        return
      }

      window.removeEventListener('message', handleResponse)
      pendingMap.delete(requestId)

      if (event.data.error) {
        reject(event.data)
      } else {
        resolve(event.data)
      }
    }

    pendingMap.set(requestId, {
      resolve,
      reject,
      handler: handleResponse,
      createdAt: Date.now()
    })

    window.addEventListener('message', handleResponse)
    sendMessageToIframe(message)
  })
}

// send message from iframe to parent
function sendMessageToParent(event: MessageEvent) {
  // Only process messages from the expected origin
  if (event.origin !== getPropsOrigin(props.src)) return

  // Skip messages that are responses to emit method calls (they have requestId property)
  if (event.data && event.data.source === 'from-outside') return

  // Forward the actual message data to parent
  emits('data', event.data)
}

onMounted(() => {
  window.addEventListener('message', sendMessageToParent)
})

defineExpose({
  emit
})
</script>
<style scoped>
.iframe-emitter-wrapper {
  height: 50vh;
  width: 100%;
}
</style>
