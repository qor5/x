<!--   @snippet_begin(TipTapEditorVueComponent) -->
<script setup lang="ts">
import "@/lib/sass/main.scss";
import { Editor, EditorContent } from "@tiptap/vue-3";
import Blockquote from "@tiptap/extension-blockquote";
import Bold from "@tiptap/extension-bold";
import BulletList from "@tiptap/extension-bullet-list";
import Code from "@tiptap/extension-code";
import CodeBlock from "@tiptap/extension-code-block";
import HardBreak from "@tiptap/extension-hard-break";
import Heading from "@tiptap/extension-heading";
import History from "@tiptap/extension-history";
import Italic from "@tiptap/extension-italic";
import Link from "@tiptap/extension-link";
import ListItem from "@tiptap/extension-list-item";
import OrderedList from "@tiptap/extension-ordered-list";
import Strike from "@tiptap/extension-strike";
import Table from "@tiptap/extension-table";
import TableCell from "@tiptap/extension-table-cell";
import TableHeader from "@tiptap/extension-table-header";
import TableRow from "@tiptap/extension-table-row";
import Underline from "@tiptap/extension-underline";

import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import Text from "@tiptap/extension-text";

import Icon from "./Icon.vue";

import { onMounted, ref } from "vue";

const emit = defineEmits(["update:modelValue"]);
const props = defineProps({ modelValue: String });
const editor = ref();

onMounted(() => {
  editor.value = new Editor({
    content: props.modelValue,
    extensions: [
      Blockquote,
      Bold,
      BulletList,
      Code,
      CodeBlock,
      HardBreak,
      Heading,
      History,
      Document,
      Paragraph,
      Text,
      Italic,
      Link,
      ListItem,
      OrderedList,
      Strike,
      Table.configure({
        resizable: true,
      }),
      TableCell,
      TableHeader,
      TableRow,
      Underline,
    ],
    onUpdate: () => {
      emit("update:modelValue", editor.value.getHTML());
    },
  });
});
</script>

<template>
  <div class="tiptap-editor">
    <div class="menubar" v-if="editor">
      <div class="toolbar">
        <button
          @click="editor.chain().focus().toggleHeading({ level: 1 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }"
        >
          H1
        </button>

        <button
          @click="editor.chain().focus().toggleHeading({ level: 2 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }"
        >
          H2
        </button>

        <button
          @click="editor.chain().focus().toggleHeading({ level: 3 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }"
        >
          H3
        </button>
        <button
          @click="editor.chain().focus().toggleBulletList().run()"
          :class="{ 'is-active': editor.isActive('bulletList') }"
        >
          <icon name="ul" />
        </button>
        <button
          @click="editor.chain().focus().toggleOrderedList().run()"
          :class="{ 'is-active': editor.isActive('orderedList') }"
        >
          <icon name="ol" />
        </button>
        <button
          @click="editor.chain().focus().toggleBold().run()"
          :disabled="!editor.can().chain().focus().toggleBold().run()"
          :class="{ 'is-active': editor.isActive('bold') }"
        >
          <icon name="bold" />
        </button>
        <button
          @click="editor.chain().focus().toggleItalic().run()"
          :disabled="!editor.can().chain().focus().toggleItalic().run()"
          :class="{ 'is-active': editor.isActive('italic') }"
        >
          <icon name="italic" />
        </button>
        <button
          @click="editor.chain().focus().toggleStrike().run()"
          :disabled="!editor.can().chain().focus().toggleStrike().run()"
          :class="{ 'is-active': editor.isActive('strike') }"
        >
          <icon name="strike" />
        </button>

        <button
          @click="editor.chain().focus().toggleUnderline().run()"
          :disabled="!editor.can().chain().focus().toggleUnderline().run()"
          :class="{ 'is-active': editor.isActive('underline') }"
        >
          <icon name="underline" />
        </button>
        <button
          @click="editor.chain().focus().toggleCode().run()"
          :disabled="!editor.can().chain().focus().toggleCode().run()"
          :class="{ 'is-active': editor.isActive('code') }"
        >
          <icon name="code" />
        </button>
        <button
          @click="editor.chain().focus().setParagraph().run()"
          :class="{ 'is-active': editor.isActive('paragraph') }"
        >
          <icon name="paragraph" />
        </button>
        <button
          @click="editor.chain().focus().setParagraph().run()"
          :class="{ 'is-active': editor.isActive('paragraph') }"
        >
          <icon name="paragraph" />
        </button>
        <button
          @click="editor.chain().focus().toggleBlockquote().run()"
          :class="{ 'is-active': editor.isActive('blockquote') }"
        >
          <icon name="quote" />
        </button>
        <button
          @click="editor.chain().focus().toggleCodeBlock().run()"
          :class="{ 'is-active': editor.isActive('codeBlock') }"
        >
          <icon name="code_block" />
        </button>
        <button
          @click="
            editor
              .chain()
              .focus()
              .insertTable({ rows: 3, cols: 3, withHeaderRow: false })
              .run()
          "
        >
          <icon name="table" />
        </button>
        <span v-if="editor.isActive('table')">
          <button
            @click="editor.chain().focus().deleteTable().run()"
            :disabled="!editor.can().deleteTable()"
          >
            <icon name="delete_table" />
          </button>
          <button
            @click="editor.chain().focus().addColumnBefore().run()"
            :disabled="!editor.can().addColumnBefore()"
          >
            <icon name="add_col_before" />
          </button>
          <button
            @click="editor.chain().focus().addColumnAfter().run()"
            :disabled="!editor.can().addColumnAfter()"
          >
            <icon name="add_col_after" />
          </button>

          <button
            @click="editor.chain().focus().deleteColumn().run()"
            :disabled="!editor.can().deleteColumn()"
          >
            <icon name="delete_col" />
          </button>
          <button
            @click="editor.chain().focus().addRowBefore().run()"
            :disabled="!editor.can().addRowBefore()"
          >
            <icon name="add_row_before" />
          </button>
          <button
            @click="editor.chain().focus().addRowAfter().run()"
            :disabled="!editor.can().addRowAfter()"
          >
            <icon name="add_row_after" />
          </button>
          <button
            @click="editor.chain().focus().deleteRow().run()"
            :disabled="!editor.can().deleteRow()"
          >
            <icon name="delete_row" />
          </button>
          <button
            @click="editor.chain().focus().mergeCells().run()"
            :disabled="!editor.can().mergeCells()"
          >
            <icon name="combine_cells" /></button
        ></span>
      </div>
    </div>
    <editor-content :editor="editor" class="tiptap-editor__content" />
  </div>
</template>
<!-- @snippet_end -->
