<template>
  <div class="vx-chart-wrap">
    <div class="d-flex align-center justify-space-between">
      <div class="vx-chart-title">
        <slot name="title" :currentIndex="currentIndex">{{ chartTitle }}</slot>
      </div>
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

    <slot name="description" :currentIndex="currentIndex" />

    <!-- Use new FunnelChart component -->
    <template v-if="props.presets === 'funnelChart'">
      <funnel-chart
        :height="props.height"
        :data="getCurrentSeriesDataForFunnel()"
        :merge-options-callback="funnelMergeOptionsCallback"
      />
    </template>

    <div v-else :id="chartId" class="vx-chart-container"></div>
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
  watch,
  defineEmits
} from 'vue'
import {
  animationPresets,
  ChartOptions,
  chartPresets,
  ChartSeriesItem,
  lightAnimationConfig
} from './presets.config'
import FunnelChart from './FunnelChart.vue'
import { useVxChartMergeOptsCallback } from './useVxChartMergeOpts'

// Define preset types
type PresetType = 'barChart' | 'pieChart' | 'funnelChart' | ''

// Define animation types
type AnimationType =
  | 'light'
  | 'fadeInGrowth'
  | 'bounceGrowth'
  | 'waveGrowth'
  | 'sequentialGrowth'
  | ''

// Generate unique ID
const chartId = `vx-chart-${Math.random().toString(36).substring(2, 10)}`

// Current displayed chart index
const currentIndex = ref(0)

const emit = defineEmits(['on-change-index'])

// Toggle displayed chart
const toggle = (index: number) => {
  if (Array.isArray(props.options) && index >= 0 && index < props.options.length) {
    currentIndex.value = index

    // If not funnel chart, update ECharts chart
    if (props.presets !== 'funnelChart') {
      // Update chart
      const chartInstance = getChartInstance()
      if (chartInstance && Array.isArray(mergedOptions.value)) {
        chartInstance.setOption(mergedOptions.value[index] as any, true)
      }
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
  },
  // Add new props for funnel chart
  dataSource: {
    type: Object as PropType<{
      url?: string
      refreshInterval?: number
      fetchFn?: () => Promise<any[]>
    }>,
    default: () => ({})
  },
  height: {
    type: String,
    default: 'auto'
  },
  mergeOptionsCallback: {
    type: Function as PropType<
      (options: ChartOptions, mergeCallbackOptions: { seriesData: any[] }) => void
    >,
    default: () => {}
  }
})

const { invokeMergeOptionsCallback } = useVxChartMergeOptsCallback(props)

// Create merge options callback for funnel chart, passing current index information
const funnelMergeOptionsCallback = (options: any, data: any) => {
  // Call external mergeOptionsCallback and pass currentIndex
  if (props.mergeOptionsCallback) {
    props.mergeOptionsCallback(options, { ...data, currentIndex: currentIndex.value })
  }
}

// Get current series data for funnel chart component
const getCurrentSeriesDataForFunnel = () => {
  if (Array.isArray(props.options)) {
    // If options is array, return complete configuration of current selected index
    const currentOption = props.options[currentIndex.value]
    return {
      title: currentOption?.title,
      series: (currentOption?.series || []).map((s) => ({
        name: s.name || '',
        type: s.type as 'funnel' | 'line' | undefined,
        data: s.data || [],
        lineColor: s.lineColor,
        isDisabled: s.isDisabled,
        smooth: s.smooth
      }))
    }
  } else {
    // If options is object, return complete configuration
    return {
      title: props.options?.title,
      series: (props.options?.series || []).map((s) => ({
        name: s.name || '',
        type: s.type as 'funnel' | 'line' | undefined,
        data: s.data || [],
        lineColor: s.lineColor,
        isDisabled: s.isDisabled,
        smooth: s.smooth
      }))
    }
  }
}

// Extract title
const chartTitle = computed(() => {
  // Check if options is array
  if (Array.isArray(props.options)) {
    // If array, return title of currently selected chart configuration
    return props.options[currentIndex.value]?.title?.text || ''
  }
  // If object, directly return title
  return props.options.title?.text || ''
})

// Internal configuration
const enableAnimation = true
const animationType: AnimationType = 'light'

// Default configuration
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
  // Enable animation by default
  animation: true,
  series: []
}

// Get animation configuration
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

