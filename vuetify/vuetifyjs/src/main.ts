import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import './main.styl';

import VWAutocomplete from './components/VWAutocomplete';

declare var window: any;

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', VWAutocomplete);
	});
