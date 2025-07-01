import dayjs from 'dayjs'
import _ from 'lodash'
import { ref, computed } from 'vue'
import { EnhancedDateParser } from '@/lib/utils/dateParser'

export type datePickerType = 'datepicker' | 'datetimepicker'

// the duplicate logic on vx-date-picker and vx-range-picker
export function useDatePicker<EmitFnType>(props: any, emit: EmitFnType) {
  const showMenu = ref(false)
  const tempData = ref<string | number | (string | number)[]>()
  const formatMapDefault = {
    datepicker: 'YYYY-MM-DD',
    datetimepicker: 'YYYY-MM-DD HH:mm:ss'
  }

  const formatStr = computed(() => props.format || formatMapDefault[props.type as datePickerType])

  function emitModelValueFormat(value: number | string | (number | string)[]) {
    const formatedValue = Array.isArray(value)
      ? value.map((item) => {
          if (!item) return ''

          // 使用增强的日期解析器进行解析
          const parsed = EnhancedDateParser.parseDate(item)
          return parsed
            ? parsed.format(formatStr.value || formatMapDefault[props.type as datePickerType])
            : ''
        })
      : value
        ? (() => {
            const parsed = EnhancedDateParser.parseDate(value)
            return parsed
              ? parsed.format(formatStr.value || formatMapDefault[props.type as datePickerType])
              : ''
          })()
        : ''
    return formatedValue
  }

  function emitDatePickerValue(
    value: number | string | (number | string)[],
    {
      needConfirm,
      extraEmitEvents
    }: {
      needConfirm?: boolean
      extraEmitEvents?: string[]
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

    extraEmitEvents?.forEach((event) => {
      ;(emit as any)(event, emitModelValueFormat(value))
    })
  }

  return {
    showMenu,
    formatStr,
    emitDatePickerValue,
    tempData
  }
}
