import Vuetify from 'vuetify';
import './main.styl';

declare var window: any;

(window.__branVueComponentRegisters = window.__branVueComponentRegisters || []).push((Vue: any, core: any): any => {
	Vue.use(Vuetify);

	core.extendSetupFuncs({
		vuetifyVSelect: (params: any) => {
			const comp = params.vnode.componentInstance;
			const fieldName = params.fieldName;
			const form = core.form;
			const values = form.getAll(fieldName);

			comp.lazyValue = values;

			comp.$on('change', (vals: string[]) => {
				form.delete(fieldName);
				vals.forEach((v) => {
					form.append(fieldName, v);
				});
			});
		},
	});
});
