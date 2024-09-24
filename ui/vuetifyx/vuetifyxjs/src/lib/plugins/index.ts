/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
import i18n from './i18n'
import { vuetifyProTipTap } from './tiptap'

import Datepicker from '@/lib/Datepicker.vue'
import Datetimepicker from '@/lib/Datetimepicker.vue'
import SelectMany from '@/lib/SelectMany.vue'
import LinkageSelect from '@/lib/LinkageSelect.vue'
import Autocomplete from '@/lib/Autocomplete.vue'
import TextDatepicker from '@/lib/TextDatepicker.vue'
import Filter from '@/lib/Filter/index.vue'
import RestoreScrollListener from '@/lib/RestoreScrollListener.vue'
import ScrollIframe from '@/lib/ScrollIframe.vue'
import draggable from 'vuedraggable'
import SendVariables from '@/lib/SendVariables.vue'
import MessageListener from '@/lib/MessageListener.vue'
import Overlay from '@/lib/Overlay.vue'
import VXField from '@/lib/Form/VXField.vue'
import VXSelect from '@/lib/Form/VXSelect.vue'
import VXToolbar from '@/lib/Common/VXToolBar.vue'
import VXLabel from '@/lib/Common/VXLabel.vue'
import VXDialog from '@/lib/Common/VXDialog.vue'
import TiptapEditor from '@/lib/TiptapEditor/TiptapEditor.vue'
import LinkageSelectRemote from '@/lib/LinkageSelectRemote/index.vue'

declare const window: any

const vuetifyx = {
  install: (app: App) => {
    // datepicker
    app.component('vx-datepicker', Datepicker)
    app.component('vx-datetimepicker', Datetimepicker)
    app.component('vx-textdatepicker', TextDatepicker)
    // select
    app.component('vx-select', VXSelect)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-linkageselect-remote', LinkageSelectRemote)
    // field and label
    app.component('vx-field', VXField)
    app.component('vx-label', VXLabel)
    // dialog
    app.component('vx-dialog', VXDialog)
    // editor
    app.component('vx-tiptap-editor', TiptapEditor)
    // overlay
    app.component('vx-overlay', Overlay)
    // filter
    app.component('vx-filter', Filter)
    // others
    app.component('vx-toolbar', VXToolbar)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-messagelistener', MessageListener)
  }
}

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
  app.use(vuetifyx)
  app.use(vuetifyProTipTap)
  // fix warning injected property "decorationClasses" is a ref and will be auto-unwrapped
  // https://github.com/ueberdosis/tiptap/issues/1719
  // app.config.unwrapInjectedRef = true
}

export function registerVuetify2Window() {
  window.Vuetify = vuetify
}
