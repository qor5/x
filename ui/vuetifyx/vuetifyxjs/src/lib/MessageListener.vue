<script lang="ts">
export default { name: 'vx-message-listener' }
</script>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'

const props = defineProps({
  listenFunc: {
    required: true,
    type: Function
  },
  name: {
    type: String,
    required: true
  }
})

//@ts-ignore
window.vxMessageListenerFunc = window.vxMessageListenerFunc || {}
onMounted(() => {
  //@ts-ignore
  if (!window.vxMessageListenerFunc[props.name]) {
    //@ts-ignore
    window.vxMessageListenerFunc[props.name] = props.listenFunc
    //@ts-ignore
    window.addEventListener(props.name, window.vxMessageListenerFunc[props.name], false)
  }
})

onUnmounted(() => {
  //@ts-ignore
  if (window.vxMessageListenerFunc[props.name]) {
    //@ts-ignore
    window.removeEventListener(props.name, window.vxMessageListenerFunc[props.name], false)
    //@ts-ignore
    window.vxMessageListenerFunc[props.name] = null
  }
})
</script>

<template></template>

<style scoped></style>
