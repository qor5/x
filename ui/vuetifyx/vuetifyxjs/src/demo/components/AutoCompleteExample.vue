<script setup lang="ts">
import VxAutocomplete from '@/lib/Autocomplete.vue'
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
const items = ref([])
const value = ref()
</script>
<template>
  <vue-json-pretty :data="value"></vue-json-pretty>
  <!--  <vx-autocomplete-->
  <!--    v-model="value"-->
  <!--    :items="items"-->
  <!--  ></vx-autocomplete>-->
  <vx-autocomplete
    :load-data="loadData"
    sorting
    :items="items"
    has-icon
    :remote="remote"
    v-model="value"
  ></vx-autocomplete>
</template>

<style scoped></style>
