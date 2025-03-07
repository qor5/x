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
    text: 'Location'
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
  <div class="chart-container">
    <vx-chart ref="chartRef" presets="barChart" :options="chartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
  border: 1px solid #eee;
  margin-bottom: 20px;
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
    text: 'Age'
  },
  series: [
    {
      name: '年龄分布',
      data: [
        {
          value: 20,
          name: '0-18'
        },
        {
          value: 35,
          name: '18-35'
        },
        {
          value: 30,
          name: '35-60'
        },
        {
          value: 15,
          name: '60-100'
        }
      ]
    }
  ]
})
</script>
<template>
  <div class="chart-container">
    <vx-chart ref="pieChartRef" presets="pieChart" :options="pieChartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
  border: 1px solid #eee;
  margin-bottom: 20px;
}
</style>
```

:::

## 预设样式说明

### barChart 柱状图预设

柱状图预设提供了美观的柱状图样式，包括：

- 蓝色主题的柱子，带有圆角
- 顶部显示数值标签
- 隐藏了Y轴刻度和线条
- 优化的提示框样式
- 平滑的动画效果

### pieChart 饼图预设

饼图预设提供了环形饼图样式，包括：

- 蓝色渐变主题
- 底部居中的图例
- 环形图设计（内半径55%，外半径70%）
- 优化的提示框，显示名称、数值和百分比
- 平滑的动画效果

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
