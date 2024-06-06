import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'
import {resolve} from "path";

// https://vitejs.dev/config/
export default defineConfig({
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
