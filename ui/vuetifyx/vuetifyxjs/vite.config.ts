// Plugins
import Components from 'unplugin-vue-components/vite'
import { Vuetify3Resolver } from 'unplugin-vue-components/resolvers'
import Vue from '@vitejs/plugin-vue'
import Vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
import ViteFonts from 'unplugin-fonts/vite'
import { resolve } from 'path'
import vueJsx from '@vitejs/plugin-vue-jsx'

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    // minify: false,
    lib: {
      entry: resolve(__dirname, 'src/lib/main.ts'),
      formats: ['umd'],
      name: 'vuetifyxjs'
    },
    copyPublicDir: false,
    rollupOptions: {
      external: ['vue', 'vuetify'],
      output: {
        assetFileNames: (assetInfo) => {
          return 'vuetifyxjs.css'
        },
        globals: {
          vue: 'Vue',
          vuetify: 'Vuetify'
        }
      }
    }
  },

  plugins: [
    Vue({
      template: { transformAssetUrls }
    }),
    vueJsx(),
    // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
    Vuetify({ autoImport: { labs: true } }),
    Components({
      dts: true,
      dirs: ['src/demo/components', 'src/lib'],
      resolvers: [Vuetify3Resolver()],
      include: [/\.vue$/]

    }),
    ViteFonts({
      google: {
        families: [{
          name: 'Roboto',
          styles: 'wght@100;300;400;500;700;900'
        }]
      }
    })
  ],
  define: { 'process.env': {} },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: [
      '.js',
      '.json',
      '.jsx',
      '.mjs',
      '.ts',
      '.tsx',
      '.vue'
    ]
  },
  server: {
    port: 3000
  }

})
