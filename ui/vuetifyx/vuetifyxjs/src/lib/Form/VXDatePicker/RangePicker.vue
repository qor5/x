<template>
  <div class="vx-range-picker-wrap">
    <vx-label class="mb-2" :tooltip="tips" :label-for="name" :required-symbol="required">{{
      label
    }}</vx-label>

    <vx-field
      ref="inputFieldParent"
      class="vx-range-picker-field"
      :class="{ isFocus }"
      v-model="inputValue"
      v-bind="filteredAttrs"
      :focused="isFocus"
      :disabled="disabled"
      @mouseover="!disabled && (isHovering = true)"
      @mouseout="!disabled && (isHovering = false)"
    >
      <template #append-inner>
        <v-icon
          :icon="showClearIcon ? 'mdi-close-circle' : 'mdi-calendar-range-outline'"
          size="x-small"
          @click.stop="onClickAppendInner"
        />
      </template>

      <div class="vx-range-picker-group d-flex flex-1-1">
        <vx-field
          :class="{ current: current === 0 }"
          v-model="inputValue[0]"
          v-model:focused="isStartInputFocus"
          ref="startDateInput"
          :placeholder="placeholder[0]"
          variant="plain"
          class="flex-1-1"
          hide-details
          readonly
          @blur="onInputBlur($event, 0)"
          @keydown.enter="onInputBlur(inputValue[0], 0)"
          @click="onClickEditDate(0)"
        />
        <div class="separator" @click.stop="showMenu = true" />
        <vx-field
          :class="{ current: current === 1 }"
          v-model="inputValue[1]"
          v-model:focused="isEndInputFocus"
          ref="endDateInput"
          :placeholder="placeholder[1]"
          variant="plain"
          class="flex-1-1"
          hide-details
          readonly
          @blur="onInputBlur($event, 1)"
          @keydown.enter="onInputBlur(inputValue[1], 1)"
          @click="onClickEditDate(1)"
        />
      </div>

      <!-- drop down -->
      <v-overlay
        :model-value="showMenu"
        persistent
        target="parent"
        :scrim="false"
        :open-delay="0"
        :close-delay="0"
        no-click-animation
        min-width="292"
        location-strategy="connected"
        @click:outside="closeEditData()"
      >
        <div class="vx-date-picker-group elevation-5 bg-background rounded-lg">
          <date-picker-base
            class="d-inline-block overflow-hidden"
            :model-value="datePickerValue[0]"
            :format-str="formatStr"
            :type="type"
            @update:modelValue="onDatePickerValueChange($event, 0)"
            :date-picker-props="datePickerProps[0]"
            v-bind="filteredAttrs.datePickerProps?.[0]"
          />

          <date-picker-base
            class="d-inline-block overflow-hidden"
            :model-value="datePickerValue[1]"
            :format-str="formatStr"
            :type="type"
            @update:modelValue="onDatePickerValueChange($event, 1)"
            :date-picker-props="datePickerProps[1]"
            v-bind="filteredAttrs.datePickerProps?.[1]"
          />

          <div v-if="needConfirm" class="border-b mx-2" />
          <div v-if="needConfirm" class="btn-group text-right mt-2 pa-2">
            <v-btn color="primary" @click="onClickConfirm">{{ i18_save }}</v-btn>
          </div>
        </div>
      </v-overlay>
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, computed, PropType, watch, watchEffect } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import datePickerBase from './DatePickerBase.vue'
import { useDatePicker, datePickerType } from '@/lib/composables/useDatePicker'
const { filteredAttrs } = useFilteredAttrs()
import dayjs from 'dayjs'
import { useLocale } from 'vuetify'
import { useHasEventListener } from '@/lib/composables/useEventListener'
import { EnhancedDateParser } from '@/lib/utils/dateParser'
const { hasEventListener } = useHasEventListener()
const { t } = useLocale()
const i18_save = t('$vuetify.datePicker.saveBtn')
const isHovering = ref(false)
const inputFieldParent = ref()
const startDateInput = ref()
const endDateInput = ref()
const current = ref()
const isStartInputFocus = ref(false)
const isEndInputFocus = ref(false)

const emit = defineEmits(['update:modelValue', 'click:confirm', 'blur'])
const props = defineProps({
  modelValue: {
    type: Array as PropType<(string | number)[]>,
    default: ['', '']
  },
  tips: String,
  label: String,
  clearable: Boolean,
  tooltips: String,
  id: String,
  name: String,
  required: Boolean,
  disabled: Boolean,
  needConfirm: Boolean,
  placeholder: {
    type: [Array, String] as PropType<string[] | string>,
    default: ['', '']
  },
  datePickerProps: {
    type: Array as PropType<Record<string, any>>,
    default: []
  },
  type: {
    type: String as PropType<datePickerType>,
    default: 'datepicker'
  },
  format: {
    type: String,
    default: ''
  }
})
const inputValue = ref<(string | number)[]>(['', ''])
const datePickerValue = ref<(string | number)[]>(['', ''])
const { showMenu, formatStr, emitDatePickerValue, tempData } = useDatePicker(props, emit)

