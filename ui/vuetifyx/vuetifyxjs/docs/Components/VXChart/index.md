# VXChart

这是一个基本示例，你可以用 `markdown` 语法 和 `vue3`、`vuetify` 在此处写任何组件代码

:::demo

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue'
import VueJsonPretty from 'vue-json-pretty'

const chartData = ref({
  title: {
    text: 'Location'
  },
  // 优化tooltip提示框
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c}',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    borderColor: '#eee',
    borderWidth: 1,
    textStyle: {
      color: '#333'
    },
    shadowBlur: 5,
    shadowColor: 'rgba(0, 0, 0, 0.1)'
  },
  // 隐藏图例
  legend: {
    show: false
  },
  xAxis: {
    data: ['Location', 'Location', 'Location', 'Location', 'Location', 'Location'],
    // 去掉x轴的分隔线
    splitLine: {
      show: false
    },
    // 隐藏x轴刻度线
    axisTick: {
      show: false
    },
    // 显示x轴标签
    axisLabel: {
      show: true,
      color: '#666',
      fontSize: 14,
      fontWeight: 'bold',
      margin: 12
    },
    // 隐藏x轴线
    axisLine: {
      show: false
    }
  },
  yAxis: {
    type: 'value', // 明确指定y轴类型
    // 去掉y轴的分隔线
    splitLine: {
      show: false
    },
    // 隐藏y轴刻度
    axisLabel: {
      show: false
    },
    // 隐藏y轴线
    axisLine: {
      show: false
    },
    // 隐藏y轴刻度线
    axisTick: {
      show: false
    }
  },
  series: [
    {
      name: '销量',
      type: 'bar',
      data: [5, 20, 36, 10, 10, 20],
      // 在柱子顶部显示数值
      label: {
        show: true,
        position: 'top',
        fontSize: 14,
        color: '#666'
      },
      // 调整柱子样式
      itemStyle: {
        borderRadius: [8, 8, 8, 8], // 柱子顶部和底部都有8px圆角
        color: 'rgba(62, 99, 221, 1)' // 柱子颜色为蓝色
      },
      // 设置柱子宽度为固定的32px
      barWidth: 32,
      // 添加鼠标悬停效果
      emphasis: {
        itemStyle: {
          color: 'rgba(62, 99, 221, 0.8)' // 悬停时颜色稍微变淡
        }
      }
    }
  ]
})

// 添加手动触发resize的函数
const chartRef = ref(null)
const triggerResize = () => {
  if (chartRef.value && chartRef.value.$el) {
    // 手动触发一次resize事件
    window.dispatchEvent(new Event('resize'))
  }
}

// 在组件挂载后触发一次resize
onMounted(() => {
  // 延迟执行以确保图表已完全渲染
  setTimeout(triggerResize, 300)
})
</script>
<template>
  <div class="chart-container">
    <vx-chart ref="chartRef" :options="chartData"></vx-chart>
  </div>
</template>
```

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
  border: 1px solid #eee;
  margin-bottom: 20px;
}
</style>

:::

### 4. 为组件撰写必要说明和参数

目前先随意，后期会有规范

### 饼图示例：年龄分布

这个示例展示了如何使用 VXChart 创建一个饼图来可视化年龄分布数据。

:::demo

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue'

const pieChartData = ref({
  title: {
    text: 'Age'
  },
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    borderColor: '#eee',
    borderWidth: 1,
    textStyle: {
      color: '#333'
    },
    shadowBlur: 5,
    shadowColor: 'rgba(0, 0, 0, 0.1)'
  },
  legend: {
    orient: 'horizontal',
    bottom: 10,
    left: 'center',
    itemWidth: 10,
    itemHeight: 10,
    textStyle: {
      fontSize: 12,
      color: '#666'
    }
  },
  series: [
    {
      name: '年龄分布',
      type: 'pie',
      radius: ['55%', '70%'], // 进一步调整环形图内外半径，使环形更瘦
      center: ['50%', '50%'],
      avoidLabelOverlap: true,
      itemStyle: {
        borderRadius: 0, // 移除圆角，使扇区之间没有间隔
        borderColor: '#fff',
        borderWidth: 0 // 移除边框，消除扇区间隔
      },
      label: {
        show: false // 隐藏标签
      },
      emphasis: {
        label: {
          show: false // 高亮时也不显示标签
        },
        itemStyle: {
          color: 'rgba(62, 99, 221, 1)', // 与柱状图相同的高亮颜色
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      labelLine: {
        show: false // 隐藏引导线
      },
      data: [
        {
          value: 20,
          name: '0-18',
          itemStyle: {
            color: 'rgba(230, 237, 254, 1)' // 最浅的蓝色
          }
        },
        {
          value: 35,
          name: '18-35',
          itemStyle: {
            color: 'rgba(200, 216, 248, 1)' // 稍微深一点的蓝色
          }
        },
        {
          value: 30,
          name: '35-60',
          itemStyle: {
            color: 'rgba(170, 195, 242, 1)' // 中等深度的蓝色
          }
        },
        {
          value: 15,
          name: '60-100',
          itemStyle: {
            color: 'rgba(140, 174, 236, 1)' // 最深的蓝色
          }
        }
      ]
    }
  ]
})

// 添加手动触发resize的函数
const pieChartRef = ref(null)
const triggerPieResize = () => {
  if (pieChartRef.value && pieChartRef.value.$el) {
    // 手动触发一次resize事件
    window.dispatchEvent(new Event('resize'))
  }
}

// 在组件挂载后触发一次resize
onMounted(() => {
  // 延迟执行以确保图表已完全渲染
  setTimeout(triggerPieResize, 300)
})
</script>
<template>
  <div class="chart-container">
    <vx-chart ref="pieChartRef" :options="pieChartData"></vx-chart>
  </div>
</template>
```

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
  border: 1px solid #eee;
  margin-bottom: 20px;
}
</style>

:::

### 饼图配置说明

- **title**: 设置图表标题为"Age"
- **tooltip**: 自定义提示框，显示名称、数值和百分比
- **legend**: 在底部居中显示图例
- **series**:
  - 使用更瘦的环形饼图（内半径55%，外半径70%）
  - 移除扇区圆角和边框，使各部分之间没有间隔
  - 为各个部分设置相近的浅蓝色系列颜色，形成微妙的色彩过渡
  - 高亮色与柱状图保持一致(rgba(62, 99, 221, 1))
  - 隐藏标签和引导线，仅在悬停时显示tooltip
  - 设置鼠标悬停效果
