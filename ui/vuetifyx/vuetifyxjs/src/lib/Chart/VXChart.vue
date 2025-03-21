<template>
  <div class="vx-chart-wrap">
    <div class="d-flex align-center justify-space-between">
      <div v-if="chartTitle" class="vx-chart-title">{{ chartTitle }}</div>
      <slot
        name="action"
        :list="
          Array.isArray(props.options)
            ? Array.from({ length: props.options.length }, (_, i) => i)
            : [0]
        "
        :currentIndex="currentIndex"
        :toggle="toggle"
      ></slot>
    </div>
    <div :id="chartId" class="vx-chart-container"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import {
  computed,
  defineProps,
  nextTick,
  onBeforeUnmount,
  onMounted,
  PropType,
  ref,
  watch
} from 'vue'
import {
  animationPresets,
  ChartOptions,
  chartPresets,
  ChartSeriesItem,
  lightAnimationConfig
} from './presets.config'

// 定义预设类型
type PresetType = 'barChart' | 'pieChart' | 'funnelChart' | ''

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

// 当前显示的图表索引
const currentIndex = ref(0)

// 切换显示的图表
const toggle = (index: number) => {
  if (Array.isArray(props.options) && index >= 0 && index < props.options.length) {
    currentIndex.value = index

    // 更新图表
    const chartInstance = getChartInstance()
    if (chartInstance && Array.isArray(mergedOptions.value)) {
      chartInstance.setOption(mergedOptions.value[index] as any, true)
    }
  }
}

const props = defineProps({
  presets: {
    type: String as () => PresetType,
    validator: (value: string) => ['barChart', 'pieChart', 'funnelChart', ''].includes(value),
    default: ''
  },
  options: {
    type: [Object, Array] as PropType<ChartOptions | ChartOptions[]>,
    default: () => ({})
  },
  loading: {
    type: Boolean,
    default: false
  }
})

// 提取标题
const chartTitle = computed(() => {
  // 检查 options 是否为数组
  if (Array.isArray(props.options)) {
    // 如果是数组，返回当前选中的图表配置的标题
    return props.options[currentIndex.value]?.title?.text || ''
  }
  // 如果是对象，直接返回标题
  return props.options.title?.text || ''
})

