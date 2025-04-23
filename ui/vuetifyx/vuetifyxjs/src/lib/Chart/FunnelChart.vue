<template>
  <div class="funnel-chart-container">
    <!-- 添加顶部标题栏 Frame 427323478
    <div class="funnel-header">
      <div class="campaign-name-container">
        <h1 class="campaign-name">Campaign Name</h1>
      </div>
      <div class="badge-container">
        <div class="badge light-badge">
          <span class="badge-text">Created On: 13:28 07/12/2024</span>
        </div>
        <div class="badge light-badge">
          <span class="badge-text">Last Updated: 13:28 07/12/2024</span>
        </div>
      </div>
    </div> -->

    <!-- 顶部统计卡片 Frame 427323603 -->
    <!-- <div class="funnel-summary-cards">
      <div class="summary-card">
        <div class="summary-tag blue">
          <div class="tag-dot blue-dot"></div>
          <div class="tag-text blue-text">Planed email amount</div>
        </div>
        <div class="summary-value">3000</div>
      </div>

      <div class="summary-card">
        <div class="summary-tag red">
          <div class="tag-dot red-dot"></div>
          <div class="tag-text red-text">Dropped</div>
        </div>
        <div class="summary-desc">because of user withdrew their email consent</div>
        <div class="summary-value">89,935</div>
      </div>

      <div class="summary-card">
        <div class="summary-tag orange">
          <div class="tag-dot orange-dot"></div>
          <div class="tag-text orange-text">Aborted</div>
        </div>
        <div class="summary-desc">Campaign manually paused</div>
        <div class="summary-value">89,935</div>
      </div>
    </div> -->
    <div class="funnel-cols-container" :style="containerStyles">
      <div class="funnel-cols">
        <!-- Email Sent Column -->
        <div class="funnel-col">
          <div class="funnel-card" :style="cardStyles">
            <div class="funnel-card-text" :style="cardTextStyles">Email Sent</div>
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon icon="mdi-near-me" color="#3E63DD" :size="iconSize" />
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ formatNumber(data[0]?.value || 0) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
        </div>

        <!-- Email Delivered Column -->
        <div class="funnel-col">
          <div class="funnel-card" :style="cardStyles">
            <div class="funnel-card-text" :style="cardTextStyles">Email Delivered</div>
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon icon="mdi-email-mark-as-unread" color="#3E63DD" :size="iconSize" />
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ formatNumber(data[1]?.value || 0) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ calculateDeliveryRate() }}%
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
        </div>

        <!-- Email Opened Column -->
        <div class="funnel-col">
          <div class="funnel-card" :style="cardStyles">
            <div class="funnel-card-text" :style="cardTextStyles">Email Opened</div>
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon icon="mdi-check-all" color="#3E63DD" :size="iconSize" />
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ formatNumber(data[2]?.value || 0) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">{{ calculateOpenRate() }}%</div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-bottom-left" color="#F44336" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">-1.01% this week</span>
            </div>
          </div>
        </div>

        <!-- Link Clicked Column -->
        <div class="funnel-col">
          <div class="funnel-card" :style="cardStyles">
            <div class="funnel-card-text" :style="cardTextStyles">Link Clicked</div>
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon icon="mdi-link" color="#3E63DD" :size="iconSize" />
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ formatNumber(data[3]?.value || 0) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ calculateClickRate() }}%
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon icon="mdi-arrow-top-right" color="#4CAF50" :size="iconSize" />
              <span class="trend-text" :style="trendTextStyles">+1.01% this week</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Actual Funnel Visual with echarts -->
      <div id="funnel-echarts-container" class="funnel-visual" :style="visualStyles"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as echarts from 'echarts'
import { funnelChartPreset } from './presets.config'

interface FunnelItem {
  value: number
  name: string
}

interface FunnelChartProps {
  data: FunnelItem[]
}

const props = defineProps<FunnelChartProps>()
const chartInstance = ref<echarts.EChartsType | null>(null)
const containerRef = ref<HTMLElement | null>(null)
const containerWidth = ref(0)
const scaleFactor = computed(() => {
  // When width is less than 700px but more than 534px, scale proportionally
  if (containerWidth.value < 700 && containerWidth.value >= 534) {
    // Linear interpolation between 1.0 and 0.7 based on width between 700 and 534
    const ratio = (containerWidth.value - 534) / (700 - 534)
    return 0.7 + ratio * 0.3
  }
  // When width is less than or equal to 534px, fix scale at 70%
  else if (containerWidth.value < 534) {
    return 0.7
  }
  // Default full size
  return 1
})

