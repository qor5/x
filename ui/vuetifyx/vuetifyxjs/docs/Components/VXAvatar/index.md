# vx-avatar 头像

带有容错处理的头像组件，当图片加载失败默认会fallback到名字显示模式

## API

### Props

| Name | Introduction                   | Type     | Default Value |
| ---- | ------------------------------ | -------- | ------------- |
| name | avatar要展示的文字（取首字符） | `String` | -             |
| img  | 要展示的头像url                | `String` | -             |
| size | 预设尺寸，见示例               | `String` | `default`     |

## Size

- 5 中预设：`x-small`、`small`、`default`、`large`、`x-large`
- 也可以自定义大小传入数值

:::demo

```vue
<template>
  <h1 class="mb-2"><b>预设尺寸:</b></h1>
  <v-row>
    <v-col cols="2 text-center">
      <div class="mb-2">x-small</div>
      <vx-avatar name="ShaoXing" size="x-small" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">small</div>
      <vx-avatar name="ShaoXing" size="small" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">default</div>
      <vx-avatar name="ShaoXing" size="default" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">large</div>
      <vx-avatar name="ShaoXing" size="large" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">x-large</div>
      <vx-avatar name="ShaoXing" size="x-large" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
  </v-row>

  <h1 class="mb-2"><b>自定义尺寸:</b></h1>
  <v-row>
    <v-col cols="2 text-center">
      <div class="mb-2">16px</div>
      <vx-avatar class="mb-1" name="ShaoXing" size="16" />
      <vx-avatar name="ShaoXing" size="16" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">24px</div>
      <vx-avatar class="mb-1" name="ShaoXing" size="24" />
      <vx-avatar name="ShaoXing" size="24" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">32px</div>
      <vx-avatar class="mb-1" name="ShaoXing" size="32" />
      <vx-avatar name="ShaoXing" size="32" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">40px</div>
      <vx-avatar class="mb-1" name="ShaoXing" size="40" />
      <vx-avatar name="ShaoXing" size="40" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">48px</div>
      <vx-avatar class="mb-2" name="ShaoXing" size="48" />
      <vx-avatar name="ShaoXing" size="48" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
  </v-row>
  <v-row>
    <v-col cols="2 text-center">
      <div class="mb-2">64px</div>
      <vx-avatar class="mb-2" name="ShaoXing" size="64" />
      <vx-avatar name="ShaoXing" size="64" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">80px</div>
      <vx-avatar class="mb-4" name="ShaoXing" size="80" />
      <vx-avatar name="ShaoXing" size="80" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">96px</div>
      <vx-avatar class="mb-4" name="ShaoXing" size="96" />
      <vx-avatar name="ShaoXing" size="96" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">128px</div>
      <vx-avatar class="mb-4" name="ShaoXing" size="128" />
      <vx-avatar name="ShaoXing" size="128" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
    <v-col cols="2 text-center">
      <div class="mb-2">160px</div>
      <vx-avatar class="mb-4" name="ShaoXing" size="160" />
      <vx-avatar name="ShaoXing" size="160" img="/x/imgs/vx-avatar-example.png" />
    </v-col>
  </v-row>
</template>
```

:::
