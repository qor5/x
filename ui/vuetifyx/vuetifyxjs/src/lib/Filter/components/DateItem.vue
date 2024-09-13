<script setup lang="ts">
import { ref } from 'vue'
import Datepicker from '@/lib/Datepicker.vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

const datePickerVisible = ref(false)

const emit = defineEmits(['update:modelValue'])
const options = props.modelValue.dateOptions
const option = ref()
if (options) {
  if (options.length >= 1) {
    option.value = options[0]
  }
}
</script>

<template>
  <div style="width: 200px">
    <datepicker
      v-model="modelValue.valueIs"
      :visible="datePickerVisible"
      :hide-details="true"
      :clear-text="translations['clear']"
      :ok-text="translations['ok']"
      v-bind="option"
    />
  </div>
</template>

<style scoped></style>
