import {VueConstructor} from 'vue';
import {VSelect,} from 'vuetify/lib';

import Autocomplete from './components/Autocomplete';
import {WithField} from './components/WithField';
import {Core, SelectedItems} from './components/Helpers';
import {Filter} from './components/Filter';
// @ts-ignore
import DatetimePicker from './components/DateTimePicker.vue';
import draggable from 'vuedraggable'

declare var window: any;

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
	Vue.component('vx-autocomplete', Autocomplete);
	Vue.component('vx-select', WithField(VSelect, undefined, [Core, SelectedItems]));
	Vue.component('vx-filter', Filter);
	Vue.component('vx-datetimepicker', DatetimePicker);
	Vue.component('vx-draggable', draggable);
});
