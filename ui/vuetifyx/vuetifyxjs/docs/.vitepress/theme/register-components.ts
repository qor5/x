import { App } from 'vue'
import DemoPreview, { useComponents } from '@vitepress-code-preview/container'
import '@vitepress-code-preview/container/dist/style.css'
import { registerPlugins } from '@/lib/plugins'

export function registerComponents(app: App) {
  app.use(registerPlugins)
  useComponents(app, DemoPreview)
}
