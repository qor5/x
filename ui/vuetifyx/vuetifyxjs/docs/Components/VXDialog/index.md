# vx-dialog 弹窗

## 唤起弹窗方式

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">使用变量控制</div>
        <vx-dialog
          v-model="dialogVisible"
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        />
        <v-btn color="primary" @click="dialogVisible = true">Open Dialog</v-btn>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">使用slot</div>
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

## 五种预设样式

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="3" class="text-center">
        <div class="mb-2">普通样式</div>
        <vx-dialog
          title="Confirm"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="secondary">普通Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">提示样式</div>
        <vx-dialog
          title="Info"
          type="info"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="primary">提示Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">成功样式</div>
        <vx-dialog
          title="Success"
          type="success"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="success">提示Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">警告样式</div>
        <vx-dialog
          title="Warning"
          type="warn"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="warning">警告Dialog</v-btn>
          </template>
        </vx-dialog>
      </v-col>

      <v-col cols="3" class="text-center">
        <div class="mb-2">错误样式</div>
        <vx-dialog
          title="Error"
          type="error"
          text="This is an info description line This is an info description lineThis is an info description lineThis is an info description lineThis is an info description line"
        >
          <template v-slot:activator="{ props: { activatorProps } }">
            <v-btn v-bind="activatorProps" color="error">错误Dialog</v-btn>
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

## 不同尺寸

- 默认尺寸是 `size: "default"`, 此时固定宽度，用作大多数场景的及时反馈
- 大尺寸是 `size: "large"`, 此时最大宽度 665px，比较适合一些操作反馈类的弹窗
- 除此之外，还可使用 `width` 和 `maxWidth`, `contentHeight` 自定义宽度和内容高度

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

        <v-col cols="3" class="text-center">
        <div class="mb-2">隐藏按钮</div>
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
