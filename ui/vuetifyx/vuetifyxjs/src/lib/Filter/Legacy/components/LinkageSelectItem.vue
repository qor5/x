<script setup lang="ts">
import VxLinkageSelect from '@/lib/LinkageSelect.vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <vx-linkage-select
    style="width: 500px"
    v-model="modelValue.valuesAre"
    :items="modelValue.linkageSelectData.items"
    :labels="modelValue.linkageSelectData.labels"
    :select-out-of-order="modelValue.linkageSelectData.selectOutOfOrder"
    row
    hide-details
  />
</template>

<style scoped></style>
