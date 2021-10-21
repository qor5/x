import { VueConstructor } from 'vue';
import {
	VCheckbox,
	VSelect,
} from 'vuetify/lib';

import Autocomplete from './components/Autocomplete';
import { WithField } from './components/WithField';
import { Core, SelectedItems } from './components/Helpers';
import { Filter } from './components/Filter';
// @ts-ignore
import DatetimePicker from './components/DateTimePicker.vue';

declare var window: any;

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		Vue.component('vx-autocomplete', Autocomplete);
		Vue.component('vx-checkbox', WithField(VCheckbox, 'input-value'));
		Vue.component('vx-select', WithField(VSelect, undefined, [Core, SelectedItems]));
		Vue.component('vx-filter', Filter);
		Vue.component('vx-datetimepicker', DatetimePicker);
	});
