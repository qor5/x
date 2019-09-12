import { VueConstructor } from 'vue';
import './css/styles.scss';


declare var window: any;

(window.__branVueComponentRegisters =
	window.__branVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		return;
	});
