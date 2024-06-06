import { describe, it, expect } from 'vitest'

import SelectMany from '../SelectMany.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick, watch } from 'vue'

it('SelectMany modelValue', async () => {
  const wrapper = mountTemplate(SelectMany, {
    items: [
      {
        id: '1',
        text: 'ScanDa Adams',
        image: 'https://cdn.vuetifyjs.com/images/lists/1.jpg'
      },
      {
        id: '2',
        text: 'Ali Connors',
        image: 'https://cdn.vuetifyjs.com/images/lists/2.jpg'
      },
      {
        id: '3',
        text: 'Ali DE',
        image: 'https://cdn.vuetifyjs.com/images/lists/3.jpg'
      },
      {
        id: '4',
        text: 'Bogn',
        image: 'https://cdn.vuetifyjs.com/images/lists/4.jpg'
      }
    ],
    modelValue: ['1', '2']
  })
  await nextTick()
  expect(wrapper.html()).toContain('ScanDa Adams')
  expect(wrapper.html()).toContain('Ali Connors')
  expect(wrapper.html()).not.toContain('Bogn')
})
