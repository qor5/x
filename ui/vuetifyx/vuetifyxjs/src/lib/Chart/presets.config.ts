// 定义图表选项类型
export interface ChartSeriesItem {
  type?: string
  name?: string
  data?: any[]
  radius?: string | string[]
  [key: string]: any
}

export interface ChartOptions {
  tooltip?: any
  grid?: any
  xAxis?: any
  yAxis?: any
  series?: ChartSeriesItem[]
  animation?: boolean
  animationThreshold?: number
  animationDuration?: number
  animationEasing?: string
  animationDelay?: number | Function
  animationDurationUpdate?: number
  animationEasingUpdate?: string
  animationDelayUpdate?: number | Function
  [key: string]: any
}

// 轻量级动画配置 - 默认使用
export const lightAnimationConfig = {
  animation: true,
  // 设置较高的动画阈值，只有数据量小时才启用动画
  animationThreshold: 2000,
  // 减少动画持续时间
  animationDuration: 300,
  animationEasing: 'cubicOut',
  // 减少延迟时间
  animationDelay: (idx: number) => Math.max(idx * 10, 0),
  // 更新动画更快
  animationDurationUpdate: 200,
  animationEasingUpdate: 'cubicInOut',
  animationDelayUpdate: (idx: number) => Math.max(idx * 5, 0)
}

// 增长动画预设 - 渐入增长效果
export const fadeInGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1000,
  animationEasing: 'cubicIn', // 先慢后快
  animationDelay: (idx: number) => idx * 100,
  animationDurationUpdate: 500,
  animationEasingUpdate: 'cubicInOut'
}

// 增长动画预设 - 弹性增长效果
export const bounceGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1200,
  animationEasing: 'elasticOut', // 弹性效果
  animationDelay: (idx: number) => idx * 120,
  animationDurationUpdate: 600,
  animationEasingUpdate: 'elasticOut'
}

// 增长动画预设 - 波浪增长效果
export const waveGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1500,
  animationEasing: 'backOut', // 回弹效果
  animationDelay: (idx: number) => idx * 80,
  animationDurationUpdate: 800,
  animationEasingUpdate: 'backOut'
}

// 增长动画预设 - 顺序增长效果
export const sequentialGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 800,
  animationEasing: 'linear',
  // 较长的延迟，使每个元素依次显示
  animationDelay: (idx: number) => idx * 200,
  animationDurationUpdate: 400,
  animationEasingUpdate: 'linear',
  animationDelayUpdate: (idx: number) => idx * 100
}

// 柱状图预设配置
export const barChartPreset: ChartOptions = {
  // 默认使用动画配置
  ...lightAnimationConfig,
  grid: {
    top: '40px',
    left: '3%',
    right: '3%',
    bottom: '8%',
    containLabel: true
  },
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c}',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    borderColor: '#eee',
    borderWidth: 1,
    textStyle: {
      color: '#333'
    },
    shadowBlur: 5,
    shadowColor: 'rgba(0, 0, 0, 0.1)'
  },
  legend: {
    show: false
  },
  xAxis: {
    splitLine: {
      show: false
    },
    axisTick: {
      show: false
    },
    axisLabel: {
      show: true,
      color: 'rgb(97,97,97)',
      fontSize: 14,
      fontWeight: 'bold',
      margin: 8
    },
    axisLine: {
      show: false
    }
  },
  yAxis: {
    type: 'value',
    splitLine: {
      show: true,
      lineStyle: {
        type: 'dashed',
        color: '#E5E7EB'
      }
    },
    axisLabel: {
      show: true,
      color: '#666',
      fontSize: 12,
      margin: 8
    },
    axisLine: {
      show: false
    },
    axisTick: {
      show: false
    }
  },
  series: [
    {
      type: 'bar',
      label: {
        show: true,
        position: 'top',
        fontSize: 16,
        color: 'rgb(97,97,97)'
      },
      itemStyle: {
        borderRadius: [0, 0, 0, 0],
        color: 'rgba(62, 99, 221, 1)'
      },
      barWidth: 32,
      barGap: '30%',
      emphasis: {
        itemStyle: {
          color: 'rgba(62, 99, 221, 0.8)'
        }
      }
    }
  ]
}

