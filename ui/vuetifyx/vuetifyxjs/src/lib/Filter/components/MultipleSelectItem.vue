<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

props.modelValue.modifier = props.modelValue.modifier || 'in'

const t = props.translations

const emit = defineEmits(['update:modelValue'])
const items = ref([
  { text: t.in, value: 'in' },
  { text: t.notIn, value: 'notIn' }
])
</script>

<template>
  <div>
    <div v-if="!props.modelValue.disableChooseModifier">
      <v-select
        class="d-inline-block"
        style="width: 200px"
        v-model="props.modelValue.modifier"
        :items="items"
        item-title="text"
        item-value="value"
        variant="underlined"
        hide-details
      ></v-select>
    </div>
    <div style="max-height: 160px; overflow-y: scroll">
      <v-checkbox
        v-for="opt in modelValue.options"
        v-model="modelValue.valuesAre"
        :label="opt.text"
        :value="opt.value"
        hide-details
        density="comfortable"
      ></v-checkbox>
    </div>
  </div>
</template>

<style scoped></style>
