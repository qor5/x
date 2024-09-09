import { fileURLToPath, URL } from 'node:url'
import { UserConfig } from 'vitepress'
import sidebar from '../sidebar.ts'
import { demoPreviewPlugin } from '@vitepress-code-preview/plugin'
import { defineConfig } from 'vitepress'
import Vuetify from 'vite-plugin-vuetify'

const nav = [
  { text: '组件文档', link: '/Components/Button/', target: '_self' },
  // { text: 'playground', link: '/playground/' },
  {
    text: 'Github',
    link: 'https://github.com/qor5/x/tree/master/ui/vuetifyx/vuetifyxjs',
    target: '_blank',
    rel: ''
  }
]

const config: UserConfig = {
  vite: {
    plugins: [Vuetify({
      autoImport: { labs: true },
      styles: {
        configFile: '../src/lib/scss/_vuetify.scss'
      }
    }),],
    ssr: {
      noExternal: ['vuetify']
    }
  },
  themeConfig: {
    sidebar,
    nav,
    search: true,
    logo: "/logo.svg"
  },

  title: 'VuetifyX UI',
  lang: 'zh-CN',
  description: '一个基于vuetify的企业级组件库',
  markdown: {
    config(md) {
      const docRoot = fileURLToPath(new URL('../', import.meta.url))
      md.use(demoPreviewPlugin, { docRoot })
    },
  }
}

export default defineConfig(config)
