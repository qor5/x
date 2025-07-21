# vx-chart 图表组件

基于ECharts封装的图表组件，提供了常用的图表预设和配置选项。

## API

### Props

| 参数名               | 说明                                                                                                        | 类型              | 默认值   |
| -------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------- | -------- |
| presets              | 预设样式，可选值：'barChart'、'pieChart'、'funnelChart'                                                     | String            | ''       |
| options              | 图表配置项，会与预设样式合并                                                                                | Object \ Object[] | {}       |
| height               | 设置图表高度                                                                                                | String            | 'auto'   |
| mergeOptionsCallback | 可以使用这个回调来修改当前的配置参数, 当需要自定义vx-chart配置的时候格外有用，详见 [#饼图示例](./#饼图示例) | Function          | () => {} |
| loading              | 是否显示加载状态                                                                                            | Boolean           | false    |

### Slots

| 名称        | 说明                                     | 插槽 Props                                                                                                        |
| ----------- | ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| title       | 支持自定义标题                           | currentIndex: number（当前索引）                                                                                  |
| description | 图表标题和图表之间的区域                 | currentIndex: number（当前索引）                                                                                  |
| action      | 图表操作区域，通常用于切换不同的图表配置 | list: number[]（可用索引列表）<br>currentIndex: number（当前索引）<br>toggle: (index: number) => void（切换函数） |

## 预设类型

VXChart 组件提供了三种预设类型，可以通过 `presets` 属性指定：

- `barChart`：柱状图预设，适用于展示分类数据的数量对比
- `pieChart`：饼图预设，适用于展示占比数据
- `funnelChart`：漏斗图预设，适用于展示转化数据，**支持无限多列的智能缩放**

## 基础示例

### 柱状图示例

使用 `barChart` 预设可以快速创建美观的柱状图：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const barChartData = ref({
  title: {
    text: '年龄分布'
  },
  xAxis: {
    data: ['0-18', '18-25', '25-65', '65+']
  },
  series: [
    {
      name: '人数',
      data: [100, 300, 500, 200]
    }
  ]
})
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="barChart" :options="barChartData"></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

### 饼图示例

使用 `pieChart` 预设可以快速创建美观的饼图：

当你想自定义图例时可以使用 `mergeOptionsCallback`, 回调函数支持两个参数

- options - 当前图表配置项
- data - 一些图表数据对象

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const pieChartData = ref({
  title: {
    text: '性别比例'
  },
  series: [
    {
      name: '性别分布',
      data: [
        {
          value: 10.111,
          name: '男性'
        },
        {
          value: 89.889,
          name: '女性'
        }
      ]
    }
  ]
})

const mergeOptionsCallback = function (options, { seriesData }) {
  options.legend = {
    ...options.legend,
    formatter: (name) => {
      const item = seriesData.find((i) => i.name === name)
      const percent = ((item.value / 100) * 100).toFixed(2)
      return `${name} ${percent}%`
    }
  }
}
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart
      presets="pieChart"
      :options="pieChartData"
      :merge-options-callback="mergeOptionsCallback"
    ></vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

### 漏斗图

使用 `funnelChart` 预设可以快速创建美观的漏斗图，用于展示转化流程和各环节的数据。以下展示最简单的用法

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const funnelChartData = ref([
  {
    title: {
      text: 'Event Funnel Chart(Past 7 days)'
    },
    series: [
      {
        name: '邮件营销',
        data: [
          {
            value: 32502,
            name: 'Email Sent',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '1000'
                }
              ]
            }
          },
          {
            value: 1000,
            name: 'Email Delivered',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '8,500'
                }
              ]
            }
          },
          {
            value: 50,
            name: 'Link Clicked',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                }
              ]
            }
          },
          {
            value: 100,
            name: 'Link Clicked1',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                }
              ]
            }
          }
        ]
      }
    ]
  },
  {
    title: {
      text: 'Event Funnel Chart(Past 14 days)'
    },
    series: [
      {
        name: '邮件营销',
        data: [
          {
            value: 500,
            name: 'Email Sent',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '1000'
                }
              ]
            }
          },
          {
            value: 200,
            name: 'Email Delivered',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '8,500'
                }
              ]
            }
          },
          {
            value: 50,
            name: 'Link Clicked',
            extraData: {
              style: 'plain',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                }
              ]
            }
          }
        ]
      }
    ]
  }
])
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="funnelChart" :options="funnelChartData" height="380">
      <template #action="{ list, currentIndex, toggle }">
        <div
          class="d-flex align-center bg-grey-lighten-3 rounded pa-1 mr-4 mt-4"
          style="height: 32px;"
        >
          <button
            v-for="(_, idx) in list"
            :key="idx"
            class="text-body-2 text-no-wrap border-0 flex-grow-1 d-flex align-center justify-center rounded px-2"
            style="height: 24px; cursor: pointer; transition: all 0.3s;"
            :style="
              currentIndex === idx
                ? 'background-color: #fff; color: #4a4a4a;'
                : 'background-color: transparent; color: rgb(117, 117, 117);'
            "
            @click="toggle(idx)"
          >
            {{ idx === 0 ? 'Past 7 Days' : 'Past 14 Days' }}
          </button>
        </div>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

