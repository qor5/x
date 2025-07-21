<template>
  <div class="funnel-chart-container">
    <div class="funnel-cols-container mt-6" :style="containerStyles">
      <div class="funnel-cols" :style="colsStyles">
        <!-- 动态生成漏斗列 -->
        <div
          class="funnel-col"
          v-for="(item, index) in internalData"
          :key="index"
          :style="colStyles(item)"
        >
          <!-- 处理简单样式 -->
          <div v-if="item.extraData?.style === 'plain'">
            <div class="funnel-card-text" :style="cardTextStylesForPlainStyle(item)">
              {{ item.name }}
            </div>
          </div>
          <!-- 处理带icon的样式 -->
          <div v-else class="funnel-card" :style="cardStyles" :title="item.name">
            <vx-label
              class="funnel-card-text cardText mr-2"
              :tooltip="item.tooltip"
              tooltipLocation="top"
              tooltip-icon-color="#757575"
              >{{ item.name }}</vx-label
            >
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon :icon="item.extraData?.icon" color="#3E63DD" :size="iconSize" />
            </div>
          </div>

          <!-- 主数值卡片 本周 -->
          <div
            v-if="getStatObject(item, index, 0)?.text"
            class="funnel-stat-card"
            :style="statCardStyles(item)"
          >
            <!-- plain 布局样式稍微有所不同 -->
            <span
              v-if="item.extraData?.style !== 'plain' && getStatObject(item, index, 0)?.labelName"
              class="trend-text"
              :style="ThisWeekTextStyles"
              >{{ getStatObject(item, index, 0)?.labelName }}</span
            >
            <!-- 主数值 -->
            <div style="line-height: 1">
              <div
                class="funnel-stat-value"
                :style="[statValueStyles, getStatObject(item, index, 0)?.textStyle]"
              >
                {{ getStatObject(item, index, 0)?.text }}
              </div>
              <!-- 趋势 -->
              <div class="funnel-stat-trend" :style="statTrendStyles">
                <v-icon
                  :icon="getStatTrend(item, index, 0).icon"
                  :color="getStatTrend(item, index, 0).color"
                  :size="trendIconSize"
                />
                <span class="trend-text" :style="trendTextStyles">{{
                  getStatTrend(item, index, 0).text
                }}</span>
              </div>
            </div>
          </div>

          <!-- 主数值卡片 上周-->
          <div
            v-if="getStatObject(item, index, 1)?.text"
            class="funnel-stat-card"
            :style="statCardStyles(item)"
          >
            <!-- plain 布局样式稍微有所不同 -->
            <span
              v-if="item.extraData?.style !== 'plain' && getStatObject(item, index, 1)?.labelName"
              class="trend-text"
              :style="ThisWeekTextStyles"
              >{{ getStatObject(item, index, 1)?.labelName }}</span
            >
            <!-- 主数值 -->
            <div style="line-height: 1">
              <div
                class="funnel-stat-value"
                :style="[statValueStyles, getStatObject(item, index, 1)?.textStyle]"
              >
                {{ getStatObject(item, index, 1)?.text }}
              </div>
              <!-- 趋势 -->
              <div class="funnel-stat-trend" :style="statTrendStyles">
                <v-icon
                  :icon="getStatTrend(item, index, 1).icon"
                  :color="getStatTrend(item, index, 1).color"
                  :size="trendIconSize"
                />
                <span class="trend-text" :style="trendTextStyles">{{
                  getStatTrend(item, index, 1)?.text
                }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Actual Funnel Visual with echarts -->
      <div :id="chartId" class="funnel-visual" :style="visualStyles"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, PropType, onBeforeUnmount, watch } from 'vue'
import * as echarts from 'echarts'
import { funnelChartPreset, presetsLineInFunnel } from './presets.config'

interface LabelItem {
  labelName?: string
  type: string
  text: string
  textStyle?: string
  icon?: string
}

interface ExtraData {
  icon?: string
  style?: 'plain' | 'default'
  labelList?: LabelItem[]
}

interface FunnelItem {
  value: number
  name: string
  labelName?: string
  extraData?: ExtraData
  tooltip?: string
}

interface SeriesData {
  name: string
  type?: 'line' | 'funnel'
  data: FunnelItem[] | number[]
  isDisabled?: boolean
  smooth?: boolean
  lineColor?: string
}

interface ChartData {
  title?: {
    text: string
  }
  series: SeriesData[]
}

