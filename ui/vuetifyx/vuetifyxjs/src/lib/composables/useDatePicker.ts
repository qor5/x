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

  function emitModelValueFormat(value: number | string | (number | string)[]) {
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
    // console.log(formatedValue, formatStr)
    return formatedValue
  }

  function emitDatePickerValue(
    value: number | string | (number | string)[],
    {
      needConfirm
    }: {
      needConfirm?: boolean
    } = {
      needConfirm: false
    }
  ) {
    if (needConfirm) {
      if (Array.isArray(value)) {
        if (!tempData.value || !Array.isArray(tempData.value)) {
          tempData.value = emitModelValueFormat(value)
        } else {
          tempData.value = emitModelValueFormat(
            value.map((item, i) => item || (tempData.value as (number | string)[])[i])
          )
        }
      } else {
        tempData.value = emitModelValueFormat(value)
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
