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
    <div class="funnel-cols-container mt-6" :style="containerStyles">
      <div class="funnel-cols">
        <!-- 动态生成漏斗列 -->
        <div class="funnel-col" v-for="(item, index) in internalData" :key="index">
          <div class="funnel-card" :style="cardStyles">
            <div class="funnel-card-text" :style="cardTextStyles">{{ item.name }}</div>
            <div class="funnel-card-icon" :style="iconStyles">
              <v-icon :icon="item.extraData?.icon" color="#3E63DD" :size="iconSize" />
            </div>
          </div>
          <!-- 主数值卡片 -->
          <div class="funnel-stat-card" :style="statCardStyles">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ getStatValue(item, index, 0) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon
                :icon="getStatTrend(item, index, 0).icon"
                :color="getStatTrend(item, index, 0).color"
                :size="iconSize"
              />
              <span class="trend-text" :style="trendTextStyles">{{
                getStatTrend(item, index, 0).text
              }}</span>
            </div>
          </div>
          <!-- 转化率卡片 (对第一个阶段不显示) -->
          <div class="funnel-stat-card" :style="statCardStyles" v-if="index > 0">
            <div class="funnel-stat-value" :style="statValueStyles">
              {{ getStatValue(item, index, 1) }}
            </div>
            <div class="funnel-stat-trend" :style="statTrendStyles">
              <v-icon
                :icon="getStatTrend(item, index, 1).icon"
                :color="getStatTrend(item, index, 1).color"
                :size="iconSize"
              />
              <span class="trend-text" :style="trendTextStyles">{{
                getStatTrend(item, index, 1).text
              }}</span>
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

interface LabelItem {
  type: string
  text: string
  icon?: string
}

interface ExtraData {
  icon?: string
  labelList?: LabelItem[]
}

interface FunnelItem {
  value: number
  name: string
  extraData?: ExtraData
}

interface FunnelChartProps {
  data: FunnelItem[]
}

const props = defineProps<FunnelChartProps>()
const chartInstance = ref<echarts.EChartsType | null>(null)
const containerRef = ref<HTMLElement | null>(null)
const containerWidth = ref(0)
const internalData = ref<FunnelItem[]>([])
const observer = ref<ResizeObserver | null>(null)

watch(
  () => props.data,
  (newData) => {
    internalData.value = [...newData]
  },
  { immediate: true }
)

// 根据item和索引位置获取统计值
const getStatValue = (item: FunnelItem, index: number, statIndex: number) => {
  // statIndex 0: 主数值, 1: 转化率
  if (item.extraData?.labelList && item.extraData.labelList.length > statIndex * 2) {
    // 找到对应的主要标签（type为primary）
    const primaryLabels = item.extraData.labelList.filter((l) => l.type === 'primary')
    if (primaryLabels.length > statIndex) {
      return primaryLabels[statIndex].text
    }
  }

  return ''
}

// 根据item和索引位置获取趋势信息
const getStatTrend = (item: FunnelItem, index: number, statIndex: number) => {
  // 查找对应的次要标签（type为secondary）
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

const scaleFactor = computed(() => {
  if (containerWidth.value < 700 && containerWidth.value >= 534) {
    const ratio = (containerWidth.value - 534) / (700 - 534)
    return 0.7 + ratio * 0.3
  } else if (containerWidth.value < 534) {
    return 0.7
  }
  return 1
})

const containerStyles = computed(() => {
  return {
    height: `${541 * scaleFactor.value}px`
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
    minWidth: '534px'
  }
})

const iconSize = computed(() => {
  return Math.round(20 * scaleFactor.value)
})

const formatNumber = (num: number): string => {
  return num.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',')
}

const initECharts = () => {
  const chartDom = document.getElementById('funnel-echarts-container')
  if (!chartDom) return

  if (!internalData.value || internalData.value.length < 2) {
    console.warn('FunnelChart: Missing data segments', internalData.value?.length)
  }

  const containerWidth = chartDom.clientWidth
  const chartWidth = containerWidth < 534 ? 534 : containerWidth

  chartInstance.value = echarts.init(chartDom, null, {
    width: chartWidth,
    height: 280 * scaleFactor.value
  })

  updateEChartsOptions()

  setTimeout(() => {
    if (chartInstance.value) {
      handleResize()
    }
  }, 100)
}

const updateEChartsOptions = () => {
  if (!chartInstance.value) return

  const ensuredData = [...(internalData.value || [])]

  const options = JSON.parse(JSON.stringify(funnelChartPreset))

  options.series[0].data = ensuredData

  chartInstance.value.clear()
  chartInstance.value.setOption(options)
}

const handleResize = () => {
  if (chartInstance.value) {
    const containerWidth = document.getElementById('funnel-echarts-container')?.clientWidth || 0
    // console.log('containerWidth', containerWidth)
    chartInstance.value.resize({
      width: containerWidth,
      height: 280 * scaleFactor.value * (containerWidth < 797 ? 0.8 : 0.93)
    })
  }
}

const updateContainerWidth = () => {
  if (containerRef.value) {
    containerWidth.value = containerRef.value.clientWidth

    if (chartInstance.value) {
      handleResize()
    }
  }
}

watch(
  internalData,
  () => {
    updateEChartsOptions()
  },
  { deep: true }
)

watch(scaleFactor, () => {
  if (chartInstance.value) {
    handleResize()
  }
})

onMounted(() => {
  containerRef.value = document.querySelector('.funnel-chart-container')
  updateContainerWidth()

  observer.value = new ResizeObserver((entries) => {
    for (const entry of entries) {
      containerWidth.value = entry.contentRect.width
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
  width: 100%;
  position: absolute;
  bottom: 0;
  z-index: 1;
}

.funnel-cols-container {
  position: relative;
  height: 541px;
  min-width: 534px;
  .funnel-cols {
    pointer-events: none;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;

    z-index: 2;
  }
}
</style>
