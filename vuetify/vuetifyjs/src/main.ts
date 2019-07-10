import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import { VTextField, VTextarea } from 'vuetify/lib';

import './main.styl';

import VWAutocomplete from './components/VWAutocomplete';
import { WithField } from './components/WithField';

declare var window: any;

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', VWAutocomplete);
		Vue.component('vw-text-field', WithField(VTextField));
		Vue.component('vw-textarea', WithField(VTextarea));
	});
