<script setup lang="ts">
//@ts-ignore
import { ActionButton } from 'vuetify-pro-tiptap'
import type { Editor } from '@tiptap/vue-3'
import { Extension } from '@tiptap/core'
import { computed } from 'vue'

const props = defineProps<{
  editor: Editor
  extension: Extension<any, any>
  t: (path: string) => string
}>()

const isDisabled = computed(() => {
  if (!!props.extension.options.isDisabled) {
    return props.extension.options.isDisabled({
      editor: props.editor,
      extension: props.extension,
      t: props.t
    })
  }
  return false
})

const isActive = computed(() => {
  if (!!props.extension.options.isActive) {
    return props.extension.options.isActive({
      editor: props.editor,
      extension: props.extension,
      t: props.t
    })
  }
  return () => {
    return false
  }
})

function onAction() {
  if (!!props.extension.options.onAction) {
    props.extension.options.onAction({
      editor: props.editor,
      extension: props.extension,
      t: props.t
    })
  }
}
</script>

<template>
  <div>
    <ActionButton
      :tooltip="extension.options.tooltip"
      :disabled="isDisabled"
      :is-active="isActive"
      :action="onAction"
    >
      <VIcon v-if="!!extension.options.icon" :icon="extension.options.icon"></VIcon>
    </ActionButton>
  </div>
</template>
