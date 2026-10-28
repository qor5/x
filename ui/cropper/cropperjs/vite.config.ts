import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'
import {resolve} from "path";

// https://vitejs.dev/config/
export default defineConfig({
  // Since vite 7.3, minification strips third-party license banners by default.
  // This directory's dist is go:embed-ed and redistributed with qor5/x, and MIT
  // and friends require the copyright notice to survive redistribution — so keep
  // them explicitly rather than relying on a default that has already changed once.
  esbuild: { legalComments: 'inline' },
    build: {
        // minify: false,
        lib: {
            entry: resolve(__dirname, 'src/lib/main.ts'),
            formats: ['umd'],
            name: 'cropper'
        },
        copyPublicDir: false,
        rollupOptions: {
            external: ['vue'],
            output: {
                assetFileNames: (assetInfo) => {
                    return 'cropperjs.css'
                },
                globals: {
                    vue: 'Vue',
                }
            }
        }
    },
    plugins: [
        vue(),
        VueDevTools(),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
