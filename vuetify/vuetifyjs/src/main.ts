import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import {
	VTextField,
	VTextarea,
	VCheckbox,
	VRadioGroup,
	VSwitch,
	VSlider,
	VSelect,
	VChipGroup,
	VItemGroup,
	VFileInput,
	VCombobox,
} from 'vuetify/lib';

import 'vuetify/dist/vuetify.min.css';
import './main.css';

// import './main.styl';

import Autocomplete from './components/Autocomplete';
import { WithField } from './components/WithField';
import { Core, SelectedItems } from './components/Helpers';
import { Filter } from './components/Filter';
import Pagination from './components/Pagination';

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
		Vue.component('vw-text-field', WithField(VTextField));
		Vue.component('vw-textarea', WithField(VTextarea));
		Vue.component('vw-checkbox', WithField(VCheckbox, 'input-value'));
		Vue.component('vw-combobox', WithField(VCombobox));
		Vue.component('vw-switch', WithField(VSwitch, 'input-value'));
		Vue.component('vw-radio-group', WithField(VRadioGroup));
		Vue.component('vw-slider', WithField(VSlider));
		Vue.component('vw-file-input', WithField(VFileInput));
		Vue.component('vw-chip-group', WithField(VChipGroup));
		Vue.component('vw-item-group', WithField(VItemGroup));
		Vue.component('vw-select', WithField(VSelect, undefined, [Core, SelectedItems]));
		Vue.component('vw-filter', Filter);
		Vue.component('vw-pagination', Pagination);

		vueOptions.vuetify = vuetify;
	});