// 内部使用的配置
const enableAnimation = true
const animationType: AnimationType = 'light'

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
  if (!enableAnimation) {
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

  if (animationType && animationType in animationPresets) {
    return animationPresets[animationType as keyof typeof animationPresets]
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

  // 处理 options 是数组的情况
  if (Array.isArray(props.options)) {
    // 对数组中的每一项应用预设配置
    const mergedOptionsArray = props.options.map((optionItem) => {
      // 处理单个配置项的 series 数据
      let itemMergedSeries: ChartSeriesItem[] = []

      if (optionItem.series && optionItem.series.length > 0) {
        // 如果用户提供了 series 数据，则使用用户的数据
        itemMergedSeries = optionItem.series.map((userSeries: ChartSeriesItem, index: number) => {
          // 获取对应的预设 series 配置
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
        // 如果用户没有提供 series，则使用预设的 series
        itemMergedSeries = JSON.parse(JSON.stringify(baseOptions.series))
      }

      // 合并其他配置
      const itemResult = deepMerge(baseOptions, optionItem)

      // 确保 series 被正确设置
      itemResult.series = itemMergedSeries

      // 为漏斗图类型自动设置 legend.data
      if (
        props.presets === 'funnelChart' &&
        itemMergedSeries.length > 0 &&
        itemMergedSeries[0].data
      ) {
        // 从 series 数据中提取 legend 数据
        if (!itemResult.legend) itemResult.legend = {}
        itemResult.legend.data = itemMergedSeries[0].data.map((item: any) => item.name)
      }

      // 从 ECharts 配置中移除标题
      if (itemResult.title) {
        delete itemResult.title
      }

      // 如果用户没有明确设置 animation，则应用动画配置
      if (optionItem.animation === undefined) {
        const animationConfig = getAnimationConfig()
        Object.assign(itemResult, animationConfig)
      }

      return itemResult
    })

    return mergedOptionsArray
  }

  // 处理 options 是对象的情况
  // 处理 series 数据
  let mergedSeries: ChartSeriesItem[] = []

  if (props.options.series && props.options.series.length > 0) {
    // 如果用户提供了 series 数据，则使用用户的数据
    mergedSeries = props.options.series.map((userSeries: ChartSeriesItem, index: number) => {
      // 获取对应的预设 series 配置
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
    // 如果用户没有提供 series，则使用预设的 series
    mergedSeries = JSON.parse(JSON.stringify(baseOptions.series))
  }

  // 合并其他配置
  const result = deepMerge(baseOptions, props.options)

  // 确保 series 被正确设置
  result.series = mergedSeries

  // 为漏斗图类型自动设置 legend.data
  if (props.presets === 'funnelChart' && mergedSeries.length > 0 && mergedSeries[0].data) {
    // 从 series 数据中提取 legend 数据
    if (!result.legend) result.legend = {}
    result.legend.data = mergedSeries[0].data.map((item: any) => item.name)
  }

  // 从 ECharts 配置中移除标题
  if (result.title) {
    delete result.title
  }

  // 如果用户没有明确设置 animation，则应用动画配置
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

  return echarts.getInstanceByDom(chartDom) || echarts.init(chartDom)
}

// 初始化图表
const initChart = async () => {
  // 确保DOM已更新
  await nextTick()

  const chartInstance = getChartInstance()
  if (!chartInstance) return

  // 设置图表配置
  // 判断 mergedOptions 是否为数组
  if (Array.isArray(mergedOptions.value)) {
    // 如果是数组，使用当前索引的配置项初始化图表
    chartInstance.setOption(mergedOptions.value[currentIndex.value] as any, true)
  } else {
    chartInstance.setOption(mergedOptions.value as any, true)
  }

  // 设置加载状态
  if (props.loading) {
    chartInstance.showLoading()
  } else {
    chartInstance.hideLoading()
  }
}

// 监听配置变化
watch(
  () => [props.options, props.presets],
  () => {
    const chartInstance = getChartInstance()
    if (chartInstance) {
      // 判断 mergedOptions 是否为数组
      if (Array.isArray(mergedOptions.value)) {
        // 如果是数组，使用当前索引的配置项更新图表
        chartInstance.setOption(mergedOptions.value[currentIndex.value] as any, true)
      } else {
        chartInstance.setOption(mergedOptions.value as any, true)
      }
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

// 监听当前索引变化
watch(
  () => currentIndex.value,
  (newIndex) => {
    const chartInstance = getChartInstance()
    if (chartInstance && Array.isArray(mergedOptions.value)) {
      chartInstance.setOption(mergedOptions.value[newIndex] as any, true)
    }
  }
)

// 处理窗口大小变化
const handleResize = () => {
  const chartInstance = getChartInstance()
  if (chartInstance) {
    chartInstance.resize()
  }
}
const resizeObserver = new ResizeObserver(() => {
  handleResize()
})

// 组件挂载后初始化图表
onMounted(async () => {
  await initChart()
  const chartDom = document.getElementById(chartId)
  if (chartDom) {
    resizeObserver.observe(chartDom)
  }
  // fix: sometimes the chart is not displayed correctly
  handleResize()
  window.addEventListener('resize', handleResize)
})

// 组件卸载前清理资源
onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)

  const chartDom = document.getElementById(chartId)
  if (chartDom) {
    resizeObserver.unobserve(chartDom)
    resizeObserver.disconnect()
    const instance = echarts.getInstanceByDom(chartDom)
    if (instance) {
      instance.dispose()
    }
  }
})
defineExpose({
  handleResize
})
</script>

<style lang="scss" scoped>
.vx-chart-wrap {
  width: 100%;
  height: 100%;
  position: relative;
  display: flex;
  flex-direction: column;

  .vx-chart-title {
    font-size: 18px;

    font-weight: 510;
    color: rgb(33, 33, 33);
    margin: 16px 0 0 16px;
    text-align: left;
  }

  .vx-chart-container {
    flex: 1;
    width: 100%;
    min-height: 300px;
  }
}
</style>