// Scaling configuration constants
const SCALE_CONFIG = {
  // Basic size configuration
  BASE_CONTAINER_HEIGHT: 541,
  BASE_CARD_HEIGHT: 68,
  BASE_CARD_PADDING: { vertical: 16, horizontal: 14 },
  BASE_ICON_SIZE: 44,
  BASE_ICON_BORDER_RADIUS: 12,
  BASE_MARGIN_BOTTOM: 24,
  STAT_CARD_MARGIN_BOTTOM: 14,
  BASE_STAT_PADDING: 12,

  // Font size configuration
  BASE_FONT_SIZES: {
    thisWeekText: 14,
    cardText: 18,
    statValue: 28,
    statTrend: 14,
    icon: 20
  },

  // Line height configuration
  BASE_LINE_HEIGHTS: {
    cardText: 20,
    statValue: 32,
    statTrend: 18
  },

  // Spacing configuration
  BASE_SPACINGS: {
    trendTextMargin: 6,
    trendMargin: 8
  },

  // Breakpoint configuration
  BREAKPOINTS: {
    // Minimum width, use minimum scaling below this width
    MIN_WIDTH: 400,
    // Comfortable width, use standard scaling above this width
    COMFORTABLE_WIDTH: 800,
    // Ideal minimum width per column
    MIN_COL_WIDTH: 120,
    // Ideal maximum width per column
    MAX_COL_WIDTH: 200
  },

  // Scaling range
  SCALE_RANGE: {
    MIN: 0.6,
    MAX: 1.2
  }
}

const props = defineProps({
  data: {
    type: [Array, Object] as PropType<FunnelItem[] | ChartData>,
    required: true
  },
  height: {
    type: String,
    default: 'auto'
  },
  mergeOptionsCallback: {
    type: Function as PropType<(options: any, mergeCallbackOptions: any) => void>,
    required: false
  }
})

if (props.height !== 'auto') {
  SCALE_CONFIG.BASE_CONTAINER_HEIGHT = parseInt(props.height)
}

watch(
  () => props.height,
  (newHeight) => {
    if (newHeight !== 'auto') {
      SCALE_CONFIG.BASE_CONTAINER_HEIGHT = parseInt(newHeight)
    } else {
      SCALE_CONFIG.BASE_CONTAINER_HEIGHT = 541
    }
  }
)

// Generate unique chart ID
const chartId = `funnel-chart-${Math.random().toString(36).substring(2, 10)}`

const chartInstance = ref<echarts.EChartsType | null>(null)
const containerRef = ref<HTMLElement | null>(null)
const containerWidth = ref(0)
const internalData = ref<FunnelItem[]>([])
const lineSeriesData = ref<SeriesData[]>([])
const lineChartOverlays = ref<echarts.EChartsType[]>([])
const observer = ref<ResizeObserver | null>(null)

// Process data format, support both old and new formats
const processData = (data: FunnelItem[] | ChartData) => {
  if (Array.isArray(data)) {
    // Old format: directly FunnelItem array
    internalData.value = [...data]
    lineSeriesData.value = []
  } else {
    // New format: ChartData object
    const funnelSeries = data.series.find((s) => !s.type || s.type === 'funnel')
    const lineSeries = data.series.filter((s) => s.type === 'line')

    if (funnelSeries && Array.isArray(funnelSeries.data)) {
      internalData.value = [...(funnelSeries.data as FunnelItem[])]
    } else {
      internalData.value = []
    }

    lineSeriesData.value = [...lineSeries]
  }
}

// Check if funnel chart visualization is disabled
const isFunnelChartDisabled = computed(() => {
  if (Array.isArray(props.data)) {
    return false // Old format doesn't support disable
  } else {
    // New format: check if funnel chart series is disabled
    const funnelSeries = props.data.series?.find((s) => !s.type || s.type === 'funnel')
    return funnelSeries?.isDisabled === true
  }
})

watch(
  () => props.data,
  (newData) => {
    processData(newData)
  },
  { immediate: true }
)

// Get statistical value based on item and index position
const getStatObject = (item: FunnelItem, index: number, statIndex: number) => {
  // statIndex 0: 1: main value this week, 2: main value last week
  if (item.extraData?.labelList && item.extraData.labelList.length >= statIndex * 2) {
    // Find corresponding primary labels (type is primary)
    const primaryLabels = item.extraData.labelList.filter((l) => l.type === 'primary')
    if (primaryLabels.length > statIndex) {
      return primaryLabels[statIndex]
    }
  }

  return {
    labelName: '',
    type: '',
    text: '',
    textStyle: ''
  }
}

