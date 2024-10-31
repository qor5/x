<script setup lang="ts">
const props = defineProps<{
  modelValue: any
}>()

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div style="width: 200px">
    <v-autocomplete
      v-model="modelValue.valueIs"
      :items="modelValue.options"
      style="width: 200px"
      class="d-inline-block"
      item-title="text"
      variant="underlined"
      item-value="value"
      hide-details
    />
  </div>
</template>

<style scoped></style>
