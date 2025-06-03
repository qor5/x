/**
 * plugins/index.js
 */

// Plugins
import { type App } from 'vue'
import vuetify from './vuetify'
import i18n from './i18n'
import { vuetifyProTipTap } from './tiptap'

import DatepickerOld from '@/lib/Datepicker.vue'
import DatetimepickerOld from '@/lib/Datetimepicker.vue'
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
import VXRangepicker from '@/lib/Form/VXDatePicker/RangePicker.vue'
import VXDatepicker from '@/lib/Form/VXDatePicker/DatePicker.vue'
import VXTimePicker from '@/lib/Form/VXDatePicker/TimePicker.vue'
import VXField from '@/lib/Form/VXField.vue'
import VXSelect from '@/lib/Form/VXSelect.vue'
import VXCheckbox from '../Form/VXCheckbox.vue'
import VXToolbar from '@/lib/Common/VXToolBar.vue'
import VXLabel from '@/lib/Common/VXLabel.vue'
import VXDialog from '@/lib/Common/VXDialog.vue'
import VXPagination from '@/lib/Common/VXPagination.vue'
import VXBtn from '@/lib/Common/VXBtn.vue'
import VXBtnGroup from '@/lib/Common/VXBtnGroup.vue'
import VXChip from '@/lib/Common/VXChip.vue'
import VXAvatar from '@/lib/VXAvatar.vue'
import VXTabs from '@/lib/Tabs/VXTabs.vue'
import VXIframeEmitter from '@/lib/IframeEmitter.vue'
import VXBreadcrumbs from '@/lib/Breadcrumbs/VXBreadcrumbs.vue'
import VXTreeview from '@/lib/Treeview/VXTreeview.vue'
import TiptapEditor from '@/lib/TiptapEditor/TiptapEditor.vue'
import LinkageSelectRemote from '@/lib/LinkageSelectRemote/index.vue'
import VXSegmentForm from '@/lib/Form/VXSegmentForm/index.vue'
import VXModelProxy from '@/lib/VXModelProxy.vue'
import VXChart from '@/lib/Chart/VXChart.vue'

declare const window: any

const vuetifyx = {
  install: (app: App) => {
    app.component('vx-chart', VXChart)
    app.component('vx-date-picker', VXDatepicker)
    app.component('vx-time-picker', VXTimePicker)
    app.component('vx-range-picker', VXRangepicker)
    app.component('vx-datepicker', DatepickerOld)
    app.component('vx-datetimepicker', DatetimepickerOld)
    app.component('vx-textdatepicker', TextDatepicker)
    app.component('vx-btn', VXBtn)
    app.component('vx-btn-group', VXBtnGroup)
    app.component('vx-chip', VXChip)
    app.component('vx-select', VXSelect)
    app.component('vx-selectmany', SelectMany)
    app.component('vx-linkageselect', LinkageSelect)
    app.component('vx-autocomplete', Autocomplete)
    app.component('vx-linkageselect-remote', LinkageSelectRemote)
    app.component('vx-checkbox', VXCheckbox)
    app.component('vx-field', VXField)
    app.component('vx-label', VXLabel)
    app.component('vx-dialog', VXDialog)
    app.component('vx-tiptap-editor', TiptapEditor)
    app.component('vx-overlay', Overlay)
    app.component('vx-filter', Filter)
    app.component('vx-tabs', VXTabs)
    app.component('vx-pagination', VXPagination)
    app.component('vx-avatar', VXAvatar)
    app.component('vx-toolbar', VXToolbar)
    app.component('vx-draggable', draggable)
    app.component('vx-restore-scroll-listener', RestoreScrollListener)
    app.component('vx-scroll-iframe', ScrollIframe)
    app.component('vx-send-variables', SendVariables)
    app.component('vx-breadcrumbs', VXBreadcrumbs)
    app.component('vx-treeview', VXTreeview)
    app.component('vx-messagelistener', MessageListener)
    app.component('vx-segment-form', VXSegmentForm)
    app.component('vx-iframe-emitter', VXIframeEmitter)
    app.component('vx-model-proxy', VXModelProxy)
  }
}

export function registerPlugins(app: App) {
  app.use(i18n)
  app.use(vuetify)
  app.use(vuetifyx)
  app.use(vuetifyProTipTap as any)
  // fix warning injected property "decorationClasses" is a ref and will be auto-unwrapped
  // https://github.com/ueberdosis/tiptap/issues/1719
  // app.config.unwrapInjectedRef = true
}

export function registerVuetify2Window() {
  window.Vuetify = vuetify
}