// Get trend information based on item and index position
const getStatTrend = (item: FunnelItem, index: number, statIndex: number) => {
  // Find corresponding secondary labels (type is secondary)
  if (item.extraData?.labelList) {
    const secondaryLabels = item.extraData.labelList.filter((l) => l.type === 'secondary')
    if (secondaryLabels.length > statIndex) {
      const label = secondaryLabels[statIndex]
      return {
        icon: label.icon,
        text: label.text,
        color: label.icon?.includes('bottom') ? '#F44336' : '#4CAF50'
      }
    }
  }

  return { icon: '', text: '', color: '' }
}

// Smart scaling algorithm
const scalingMetrics = computed(() => {
  const colCount = internalData.value.length
  const width = containerWidth.value

  if (colCount === 0 || width === 0) {
    return {
      scaleFactor: 1,
      colWidth: SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH,
      isCompact: false,
      adaptiveSpacing: 1,
      useFixedSize: false
    }
  }

  // If 5+ columns, use fixed size mode
  const useFixedSize = colCount >= 6

  if (useFixedSize) {
    return {
      scaleFactor: 1, // Not used in fixed size mode
      colWidth: SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH,
      isCompact: false,
      adaptiveSpacing: 1,
      useFixedSize: true
    }
  }

  // For 4 columns and below, use original scaling logic
  const idealColWidth = Math.max(
    SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH,
    Math.min(SCALE_CONFIG.BREAKPOINTS.MAX_COL_WIDTH, width / colCount)
  )

  let scaleFactor = 1

  if (colCount <= 3) {
    // 3 columns and below, use standard scaling
    scaleFactor = Math.max(
      SCALE_CONFIG.SCALE_RANGE.MIN,
      Math.min(SCALE_CONFIG.SCALE_RANGE.MAX, width / SCALE_CONFIG.BREAKPOINTS.COMFORTABLE_WIDTH)
    )
  } else if (colCount === 4) {
    // 4 columns, standard scaling (default style)
    const baseScale = width / SCALE_CONFIG.BREAKPOINTS.COMFORTABLE_WIDTH
    scaleFactor = Math.max(SCALE_CONFIG.SCALE_RANGE.MIN, Math.min(1, baseScale))
  }

  // Width constraint adjustment for 4 columns and below
  if (width < SCALE_CONFIG.BREAKPOINTS.MIN_WIDTH) {
    scaleFactor *= 0.8
  } else if (width < SCALE_CONFIG.BREAKPOINTS.COMFORTABLE_WIDTH) {
    const widthRatio =
      (width - SCALE_CONFIG.BREAKPOINTS.MIN_WIDTH) /
      (SCALE_CONFIG.BREAKPOINTS.COMFORTABLE_WIDTH - SCALE_CONFIG.BREAKPOINTS.MIN_WIDTH)
    scaleFactor *= 0.8 + widthRatio * 0.2
  }

  // Ensure scaling factor is within reasonable range
  scaleFactor = Math.max(
    SCALE_CONFIG.SCALE_RANGE.MIN,
    Math.min(SCALE_CONFIG.SCALE_RANGE.MAX, scaleFactor)
  )

  const adaptiveSpacing = 1

  return {
    scaleFactor,
    colWidth: idealColWidth,
    isCompact: width < 800,
    adaptiveSpacing,
    useFixedSize: false
  }
})

// 容器样式
const containerStyles = computed(() => {
  const { scaleFactor } = scalingMetrics.value
  return {
    height: `${SCALE_CONFIG.BASE_CONTAINER_HEIGHT * scaleFactor}px`,
    minWidth: `${Math.max(400, internalData.value.length * SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH)}px`
  }
})

// 列容器样式
const colsStyles = computed(() => {
  const { isCompact } = scalingMetrics.value
  const colCount = internalData.value.length

  return {
    width: '100%', // 确保填充整个容器宽度
    display: 'flex',
    // 当列数很多时，确保每列平均分配宽度
    ...(colCount > 6 && {
      justifyContent: 'space-between'
    })
  }
})

