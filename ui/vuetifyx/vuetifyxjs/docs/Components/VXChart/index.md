# vx-chart 图表组件

基于ECharts封装的图表组件，提供了常用的图表预设和配置选项。

## API

### Props

| 参数名  | 说明                                     | 类型    | 默认值 |
| ------- | ---------------------------------------- | ------- | ------ |
| presets | 预设样式，可选值：'barChart'、'pieChart' | String  | ''     |
| options | 图表配置项，会与预设样式合并             | Object  | {}     |
| loading | 是否显示加载状态                         | Boolean | false  |

## 示例

### 柱状图示例

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

// 定义7天和14天的数据
const sevenDaysData = {
  title: {
    text: 'Daily Active Users'
  },
  xAxis: {
    data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
  },
  series: [
    {
      name: '用户数',
      data: [5, 20, 36, 10, 10, 20, 30]
    }
  ]
}

const fourteenDaysData = {
  title: {
    text: 'Daily Active Users'
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
      name: '用户数',
      data: [5, 20, 36, 10, 10, 20, 30, 15, 25, 40, 20, 15, 25, 35]
    }
  ]
}

// 当前选中的时间范围
const selectedRange = ref('7days')

// 当前图表数据
const chartData = ref(sevenDaysData)

// 切换时间范围
const switchRange = (range) => {
  selectedRange.value = range
  chartData.value = range === '7days' ? sevenDaysData : fourteenDaysData
}
</script>
<template>
  <div class="chart-container vx-border vx-border-gray-500 vx-rounded-lg vx-mb-5">
    <div class="chart-header">
      <div class="chart-tabs">
        <button
          class="tab-button"
          :class="{ active: selectedRange === '7days' }"
          @click="switchRange('7days')"
        >
          Past 7 Days
        </button>
        <button
          class="tab-button"
          :class="{ active: selectedRange === '14days' }"
          @click="switchRange('14days')"
        >
          Past 14 Days
        </button>
      </div>
    </div>
    <vx-chart ref="chartRef" presets="barChart" :options="chartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
  position: relative;
}

.chart-header {
  position: absolute;
  top: 16px;
  right: 16px;
  z-index: 10;
}

.chart-tabs {
  display: flex;
  background-color: #f5f5f5;
  border-radius: 20px;
  padding: 2px;
}

.tab-button {
  border: none;
  background: transparent;
  padding: 6px 12px;
  border-radius: 16px;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s;
}

.tab-button.active {
  background-color: #1976d2;
  color: white;
}
</style>
```

:::

### 自定义显示/隐藏元素

可以通过配置选项来控制图表元素的显示或隐藏：

```vue
<vx-chart
  presets="barChart"
  :options="{
    // 隐藏y轴标签
    yAxis: {
      axisLabel: { show: false },
      // 隐藏横向虚线
      splitLine: { show: false }
    }
    // 其他配置...
  }"
></vx-chart>
```

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const customBarChartData = ref({
  title: {
    text: 'Age'
  },
  xAxis: {
    data: ['0-18', '18-25', '25-65', '65+']
  },
  yAxis: {
    // 隐藏y轴标签
    axisLabel: {
      show: false
    },
    // 隐藏横向虚线
    splitLine: {
      show: false
    }
  },
  series: [
    {
      name: '数据',
      data: [100, 300, 500, 200]
    }
  ]
})
</script>
<template>
  <div class="chart-container vx-border vx-border-gray-500 vx-rounded-lg vx-mb-5">
    <vx-chart ref="customBarChartRef" presets="barChart" :options="customBarChartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
}
</style>
```

:::

### 饼图示例

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const pieChartData = ref({
  title: {
    text: 'Gender'
  },
  series: [
    {
      name: '年龄分布',
      data: [
        {
          value: 10,
          name: 'Male 10%'
        },
        {
          value: 90,
          name: 'Female 90%'
        }
      ]
    }
  ]
})
</script>
<template>
  <div class="chart-container vx-border vx-border-gray-500 vx-rounded-lg vx-mb-5">
    <vx-chart ref="pieChartRef" presets="pieChart" :options="pieChartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
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

### 使用加载状态

```vue
<vx-chart presets="barChart" :options="chartData" :loading="isLoading"></vx-chart>
```
