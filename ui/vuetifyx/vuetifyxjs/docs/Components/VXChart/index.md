# vx-chart 图表组件

基于ECharts封装的图表组件，提供了常用的图表预设和配置选项。

## API

### Props

| 参数名  | 说明                                                    | 类型              | 默认值 |
| ------- | ------------------------------------------------------- | ----------------- | ------ |
| presets | 预设样式，可选值：'barChart'、'pieChart'、'funnelChart' | String            | ''     |
| options | 图表配置项，会与预设样式合并                            | Object \ Object[] | {}     |
| loading | 是否显示加载状态                                        | Boolean           | false  |

### Slots

| 名称   | 说明                                     | 插槽 Props                                                                                                        |
| ------ | ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| action | 图表操作区域，通常用于切换不同的图表配置 | list: number[]（可用索引列表）<br>currentIndex: number（当前索引）<br>toggle: (index: number) => void（切换函数） |

## 预设类型

VXChart 组件提供了三种预设类型，可以通过 `presets` 属性指定：

- `barChart`：柱状图预设，适用于展示分类数据的数量对比
- `pieChart`：饼图预设，适用于展示占比数据
- `funnelChart`：漏斗图预设，适用于展示转化数据

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
          value: 10,
          name: '男性 10%'
        },
        {
          value: 90,
          name: '女性 90%'
        }
      ]
    }
  ]
})
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="pieChart" :options="pieChartData"></vx-chart>
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

使用 `funnelChart` 预设可以快速创建美观的漏斗图，用于展示转化流程和各环节的数据：

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
          value: 10010,
          name: 'Email Sent',
          extraData: {
            icon: 'mdi-near-me',
            labelList: [
              {
                type: 'primary',
                text: '10,010'
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
          value: 8500,
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
                icon: 'mdi-arrow-top-right',
                text: '+1.01% this week'
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
          value: 5000,
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
          value: 2500,
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

> **注意**：漏斗图的图例（legend）数据会根据传入的 series[0].data 中的 name 字段自动生成，不需要手动指定 legend.data。

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
