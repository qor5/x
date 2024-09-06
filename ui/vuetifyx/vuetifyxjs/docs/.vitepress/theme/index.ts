// theme/index.ts
// import DefaultTheme from 'vitepress/theme';
import theme from './theme-default/theme.ts'
import { registerComponents } from './register-components.ts'
import { h } from 'vue'
import myLayout from './myLayout'

export default {
    // ...DefaultTheme,
    ...theme,
    Layout: () => {
      return h(myLayout)
    },
    enhanceApp({app}) {
      registerComponents(app)
    },
};
