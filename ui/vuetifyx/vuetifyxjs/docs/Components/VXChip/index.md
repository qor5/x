# vx-chip 徽章

## API

### Props

| Name    | Introduction | Type      | Default Value |
| ------- | ------------ | --------- | ------------- |
| presets | 预设尺寸     | `String`  | "badge"       |
| round   | 是否使用圆角 | `Boolean` | false         |

> 除此之外所有的 [v-chip](https://vuetifyjs.com/en/api/v-chip/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Slots

#### v-slot:default

默认内容区域插槽

#### v-slot:prepend

按钮内文字前图标区域

#### v-slot:append

按钮内文字前图标区域

### 颜色和圆角

:::demo

```vue
<script setup lang="ts"></script>
<template>
  <v-row>
    <v-col cols="2" class="text-center">
      <div class="mb-2">primary</div>
      <vx-chip>badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round>badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip disabled>badge</vx-chip>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">secondary</div>
      <vx-chip color="secondary">badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round color="secondary">badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip color="secondary" disabled>badge</vx-chip>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">info</div>
      <vx-chip color="info">badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round color="info">badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip color="info" disabled>badge</vx-chip>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">success</div>
      <vx-chip color="success">badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round color="success">badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip color="success" disabled>badge</vx-chip>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">warning</div>
      <vx-chip color="warning">badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round color="warning">badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip color="warning" disabled>badge</vx-chip>
    </v-col>

    <v-col cols="2" class="text-center">
      <div class="mb-2">error</div>
      <vx-chip color="error">badge</vx-chip>

      <div class="my-2">(round)</div>
      <vx-chip round color="error">badge</vx-chip>

      <div class="my-2">(disabled)</div>
      <vx-chip color="error" disabled>badge</vx-chip>
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
      <vx-chip variant="elevated" color="primary">Badge</vx-chip>
      <vx-chip variant="elevated" color="secondary">Badge</vx-chip>
      <vx-chip variant="elevated" color="info">Badge</vx-chip>
      <vx-chip variant="elevated" color="success">Badge</vx-chip>
      <vx-chip variant="elevated" color="warning">Badge</vx-chip>
      <vx-chip variant="elevated" color="error">Badge</vx-chip>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">tonal</div>
      <vx-chip variant="tonal" color="primary">Badge</vx-chip>
      <vx-chip variant="tonal" color="secondary">Badge</vx-chip>
      <vx-chip variant="tonal" color="info">Badge</vx-chip>
      <vx-chip variant="tonal" color="success">Badge</vx-chip>
      <vx-chip variant="tonal" color="warning">Badge</vx-chip>
      <vx-chip variant="tonal" color="error">Badge</vx-chip>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">outlined</div>
      <vx-chip variant="outlined" color="primary">Badge</vx-chip>
      <vx-chip variant="outlined" color="secondary">Badge</vx-chip>
      <vx-chip variant="outlined" color="info">Badge</vx-chip>
      <vx-chip variant="outlined" color="success">Badge</vx-chip>
      <vx-chip variant="outlined" color="warning">Badge</vx-chip>
      <vx-chip variant="outlined" color="error">Badge</vx-chip>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">flat</div>
      <vx-chip variant="flat" color="primary">Badge</vx-chip>
      <vx-chip variant="flat" color="secondary">Badge</vx-chip>
      <vx-chip variant="flat" color="info">Badge</vx-chip>
      <vx-chip variant="flat" color="success">Badge</vx-chip>
      <vx-chip variant="flat" color="warning">Badge</vx-chip>
      <vx-chip variant="flat" color="error">Badge</vx-chip>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">plain</div>
      <vx-chip variant="plain" color="primary">Badge</vx-chip>
      <vx-chip variant="plain" color="secondary">Badge</vx-chip>
      <vx-chip variant="plain" color="info">Badge</vx-chip>
      <vx-chip variant="plain" color="success">Badge</vx-chip>
      <vx-chip variant="plain" color="warning">Badge</vx-chip>
      <vx-chip variant="plain" color="error">Badge</vx-chip>
    </v-col>
    <v-col cols="2" class="text-center">
      <div class="my-2">text</div>
      <vx-chip variant="plain" color="primary">text</vx-chip>
      <vx-chip variant="plain" color="secondary">text</vx-chip>
      <vx-chip variant="plain" color="info">text</vx-chip>
      <vx-chip variant="plain" color="success">text</vx-chip>
      <vx-chip variant="plain" color="warning">text</vx-chip>
      <vx-chip variant="plain" color="error">text</vx-chip>
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
      <div class="mb-2">prepend-icon</div>
      <vx-chip prepend-icon="mdi-radiobox-marked" variant="tonal" color="success">Online</vx-chip>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">append-icon</div>
      <vx-chip append-icon="mdi-close" variant="tonal" color="">Offline</vx-chip>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">prepend-icon(round)</div>
      <vx-chip round prepend-icon="mdi-radiobox-marked" color="black">Button</vx-chip>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">prepend-icon(round)</div>
      <vx-chip round append-icon="mdi-radiobox-marked" color="black">Button</vx-chip>
    </v-col>
  </v-row>

  <v-row> <v-col>通过slot 实现</v-col></v-row>
  <v-row>
    <v-col cols="3" class="text-center">
      <div class="mb-2">prepend-icon</div>
      <vx-chip variant="tonal" color="success">
        <template #prepend><v-icon icon="mdi-radiobox-marked" class="mr-1" /></template>
        Online
      </vx-chip>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">append-icon</div>
      <vx-chip variant="tonal" color="success">
        <template #append><v-icon icon="mdi-radiobox-marked" class="ml-1" /></template>
        Online
      </vx-chip>
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">prepend-icon(round)</div>
      <vx-chip round color="black">
        <template #prepend><v-icon icon="mdi-radiobox-marked" class="mr-1" /></template>
        Button</vx-chip
      >
    </v-col>

    <v-col cols="3" class="text-center">
      <div class="mb-2">prepend-icon(round)</div>
      <vx-chip round color="black">
        <template #append><v-icon icon="mdi-radiobox-marked" class="ml-1" /></template>
        Button</vx-chip
      >
    </v-col>
  </v-row>
</template>
```

:::
