import TipTapEditor from './editor.vue'

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue) => {
		Vue.component('tiptap-editor', TipTapEditor)
	});