// 单列样式
const colStyles = (item: FunnelItem) => {
  const { scaleFactor, adaptiveSpacing, isCompact, useFixedSize } = scalingMetrics.value
  const colCount = internalData.value.length

  // 计算每列应该占用的宽度百分比
  const colWidthPercent = colCount > 0 ? 100 / colCount : 100

  // Set padding based on useFixedSize: default 16px, fixed size mode 8px
  const basePadding = useFixedSize ? 8 : 16

  const baseStyles = {
    padding: `${basePadding * scaleFactor * adaptiveSpacing}px ${isCompact ? 4 : basePadding * scaleFactor * adaptiveSpacing}px`,
    minWidth: `${SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH * scaleFactor}px`
  }

  if (item.extraData?.style === 'plain') {
    baseStyles.padding = `${16 * scaleFactor * adaptiveSpacing}px`
  }

  if (colCount > 6) {
    // 当列数大于6时，使用flex-basis确保平均分配宽度
    return {
      ...baseStyles,
      flex: `1 1 ${colWidthPercent}%`,
      maxWidth: `${colWidthPercent}%`,
      width: `${colWidthPercent}%`
    }
  } else {
    // 列数较少时，使用原来的逻辑
    return {
      ...baseStyles,
      flex: '1',
      maxWidth: 'none'
    }
  }
}

// Card styles
const cardStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing, isCompact, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    // Fixed size for 5+ columns
    return {
      padding: '8px',
      // No height restriction
      marginBottom: `${SCALE_CONFIG.BASE_MARGIN_BOTTOM}px`
    }
  }

  // Original logic for 4 columns and below
  const verticalPadding = SCALE_CONFIG.BASE_CARD_PADDING.vertical * scaleFactor * adaptiveSpacing
  const horizontalPadding =
    SCALE_CONFIG.BASE_CARD_PADDING.horizontal * scaleFactor * adaptiveSpacing

  return {
    padding: isCompact
      ? `${Math.max(8, verticalPadding * 0.7)}px ${Math.max(6, horizontalPadding * 0.8)}px`
      : `${verticalPadding}px ${horizontalPadding}px`,
    height: `${SCALE_CONFIG.BASE_CARD_HEIGHT * scaleFactor * adaptiveSpacing}px`,
    marginBottom: `${SCALE_CONFIG.BASE_MARGIN_BOTTOM * scaleFactor * adaptiveSpacing}px`
  }
})

// Card text styles
const cardTextClassFontSize = computed(() => {
  const { scaleFactor, adaptiveSpacing, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    return '14px'
  }

  return `${SCALE_CONFIG.BASE_FONT_SIZES.cardText * scaleFactor * adaptiveSpacing}px`
})

const cardTextClassLineHeight = computed(() => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value

  return `${SCALE_CONFIG.BASE_LINE_HEIGHTS.cardText * scaleFactor * adaptiveSpacing}px`
})

const cardTextStylesForPlainStyle = (item: FunnelItem) => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value

  return {
    fontSize: `${SCALE_CONFIG.BASE_FONT_SIZES.cardText * 0.835 * scaleFactor * adaptiveSpacing}px`,
    lineHeight: `${SCALE_CONFIG.BASE_LINE_HEIGHTS.cardText * 0.835 * scaleFactor * adaptiveSpacing}px`,
    color: '#616161',
    marginBottom: '16px'
  }
}

// Icon styles
const iconStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    return {
      width: '32px',
      height: '32px',
      borderRadius: '8.7px'
    }
  }

  const size = SCALE_CONFIG.BASE_ICON_SIZE * scaleFactor * adaptiveSpacing

  return {
    width: `${size}px`,
    height: `${size}px`,
    borderRadius: `${SCALE_CONFIG.BASE_ICON_BORDER_RADIUS * scaleFactor * adaptiveSpacing}px`
  }
})

// 统计卡片样式
const statCardStyles = (item: FunnelItem) => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value

  return {
    lineHeight: '1.2',
    marginBottom: `${SCALE_CONFIG.STAT_CARD_MARGIN_BOTTOM * scaleFactor * adaptiveSpacing}px`,
    padding: `0 ${SCALE_CONFIG.BASE_STAT_PADDING * scaleFactor * adaptiveSpacing}px`,
    paddingLeft:
      item.extraData?.style === 'plain'
        ? '0'
        : `${SCALE_CONFIG.BASE_STAT_PADDING * scaleFactor * adaptiveSpacing}px`
  }
}

