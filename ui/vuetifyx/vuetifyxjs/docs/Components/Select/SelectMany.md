# select 选择器

## vx-selectmany

:::demo

```vue
<script setup>
import { ref } from 'vue'

const value = ref(['1', '2'])
const items = ref([
  {
    id: '1',
    text: 'ScanDa Adams',
    image: 'https://cdn.vuetifyjs.com/images/lists/1.jpg'
  },
  {
    id: '2',
    text: 'Ali Connors',
    image: 'https://cdn.vuetifyjs.com/images/lists/2.jpg'
  },
  {
    id: '3',
    text: 'Ali DE',
    image: 'https://cdn.vuetifyjs.com/images/lists/3.jpg'
  },
  {
    id: '4',
    text: 'Bogn',
    image: 'https://cdn.vuetifyjs.com/images/lists/4.jpg'
  }
])
</script>

<template>
  <p>{{ value }}</p>
  <vx-selectmany v-model="value" :items="items" />
</template>

```
:::


## vx-linkageselect

:::demo

```vue
<script setup lang="ts">
import { Ref, ref } from 'vue'

const value = ref(['2', '3', '6'])
const items: Ref<any> = ref([
  [
    {
      ID: '1',
      Name: '浙江',
      ChildrenIDs: ['1', '2']
    },
    {
      ID: '2',
      Name: '江苏',
      ChildrenIDs: ['3', '4']
    }
  ],
  [
    { ID: '1', Name: '杭州', ChildrenIDs: ['1', '2'] },
    { ID: '2', Name: '宁波', ChildrenIDs: ['3', '4'] },
    { ID: '3', Name: '南京', ChildrenIDs: ['5', '6'] },
    { ID: '4', Name: '苏州', ChildrenIDs: ['7', '8'] }
  ],
  [
    { ID: '1', Name: '拱墅区' },
    { ID: '2', Name: '西湖区' },
    { ID: '3', Name: '镇海区' },
    { ID: '4', Name: '鄞州区' },
    { ID: '5', Name: '鼓楼区' },
    { ID: '6', Name: '玄武区' },
    { ID: '7', Name: '常熟区' },
    { ID: '8', Name: '吴江区' }
  ]
])
const labels = ref(['Province', 'City', 'District'])
</script>

<template>
  <p>{{ value }}</p>
  <vx-linkageselect
    v-model="value"
    :items="items"
    :labels="labels"
    select-out-of-order
  ></vx-linkageselect>
</template>

<style scoped></style>


```
:::