import { it, expect } from 'vitest'

import TextDatepicker from '../TextDatepicker.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick } from 'vue'

it('Datetimepicker modelValue', async () => {
  const wrapper = mountTemplate(TextDatepicker, {
    modelValue: '2023-10-01'
  })
  await nextTick()
  expect(wrapper.find('input').element.value).toContain('2023-10-01')
  expect(wrapper.find('input').element.value).not.toContain('2023-10-02')
})