// Statistical value styles
const statValueStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    return {
      fontSize: '16px',
      lineHeight: '24px'
    }
  }

  return {
    fontSize: `${SCALE_CONFIG.BASE_FONT_SIZES.statValue * scaleFactor * adaptiveSpacing}px`,
    // marginBottom: `${SCALE_CONFIG.BASE_STAT_PADDING * scaleFactor * adaptiveSpacing}px`,
    lineHeight: `${SCALE_CONFIG.BASE_LINE_HEIGHTS.statValue * scaleFactor * adaptiveSpacing}px`
  }
})

// Statistical trend styles
const statTrendStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value
  return {
    fontSize: `${SCALE_CONFIG.BASE_FONT_SIZES.statTrend * scaleFactor * adaptiveSpacing}px`,
    lineHeight: `${SCALE_CONFIG.BASE_LINE_HEIGHTS.statTrend * scaleFactor * adaptiveSpacing}px`
  }
})

const ThisWeekTextStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    return {
      fontSize: '12px',
      fontWeight: '590'
    }
  }

  return {
    fontSize: `${SCALE_CONFIG.BASE_FONT_SIZES.thisWeekText * scaleFactor * adaptiveSpacing}px`,
    fontWeight: '590'
  }
})

// Trend text styles
const trendTextStyles = computed(() => {
  const { scaleFactor, adaptiveSpacing, useFixedSize } = scalingMetrics.value

  if (useFixedSize) {
    return {
      fontSize: '12px',
      marginLeft: `${SCALE_CONFIG.BASE_SPACINGS.trendTextMargin}px`,
      lineHeight: '18px'
    }
  }

  return {
    marginLeft: `${SCALE_CONFIG.BASE_SPACINGS.trendTextMargin * scaleFactor * adaptiveSpacing}px`,
    lineHeight: `${SCALE_CONFIG.BASE_LINE_HEIGHTS.statTrend * scaleFactor * adaptiveSpacing}px`
  }
})

// 可视化区域样式
const visualStyles = computed(() => {
  return {
    minWidth: `${Math.max(400, internalData.value.length * SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH)}px`
  }
})

// 图标大小
const iconSize = computed(() => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value
  return Math.round(SCALE_CONFIG.BASE_FONT_SIZES.icon * scaleFactor * adaptiveSpacing)
})

// 趋势图标大小
const trendIconSize = computed(() => {
  const { scaleFactor, adaptiveSpacing } = scalingMetrics.value
  return Math.round(SCALE_CONFIG.BASE_FONT_SIZES.icon * scaleFactor * adaptiveSpacing * 0.8)
})

const formatNumber = (num: number): string => {
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
}

const initECharts = () => {
  const chartDom = document.getElementById(chartId)
  if (!chartDom) return

  if (!internalData.value || internalData.value.length < 2) {
    console.warn('FunnelChart: Missing data segments', internalData.value?.length)
  }

  const containerWidth = chartDom.clientWidth
  const minWidth = Math.max(400, internalData.value.length * SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH)
  const chartWidth = containerWidth < minWidth ? minWidth : containerWidth

  chartInstance.value = echarts.init(chartDom, null, {
    width: chartWidth,
    height: getChartHeight()
  })

  updateEChartsOptions()

  setTimeout(() => {
    if (chartInstance.value) {
      handleResize()
    }
  }, 100)
}

// 动态计算图表高度
const getChartHeight = () => {
  const { scaleFactor } = scalingMetrics.value
  const baseHeight = 280
  const colCount = internalData.value.length

  // 根据列数调整高度
  let heightMultiplier = 1
  if (colCount > 6) {
    heightMultiplier = 0.8
  } else if (colCount > 4) {
    heightMultiplier = 0.9
  }

  return baseHeight * scaleFactor * heightMultiplier
}

const updateEChartsOptions = () => {
  if (!chartInstance.value) return

  const ensuredData = [...(internalData.value || [])]

  const options = JSON.parse(JSON.stringify(funnelChartPreset))

  // Filter out tooltip property from data to prevent ECharts from using it
  // Keep only properties that ECharts funnel chart needs: value, name, and other non-tooltip properties
  const chartData = ensuredData.map((item) => {
    const { tooltip, ...chartItem } = item
    return chartItem
  })

  // Set funnel chart data (without tooltip property)
  options.series[0].data = chartData

  // Handle merge options callback first
  if (props.mergeOptionsCallback) {
    props.mergeOptionsCallback(options, { seriesData: options.series })
  }

  // Then render funnel chart and add line chart overlay
  lineSeriesData.value.length > 0
    ? setTimeout(() => {
        addLineChartOverlay()
      }, 100)
    : cleanupLineChartOverlays()

  chartInstance.value.clear()
  !isFunnelChartDisabled.value && chartInstance.value.setOption(options)
}

