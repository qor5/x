import { defineConfig } from 'vite'
import { fileURLToPath, URL } from 'node:url'
import { viteDemoPreviewPlugin } from '@vitepress-code-preview/plugin'
import { UserConfig } from 'vitepress';

export default defineConfig({
  plugins: [
    viteDemoPreviewPlugin()
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('../src', import.meta.url))
    }
  }
} as UserConfig)
