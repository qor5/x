import { describe, it, expect } from 'vitest'

import Datepicker from '../Datepicker.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick, ref, watch } from 'vue'
import { flushPromises } from '@vue/test-utils'

it('Datepicker modelValue', async () => {
  const wrapper = mountTemplate(Datepicker, {
    modelValue: '2023-10-01'
  })
  await nextTick()
  await flushPromises()
  expect(wrapper.find('input').element.value).toContain('2023-10-01')
  expect(wrapper.find('input').element.value).not.toContain('2023-10-02')
})
