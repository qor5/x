// Define chart options types
export interface ChartSeriesItem {
  type?: string
  name?: string
  data?: any[]
  radius?: string | string[]
  isDisabled?: boolean
  lineColor?: string
  smooth?: boolean
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

// Lightweight animation configuration - used by default
export const lightAnimationConfig = {
  animation: true,
  // Set higher animation threshold, only enable animation when data size is small
  animationThreshold: 2000,
  // Reduce animation duration
  animationDuration: 300,
  animationEasing: 'cubicOut',
  // Reduce delay time
  animationDelay: (idx: number) => Math.max(idx * 10, 0),
  // Faster update animation
  animationDurationUpdate: 200,
  animationEasingUpdate: 'cubicInOut',
  animationDelayUpdate: (idx: number) => Math.max(idx * 5, 0)
}

// Growth animation preset - fade-in growth effect
export const fadeInGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1000,
  animationEasing: 'cubicIn', // slow first then fast
  animationDelay: (idx: number) => idx * 100,
  animationDurationUpdate: 500,
  animationEasingUpdate: 'cubicInOut'
}

// Growth animation preset - elastic growth effect
export const bounceGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1200,
  animationEasing: 'elasticOut', // elastic effect
  animationDelay: (idx: number) => idx * 120,
  animationDurationUpdate: 600,
  animationEasingUpdate: 'elasticOut'
}

// Growth animation preset - wave growth effect
export const waveGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 1500,
  animationEasing: 'backOut', // bounce back effect
  animationDelay: (idx: number) => idx * 80,
  animationDurationUpdate: 800,
  animationEasingUpdate: 'backOut'
}

// Growth animation preset - sequential growth effect
export const sequentialGrowthAnimation = {
  animation: true,
  animationThreshold: 5000,
  animationDuration: 800,
  animationEasing: 'linear',
  // Longer delay to make each element display sequentially
  animationDelay: (idx: number) => idx * 200,
  animationDurationUpdate: 400,
  animationEasingUpdate: 'linear',
  animationDelayUpdate: (idx: number) => idx * 100
}

// Preset configuration for line chart in funnel chart
export const presetsLineInFunnel: ChartOptions = {
  backgroundColor: 'transparent',
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
  grid: {
    left: '0%',
    right: '0%',
    top: '0%',
    bottom: '0%',
    containLabel: false
  },
  xAxis: {
    type: 'value',
    show: false,
    boundaryGap: false
  },
  yAxis: {
    type: 'value',
    show: false,
    min: 0
  },
  series: [
    {
      type: 'line',
      smooth: false,
      symbol: 'none',
      symbolSize: 0,
      showSymbol: false,
      connectNulls: false,
      lineStyle: {
        width: 2,
        color: '#3e63dd'
      },
      emphasis: {
        lineStyle: {
          width: 3
        }
      }
    }
  ]
}

// Bar chart preset configuration
export const barChartPreset: ChartOptions = {
  // Use animation configuration by default
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

// Pie chart preset configuration
export const pieChartPreset: ChartOptions = {
  // Use animation configuration by default
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
        'rgba(52, 81, 178, 1)', // Darkest color
        'rgba(141, 164, 239, 1)', // Second darkest color
        'rgba(162, 181, 243, 1)', // Lighten 1
        'rgba(183, 198, 247, 1)', // Lighten 2
        'rgba(204, 215, 251, 1)', // Lighten 3
        'rgba(225, 232, 255, 1)', // Lighten 4
        'rgba(236, 241, 255, 1)' // Lighten 5
      ]
    }
  ]
}

// Funnel chart preset configuration
export const funnelChartPreset: ChartOptions = {
  // Use animation configuration by default
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
    show: false // Hide legend because we have custom UI
  },
  title: {
    left: 'left',
    textAlign: 'left',
    padding: [10, 0, 0, 10]
  },
  // Set matching colors for four funnel segments
  color: ['#e7edfc', '#91a3e9', '#4662d5', '#3a50ac', '#181d46'],
  series: [
    {
      name: '', // Use generic name, will be overridden by user-provided series name
      type: 'funnel',
      orient: 'horizontal',
      funnelAlign: 'bottom',
      left: '0',
      top: '0',
      bottom: '5%',
      width: '100%',
      height: '100%',
      // Ensure small values can be seen
      min: 0,
      minSize: '1%',
      maxSize: '100%',
      // No sorting, keep original data order
      sort: 'none',
      // Increase spacing between segments
      gap: 0,
      // Label configuration
      labelLine: {
        show: false
      }
    }
  ]
}

// Export all presets
export const chartPresets = {
  barChart: barChartPreset,
  pieChart: pieChartPreset,
  funnelChart: funnelChartPreset,
  lineInFunnel: presetsLineInFunnel
}

// Export all animation presets
export const animationPresets = {
  light: lightAnimationConfig,
  fadeInGrowth: fadeInGrowthAnimation,
  bounceGrowth: bounceGrowthAnimation,
  waveGrowth: waveGrowthAnimation,
  sequentialGrowth: sequentialGrowthAnimation
}

export default chartPresets