const isFocus = computed(() => (isStartInputFocus.value || isEndInputFocus.value) && showMenu)
const showClearIcon = computed(
  () =>
    (isHovering.value || isFocus.value || showMenu.value) &&
    props.clearable &&
    inputValue.value.some((item) => Boolean(item))
)

// this flag is used to format initial modelValue with formatStr
let onceEmitFlag = true

watch(
  () => props.modelValue,
  () => {
    // debugger
    convertValueForInputAndDatePicker({ value: props.modelValue, shouldEmit: onceEmitFlag })
    onceEmitFlag = false
  },
  { immediate: true }
)

// func: reset all temporal datepicker selected data when showMenu is false
watch(showMenu, (value) => {
  // console.log(value, 'showMenu')
  if (!value) {
    setTimeout(
      () =>
        convertValueForInputAndDatePicker({
          value: props.modelValue,
          extraEmitEvents: ['blur', 'update:modelValue']
        }),
      300
    )
  }
})

watch(
  () => tempData,
  (value) => {
    console.log('needConfirm')
    props.needConfirm &&
      convertValueForInputAndDatePicker({ value: tempData.value as (string | number)[] })
  }
)

function onDatePickerValueChange(value: number, position: 0 | 1) {
  let data = props.modelValue
  current.value = position

  // Handle case where only time is selected and no date
  if (!value && props.type === 'datetimepicker') {
    // If only time is provided but no date, use current date
    const currentDate = new Date()
    value = currentDate.valueOf()
  }

  if (datePickerValue.value.length === 0) {
    if (position === 0) data.push(value)
    else data = ['', value]
  } else if (datePickerValue.value.length === 1) {
    if (position === 0) data[0] = value
    else data.push(value)
  } else {
    datePickerValue.value[position] = value

    data = datePickerValue.value.map((item, i) => (position === i ? value : item))
  }

  emitDatePickerValue(data, { needConfirm: props.needConfirm })
}

function convertValueForInputAndDatePicker({
  value,
  shouldEmit,
  extraEmitEvents
}: {
  value: (string | number)[]
  shouldEmit?: boolean
  extraEmitEvents?: string[]
}) {
  // console.log('=== convertValueForInputAndDatePicker (RangePicker) ===', {
  //   value,
  //   shouldEmit,
  //   extraEmitEvents,
  //   currentDatePickerValue: datePickerValue.value,
  //   currentInputValue: inputValue.value
  // })

  //case: no init value
  if (!value || value.length === 0) {
    inputValue.value = ['', '']
    datePickerValue.value = ['', '']
  } else {
    if (Array.isArray(value)) {
      inputValue.value = value.map((item) => {
        if (!item) return ''

        try {
          // First try relative date parsing
          let parsedDate = EnhancedDateParser.parseRelativeDate(String(item))

          // If relative date parsing fails, use the enhanced date parser
          if (!parsedDate) {
            parsedDate = EnhancedDateParser.parseDate(item)
          }

          if (parsedDate && parsedDate.isValid()) {
            return parsedDate.format(formatStr.value)
          } else {
            console.warn('Failed to parse date with enhanced parser:', item)
            // If enhanced parsing fails, try original dayjs parsing as fallback
            const fallbackDate = dayjs(item)
            if (fallbackDate.isValid()) {
              return fallbackDate.format(formatStr.value)
            } else {
              return String(item)
            }
          }
        } catch (error) {
          console.error('Date conversion error:', error)
          return String(item)
        }
      })

      datePickerValue.value = value.map((item) => {
        if (!item) return item

        try {
          // First try relative date parsing
          let parsedDate = EnhancedDateParser.parseRelativeDate(String(item))

          // If relative date parsing fails, use the enhanced date parser
          if (!parsedDate) {
            parsedDate = EnhancedDateParser.parseDate(item)
          }

          if (parsedDate && parsedDate.isValid()) {
            return parsedDate.valueOf()
          } else {
            // If enhanced parsing fails, try original dayjs parsing as fallback
            const fallbackDate = dayjs(item)
            if (fallbackDate.isValid()) {
              return fallbackDate.valueOf()
            } else {
              return ''
            }
          }
        } catch (error) {
          console.error('Date conversion error:', error)
          return ''
        }
      })
    } else {
      try {
        // First try relative date parsing
        let parsedDate = EnhancedDateParser.parseRelativeDate(String(value))

        // If relative date parsing fails, use the enhanced date parser
        if (!parsedDate) {
          parsedDate = EnhancedDateParser.parseDate(value)
        }

        if (parsedDate && parsedDate.isValid()) {
          inputValue.value = [parsedDate.format(formatStr.value)]
          datePickerValue.value = [parsedDate.valueOf()]
        } else {
          console.warn('Failed to parse single date with enhanced parser:', value)
          // If enhanced parsing fails, try original dayjs parsing as fallback
          const fallbackDate = dayjs(value)
          if (fallbackDate.isValid()) {
            inputValue.value = [fallbackDate.format(formatStr.value)]
            datePickerValue.value = [fallbackDate.valueOf()]
          } else {
            inputValue.value = [String(value)]
            datePickerValue.value = ['']
          }
        }
      } catch (error) {
        console.error('Date conversion error:', error)
        inputValue.value = [String(value)]
        datePickerValue.value = ['']
      }
    }
  }

  shouldEmit &&
    emitDatePickerValue(datePickerValue.value, { needConfirm: props.needConfirm, extraEmitEvents })
}

