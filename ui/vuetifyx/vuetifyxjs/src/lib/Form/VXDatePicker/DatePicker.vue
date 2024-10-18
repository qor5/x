<template>
  <div class="vx-datepicker-wrap">
    <vx-field
      v-model="inputValue"
      :placeholder="placeholder"
      :label="label"
      ref="inputRef"
      @blur="onInputBlur"
      @keydown.enter="onInputBlur(inputValue, true)"
      v-bind="filteredAttrs"
      :style="minWidth"
    >
      <!-- calendar icon -->
      <template v-if="!hideAppendInner" #append-inner
        ><v-icon icon="mdi-calendar-range-outline"
      /></template>

      <!-- drop down -->
      <v-overlay
        v-model="showMenu"
        open-on-click
        :close-on-content-click="false"
        activator="parent"
        :scrim="false"
        :open-delay="0"
        :close-delay="0"
        min-width="292"
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
const { filteredAttrs } = useFilteredAttrs()
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
  },
  hideAppendInner: Boolean
})
const showMenu = ref(false)
const inputValue = ref()
const inputRef = ref()
const datePickerValue = ref()
const emit = defineEmits(['update:modelValue'])
const minWidth = computed(() => ({
  minWidth: props.type === 'datepicker' ? '140px' : '190px'
}))

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
  },
  { immediate: true }
)

function emitModelValueFormat(value: any) {
  const formatStr = getFormatStr()

  if (value === props.modelValue) return value

  return dayjs(dayjs(value).format(formatStr)).valueOf()
}

function onDatePickerValueChange(value: any) {
  emit('update:modelValue', emitModelValueFormat(value))
}

function onInputBlur(obj: FocusEvent | string, closeMenu: boolean = false) {
  // fix blur event is more quick than modelValue change event
  // dropdown visible condition is conflict with this event
  if (closeMenu) {
    showMenu.value = false
  }

  let value

  if (obj instanceof FocusEvent) {
    const target = obj.target as HTMLInputElement
    value = target.value
  } else {
    value = obj
  }

  // the first time select date will trigger blur event
  if (!value) return

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
