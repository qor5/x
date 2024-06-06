<script setup lang="ts">
import Datepicker from '@/lib/Datepicker.vue'
import { ref } from 'vue'

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: { type: String },
  visible: { type: Boolean, default: false }
})
const value = ref(props.modelValue)

const internalVisible = ref(props.visible)
const toggle = () => {
  internalVisible.value = !internalVisible.value
}

const change = () => {
  emit('update:modelValue', value.value)
  toggle()
}
</script>

<template>
  <v-menu
    class="d-inline-block"
    min-width="290px"
    eager
    v-model="internalVisible"
    location="end bottom"
    @input="toggle"
  >
    <template v-slot:activator="{ props }">
      <v-text-field
        class="d-inline-block"
        v-bind="props"
        style="width: 180px"
        hide-details
        variant="underlined"
        v-model="value"
        prepend-inner-icon="mdi-event"
      ></v-text-field>
    </template>

    <datepicker v-model="value" @update:modelValue="change"></datepicker>
  </v-menu>
</template>

<style scoped></style>
