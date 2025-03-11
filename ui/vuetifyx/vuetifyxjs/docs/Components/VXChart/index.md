# vx-chart 图表组件

基于ECharts封装的图表组件，提供了常用的图表预设和配置选项。

## API

### Props

| 参数名  | 说明                                     | 类型               | 默认值 |
| ------- | ---------------------------------------- | ------------------ | ------ |
| presets | 预设样式，可选值：'barChart'、'pieChart' | String             | ''     |
| options | 图表配置项，会与预设样式合并             | Object \| Object[] | {}     |
| loading | 是否显示加载状态                         | Boolean            | false  |

### Slots

| 名称   | 说明                                     | 插槽 Props                                                                                                         |
| ------ | ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------ |
| action | 图表操作区域，通常用于切换不同的图表配置 | list: number[]（可用索引列表）<br>current-index: number（当前索引）<br>toggle: (index: number) => void（切换函数） |

## 示例

### 柱状图示例

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
    <vx-chart ref="chartRef" presets="barChart" :options="chartData">
      <template #action="{ list, 'current-index': currentIndex, toggle }">
        <div class="d-flex rounded-pill bg-grey-lighten-4 pa-1 mt-4 mr-4">
          <button
            v-for="(_, idx) in list"
            :key="idx"
            class="text-body-2 rounded-pill px-3 py-1 text-no-wrap border-0"
            :class="
              currentIndex === idx ? 'bg-primary text-white' : 'bg-transparent text-medium-emphasis'
            "
            style="cursor: pointer; transition: all 0.3s;"
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
  height: 400px;
  position: relative;
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
  <div class="chart-container border border-gray-500 rounded-lg">
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
  <div class="chart-container border border-gray-500 rounded-lg">
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

### 多图表切换示例

使用 `options` 数组和 `action` 插槽可以实现多图表切换功能：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

// 定义多个图表配置
const multiChartData = ref([
  {
    title: {
      text: '销售额'
    },
    xAxis: {
      data: ['1月', '2月', '3月', '4月', '5月', '6月']
    },
    series: [
      {
        type: 'bar',
        name: '销售额',
        data: [120, 200, 150, 80, 70, 110]
      }
    ]
  },
  {
    title: {
      text: '利润'
    },
    xAxis: {
      data: ['1月', '2月', '3月', '4月', '5月', '6月']
    },
    series: [
      {
        type: 'line',
        name: '利润',
        data: [20, 40, 35, 15, 10, 25],
        smooth: true,
        itemStyle: {
          color: '#4CAF50'
        },
        lineStyle: {
          width: 3,
          color: '#4CAF50'
        }
      }
    ]
  },
  {
    title: {
      text: '用户数'
    },
    xAxis: {
      data: ['1月', '2月', '3月', '4月', '5月', '6月']
    },
    series: [
      {
        type: 'bar',
        name: '用户数',
        data: [500, 800, 1200, 1500, 1800, 2200],
        itemStyle: {
          color: '#FF9800'
        }
      }
    ]
  }
])

// 标签文本
const tabLabels = ['销售额', '利润', '用户数']
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="barChart" :options="multiChartData">
      <template #action="{ list, 'current-index': currentIndex, toggle }">
        <div class="d-flex rounded-pill bg-grey-lighten-4 pa-1 mt-4 mr-4">
          <button
            v-for="idx in list"
            :key="idx"
            class="text-body-2 rounded-pill px-3 py-1 text-no-wrap border-0"
            :class="
              currentIndex === idx ? 'bg-primary text-white' : 'bg-transparent text-medium-emphasis'
            "
            style="cursor: pointer; transition: all 0.3s;"
            @click="toggle(idx)"
          >
            {{ tabLabels[idx] }}
          </button>
        </div>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
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

### 使用加载状态

```vue
<vx-chart presets="barChart" :options="chartData" :loading="isLoading"></vx-chart>
```
