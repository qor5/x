import Drawer from './components/BranDrawer.vue'
import Popover from 'ant-design-vue/es/popover'
import 'ant-design-vue/lib/popover/style/index.css'

(window.__branVueComponentRegisters = (window.__branVueComponentRegisters || [])).push(function(Vue){
    Vue.component("bran-drawer", Drawer);
    Vue.component("bran-popover", Popover);
})

export default {
	Drawer,
	Popover,
}
