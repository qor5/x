import { VueConstructor } from 'vue';
import './css/styles.scss';


declare var window: any;

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push((Vue: VueConstructor, vueOptions: any): any => {
		return;
	});
