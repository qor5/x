<template>
  <div class="vx-chart-wrap">
    <div ref="vxChartRoot"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import {
  ref,
  onMounted,
  defineProps,
  onBeforeUnmount,
  watch,
  computed,
  shallowRef,
  nextTick
} from 'vue'
import { chartPresets, ChartOptions, ChartSeriesItem } from './presets.config'

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
  },
  theme: {
    type: String,
    default: ''
  },
  autoResize: {
    type: Boolean,
    default: true
  },
  loading: {
    type: Boolean,
    default: false
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
  animation: true,
  animationDuration: 800,
  animationEasing: 'cubicInOut',
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
    if (source[key] !== undefined) {
      if (
        source[key] instanceof Object &&
        key in target &&
        target[key] instanceof Object &&
        !(source[key] instanceof Array)
      ) {
        result[key] = deepMerge(target[key], source[key])
      } else {
        result[key] = source[key]
      }
    }
  }

  return result
}

// 合并配置
const mergedOptions = computed(() => {
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
            ? { ...baseOptions.series[0] }
            : {}

      // 合并预设和用户配置
      return deepMerge(presetSeries, userSeries)
    })
  } else if (baseOptions.series) {
    // 如果用户没有提供series，则使用预设的series
    mergedSeries = JSON.parse(JSON.stringify(baseOptions.series))
  }

  // 合并其他配置
  const result = deepMerge(baseOptions, props.options)

  // 确保series被正确设置
  result.series = mergedSeries

  return result
})

// 初始化图表
const initChart = async () => {
  if (!vxChartRoot.value) return

  // 确保DOM已更新
  await nextTick()

  // 如果已经存在图表实例，先销毁
  if (vxChart.value) {
    vxChart.value.dispose()
  }

  // 创建新的图表实例
  vxChart.value = echarts.init(vxChartRoot.value, props.theme)

  // 设置图表配置
  vxChart.value.setOption(mergedOptions.value as any)

  // 设置加载状态
  if (props.loading) {
    vxChart.value.showLoading()
  } else {
    vxChart.value.hideLoading()
  }
}

// 监听配置变化
watch(
  () => [props.options, props.presets],
  () => {
    if (vxChart.value) {
      vxChart.value.setOption(mergedOptions.value as any, true)
    }
  },
  { deep: true }
)

// 监听加载状态
watch(
  () => props.loading,
  (loading) => {
    if (!vxChart.value) return

    if (loading) {
      vxChart.value.showLoading()
    } else {
      vxChart.value.hideLoading()
    }
  }
)

// 监听主题变化
watch(
  () => props.theme,
  () => {
    initChart()
  }
)

// 处理窗口大小变化
const handleResize = () => {
  if (vxChart.value) {
    vxChart.value.resize()
  }
}

// 组件挂载后初始化图表
onMounted(async () => {
  await initChart()

  if (props.autoResize) {
    window.addEventListener('resize', handleResize)
  }
})

// 组件卸载前清理资源
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
  position: relative;

  div {
    width: 100%;
    height: 100%;
    min-height: 300px;
  }
}
</style>
