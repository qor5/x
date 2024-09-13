import { App } from 'vue'
import { registerPlugins, registerVuetify2Window } from '@/lib/plugins'

declare const window: any

// export vuetifyInstance to window, thus qor5/web/core.js can use it.
registerVuetify2Window()

window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || []
window.__goplaidVueComponentRegisters.push((app: App, vueOptions: any): any => {
  registerPlugins(app)
})