#### 漏斗图进阶配置

使用 `funnelChart` 预设可以快速创建美观的漏斗图，用于展示转化流程和各环节的数据。**新版本支持无限多列的智能缩放算法**，能够根据列数和容器宽度自动调整元素大小和布局：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

const addtionalDataList = ref([
  {
    name: 'Dropped',
    icon: 'mdi-cancel',
    valueStr: '89,935'
  },
  {
    name: 'Aborted',
    icon: 'mdi-close-octagon-outline',
    tips: 'it is tips',
    valueStr: '89,935'
  },
  {
    name: 'Bounced',
    icon: 'mdi-lock-reset',
    valueStr: '89,935'
  },
  {
    name: 'Complaint',
    icon: 'mdi-emoticon-sad-outline',
    valueStr: '89,935'
  }
])

const funnelChartData = ref([
  {
    title: {
      text: '邮件营销漏斗'
    },
    series: [
      {
        name: '邮件营销',
        data: [
          {
            value: 1000,
            name: 'Email Sent',
            extraData: {
              icon: 'mdi-near-me',
              hideLabel: true,
              labelList: [
                {
                  type: 'primary',
                  text: '1000'
                }
              ]
            }
          },
          {
            value: 800,
            name: 'Email Delivered',
            tooltip:
              'SES delivery events confirm SMTP-level acceptance, but do not guarantee inbox placement or user visibility tooltip',
            extraData: {
              icon: 'mdi-email-mark-as-unread',
              hideLabel: true,
              labelList: [
                {
                  type: 'primary',
                  text: '8,500'
                }
              ]
            }
          },
          {
            value: 400,
            name: 'Email Opened',
            extraData: {
              icon: 'mdi-check-all',
              hideLabel: true,
              labelList: [
                {
                  type: 'primary',
                  text: '5,000'
                }
              ]
            }
          },
          {
            value: 200,
            name: 'Link Clicked',
            hideLabel: true,
            extraData: {
              icon: 'mdi-link',
              hideLabel: true,
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                }
              ]
            }
          },
          {
            value: 100,
            name: 'Link Clicked2',
            extraData: {
              icon: 'mdi-link',
              hideLabel: true,
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                }
              ]
            }
          }
        ]
      }
    ]
  },
  {
    title: {
      text: '邮件营销漏斗 - 带折线图'
    },
    series: [
      {
        name: '邮件营销',
        isDisabled: true,
        type: 'funnel', // 明确指定为漏斗图
        data: [
          {
            value: 1200,
            name: 'Email Sent',
            extraData: {
              icon: 'mdi-near-me',
              labelList: [
                {
                  labelName: 'This Week',
                  type: 'primary',
                  text: '1000'
                },
                {
                  labelName: 'Last Week',
                  type: 'primary',
                  text: '4000',
                  textStyle: 'color: #9e9e9e;'
                }
              ]
            }
          },
          {
            value: 700,
            name: 'Email Delivered',
            extraData: {
              icon: 'mdi-email-mark-as-unread',
              labelList: [
                {
                  type: 'primary',
                  text: '8,500'
                },
                {
                  type: 'primary',
                  text: '4000'
                }
              ]
            }
          },
          {
            value: 300,
            name: 'Email Opened',
            extraData: {
              icon: 'mdi-check-all',
              labelList: [
                {
                  type: 'primary',
                  text: '5,000'
                },
                {
                  type: 'primary',
                  text: '5,000'
                }
              ]
            }
          },
          {
            value: 100,
            name: 'Link Clicked',
            extraData: {
              icon: 'mdi-link',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                },
                {
                  type: 'secondary',
                  icon: 'mdi-arrow-top-right',
                  text: '+1.01%'
                },
                {
                  type: 'primary',
                  text: '3000'
                }
              ]
            }
          },
          {
            value: 50,
            name: 'Link Clicked2',
            extraData: {
              icon: 'mdi-link',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                },
                {
                  type: 'secondary',
                  icon: 'mdi-arrow-top-right',
                  text: '+1.01%'
                },
                {
                  type: 'primary',
                  text: '2000'
                }
              ]
            }
          },
          {
            value: 50,
            name: 'Link Clicked2',
            extraData: {
              icon: 'mdi-link',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                },
                {
                  type: 'secondary',
                  icon: 'mdi-arrow-top-right',
                  text: '+1.01%'
                },
                {
                  type: 'primary',
                  text: '2000'
                }
              ]
            }
          },
          {
            value: 50,
            name: 'Link Clicked2',
            extraData: {
              icon: 'mdi-link',
              labelList: [
                {
                  type: 'primary',
                  text: '2,500'
                },
                {
                  type: 'secondary',
                  icon: 'mdi-arrow-top-right',
                  text: '+1.01%'
                },
                {
                  type: 'primary',
                  text: '2000'
                }
              ]
            }
          }
        ]
      },
      {
        name: '转化率趋势',
        type: 'line',
        data: [120, 132, 101, 134, 90, 70, 100, 50],
        lineColor: 'green',
        smooth: false // 可选：是否平滑曲线
      }
    ]
  }
])

