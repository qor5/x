import { App } from 'vue'

import Datepicker from '@/lib/Datepicker.vue'
import Datetimepicker from '@/lib/Datetimepicker.vue'
import SelectMany from '@/lib/SelectMany.vue'
import LinkageSelect from '@/lib/LinkageSelect.vue'
import LinkageSelectRemote from '@/lib/LinkageSelectRemote/index.vue'
import Autocomplete from '@/lib/Autocomplete.vue'
import TextDatepicker from '@/lib/TextDatepicker.vue'
import Filter from '@/lib/Filter/index.vue'
import RestoreScrollListener from '@/lib/RestoreScrollListener.vue'
import ScrollIframe from '@/lib/ScrollIframe.vue'
import draggable from 'vuedraggable'
import SendVariables from '@/lib/SendVariables.vue'
import MessageListener from '@/lib/MessageListener.vue'
import Overlay from '@/lib/Overlay.vue'
import TiptapEditor from './TiptapEditor.vue'
import { registerPlugins, registerVuetify2Window } from '@/lib/plugins'

const vuetifyx = {
  install: (app: App) => {
    app.component('vx-datepicker', Datepicker)
    app.component('vx-datetimepicker', Datetimepicker)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-linkageselect-remote', LinkageSelectRemote)
    app.component('vx-filter', Filter)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-textdatepicker', TextDatepicker)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-messagelistener', MessageListener)
    app.component('vx-overlay', Overlay)
    app.component('vx-tiptap-editor', TiptapEditor)
  }
}
declare const window: any

// export vuetifyInstance to window, thus qor5/web/core.js can use it.
registerVuetify2Window()

window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || []
window.__goplaidVueComponentRegisters.push((app: App, vueOptions: any): any => {
  // register any theme、build-in components from vuetify
  registerPlugins(app)
  // register any business components based on vuetify build-in components
  app.use(vuetifyx)
})
