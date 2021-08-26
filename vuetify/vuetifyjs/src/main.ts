import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import {
	VCheckbox,
	VSelect,
} from 'vuetify/lib';

import 'vuetify/dist/vuetify.min.css';
import './main.css';

// import './main.styl';

import Autocomplete from './components/Autocomplete';
import { WithField } from './components/WithField';
import { Core, SelectedItems } from './components/Helpers';
import { Filter } from './components/Filter';

const vuetify = new Vuetify({
	icons: {
		iconfont: 'md', // 'mdi' || 'mdiSvg' || 'md' || 'fa' || 'fa4'
	},
});

declare var window: any;

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', Autocomplete);
		Vue.component('vw-checkbox', WithField(VCheckbox, 'input-value'));
		Vue.component('vw-select', WithField(VSelect, undefined, [Core, SelectedItems]));
		Vue.component('vw-filter', Filter);

		vueOptions.vuetify = vuetify;
	});
