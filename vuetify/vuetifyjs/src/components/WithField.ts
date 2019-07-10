import Vue, { CreateElement, VNode, VNodeData, Component, VueConstructor } from 'vue';
import Core from './Core';

export const WithField = (comp: Component): VueConstructor => {
	return Vue.extend({
		mixins: [Core],

		render(h: CreateElement): VNode {
			const self = this;
			const {
				fieldName,
			} = self.$props;

			const formValue = self.core.getFormValue(fieldName);

			const data: VNodeData = {
				props: {
					...{
						value: formValue,
					},
					...self.$attrs,
				},

				on: {
					change: (val: string) => {
						self.core.setFormValue(fieldName, val);
					},
				},
			};
			return h(comp, data);
		},
	});
};

