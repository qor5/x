import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from "path";
import svgLoader from 'vite-svg-loader';

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
    return {
        define: {
            'process.env.NODE_ENV': JSON.stringify(mode),
        },
        build: {
            // minify: false,
            lib: {
                entry: resolve(__dirname, 'src/lib/main.ts'),
                formats: ['umd'],
                name: 'tiptap'
            },
            copyPublicDir: false,
            // @snippet_begin(TipTapVueConfig)
            rollupOptions: {
                external: ['vue'],
                output: {
                    assetFileNames: (assetInfo) => {
                        return 'tiptap.css'
                    },
                    globals: {
                        vue: 'Vue',
                    }
                }
            }
            // @snippet_end
        },
        plugins: [
            vue(),
            svgLoader()
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        },
    }
})
