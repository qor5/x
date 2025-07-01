# vx-label 标签

提供多功能的标签组件，支持基本标签显示、工具提示、必填标记以及可折叠的内容展示。

## API

### Props

| Name             | Introduction          | Type                                                    | Default Value             |
| ---------------- | --------------------- | ------------------------------------------------------- | ------------------------- |
| tooltip          | 工具提示内容          | `string`                                                | -                         |
| tooltipLocation  | 工具提示位置          | `"top"` `"bottom"` `"left"` `"right"` `"start"` `"end"` | "end"                     |
| toggleLabel      | 是否启用可折叠标签    | `boolean`                                               | false                     |
| toggleIconSize   | 折叠图标尺寸          | `string`                                                | "default"                 |
| tooltipIconColor | 工具提示图标颜色      | `string`                                                | -                         |
| labelFor         | 标签关联的表单控件 ID | `string`                                                | -                         |
| icon             | 工具提示图标          | `string`                                                | "mdi-information-outline" |
| iconSize         | 图标尺寸              | `string`                                                | "small"                   |
| requiredSymbol   | 是否显示必填标记      | `boolean`                                               | false                     |

### Slots

#### v-slot:default

默认内容区域插槽，显示标签文本

#### v-slot:prepend

标签前置内容插槽

#### v-slot:toggle-content

可折叠内容插槽（仅在 `toggleLabel` 为 true 时可用）

## 基本用法

:::demo

```vue
<template>
  <div class="mb-4">
    <vx-label>简单标签</vx-label>
  </div>
</template>
```

:::

## 关联表单控件

:::demo

```vue
<template>
  <div>
    <vx-label labelFor="input1">关联表单的标签</vx-label>
    <vx-field
      id="input1"
      placeholder="输入内容"
      variant="outlined"
      density="compact"
      class="mt-2"
    />
  </div>
</template>
```

:::

## 必填标记

:::demo

```vue
<template>
  <div>
    <vx-label requiredSymbol>必填字段标签</vx-label>
    <vx-field placeholder="必填输入框" variant="outlined" density="compact" class="mt-2" />
  </div>
</template>
```

:::

## 基础工具提示

:::demo

```vue
<template>
  <div>
    <vx-label
      icon="mdi-help-circle-outline"
      icon-size="default"
      tooltip="这是一个简单的提示信息"
      tooltip-icon-color="primary"
    >
      带工具提示的标签
    </vx-label>
  </div>
</template>
```

:::

## 工具提示位置

:::demo

```vue
<template>
  <div class="mb-4">
    <vx-label icon="mdi-information-outline" tooltip="提示在顶部" tooltipLocation="top">
      顶部提示
    </vx-label>
  </div>

  <div class="mb-4">
    <vx-label icon="mdi-information-outline" tooltip="提示在右侧" tooltipLocation="right">
      右侧提示
    </vx-label>
  </div>

  <div class="mb-4">
    <vx-label icon="mdi-information-outline" tooltip="提示在左侧" tooltipLocation="left">
      左侧提示
    </vx-label>
  </div>
</template>
```

:::

## 格式化工具提示

:::demo

```vue
<template>
  <div>
    <vx-label
      icon="mdi-alert-circle-outline"
      :tooltip="formatTooltip"
      tooltip-icon-color="warning"
      tooltipLocation="right"
    >
      格式化 JSON 提示
    </vx-label>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const formatTooltip = ref(`{
  "defaultValue": {
    "value": "3426"
  },
  "valueType": "STRING"
}`)
</script>
```

:::

## 自定义样式

:::demo

```vue
<template>
  <div>
    <vx-label icon="mdi-palette-outline" icon-size="default">
      <span style="font-size: 20px; color: #e91e63; font-weight: bold;"> 自定义样式标签 </span>
    </vx-label>
  </div>
</template>
```

:::

## 混合内容

:::demo

```vue
<template>
  <div>
    <vx-label icon="mdi-star-outline" tooltip="带有多种元素的标签">
      <v-chip color="primary" size="small" class="mr-2">标签</v-chip>
      <span class="text-h6">混合内容标签</span>
    </vx-label>
  </div>
</template>
```

:::

## 基础折叠

:::demo

```vue
<template>
  <div>
    <vx-label toggleLabel>
      <template #prepend>
        <v-icon icon="mdi-folder-outline" class="mr-2" color="orange" />
      </template>
      简单可折叠标签
      <template #toggle-content>
        <v-card class="mt-3 pa-4" variant="outlined">
          <p>这是折叠内容区域</p>
          <p>可以放置任何内容，比如表单、列表、图片等</p>
        </v-card>
      </template>
    </vx-label>
  </div>
</template>
```

:::

## 可排序折叠列表

:::demo

