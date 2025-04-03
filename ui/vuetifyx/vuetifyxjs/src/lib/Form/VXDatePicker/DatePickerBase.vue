<template>
  <div class="v-picker-ui-wrap" :class="{ 'pb-5': useTimeSelect }">
    <v-date-picker
      class="v-picker-wrap"
      hide-header
      flat
      show-adjacent-months
      v-model="dateOfPicker"
      @update:year="onYearOrMonthChange($event, 'year')"
      @update:month="onYearOrMonthChange($event, 'month')"
      @update:modelValue="onYearOrMonthChange($event, 'date')"
      v-bind="combinedProps"
    />

    <time-select
      v-if="useTimeSelect"
      class="time-select-wrap"
      :formatStr="formatStr"
      v-model="timeStr"
      v-bind="propsForTimeSelect"
      @update:modelValue="onTimeSelected"
    />
  </div>
</template>

<script setup lang="ts">
import dayjs from 'dayjs'
import { ref, defineEmits, defineProps, PropType, computed, watch, Prop } from 'vue'
import TimeSelect from './TimeSelect.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import { format } from 'path'
const { filteredAttrs } = useFilteredAttrs()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: [String, Number] as PropType<string | number>,
  type: {
    type: String,
    default: 'datepicker'
  },
  datePickerProps: Object as PropType<any>,
  formatStr: String,
  disableSecond: Boolean,
  disableMinute: Boolean,
  disableHour: Boolean
})
const timeStr = ref()
const dateOfPicker = ref()

const dateStr = computed(() => dayjs(dateOfPicker.value).format('YYYY-MM-DD'))
const useTimeSelect = computed(() => props.type === 'datetimepicker')

watch(
  () => props.modelValue,
  (value) => {
    timeStr.value = value ? dayjs(value).format('HH:mm:ss') : '00:00:00'

    if (value) {
      const date = new Date(value)
      if (!isNaN(date.getTime())) {
        dateOfPicker.value = date
      }
    }
  },
  { immediate: true }
)

function emitValue(date: string, time: string) {
  // console.log('emitValue', `${date} ${time}`)
  return dayjs(`${date} ${time}`).valueOf()
}

function onYearOrMonthChange(value: number | unknown | Date, type: 'year' | 'month' | 'date') {
  let newDate = ''
  let newTimeStr = timeStr.value
  let emitValueImmediate = false

  if (type === 'year') {
    newDate = dateStr.value
      .split('-')
      .map((item: string, index: number) => (index === 0 ? value : item))
      .join('-')

    emitValueImmediate = !!props.modelValue
  } else if (type === 'month') {
    newDate = dateStr.value
      .split('-')
      .map((item: string, index: number) => (index === 1 ? (value as number) + 1 : item))
      .join('-')
    emitValueImmediate = !!props.modelValue
  } else {
    newDate = dayjs(value as number).format('YYYY-MM-DD')
    emitValueImmediate = true
  }

  // only emitValue when date is selected
  emitValueImmediate && emit('update:modelValue', emitValue(newDate, newTimeStr))
}

function onTimeSelected(time: string) {
  // If no date is selected, use current date
  if (!dateOfPicker.value) {
    // Set current date
    dateOfPicker.value = new Date()
  }

  // now emitValue when time is selected (with current date if none selected)
  emit('update:modelValue', emitValue(dateStr.value, timeStr.value))
}

const combinedProps = computed(() => ({
  ...props.datePickerProps,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))

const propsForTimeSelect = computed(() => {
  const { disableSecond, disableMinute, disableHour, ...rest } = props.datePickerProps || {}
  return { disableSecond, disableMinute, disableHour }
})
</script>

<style lang="scss" scoped>
.v-picker-wrap {
  padding: 8px 0;
  width: 292px;

  &:deep(.v-date-picker-years),
  &:deep(.v-date-picker-months) {
    height: 256px;
  }

  &:deep(.v-date-picker-month__days) {
    flex: initial;
    column-gap: 2px;

    .v-btn.v-btn--variant-outlined {
      color: rgb(var(--v-theme-primary));
    }

    .v-date-picker-month__day {
      width: 36px;
      height: 36px;
    }

    .v-btn.v-date-picker-month__day-btn {
      --v-btn-size: 14px;
      --v-btn-height: 20px;
    }

    .v-date-picker-month__weekday {
      color: rgb(var(--v-theme-grey-darken-1));
    }

    .v-date-picker-month__day--selected {
      .v-btn {
        background-color: rgb(var(--v-theme-primary));
      }
    }
  }

  // &:deep(.v-date-picker-month__day--week-end) ~ .v-date-picker-month__day--hide-adjacent {
  //   display: none;
  // }

  &:deep(.v-date-picker-controls) {
    display: block;
    position: relative;
    padding-inline-start: 15px;
    padding-inline-end: 15px;
    padding-bottom: 20px;

    .v-date-picker-controls__month-btn {
      position: absolute;
      left: 50%;
      margin-left: -10px;
      transform: translateX(-50%);
      top: 2px;
      .v-btn__content {
        color: rgb(var(--v-theme-on-surface));
      }
    }

    .v-date-picker-controls__mode-btn {
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
      margin-left: 65.5px;
      top: 2px;
    }

    .v-date-picker-controls__month {
      display: flex;
      width: 100%;

      .v-btn {
        --v-btn-height: 20px;
      }

      & > :nth-child(2) {
        margin-left: auto;
      }
    }
  }
}

.time-select-wrap {
  justify-content: center;
  width: 292px;
}
</style>
