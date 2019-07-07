import Vue, { VNode, VNodeDirective } from 'vue';
import { Core } from './core';

const app = document.getElementById('branRoot');
if (!app) {
	throw new Error('#branRoot required');
}

declare var window: any;

const core = new Core();

for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue, core);
}

Vue.component('BranLazyLoader', {
	name: 'BranLazyLoader',
	props: ['loaderFunc', 'visible', 'afterLoaded'],
	template: `
		<div class="bran-lazy-loader" v-if="visible">
			<component :is="lazyloader"></component>
		</div>
	`,
	data() {
		const ef = this.loaderFunc;
		const afterLoaded = this.afterLoaded;
		if (!ef) {
			return {
				lazyloader: {
					render() {
						return null;
					},
				},
			};
		}
		return {
			lazyloader(): any {
				return core.fetchEvent(ef, {})
					.then((r) => {
						return core.componentByTemplate(r.schema, afterLoaded);
					});
			},
		};
	},
});

// Vue.mixin({
// 	props: ['bran-field-name'],
// 	created: function () {
// 		// var myOption = this.$options
// 		// console.log("props", JSON.stringify(this.$props))
// 		console.log("this.$el", this.$el, this)
// 	}
// })

Vue.directive('bran', {
	// When the bound element is inserted into the DOM...
	inserted: (el: HTMLElement, binding: VNodeDirective, vnode: VNode) => {
		core.callSetupFunc(binding.value.setupFunc, el, binding, vnode);
	},
});

const vm = new Vue({
	data: {
		current: core.componentByTemplate(app.innerHTML),
	},
	template: `
	<div id="branRoot">
		<component :is="current"></component>
	</div>
`,
	methods: {
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},
});

vm.$mount('#branRoot');
