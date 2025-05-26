# vx-chart 图表组件

基于ECharts封装的图表组件，提供了常用的图表预设和配置选项。

## API

### Props

| 参数名               | 说明                                                                                                                                          | 类型              | 默认值   |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | ----------------- | -------- |
| presets              | 预设样式，可选值：'barChart'、'pieChart'、'funnelChart'                                                                                       | String            | ''       |
| options              | 图表配置项，会与预设样式合并                                                                                                                  | Object \ Object[] | {}       |
| mergeOptionsCallback | 可以使用这个回调来修改当前的配置参数, 当需要自定义vx-chart配置的时候格外有用，详见 [#饼图示例](./#饼图示例) ，目前只支持 pieChart 和 barChart | Function          | () => {} |
| loading              | 是否显示加载状态                                                                                                                              | Boolean           | false    |

### Slots

| 名称   | 说明                                     | 插槽 Props                                                                                                        |
| ------ | ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| action | 图表操作区域，通常用于切换不同的图表配置 | list: number[]（可用索引列表）<br>currentIndex: number（当前索引）<br>toggle: (index: number) => void（切换函数） |

## 预设类型

VXChart 组件提供了三种预设类型，可以通过 `presets` 属性指定：

- `barChart`：柱状图预设，适用于展示分类数据的数量对比
- `pieChart`：饼图预设，适用于展示占比数据
- `funnelChart`：漏斗图预设，适用于展示转化数据，**支持无限多列的智能缩放**

## 基础示例

### 柱状图示例

使用 `barChart` 预设可以快速创建美观的柱状图：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const barChartData = ref({
  title: {
    text: '年龄分布'
  },
  xAxis: {
    data: ['0-18', '18-25', '25-65', '65+']
  },
  series: [
    {
      name: '人数',
      data: [100, 300, 500, 200]
    }
  ]
})
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="barChart" :options="barChartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

### 饼图示例

使用 `pieChart` 预设可以快速创建美观的饼图：

当你想自定义图例时可以使用 `mergeOptionsCallback`, 回调函数支持两个参数

- options - 当前图表配置项
- data - 一些图表数据对象

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const pieChartData = ref({
  title: {
    text: '性别比例'
  },
  series: [
    {
      name: '性别分布',
      data: [
        {
          value: 10.111,
          name: '男性'
        },
        {
          value: 89.889,
          name: '女性'
        }
      ]
    }
  ]
})

const mergeOptionsCallback = function (options, { seriesData }) {
  options.legend = {
    ...options.legend,
    formatter: (name) => {
      const item = seriesData.find((i) => i.name === name)
      const percent = ((item.value / 100) * 100).toFixed(2)
      return `${name} ${percent}%`
    }
  }
}
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart
      presets="pieChart"
      :options="pieChartData"
      :merge-options-callback="mergeOptionsCallback"
    ></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

### 漏斗图示例

使用 `funnelChart` 预设可以快速创建美观的漏斗图，用于展示转化流程和各环节的数据。**新版本支持无限多列的智能缩放算法**，能够根据列数和容器宽度自动调整元素大小和布局：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const funnelChartData = ref({
  title: {
    text: '邮件营销漏斗'
  },
  series: [
    {
      name: '邮件营销',
      data: [
        {
          value: 1000,
          name: 'Email Sent',
          extraData: {
            icon: 'mdi-near-me',
            labelList: [
              {
                type: 'primary',
                text: '1000'
              },
              {
                type: 'secondary',
                icon: '',
                text: ''
              }
            ]
          }
        },
        {
          value: 800,
          name: 'Email Delivered',
          extraData: {
            icon: 'mdi-email-mark-as-unread',
            labelList: [
              {
                type: 'primary',
                text: '8,500'
              },
              {
                type: 'secondary',
                icon: '',
                text: ''
              },
              {
                type: 'primary',
                text: '84.9%'
              },
              {
                type: 'secondary',
                icon: 'mdi-arrow-top-right',
                text: '+1.01% this week'
              }
            ]
          }
        },
        {
          value: 400,
          name: 'Email Opened',
          extraData: {
            icon: 'mdi-check-all',
            labelList: [
              {
                type: 'primary',
                text: '5,000'
              },
              {
                type: 'secondary',
                icon: 'mdi-arrow-top-right',
                text: '+1.01% this week'
              },
              {
                type: 'primary',
                text: '58.8%'
              },
              {
                type: 'secondary',
                icon: 'mdi-arrow-bottom-left',
                text: '-1.01% this week'
              }
            ]
          }
        },
        {
          value: 200,
          name: 'Link Clicked',
          extraData: {
            icon: 'mdi-link',
            labelList: [
              {
                type: 'primary',
                text: '2,500'
              },
              {
                type: 'secondary',
                icon: 'mdi-arrow-top-right',
                text: '+1.01% this week'
              },
              {
                type: 'primary',
                text: '50.0%'
              },
              {
                type: 'secondary',
                icon: 'mdi-arrow-top-right',
                text: '+1.01% this week'
              }
            ]
          }
        }
      ]
    }
  ]
})
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="funnelChart" :options="funnelChartData">
      <template #action>
        <span class="text-caption mr-4 px-1 py-0 rounded" style="background:#F5F5F5;"
          >Data updates on everyday's 00:00
        </span>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

### 多列漏斗图示例（智能缩放）

展示新的智能缩放算法如何处理更多列的情况：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const multiColumnFunnelData = ref({
  title: {
    text: '完整用户转化漏斗 (8个阶段)'
  },
  series: [
    {
      name: '用户转化',
      data: [
        {
          value: 10000,
          name: 'Visitors',
          extraData: {
            icon: 'mdi-account-group',
            labelList: [
              { type: 'primary', text: '10,000' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+5.2% this month' }
            ]
          }
        },
        {
          value: 8500,
          name: 'Page Views',
          extraData: {
            icon: 'mdi-eye',
            labelList: [
              { type: 'primary', text: '8,500' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+3.1% this month' },
              { type: 'primary', text: '85.0%' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+2.1% conversion' }
            ]
          }
        },
        {
          value: 6200,
          name: 'Engaged Users',
          extraData: {
            icon: 'mdi-heart',
            labelList: [
              { type: 'primary', text: '6,200' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+1.8% this month' },
              { type: 'primary', text: '72.9%' },
              { type: 'secondary', icon: 'mdi-arrow-bottom-left', text: '-1.2% conversion' }
            ]
          }
        },
        {
          value: 4800,
          name: 'Sign Ups',
          extraData: {
            icon: 'mdi-account-plus',
            labelList: [
              { type: 'primary', text: '4,800' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+4.5% this month' },
              { type: 'primary', text: '77.4%' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+3.2% conversion' }
            ]
          }
        },
        {
          value: 3600,
          name: 'Email Verified',
          extraData: {
            icon: 'mdi-email-check',
            labelList: [
              { type: 'primary', text: '3,600' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+2.1% this month' },
              { type: 'primary', text: '75.0%' },
              { type: 'secondary', icon: 'mdi-arrow-bottom-left', text: '-0.8% conversion' }
            ]
          }
        },
        {
          value: 2000,
          name: 'Email Verified2',
          extraData: {
            icon: 'mdi-email-check',
            labelList: [
              { type: 'primary', text: '3,600' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+2.1% this month' },
              { type: 'primary', text: '75.0%' },
              { type: 'secondary', icon: 'mdi-arrow-bottom-left', text: '-0.8% conversion' }
            ]
          }
        },
        {
          value: 1000,
          name: 'Email Verified3',
          extraData: {
            icon: 'mdi-email-check',
            labelList: [
              { type: 'primary', text: '3,600' },
              { type: 'secondary', icon: 'mdi-arrow-top-right', text: '+2.1% this month' },
              { type: 'primary', text: '75.0%' },
              { type: 'secondary', icon: 'mdi-arrow-bottom-left', text: '-0.8% conversion' }
            ]
          }
        }
      ]
    }
  ]
})
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="funnelChart" :options="multiColumnFunnelData">
      <template #action>
        <span class="text-caption mr-4 px-1 py-0 rounded" style="background:#E3F2FD;"
          >智能缩放算法自动适配 8 列布局
        </span>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  min-height: 400px;
}
</style>
```

:::

> **注意**：漏斗图的图例（legend）数据会根据传入的 series[0].data 中的 name 字段自动生成，不需要手动指定 legend.data。

## 漏斗图智能缩放算法

新版本的漏斗图组件采用了智能缩放算法，具有以下特性：

### 🎯 核心特性

- **无限列支持**：支持任意数量的列，从 2 列到 20+ 列
- **智能缩放**：根据列数和容器宽度自动计算最佳缩放比例
- **自适应布局**：元素大小、间距、字体大小都会根据列数智能调整
- **响应式设计**：在不同屏幕尺寸下都能保持良好的显示效果

### 📐 缩放策略

| 列数范围 | 缩放策略 | 特点                                |
| -------- | -------- | ----------------------------------- |
| 1-3 列   | 标准缩放 | 保持最佳视觉效果，元素大小适中      |
| 4-6 列   | 适度缩放 | 每增加一列减少 10% 大小，保持可读性 |
| 7+ 列    | 激进缩放 | 更大幅度缩放，启用紧凑模式          |

### 🔧 技术细节

- **最小宽度保护**：每列最小宽度 120px，确保内容可读
- **自适应间距**：列数超过 4 列时自动减少间距
- **平滑过渡**：所有缩放变化都有 0.3s 的过渡动画
- **性能优化**：使用 computed 属性缓存计算结果

### 📱 响应式支持

- **移动端优化**：在小屏幕上自动切换为垂直布局
- **容器适配**：根据父容器宽度动态调整
- **最小宽度限制**：确保在任何情况下都不会过度压缩

## 功能扩展

### 多图表切换

使用 `options` 数组和 `action` 插槽可以实现多图表切换功能：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

// 定义7天和14天的数据
const chartData = ref([
  {
    title: {
      text: 'Daily Active Users (7 Days)'
    },
    xAxis: {
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    series: [
      {
        type: 'bar',
        name: '用户数',
        data: [5, 20, 36, 10, 10, 20, 30]
      }
    ]
  },
  {
    title: {
      text: 'Daily Active Users (14 Days)'
    },
    xAxis: {
      data: [
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2'
      ]
    },
    series: [
      {
        type: 'bar',
        name: '用户数',
        data: [5, 20, 36, 10, 10, 20, 30, 15, 25, 40, 20, 15, 25, 35]
      }
    ]
  }
])
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="barChart" :options="chartData">
      <template #action="{ list, currentIndex, toggle }">
        <div
          class="d-flex align-center bg-grey-lighten-3 rounded pa-1 mr-4 mt-4"
          style="height: 32px;"
        >
          <button
            v-for="(_, idx) in list"
            :key="idx"
            class="text-body-2 rounded text-no-wrap border-0 flex-grow-1 d-flex align-center justify-center rounded px-2"
            style="height: 24px; cursor: pointer; transition: all 0.3s;"
            :style="
              currentIndex === idx
                ? 'background-color: #fff; color: #4a4a4a;'
                : 'background-color: transparent; color: rgb(117, 117, 117);'
            "
            @click="toggle(idx)"
          >
            {{ idx === 0 ? 'Past 7 Days' : 'Past 14 Days' }}
          </button>
        </div>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;

  position: relative;
}
</style>
```

:::

## 使用说明

### 基本使用

只需传入数据，使用预设样式：

```vue
<vx-chart
  presets="barChart"
  :options="{
    xAxis: { data: ['A', 'B', 'C'] },
    series: [{ data: [10, 20, 30] }]
  }"
></vx-chart>
```

### 完全自定义

不使用预设，完全自定义配置：

```vue
<vx-chart :options="customOptions"></vx-chart>
```

### 混合使用

使用预设，但覆盖部分配置：

```vue
<vx-chart
  presets="pieChart"
  :options="{
    title: { text: '自定义标题' },
    series: [{ data: customData }]
  }"
></vx-chart>
```

自定义漏斗图示例:

```vue
<vx-chart
  presets="funnelChart"
  :options="{
    title: { text: '自定义漏斗图' },
    tooltip: {
      formatter: '{b}: {c}人 ({d}%)' // 自定义提示格式
    },
    color: ['#FF6B6B', '#FFD166', '#06D6A0', '#118AB2', '#073B4C'], // 自定义颜色
    series: [
      {
        name: '用户行为',
        data: myFunnelData,
        label: {
          position: 'right' // 将标签放在右侧（默认在左侧）
        }
      }
    ]
  }"
></vx-chart>
```

### 使用加载状态

```vue
<vx-chart presets="barChart" :options="chartData" :loading="isLoading"></vx-chart>
```
