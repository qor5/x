<template>
  <div class="rounded tiptap-wrapper">
    <!-- <VuetifyViewer v-if="readonly" v-bind="processedAttrs" :value="model">
    </VuetifyViewer> -->
    <VuetifyTiptap ref="tiptapRef" v-bind="processedAttrs" v-model="model">
      <template #bottom>
        <div style="display: none"></div>
      </template>
    </VuetifyTiptap>
  </div>
</template>
<script setup lang="ts">
import { computed, defineExpose, provide, ref, useAttrs, watchEffect } from 'vue'
import { useLocale } from 'vuetify'
import {
  BaseKit,
  Blockquote,
  Bold,
  BulletList,
  Clear,
  Code,
  CodeBlock,
  Color,
  FontFamily,
  FontSize,
  Fullscreen,
  Heading,
  Highlight,
  History,
  HorizontalRule,
  Image,
  Indent,
  Italic,
  Link,
  locale,
  OrderedList,
  Strike,
  SubAndSuperScript,
  Table,
  TaskList,
  TextAlign,
  Underline,
  useContext,
  Video,
  //@ts-ignore
  VuetifyTiptap,
  //@ts-ignore
  HtmlView,
  Paragraph
} from 'vuetify-pro-tiptap'

import { Extension } from '@tiptap/core'
import ImageGlue from '@/lib/TiptapEditor/Extensions/ImageGlue.vue'
import Callback from '@/lib/TiptapEditor/Extensions/CallbackActionButton'
// i18n
const { current: currentLocale } = useLocale()
watchEffect(() => {
  locale.setLang(currentLocale.value)
})
const tiptapRef = ref()
defineExpose({
  focus() {
    const editorInstance = tiptapRef.value?.editor
    if (editorInstance) {
      editorInstance.commands.focus()
    }
  },
  setSelectionRange(from: number, to: number) {
    const editorInstance = tiptapRef.value?.editor
    if (editorInstance) {
      editorInstance.commands.focus(from)
    }
  },
  editor: computed(() => {
    return tiptapRef.value?.editor
  })
})
const extensionMap = {
  BaseKit,
  Bold,
  Italic,
  Underline,
  Strike,
  Color,
  Highlight,
  Heading,
  TextAlign,
  FontFamily,
  FontSize,
  SubAndSuperScript,
  BulletList,
  OrderedList,
  TaskList,
  Indent,
  Link,
  Image,
  Video,
  Table,
  Blockquote,
  HorizontalRule,
  Code,
  CodeBlock,
  Clear,
  Fullscreen,
  History,
  Callback,
  ImageGlue,
  HtmlView,
  Paragraph
}
type ExtensionName = keyof typeof extensionMap

const model: string | object | undefined = defineModel()
const attrs = useAttrs()

const props = withDefaults(
  defineProps<{
    readonly?: boolean | undefined
  }>(),
  {
    readonly: false
  }
)

function resolvedExtensions(
  extensions: Array<{ name: ExtensionName; options?: any }>
): Array<Extension<any, any>> {
  return extensions
    .map((extension) => {
      const extensionInstance = extensionMap[extension.name]
      if (extensionInstance) {
        return extension.options
          ? extensionInstance.configure(extension.options)
          : extensionInstance
      }
      console.warn(`Extension ${extension.name} not found in extensionMap.`)
      return null
    })
    .filter((extension): extension is Extension<any, any> => extension !== null)
}

const defaultExtensions: Array<{ name: ExtensionName; options?: any }> = [
  {
    name: 'BaseKit',
    options: {
      placeholder: {
        placeholder: 'Enter some text...'
      }
    }
  },
  { name: 'Bold' },
  { name: 'Italic' },
  { name: 'Underline' },
  { name: 'Strike' },
  { name: 'Code', options: { divider: true } },
  { name: 'Heading' },
  { name: 'TextAlign', options: { types: ['heading', 'paragraph', 'image'] } },
  // { name: 'FontFamily' },
  // { name: 'FontSize' },
  { name: 'Color' },
  { name: 'Highlight', options: { divider: true } },
  // { name: 'SubAndSuperScript', options: { divider: true } },
  { name: 'BulletList' },
  { name: 'OrderedList' },
  // { name: 'TaskList' },
  { name: 'Indent', options: { divider: true } },
  { name: 'Link' },
  { name: 'Video', options: { divider: true } },
  // { name: 'Table', options: { divider: true } },
  { name: 'Blockquote' },
  { name: 'HorizontalRule' },
  { name: 'CodeBlock', options: { divider: true } },
  { name: 'Clear' },
  { name: 'History', options: { divider: true } }
  // { name: 'Fullscreen' },
]

let imageGlueClick: any = undefined

provide('__imageGlueClick__', ({ editor, value }: { editor: any; value: any }) => {
  if (imageGlueClick) {
    imageGlueClick({ editor, value, window: window })
  }
})

const processedAttrs = computed(() => {
  let extensions = (attrs.extensions as { name: ExtensionName; options?: any }[]) || []
  if (extensions.length <= 0) {
    extensions = [...defaultExtensions]
  }
  if (!extensions.some((extension) => extension.name === 'BaseKit')) {
    extensions = [{ name: 'BaseKit' }, ...extensions]
  }
  if (!extensions.some((extension) => extension.name === 'History')) {
    if (extensions.length > 0) {
      const lastExtension = extensions[extensions.length - 1]
      if (!lastExtension.options) {
        lastExtension.options = {}
      }
      lastExtension.options.divider = true
    }
    // TODO: hideable ?
    extensions = [...extensions, { name: 'History' }]
  }
  const imageGlueIdx = extensions.findIndex((extension) => extension.name === 'ImageGlue')
  if (imageGlueIdx >= 0) {
    const imageGlueOptions = extensions[imageGlueIdx].options
    imageGlueClick = imageGlueOptions?.onClick

    extensions[imageGlueIdx] = {
      name: 'Image',
      options: {
        ...imageGlueOptions,
        dialogComponent: () => {
          return ImageGlue
        }
      }
    }
  }
  return {
    ...attrs,
    disabled: !!attrs.disabled || props.readonly,
    'disable-toolbar': !!attrs['disable-toolbar'] || !!attrs.disabled || props.readonly,
    'hide-toolbar': !!attrs['hide-toolbar'] || props.readonly,
    'hide-bubble':
      attrs['hide-bubble'] !== undefined ? !!attrs['hide-bubble'] || props.readonly : true,
    extensions: resolvedExtensions(extensions),
    style: '',
    class: ''
  }
})

// vuetify-pro-tiptap bug
// we need to force update the theme
const { state } = useContext()
watchEffect(() => {
  state.defaultMarkdownTheme = (attrs['markdown-theme'] as string) || 'default'
})
</script>

<style lang="scss">
.tiptap-wrapper .vuetify-pro-tiptap-editor__content {
  cursor: text !important;
}
</style>
