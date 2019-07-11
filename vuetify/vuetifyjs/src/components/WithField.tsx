import Vue, {
	CreateElement,
	VNode,
	VNodeData,
	Component,
	VueConstructor,
	ComponentOptions,
} from 'vue';
import { Core, slotTemplates, selectValue } from './Helpers';


export const WithField = (
	comp: Component,
	valuePropsFunc?: (formValue: string, props: Record<string, any>) => any,
	mixins?: Array<ComponentOptions<Vue> | typeof Vue>,
): VueConstructor => {
	const m = mixins || [Core];
	return Vue.extend({
		mixins: m,

		render(h: CreateElement): VNode {
			const self = this;
			const {
				fieldName,
			} = self.$props;


			let valueProps = {};
			if (valuePropsFunc) {
				const formValue = self.core.getFormValue(fieldName);
				valueProps = valuePropsFunc(formValue, this.$props);
			} else {
				valueProps = {
					value: selectValue(self.core, this.$props),
				};
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