// Get preset configuration
const getChatOptions = (): ChartOptions => {
  let chatOptions: ChartOptions = {}

  if (props.presets && props.presets in chartPresets) {
    const presetOptions = chartPresets[props.presets as keyof typeof chartPresets]

    // Get animation configuration
    const animationConfig = getAnimationConfig()

    // Merge preset and animation configuration
    chatOptions = {
      ...presetOptions,
      ...animationConfig
    }
  } else {
    chatOptions = defaultOptions
  }
  return chatOptions
}

// Deep merge objects
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

// Merge configuration
const mergedOptions = computed(() => {
  const baseOptions = getChatOptions()

  // Handle case where options is array
  if (Array.isArray(props.options)) {
    // Apply preset configuration to each item in array
    const mergedOptionsArray = props.options.map((optionItem) => {
      // Process series data for single configuration item
      let itemMergedSeries: ChartSeriesItem[] = []

      if (optionItem.series && optionItem.series.length > 0) {
        // If user provided series data, use user's data
        itemMergedSeries = optionItem.series.map((userSeries: ChartSeriesItem, index: number) => {
          // Get corresponding preset series configuration
          const presetSeries =
            baseOptions.series && baseOptions.series[index]
              ? baseOptions.series[index]
              : baseOptions.series && baseOptions.series[0]
                ? { ...baseOptions.series[0] }
                : {}

          // Merge preset and user configuration
          return deepMerge(presetSeries, userSeries)
        })
      } else if (baseOptions.series) {
        // If user didn't provide series, use preset series
        itemMergedSeries = JSON.parse(JSON.stringify(baseOptions.series))
      }

      // Merge other configurations
      const itemResult = deepMerge(baseOptions, optionItem)

      // Ensure series is set correctly
      itemResult.series = itemMergedSeries

      // Automatically set legend.data for funnel chart type
      if (
        props.presets === 'funnelChart' &&
        itemMergedSeries.length > 0 &&
        itemMergedSeries[0].data
      ) {
        // Extract legend data from series data
        if (!itemResult.legend) itemResult.legend = {}
        itemResult.legend.data = itemMergedSeries[0].data.map((item: any) => item.name)
      }

      // Remove title from ECharts configuration
      if (itemResult.title) {
        delete itemResult.title
      }

      // If user didn't explicitly set animation, apply animation configuration
      if (optionItem.animation === undefined) {
        const animationConfig = getAnimationConfig()
        Object.assign(itemResult, animationConfig)
      }

      return itemResult
    })

    return mergedOptionsArray
  }

  // Handle case where options is object
  // Process series data
  let mergedSeries: ChartSeriesItem[] = []

  if (props.options.series && props.options.series.length > 0) {
    // If user provided series data, use user's data
    mergedSeries = props.options.series.map((userSeries: ChartSeriesItem, index: number) => {
      // Get corresponding preset series configuration
      const presetSeries =
        baseOptions.series && baseOptions.series[index]
          ? baseOptions.series[index]
          : baseOptions.series && baseOptions.series[0]
            ? { ...baseOptions.series[0] }
            : {}

      // Merge preset and user configuration
      return deepMerge(presetSeries, userSeries)
    })
  } else if (baseOptions.series) {
    // If user didn't provide series, use preset series
    mergedSeries = JSON.parse(JSON.stringify(baseOptions.series))
  }

  // Merge other configurations
  const result = deepMerge(baseOptions, props.options)

  // Ensure series is set correctly
  result.series = mergedSeries

  // Automatically set legend.data for funnel chart type
  if (props.presets === 'funnelChart' && mergedSeries.length > 0 && mergedSeries[0].data) {
    // Extract legend data from series data
    if (!result.legend) result.legend = {}
    result.legend.data = mergedSeries[0].data.map((item: any) => item.name)
  }

  // Remove title from ECharts configuration
  if (result.title) {
    delete result.title
  }

  // If user didn't explicitly set animation, apply animation configuration
  if (props.options.animation === undefined) {
    const animationConfig = getAnimationConfig()
    Object.assign(result, animationConfig)
  }

  return result
})

// Get current series data
const currentSeriesData = computed(() => {
  if (Array.isArray(mergedOptions.value)) {
    return mergedOptions.value[currentIndex.value]?.series?.[0]?.data || []
  }
  return mergedOptions.value?.series?.[0]?.data || []
})

const chartHeight = computed(() => {
  if (props.height !== 'auto') {
    return props.height + 'px'
  }
  return '300px'
})

