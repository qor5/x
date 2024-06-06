<script setup lang="ts">
import { ref } from 'vue'
import selectMany from '@/lib/SelectMany.vue'

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
  <select-many v-model="value" :items="items" />
</template>

<style scoped></style>
