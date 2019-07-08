import Vue, { VNode, VNodeDirective } from 'vue';
import { Core } from './core';

const app = document.getElementById('app');
if (!app) {
	throw new Error('#app required');
}

declare var window: any;

const core = new Core();

for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue);
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
// 	beforeCreate: function () {
// 		// var myOption = this.$options
// 		// console.log("props", JSON.stringify(this.$props))
// 		const tag = this.$options._componentTag;
// 		let watch = (this.$options.watch = this.$options.watch || {})

// 		watch.search = function (val) {
// 			console.log("val", val)
// 		}

// 		console.log("beforeCreate this", tag, this, this.$options)

// 	}
// })

// Vue.directive('bran', {
// 	// When the bound element is inserted into the DOM...
// 	bind: (el: HTMLElement, binding: VNodeDirective, vnode: VNode) => {
// 		core.callSetupFunc(binding.value.setupFunc, el, binding, vnode);
// 	},
// });

const vm = new Vue({
	provide: {
		core,
	},
	data: {
		current: core.componentByTemplate(app.innerHTML),
	},
	template: `
	<div id="app">
		<component :is="current"></component>
	</div>
`,
	methods: {
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},
});

vm.$mount('#app');
