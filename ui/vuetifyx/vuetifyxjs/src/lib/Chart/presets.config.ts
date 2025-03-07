// 定义图表选项类型
interface ChartSeriesItem {
  type?: string
  name?: string
  data?: any[]
  radius?: string | string[]
  [key: string]: any
}

interface ChartOptions {
  tooltip?: any
  grid?: any
  xAxis?: any
  yAxis?: any
  series?: ChartSeriesItem[]
  [key: string]: any
}

// 柱状图预设配置
export const barChartPreset: ChartOptions = {
  title: {
    text: ''
  },
  // 优化tooltip提示框
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
  // 隐藏图例
  legend: {
    show: false
  },
  xAxis: {
    // 去掉x轴的分隔线
    splitLine: {
      show: false
    },
    // 隐藏x轴刻度线
    axisTick: {
      show: false
    },
    // 显示x轴标签
    axisLabel: {
      show: true,
      color: '#666',
      fontSize: 14,
      fontWeight: 'bold',
      margin: 12
    },
    // 隐藏x轴线
    axisLine: {
      show: false
    }
  },
  yAxis: {
    type: 'value', // 明确指定y轴类型
    // 去掉y轴的分隔线
    splitLine: {
      show: false
    },
    // 隐藏y轴刻度
    axisLabel: {
      show: false
    },
    // 隐藏y轴线
    axisLine: {
      show: false
    },
    // 隐藏y轴刻度线
    axisTick: {
      show: false
    }
  },
  series: [
    {
      type: 'bar',
      // 在柱子顶部显示数值
      label: {
        show: true,
        position: 'top',
        fontSize: 14,
        color: '#666'
      },
      // 调整柱子样式
      itemStyle: {
        borderRadius: [8, 8, 8, 8], // 柱子顶部和底部都有8px圆角
        color: 'rgba(62, 99, 221, 1)' // 柱子颜色为蓝色
      },
      // 设置柱子宽度为固定的32px
      barWidth: 32,
      // 添加鼠标悬停效果
      emphasis: {
        itemStyle: {
          color: 'rgba(62, 99, 221, 0.8)' // 悬停时颜色稍微变淡
        }
      }
    }
  ]
}

// 饼图预设配置
export const pieChartPreset: ChartOptions = {
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
      radius: ['55%', '70%'], // 进一步调整环形图内外半径，使环形更瘦
      center: ['50%', '50%'],
      avoidLabelOverlap: true,
      itemStyle: {
        borderRadius: 0, // 移除圆角，使扇区之间没有间隔
        borderColor: '#fff',
        borderWidth: 0 // 移除边框，消除扇区间隔
      },
      label: {
        show: false // 隐藏标签
      },
      emphasis: {
        label: {
          show: false // 高亮时也不显示标签
        },
        itemStyle: {
          color: 'rgba(62, 99, 221, 1)', // 与柱状图相同的高亮颜色
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      labelLine: {
        show: false // 隐藏引导线
      },
      // 默认颜色配置
      color: [
        'rgba(230, 237, 254, 1)', // 最浅的蓝色
        'rgba(200, 216, 248, 1)', // 稍微深一点的蓝色
        'rgba(170, 195, 242, 1)', // 中等深度的蓝色
        'rgba(140, 174, 236, 1)', // 较深的蓝色
        'rgba(110, 153, 230, 1)' // 最深的蓝色
      ]
    }
  ]
}

// 导出所有预设
export const chartPresets = {
  barChart: barChartPreset,
  pieChart: pieChartPreset
}

export default chartPresets
