import { describe, it, expect } from 'vitest'

import Datetimepicker from '../Datetimepicker.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick, watch } from 'vue'
import { flushPromises } from '@vue/test-utils'

it('Datetimepicker modelValue', async () => {
  const wrapper = mountTemplate(Datetimepicker, {
    modelValue: '2023-10-01 22:11'
  })
  await nextTick()
  await flushPromises()
  expect(wrapper.find('input').element.value).toContain('2023-10-01 22:11')
  expect(wrapper.find('input').element.value).not.toContain('2023-10-01 22:22')
})
