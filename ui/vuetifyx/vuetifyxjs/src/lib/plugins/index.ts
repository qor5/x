/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
import i18n from './i18n'
declare const window: any

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
}

export function registerVuetify2Window() {
  window.Vuetify = vuetify
}
