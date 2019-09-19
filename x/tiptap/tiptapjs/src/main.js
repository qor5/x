// @snippet_begin(GoPlaidRegisterVueComponentSample)
import TipTapEditor from './editor.vue'

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue) => {
		Vue.component('tiptap-editor', TipTapEditor)
	});

// @snippet_end
