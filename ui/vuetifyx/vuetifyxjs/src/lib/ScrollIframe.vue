<script setup lang="ts">
import {nextTick, onMounted, onUnmounted, ref, watch} from 'vue'

const emit = defineEmits(['load'])
const iframe = ref()
const virtualEle = ref()
const parentEle = ref()
const currentEle = ref()
const container = ref()
const height = ref()
let resizable = false
const props = defineProps({
  srcdoc: {type: String, required: true},
  width: {type: String},
  virtualElementText: {type: String, default: 'New Component'},
  backgroundColor: {type: String, default: ''},
  virtualElementHeight: {type: Number, default: 100}
})
const virtualHeight = props.virtualElementHeight
const resizeContainer = (entry: ResizeObserverEntry) => {
  if (!container.value) {
    return
  }
  if (
    iframe.value &&
    iframe.value.contentWindow &&
    entry.contentRect.width >= iframe.value.contentWindow.document.documentElement.scrollWidth
  ) {
    container.value.style.display = 'flex'
    container.value.style.justifyContent = 'center'
  } else {
    container.value.style.display = ''
    container.value.style.justifyContent = ''
  }
}
const resizeObserver = new ResizeObserver((entries) => {
  for (let entry of entries) {
    resizeContainer(entry)

  }
})
const setIframeDisplay = () => {
  const iframeWidth = iframe.value.style.width.replace('px', '')
  const containerWidth = container.value.offsetWidth
  if (iframeWidth <= containerWidth) {
    container.value.style.display = 'flex'
    container.value.style.justifyContent = 'center'
  } else {
    container.value.style.display = ''
    container.value.style.justifyContent = ''
  }
}

onMounted(() => {
  nextTick(() => {
    setIframeDisplay()
  })
})
watch(
  () => props.width,
  () => {
    iframe.value.style.width = props.width
    nextTick(() => {
      setIframeDisplay()
    })
  }
)
onUnmounted(() => {
  if (!container.value) {
    return
  }
  resizeObserver.unobserve(container.value)
  resizeObserver.disconnect()
  resizable = false
})

const setIframeHeight = () => {
  const bodyEle = iframe.value.contentWindow.document.querySelector('body')
  height.value = bodyEle.scrollHeight
  if (height.value < virtualHeight) {
    height.value = virtualHeight
  }
  setIframeContainerHeight(0)
}

const load = (event: any) => {
  if (!iframe.value || !iframe.value.contentWindow) {
    emit('load', event)
    return
  }
  setIframeHeight()
  if (!resizable) {
    resizeObserver.observe(container.value)
    resizable = true
  }
  const iframeDoc = iframe.value.contentWindow.document
  emit('load', event)
}
const removeHighlightClass = () => {
  const iframeDocument = iframe.value.contentWindow.document
  const elements = iframeDocument.querySelectorAll('.highlight')
  elements.forEach((el: Element) => (el as HTMLElement).classList.remove('highlight'))
}

const setIframeContainerHeight = (h: number) => {
  iframe.value.style.height = height.value + h + 'px'
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
  virtualEle.value.innerHTML = props.virtualElementText
  setIframeContainerHeight(virtualHeight)
}
const addVirtualElement = (data: any) => {
  if (!iframe.value) {
    return
  }
  const current = findContainerByDataID(data)
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
      container.value.scrollTo({top: virtualEle.value.offsetTop, behavior: 'smooth'})
    }
    return
  }
  removeVirtualElement()
  createVirtualElement()
  currentEle.value = app
  parentEle.value = app
  app.appendChild(virtualEle.value)
  container.value.scrollTo({top: virtualEle.value.offsetTop, behavior: 'smooth'})
}
const querySelector = (val: any) => {
  return container.value.querySelector(val)
}

const scrollToCurrentContainer = (data: any, isUpdate: boolean) => {
  if (!iframe.value || !data) {
    return
  }
  removeHighlightClass()
  const current = findContainerByDataID(data)
  if (!current) {
    return
  }
  current.classList.add('highlight')
  const inView = isElementInViewport(current)
  if (isUpdate && inView) {
    return
  }
  container.value.scrollTo({top: current.offsetTop, behavior: 'smooth'})
}
const findContainerByDataID = (containerDataID: string): HTMLElement | undefined => {
  if (!iframe.value) {
    return
  }
  const iframeDocument = iframe.value.contentWindow.document
  if (!iframeDocument) {
    return
  }
  return iframeDocument.querySelector("div[data-container-id='" + containerDataID + "']")
}
const isElementInViewport = (element: HTMLElement) => {
  if (!element) return false

  const containerScrollTop = container.value.scrollTop
  const containerHeight = container.value.clientHeight
  const targetOffsetTop = element.offsetTop
  const targetHeight = element.offsetHeight
  const containerTop = containerScrollTop
  const containerBottom = containerScrollTop + containerHeight
  const targetTop = targetOffsetTop
  const targetBottom = targetOffsetTop + targetHeight
  return targetBottom >= containerTop && targetTop <= containerBottom
}
const preloadImage = (src: string) => {
  return new Promise((resolve, reject) => {
    const img = new Image();
    img.src = src;
    img.onload = () => resolve(src);
    img.onerror = () => reject(new Error(`Failed to load image: ${src}`));
  });
}


const updateBody = (data: { body: string; containerDataID: string; isUpdate: boolean }) => {
  if (!iframe.value) {
    return
  }
  const iframeDocument = iframe.value.contentWindow.document
  if (!iframeDocument) {
    return
  }
  const bodyEle = iframeDocument.querySelector('body')
  bodyEle.innerHTML = data.body;
  setTimeout(() => {
    setIframeHeight()
    setIframeDisplay()
    scrollToCurrentContainer(data.containerDataID, data.isUpdate)
  }, 200)
}
const updateIframeBody = (data: { body: string; containerDataID: string; isUpdate: boolean }) => {
  const tempDiv = document.createElement("div");
  tempDiv.innerHTML = data.body;
  const imgElements = tempDiv.querySelectorAll("img");
  const imageSrcs = Array.from(imgElements)
    .map(img => img.src)
    .filter(src => src);
  if (imageSrcs.length === 0) {
    updateBody(data)
    return;
  }
  Promise.all(imageSrcs.map(preloadImage))
    .then(() => {
      updateBody(data)
    })
    .catch(err => {
      updateBody(data)
    });
}

defineExpose({
  scrollToCurrentContainer,
  addVirtualElement,
  removeVirtualElement,
  appendVirtualElement,
  querySelector,
  updateIframeBody
})
</script>

<template>
  <div
    ref="container"
    :style="{
      height: 'calc(100vh - 88px)',
      width: '100%',
      overflow: 'auto',
      backgroundColor: backgroundColor
    }"
  >
    <iframe
      ref="iframe"
      :srcdoc="srcdoc"
      frameborder="0"
      scrolling="no"
      @load="load"
      :style="{
        width: width,
        display: 'block',
        border: 'none',
        padding: 0,
        margin: 0,
        backgroundColor: '#FFF'
      }"
    >
    </iframe>
  </div>
</template>

<style scoped></style>
