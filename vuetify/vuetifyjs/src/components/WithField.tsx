import Vue, { CreateElement, VNode, VNodeData, Component, VueConstructor } from 'vue';
import { Core, slotTemplates } from './Helpers';


export const WithField = (comp: Component, valuePropsFunc?: (formValue: string) => any): VueConstructor => {
	return Vue.extend({
		mixins: [Core],

		render(h: CreateElement): VNode {
			const self = this;
			const {
				fieldName,
			} = self.$props;


			const formValue = self.core.getFormValue(fieldName);
			let valueProps = {
				value: formValue,
			};
			if (valuePropsFunc) {
				valueProps = valuePropsFunc(formValue);
			}

			const data: VNodeData = {
				props: {
					...valueProps,
					...self.$attrs,
				},

				on: {
					change: (val: any) => {
						self.core.setFormValue(fieldName, val);
					},
				},

				scopedSlots: this.$scopedSlots,
			};
			return (
				<comp {...data}>
					{slotTemplates(h, self.$slots)}
				</comp>
			);
		},
	});
};
