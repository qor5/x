import { mount, VueWrapper } from '@vue/test-utils'
import { Component } from 'vue'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives
})

global.ResizeObserver = require('resize-observer-polyfill')

export function mountTemplate(component: Component, props: {}): VueWrapper {
  return mount(component, {
    props: {
      ...props
    },
    global: {
      plugins: [vuetify]
    }
  })
}