// Computed styles based on scale factor
const containerStyles = computed(() => {
  return {
    height: `${580 * scaleFactor.value}px`
  }
})

const cardStyles = computed(() => {
  return {
    padding: scaleFactor.value === 1 ? '22px 12px' : '15px 8px',
    height: `${68 * scaleFactor.value}px`,
    marginBottom: `${24 * scaleFactor.value}px`
  }
})

const cardTextStyles = computed(() => {
  return {
    fontSize: `${16 * scaleFactor.value}px`,
    lineHeight: `${20 * scaleFactor.value}px`
  }
})

const iconStyles = computed(() => {
  return {
    width: `${44 * scaleFactor.value}px`,
    height: `${44 * scaleFactor.value}px`,
    borderRadius: `${12 * scaleFactor.value}px`
  }
})

const statCardStyles = computed(() => {
  return {
    marginBottom: `${24 * scaleFactor.value}px`,
    padding: `0 ${12 * scaleFactor.value}px`
  }
})

const statValueStyles = computed(() => {
  return {
    fontSize: `${28 * scaleFactor.value}px`,
    marginBottom: `${12 * scaleFactor.value}px`,
    lineHeight: `${32 * scaleFactor.value}px`
  }
})

const statTrendStyles = computed(() => {
  return {
    fontSize: `${14 * scaleFactor.value}px`,
    lineHeight: `${18 * scaleFactor.value}px`
  }
})

const trendTextStyles = computed(() => {
  return {
    marginLeft: `${12 * scaleFactor.value}px`,
    lineHeight: `${18 * scaleFactor.value}px`
  }
})

const visualStyles = computed(() => {
  return {
    height: `${300 * scaleFactor.value}px`,
    minWidth: '534px'
  }
})

// Computed icon size based on scale factor
const iconSize = computed(() => {
  return Math.round(20 * scaleFactor.value)
})

// Format number to include commas
const formatNumber = (num: number): string => {
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
}

// Calculate delivery rate (second value / first value * 100)
const calculateDeliveryRate = () => {
  if (!props.data || props.data.length < 2 || !props.data[0].value) return 0
  return ((props.data[1].value / props.data[0].value) * 100).toFixed(1)
}

// Calculate open rate (third value / second value * 100)
const calculateOpenRate = () => {
  if (!props.data || props.data.length < 3 || !props.data[1].value) return 0
  return ((props.data[2].value / props.data[1].value) * 100).toFixed(1)
}

// Calculate click rate (fourth value / third value * 100)
const calculateClickRate = () => {
  if (!props.data || props.data.length < 4 || !props.data[2].value) return 0
  return ((props.data[3].value / props.data[2].value) * 100).toFixed(1)
}

// Initialize and configure the echarts instance
const initECharts = () => {
  const chartDom = document.getElementById('funnel-echarts-container')
  if (!chartDom) return

  console.log('Initializing funnel chart with data:', JSON.stringify(props.data))

  // Ensure we have 4 data points for the funnel chart
  if (!props.data || props.data.length < 4) {
    console.warn('FunnelChart: Missing data segments, expected 4 but got:', props.data?.length)
  }

  // Create echarts instance
  chartInstance.value = echarts.init(chartDom)

  // Configure chart
  updateEChartsOptions()
}

// Update chart options when data changes
const updateEChartsOptions = () => {
  if (!chartInstance.value) return

  console.log('Updating funnel chart with data:', JSON.stringify(props.data))

  // Get data with fallbacks to ensure we always have 4 segments
  const ensuredData = [...(props.data || [])]

  console.log('Final funnel chart data:', JSON.stringify(ensuredData))

  // Start with preset options
  const options = JSON.parse(JSON.stringify(funnelChartPreset))

  // Update series data with our ensured data
  options.series[0].data = ensuredData

  // Log the applied options
  console.log('Applied funnel chart options:', JSON.stringify(options.series[0]))

  // Apply options
  chartInstance.value.clear()
  chartInstance.value.setOption(options)
}

