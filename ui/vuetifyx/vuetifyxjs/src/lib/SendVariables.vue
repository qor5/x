<script setup lang="ts">
import { ref } from 'vue'

const vnode = ref()

const tagInputsFocus = (v: any) => {
  vnode.value =  v.$.ctx
}
const addTags = (tag: any) => {
  if (!vnode.value) {
    return
  }
  let lazyValue = vnode.value.modelValue
  let selectionStart = vnode.value.selectionStart
  let selectionEnd = vnode.value.selectionEnd
  const inputFiled = vnode.value.$el.querySelector('input') || vnode.value.$el.querySelector('textarea')
  if (inputFiled) {
    selectionStart = inputFiled.selectionStart
    selectionEnd = inputFiled.selectionEnd
  }
  let startString = lazyValue.substring(0, selectionStart)
  let endString = lazyValue.substring(selectionEnd, lazyValue.length)

  vnode.value.$emit('update:modelValue', startString + '{{' + tag + '}}' + endString)
  inputFiled.focus()
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
