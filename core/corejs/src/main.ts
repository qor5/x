import Vue from 'vue';
import { Core } from './core';
import { newFormWithStates } from './form';

const app = document.getElementById('app');
if (!app) {
	throw new Error('#app required');
}

declare var window: any;

for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue);
}

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
		const rootChangeCurrent = (this.$root as any).changeCurrent;
		const core = new Core(new FormData(), rootChangeCurrent, this.changeCurrent);

		core.fetchEvent(ef, {})
			.then((r) => {
				self.current = core.componentByTemplate(r.schema, afterLoaded);
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
		const ssd = window.__serverSideData__;
		const states = (ssd && ssd.states) || {};
		const core = new Core(newFormWithStates(states), this.changeCurrent, this.changeCurrent);
		this.current = core.componentByTemplate(app.innerHTML);
	},

	data() {
		return {
			current: null,
		};
	},

});

vm.$mount('#app');
