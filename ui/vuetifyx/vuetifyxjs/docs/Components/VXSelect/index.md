# vx-select 选择器

## vx-select

### 基本用法

```html
<vx-select model-value="Hello World" label="field1" />
```

:::demo

```vue
<template>
  <vx-select
    type="autocomplete"
    v-model="value"
    multiple
    chips
    clearable
    label="autoComplete Select"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <p>selected value: {{ value }}</p>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const value = ref([1,2,3])
const srcs = {
  1: 'https://cdn.vuetifyjs.com/images/lists/1.jpg',
  2: 'https://cdn.vuetifyjs.com/images/lists/2.jpg',
  3: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
  4: 'https://cdn.vuetifyjs.com/images/lists/4.jpg',
  5: 'https://cdn.vuetifyjs.com/images/lists/5.jpg'
}
const items = ref([
  { id:1, name: 'Sandra Adams', group: 'Group 1', avatar: srcs[1] },
  { id:2, name: 'Ali Connors', group: 'Group 1', avatar: srcs[2] },
  { id:3, name: 'Trevor Hansen', group: 'Group 1', avatar: srcs[3] },
  { id:4, name: 'Tucker Smith', group: 'Group 1', avatar: srcs[2] },
  { id:5, name: 'Britta Holt', group: 'Group 2', avatar: srcs[4] },
  { id:6, name: 'Jane Smith ', group: 'Group 2', avatar: srcs[5] },
  { id:7, name: 'John Smith', group: 'Group 2', avatar: srcs[1] },
  { id:8, name: 'Sandra Williams', group: 'Group 2', avatar: srcs[3] }
])
</script>

<style scoped></style>
```

:::

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
