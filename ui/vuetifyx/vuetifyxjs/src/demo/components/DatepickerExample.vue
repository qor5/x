<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'

const value = ref('2023-10-01')
</script>

<template>
  <p>{{ value }}</p>
  <datepicker v-model="value" />
</template>

<style scoped></style>
