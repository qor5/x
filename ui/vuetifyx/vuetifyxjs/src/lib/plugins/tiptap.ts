import {
  VuetifyTiptap,
  VuetifyViewer,
  createVuetifyProTipTap
  // @ts-ignore
} from '../TiptapEditor/lib/vuetify-pro-tiptap'
import '../TiptapEditor/lib/style.css'

export const vuetifyProTipTap = createVuetifyProTipTap({
  lang: 'en',
  components: {
    VuetifyTiptap,
    VuetifyViewer
  },
  extensions: []
})
