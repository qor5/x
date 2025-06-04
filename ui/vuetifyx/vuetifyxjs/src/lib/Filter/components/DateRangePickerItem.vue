<script setup lang="ts">
import { ref } from 'vue'
import * as constants from '../Constants'
import Datepicker from '@/lib/Datepicker.vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

props.modelValue.modifier = props.modelValue.modifier || constants.ModifierBetween

const datePickerVisible = ref(false)

const emit = defineEmits(['update:modelValue'])
const modifier = props.modelValue.modifier
const options = props.modelValue.dateOptions
const fromOption = ref()
const toOption = ref()
if (options) {
  if (options.length >= 1) {
    fromOption.value = options[0]
  }
  if (options.length >= 2) {
    toOption.value = options[1]
  }
}
</script>

<template>
  <div style="width: 200px">
    <datepicker
      v-model="modelValue.valueFrom"
      :key="modifier + 'form'"
      :visible="datePickerVisible"
      :hide-details="true"
      :clear-text="translations['clear']"
      :ok-text="translations['ok']"
      :label="translations['startAt']"
      v-bind="fromOption"
    />
    <div style="height: 34px" class="pl-2 pt-4">
      <span>{{ translations['to'] }}</span>
    </div>
    <datepicker
      v-model="modelValue.valueTo"
      :key="modifier + 'to'"
      :hide-details="true"
      :clear-text="translations['clear']"
      :ok-text="translations['ok']"
      :label="translations['endAt']"
      v-bind="toOption"
    />
  </div>
</template>

<style scoped></style>
