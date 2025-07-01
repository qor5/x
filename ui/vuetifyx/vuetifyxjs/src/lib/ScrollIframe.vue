<script setup lang="ts">
import { nextTick, onMounted, onUnmounted, reactive, ref, watch } from 'vue'

const emit = defineEmits(['load'])
const iframe = ref()
const virtualEle = ref()
const parentEle = ref()
const currentEle = ref()
const container = ref()
const storeAddVirtualEle = reactive({
  height: 0,
  containerDataID: ''
})
let resizable = false
const props = defineProps({
  srcdoc: { type: String, required: true },
  width: { type: String },
  virtualElementText: { type: String, default: 'New Component' },
  backgroundColor: { type: String, default: '' },
  virtualElementHeight: { type: Number, default: 100 },
  updateDifferent: { type: Boolean, default: false }
})
const virtualHeight = props.virtualElementHeight

const diffAndUpdate = (oldNode: Node, newNode: Node, deep: number = 0) => {
  if (deep > 0) {
    if (oldNode.nodeType !== newNode.nodeType || oldNode.nodeName !== newNode.nodeName) {
      const parent = oldNode.parentNode
      if (parent) {
        parent.replaceChild(newNode.cloneNode(true), oldNode)
      }
      return
    }

    if (oldNode.nodeType === Node.TEXT_NODE) {
      if (oldNode.nodeValue !== newNode.nodeValue) {
        oldNode.nodeValue = newNode.nodeValue
      }
      return
    }
    const oldElement = oldNode as Element
    const newElement = newNode as Element
    const oldAttrs = oldElement.attributes
    const newAttrs = newElement.attributes
    Array.from(oldAttrs).forEach((attr) => {
      if (!newElement.hasAttribute(attr.name)) {
        oldElement.removeAttribute(attr.name)
      }
    })
    Array.from(newAttrs).forEach((attr) => {
      if (oldElement.getAttribute(attr.name) !== attr.value) {
        oldElement.setAttribute(attr.name, attr.value)
      }
    })
  }

  const oldChildren = Array.from(oldNode.childNodes)
  const newChildren = Array.from(newNode.childNodes)
  const maxLength = Math.max(oldChildren.length, newChildren.length)
  for (let i = 0; i < maxLength; i++) {
    if (!oldChildren[i] && newChildren[i]) {
      oldNode.appendChild(newChildren[i].cloneNode(true))
    } else if (oldChildren[i] && !newChildren[i]) {
      oldChildren[i].remove()
    } else if (oldChildren[i] && newChildren[i]) {
      diffAndUpdate(oldChildren[i], newChildren[i], deep + 1)
    }
  }
}

