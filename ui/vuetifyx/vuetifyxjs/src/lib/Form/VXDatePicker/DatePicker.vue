<template>
  <div class="vx-datepicker-wrap">
    <vx-field
      v-model="inputValue"
      :placeholder="placeholder"
      :label="label"
      ref="inputRef"
      @blur="onInputBlur"
    >
      <!-- @keydown.enter="onInputKeyEnter($refs.inputRef)" -->
      <!-- calendar icon -->
      <template #append-inner><v-icon icon="mdi-calendar-range-outline" /></template>

      <!-- drop down -->
      <v-overlay
        open-on-click
        :close-on-content-click="false"
        activator="parent"
        :scrim="false"
        :open-delay="0"
        :close-delay="0"
        max-width="300"
        location-strategy="connected"
      >
        <date-picker-base
          class="elevation-2 d-inline-block bg-background rounded-lg overflow-hidden"
          :model-value="datePickerValue"
          :use-time-select="type === 'datetimepicker'"
          @update:modelValue="onDatePickerValueChange"
        />
      </v-overlay>
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, computed, PropType, watch } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import datePickerBase from './DatePickerBase.vue'
// const { filteredAttrs } = useFilteredAttrs()
import dayjs from 'dayjs'
const formatMapDefault = {
  datepicker: 'YYYY-MM-DD',
  datetimepicker: 'YYYY-MM-DD HH:mm'
}
const props = defineProps({
  modelValue: [String, Number],
  label: String,
  placeholder: String,
  type: {
    type: String as PropType<'datepicker' | 'datetimepicker'>,
    default: 'datepicker'
  },
  format: {
    type: String,
    default: ''
  }
})
const showDropDown = ref(false)
const inputValue = ref()
const inputRef = ref()
const datePickerValue = ref()
const emit = defineEmits(['update:modelValue'])

const getFormatStr = () => formatMapDefault[props.type] || props.format

const convertValueForInputAndDatePicker = (
  value: string | number | undefined,
  shouldEmit?: boolean
) => {
  const formatStr = getFormatStr()

  //case: no init value
  if (!value) {
    inputValue.value = ''
    datePickerValue.value = ''
  } else {
    inputValue.value = formatStr ? dayjs(value).format(formatStr) : value
    datePickerValue.value = value ? dayjs(value).valueOf() : ''
  }

  // console.log('datePickerValue', dayjs(datePickerValue.value).format('YYYY-MM-DD HH:mm:ss'))

  if (shouldEmit) {
    emit('update:modelValue', emitModelValueFormat(datePickerValue.value))
  }
}

watch(
  () => props.modelValue,
  () => {
    convertValueForInputAndDatePicker(props.modelValue)
  }
)

function emitModelValueFormat(value: any) {
  const formatStr = getFormatStr()

  if (value === props.modelValue) return value

  return dayjs(dayjs(value).format(formatStr)).valueOf()
}

function onDatePickerValueChange(value: any) {
  emit('update:modelValue', emitModelValueFormat(value))
  // showDropDown.value = false
}

// function onInputKeyEnter(instance: any) {
//   showDropDown.value = false
//   instance.blur()
// }

function onInputBlur(obj: FocusEvent | string) {
  // fix blur event is more quick than modelValue change event
  // dropdown visible condition is conflict with this event
  // if (showDropDown.value) return

  let value

  if (obj instanceof FocusEvent) {
    const target = obj.target as HTMLInputElement
    value = target.value
  } else {
    value = obj
  }

  // the first time select date will trigger blur event
  if (!value) return

  // showDropDown.value = false

  convertValueForInputAndDatePicker(dayjs(value).valueOf(), true)
}
</script>

<style lang="scss" scoped>
.v-menu {
  &:deep(.v-overlay__content) {
    border-radius: 8px !important;
  }
}
</style>
