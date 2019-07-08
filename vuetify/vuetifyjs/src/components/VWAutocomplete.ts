import { VAutocomplete } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData } from 'vue';
import Core from './Core';

export default Vue.extend({
	name: 'vw-autocomplete',
	mixins: [Core],
	props: {
		items: {
			type: Array,
			default: () => [],
		},
		multiple: Boolean,
	},
	render(h: CreateElement): VNode {
		const props = this.$props;
		const fieldName = props.fieldName;
		const form = this.core.form;
		const values = form.getAll(fieldName);

		const data = {
			props: {
				solo: true,
				items: this.$props.items,
				multiple: true,
				chips: true,
				deletableChips: true,
				value: values,
			},
			on: {
				change: (vals: any) => {
					form.delete(fieldName);
					if (typeof vals === 'string') {
						vals = [vals];
					}
					vals.forEach((v: string) => {
						form.append(fieldName, v);
					});
				},
			},
		};
		return h(VAutocomplete, data);
	},
});

