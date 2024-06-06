<script setup lang="ts">
import { ref } from 'vue'
import TextDatepicker from '@/lib/TextDatepicker.vue'

const value = ref('2023-10-01')
</script>

<template>
  <p>{{ value }}</p>
  <text-datepicker v-model="value" />
</template>

<style scoped></style>
