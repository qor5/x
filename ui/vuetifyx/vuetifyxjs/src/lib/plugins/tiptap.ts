import { VuetifyTiptap, VuetifyViewer, createVuetifyProTipTap } from 'vuetify-pro-tiptap'
import 'vuetify-pro-tiptap/style.css'

export const vuetifyProTipTap = createVuetifyProTipTap({
  lang: 'en', 
  components: {
    VuetifyTiptap,
    VuetifyViewer
  },
  extensions: []
})
