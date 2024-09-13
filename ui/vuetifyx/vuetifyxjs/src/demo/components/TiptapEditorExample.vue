<template>
  <VApp id="app">
    <VContainer>
      <VLocaleProvider locale="zhHans">
        <TiptapEditor
          v-model="content"
          label=""
          :min-height="200"
          :max-height="265"
          :hide-bubble="true"
          :extensions="extensions"
          :disabled="false"
          :readonly="false"
        />
        <p class="mt-4">{{ content }}</p>
      </VLocaleProvider>
    </VContainer>
  </VApp>
</template>

<script setup lang="ts">
import TiptapEditor from '@/lib/TiptapEditor/TiptapEditor.vue'
import { ref } from 'vue'

const extensions = ref([
//   { 
//     name: 'BaseKit', 
//     options: { 
//         placeholder: { 
//             placeholder: 'Enter some text...' 
//         }
//     }
//   },
  { name: 'Bold' },
  { name: 'Italic' },
  { name: 'Underline' },
  { name: 'Strike' },
  { name: 'Code', options: { divider: true } },
  { name: 'Heading' },
  { name: 'TextAlign', options: { types: ['heading', 'paragraph', 'image'] } },
  { name: 'FontFamily' },
  { name: 'FontSize' },
  { name: 'Color' },
  { name: 'Highlight', options: { divider: true } },
//   // { name: 'SubAndSuperScript', options: { divider: true } },
  { name: 'BulletList' },
  { name: 'OrderedList', options: { divider: true } },
//   // { name: 'TaskList' },
  { name: 'Indent', options: { divider: true } },
  { name: 'Link' },
  { name: 'Image', options: { hiddenTabs: ["upload"] } },
  { name: 'Video', options: { divider: true } },
  // { name: 'Table', options: { divider: true } },
  { name: 'Blockquote' },
  { name: 'HorizontalRule' },
  { name: 'CodeBlock', options: { divider: true } },
  { name: 'Clear' },
  { name: 'History', options: { divider: true } },
//   { name: 'Fullscreen' },
  // { 
  //   name: 'Callback', 
  //   options: { 
  //       divider: true,
  //       tooltip: "Image",
  //       icon:"mdi-image",
  //       isDisabled: ({editor}: {editor: any, extension: any}) => {
  //           return !editor.can().setImage({})
  //       },
  //       isActive: ({editor}: {editor: any, extension: any}) => {
  //           return () => editor.isActive('image') || false
  //       },
  //       onAction: ({editor}: {editor: any, extension: any}) => {
  //           editor.chain().focus().setImage({
  //               display: 'block', // 'block' 'inline' 'left' 'right'
  //               src: 'https://picsum.photos/300/200',
  //               alt: 'Random image',
  //               width: 300,
  //               height: 200,
  //           }).run()
  //       }
  //   }
  // },
])

// slack-like
// const extensions = ref([
//   { 
//     name: 'BaseKit', 
//     options: { 
//         placeholder: { 
//             placeholder: 'Jot something down...' 
//         }
//     }
//   },
//   { name: 'Bold' },
//   { name: 'Italic' },
//   { name: 'Strike', options: { divider: true } },
//   { name: 'Link' , options: { divider: true } },
//   { name: 'OrderedList'},
//   { name: 'BulletList', options: { divider: true } },
//   { name: 'Blockquote',  options: { divider: true }  },
//   { name: 'Code'},
//   { name: 'CodeBlock'},
// ])

const content = ref(`<h2>
            Hi there,
          </h2>
          <p>
            this is a <em>basic</em> example of <strong>Tiptap</strong>. Sure, there are all kind of basic text styles you’d probably expect from a text editor. But wait until you see the lists:
          </p>
          <ul>
            <li>
              That’s a bullet list with one …
            </li>
            <li>
              … or two list items.
            </li>
          </ul>
          <p>
            Isn’t that great? And all of that is editable. But wait, there’s more. Let’s try a code block:
          </p>
          <pre><code class="language-css">body {
    display: none;
  }</code></pre>
          <p>
            I know, I know, this is impressive. It’s only the tip of the iceberg though. Give it a try and click a little bit around. Don’t forget to check the other examples too.
          </p>
          <blockquote>
            Wow, that’s amazing. Good work, boy! 👏
            <br />
            — Mom
          </blockquote>`)
</script>
