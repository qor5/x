import { VTextField } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData } from 'vue';
import Core from './Core';

export default Vue.extend({
	name: 'vw-text-field',
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
		return h(VTextField, data);
	},
});