// Get chart instance
const getChartInstance = (): echarts.EChartsType | null => {
  // If funnel chart, don't initialize echarts instance
  if (props.presets === 'funnelChart') {
    return null
  }

  const chartDom = document.getElementById(chartId)
  if (!chartDom) return null

  return echarts.getInstanceByDom(chartDom) || echarts.init(chartDom)
}

// Initialize chart
const initChart = async () => {
  // If funnel chart, don't initialize echarts chart
  if (props.presets === 'funnelChart') {
    return
  }

  // Ensure DOM is updated
  await nextTick()

  const chartInstance = getChartInstance()
  if (!chartInstance) return

  // Set chart configuration
  // Check if mergedOptions is array
  if (Array.isArray(mergedOptions.value)) {
    invokeMergeOptionsCallback(mergedOptions.value[currentIndex.value], {
      seriesData: currentSeriesData.value
    })
    // If array, use current index configuration item to initialize chart
    chartInstance.setOption(mergedOptions.value[currentIndex.value] as any, true)
  } else {
    invokeMergeOptionsCallback(mergedOptions.value, {
      seriesData: currentSeriesData.value
    })
    chartInstance.setOption(mergedOptions.value as any, true)
  }

  // Set loading state
  if (props.loading) {
    chartInstance.showLoading()
  } else {
    chartInstance.hideLoading()
  }
}

// Watch configuration changes
watch(
  () => [props.options, props.presets],
  () => {
    // If funnel chart, don't update echarts chart
    if (props.presets === 'funnelChart') {
      return
    }

    const chartInstance = getChartInstance()
    if (chartInstance) {
      // Check if mergedOptions is array
      if (Array.isArray(mergedOptions.value)) {
        // If array, use current index configuration item to update chart
        invokeMergeOptionsCallback(mergedOptions.value[currentIndex.value], {
          seriesData: currentSeriesData.value
        })
        chartInstance.setOption(mergedOptions.value[currentIndex.value] as any, true)
      } else {
        invokeMergeOptionsCallback(mergedOptions.value, {
          seriesData: currentSeriesData.value
        })
        chartInstance.setOption(mergedOptions.value as any, true)
      }
    }
  },
  { deep: true }
)

// Watch loading state
watch(
  () => props.loading,
  (loading) => {
    // If funnel chart, don't update echarts loading state
    if (props.presets === 'funnelChart') {
      return
    }

    const chartInstance = getChartInstance()
    if (!chartInstance) return

    if (loading) {
      chartInstance.showLoading()
    } else {
      chartInstance.hideLoading()
    }
  }
)

// Watch current index changes
watch(
  () => currentIndex.value,
  (oldIndex, newIndex) => {
    emit('on-change-index', newIndex)

    // If funnel chart, don't update echarts chart
    if (props.presets === 'funnelChart') {
      return
    }

    const chartInstance = getChartInstance()
    if (chartInstance && Array.isArray(mergedOptions.value)) {
      invokeMergeOptionsCallback(mergedOptions.value[newIndex], {
        seriesData: currentSeriesData.value
      })
      chartInstance.setOption(mergedOptions.value[newIndex] as any, true)
    }
  }
)

// Handle window resize
const handleResize = () => {
  // If funnel chart, don't resize echarts chart
  if (props.presets === 'funnelChart') {
    return
  }

  const chartInstance = getChartInstance()
  if (chartInstance) {
    chartInstance.resize()
  }
}
const resizeObserver = new ResizeObserver(() => {
  handleResize()
})

// Initialize chart after component mounted
onMounted(async () => {
  // If not funnel chart, initialize echarts chart
  if (props.presets !== 'funnelChart') {
    await initChart()
    const chartDom = document.getElementById(chartId)
    if (chartDom) {
      resizeObserver.observe(chartDom)
    }
    // fix: sometimes the chart is not displayed correctly
    handleResize()
  }

  window.addEventListener('resize', handleResize)
})

// Clean up resources before component unmounted
onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)

  // If not funnel chart, clean up echarts chart resources
  if (props.presets !== 'funnelChart') {
    const chartDom = document.getElementById(chartId)
    if (chartDom) {
      resizeObserver.unobserve(chartDom)
      resizeObserver.disconnect()
      const instance = echarts.getInstanceByDom(chartDom)
      if (instance) {
        instance.dispose()
      }
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
    min-height: v-bind(chartHeight);
  }
}
</style>
