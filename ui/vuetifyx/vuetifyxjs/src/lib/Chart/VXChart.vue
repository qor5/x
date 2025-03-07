<template>
  <div class="vx-chart-wrap">
    <div ref="vxChartRoot"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { ref, onMounted, defineProps, onBeforeUnmount, watch, computed, shallowRef } from 'vue'

// 定义图表选项类型
interface ChartSeriesItem {
  type?: string
  name?: string
  data?: any[]
  radius?: string
  [key: string]: any
}

interface ChartOptions {
  tooltip?: any
  grid?: any
  xAxis?: any
  yAxis?: any
  series?: ChartSeriesItem[]
  [key: string]: any
}

const props = defineProps({
  presets: {
    type: String,
    validator: (value: string) => ['barChart', 'pieChart', ''].includes(value),
    default: ''
  },
  options: {
    type: Object as () => ChartOptions,
    default: () => ({})
  }
})

const vxChart = shallowRef<echarts.EChartsType | null>(null)
const vxChartRoot = ref<HTMLElement | null>(null)

// 默认配置，确保基础属性都存在
const defaultOptions: ChartOptions = {
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  series: []
}

// 预设配置
const getPresetOptions = (): ChartOptions => {
  switch (props.presets) {
    case 'barChart':
      return {
        ...defaultOptions,
        xAxis: {
          type: 'category',
          data: ['类别1', '类别2', '类别3', '类别4', '类别5']
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            type: 'bar',
            data: [10, 20, 30, 40, 50]
          }
        ]
      }
    case 'pieChart':
      return {
        ...defaultOptions,
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b}: {c} ({d}%)'
        },
        series: [
          {
            name: '数据',
            type: 'pie',
            radius: '50%',
            data: [
              { value: 335, name: '直接访问' },
              { value: 310, name: '邮件营销' },
              { value: 234, name: '联盟广告' },
              { value: 135, name: '视频广告' },
              { value: 1548, name: '搜索引擎' }
            ]
          }
        ]
      }
    default:
      return defaultOptions
  }
}

// 合并配置
const mergedOptions = computed((): ChartOptions => {
  const baseOptions = props.presets ? getPresetOptions() : defaultOptions
  return {
    ...baseOptions,
    ...props.options,
    // 确保series中的每个项目都有type属性
    series: ((props.options.series || baseOptions.series || []) as ChartSeriesItem[]).map(
      (item: ChartSeriesItem) => {
        if (!item.type) {
          // 如果series项没有type，根据preset设置默认type
          return {
            ...item,
            type: props.presets === 'pieChart' ? 'pie' : 'bar'
          }
        }
        return item
      }
    )
  }
})

const initChart = () => {
  if (vxChartRoot.value) {
    vxChart.value = echarts.init(vxChartRoot.value)
    vxChart.value.setOption(mergedOptions.value)
  }
}

watch(
  () => [props.options, props.presets],
  () => {
    if (vxChart.value) {
      vxChart.value.setOption(mergedOptions.value, true)
    }
  },
  { deep: true }
)

const handleResize = () => {
  if (vxChart.value) {
    vxChart.value.resize()
  }
}

onMounted(() => {
  initChart()
  window.addEventListener('resize', handleResize)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
  if (vxChart.value) {
    vxChart.value.dispose()
    vxChart.value = null
  }
})
</script>

<style lang="scss" scoped>
.vx-chart-wrap {
  width: 100%;
  height: 100%;

  div {
    width: 100%;
    height: 100%;
    min-height: 300px;
  }
}
</style>
