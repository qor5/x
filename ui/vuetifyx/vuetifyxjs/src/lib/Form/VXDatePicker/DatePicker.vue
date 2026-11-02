<template>
  <div class="vx-datepicker-wrap">
    <vx-label
      v-if="label"
      class="mb-2"
      :tooltip="tips"
      :label-for="name"
      :required-symbol="required"
      >{{ label }}</vx-label
    >

    <vx-field
      v-model="inputValue"
      :placeholder="placeholder"
      :focused="isFocus"
      ref="inputRef"
      @blur="onInputBlur"
      @mouseover="isHovering = true"
      @mouseout="isHovering = false"
      @keydown.enter="onInputBlur(inputValue, true)"
      v-bind="filteredAttrs"
      :style="minWidth"
    >
      <!-- calendar icon -->
      <template v-if="!hideAppendInner" #append-inner
        ><v-icon
          :icon="showClearIcon ? 'mdi-close-circle' : 'mdi-calendar-range-outline'"
          @click.stop="onClickAppendInner"
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
        scroll-strategy="close"
      >
        <date-picker-base
          class="elevation-5 d-inline-block bg-background rounded-lg overflow-hidden"
          :model-value="datePickerValue"
          :type="type"
          @update:modelValue="onDatePickerUpdate"
          :format-str="formatStr"
          :datePickerProps="datePickerProps"
        />
      </v-overlay>

      <input readonly class="input-cover" :value="inputValue" :placeholder="placeholder" />
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, PropType, watchEffect, watch, ref } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import datePickerBase from './DatePickerBase.vue'
import { useDatePicker, datePickerType } from '@/lib/composables/useDatePicker'
import dayjs from 'dayjs'
import { EnhancedDateParser } from '@/lib/utils/dateParser'

const { filteredAttrs } = useFilteredAttrs()

const props = defineProps({
  modelValue: [String, Number, Date],
  tips: String,
  name: String,
  required: Boolean,
  label: String,
  placeholder: String,
  type: {
    type: String as PropType<datePickerType>,
    default: 'datepicker'
  },
  clearable: Boolean,
  format: {
    type: String,
    default: ''
  },
  hideAppendInner: Boolean,
  datePickerProps: Object
})
const inputValue = ref()
const inputRef = ref()
const datePickerValue = ref()
const isHovering = ref(false)
const isFocus = computed(() => showMenu.value)
const emit = defineEmits(['update:modelValue', 'blur'])
const { showMenu, formatStr, emitDatePickerValue } = useDatePicker(props, emit)

const showClearIcon = computed(
  () => (isHovering.value || showMenu.value) && inputValue.value && props.clearable
)
const minWidth = computed(() => ({
  minWidth: props.type === 'datepicker' ? '140px' : '190px'
}))

watch(
  () => showMenu.value,
  (newVal, oldVal) => {
    // When the dropdown closes, emit blur event (if there is a value)
    if (oldVal && !newVal && datePickerValue.value) {
      const formattedValue = dayjs(datePickerValue.value).format(formatStr.value)
      emit('blur', formattedValue)
    }
  }
)

// Add watch for time selection with no date
watch(
  () => props.modelValue,
  (value) => {
    if (value && props.type === 'datetimepicker') {
      // When value is changed and it has a time component but no date
      const val = dayjs(value)
      if (val.hour() !== 0 || val.minute() !== 0 || val.second() !== 0) {
        // Has time component
        if (val.year() === 1970 && val.month() === 0 && val.date() === 1) {
          // Likely only time was set (Unix epoch date)
          const currentDate = new Date()
          const newDate = dayjs(currentDate)
            .hour(val.hour())
            .minute(val.minute())
            .second(val.second())

          emitDatePickerValue(newDate.valueOf())
        }
      }
    }
  }
)

watchEffect(() => {
  convertValueForInputAndDatePicker({ value: props.modelValue })
})

