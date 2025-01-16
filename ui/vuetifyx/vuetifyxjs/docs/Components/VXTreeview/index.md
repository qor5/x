# vx-treeview

树状视图组件

## API

### Props

| Name  | Introduction           | Type    | Default Value |
| ----- | ---------------------- | ------- | ------------- |
| items | 树状视图项目组成的数组 | `Array` | undefined     |

> 除此之外所有的 [v-treeview](https://vuetifyjs.com/en/api/v-treeview/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Slots

#### v-slot:default

默认内容区域插槽

#### v-slot:prepend

前置插槽

#### v-slot:divider

分割线插槽

### Slot Usage

:::demo

```vue
<script setup lang="ts">
const linkUsageItems = [
  {
    id: 1,
    title: 'Node 1 :'
  },
  {
    id: 5,
    title: 'Node 5 :'
  },

  {
    id: 19,
    title: 'Node 19 :',
    children: [
      {
        id: 20,
        title: 'Node 20 :',
        children: [
          { id: 21, title: 'Node 21 :' },
          { id: 22, title: 'Node 22 :' },
          { id: 23, title: 'Node 23 :' }
        ]
      },
      { id: 24, title: 'Node 24 :' },
      { id: 25, title: 'Node 25 :' }
    ]
  }
]
</script>
<template>
  <v-row>
    <v-col cols="6">
      <vx-treeview :items="linkUsageItems">
        <template v-slot:prepend="{ item, isOpen }">
          <v-icon icon="mdi-plus-circle-outline" size="small" color="grey-darken-3" />
        </template>
      </vx-treeview>
    </v-col>
  </v-row>
</template>
```

:::