// 清理折线图叠加层
const cleanupLineChartOverlays = () => {
  // 清理所有折线图实例
  lineChartOverlays.value.forEach((chart) => {
    if (chart) {
      chart.dispose()
    }
  })
  lineChartOverlays.value = []

  // 移除DOM元素
  const chartDom = document.getElementById(chartId)
  if (chartDom) {
    const overlays = chartDom.querySelectorAll('.line-chart-overlay')
    overlays.forEach((overlay) => overlay.remove())
  }
}

// 添加折线图叠加层
const addLineChartOverlay = () => {
  if (!chartInstance.value || lineSeriesData.value.length === 0) return

  // 先清理之前的叠加层
  cleanupLineChartOverlays()

  const chartDom = document.getElementById(chartId)
  if (!chartDom) return

  // 创建叠加层容器
  const overlay = document.createElement('div')
  overlay.className = 'line-chart-overlay'
  overlay.style.cssText = `
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    pointer-events: auto;
    z-index: 10;
  `

  // 创建折线图容器
  const lineChartContainer = document.createElement('div')
  lineChartContainer.style.cssText = `
    width: 100%;
    height: 100%;
  `
  overlay.appendChild(lineChartContainer)
  chartDom.appendChild(overlay)

  // 创建独立的折线图实例
  const lineChart = echarts.init(lineChartContainer, null, {
    width: chartDom.clientWidth,
    height: chartDom.clientHeight
  })

  // 配置独立的折线图
  const ensuredData = [...(internalData.value || [])]
  const baseSeriesConfig = presetsLineInFunnel.series?.[0] || {
    type: 'line',
    smooth: false,
    symbol: 'none',
    symbolSize: 0,
    showSymbol: false,
    connectNulls: false,
    lineStyle: { width: 2, color: '#3e63dd' }
  }

  const lineChartOption = {
    backgroundColor: presetsLineInFunnel.backgroundColor,
    tooltip: presetsLineInFunnel.tooltip,
    grid: presetsLineInFunnel.grid,
    xAxis: {
      ...presetsLineInFunnel.xAxis,
      min: 0,
      max: ensuredData.length
    },
    yAxis: presetsLineInFunnel.yAxis,
    series: lineSeriesData.value
      .map((lineSeries) => {
        // 处理折线图数据
        let lineData = lineSeries.data

        if (!Array.isArray(lineData)) {
          console.warn('Line chart data must be an array')
          return null
        }

        // 格式化数据为中心坐标
        const formattedData = lineData.map((item, index) => {
          let value = 0
          if (typeof item === 'number') {
            value = item
          } else if (typeof item === 'object' && item !== null && 'value' in item) {
            value = item.value
          } else if (typeof item === 'string') {
            const num = parseFloat(item)
            value = isNaN(num) ? 0 : num
          }

          // 返回 [x, y] 坐标格式，x为列的真正中心位置
          return [index + 0.5, value]
        })

        // 调整数据长度
        while (formattedData.length < ensuredData.length) {
          const nextIndex = formattedData.length
          formattedData.push([nextIndex + 0.5, 0])
        }
        if (formattedData.length > ensuredData.length) {
          formattedData.splice(ensuredData.length)
        }

        // 以预设配置为基准，只覆盖数据相关属性
        return {
          ...baseSeriesConfig,
          name: lineSeries.name,
          data: formattedData,
          smooth: lineSeries.smooth || baseSeriesConfig.smooth,
          lineStyle: {
            ...baseSeriesConfig.lineStyle,
            color: lineSeries.lineColor || baseSeriesConfig.lineStyle?.color || '#3e63dd'
          }
        }
      })
      .filter(Boolean)
  }

  lineChart.setOption(lineChartOption)

  // 保存折线图实例以便后续清理
  lineChartOverlays.value.push(lineChart)
}

