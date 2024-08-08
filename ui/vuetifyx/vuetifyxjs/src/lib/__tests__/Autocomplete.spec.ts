import { expect, it } from 'vitest'

import Autocomplete from '../Autocomplete.vue'
import { mountTemplate } from '@/lib/__tests__/testutils'
import { nextTick } from 'vue'
import { flushPromises } from '@vue/test-utils'
import { setupServer } from 'msw/node'
import { autocompleteHandle } from '@/lib/__tests__/handler'

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
    modelValue: 'test_1_1'
  })
  await nextTick()
  expect(wrapper.html()).toContain('test_1_1')
  expect(wrapper.html()).not.toContain('test_2_1')
})

it('Autocomplete loadData', async () => {
  const server = setupServer(...autocompleteHandle)
  server.listen({ onUnhandledRequest: 'error' })
  const wrapper = mountTemplate(Autocomplete, {
    sorting: true,
    remoteUrl: '/autocomplete',
    itemsKey: 'data',
    itemLabel: 'title',
    itemValue: 'id',
    isPaging: true
  })
  await nextTick()
  await flushPromises()
  server.close()
})
