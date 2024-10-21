import dayjs from 'dayjs'
import { isEqual } from 'lodash'
import { ref } from 'vue'
export type datePickerType = 'datepicker' | 'datetimepicker'

// the duplicate logic on vx-date-picker and vx-range-picker
export function useDatePicker<EmitFnType>(props: any, emit: EmitFnType) {
  const showMenu = ref(false)
  const tempData = ref<string | number | (string | number)[]>()
  const formatMapDefault = {
    datepicker: 'YYYY-MM-DD',
    datetimepicker: 'YYYY-MM-DD HH:mm'
  }

  const formatStr = props.format || formatMapDefault[props.type as datePickerType]

  function emitModelValueFormat(value: number | string | (number | string)[]) {
    if (isEqual(value, props.modelValue)) return value

    return Array.isArray(value)
      ? value.map((item) => dayjs(item).valueOf())
      : dayjs(value).valueOf()
  }

  function emitDatePickerValue(
    value: number | string | (number | string)[],
    needConfirm?: boolean
  ) {
    if (needConfirm) {
      if (Array.isArray(value)) {
        if (!tempData.value || !Array.isArray(tempData.value)) {
          tempData.value = value
        } else {
          tempData.value = value.map(
            (item, i) => item || (tempData.value as (number | string)[])[i]
          )
        }
      } else {
        ;(tempData.value as string | number) = value
      }
    } else {
      ;(emit as any)('update:modelValue', emitModelValueFormat(value))
    }
  }

  return {
    showMenu,
    formatStr,
    emitDatePickerValue,
    tempData
  }
}
