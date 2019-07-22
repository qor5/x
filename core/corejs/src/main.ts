import Vue from 'vue';
import { Core } from './core';

const app = document.getElementById('app');
if (!app) {
	throw new Error('#app required');
}

declare var window: any;

for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue);
}

const core = new Core();

Vue.component('BranLazyLoader', {
	name: 'BranLazyLoader',
	props: ['loaderFunc', 'visible', 'afterLoaded'],
	template: `
		<div class="bran-lazy-loader" v-if="visible">
			<component :is="current"></component>
		</div>
	`,

	mounted() {
		const ef = this.loaderFunc;
		const afterLoaded = this.afterLoaded;
		const self = this;
		const doFunc = this.changeCurrent;
		core.fetchEvent(ef, {})
			.then((r) => {
				self.current = core.componentByTemplate(doFunc, r.schema, afterLoaded);
			});
	},

	data() {
		return {
			current: null,
		};
	},

	methods: {
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},
});


const vm = new Vue({
	provide: {
		core,
	},
	template: `
	<div id="app" v-cloak>
		<component :is="current"></component>
	</div>
`,
	methods: {
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},

	mounted() {
		core.rootChangeCurrent = this.changeCurrent;
		this.current = core.componentByTemplate(this.changeCurrent, app.innerHTML);
	},

	data() {
		return {
			current: null,
		};
	},

});

vm.$mount('#app');