// 通过mergeOptionsCallback控制漏斗图颜色
const funnelMergeOptionsCallback = (options, { currentIndex }) => {
  // 当切换到第二个配置(idx = 1)时，将所有漏斗段设置为浅蓝色
  if (currentIndex === 1 && options.series && options.series[0] && options.series[0].data) {
    console.log(options)
    options.series[0].data = options.series[0].data.map((item) => ({
      ...item,
      itemStyle: {
        color: '#e6edfe'
      }
    }))
  }
}
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart
      presets="funnelChart"
      :options="funnelChartData"
      :merge-options-callback="funnelMergeOptionsCallback"
    >
      <template #title="{ currentIndex }"
        ><div class="d-flex align-center" style="font-size:35px">
          Campaign Name
          <span
            v-if="currentIndex === 1"
            class="rounded pa-1 ml-4 text-caption"
            style="font-size:16px;background:#eee;font-weight:400;line-height:1;"
            >Weekly</span
          >
        </div></template
      >
      <template #description="{ currentIndex }">
        <div class="mt-6 ml-3">
          <span
            v-if="currentIndex === 0"
            class="text-caption mr-4 px-1 py-1 rounded"
            style="background:#F5F5F5;"
            >Last Updated: 0:05 25/05/09; Data will be updated at 0:05
          </span>

          <div v-else class="d-flex align-center">
            <span
              class="d-inline-flex align-center text-caption mr-2 px-1 py-1 rounded"
              style="background:#F5F5F5;"
              ><i
                class="d-inline-block rounded mr-2"
                style="width:12px; height:12px; background: #e6edfe"
              /><b class="mr-2" style="color:#616161;">This Week</b>
              25/05/01-25/05/07
            </span>
            <span
              class="d-inline-flex align-center text-caption mr-4 px-1 py-1 rounded"
              style="background:#F5F5F5;"
              ><i class="d-inline-block mr-2" style="width:12px;height:2px;background:#3e63dd" /><b
                class="mr-2"
                style="color:#616161;"
                >Last Week</b
              >
              25/05/01-25/05/07
            </span>
          </div>
        </div>
      </template>

      <template #action="{ list, currentIndex, toggle }">
        <div
          class="d-flex align-center bg-grey-lighten-3 rounded pa-1 mr-4 mt-4"
          style="height: 32px;"
        >
          <button
            v-for="(_, idx) in list"
            :key="idx"
            class="text-body-2 rounded text-no-wrap border-0 flex-grow-1 d-flex align-center justify-center rounded px-2"
            style="height: 24px; cursor: pointer; transition: all 0.3s;"
            :style="
              currentIndex === idx
                ? 'background-color: #fff; color: #4a4a4a;'
                : 'background-color: transparent; color: rgb(117, 117, 117);'
            "
            @click="toggle(idx)"
          >
            {{ idx === 0 ? 'Summary' : 'By Week' }}
          </button>
        </div>
      </template>
    </vx-chart>

    <div class="mt-4">
      <div class="d-flex justify-space-between w-100 ga-4">
        <div v-for="(item, i) in addtionalDataList" class="border pa-3 rounded-lg" style="flex:1;">
          <div
            class="d-flex border pa-2 rounded-lg justify-space-between align-center"
            style="background: #f9e6e4;border-color:#eb9091!important;"
          >
            <vx-label :tooltip="item.tips" tooltip-icon-color="error">
              <span style="color: #e6484e;">{{ item.name }}</span>
            </vx-label>

            <div
              class="d-flex rounded-lg justify-center align-center"
              style="background:#fff;width:32px; height:32px;"
            >
              <v-icon :icon="item.icon" size="16" color="error" />
            </div>
          </div>

          <div v-if="i === 0" class="mt-8 text-bold " style="font-size: 20px;font-weight: 510;">
            <div class="pb-3">
              <div style="font-size:12px; font-weight:510;color: #616161">This Week</div>
              <div>89,935</div>
            </div>
            <div>
              <div style="font-size:12px; font-weight:510;color: #616161">Last Week</div>
              <div style="color: #9e9e9e">89,935</div>
            </div>
          </div>

          <div v-else class="mt-8 text-bold" style="font-size: 24px;font-weight: 510;">
            {{ item.valueStr }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;
}
</style>
```

:::

## 功能扩展

### 多图表切换

使用 `options` 数组和 `action` 插槽可以实现多图表切换功能：

:::demo

```vue
<script setup lang="ts">
import { ref } from 'vue'

