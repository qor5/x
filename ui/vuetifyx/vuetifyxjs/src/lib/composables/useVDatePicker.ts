import { ref, computed, Ref } from 'vue'
import dayjs from 'dayjs'

export function useVDatePickerTimeChange(dateOfPicker: Ref) {
  const displayedMonth = ref()
  const displayedYear = ref()

  function setDisplayedYearAndMonth(timeStamp?: string | number | Date) {
    let year: number, month: number

    if (!timeStamp) {
      year = new Date().getFullYear()
      month = new Date().getMonth()
    } else {
      ;[year, month] = dayjs(timeStamp).format('YYYY-MM').split('-').map(Number)
      month = month - 1
    }

    displayedYear.value = year
    displayedMonth.value = month
  }

  function onYearOrMonthChange(value: number | unknown, type: 'year' | 'month' | 'modelValue') {
    if (type === 'modelValue') {
      // console.log('modelValue', value)
      return
    }
    dateOfPicker.value = null
  }

  const valueChangedWithoutSaved = computed(() => {
    return dateOfPicker.value === null
  })

  return {
    displayedMonth,
    displayedYear,
    setDisplayedYearAndMonth,
    onYearOrMonthChange,
    valueChangedWithoutSaved
  }
}
