// Plugins
import Components from 'unplugin-vue-components/vite'
import { Vuetify3Resolver } from 'unplugin-vue-components/resolvers'
import Vue from '@vitejs/plugin-vue'
import Vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
// import ViteFonts from 'unplugin-fonts/vite'
import { resolve } from 'path'
import vueJsx from '@vitejs/plugin-vue-jsx'

// Utilities
import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    cssCodeSplit:false,
    /**
     * ⚠️ warning: ⚠️
     * if configured lib options, vite will always inline assets in css includes, such as fonts,
     * it will create a huge size css about 7mb!
     */
    // lib: {
    //   entry: resolve(__dirname, 'src/lib/main.ts'),
    //   formats: ['umd'],
    //   name: 'vuetifyx'
    // },
    copyPublicDir: false,
    rollupOptions: {
      input: resolve(__dirname, 'src/lib/main.ts'),
      external: ['vue'],
      output: {
        format: 'umd',
        name: 'vuetifyx',

        globals: {
          vue: 'Vue',
          vuetify: 'Vuetify'
        },
        chunkFileNames: `[name].js`,
        entryFileNames: 'vuetifyx.min.js',
        assetFileNames: assetInfo => {
          if (assetInfo.name && assetInfo.name.endsWith('.css')) {
            return 'assets/vuetifyx.min.css';
          }
          return 'assets/[name].[ext]';
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
    Vuetify({
      autoImport: { labs: true },
      styles: {
        configFile: 'src/lib/scss/_vuetify.scss'
      }
    }),
    Components({
      dts: true,
      dirs: ['src/demo/components', 'src/lib'],
      resolvers: [Vuetify3Resolver()],
      include: [/\.vue$/]
    }),
    // ViteFonts({
    //   google: {
    //     families: [{
    //       name: 'Roboto',
    //       styles: 'wght@100;300;400;500;700;900'
    //     }]
    //   }
    // })
  ],

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
    extensions: ['.js', '.json', '.jsx', '.mjs', '.ts', '.tsx', '.vue']
  },
  server: {
    port: 3000
  }
})
