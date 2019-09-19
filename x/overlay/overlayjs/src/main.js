import BranDrawer from "./components/BranDrawer.jsx";
import BranPopover from "./components/BranPopover.jsx";
import BranDialog from "./components/BranDialog.jsx";

import Vue from "vue";
if (!window.Vue) {
	window.Vue = Vue;
}

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push(function(Vue) {
	Vue.component("BranDrawer", BranDrawer);
	Vue.component("BranPopover", BranPopover);
	Vue.component("BranDialog", BranDialog);
});

export default {
	BranDrawer,
	BranPopover,
	BranDialog
};
