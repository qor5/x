# vx-btn 按钮

提供预设和样式微调的v-btn, 最重要的部份在presets，css根据presets进行了大量的微调

## API

### Props

| Name    | Introduction | Type                                           | Default Value |
| ------- | ------------ | ---------------------------------------------- | ------------- |
| presets | 预设尺寸     | `"default"`、`"x-small"`、`"small"`、`"large"` | "default"     |

> 除此之外所有的 [v-btn](https://vuetifyjs.com/en/api/v-btn/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Slots

#### v-slot:default

默认内容区域插槽

#### v-slot:prepend

按钮内文字前图标区域

#### v-slot:append

按钮内文字前图标区域

### 预设尺寸

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">x-small</div>
      <vx-btn presets="x-small">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn presets="x-small" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">small</div>
      <vx-btn presets="small">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn presets="small" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn>Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn disabled>Button</vx-btn>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn presets="large">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn presets="large" disabled>Button</vx-btn>
    </v-col>
  </v-row>
</template>
```

:::

### 颜色

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row>
    <v-col cols="2" class="text-center">
      <div class="mb-2">primary</div>
      <vx-btn color="primary">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="primary" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">secondary</div>
      <vx-btn color="secondary">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="secondary" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">info</div>
      <vx-btn color="info">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="info" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">success</div>
      <vx-btn color="success">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="success" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">warning</div>
      <vx-btn color="warning">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="warning" disabled>Button</vx-btn>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">error</div>
      <vx-btn color="error">Button</vx-btn>

      <div class="my-2">(disabled)</div>
      <vx-btn color="error" disabled>Button</vx-btn>
    </v-col>
  </v-row>
</template>
```

:::

### variant

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row>
    <v-col cols="2" class="text-center">
      <div class="my-2">elevated</div>
      <vx-btn variant="elevated" color="primary">Button</vx-btn>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">tonal</div>
      <vx-btn variant="tonal" color="primary">Button</vx-btn>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">outlined</div>
      <vx-btn variant="outlined" color="primary">Button</vx-btn>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">flat</div>
      <vx-btn variant="flat" color="primary">Button</vx-btn>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">plain</div>
      <vx-btn variant="plain" color="primary">Button</vx-btn>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">text</div>
      <vx-btn variant="plain" color="primary">text</vx-btn>
    </v-col>
  </v-row>
</template>
```

:::

### 带图标

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row> <v-col>通过props 实现</v-col></v-row>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">x-small</div>
      <vx-btn prepend-icon="mdi-refresh" append-icon="mdi-refresh" presets="x-small" color="black"
        >Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn
        class="mt-2"
        prepend-icon="mdi-refresh"
        append-icon="mdi-close"
        presets="x-small"
        color="black"
        stacked
        >Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">small</div>
      <vx-btn color="black" prepend-icon="mdi-refresh" append-icon="mdi-refresh" presets="small"
        >Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn
        class="mt-2"
        prepend-icon="mdi-refresh"
        append-icon="mdi-close"
        presets="small"
        color="black"
        stacked
        >Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn color="black" prepend-icon="mdi-refresh" append-icon="mdi-refresh">Button</vx-btn>
      <div class="my-2">(stacked)</div>
      <vx-btn class="mt-2" prepend-icon="mdi-refresh" append-icon="mdi-close" color="black" stacked
        >Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn color="black" prepend-icon="mdi-refresh" append-icon="mdi-refresh" presets="large"
        >Button</vx-btn
      >

      <div class="my-2">(stacked)</div>
      <vx-btn
        class="mt-2"
        prepend-icon="mdi-refresh"
        append-icon="mdi-close"
        presets="large"
        color="black"
        stacked
        >Button</vx-btn
      >
    </v-col>
  </v-row>

  <v-row> <v-col>通过slot 实现</v-col></v-row>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">x-small</div>
      <vx-btn presets="x-small" color="black">
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn presets="x-small" color="black" stacked>
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">small</div>
      <vx-btn color="black" presets="small">
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn presets="small" color="black" stacked>
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn color="black">
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn color="black" stacked>
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn color="black" presets="large">
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
      <div class="my-2">(stacked)</div>
      <vx-btn presets="large" color="black" stacked>
        <template #prepend><v-icon icon="mdi-refresh" /></template>
        <template #append><v-icon icon="mdi-refresh" /></template>
        Button</vx-btn
      >
    </v-col>
  </v-row>
</template>
```

:::

### 纯图标

- 使用 `rounded` 可以使图标按钮变方

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row> <v-col>通过props 实现</v-col></v-row>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">x-small</div>
      <vx-btn presets="x-small" icon="mdi-plus-circle-outline" />
      <div class="my-2">(rounded)</div>
      <vx-btn presets="x-small" icon="mdi-plus-circle-outline" rounded />
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">small</div>

      <vx-btn icon="mdi-plus-circle-outline" presets="small" />
      <div class="my-2">(rounded)</div>
      <vx-btn icon="mdi-plus-circle-outline" rounded presets="small" />
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn icon="mdi-plus-circle-outline" />
      <div class="my-2">(rounded)</div>
      <vx-btn icon="mdi-plus-circle-outline" rounded />
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <vx-btn icon="mdi-plus-circle-outline" presets="large" />
      <div class="my-2">(rounded)</div>
      <vx-btn icon="mdi-plus-circle-outline" presets="large" rounded />
    </v-col>
  </v-row>
</template>
```

:::
