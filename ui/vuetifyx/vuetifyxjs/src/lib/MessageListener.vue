<script lang="ts">
export default { name: 'vx-message-listener' }
</script>

<script setup lang="ts">
import { onMounted } from 'vue'

const props = defineProps({
  listenFunc: {
    required: true,
    type: Function
  }
})
onMounted(() => {
  //@ts-ignore
  if (!window.vxMessageListener) {
    //@ts-ignore
    window.addEventListener('message', props.listenFunc, false)
    //@ts-ignore
    window.vxMessageListener = true
  }
})
</script>

<template></template>

<style scoped></style>
