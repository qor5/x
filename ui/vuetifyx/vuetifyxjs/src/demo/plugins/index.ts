/**
 * plugins/index.ts
 *
 * Automatically included in `./src/main.ts`
 */

// Plugins
import vuetify from './vuetify'
import i18n from './i18n'
import { vuetifyProTipTap } from './tiptap'

// Types
import type { App } from 'vue'

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
  app.use(vuetifyProTipTap)
  // fix warning injected property "decorationClasses" is a ref and will be auto-unwrapped
  // https://github.com/ueberdosis/tiptap/issues/1719
  // app.config.unwrapInjectedRef = true
}
