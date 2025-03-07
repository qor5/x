<template>
  <div class="vx-chart-wrap">
    <div :id="chartId"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { ref, onMounted, defineProps, onBeforeUnmount, watch, computed, nextTick } from 'vue'
import {
  chartPresets,
  ChartOptions,
  ChartSeriesItem,
  lightAnimationConfig,
  animationPresets
} from './presets.config'

// 定义预设类型
type PresetType = 'barChart' | 'pieChart' | ''

// 定义动画类型
type AnimationType =
  | 'light'
  | 'fadeInGrowth'
  | 'bounceGrowth'
  | 'waveGrowth'
  | 'sequentialGrowth'
  | ''

// 生成唯一ID
const chartId = `vx-chart-${Math.random().toString(36).substring(2, 10)}`

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
  },
  // 是否启用动画，默认开启
  enableAnimation: {
    type: Boolean,
    default: true
  },
  // 动画类型
  animationType: {
    type: String as () => AnimationType,
    validator: (value: string) =>
      ['light', 'fadeInGrowth', 'bounceGrowth', 'waveGrowth', 'sequentialGrowth', ''].includes(
        value
      ),
    default: 'light'
  }
})

// 默认配置
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
  // 默认开启动画
  animation: true,
  series: []
}

// 获取动画配置
const getAnimationConfig = (): any => {
  if (!props.enableAnimation) {
    return {
      animation: false,
      animationDuration: 0,
      animationEasing: undefined,
      animationDelay: undefined,
      animationDurationUpdate: 0,
      animationEasingUpdate: undefined,
      animationDelayUpdate: undefined
    }
  }

  if (props.animationType && props.animationType in animationPresets) {
    return animationPresets[props.animationType as keyof typeof animationPresets]
  }

  return lightAnimationConfig
}

// 获取预设配置
const getPresetOptions = (): ChartOptions => {
  if (props.presets && props.presets in chartPresets) {
    const presetOptions = chartPresets[props.presets as keyof typeof chartPresets]

    // 获取动画配置
    const animationConfig = getAnimationConfig()

    // 合并预设和动画配置
    return {
      ...presetOptions,
      ...animationConfig
    }
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

  // 如果用户没有明确设置animation，则应用动画配置
  if (props.options.animation === undefined) {
    const animationConfig = getAnimationConfig()
    Object.assign(result, animationConfig)
  }

  return result
})

// 获取图表实例
const getChartInstance = (): echarts.EChartsType | null => {
  const chartDom = document.getElementById(chartId)
  if (!chartDom) return null

  return echarts.getInstanceByDom(chartDom) || echarts.init(chartDom, props.theme)
}

// 初始化图表
const initChart = async () => {
  // 确保DOM已更新
  await nextTick()

  const chartInstance = getChartInstance()
  if (!chartInstance) return

  // 设置图表配置
  chartInstance.setOption(mergedOptions.value as any, true)

  // 设置加载状态
  if (props.loading) {
    chartInstance.showLoading()
  } else {
    chartInstance.hideLoading()
  }
}

// 监听配置变化
watch(
  () => [props.options, props.presets, props.enableAnimation, props.animationType],
  () => {
    const chartInstance = getChartInstance()
    if (chartInstance) {
      chartInstance.setOption(mergedOptions.value as any, true)
    }
  },
  { deep: true }
)

// 监听加载状态
watch(
  () => props.loading,
  (loading) => {
    const chartInstance = getChartInstance()
    if (!chartInstance) return

    if (loading) {
      chartInstance.showLoading()
    } else {
      chartInstance.hideLoading()
    }
  }
)

// 监听主题变化
watch(
  () => props.theme,
  () => {
    // 主题变化时，需要重新初始化图表
    const chartDom = document.getElementById(chartId)
    if (chartDom) {
      const instance = echarts.getInstanceByDom(chartDom)
      if (instance) {
        instance.dispose()
      }
    }
    initChart()
  }
)

// 处理窗口大小变化
const handleResize = () => {
  const chartInstance = getChartInstance()
  if (chartInstance) {
    chartInstance.resize()
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

  const chartDom = document.getElementById(chartId)
  if (chartDom) {
    const instance = echarts.getInstanceByDom(chartDom)
    if (instance) {
      instance.dispose()
    }
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
