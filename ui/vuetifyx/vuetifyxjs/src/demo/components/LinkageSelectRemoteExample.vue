<script setup lang="ts">
import {ref} from 'vue'
import LinkageSelectRemote from '@/lib/LinkageSelectRemote/index.vue'

const value = ref([])
const labels = ref(['Province', 'City', 'District'])
const value2 = ref([{ID: "1", Name: "浙江"}])
</script>

<template>
  <v-divider class="mb-6"></v-divider>

  <span class="text-h4 mb-4">linkage-select-remote-example-select-out-of-order</span>
  <p>{{ value }}</p>
  <linkage-select-remote
    v-model="value"
    :level-start="1"
    :page-size="2"
    remote-url="http://localhost:7800/examples/api/linkage-select-server"
    :labels="labels"
    :select-out-of-order="true"

  ></linkage-select-remote>

  <v-divider class="mb-6"></v-divider>
  <span class="text-h4 mb-4">linkage-select-remote-example-row</span>
  <p>{{ value2 }}</p>
  <linkage-select-remote
    v-model="value2"
    :is-paging="true"
    :level-start="1"
    row
    :page-size="2"
    remote-url="http://localhost:7800/examples/api/linkage-select-server"
    :labels="labels"
  ></linkage-select-remote>
</template>

<style scoped></style>
