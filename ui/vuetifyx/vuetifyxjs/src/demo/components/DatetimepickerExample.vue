<script setup lang="ts">
import { ref } from 'vue'
import Datetimepicker from '@/lib/Datetimepicker.vue'

const value = ref('2023-10-01 22:33')
</script>

<template>
  <p>{{ value }}</p>
  <datetimepicker v-model="value" :date-picker-props="{ min: '2016-06-15', max: '2028-03-20' }" />
</template>

<style scoped></style>
