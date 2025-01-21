import {
  VuetifyTiptap,
  VuetifyViewer,
  createVuetifyProTipTap
  // @ts-ignore
} from '../TiptapEditor/source/lib/vuetify-pro-tiptap'
import '../TiptapEditor/source/lib/style.css'

export const vuetifyProTipTap = createVuetifyProTipTap({
  lang: 'en',
  components: {
    VuetifyTiptap,
    VuetifyViewer
  },
  extensions: []
})
