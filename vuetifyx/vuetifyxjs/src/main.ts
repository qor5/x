import { VueConstructor } from 'vue';
import {
	VCheckbox,
	VSelect,
} from 'vuetify/lib';

import Autocomplete from './components/Autocomplete';
import { WithField } from './components/WithField';
import { Core, SelectedItems } from './components/Helpers';
import { Filter } from './components/Filter';

declare var window: any;

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		Vue.component('vx-autocomplete', Autocomplete);
		Vue.component('vx-checkbox', WithField(VCheckbox, 'input-value'));
		Vue.component('vx-select', WithField(VSelect, undefined, [Core, SelectedItems]));
		Vue.component('vx-filter', Filter);

	});
