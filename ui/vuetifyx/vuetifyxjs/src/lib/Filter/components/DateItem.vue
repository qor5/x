<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

const datePickerVisible = ref(false)

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div style="width: 200px">
    <datepicker v-model="modelValue.valueIs" :visible="datePickerVisible" :hide-details="true" />
  </div>
</template>

<style scoped></style>
