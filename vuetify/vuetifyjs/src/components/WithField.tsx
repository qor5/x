import Vue, {
	CreateElement,
	VNode,
	VNodeData,
	Component,
	VueConstructor,
	ComponentOptions,
} from 'vue';
import { Core, slotTemplates } from './Helpers';


export const WithField = (
	comp: Component,
	valueField?: string,
	mixins?: Array<ComponentOptions<Vue> | typeof Vue>,
): VueConstructor => {
	const m = mixins || [Core];
	return Vue.extend({
		mixins: m,

		mounted() {
			const val = this.$attrs[valueField || 'value'];
			this.core.setFormValue(this.$props.fieldName, val);
		},

		render(h: CreateElement): VNode {
			const self = this;
			const {
				fieldName,
			} = self.$props;

			const data: VNodeData = {
				props: {
					...self.$attrs,
				},

				on: {
					...{
						change: (val: any) => {
							self.core.setFormValue(fieldName, val);
						},
					},
					...this.$listeners,
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
