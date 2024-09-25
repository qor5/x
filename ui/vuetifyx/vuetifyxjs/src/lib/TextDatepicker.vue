<script setup lang="ts">
import Datepicker from '@/lib/Datepicker.vue'
import { ref, watch } from 'vue'
import { useVDatePickerTimeChange } from '@/lib/composables/useVDatePicker'

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: { type: String },
  visible: { type: Boolean, default: false }
})
const value = ref(props.modelValue)

const {
  displayedMonth,
  displayedYear,
  setDisplayedYearAndMonth,
  onYearOrMonthChange,
  valueChangedWithoutSaved
} = useVDatePickerTimeChange(value)
setDisplayedYearAndMonth(value.value)

const internalVisible = ref(props.visible)
const toggle = () => {
  internalVisible.value = !internalVisible.value
}

watch(
  () => value,
  (newVal) => {
    emit('update:modelValue', value.value)
    toggle()
  }
)
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

    <datepicker
      v-model="value"
      :year="displayedYear"
      :month="displayedMonth"
      @update:year="onYearOrMonthChange($event, 'year')"
      @update:month="onYearOrMonthChange($event, 'month')"
      @update:modelValue="onYearOrMonthChange($event, 'modelValue')"
    />
  </v-menu>
</template>

<style scoped></style>