```vue
<template>
  <div>
    <div v-for="(item, index) in sortableItems" :key="item.id" class="mb-3">
      <vx-label toggleLabel class="sortable-label">
        <template #prepend>
          <div class="section-sortable-area">
            <v-btn-group
              border
              density="comfortable"
              class="d-flex flex-column h-auto section-sortable-group"
              elevation="2"
            >
              <vx-btn
                icon="mdi-arrow-up"
                size="x-small"
                :class="['rounded-0', 'border-b', { 'btn-disabled': index === 0 }]"
                style="width:22px;height:30px;font-size:10px"
                :disabled="index === 0"
                @click="moveItemUp(index)"
              />
              <vx-btn
                icon="mdi-arrow-down"
                size="x-small"
                :class="['rounded-0', { 'btn-disabled': index === sortableItems.length - 1 }]"
                style="width:22px;height:30px;font-size:10px"
                :disabled="index === sortableItems.length - 1"
                @click="moveItemDown(index)"
              />
            </v-btn-group>
          </div>
        </template>
        {{ item.title }} ({{ index + 1 }}/{{ sortableItems.length }})
        <template #toggle-content>
          <div
            class="timeline-block pa-4 ma-3"
            style="background-color: #f5f5f5; border-radius: 8px;"
          >
            <h4>{{ item.title }}</h4>
            <p>{{ item.content }}</p>
          </div>
        </template>
      </vx-label>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const sortableItems = ref([
  {
    id: 1,
    title: '任务一',
    content: '这是第一个任务的详细描述，包含了任务的目标、要求和注意事项。'
  },
  {
    id: 2,
    title: '任务二',
    content: '这是第二个任务的详细描述，涉及到数据处理和分析工作。'
  },
  {
    id: 3,
    title: '任务三',
    content: '这是第三个任务的详细描述，主要负责界面设计和用户体验优化。'
  }
])

const moveItemUp = (index: number) => {
  if (index > 0) {
    const items = [...sortableItems.value]
    const temp = items[index]
    items[index] = items[index - 1]
    items[index - 1] = temp
    sortableItems.value = items
  }
}

const moveItemDown = (index: number) => {
  if (index < sortableItems.value.length - 1) {
    const items = [...sortableItems.value]
    const temp = items[index]
    items[index] = items[index + 1]
    items[index + 1] = temp
    sortableItems.value = items
  }
}
</script>

<style scoped>
.section-sortable-area {
  margin-right: 8px;
}

.section-sortable-group {
  border-radius: 4px;
}

.timeline-block {
  border-left: 3px solid #2196f3;
}

.btn-disabled {
  opacity: 0.3 !important;
  background-color: #f5f5f5 !important;
  color: #9e9e9e !important;
  cursor: not-allowed !important;
}

.btn-disabled:hover {
  opacity: 0.3 !important;
  background-color: #f5f5f5 !important;
}
</style>
```

:::

## 表单展开面板

:::demo

```vue
<template>
  <div>
    <vx-label toggleLabel>
      <template #prepend>
        <v-avatar size="24" class="mr-2">
          <v-icon icon="mdi-account" />
        </v-avatar>
      </template>
      用户信息展开面板
      <template #toggle-content>
        <v-card class="mt-3" variant="outlined">
          <v-card-text>
            <v-row>
              <v-col cols="6">
                <vx-field label="姓名" variant="outlined" density="compact" />
              </v-col>
              <v-col cols="6">
                <vx-field label="邮箱" variant="outlined" density="compact" />
              </v-col>
              <v-col cols="12">
                <vx-field label="描述" variant="outlined" density="compact" rows="3" />
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </template>
    </vx-label>
  </div>
</template>
```

:::

## 组合使用

:::demo

```vue
<template>
  <v-form>
    <v-row>
      <v-col cols="12">
        <vx-label
          requiredSymbol
          icon="mdi-account-outline"
          tooltip="请输入用户的真实姓名"
          labelFor="username"
        >
          用户姓名
        </vx-label>
        <vx-field
          id="username"
          placeholder="请输入姓名"
          variant="outlined"
          density="compact"
          class="mt-2"
        />
      </v-col>

      <v-col cols="12">
        <vx-label
          requiredSymbol
          icon="mdi-email-outline"
          tooltip="用于接收系统通知和密码重置"
          tooltipLocation="left"
          labelFor="email"
        >
          邮箱地址
        </vx-label>
        <vx-field
          id="email"
          placeholder="请输入邮箱"
          variant="outlined"
          density="compact"
          class="mt-2"
        />
      </v-col>

      <v-col cols="12">
        <vx-label toggleLabel>
          <template #prepend>
            <v-icon icon="mdi-cog-outline" class="mr-2" />
          </template>
          高级设置
          <template #toggle-content>
            <v-card class="mt-3 pa-4" variant="outlined">
              <v-row>
                <v-col cols="6">
                  <vx-label tooltip="选择用户角色">角色</vx-label>
                  <vx-select
                    :items="['管理员', '普通用户', '访客']"
                    variant="outlined"
                    density="compact"
                    class="mt-2"
                  />
                </v-col>
                <v-col cols="6">
                  <vx-label tooltip="设置用户状态">状态</vx-label>
                  <vx-select
                    :items="['启用', '禁用', '待审核']"
                    variant="outlined"
                    density="compact"
                    class="mt-2"
                  />
                </v-col>
              </v-row>
            </v-card>
          </template>
        </vx-label>
      </v-col>
    </v-row>
  </v-form>
</template>
```

:::
