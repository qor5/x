import { postcssIsolateStyles } from 'vitepress'

export default {
  plugins: [postcssIsolateStyles({
    includeFiles: [/vp-doc\.css/, /reset\.css/, /base\.css/]
  })]
}