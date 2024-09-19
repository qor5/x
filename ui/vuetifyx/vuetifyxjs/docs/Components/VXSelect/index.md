# vx-select 选择器

## vx-select

:::demo

```vue
<template>
  <vx-select
    type="autocomplete"
    v-model="valueAutoComplete"
    multiple
    chips
    clearable
    label="autoComplete Select"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    v-model="valueNormal"
    label="Normal Select"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    v-model="valueWithErrorMsg"
    label="Select with Error messages"
    error-messages="This is an error message explanation"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <p>selected value: {{ value }}</p>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const valueAutoComplete = ref([1, 2, 3])
const valueNormal = ref([1])
const valueWithErrorMsg = ref([1])
const srcs = {
  1: 'https://cdn.vuetifyjs.com/images/lists/1.jpg',
  2: 'https://cdn.vuetifyjs.com/images/lists/2.jpg',
  3: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
  4: 'https://cdn.vuetifyjs.com/images/lists/4.jpg',
  5: 'https://cdn.vuetifyjs.com/images/lists/5.jpg'
}
const items = ref([
  { id: 1, name: 'Sandra Adams', group: 'Group 1', avatar: srcs[1] },
  { id: 2, name: 'Ali Connors', group: 'Group 1', avatar: srcs[2] },
  { id: 3, name: 'Trevor Hansen', group: 'Group 1', avatar: srcs[3] },
  { id: 4, name: 'Tucker Smith', group: 'Group 1', avatar: srcs[2] },
  { id: 5, name: 'Britta Holt', group: 'Group 2', avatar: srcs[4] },
  { id: 6, name: 'Jane Smith ', group: 'Group 2', avatar: srcs[5] },
  { id: 7, name: 'John Smith', group: 'Group 2', avatar: srcs[1] },
  { id: 8, name: 'Sandra Williams', group: 'Group 2', avatar: srcs[3] }
])
</script>

<style scoped></style>
```

:::

## vx-selectmany

> legacy component

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

> legacy component

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

## vx-autocomplete

> legacy component

:::demo

```vue
<script setup lang="ts">
import { reactive, ref } from 'vue'
import VueJsonPretty from 'vue-json-pretty'

const remote = reactive({
  pageSize: 5,
  page: 1,
  search: ''
})

const getItems = () => {
  const items = []
  for (let i = 1; i <= remote.pageSize; i++) {
    items.push({
      icon: `https://cdn.vuetifyjs.com/images/lists/${i}.jpg`,
      text: `test_${remote.page}_${i}`,
      value: (remote.pageSize * (remote.page - 1) + i).toFixed()
    })
  }
  return items
}
const loadData = (): Promise<any> => {
  return new Promise((resolve, reject) => {
    resolve({
      data: {
        pages: 4,
        total: 20,
        current: remote.page * remote.pageSize,
        items: getItems()
      }
    })
  })
}
// const items = ref( [{ 'text': '高节', 'value': '1' }, { 'text': '地界', 'value': '3' }],)
const items = ref(getItems())
const value = ref()
</script>
<template>
  <vue-json-pretty :data="value"></vue-json-pretty>
  <!--  <vx-autocomplete-->
  <!--    v-model="value"-->
  <!--    :items="items"-->
  <!--  ></vx-autocomplete>-->
  <vx-autocomplete
    sorting
    :items="items"
    has-icon
    :remote="remote"
    v-model="value"
  ></vx-autocomplete>
</template>

<style scoped></style>
```

:::
