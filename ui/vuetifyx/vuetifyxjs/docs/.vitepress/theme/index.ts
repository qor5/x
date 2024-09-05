// theme/index.ts
import DefaultTheme from 'vitepress/theme';
import { registerComponents } from './register-components.ts'

export default {
    ...DefaultTheme,
    enhanceApp({app}) {
      registerComponents(app)
    },
};
