import { describe, it, expect } from 'vitest'
import dayjs from 'dayjs'

import DatePickerBase from '@/lib/Form/VXDatePicker/DatePickerBase.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'

function lastEmittedModelValue(wrapper: ReturnType<typeof mountTemplate>): number {
  const events = wrapper.emitted('update:modelValue') as unknown as [number][]
  return events[events.length - 1][0]
}

describe('DatePickerBase arrow navigation across a year boundary (KGM-3560)', () => {
  it('going back from January lands on December of the previous year', async () => {
    const wrapper = mountTemplate(DatePickerBase, {
      modelValue: dayjs('2026-01-11').valueOf(),
      type: 'datepicker'
    })
    await wrapper.find('[data-testid="prev-month"]').trigger('click')

    expect(dayjs(lastEmittedModelValue(wrapper)).format('YYYY-MM-DD')).toBe('2025-12-11')
  })

  it('going forward from December lands on January of the next year', async () => {
    const wrapper = mountTemplate(DatePickerBase, {
      modelValue: dayjs('2025-12-11').valueOf(),
      type: 'datepicker'
    })
    await wrapper.find('[data-testid="next-month"]').trigger('click')

    expect(dayjs(lastEmittedModelValue(wrapper)).format('YYYY-MM-DD')).toBe('2026-01-11')
  })

  it('navigating within the same year is unaffected', async () => {
    const wrapper = mountTemplate(DatePickerBase, {
      modelValue: dayjs('2026-05-11').valueOf(),
      type: 'datepicker'
    })
    await wrapper.find('[data-testid="prev-month"]').trigger('click')

    expect(dayjs(lastEmittedModelValue(wrapper)).format('YYYY-MM-DD')).toBe('2026-04-11')
  })
})
