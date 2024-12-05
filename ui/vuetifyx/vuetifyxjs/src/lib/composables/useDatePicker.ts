import dayjs from 'dayjs'
import _ from 'lodash'
import { ref } from 'vue'
export type datePickerType = 'datepicker' | 'datetimepicker'

// the duplicate logic on vx-date-picker and vx-range-picker
export function useDatePicker<EmitFnType>(props: any, emit: EmitFnType) {
  const showMenu = ref(false)
  const tempData = ref<string | number | (string | number)[]>()
  const formatMapDefault = {
    datepicker: 'YYYY-MM-DD',
    datetimepicker: 'YYYY-MM-DD HH:mm:ss'
  }

  const formatStr = props.format || formatMapDefault[props.type as datePickerType]

  function emitModelValueFormat(value: number | string | (number | string)[], formatStr?: string) {
    if (_.isEqual(value, dayjs(props.modelValue).format(formatStr))) return value

    const formatedValue = Array.isArray(value)
      ? value.map((item) =>
          item
            ? dayjs(item).format(formatStr || formatMapDefault[props.type as datePickerType])
            : ''
        )
      : value
        ? dayjs(value).format(formatStr || formatMapDefault[props.type as datePickerType])
        : ''

    return formatedValue
  }

  function emitDatePickerValue(
    value: number | string | (number | string)[],
    {
      needConfirm,
      formatStr = formatMapDefault[props.type as datePickerType]
    }: {
      needConfirm?: boolean
      formatStr?: string
    } = {
      needConfirm: false,
      formatStr: formatMapDefault[props.type as datePickerType]
    }
  ) {
    if (needConfirm) {
      if (Array.isArray(value)) {
        if (!tempData.value || !Array.isArray(tempData.value)) {
          tempData.value = emitModelValueFormat(value, formatStr)
        } else {
          tempData.value = emitModelValueFormat(
            value.map((item, i) => item || (tempData.value as (number | string)[])[i]),
            formatStr
          )
        }
      } else {
        tempData.value = emitModelValueFormat(value, formatStr)
      }
    } else {
      ;(emit as any)('update:modelValue', emitModelValueFormat(value, formatStr))
    }
  }

  return {
    showMenu,
    formatStr,
    emitDatePickerValue,
    tempData
  }
}
