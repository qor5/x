import { mergeConfig, defineConfig, configDefaults } from 'vitest/config'
import viteConfig from './vite.config'

export default mergeConfig(
  viteConfig,
  defineConfig({
    test: {
      server: {
        deps: {
          inline: ['vuetify'],
        },
      },
      environment: 'jsdom',
      exclude: [...configDefaults.exclude, 'e2e/*'],
    }
  })
)
