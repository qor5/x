import { describe, it, expect } from 'vitest'
import dayjs from 'dayjs'
import { nextTick } from 'vue'
import { flushPromises } from '@vue/test-utils'

import DatePickerBase from '../Form/VXDatePicker/DatePickerBase.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'

async function mountWithDate(modelValue: string) {
  const wrapper = mountTemplate(DatePickerBase, { modelValue })
  await nextTick()
  await flushPromises()
  return wrapper
}

function lastEmittedDate(wrapper: ReturnType<typeof mountTemplate>) {
  const emitted = wrapper.emitted('update:modelValue')!
  const ts = emitted[emitted.length - 1][0] as number
  return dayjs(ts).format('YYYY-MM-DD')
}

describe('DatePickerBase arrow navigation across a year boundary', () => {
  // v-date-picker updates year first, then month, in the same tick
  // (see Vuetify's onClickPrevMonth/onClickNextMonth), so the handler
  // receives update:year and update:month back-to-back before Vue's
  // reactivity flushes.
  it('keeps the year when navigating backward from January', async () => {
    const wrapper = await mountWithDate('2026-01-11')
    const picker = wrapper.findComponent({ name: 'VDatePicker' })

    picker.vm.$emit('update:year', 2025)
    picker.vm.$emit('update:month', 11)

    expect(lastEmittedDate(wrapper)).toBe('2025-12-11')
  })

  it('keeps the year when navigating forward from December', async () => {
    const wrapper = await mountWithDate('2025-12-11')
    const picker = wrapper.findComponent({ name: 'VDatePicker' })

    picker.vm.$emit('update:year', 2026)
    picker.vm.$emit('update:month', 0)

    expect(lastEmittedDate(wrapper)).toBe('2026-01-11')
  })

  it('navigates within the same year', async () => {
    const wrapper = await mountWithDate('2026-05-11')
    const picker = wrapper.findComponent({ name: 'VDatePicker' })

    picker.vm.$emit('update:month', 3)

    expect(lastEmittedDate(wrapper)).toBe('2026-04-11')
  })
})
