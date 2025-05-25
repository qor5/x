# vx-tiptap-editor 富文本编辑器

## 基本用法

:::demo

```vue
<template>
  <VApp id="app">
    <VContainer>
      <VLocaleProvider locale="zhHans">
        <div class="border-thin">
          <vx-tiptap-editor
            v-model="content"
            label=""
            :min-height="200"
            :max-height="365"
            :hide-bubble="true"
            :extensions="extensions"
            :disabled="false"
            :readonly="false"
          />
        </div>
        <p class="mt-4">{{ content }}</p>
      </VLocaleProvider>
    </VContainer>
  </VApp>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const extensions = ref([
  // {
  //   name: 'BaseKit',
  //   options: {
  //     placeholder: {
  //       placeholder: 'Enter some text...'
  //     }
  //   }
  // },
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
  // { name: 'SubAndSuperScript', options: { divider: true } },
  { name: 'BulletList' },
  { name: 'OrderedList', options: { divider: true } },
  // { name: 'TaskList' },
  { name: 'Indent', options: { divider: true } },
  { name: 'Link', options: { divider: true } },
  { name: 'Image' },
  // { name: 'ImageGlue', options: {
  //   onClick:({editor, value, window}:{editor: any, value:any, window: any}) => {
  //     console.log('ImageGlue clicked with editor:', editor);
  //     console.log('ImageGlue clicked with value:', value);
  //     console.log('ImageGlue clicked with window:', window);
  //   },
  // } },
  { name: 'Video', options: { divider: true } },
  // { name: 'Table', options: { divider: true } },
  { name: 'Blockquote' },
  { name: 'HorizontalRule' },
  { name: 'CodeBlock', options: { divider: true } },
  {
    name: 'HtmlView',
    options: {
      divider: true,
      allowedAttributes: ['class', 'style', 'id', 'data-abc']
    }
  },
  { name: 'Clear' },
  { name: 'History', options: { divider: true } }
])

const content = ref(`<h2>
            Hi there,
          </h2>
          <p>
            this is a <em>basic</em> example of <strong>Tiptap</strong>. Sure, there are all kind of basic text styles you'd probably expect from a text editor. But wait until you see the lists:
          </p>
          <ul>
            <li>
              That's a bullet list with one …
            </li>
            <li>
              … or two list items.
            </li>
          </ul>
          <p>
            Isn't that great? And all of that is editable. But wait, there's more. Let's try a code block:
          </p>
          <pre><code class="language-css">body {
    display: none;
  }</code></pre>
          <p>
            I know, I know, this is impressive. It's only the tip of the iceberg though. Give it a try and click a little bit around. Don't forget to check the other examples too.
          </p>
          <blockquote>
            Wow, that's amazing. Good work, boy! 👏
            <br />
            — Mom
          </blockquote>`)
</script>
```

:::

## 属性继承机制

编辑器现在支持统一的属性管理机制：

### 1. 全局属性配置

通过 `HtmlView` 扩展的 `allowedAttributes` 选项，可以全局控制所有支持的扩展允许哪些 HTML 属性：

```javascript
{
  name: 'HtmlView',
  options: {
    allowedAttributes: ['class', 'style', 'id', 'data-testid', 'title', 'aria-label']
  }
}
```

### 2. 自动继承

支持属性继承的扩展（如 `Heading`、`Blockquote`）会自动从全局配置中继承 `allowedAttributes`，无需单独配置。

### 3. 局部覆盖

如果需要为特定扩展设置不同的属性配置，可以在扩展的选项中单独指定：

```javascript
{
  name: 'Heading',
  options: {
    allowedAttributes: ['class', 'id'] // 覆盖全局配置
  }
}
```

### 4. 支持的扩展

目前支持属性继承的扩展包括：

- `Heading` - 标题元素
- `Blockquote` - 引用块元素
- 更多扩展正在逐步支持中...

### 5. 使用示例

1. 在富文本模式下创建标题或引用块
2. 切换到 HTML 视图模式
3. 手动添加配置的属性（如 `class`、`style`、`data-testid` 等）
4. 切换回富文本模式 - 属性会被保留