// 定义7天和14天的数据
const chartData = ref([
  {
    title: {
      text: 'Daily Active Users (7 Days)'
    },
    xAxis: {
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
    },
    series: [
      {
        type: 'bar',
        name: '用户数',
        data: [5, 20, 36, 10, 10, 20, 30]
      }
    ]
  },
  {
    title: {
      text: 'Daily Active Users (14 Days)'
    },
    xAxis: {
      data: [
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 1',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2',
        'Week 2'
      ]
    },
    series: [
      {
        type: 'bar',
        name: '用户数',
        data: [5, 20, 36, 10, 10, 20, 30, 15, 25, 40, 20, 15, 25, 35]
      }
    ]
  }
])
</script>
<template>
  <div class="chart-container border border-gray-500 rounded-lg">
    <vx-chart presets="barChart" :options="chartData">
      <template #action="{ list, currentIndex, toggle }">
        <div
          class="d-flex align-center bg-grey-lighten-3 rounded pa-1 mr-4 mt-4"
          style="height: 32px;"
        >
          <button
            v-for="(_, idx) in list"
            :key="idx"
            class="text-body-2 rounded text-no-wrap border-0 flex-grow-1 d-flex align-center justify-center rounded px-2"
            style="height: 24px; cursor: pointer; transition: all 0.3s;"
            :style="
              currentIndex === idx
                ? 'background-color: #fff; color: #4a4a4a;'
                : 'background-color: transparent; color: rgb(117, 117, 117);'
            "
            @click="toggle(idx)"
          >
            {{ idx === 0 ? 'Past 7 Days' : 'Past 14 Days' }}
          </button>
        </div>
      </template>
    </vx-chart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 100%;

  position: relative;
}
</style>
```

:::

## 使用说明

### 基本使用

只需传入数据，使用预设样式：

```vue
<vx-chart
  presets="barChart"
  :options="{
    xAxis: { data: ['A', 'B', 'C'] },
    series: [{ data: [10, 20, 30] }]
  }"
></vx-chart>
```

### 完全自定义

不使用预设，完全自定义配置：

```vue
<vx-chart :options="customOptions"></vx-chart>
```

### 混合使用

使用预设，但覆盖部分配置：

```vue
<vx-chart
  presets="pieChart"
  :options="{
    title: { text: '自定义标题' },
    series: [{ data: customData }]
  }"
></vx-chart>
```

自定义漏斗图示例:

```vue
<vx-chart
  presets="funnelChart"
  :options="{
    title: { text: '自定义漏斗图' },
    tooltip: {
      formatter: '{b}: {c}人 ({d}%)' // 自定义提示格式
    },
    color: ['#FF6B6B', '#FFD166', '#06D6A0', '#118AB2', '#073B4C'], // 自定义颜色
    series: [
      {
        name: '用户行为',
        data: myFunnelData,
        label: {
          position: 'right' // 将标签放在右侧（默认在左侧）
        }
      }
    ]
  }"
></vx-chart>
```

### 使用加载状态

```vue
<vx-chart presets="barChart" :options="chartData" :loading="isLoading"></vx-chart>
```
