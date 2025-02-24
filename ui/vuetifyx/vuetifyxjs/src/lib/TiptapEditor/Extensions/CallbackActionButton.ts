import type { GeneralOptions } from 'vuetify-pro-tiptap'
import { Extension } from '@tiptap/core'

import CallbackActionButton from './CallbackActionButton.vue'

export type CallbackOptions = GeneralOptions<CallbackOptions>

export default Extension.create<CallbackOptions>({
  name: 'callback',
  addOptions() {
    return {
      divider: false,
      spacer: false,
      button: ({ editor, extension, t }) => {
        return {
          component: CallbackActionButton,
          componentProps: {
            editor: editor,
            extension: extension,
            t: t
          }
        }
      }
    }
  }
})
