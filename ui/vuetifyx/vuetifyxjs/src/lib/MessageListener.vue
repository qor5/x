<script lang="ts">
export default { name: 'vx-message-listener' }
</script>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'

const props = defineProps({
  listenFunc: {
    required: true,
    type: Function
  }
})

onMounted(() => {
  // Initialize global state
  //@ts-ignore
  if (typeof window.vxMessageListenerCount === 'undefined') {
    //@ts-ignore
    window.vxMessageListenerCount = 0
    //@ts-ignore
    window.vxMessageListenerFunc = null
  }
  
  // Add listener if this is the first instance
  //@ts-ignore
  if (window.vxMessageListenerCount === 0) {
    //@ts-ignore
    window.vxMessageListenerFunc = props.listenFunc
    //@ts-ignore
    window.addEventListener('message', window.vxMessageListenerFunc, false)
  }
  
  // Increment reference count
  //@ts-ignore
  window.vxMessageListenerCount++
})

onUnmounted(() => {
  // Decrement reference count
  //@ts-ignore
  if (window.vxMessageListenerCount > 0) {
    //@ts-ignore
    window.vxMessageListenerCount--
  }
  
  // Remove listener if this is the last instance
  //@ts-ignore
  if (window.vxMessageListenerCount === 0 && window.vxMessageListenerFunc) {
    //@ts-ignore
    window.removeEventListener('message', window.vxMessageListenerFunc, false)
    //@ts-ignore
    window.vxMessageListenerFunc = null
  }
})
</script>

<template></template>

<style scoped></style>
