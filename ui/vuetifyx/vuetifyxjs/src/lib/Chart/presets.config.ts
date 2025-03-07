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
  title: {
    text: ''
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
      color: '#666',
      fontSize: 14,
      fontWeight: 'bold',
      margin: 12
    },
    axisLine: {
      show: false
    }
  },
  yAxis: {
    type: 'value',
    splitLine: {
      show: false
    },
    axisLabel: {
      show: false
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
        fontSize: 14,
        color: '#666'
      },
      itemStyle: {
        borderRadius: [8, 8, 8, 8],
        color: 'rgba(62, 99, 221, 1)'
      },
      barWidth: 32,
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
  title: {
    text: ''
  },
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
    bottom: 10,
    left: 'center',
    itemWidth: 10,
    itemHeight: 10,
    textStyle: {
      fontSize: 12,
      color: '#666'
    }
  },
  series: [
    {
      type: 'pie',
      radius: ['55%', '70%'],
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
          color: 'rgba(62, 99, 221, 1)',
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      labelLine: {
        show: false
      },
      color: [
        'rgba(230, 237, 254, 1)',
        'rgba(200, 216, 248, 1)',
        'rgba(170, 195, 242, 1)',
        'rgba(140, 174, 236, 1)',
        'rgba(110, 153, 230, 1)'
      ]
    }
  ]
}

// 导出所有预设
export const chartPresets = {
  barChart: barChartPreset,
  pieChart: pieChartPreset
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
