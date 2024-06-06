<script setup lang="ts">
import Autocomplete from '@/lib/Autocomplete.vue'

const props = defineProps<{
  modelValue: any
}>()

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div style="width: 200px">
    <autocomplete
      v-model="modelValue.valueIs"
      :items="modelValue.options"
      style="width: 200px"
      class="d-inline-block"
      hide-details
    />
  </div>
</template>

<style scoped></style>
