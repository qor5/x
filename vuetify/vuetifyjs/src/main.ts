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
} from 'vuetify/lib';

import 'vuetify/dist/vuetify.min.css';
import './main.css';

// import './main.styl';

import VWAutocomplete from './components/VWAutocomplete';
import { WithField } from './components/WithField';
import { Core, SelectedItems } from './components/Helpers';

const vuetify = new Vuetify({
	icons: {
		iconfont: 'md', // 'mdi' || 'mdiSvg' || 'md' || 'fa' || 'fa4'
	},
});

declare var window: any;

const inputValueFunc = (formValue: string, props: Record<string, any>): any => {
	return {
		inputValue: formValue,
	};
};

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', VWAutocomplete);
		Vue.component('vw-text-field', WithField(VTextField));
		Vue.component('vw-textarea', WithField(VTextarea));
		Vue.component('vw-checkbox', WithField(VCheckbox, inputValueFunc));
		Vue.component('vw-switch', WithField(VSwitch, inputValueFunc));
		Vue.component('vw-radio-group', WithField(VRadioGroup));
		Vue.component('vw-slider', WithField(VSlider));
		Vue.component('vw-select', WithField(VSelect, undefined, [Core, SelectedItems]));

		vueOptions.vuetify = vuetify;
	});
