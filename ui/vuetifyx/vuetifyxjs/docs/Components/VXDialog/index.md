# vx-dialog 弹窗

带样式预设的通用弹窗，即时反馈类的仅需配置较少参数即可使用，也支持复杂场景的控制

## API

### Props

| Name             | Introduction                                                | Type                                                            | Default Value |
| ---------------- | ----------------------------------------------------------- | --------------------------------------------------------------- | ------------- |
| title            | 弹窗标题                                                    | `String`                                                        | -             |
| text             | 弹窗内容，适合纯文本                                        | `String`                                                        | -             |
| size             | [弹窗尺寸](./#弹窗尺寸), 默认尺寸定宽，large 尺寸定最大宽度 | `'default'`｜ `'large'`                                         | `'default'`   |
| type             | 弹窗[预设样式](./#预设样式)                                 | `'default'` ｜ `'info'` ｜ `'success'` ｜ `'warn'` ｜ `'error'` | `'default'`   |
| okText           | 确认按钮文案                                                | `String`                                                        | `OK`          |
| cancelText       | 取消按钮文案                                                | `String`                                                        | `Cancel`      |
| width            | 弹窗宽度                                                    | `Number`                                                        | -             |
| maxWidth         | 弹窗最大宽度                                                | `Number`                                                        | -             |
| contentPadding   | 弹窗内的间隙padding（请使用css的样式修改）                  | `String`                                                        | -             |
| contentOnlyMode  | 该模式可以自定义弹窗内所有的内容                            | `Boolean`                                                       | false         |
| contentHeight    | 弹窗内容高度                                                | `Number`                                                        | -             |
| disableOk        | 禁用确认按钮                                                | `Boolean`                                                       | `false`       |
| loadingOk        | 确认按钮loading                                             | `Boolean`                                                       | `false`       |
| hideOk           | 隐藏确认按钮                                                | `Boolean`                                                       | `false`       |
| hideCancel       | 隐藏取消按钮                                                | `Boolean`                                                       | `false`       |
| hideClose        | 隐藏右上角关闭按钮                                          | `Boolean`                                                       | `false`       |
| hideFooter       | 隐藏底部操作按钮区域                                        | `Boolean`                                                       | `false`       |
| noClickAnimation | 取消点击弹窗外的弹性动效                                    | `Boolean`                                                       | `false`       |
| model-value      | 控制弹窗显示与否                                            | `Boolean`                                                       | `false`       |

> 除此之外所有的 [v-dialog](https://vuetifyjs.com/en/api/v-dialog/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Events

| Name              | Payload                                              | Introduction                                                                                                                      |
| ----------------- | ---------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| click:ok          | `{ isActive: Ref<isActive>, isLoading: Ref<false> }` | 点击 OK 按钮时触发，使用此事件回调需要[手动关闭弹窗](./#按钮及事件回调) , 可以通过回调函数接收 isLoading 控制 okbtn 的loading状态 |
| click:cancel      | `{ isActive: Ref<isActive> }`                        | 点击 Cancel 按钮时触发，使用此事件回调需要[手动关闭弹窗](./#按钮及事件回调)                                                       |
| click:close       | `{ isActive: Ref<isActive> }`                        | 点击右上角关闭图标时触发，使用此事件回调需要[手动关闭弹窗](./#按钮及事件回调)                                                     |
| click:outside     | `PointerEvent`                                       | 点击弹窗以外的区域触发                                                                                                            |
| update:modelValue | `boolean`                                            | 弹窗 model 值改变时触发                                                                                                           |

### Slots

#### v-slot:activator

[激活插槽](./#激活弹窗)，使用该插槽不需要绑定 `v-model`

##### scope value

```js
{
  isActive: boolean,
  props: { activatorProps:Record<string, any> }
}
```

#### v-slot:default

默认内容区域插槽

##### scope value

```js
{ isActive: Ref<boolean> }
```

#### v-slot:action-btn

底部按钮区域插槽

##### scope value

```js
{ isActive: Ref<boolean> }
```

## 激活弹窗

在以下两个极简的示例里，我们可以使用两种方法来唤起弹窗。

- 1. 使用 `v-model` 去绑定值，并通过一个变量去驱动它
- 2. 使用 **activator** slot 去渲染按钮，并给按钮 `v-bind` **props** 对象的 `activatorProps`

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">1.使用v-model</div>
        <vx-dialog
          v-model="dialogVisible"
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        />
        <v-btn color="primary" @click="dialogVisible = true">Open Dialog</v-btn>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">2.使用 activator slot</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const dialogVisible = ref(false)
</script>
```

:::

## 预设样式

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">default</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">info</div>
        <vx-dialog
          title="Info"
          type="info"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="primary">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">success</div>
        <vx-dialog
          title="Success"
          type="success"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="success">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">warn</div>
        <vx-dialog
          title="Warning"
          type="warn"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="warning">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">error</div>
        <vx-dialog
          title="Error"
          type="error"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="error">Open Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const dialogVisible = ref(false)
</script>
```

:::

## 弹窗尺寸

- 预设尺寸：
  - 默认尺寸是 `size: "default"`, 此时**固定宽度**，用作大多数场景的及时反馈
  - 大尺寸是 `size: "large"`, 此时宽度靠内容撑开，最大宽度 665px，比较适合一些操作反馈类的弹窗

- 如果要自定义尺寸：还可使用 `width` 和 `maxWidth`, `contentHeight` 自定义宽度和内容高度

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">默认尺寸</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
        <div class="mt-2 text-caption">适用大部分场景，不需要指定size</div>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">大尺寸</div>
        <vx-dialog
          title="Confirm"
          size="large"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">自定义宽度</div>
        <vx-dialog
          title="Confirm"
          width="1200"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">内容滚动条（定高）</div>
        <vx-dialog
          title="Confirm"
          size="large"
          height="400"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">内容区域滚动条</div>
        <vx-dialog
          title="Confirm"
          width="300"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description lineThis is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
        <div class="mt-2 text-caption">
          内容区域超出高度会出现滚动条,如果不指定contentHeight则是以屏幕高度为准
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const dialogVisible = ref(false)
</script>
```

:::

## 只有内容区域的模式（contentOnlyMode）

- 当希望完全使用内部组件时，contentOnlyMode: true 配合使用 contentPadding 控制间距大小

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">contentOnlyMode</div>
        <vx-dialog contentOnlyMode contentPadding="0">
          This is an info description line This is an info description lineThis is an info
          description lineThis is an info description lineThis is an info description line

          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const dialogVisible = ref(false)
</script>
```

:::

## 按钮及事件回调

- 提供三个事件回调来控制弹窗的开启或者关闭 `click:ok`, `click:close`, `click:cancel`,
- 当传入自定义事件后，弹窗不再自动点击关闭，需要自己控制

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">自定义事件</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          @click:cancel="onCancel"
          @click:ok="onOK"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
        <div class="mt-2 text-caption">此弹窗关闭是自己控制的</div>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">自定义按钮文案</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          okText="保存"
          cancelText="取消"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const dialogVisible = ref(false)
function onOK({ isActive, isLoading }) {
  isLoading.value = true
  setTimeout(() => {
    isActive.value = false
    isLoading.value = false
  }, 1000)
}

function onCancel({ isActive }) {
  alert('not allowed to close dialog')
  // isActive.value = false
}
</script>
```

:::

## 区域隐藏/loading/禁用/显示

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">hideCancel</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          hide-cancel
          ok-text="Fine"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">hideOk</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          hide-ok
          ok-text="Fine"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">hideClose</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          hide-close
          ok-text="Fine"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">hideFooter</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          hideFooter
          ok-text="Fine"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">disableOk</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          disableOk
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">loadingOk</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description"
          loadingOk
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const dialogVisible = ref(false)
function onOK({ isActive, isLoading }) {
  isLoading.value = true
  setTimeout(() => {
    isActive.value = false
    isLoading.value = false
  }, 1000)
}

function onCancel({ isActive }) {
  alert('not allowed to close dialog')
  // isActive.value = false
}
</script>
```

:::
