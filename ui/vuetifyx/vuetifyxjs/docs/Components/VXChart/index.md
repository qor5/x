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

const chartData = ref({
  title: {
    text: 'Daily Active Users'
  },
  xAxis: {
    data: ['Location', 'Location', 'Location', 'Location', 'Location', 'Location']
  },
  series: [
    {
      name: '销量',
      data: [5, 20, 36, 10, 10, 20]
    }
  ]
})
</script>
<template>
  <div class="chart-container vx-border vx-border-gray-500 vx-rounded-lg vx-mb-5">
    <vx-chart ref="chartRef" presets="barChart" :options="chartData"></vx-chart>
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
