<template>
  <div class="border-thin rounded">
    <VuetifyTiptap v-bind="processedAttrs" v-model="model">
      <template #bottom>
        <div style="display: none"></div>
      </template>
    </VuetifyTiptap>
  </div>
</template>

<script setup lang="ts">
import { computed, useAttrs } from 'vue'
import { Extension } from '@tiptap/core'
import { VuetifyTiptap } from 'vuetify-pro-tiptap'
import { BaseKit, Bold, Italic, Underline, Strike, Color, Highlight, Heading, TextAlign, FontFamily, FontSize, SubAndSuperScript, BulletList, OrderedList, TaskList, Indent, Link, Image, Video, Table, Blockquote, HorizontalRule, Code, CodeBlock, Clear, Fullscreen, History } from 'vuetify-pro-tiptap'
const extensionMap = { BaseKit, Bold, Italic, Underline, Strike, Color, Highlight, Heading, TextAlign, FontFamily, FontSize, SubAndSuperScript, BulletList, OrderedList, TaskList, Indent, Link, Image, Video, Table, Blockquote, HorizontalRule, Code, CodeBlock, Clear, Fullscreen, History }
type ExtensionName = keyof typeof extensionMap

const model: string | object | undefined = defineModel()
const attrs = useAttrs()

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
  { name: 'TextAlign' }, // TODO: unavailable
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
  { name: 'History', options: { divider: true } },
  { name: 'Fullscreen' },
]

const processedAttrs = computed(() => {
  let extensions = (attrs.extensions as { name: ExtensionName; options?: any }[]) || []
  if (extensions.length <=0 ) {
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
    extensions = [...extensions,  { name: 'History' }]
  }
  return {
    ...attrs,
    'hide-bubble': attrs['hide-bubble'] !== undefined ? attrs['hide-bubble'] : true,
    extensions: resolvedExtensions(extensions)
  }
})
</script>
