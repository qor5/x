<script setup lang="ts">
import VxLinkageSelectRemote from '@/lib/LinkageSelectRemote/index.vue'
import { ref } from 'vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

const emit = defineEmits(['update:modelValue'])
const ads = ref(props.modelValue.linkageSelectData.linkageSelectRemoteOptions)
</script>

<template>
  <vx-linkage-select-remote
    style="width: 500px"
    v-model="modelValue.valuesAre"
    :labels="modelValue.linkageSelectData.labels"
    :select-out-of-order="modelValue.linkageSelectData.selectOutOfOrder"
    row
    hide-details
    v-bind="ads"
  />
</template>

<style scoped></style>