const resizeObserver = new ResizeObserver((entries) => {
  for (let entry of entries) {
    setIframeDisplay()
  }
})
const setIframeDisplay = () => {
  if (!iframe.value) return

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
const iframeDoc = () => {
  return iframe.value.contentDocument || iframe.value.contentWindow?.document
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

const load = (event: any) => {
  if (!iframe.value || !iframe.value.contentWindow) {
    emit('load', event)
    return
  }
  if (!resizable) {
    resizeObserver.observe(container.value)
    resizable = true
  }

  // Add click event listener to iframe body
  addIframeClickListener()

  emit('load', event)
}

// Check if element or its parents have wrapper-shadow class
const hasWrapperShadowInParentChain = (element: Element): boolean => {
  let current: Element | null = element
  while (current && current !== iframeDoc().body) {
    if (current.classList && current.classList.contains('wrapper-shadow')) {
      return true
    }
    current = current.parentElement
  }
  return false
}

// Add click event listener to iframe body
const addIframeClickListener = () => {
  const body = iframeDoc().querySelector('body')
  if (!body) return

  // Remove existing listener if any
  body.removeEventListener('click', handleIframeClick)

  // Add new click listener
  body.addEventListener('click', handleIframeClick)
}

// Handle iframe body click events
const handleIframeClick = (event: Event) => {
  const target = event.target as Element
  if (!target) return

  // Check if clicked element or its parents have wrapper-shadow class
  if (!hasWrapperShadowInParentChain(target)) {
    // Send message to parent window
    window.parent.postMessage(
      {
        msg_type: 'clickOutsideWrapperShadow'
      },
      '*'
    )
  }
}

const removeHighlightClass = () => {
  const elements = iframeDoc().querySelectorAll('.highlight')
  elements.forEach((el: Element) => (el as HTMLElement).classList.remove('highlight'))
}

const createVirtualElement = () => {
  storeScrollHeight()
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
  removeVirtualElement(true)
  createVirtualElement()
  currentEle.value = current
  parentEle.value = current.parentElement
  parentEle.value?.insertBefore(virtualEle.value, current.nextSibling)
  if (!isElementCompletelyInViewport(virtualEle.value)) {
    scrollTo({ top: virtualEle.value.offsetTop, behavior: 'smooth' })
  }
}

const removeVirtualElement = (createEle: boolean = false) => {
  if (virtualEle.value && parentEle.value) {
    parentEle.value.removeChild(virtualEle.value)
    virtualEle.value = null
    parentEle.value = null
    currentEle.value = null
    if (!createEle) {
      restoreScrollHeight()
    }
  }
}
const appendVirtualElement = () => {
  const app = iframeDoc().getElementById('app') as HTMLElement
  if (!app) {
    return
  }
  if (app == currentEle.value) {
    if (virtualEle.value) {
      scrollTo({ top: virtualEle.value.offsetTop, behavior: 'smooth' })
    }
    return
  }
  removeVirtualElement(true)
  createVirtualElement()
  currentEle.value = app
  parentEle.value = app
  app.appendChild(virtualEle.value)
  scrollTo({ top: virtualEle.value.offsetTop, behavior: 'smooth' })
}
const querySelector = (val: any) => {
  return container.value.querySelector(val)
}

const scrollTo = (data: { top: number; behavior: string }) => {
  const mainElement = iframeDoc().querySelector('.pagebuilder-main') as HTMLElement
  if (mainElement) {
    const mainPaddingTop = parseFloat(window.getComputedStyle(mainElement).paddingTop) || 0
    data.top -= mainPaddingTop
  }
  iframe.value.contentWindow.scrollTo(data)
}
const storeScrollHeight = () => {
  storeAddVirtualEle.height = iframe.value.contentWindow.scrollY
  const hl = iframeDoc().querySelector('.highlight')
  storeAddVirtualEle.containerDataID = ''
  if (hl) {
    storeAddVirtualEle.containerDataID = hl.getAttribute('data-container-id')
  }
}
const restoreScrollHeight = () => {
  scrollTo({ top: storeAddVirtualEle.height, behavior: 'smooth' })
  storeAddVirtualEle.height = 0
  const el = findContainerByDataID(storeAddVirtualEle.containerDataID)
  if (el) {
    el.classList.add('highlight')
  }
  storeAddVirtualEle.containerDataID = ''
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
  scrollTo({ top: current.offsetTop, behavior: 'smooth' })
}
const findContainerByDataID = (containerDataID: string): HTMLElement | undefined => {
  if (!iframe.value) {
    return
  }
  return iframeDoc().querySelector("div[data-container-id='" + containerDataID + "']")
}
const isElementInViewport = (element: HTMLElement) => {
  if (!element) return false
  const { containerTop, containerBottom, targetTop, targetBottom } = getPositionInfo(element)

  return targetBottom >= containerTop && targetTop <= containerBottom
}
const getPositionInfo = (element: HTMLElement) => {
  const doc = iframeDoc()
  const containerScrollTop = doc.documentElement.scrollTop || doc.body.scrollTop
  const containerHeight = iframe.value.clientHeight
  const targetOffsetTop = element.offsetTop
  const targetHeight = element.offsetHeight

  const containerTop = containerScrollTop
  const containerBottom = containerScrollTop + containerHeight
  const targetTop = targetOffsetTop
  const targetBottom = targetOffsetTop + targetHeight
  return { containerTop, containerBottom, targetTop, targetBottom }
}

const isElementCompletelyInViewport = (element: HTMLElement) => {
  if (!element) return false
  const { containerTop, containerBottom, targetTop, targetBottom } = getPositionInfo(element)
  return targetTop >= containerTop && targetBottom <= containerBottom
}
const preloadImage = (src: string) => {
  return new Promise((resolve, reject) => {
    const img = new Image()
    img.src = src
    img.onload = () => resolve(src)
    img.onerror = () => reject(new Error(`Failed to load image: ${src}`))
  })
}

const updateBody = (
  data: { body: string; containerDataID: string; isUpdate: boolean; eventName: string },
  temp: Node
) => {
  if (!iframe.value) {
    return
  }
  iframe.value.contentWindow.postMessage(
    {
      eventName: data.eventName
    },
    '*'
  )
  const bodyEle = iframeDoc().querySelector('body')
  bodyEle.innerHTML = data.body
  setTimeout(() => {
    setIframeDisplay()
    scrollToCurrentContainer(data.containerDataID, data.isUpdate)
  }, 200)
}
const updateIframeBody = (data: {
  body: string
  containerDataID: string
  isUpdate: boolean
  eventName: string
}) => {
  const temp = document.createElement('body')
  temp.innerHTML = data.body
  const imgElements = temp.querySelectorAll('img')
  const imageSrcs = Array.from(imgElements)
    .map((img) => img.src)
    .filter((src) => src)
  if (imageSrcs.length === 0) {
    updateBody(data, temp)
    return
  }
  Promise.all(imageSrcs.map(preloadImage))
    .then(() => {
      updateBody(data, temp)
    })
    .catch((err) => {
      updateBody(data, temp)
    })
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
      width: '100%',
      height: '100%',
      overflow: 'auto',
      backgroundColor: backgroundColor
    }"
  >
    <iframe
      ref="iframe"
      :srcdoc="srcdoc"
      frameborder="0"
      @load="load"
      scrolling="yes"
      :style="{
        height: 'calc(100vh - 88px)',
        width: width,
        display: 'block',
        border: 'none',
        padding: 0,
        margin: 0,
        overflow: 'auto',
        backgroundColor: '#FFF'
      }"
    >
    </iframe>
  </div>
</template>

<style scoped></style>
