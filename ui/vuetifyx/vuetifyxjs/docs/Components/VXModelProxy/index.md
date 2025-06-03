# VXModelProxy

VXModelProxy 是一个代理组件，用于代理绑定到该组件的事件和方法，同时提供 JSON 字符串与对象之间的转换功能。

## API

### Props

| 名称        | 介绍                                         | 类型     | 默认值 |
| ----------- | -------------------------------------------- | -------- | ------ |
| modelValue  | 绑定的模型值，使用 v-model 双向绑定          | `String` | -      |
| formatModel | 模型值格式化方法，目前仅支持 'jsonStringify' | `String` | -      |

### Events

| 名称              | 载荷     | 介绍                                                                     |
| ----------------- | -------- | ------------------------------------------------------------------------ |
| update:modelValue | `String` | 当内部组件的 v-model 值更新时触发，会根据 formatModel 属性进行适当的转换 |

## 代理复杂组件

VXModelProxy 可以代理任何支持 v-model 的组件，包括复杂组件如 VXSegmentForm。

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

// 假设这是从数据库获取的已保存的过滤器配置（JSON字符串）
const savedFilter = ref(
  '{"demographics":[{"id":"user_gender","view":{"operator":"EQ","value":"MALE"}}]}'
)

// 简化版的过滤器选项
const options = ref([
  {
    id: 'demographics',
    name: '人口统计',
    description: '人口统计过滤器',
    builders: [
      {
        id: 'user_gender',
        name: '用户性别',
        description: '按性别过滤用户',
        categoryID: 'demographics',
        view: {
          fragments: [
            {
              defaultValue: 'EQ',
              key: 'operator',
              multiple: false,
              options: [
                { label: '等于', value: 'EQ' },
                { label: '不等于', value: 'NE' }
              ],
              required: true,
              type: 'SELECT'
            },
            {
              defaultValue: null,
              key: 'value',
              multiple: false,
              options: [
                { label: '男', value: 'MALE' },
                { label: '女', value: 'FEMALE' },
                { label: '其他', value: 'OTHER' }
              ],
              required: true,
              type: 'SELECT'
            }
          ]
        }
      }
    ]
  }
])
</script>

<template>
  <div>
    <h3>当前编辑的 JSON 字符串</h3>
    <pre style="max-height: 150px; overflow: auto">{{ savedFilter }}</pre>

    <v-divider class="my-4"></v-divider>

    <h3>VXSegmentForm 使用 VXModelProxy</h3>
    <vx-model-proxy
      v-model="savedFilter"
      format-model="jsonStringify"
      v-slot="{ modelValue, 'onUpdate:modelValue': updateModelValue }"
    >
      <vx-segment-form
        :model-value="modelValue"
        :options="options"
        @update:model-value="updateModelValue"
      />
    </vx-model-proxy>
  </div>
</template>
```

:::
