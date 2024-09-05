import { vuePlugin, addImportMap } from 'vitepress-demo-editor'
import { App } from 'vue'
import 'vitepress-demo-editor/dist/style.css'
import { registerPlugins } from '../../../src/lib/plugins'
let first = true
export function registerComponents(app: App) {
  // addImportMap('@vicons/ionicons5', icons)
  app.use(registerPlugins)
  app.use(vuePlugin, {
    onMonacoCreated(monaco) {
      // monaco.languages.typescript.typescriptDefaults.addExtraLib(
      //   `
      //   declare module 'promiseui-vue' { ${promiseuiType} }
      // `,
      //   `promiseui-vue`
      // )
    }
  })

  app.mixin({
    async mounted() {
      if (!first) return
      first = false
      // await import('../../../promiseui').then((promiseUI) => {
      //   addImportMap('promiseui-vue', promiseUI)

      //   app.use(promiseUI.default)
      //   emitUILoaded()
      // })
    }
  })
}
