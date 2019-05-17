import BranDrawer from './components/BranDrawer.jsx'
import BranPopover from './components/BranPopover.jsx'

import Vue from 'vue'
if (!window.Vue) {
	window.Vue = Vue
}

(window.__branVueComponentRegisters = (window.__branVueComponentRegisters || [])).push(function(Vue){
    Vue.component("bran-drawer", BranDrawer);
    Vue.component("bran-popover", BranPopover);
})

export default {
	BranDrawer,
	BranPopover,
}
