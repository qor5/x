<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  modelValue: any
  translations: any
}>()

props.modelValue.modifier = props.modelValue.modifier || 'contains'

const emit = defineEmits(['update:modelValue'])
const t = props.translations
const items = ref([
  { text: t.equals, value: 'equals' },
  { text: t.contains, value: 'contains' }
])
</script>

<template>
  <div>
    <div class="d-flex" v-if="!props.modelValue.disableChooseModifier">
      <v-select
        class="d-inline-block"
        style="min-width: 200px"
        v-model="props.modelValue.modifier"
        :items="items"
        item-title="text"
        item-value="value"
        variant="underlined"
        hide-details
      ></v-select>
    </div>
    <div class="d-flex">
      <v-icon
        class="py-8 mr-4 float-md-start"
        icon="mdi-subdirectory-arrow-right"
        size="large"
      ></v-icon>
      <v-text-field
        class="d-inline-block"
        style="min-width: 120px"
        variant="underlined"
        type="text"
        v-model="props.modelValue.valueIs"
        hide-etails
      />
    </div>
  </div>
</template>

<style scoped></style>
