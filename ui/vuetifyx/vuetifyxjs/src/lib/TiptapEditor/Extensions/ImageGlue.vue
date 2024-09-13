<script setup lang="ts">
import type { Editor } from '@tiptap/vue-3'
import { ref, inject, onMounted, onBeforeUnmount } from 'vue'

interface Props {
  value?: any
  editor: Editor
  destroy?: () => void
}

const props = withDefaults(defineProps<Props>(), {
  value: () => ({}),
  destroy: undefined
})

const elementRef = ref<HTMLElement | null>(null)

// 定义 click 函数的参数类型
interface ClickParams {
  editor: Editor
  value: any
}

// 修改 inject 的类型参数，明确指定 editor 和 value 的类型
const click = inject<(params: ClickParams) => void>('__imageGlueClick__')

const emitClick = () => {
  if (click) {
    click({ editor: props.editor, value: props.value })
  }
}

onMounted(() => {
  const parentElement = elementRef.value?.parentElement
  if (parentElement) {
    parentElement.addEventListener('click', emitClick)
  }
})

onBeforeUnmount(() => {
  const parentElement = elementRef.value?.parentElement
  if (parentElement) {
    parentElement.removeEventListener('click', emitClick)
  }
})
</script>

<template>
  <div ref="elementRef">
  </div>
</template>
