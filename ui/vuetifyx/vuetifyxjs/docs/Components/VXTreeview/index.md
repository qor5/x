# VXTreeview

这是一个基本示例，你可以用 `markdown` 语法 和 `vue3`、`vuetify` 在此处写任何组件代码


## vuetifyx 组件开发流程

### 1.新建组件
在 `qor5/x/vuetifyx/src/lib` 里新建任何 vue 组件, 比如 `qor5/x/vuetifyx/src/lib/richEditor/index.vue`

### 2.注册组件
在 `qor5/x/vuetifyx/src/lib/plugins/index.vue` 里注册组件，这样注册可以确保生产环境和本地环境都可用

1. vuetify 组件不用注册，直接用就行，比如 `v-btn`
2. vuetifyx 组件请以 `vx-` 开头

```ts
// qor5/x/vuetifyx/src/lib/plugins/index.vue

import TextField from '@/lib/Form/TextFiled.vue'

const vuetifyx = {
  install: (app: App) => {
    app.component('vx-datepicker', Datepicker)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-filter', Filter)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-textdatepicker', TextDatepicker)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-messagelistener', MessageListener)
    app.component('vx-overlay', Overlay)
    app.component('vx-text-field', TextField)
    // 在此注册你的新组件
    // app.component('vx-rich-editor', TextField)
  }
}
```

### 3. 直接在当前文档使用

可以在当前 markdown 里倒入任何本地局部示例组件，比如

`import VueJsonPretty from 'vue-json-pretty'`

:::demo

```vue
<script setup lang="ts">
import VueJsonPretty from 'vue-json-pretty'

</script>
<template>
  <v-btn color="primary">hello world</v-btn>

  <VueJsonPretty :data="value"></VueJsonPretty>
  你可以任意更改这里的代码
</template>
```

<style scoped></style>
:::

### 4. 为组件撰写必要说明和参数
目前先随意，后期会有规范
