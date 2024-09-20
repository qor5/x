<template>
  <div class="vx-dialog-wrap">
    <v-dialog :model-value="dialogVisible" v-bind="filteredAttrs"></v-dialog>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs } = useFilteredAttrs()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: Boolean
})

const dialogVisible = ref(props.modelValue)

watch(
  () => props.modelValue,
  (newValue) => {
    dialogVisible.value = newValue
  }
)

function onUpdateModelValue(value: any) {
  emit('update:modelValue', value)
  dialogVisible.value = value
}
</script>

<style lang="sass" scoped></style>
