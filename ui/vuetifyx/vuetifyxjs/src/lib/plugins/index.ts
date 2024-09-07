/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
import i18n from './i18n'
import { vuetifyProTipTap } from './tiptap'

declare const window: any

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
  app.use(vuetifyProTipTap)
  // fix warning injected property "decorationClasses" is a ref and will be auto-unwrapped
  // https://github.com/ueberdosis/tiptap/issues/1719
  // app.config.unwrapInjectedRef = true
}

export function registerVuetify2Window() {
  window.Vuetify = vuetify
}
