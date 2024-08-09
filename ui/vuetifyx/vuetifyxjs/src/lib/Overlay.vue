<script setup lang="ts">
import { computed, onMounted, onUnmounted, reactive, ref, watch } from 'vue'

const props = defineProps({
  modelValue: { type: Boolean },
  maxWith: { type: Number, default: 400 }
})

const currentElement = ref()
const rectVal = ref()
const elementStyle = reactive({
  top: 0,
  left: 0,
  height: 0,
  width: 0,
  center: false
})
const emit = defineEmits(['update:modelValue', 'afterLeave'])
const content = ref()
const visible = ref(false)
const offsetTop = computed(() => {
  const top = elementStyle.top + elementStyle.height / 2
  if (
    top + contentHeight.value > window.innerHeight &&
    window.innerHeight - top - contentHeight.value
  ) {
    if (top - contentHeight.value < 0) {
      return '0'
    }
    return top - contentHeight.value + 'px'
  }
  return top + 'px'
})

const offsetLeft = computed(() => {
  const left = elementStyle.left + elementStyle.width / 2
  if (
    left + contentWidth.value > window.innerWidth &&
    window.innerWidth - left - contentWidth.value
  ) {
    if (left - contentWidth.value < 0) {
      return '0'
    }
    return left - contentWidth.value + 'px'
  }
  return left + 'px'
})
const contentWidth = computed(() => {
  if (content.value) {
    return content.value.offsetWidth
  }
  return 0
})
const contentHeight = computed(() => {
  if (content.value) {
    return content.value.offsetHeight
  }
  return 0
})
const showCenter = () => {
  elementStyle.center = true
}

const showByElement = (e: any) => {
  elementStyle.center = false
  currentElement.value = e.currentTarget
  reloadRect()
}

const reloadRect = () => {
  const rect = currentElement.value.getBoundingClientRect()
  elementStyle.top = rect.top
  elementStyle.left = rect.left
  elementStyle.width = rect.width
  elementStyle.height = rect.height
}
const reloadIframeRect = () => {
  const rect = rectVal.value
  const iframeRect = currentElement.value.getBoundingClientRect()
  elementStyle.top = rect.top + iframeRect.top
  elementStyle.left = rect.left + iframeRect.left
  elementStyle.width = rect.width
  elementStyle.height = rect.height
}
const showByIframe = (
  element: HTMLElement,
  rect: {
    top: number
    left: number
    width: number
    height: number
  }
) => {
  elementStyle.center = false
  if (!element) {
    return
  }
  if (element.tagName === 'IFRAME') {
    currentElement.value = element
  } else {
    currentElement.value = element.querySelector('iframe')
    if (!currentElement.value) {
      return
    }
  }
  rectVal.value = rect
  reloadIframeRect()
}

defineExpose({ showByElement, showByIframe, showCenter })

watch(
  () => props.modelValue,
  () => {
    visible.value = props.modelValue
  }
)
const emitUnVisible = () => {
  emit('update:modelValue', false)
}
const afterLeave = (e: any) => {
  emit('afterLeave', e)
  currentElement.value = null
}
const loadWindowSize = () => {
  if (!currentElement.value) {
    return
  }
  if (currentElement.value?.tagName === 'IFRAME') {
    reloadIframeRect()
    return
  }
  reloadRect()
}

onMounted(() => {
  window.addEventListener('resize', loadWindowSize)
})
onUnmounted(() => {
  window.removeEventListener('resize', loadWindowSize)
})
</script>

<template>
  <v-dialog
    v-if="elementStyle.center"
    v-model="visible"
    :max-width="maxWith"
    @update:model-value="emitUnVisible"
    @after-leave="afterLeave($event)"
  >
    <div ref="content">
      <slot></slot>
    </div>
  </v-dialog>
  <v-overlay
    v-else
    v-model="visible"
    :style="{ top: offsetTop, left: offsetLeft }"
    @update:model-value="emitUnVisible"
    @after-leave="afterLeave($event)"
  >
    <div ref="content">
      <slot></slot>
    </div>
  </v-overlay>
</template>

<style scoped></style>
