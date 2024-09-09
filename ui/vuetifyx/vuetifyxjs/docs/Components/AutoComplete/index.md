# v-autocomplete 自动填充

## 基本用法(vuetify)

:::demo

```vue
<template>
  <v-row>
    <v-col cols="12" sm="6">
      <v-autocomplete
        clearable
        label="Autocomplete"
        :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
        multiple
      ></v-autocomplete>
    </v-col>
    <v-col cols="12" sm="6">
      <v-autocomplete
        clearable
        chips
        density="compact"
        :items="['California', 'Colorado', 'Florida', 'Georgia', 'Texas', 'Wyoming']"
        multiple
        variant="outlined"
      ></v-autocomplete>
    </v-col>
  </v-row>
</template>

```
:::

## vx-autocomplete
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