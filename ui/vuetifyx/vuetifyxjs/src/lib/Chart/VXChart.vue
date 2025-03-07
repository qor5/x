<template>
  <div class="vx-chart-wrap">
    <div ref="vxChartRoot"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { ref, onMounted, defineProps, onBeforeUnmount, watch, computed, shallowRef } from 'vue'
import { chartPresets } from './presets.config'

// 定义图表选项类型
interface ChartSeriesItem {
  type?: string
  name?: string
  data?: any[]
  radius?: string | string[]
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

// 定义预设类型
type PresetType = 'barChart' | 'pieChart' | ''

const props = defineProps({
  presets: {
    type: String as () => PresetType,
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

// 获取预设配置
const getPresetOptions = (): ChartOptions => {
  if (props.presets && props.presets in chartPresets) {
    return chartPresets[props.presets as keyof typeof chartPresets]
  }
  return defaultOptions
}

// 深度合并对象
const deepMerge = (target: any, source: any): any => {
  const result = { ...target }

  for (const key in source) {
    if (source[key] instanceof Object && key in target && target[key] instanceof Object) {
      result[key] = deepMerge(target[key], source[key])
    } else {
      result[key] = source[key]
    }
  }

  return result
}

// 合并配置
const mergedOptions = computed((): ChartOptions => {
  const baseOptions = props.presets ? getPresetOptions() : defaultOptions

  // 处理series数据
  let mergedSeries: ChartSeriesItem[] = []

  if (props.options.series && props.options.series.length > 0) {
    // 如果用户提供了series数据，则使用用户的数据
    mergedSeries = props.options.series.map((userSeries: ChartSeriesItem, index: number) => {
      // 获取对应的预设series配置
      const presetSeries =
        baseOptions.series && baseOptions.series[index]
          ? baseOptions.series[index]
          : baseOptions.series && baseOptions.series[0]
            ? baseOptions.series[0]
            : {}

      // 合并预设和用户配置
      return deepMerge(presetSeries, userSeries)
    })
  } else if (baseOptions.series) {
    // 如果用户没有提供series，则使用预设的series
    mergedSeries = baseOptions.series
  }

  // 合并其他配置
  const result = deepMerge(baseOptions, props.options)

  // 确保series被正确设置
  result.series = mergedSeries

  return result
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
