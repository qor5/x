import { VueConstructor } from 'vue';
import Vuetify from 'vuetify';
import './main.styl';

import VWAutocomplete from './components/VWAutocomplete';
import VWTextField from './components/VWTextField';

declare var window: any;

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor): any => {
		Vue.use(Vuetify);
		Vue.component('vw-autocomplete', VWAutocomplete);
		Vue.component('vw-text-field', VWTextField);
	});