const handleResize = () => {
  if (chartInstance.value) {
    const containerWidth = document.getElementById(chartId)?.clientWidth || 0
    const minWidth = Math.max(
      400,
      internalData.value.length * SCALE_CONFIG.BREAKPOINTS.MIN_COL_WIDTH
    )
    const chartWidth = Math.max(containerWidth, minWidth)

    chartInstance.value.resize({
      width: chartWidth,
      height: getChartHeight()
    })

    // 同时调整折线图叠加层的尺寸
    lineChartOverlays.value.forEach((lineChart) => {
      if (lineChart) {
        lineChart.resize({
          width: chartWidth,
          height: getChartHeight()
        })
      }
    })
  }
}

const updateContainerWidth = () => {
  if (containerRef.value) {
    containerWidth.value = containerRef.value.clientWidth

    if (chartInstance.value) {
      handleResize()

      // 如果有折线图数据，重新创建叠加层以确保正确的位置和尺寸
      if (lineSeriesData.value.length > 0) {
        setTimeout(() => {
          addLineChartOverlay()
        }, 100)
      }
    }
  }
}

watch(
  [internalData, lineSeriesData],
  () => {
    updateEChartsOptions()
  },
  { deep: true }
)

// 监听折线图数据变化，当没有折线图时清理叠加层
watch(
  () => lineSeriesData.value,
  (newLineData) => {
    if (!newLineData || newLineData.length === 0) {
      // 如果没有折线图数据，清理所有叠加层
      cleanupLineChartOverlays()
    }
  },
  { deep: true }
)

watch(
  scalingMetrics,
  () => {
    if (chartInstance.value) {
      handleResize()
    }
  },
  { deep: true }
)

onMounted(() => {
  containerRef.value = document.querySelector('.funnel-chart-container')
  updateContainerWidth()

  observer.value = new ResizeObserver((entries) => {
    for (const entry of entries) {
      containerWidth.value = entry.contentRect.width

      // 立即触发resize处理
      if (chartInstance.value) {
        handleResize()
      }
    }
  })

  if (containerRef.value && observer.value) {
    observer.value.observe(containerRef.value)
  }

  initECharts()
  window.addEventListener('resize', () => {
    handleResize()
    updateContainerWidth()
  })
})

// Cleanup observer
onBeforeUnmount(() => {
  if (containerRef.value && observer.value) {
    observer.value.unobserve(containerRef.value)
    observer.value.disconnect()
  }

  window.removeEventListener('resize', handleResize)

  // 清理折线图叠加层
  cleanupLineChartOverlays()

  if (chartInstance.value) {
    chartInstance.value.dispose()
    chartInstance.value = null
  }
})
</script>

<style lang="scss" scoped>
.funnel-chart-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

/* 顶部标题栏样式 */
.funnel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 24px;
  height: 56px;
  width: 100%;
  box-sizing: border-box;
}

.campaign-name-container {
  display: flex;
  align-items: center;
  height: 40px;
}

.campaign-name {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 35px;
  font-weight: 510;
  line-height: 40px;
  color: #212121;
  letter-spacing: -0.16px;
  margin: 0;
}

.badge-container {
  display: flex;
  gap: 8px;
}

.badge {
  height: 20px;
  padding: 2px 8px;
  border-radius: 4px;
  display: flex;
  align-items: center;

  &.light-badge {
    background-color: #f5f5f5;
  }
}

.badge-text {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 12px;
  font-weight: 400;
  color: #424242;
  letter-spacing: 0.04px;
}

/* 顶部统计卡片样式 */
.funnel-summary-cards {
  display: flex;
  margin: 16px 0;
  gap: 16px;
  padding: 0 24px;
}

.summary-card {
  flex: 1;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  padding: 12px;
  background-color: #ffffff;
  height: 128px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.summary-tag {
  display: flex;
  align-items: center;
  padding: 4px 8px;
  border-radius: 4px;
  width: fit-content;
  height: 24px;

  &.blue {
    background-color: #e4ecfe;
  }

  &.red {
    background-color: #ffe5e5;
  }

  &.orange {
    background-color: #ffe8d7;
  }
}

.tag-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  margin-right: 8px;

  &.blue-dot {
    background-color: #3e63dd;
  }

  &.red-dot {
    background-color: #e5484d;
  }

  &.orange-dot {
    background-color: #f76808;
  }
}

.tag-text {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 16px;
  font-weight: 400;

  &.blue-text {
    color: #3e63dd;
  }

  &.red-text {
    color: #e5484d;
  }

  &.orange-text {
    color: #f76808;
  }
}

