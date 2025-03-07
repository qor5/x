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
      color: '#999',
      fontSize: 12,
      margin: 12 // 增加与轴线的距离
    },
    // 隐藏x轴线
    axisLine: {
      show: false
    }
  },
  yAxis: {
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
