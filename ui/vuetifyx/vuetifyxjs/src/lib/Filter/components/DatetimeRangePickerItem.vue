<script setup lang="ts">
import { ref } from 'vue'
import * as constants from '../Constants'
import RangePicker from '@/lib/Form/VXDatePicker/RangePicker.vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

props.modelValue.modifier = props.modelValue.modifier || constants.ModifierBetween

const datePickerVisible = ref(false)

const emit = defineEmits(['update:modelValue'])
const options = props.modelValue.dateOptions
const option = ref()
if (options) {
  if (options.length >= 1) {
    option.value = options[0]
  }
}
const value = ref([props.modelValue.valueFrom, props.modelValue.valueTo])

const updateModelValue = (val: any) => {
  value.value = val
  props.modelValue.valueFrom = val[0]
  props.modelValue.valueTo = val[1]
  emit('update:modelValue', props.modelValue)
}
</script>

<template>
  <div style="width: 525px">
    <range-picker
      :visible="datePickerVisible"
      clearable
      @update:model-value="updateModelValue"
      v-model="value"
      type="datetimepicker"
      v-bind="option"
    ></range-picker>
  </div>
</template>

<style scoped></style>
