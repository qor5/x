# vx-select 选择器

## API

### Slot

#### v-slot:prepend-inner

可以自定义input文字前固定位置的图标的插槽, `selectedItems` 可以用来控制当前选中项

slot scope

```js
{
  isActive: Ref<boolean>
  isFocused: Ref<boolean>
  controlRef: Ref<HTMLElement | undefined>
  selectedItems: Array<{string, any}>
  focus: () => void
  blur: () => void
}
```

#### v-slot:item

每一项的插槽，用来自定义渲染每一项，根元素必须使用 `v-list-item` 及绑定上 `v-bind="props"`

slot scope

```js
{ item: ListItem; index: number; props: Record<string, unknown> }
```

见 [example](#slot-item)

## 示例（vx-select）

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
    closable-chips
  />

  <vx-select
    type="autocomplete"
    multiple
    required
    :rules="[(value) => (value && value.length > 0) || 'At least select a Item']"
    clearable
    label="autocomplete Select + required rules"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    type="autocomplete"
    v-model="valueAutoComplete"
    multiple
    chips
    clearable
    disabled
    label="disabled autocomplete Select"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    type="autocomplete"
    v-model="valueAutoComplete"
    multiple
    chips
    clearable
    error-messages="error message"
    label="autoComplete Select(state with error)"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
    closable-chips
  >
  </vx-select>

  <vx-select
    v-model="valueNormal"
    label="Normal Select"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  >
  </vx-select>

  <vx-select
    label="Normal Select + required rule"
    :items="items"
    required
    :rules="[(value) => !!value || 'At least select a Item']"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    v-model="valueWithErrorMsg"
    label="Select with Error messages"
    :error-messages="['This is an error message explanation']"
    :items="items"
    clearable
    item-title="name"
    item-value="id"
    placeholder="choose a item"
  />

  <vx-select
    v-model="valueWithErrorMsg"
    label="vx-select multiple + closable-chips"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
    multiple
    chips
    closable-chips
    clearable
  />

  <vx-select
    v-model="valueWithErrorMsg"
    label="vx-select disabled"
    :items="items"
    item-title="name"
    item-value="id"
    placeholder="choose a item"
    multiple
    disabled
    chips
    closable-chips
    clearable
  />
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
  { id: 1, name: 'Sandra Adams', group: 'Group 1', avatar: srcs[1], icon: 'mdi-wifi' },
  { id: 2, name: 'Ali Connors', group: 'Group 1', avatar: srcs[2], icon: 'mdi-wifi' },
  { id: 3, name: 'Trevor Hansen', group: 'Group 1', avatar: srcs[3], icon: 'mdi-wifi' },
  { id: 4, name: 'Tucker Smith', group: 'Group 1', avatar: srcs[2], icon: 'mdi-wifi' },
  { id: 5, name: 'Britta Holt', group: 'Group 2', avatar: srcs[4], icon: 'mdi-wifi' },
  { id: 6, name: 'Jane Smith ', group: 'Group 2', avatar: srcs[5], icon: 'mdi-wifi' },
  { id: 7, name: 'John Smith', group: 'Group 2', avatar: srcs[1], icon: 'mdi-wifi' },
  { id: 8, name: 'Sandra Williams', group: 'Group 2', avatar: srcs[3], icon: 'mdi-wifi' }
])
</script>

<style scoped></style>
```

:::

### Slot（item）

item 的原始数据在 item.raw 里， 如果不希望 prepend 的区域i元素变暗，使用`style: "--v-medium-emphasis-opacity:1"` 来控制

:::demo

```vue
<template>
  <div class="mb-4">1. item with prepend element + prepend inner element</div>

  <v-row>
    <v-col cols="6">
      <vx-select
        type="autocomplete"
        v-model="valueAutoComplete"
        label="autoComplete Select"
        :items="items"
        item-title="name"
        item-value="id"
        placeholder="choose a item"
        closable-chips
      >
        <template v-slot:prepend-inner="{ selectedItems }">
          <v-icon :icon="selectedItems[0].icon" style="--v-medium-emphasis-opacity:1" />
        </template>

        <template v-slot:item="{ props, item }">
          <v-list-item v-bind="props" :title="item.title">
            <template v-slot:prepend>
              <v-icon :icon="item.raw.icon" style="--v-medium-emphasis-opacity:1" />
            </template>
          </v-list-item>
        </template> </vx-select
    ></v-col>
    <v-col cols="6">
      <vx-select
        v-model="valueNormal"
        label="Normal Select"
        :items="items"
        item-title="name"
        item-value="id"
        placeholder="choose a item"
      >
        <template v-slot:prepend-inner="{ selectedItems }">
          <v-icon :icon="selectedItems[0].icon" />
        </template>

        <template v-slot:item="{ props, item }">
          <v-list-item v-bind="props" :title="item.title">
            <template v-slot:prepend>
              <v-icon :icon="item.raw.icon" />
            </template>
          </v-list-item>
        </template>
      </vx-select>
    </v-col>
  </v-row>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const valueAutoComplete = ref([1])
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
  { id: 1, name: 'Sandra Adams', group: 'Group 1', avatar: srcs[1], icon: 'mdi-wifi' },
  { id: 2, name: 'Ali Connors', group: 'Group 1', avatar: srcs[2], icon: 'mdi-plus' },
  { id: 3, name: 'Trevor Hansen', group: 'Group 1', avatar: srcs[3], icon: 'mdi-information' },
  { id: 4, name: 'Tucker Smith', group: 'Group 1', avatar: srcs[2], icon: 'mdi-alert' },
  { id: 5, name: 'Britta Holt', group: 'Group 2', avatar: srcs[4], icon: 'mdi-alert-circle' },
  { id: 6, name: 'Jane Smith ', group: 'Group 2', avatar: srcs[5], icon: 'mdi-domain' },
  { id: 7, name: 'John Smith', group: 'Group 2', avatar: srcs[1], icon: 'mdi-message-text' },
  { id: 8, name: 'Sandra Williams', group: 'Group 2', avatar: srcs[3], icon: 'mdi-dialpad' }
])
</script>

<style scoped></style>
```

:::

## 示例（vx-selectmany）

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

## vx-linkageselect-remote

> legacy component

:::demo

```vue
<script setup lang="ts">
import { Ref, ref } from 'vue'

const value = ref([])

const labels = ref(['Province', 'City', 'District'])
</script>

<template>
  <p>{{ value }}</p>
  <vx-linkageselect-remote
    v-model="value"
    remote-url="http://localhost:7800/examples/api/linkage-select-server"
    :level-start="1"
    :labels="labels"
    select-out-of-order
  ></vx-linkageselect-remote>
</template>

<style scoped></style>
```

:::
