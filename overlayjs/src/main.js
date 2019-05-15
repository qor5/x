import BranDrawer from './components/BranDrawer.vue'
import BranLazyLoader from './components/BranLazyLoader.vue'
import BranPopover from 'ant-design-vue/es/popover'
import 'ant-design-vue/lib/popover/style/index.css'

(window.__branVueComponentRegisters = (window.__branVueComponentRegisters || [])).push(function(Vue){
    Vue.component("bran-drawer", BranDrawer);
    Vue.component("bran-popover", BranPopover);
    Vue.component("bran-lazy-loader", BranLazyLoader);
})

export default {
	BranDrawer,
	BranPopover,
	BranLazyLoader,
}
