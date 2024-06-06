import { describe, it, expect } from 'vitest'

import Filter from '../Filter/index.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick, ref, watch } from 'vue'

it('Filter modelValue', async () => {
  const internalValue = [
    {
      key: 'TestStringItem',
      label: 'TestStringItem',
      itemType: 'StringItem',
      selected: true,
      folded: true,
      valueIs: 'active'
    }
  ]
  const value = ref('')
  const wrapper = mountTemplate(Filter, {
    internalValue: internalValue,
    modelValue: value
  })
  await nextTick()
  expect(wrapper.html()).contain('active')
  expect(wrapper.html()).not.contain('deactivate')
})