function onInputBlur(obj: FocusEvent | string, closeMenu: boolean = false) {
  if (closeMenu) {
    showMenu.value = false
    inputRef.value.blur()
  }

  if (obj instanceof FocusEvent) {
    const target = obj.target as HTMLInputElement
    const inputText = target.value

    // If the user did not enter anything, and there is already a selected value, keep the current value unchanged
    if (!inputText && datePickerValue.value) {
      return
    }

    // If the user entered something, try to parse it
    if (inputText) {
      convertValueForInputAndDatePicker({
        value: inputText,
        shouldEmit: true,
        extraEmitEvents: ['blur']
      })
    }
  }
}

function convertValueForInputAndDatePicker({
  value,
  shouldEmit,
  extraEmitEvents
}: {
  value: string | number | undefined | Date
  shouldEmit?: boolean
  extraEmitEvents?: string[]
}) {
  // console.log('=== convertValueForInputAndDatePicker ===', {
  //   value,
  //   shouldEmit,
  //   extraEmitEvents,
  //   currentDatePickerValue: datePickerValue.value,
  //   currentInputValue: inputValue.value
  // })

  if (!value) {
    inputValue.value = ''
    datePickerValue.value = null
  } else {
    try {
      // First try relative date parsing
      let parsedDate = EnhancedDateParser.parseRelativeDate(String(value))

      // If relative date parsing fails, use the enhanced date parser
      if (!parsedDate) {
        parsedDate = EnhancedDateParser.parseDate(value)
      }

      if (parsedDate && parsedDate.isValid()) {
        datePickerValue.value = parsedDate.valueOf()

        // Display value is formatted according to format
        const currentFormatStr = formatStr.value
        if (currentFormatStr) {
          inputValue.value = parsedDate.format(currentFormatStr)
        } else {
          inputValue.value = parsedDate.format('YYYY-MM-DD') // Default format
        }
      } else {
        console.warn('Failed to parse date with enhanced parser:', value)
        // If enhanced parsing fails, try original dayjs parsing as fallback
        const fallbackDate = dayjs(value)
        if (fallbackDate.isValid()) {
          datePickerValue.value = fallbackDate.valueOf()
          const currentFormatStr = formatStr.value
          if (currentFormatStr) {
            inputValue.value = fallbackDate.format(currentFormatStr)
          } else {
            inputValue.value = fallbackDate.format('YYYY-MM-DD')
          }
        } else {
          // If all parsing fails, keep the original input but clear internal value
          inputValue.value = String(value)
          datePickerValue.value = null
        }
      }
    } catch (error) {
      console.error('Date conversion error:', error)
      inputValue.value = String(value)
      datePickerValue.value = null
    }
  }

  // When emitting event, emit formatted value if there is a value, otherwise emit an empty string
  if (shouldEmit) {
    const emitValue = datePickerValue.value
      ? dayjs(datePickerValue.value).format(formatStr.value)
      : ''
    emitDatePickerValue(emitValue, { extraEmitEvents })
  }
}

function onClickAppendInner() {
  if (showClearIcon.value) {
    // Explicitly mark this as a clear operation
    inputValue.value = ''
    datePickerValue.value = null
    emitDatePickerValue('', { extraEmitEvents: ['clear'] })
    showMenu.value = false
  } else {
    showMenu.value = true
  }
}

function onDatePickerUpdate(value: number) {
  // Internally store the real value (timestamp)
  datePickerValue.value = value

  // Display value is formatted according to format
  if (value) {
    inputValue.value = dayjs(value).format(formatStr.value)
  } else {
    inputValue.value = ''
  }

  // Emit formatted value to parent component
  const formattedValue = value ? dayjs(value).format(formatStr.value) : ''
  emitDatePickerValue(formattedValue)
}
</script>

<style lang="scss" scoped>
.v-menu {
  &:deep(.v-overlay__content) {
    border-radius: 8px !important;
  }
}

.vx-datepicker-wrap {
  &:deep(.v-input) {
    .input-cover {
      position: absolute;
      width: 100%;
      height: 100%;
      z-index: 1;
      pointer-events: none;
    }

    & input:not(.input-cover) {
      display: none;
    }

    .v-field {
      cursor: pointer;
    }
  }
}
</style>

<style lang="scss">
.v-picker-wrap .v-date-picker-month__days .v-date-picker-month__day--selected .v-btn[disabled] {
  color: #fff;
}
</style>
