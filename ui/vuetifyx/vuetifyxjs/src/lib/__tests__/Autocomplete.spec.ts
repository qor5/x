import { describe, it, expect } from 'vitest'

import Autocomplete from '../Autocomplete.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick, watch } from 'vue'

it('Autocomplete modelValue', async () => {
  const pageSize = 1
  const page = 1
  const items = []
  for (let i = 1; i <= pageSize; i++) {
    items.push({
      icon: `https://cdn.vuetifyjs.com/images/lists/${i}.jpg`,
      text: `test_${page}_${i}`,
      value: (pageSize * (page - 1) + i).toFixed()
    })
  }
  const wrapper = mountTemplate(Autocomplete, {
    items: items,
    sorting: true,
    modelValue: '1'
  })
  await nextTick()
  expect(wrapper.html()).toContain('test_1_1')
  expect(wrapper.html()).not.toContain('test_2_1')
})
