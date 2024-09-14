/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
import i18n from './i18n'
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
import VXToolbar from '@/lib/Common/ToolBar.vue'
import VXLabel from '@/lib/Common/VXLabel.vue'
declare const window: any

const vuetifyx = {
  install: (app: App) => {
    app.component('vx-datepicker', Datepicker)
    app.component('vx-datetimepicker', Datetimepicker)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-filter', Filter)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-textdatepicker', TextDatepicker)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-messagelistener', MessageListener)
    app.component('vx-overlay', Overlay)
    app.component('vx-field', VXField)
    app.component('vx-select', VXSelect)
    app.component('vx-toolbar', VXToolbar)
    app.component('vx-label', VXLabel)
  }
}

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
  app.use(vuetifyx)
}

export function registerVuetify2Window() {
  window.Vuetify = vuetify
}
