<template>
  <div class="tiptap-editor">

    <editor-menu-bar v-slot="{ commands, isActive }"
                     :editor="editor">
      <div class="menubar">
        <div class="toolbar">

          <button :class="{ 'is-active': isActive.heading({ level: 1 }) }"
                  class="menubar__button"
                  @click="commands.heading({ level: 1 })">
            H1
          </button>

          <button :class="{ 'is-active': isActive.heading({ level: 2 }) }"
                  class="menubar__button"
                  @click="commands.heading({ level: 2 })">
            H2
          </button>

          <button :class="{ 'is-active': isActive.heading({ level: 3 }) }"
                  class="menubar__button"
                  @click="commands.heading({ level: 3 })">
            H3
          </button>

          <button :class="{ 'is-active': isActive.bullet_list() }" class="menubar__button"
                  @click="commands.bullet_list">
            <icon name="ul"/>
          </button>

          <button :class="{ 'is-active': isActive.ordered_list() }"
                  class="menubar__button"
                  @click="commands.ordered_list">
            <icon name="ol"/>
          </button>

          <button :class="{ 'is-active': isActive.bold() }" class="menubar__button"
                  @click="commands.bold">
            <icon name="bold"/>
          </button>

          <button :class="{ 'is-active': isActive.italic() }" class="menubar__button"
                  @click="commands.italic">
            <icon name="italic"/>
          </button>

          <button :class="{ 'is-active': isActive.strike() }" class="menubar__button"
                  @click="commands.strike">
            <icon name="strike"/>
          </button>

          <button :class="{ 'is-active': isActive.underline() }" class="menubar__button"
                  @click="commands.underline">
            <icon name="underline"/>
          </button>

          <button :class="{ 'is-active': isActive.code() }" class="menubar__button"
                  @click="commands.code">
            <icon name="code"/>
          </button>

          <button :class="{ 'is-active': isActive.paragraph() }" class="menubar__button"
                  @click="commands.paragraph">
            <icon name="paragraph"/>
          </button>

          <button :class="{ 'is-active': isActive.blockquote() }" class="menubar__button"
                  @click="commands.blockquote">
            <icon name="quote"/>
          </button>

          <button :class="{ 'is-active': isActive.code_block() }" class="menubar__button"
                  @click="commands.code_block">
            <icon name="code_block"/>
          </button>

          <button
              class="menubar__button"
              @click="commands.createTable({rowsCount: 3, colsCount: 3, withHeaderRow: false })"
          >
            <icon name="table"/>
          </button>

          <span v-if="isActive.table()">
			<button
          class="menubar__button"
          @click="commands.deleteTable"
      >
				<icon name="delete_table"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.addColumnBefore"
      >
				<icon name="add_col_before"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.addColumnAfter"
      >
				<icon name="add_col_after"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.deleteColumn"
      >
				<icon name="delete_col"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.addRowBefore"
      >
				<icon name="add_row_before"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.addRowAfter"
      >
				<icon name="add_row_after"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.deleteRow"
      >
				<icon name="delete_row"/>
			</button>
			<button
          class="menubar__button"
          @click="commands.toggleCellMerge"
      >
				<icon name="combine_cells"/>
			</button>
		</span>
        </div>
      </div>
    </editor-menu-bar>

    <editor-content :editor="editor" class="tiptap-editor__content"/>
  </div>

</template>

<script>
import './sass/main.scss'
import {Editor, EditorContent, EditorMenuBar} from 'tiptap';
import Icon from './icon';
import {
  Blockquote,
  Bold,
  BulletList,
  Code,
  CodeBlock,
  HardBreak,
  Heading,
  History,
  Italic,
  Link,
  ListItem,
  OrderedList,
  Strike,
  Table,
  TableCell,
  TableHeader,
  TableRow,
  TodoItem,
  TodoList,
  Underline,
} from 'tiptap-extensions'

function extensions() {
  return [
    new Blockquote(),
    new BulletList(),
    new CodeBlock(),
    new HardBreak(),
    new Heading(),
    new ListItem(),
    new OrderedList(),
    new TodoItem(),
    new TodoList(),
    new Link(),
    new Table({
      resizable: true,
    }),
    new TableHeader(),
    new TableCell(),
    new TableRow(),
    new Bold(),
    new Code(),
    new Italic(),
    new Strike(),
    new Underline(),
    new History(),
  ]
}

// @snippet_begin(TipTapEditorVueComponent)
export default {

  components: {
    EditorContent,
    EditorMenuBar,
    Icon,
  },

  props: {
    value: String,
  },

  data() {
    return {
      editor: new Editor({
        content: this.$props.value,
        extensions: extensions(),
        onUpdate: ({getHTML}) => {
          const html = getHTML();
          this.$emit("input", html)
        },
      })
    }
  },

  beforeDestroy() {
    this.editor.destroy()
  }
}
// @snippet_end

</script>
