import BranDrawer from "./components/BranDrawer.jsx";
import BranPopover from "./components/BranPopover.jsx";
import BranDialog from "./components/BranDialog.jsx";

import Vue from "vue";
if (!window.Vue) {
	window.Vue = Vue;
}

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push(function(Vue) {
	Vue.component("bran-drawer", BranDrawer);
	Vue.component("bran-popover", BranPopover);
	Vue.component("bran-dialog", BranDialog);
});

export default {
	BranDrawer,
	BranPopover,
	BranDialog
};
