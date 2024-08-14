/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
declare const window: any

export function registerPlugins (app:App) {
  app.use(vuetify)
}

export function registerVuetify2Window() {
  window.Vuerify = vuetify
}
