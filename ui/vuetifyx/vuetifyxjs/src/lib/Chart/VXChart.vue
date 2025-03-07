<template>
  <div class="vx-chart-wrap">
    <div ref="vxChartRoot"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts'
import { ref, onMounted, defineProps, onBeforeUnmount, watch } from 'vue'

const props = defineProps({
  presets: {
    type: String,
    validator: (value: string) => ['barChart', 'pieChart', ''].includes(value),
    default: ''
  },
  options: {
    type: Object,
    default: () => ({})
  }
})

const vxChart = ref<echarts.EChartsType | null>(null)
const vxChartRoot = ref<HTMLElement | null>(null)

const initChart = () => {
  if (vxChartRoot.value) {
    vxChart.value = echarts.init(vxChartRoot.value)
    vxChart.value.setOption(props.options)
  }
}

watch(
  () => props.options,
  (newOptions) => {
    if (vxChart.value) {
      vxChart.value.setOption(newOptions)
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