// 饼图预设配置
export const pieChartPreset: ChartOptions = {
  // 默认使用动画配置
  ...lightAnimationConfig,
  tooltip: {
    trigger: 'item',
    formatter: '{b}: {c} ({d}%)',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    borderColor: '#eee',
    borderWidth: 1,
    textStyle: {
      color: '#333'
    },
    shadowBlur: 5,
    shadowColor: 'rgba(0, 0, 0, 0.1)'
  },
  legend: {
    orient: 'horizontal',
    bottom: 0,
    left: 'center',
    itemWidth: 10,
    itemHeight: 10,
    itemGap: 24,
    icon: 'circle',
    textStyle: {
      fontSize: 16,
      color: 'rgb(97,97,97)',
      fontWeight: 510
    }
  },
  grid: {
    top: '10%',
    bottom: '15%',
    containLabel: true
  },
  series: [
    {
      type: 'pie',
      radius: ['65%', '80%'],
      center: ['50%', '50%'],
      avoidLabelOverlap: true,
      itemStyle: {
        borderRadius: 0,
        borderColor: '#fff',
        borderWidth: 0
      },
      label: {
        show: false
      },
      emphasis: {
        label: {
          show: false
        },
        itemStyle: {
          shadowBlur: 5,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.2)'
        }
      },
      labelLine: {
        show: false
      },
      color: [
        'rgba(52, 81, 178, 1)', // 最深色
        'rgba(141, 164, 239, 1)', // 次深色
        'rgba(162, 181, 243, 1)', // 减淡1
        'rgba(183, 198, 247, 1)', // 减淡2
        'rgba(204, 215, 251, 1)', // 减淡3
        'rgba(225, 232, 255, 1)', // 减淡4
        'rgba(236, 241, 255, 1)' // 减淡5
      ]
    }
  ]
}

// 漏斗图预设配置
export const funnelChartPreset: ChartOptions = {
  // 默认使用动画配置
  ...lightAnimationConfig,
  tooltip: {
    trigger: 'item',
    formatter: '{b} : {c}',
    backgroundColor: 'rgba(255, 255, 255, 0.9)',
    borderColor: '#eee',
    borderWidth: 1,
    textStyle: {
      color: '#333'
    },
    shadowBlur: 5,
    shadowColor: 'rgba(0, 0, 0, 0.1)'
  },
  legend: {
    show: false // 隐藏图例，因为我们有自定义UI
  },
  title: {
    left: 'left',
    textAlign: 'left',
    padding: [10, 0, 0, 10]
  },
  // 为四个漏斗段设置匹配的颜色
  color: ['#e7edfc', '#91a3e9', '#4662d5', '#3a50ac', '#181d46'],
  series: [
    {
      name: '', // 使用通用名称，会被用户传入的series name覆盖
      type: 'funnel',
      orient: 'horizontal',
      funnelAlign: 'bottom',
      left: '0',
      top: '0',
      bottom: '5%',
      width: '100%',
      height: '100%',
      // 确保小值也能被看到
      min: 0,
      minSize: '1%',
      maxSize: '100%',
      // 不排序，保持原始数据顺序
      sort: 'none',
      // 增加段之间的间隔
      gap: 0,
      // 标签配置
      labelLine: {
        show: false
      }
    }
  ]
}

// 导出所有预设
export const chartPresets = {
  barChart: barChartPreset,
  pieChart: pieChartPreset,
  funnelChart: funnelChartPreset
}

// 导出所有动画预设
export const animationPresets = {
  light: lightAnimationConfig,
  fadeInGrowth: fadeInGrowthAnimation,
  bounceGrowth: bounceGrowthAnimation,
  waveGrowth: waveGrowthAnimation,
  sequentialGrowth: sequentialGrowthAnimation
}

export default chartPresets
