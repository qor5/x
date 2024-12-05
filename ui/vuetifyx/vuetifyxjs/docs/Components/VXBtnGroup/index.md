# vx-btn-group 按钮组

适配vx-btn 的 vx-btn-group

## API

### Props

| Name         | Introduction                                                | Type      | Default Value           |
| ------------ | ----------------------------------------------------------- | --------- | ----------------------- |
| divided      | 是否显示分割线                                              | `Boolean` | `false`                 |
| dividerColor | 分割线颜色，支持预设颜色和普通hex及rgba颜色,比如 primary 等 | `String`  | `"rgba(0, 0, 0, 0.12)"` |
| dividerWidth | 是否显示分割线                                              | `Number`  | `1`                     |

> 除此之外所有的 [v-btn](https://vuetifyjs.com/en/api/v-btn/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Slots

#### v-slot:default

默认内容区域插槽

### 预设尺寸

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">x-small</div>
      <!-- primary -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>
      <!-- secondary -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small" color="secondary">Button</vx-btn>

        <vx-btn presets="x-small" icon="mdi-menu-down" color="secondary" />
      </vx-btn-group>
      <!-- info -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small" color="info">Button</vx-btn>

        <vx-btn presets="x-small" icon="mdi-menu-down" color="info" />
      </vx-btn-group>
      <!-- success -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small" color="success">Button</vx-btn>

        <vx-btn presets="x-small" icon="mdi-menu-down" color="success" />
      </vx-btn-group>

      <!-- warning -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small" color="warning">Button</vx-btn>

        <vx-btn presets="x-small" icon="mdi-menu-down" color="warning" />
      </vx-btn-group>

      <!-- error -->
      <vx-btn-group presets="x-small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small" color="error">Button</vx-btn>

        <vx-btn presets="x-small" icon="mdi-menu-down" color="error" />
      </vx-btn-group>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">small</div>
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- secondary -->
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small" color="secondary">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" color="secondary" />
      </vx-btn-group>
      <!-- info -->
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small" color="info">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" color="info" />
      </vx-btn-group>
      <!-- success -->
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small" color="success">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" color="success" />
      </vx-btn-group>

      <!-- warning -->
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small" color="warning">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" color="warning" />
      </vx-btn-group>

      <!-- error -->
      <vx-btn-group presets="small" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small" color="error">Button</vx-btn>

        <vx-btn presets="small" icon="mdi-menu-down" color="error" />
      </vx-btn-group>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">default</div>
      <!-- 默认presets就是default -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn presets="default">Button</vx-btn>

        <vx-btn presets="default" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- secondary -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn color="secondary">Button</vx-btn>

        <vx-btn icon="mdi-menu-down" color="secondary" />
      </vx-btn-group>
      <!-- info -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn color="info">Button</vx-btn>

        <vx-btn icon="mdi-menu-down" color="info" />
      </vx-btn-group>
      <!-- success -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn color="success">Button</vx-btn>

        <vx-btn icon="mdi-menu-down" color="success" />
      </vx-btn-group>

      <!-- warning -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn color="warning">Button</vx-btn>

        <vx-btn icon="mdi-menu-down" color="warning" />
      </vx-btn-group>

      <!-- error -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn color="error">Button</vx-btn>

        <vx-btn icon="mdi-menu-down" color="error" />
      </vx-btn-group>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">large</div>
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- secondary -->
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large" color="secondary">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" color="secondary" />
      </vx-btn-group>
      <!-- info -->
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large" color="info">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" color="info" />
      </vx-btn-group>
      <!-- success -->
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large" color="success">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" color="success" />
      </vx-btn-group>

      <!-- warning -->
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large" color="warning">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" color="warning" />
      </vx-btn-group>

      <!-- error -->
      <vx-btn-group presets="large" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large" color="error">Button</vx-btn>

        <vx-btn presets="large" icon="mdi-menu-down" color="error" />
      </vx-btn-group>
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
    <v-col cols="3" class="text-center">
      <div class="mb-2">elevated</div>
      <!-- x-small -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group class="my-1" divided divider-color="#fff">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
    <v-col cols="3" class="text-center">
      <div class="mb-2">tonal</div>

      <!-- x-small -->
      <vx-btn-group variant="tonal" class="my-1" divided divider-color="#fff">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group variant="tonal" class="my-1" divided divider-color="#fff">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group variant="tonal" class="my-1" divided divider-color="#fff">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group variant="tonal" class="my-1" divided divider-color="#fff">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
    <v-col cols="3" class="text-center">
      <div class="mb-2">outlined</div>
      <!-- x-small -->
      <vx-btn-group variant="outlined" class="my-1" divided divider-color="primary">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group variant="outlined" class="my-1" divided divider-color="primary">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group variant="outlined" class="my-1" divided divider-color="primary">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group variant="outlined" class="my-1" divided divider-color="primary">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
    <v-col cols="3" class="text-center">
      <div class="mb-2">plain</div>
      <!-- x-small -->
      <vx-btn-group variant="plain" class="my-1">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group variant="plain" class="my-1">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group variant="plain" class="my-1">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group variant="plain" class="my-1">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
    <v-col cols="3" class="text-center">
      <div class="mb-2">text</div>
      <!-- x-small -->
      <vx-btn-group variant="text" class="my-1">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group variant="text" class="my-1">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group variant="text" class="my-1">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group variant="text" class="my-1">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
    <v-col cols="3" class="text-center">
      <div class="mb-2">flat</div>
      <!-- x-small -->
      <vx-btn-group variant="flat" class="my-1">
        <vx-btn presets="x-small">Button</vx-btn>
        <vx-btn presets="x-small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- small -->
      <vx-btn-group variant="flat" class="my-1">
        <vx-btn presets="small">Button</vx-btn>
        <vx-btn presets="small" icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- default -->
      <vx-btn-group variant="flat" class="my-1">
        <vx-btn>Button</vx-btn>
        <vx-btn icon="mdi-menu-down" />
      </vx-btn-group>

      <!-- large -->
      <vx-btn-group variant="flat" class="my-1">
        <vx-btn presets="large">Button</vx-btn>
        <vx-btn presets="large" icon="mdi-menu-down" />
      </vx-btn-group>
    </v-col>
  </v-row>
</template>
```

:::
