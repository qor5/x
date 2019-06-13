import Vue from 'vue'
import VueHighlightJS from 'vue-highlightjs'
import 'highlight.js/styles/github.css'

Vue.use(VueHighlightJS);

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push(function (Vue) {
		Vue.component("BranCode", {
			name: "BranCode",
			props: ['language'],
			template: `<pre v-highlightjs><code :class="language"><slot name="sourcecode"></slot></code></pre>`,
		});
	});
