import Vue, { VueConstructor } from 'vue';
import { Core } from './core';
import { newFormWithStates } from './utils';

const app = document.getElementById('app');
if (!app) {
	throw new Error('#app required');
}

declare var window: any;

const vueOptions = {};
for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue, vueOptions);
}

window.branLazyPortals = {};

const ssd = window.__serverSideData__;
const states = (ssd && ssd.states) || {};

const form = newFormWithStates(states);

interface DynaCompData {
	current: VueConstructor | null;
}

Vue.component('BranLazyPortal', {
	name: 'BranLazyPortal',
	props: ['loaderFunc', 'visible', 'afterLoaded', 'portalName'],
	template: `
		<div class="bran-lazy-loader" v-if="visible">
			<component :is="current"></component>
		</div>
	`,

	mounted() {
		const pn = this.$props.portalName;
		if (pn) {
			window.branLazyPortals[pn] = this;
		}
		this.reload();
	},

	data(): DynaCompData {
		return {
			current: null,
		};
	},

	methods: {
		reload() {
			const ef = this.loaderFunc;
			const afterLoaded = this.afterLoaded;
			const self = this;
			const rootChangeCurrent = (this.$root as any).changeCurrent;
			const core = new Core(form, rootChangeCurrent, this.changeCurrent);

			core.fetchEvent(ef, {})
				.then((r) => {
					self.current = core.componentByTemplate(r.schema, afterLoaded);
				});
		},
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},
});


const vm = new Vue({
	...{
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
			const core = new Core(form, this.changeCurrent, this.changeCurrent);
			this.current = core.componentByTemplate(app.innerHTML);
		},

		data(): DynaCompData {
			return {
				current: null,
			};
		},
	},
	...vueOptions,
});

vm.$mount('#app');
