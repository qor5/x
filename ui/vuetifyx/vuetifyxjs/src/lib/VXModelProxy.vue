<template>
  <div>
    <slot :model-value="internalModelValue" @update:model-value="handleUpdateModelValue"></slot>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Use inherited-attrs: false to prevent attrs from being applied to the root element
defineOptions({
  inheritAttrs: false
})

const props = defineProps<{
  modelValue?: string
  formatModel?: 'jsonStringify'
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

// Convert from JSON string to object when component receives a string value
const internalModelValue = computed(() => {
  if (props.formatModel === 'jsonStringify' && props.modelValue) {
    try {
      return JSON.parse(props.modelValue)
    } catch (error) {
      console.error('Failed to parse JSON string:', error)
      return {}
    }
  }
  return props.modelValue
})

// Update and emit when the inner component updates its model value
const handleUpdateModelValue = (value: any) => {
  if (props.formatModel === 'jsonStringify') {
    try {
      const jsonString = JSON.stringify(value)
      emit('update:modelValue', jsonString)
    } catch (error) {
      console.error('Failed to stringify JSON object:', error)
    }
  } else {
    emit('update:modelValue', value)
  }
}
</script>

<style scoped></style>
