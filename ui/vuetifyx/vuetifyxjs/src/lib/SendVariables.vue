<script setup lang="ts">
import { ref } from 'vue'

const vnode = ref()

const tagInputsFocus = (v: any) => {
  vnode.value = v
}
const addTags = (tag: any) => {
  if (!vnode.value) {
    return
  }
  let lazyValue = vnode.value.modelValue
  let selectionStart = vnode.value.selectionStart
  let selectionEnd = vnode.value.selectionEnd
  const input = vnode.value.$el.querySelector('input')
  if (input) {
    selectionStart = input.selectionStart
    selectionEnd = input.selectionEnd
  }
  let startString = lazyValue.substring(0, selectionStart)
  let endString = lazyValue.substring(selectionEnd, lazyValue.length)

  vnode.value.$emit('update:modelValue', startString + '{{' + tag + '}}' + endString)
  vnode.value.focus()
}
defineExpose({
  tagInputsFocus,
  addTags
})
</script>

<template>
  <div>
    <slot></slot>
  </div>
</template>

<style scoped></style>