function onInputBlur(obj: FocusEvent | string | number, position: 0 | 1) {
  inputFieldParent.value.blur()

  let value

  if (obj instanceof FocusEvent) {
    const target = obj.target as HTMLInputElement
    value = target.value
  } else {
    value = obj
  }

  // the first time select date will trigger blur event
  if (!value) return

  // If the user entered something, try to use the enhanced date parser
  if (value) {
    try {
      // First try relative date parsing
      let parsedDate = EnhancedDateParser.parseRelativeDate(String(value))

      // If relative date parsing fails, use the enhanced date parser
      if (!parsedDate) {
        parsedDate = EnhancedDateParser.parseDate(value)
      }

      if (parsedDate && parsedDate.isValid()) {
        value = parsedDate.valueOf()
      } else {
        console.warn('Failed to parse date with enhanced parser:', value)
        // If enhanced parsing fails, try original dayjs parsing as fallback
        const fallbackDate = dayjs(value)
        if (fallbackDate.isValid()) {
          value = fallbackDate.valueOf()
        } else {
          console.error('Failed to parse date:', value)
          return
        }
      }
    } catch (error) {
      console.error('Date conversion error:', error)
      return
    }
  }

  if (props.datePickerProps[position]) {
    const currentConfig = props.datePickerProps[position]
    const numericValue = Number(value)

    if (currentConfig.max) {
      const maxTimestamp = currentConfig.max ? dayjs(currentConfig.max).valueOf() : 0
      value = Math.min(numericValue, maxTimestamp)
    }

    if (currentConfig.min) {
      const minTimestamp = currentConfig.min ? dayjs(currentConfig.min).valueOf() : 0
      value = Math.max(minTimestamp, numericValue)
    }
  }

  convertValueForInputAndDatePicker({
    value: inputValue.value.map((item, i) => (i === position ? value : item)),
    shouldEmit: true,
    extraEmitEvents: ['blur']
  })
}

function onClickEditDate(index: number) {
  showMenu.value = true
}

function reset() {
  current.value = null
  tempData.value = ['', '']
}

function onClickAppendInner() {
  if (showClearIcon.value) {
    emitDatePickerValue(['', ''])
    reset()
    showMenu.value = false
  } else {
    showMenu.value = true
  }
}

function closeEditData() {
  if (isFocus.value) return
  showMenu.value = false
  current.value = null
}

function onClickConfirm() {
  if (hasEventListener('click:confirm')) {
    new Promise((resolve) => {
      emit('click:confirm', { next: resolve, value: tempData.value || [] })
    }).then(() => {
      emitDatePickerValue(tempData.value || [])
      reset()
      showMenu.value = false
    })
  } else {
    console.log(tempData.value, 'tempData.value')
    emitDatePickerValue(tempData.value || [])
    reset()
    showMenu.value = false
  }
}
</script>

<style lang="scss" scoped>
.vx-range-picker-field {
  .current > :deep(.v-input):not(.v-input--error) {
    .v-field {
      &::after {
        height: 3px !important;
        transition: all ease 0.3s;
        background: #3e63dd;
        width: calc(100% - 24px);
      }
    }
  }

  & > :deep(.v-input) {
    .v-field--variant-plain {
      padding: 0 12px;
      .v-field__input {
        padding-bottom: 8px;
      }
    }

    & .vx-range-picker-group + input {
      display: none;
    }

    & > .v-input__control > .v-field {
      padding-inline-start: 0;

      & > .v-field__field > .v-field__input {
        padding: 0;
      }

      .v-field {
        position: relative;
        &::after {
          transition: all ease 0.3s;
          position: absolute;
          height: 0;
          content: '';
          bottom: -2px;
          left: 12px;
        }
      }
    }
  }
}

.vx-range-picker-wrap {
  :deep(.v-field) {
    &,
    * {
      cursor: pointer;
    }
  }
}

.vx-date-picker-group {
  &:deep(.v-picker-wrap) {
    padding-bottom: 0;
    .v-date-picker-month {
      padding-bottom: 4px;
    }
  }
}

.separator {
  display: flex;
  justify-content: center;
  align-items: center;
  &::before {
    display: block;
    content: '';
    height: 1px;
    background: rgb(var(--v-theme-grey));
    width: 16px;
  }
}
</style>
