# vx-breadcrumbs

面包屑组件

## API

### Props

| Name    | Introduction         | Type     | Default Value |
| ------- | -------------------- | -------- | ------------- |
| items   | 面包屑项目组成的数组 | `Array`  | undefined     |
| divider | 分割线类型           | `String` | "»"           |

> 除此之外所有的 [v-breadcrumbs](https://vuetifyjs.com/en/api/v-breadcrumbs/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Slots

#### v-slot:default

默认内容区域插槽

#### v-slot:divider

分割线插槽

#### v-slot:prepend

前置插槽

#### v-slot:title

标题插槽

### Simplest Usage (纯文本)

:::demo

```vue
<template>
  <v-row>
    <v-col cols="12">
      <vx-breadcrumbs :items="['Link_1', 'Link_2', 'Link_3']" />
    </v-col>
  </v-row>
</template>
```

:::

### Link Usage (链接用法)

:::demo

```vue
<script setup lang="ts">
const linkUsageItems = [
  {
    title: 'Link_1',
    disabled: false,
    href: 'breadcrumbs_dashboard'
  },
  {
    title: 'Link_2',
    disabled: false,
    href: 'breadcrumbs_link_1'
  },
  {
    title: 'Link_3',
    disabled: true,
    href: 'breadcrumbs_link_2'
  }
]
</script>
<template>
  <v-row>
    <v-col cols="12">
      <vx-breadcrumbs :items="linkUsageItems" />
    </v-col>
  </v-row>
</template>
```

:::

### Slot Usage

:::demo

```vue
<script setup lang="ts">
const linkUsageItems = [
  {
    title: 'Link_1',
    disabled: false,
    href: 'breadcrumbs_dashboard'
  },
  {
    title: 'Link_2',
    disabled: false,
    href: 'breadcrumbs_link_1'
  },
  {
    title: 'Link_3',
    disabled: true,
    href: 'breadcrumbs_link_2'
  }
]
</script>
<template>
  <v-row>
    <v-col cols="6">
      <div class="text-caption">分隔符插槽</div>
      <vx-breadcrumbs :items="linkUsageItems">
        <template v-slot:divider>
          <v-icon icon="mdi-chevron-right"></v-icon>
        </template>
      </vx-breadcrumbs>
    </v-col>

    <v-col cols="6">
      <div class="text-caption">前置插槽</div>
      <vx-breadcrumbs :items="linkUsageItems">
        <template v-slot:prepend>
          <v-icon icon="$vuetify" size="small"></v-icon>
        </template>
      </vx-breadcrumbs>
    </v-col>

    <v-col cols="6">
      <div class="text-caption">title插槽</div>
      <vx-breadcrumbs :items="linkUsageItems">
        <template v-slot:title="{ item, index }">
          <vx-chip>{{ item.title }}</vx-chip></template
        >
      </vx-breadcrumbs>
    </v-col>

    <v-col cols="6">
      <div class="text-caption">默认插槽</div>
      <vx-breadcrumbs>
        <v-breadcrumbs-item>Link_1</v-breadcrumbs-item>
        <v-breadcrumbs-divider> - </v-breadcrumbs-divider>
        <v-breadcrumbs-item>Link_2</v-breadcrumbs-item>
        <v-breadcrumbs-divider> - </v-breadcrumbs-divider>
        <v-breadcrumbs-item>Link_3</v-breadcrumbs-item>
      </vx-breadcrumbs>
    </v-col>
  </v-row>
</template>
```

:::