.summary-desc {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 16px;
  font-weight: 400;
  color: #616161;
  margin-top: 12px;
}

.summary-value {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 24px;
  font-weight: 510;
  letter-spacing: -0.1px;
  color: #212121;
  align-self: flex-start;
}

/* 漏斗图列样式 - 优化版 */
.funnel-cols {
  position: relative;
  display: flex;
  width: 100%;
  flex: 1;
  transition: gap 0.3s ease;
  box-sizing: border-box;
  z-index: 1;

  /* 当有gap时，移除所有边框 */
  &[style*='gap'] {
    .funnel-col {
      border-right: none !important;
      border-left: 1px solid #e0e0e0;

      &:first-child {
        border-left: none;
      }
    }
  }
}

.funnel-col {
  display: flex;
  flex-direction: column;
  border-right: 1px solid #e0e0e0;
  transition: all 0.3s ease;
  min-width: 0; /* 允许flex收缩 */
  box-sizing: border-box;
  overflow: hidden; /* 防止内容溢出 */

  &:last-child {
    border-right: none;
  }

  /* 当列数很多时，减少边框宽度 */
  &:nth-child(n + 7) {
    border-right-width: 0.5px;
  }
}

.funnel-card {
  background-color: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-sizing: border-box;
  transition: all 0.3s ease;
  min-height: 0; /* 允许高度收缩 */
}

.funnel-card-text {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-weight: 510;
  letter-spacing: 0.15px;
  color: #212121;
  flex: 1;
  min-width: 0;
  word-break: break-word;
  transition: font-size 0.3s ease;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  text-overflow: ellipsis;

  &.cardText * {
    font-size: v-bind(cardTextClassFontSize) !important;
    line-height: v-bind(cardTextClassLineHeight) !important;
  }

  &.cardText:deep(.text-subtitle-2) {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /* 高列数时允许换行 */
  .funnel-col:nth-child(n + 7) & {
    white-space: normal;
    line-height: 1.2;
  }
}

.funnel-card-icon {
  flex-shrink: 0;
  background-color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.funnel-stat-card {
  transition: all 0.3s ease;
  min-height: 0; /* 允许高度收缩 */
}

.funnel-stat-value {
  display: inline-block;
  vertical-align: middle;
  margin-right: 8px;
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-weight: 510;
  letter-spacing: -0.12px;
  color: #212121;
  word-break: break-word;
  transition: all 0.3s ease;
  overflow: hidden;
  text-overflow: ellipsis;

  &.lighter {
    color: #9e9e9e;
  }
}

.funnel-stat-trend {
  display: inline-flex;
  vertical-align: middle;
  align-items: center;
  color: #616161;
  font-weight: bold;
  transition: all 0.3s ease;
  min-height: 0;
}

.trend-text {
  transition: all 0.3s ease;
  word-break: break-word;
  min-width: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #616161;

  /* 高列数时允许换行 */
  .funnel-col:nth-child(n + 7) & {
    white-space: normal;
    line-height: 1.2;
  }
}

.funnel-visual {
  width: 100%;
  position: absolute;
  bottom: 0;
  z-index: 2;
}

.funnel-cols-container {
  position: relative;
  transition: all 0.3s ease;

  .funnel-cols {
    // pointer-events: none;
    flex-direction: row;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    z-index: 2;
  }
}

/* 响应式优化 */
@media (max-width: 768px) {
  .funnel-col {
    border-bottom: 1px solid #e0e0e0;

    &:last-child {
      border-bottom: none;
    }
  }

  .funnel-cols {
    flex-direction: column;
  }
}

/* 高列数优化 - 更精细的控制 */
.funnel-cols {
  /* 7列及以上时的特殊处理 */
  &:has(.funnel-col:nth-child(7)) {
    .funnel-col {
      border-right-width: 0.5px;

      .funnel-card-text,
      .trend-text {
        font-size: 0.9em;
        line-height: 1.2;
      }
    }
  }

  /* 10列及以上时的更激进优化 */
  &:has(.funnel-col:nth-child(10)) {
    .funnel-col {
      .funnel-card-text,
      .trend-text {
        font-size: 0.8em;
        line-height: 1.1;
      }

      .funnel-stat-value {
        font-size: 0.9em;
      }
    }
  }
}
</style>
