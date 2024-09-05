# Input 输入框

## 基本用法

`v-model`对输入值做双向绑定，`placeholder`、`autofocus` 等原生 `input` 支持的属性会被自动继承。

:::demo

```vue
<template>
  <div>
    <h4>Default</h4>
    <p-input autofocus v-model="value" />
    <p>{{ value }}</p>
    <h4>disabled</h4>
    <p-input disabled placeholder="disabled" />
    <h4>placeholder</h4>
    <p-input placeholder="请输入用户名"
  /></div>
</template>
<script setup>
  import { ref } from 'vue'
  const value = ref('Default')
</script>
```

:::

## 图标插槽

可以使用 prefix 和 suffix 插槽

:::demo

```vue
<template>
  <div>
    <p-input placeholder="搜索关键词">
      <template #prefix>
        <p-icon :component="SearchSharp" size="16" />
      </template>
    </p-input>
    <br />
    <p-input placeholder="搜索关键词">
      <template #suffix>
        <p-icon :component="SearchSharp" size="16" />
      </template>
    </p-input>
    <br />
    <p-input placeholder="请输入信息">
      <template #suffix>
        <p-icon :component="ChatboxSharp" size="16" />
      </template>
    </p-input>
    <br />
    <p-input placeholder="百度一下 , 你就知道">
      <template #prefix>
        <p-icon :component="PawSharp" size="16" />
      </template>
    </p-input>
  </div>
</template>
<script setup>
  import { SearchSharp, ChatboxSharp, PawSharp } from '@vicons/ionicons5'
</script>
```

:::

## 复合型输入框

可以在输入框前后添加一个元素，通常是标签或按钮。

:::demo

```vue
<template>
  <div>
    <p-input v-model="value1">
      <template #prepend> https:// </template>
    </p-input>
    <br />
    <p-input>
      <template #prepend> www. </template>
    </p-input>
    <br />
    <p-input placeholder="请输入邮箱">
      <template #append>@qq.com</template>
    </p-input>
    <br />
    <p-input>
      <template #append>.com</template>
    </p-input>
    <br />
    <p-input placeholder="输入网址">
      <template #prepend>www.</template>
      <template #append>.com</template>
    </p-input>
  </div>
</template>
<script setup>
  import { ref } from 'vue'
  const value1 = ref('')
  const value2 = ref('')
  const value3 = ref('')
  const value4 = ref('')
  const value5 = ref('')
</script>
```

:::

## 密码框

使用 show-password 属性可得到一个可切换显示隐藏的密码框。

:::demo

```vue
<template>
  <div>
    <p-input placeholder="请输入密码" type="password" v-model="value" />
    <br />
    <p-input placeholder="请输入密码" showPassword v-model="value" />
  </div>
</template>
<script setup>
  import { ref } from 'vue'
  const value = ref('123456789')
</script>
```

:::

## 一键清空

使用 clearable 属性可得到一个可一键清空的输入框，使用 clear 事件可在清空按钮被点击时做一些操作。

:::demo

```vue
<template>
  <p-input placeholder="请输入" v-model="value" clearable />
</template>
<script setup>
  import { ref } from 'vue'
  const value = ref('')
</script>
```

:::

## 尺寸

:::demo

```vue
<template>
  <div>
    <h4>Small</h4>
    <p-input v-model="value1" size="sm" placeholder="请输入" />
    <h4>Middle</h4>
    <p-input v-model="value2" placeholder="请输入" />
    <h4>Large</h4>
    <p-input v-model="value3" size="lg" placeholder="请输入" />
  </div>
</template>
<script setup>
  import { ref } from 'vue'
  const value1 = ref('')
  const value2 = ref('')
  const value3 = ref('')
</script>
```

:::
