<script setup lang="ts">
import { onMounted, ref } from 'vue'

const iframe = ref()
const container = ref()
const props = defineProps({
  srcdoc: { type: String, required: true },
  iframeHeightName: { type: String, required: true },
  iframeHeight: { type: String, required: true },
  width: { type: String }
})

const load = (event: any) => {
  let height = iframe.value.contentWindow.document.documentElement.scrollHeight + 'px'
  iframe.value.style.height = height
  document.cookie = `${props.iframeHeightName}=` + height
  container.value.style.height = height
}
const scrollToCurrentContainer = (data: any) => {
  if (!iframe.value) {
    return
  }
  const current = iframe.value.contentWindow.document.body.querySelector(
    "div[data-container-id='" + data + "']"
  ) as HTMLElement
  if (!current) {
    return
  }
  window.parent.scroll({ top: current.parentElement?.offsetTop, behavior: 'smooth' })
}
defineExpose({ scrollToCurrentContainer })
</script>

<template>
  <div class="mx-auto" ref="container" :style="{ height: iframeHeight, width: width }">
    <iframe
      ref="iframe"
      :srcdoc="srcdoc"
      frameborder="0"
      scrolling="no"
      @load="load"
      :style="{
        width: '100%',
        display: 'block',
        border: 'none',
        padding: 0,
        margin: 0
      }"
    >
    </iframe>
  </div>
</template>

<style scoped></style>
