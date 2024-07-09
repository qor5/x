<script setup lang="ts">
import { ref } from 'vue'

const iframe = ref()
const virtualEle = ref()
const parentEle = ref()
const currentEle = ref()
const container = ref()
const height = ref()
const props = defineProps({
  srcdoc: { type: String, required: true },
  iframeHeightName: { type: String, required: true },
  iframeHeight: { type: String, required: true },
  width: { type: String },
  virtualEleText: { type: String, default: 'New Component' },
  virtualEleHeight: { type: Number, default: 100 },
  containerDataId: { type: String }
})
const virtualHeight = props.virtualEleHeight

const load = (event: any) => {
  if (!iframe.value) {
    return
  }
  height.value = iframe.value.contentWindow.document.documentElement.scrollHeight
  if (height.value < virtualHeight) {
    height.value = virtualHeight
  }
  setIframeContainerHeight(0)
  document.cookie = `${props.iframeHeightName}=` + height.value + 'px'
  scrollToCurrentContainer(props.containerDataId)
}
const removeHighlightClass = () => {
  const iframeDocument = iframe.value.contentDocument || iframe.value.contentWindow.document
  const elements = iframeDocument.querySelectorAll('.highlight')
  elements.forEach((el: Element) => (el as HTMLElement).classList.remove('highlight'))
}

const setIframeContainerHeight = (h: number) => {
  iframe.value.style.height = height.value + h + 'px'
  container.value.style.height = height.value + h + 'px'
}
const scrollToCurrentContainer = (data: any) => {
  if (!iframe.value || !data) {
    return
  }
  removeHighlightClass()
  const current = iframe.value.contentWindow.document.body.querySelector(
    "div[data-container-id='" + data + "']"
  ) as HTMLElement
  if (!current) {
    return
  }
  current.classList.add('highlight')
  window.parent.scroll({ top: current.offsetTop, behavior: 'smooth' })
}

const createVirtualElement = () => {
  removeHighlightClass()
  virtualEle.value = document.createElement('div')
  virtualEle.value.style.height = virtualHeight + 'px'
  virtualEle.value.style.border = '2px dashed #3E63DD'
  virtualEle.value.style.margin = '8px'
  virtualEle.value.style.fontSize = '18px'
  virtualEle.value.style.color = '#3E63DD'
  virtualEle.value.style.display = 'flex'
  virtualEle.value.style.justifyContent = 'center'
  virtualEle.value.style.alignItems = 'center'
  virtualEle.value.innerHTML = props.virtualEleText
  setIframeContainerHeight(virtualHeight)
}
const addVirtualElement = (data: any) => {
  if (!iframe.value) {
    return
  }
  const current = iframe.value.contentWindow.document.body.querySelector(
    "div[data-container-id='" + data + "']"
  ) as HTMLElement
  if (!current) {
    return
  }
  if (currentEle.value == current) {
    return
  }
  removeVirtualElement()
  createVirtualElement()
  currentEle.value = current
  parentEle.value = current.parentElement
  parentEle.value?.insertBefore(virtualEle.value, current.nextSibling)
}

const removeVirtualElement = () => {
  if (virtualEle.value && parentEle.value) {
    parentEle.value.removeChild(virtualEle.value)
    container.value.style.height = height.value + 'px'
    virtualEle.value = null
    parentEle.value = null
    currentEle.value = null
  }
}
const appendVirtualElement = () => {
  const app = iframe.value.contentWindow.document.getElementById('app') as HTMLElement
  if (!app) {
    return
  }
  if (app == currentEle.value) {
    if (virtualEle.value) {
      window.parent.scroll({ top: virtualEle.value.offsetTop, behavior: 'smooth' })
    }
    return
  }
  removeVirtualElement()
  createVirtualElement()
  currentEle.value = app
  parentEle.value = app
  app.appendChild(virtualEle.value)
  window.parent.scroll({ top: virtualEle.value.offsetTop, behavior: 'smooth' })
}
defineExpose({
  scrollToCurrentContainer,
  addVirtualElement,
  removeVirtualElement,
  appendVirtualElement
})
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
