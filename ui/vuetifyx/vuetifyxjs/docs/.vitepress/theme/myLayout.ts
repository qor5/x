import { defineComponent, h } from 'vue'
import theme from './theme-default/theme'

import './theme-default/custom.scss'
const { Layout } = theme
export default defineComponent({
  async mounted() {
    // 黑暗模式设配
    // const { default: theme } = await import('../../../promiseui/theme')
    const toggleTheme = (isDark: boolean) => {
      // if (!isDark) {
      //   theme.use('light')
      // } else {
      //   theme.use('dark')
      // }
    }
    const observer = new MutationObserver((entries) => {
      entries.forEach((mutation) => {
        const target = mutation.target as HTMLHtmlElement
        toggleTheme(target.classList.contains('dark'))
      })
    })
    observer.observe(document.documentElement, {
      attributes: true,
      attributeFilter: ['class']
    })
    toggleTheme(document.documentElement.classList.contains('dark'))
  },
  render() {
    return h(Layout)
  }
})
