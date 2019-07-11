import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import {
	VTextField,
	VTextarea,
	VCheckbox,
	VRadioGroup,
	VSwitch,
	VSlider,
} from 'vuetify/lib';

import './main.styl';

import VWAutocomplete from './components/VWAutocomplete';
import { WithField } from './components/WithField';

declare var window: any;

const inputValueFunc = (formValue: string): any => {
	return {
		inputValue: formValue,
	};
};

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', VWAutocomplete);
		Vue.component('vw-text-field', WithField(VTextField));
		Vue.component('vw-textarea', WithField(VTextarea));
		Vue.component('vw-checkbox', WithField(VCheckbox, inputValueFunc));
		Vue.component('vw-switch', WithField(VSwitch, inputValueFunc));
		Vue.component('vw-radio-group', WithField(VRadioGroup));
		Vue.component('vw-slider', WithField(VSlider));
	});
