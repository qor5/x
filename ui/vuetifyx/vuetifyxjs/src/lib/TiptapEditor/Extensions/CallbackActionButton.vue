<script setup lang="ts">
import { ActionButton } from 'vuetify-pro-tiptap'
import type { Editor } from '@tiptap/vue-3'
import { Extension } from '@tiptap/core'

const props = defineProps<{
  editor: Editor
  extension: Extension<any, any>
  t: (path: string) => string
}>()

function onClick() {
  if (!!props.extension.options.click) {
    props.extension.options.click({editor:props.editor, extension:props.extension, t:props.t})
  }
}
</script>

<template>
  <ActionButton :tooltip="extension.options.tooltip" @click="onClick">
    <VIcon v-if="!!extension.options.icon" :icon="extension.options.icon"></VIcon>
  </ActionButton>
</template>
