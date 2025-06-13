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
  //@ts-ignore
  if (!window.vxMessageListenerFunc) {
    //@ts-ignore
    window.vxMessageListenerFunc = props.listenFunc
    //@ts-ignore
    window.addEventListener('message', window.vxMessageListenerFunc, false)
  }
})

onUnmounted(() => {
  //@ts-ignore
  if (window.vxMessageListenerFunc) {
    //@ts-ignore
    window.removeEventListener('message', window.vxMessageListenerFunc, false)
    //@ts-ignore
    window.vxMessageListenerFunc = null
  }
})
</script>

<template></template>

<style scoped></style>