// Resize chart when window size changes
const handleResize = () => {
  if (chartInstance.value) {
    // Get the container width
    const containerWidth = document.getElementById('funnel-echarts-container')?.clientWidth || 0

    // If container is smaller than 534px, set chart to 534px width
    if (containerWidth < 534 && chartInstance.value) {
      chartInstance.value.resize({
        width: 534,
        height: chartInstance.value.getHeight()
      })
    } else {
      chartInstance.value.resize()
    }
  }
}

// Update container width
const updateContainerWidth = () => {
  if (containerRef.value) {
    containerWidth.value = containerRef.value.clientWidth
    console.log('Container width updated:', containerWidth.value)
  }
}

// Watch for data changes
watch(
  () => props.data,
  () => {
    updateEChartsOptions()
  },
  { deep: true }
)

// Initialize chart on component mount
onMounted(() => {
  containerRef.value = document.querySelector('.funnel-chart-container')
  updateContainerWidth()

  // Set up ResizeObserver to monitor container width changes
  const observer = new ResizeObserver((entries) => {
    for (const entry of entries) {
      containerWidth.value = entry.contentRect.width
      console.log('Container width changed:', containerWidth.value)
    }
  })

  if (containerRef.value) {
    observer.observe(containerRef.value)
  }

  initECharts()
  window.addEventListener('resize', () => {
    handleResize()
    updateContainerWidth()
  })

  // Cleanup observer
  onBeforeUnmount(() => {
    if (containerRef.value) {
      observer.unobserve(containerRef.value)
    }
  })
})

// Clean up on component unmount
onBeforeUnmount(() => {
  window.removeEventListener('resize', handleResize)
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

/* 漏斗图列样式 */
.funnel-cols {
  display: flex;
  width: 100%;
  flex: 1;
}

.funnel-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 16px;
  border-right: 1px solid #e0e0e0;

  &:last-child {
    border-right: none;
  }
}

.funnel-card {
  background-color: #f5f5f5;
  border: 1px solid #e0e0e0;
  border-radius: 12px;
  padding: 22px 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  height: 68px;
  box-sizing: border-box;
}

.funnel-card-text {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 16px;
  font-weight: 510;
  letter-spacing: 0.15px;
  color: #212121;
  // margin-left: 12px;
}

.funnel-card-icon {
  flex-shrink: 0;
  width: 44px;
  height: 44px;
  background-color: #ffffff;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  // margin-right: 12px;
}

.funnel-stat-card {
  margin-bottom: 24px;
  padding: 0 12px;
}

.funnel-stat-value {
  font-family:
    'SF Pro',
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
  font-size: 28px;
  font-weight: 510;
  letter-spacing: -0.12px;
  color: #212121;
  margin-bottom: 12px;
}

.funnel-stat-trend {
  display: flex;
  align-items: center;
  color: #616161;
  font-size: 14px;
}

.trend-text {
  margin-left: 12px;
}

.funnel-visual {
  height: 300px;
  width: 100%;
  position: absolute;
  bottom: 0;
  z-index: 1;
}

.funnel-cols-container {
  position: relative;
  height: 580px;
  min-width: 534px;
  .funnel-cols {
    // pointer-events: none;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;

    z-index: 2;
  }
}

/* Media query for smaller screens */
@media screen and (max-width: 700px) {
  .funnel-cols-container {
    height: 406px; /* 580px * 0.7 = 406px */
  }

  .funnel-chart-container {
    overflow-x: auto;
  }

  .funnel-card {
    padding: 15px 8px;
    height: 48px;
    margin-bottom: 17px;
  }

  .funnel-card-text {
    font-size: 11.2px;
  }

  .funnel-card-icon {
    width: 31px;
    height: 31px;
    border-radius: 8px;
  }

  .funnel-stat-card {
    margin-bottom: 17px;
    padding: 0 8px;
  }

  .funnel-stat-value {
    font-size: 19.6px;
    margin-bottom: 8px;
  }

  .funnel-stat-trend {
    font-size: 9.8px;
  }

  .trend-text {
    margin-left: 8px;
  }

  .funnel-visual {
    height: 210px;
  }
}
</style>
