import Vue, {
	Component,
	ComponentOptions,
	CreateElement,
	VNode,
	VNodeData,
	VueConstructor,
} from 'vue';
import {Core, slotTemplates} from './Helpers';


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
			(this as any).$plaid().fieldValue(this.$props.fieldName, val);
		},

		render(h: CreateElement): VNode {
			const {
				fieldName,
			} = this.$props;

			const ch = this.$listeners["change"];

			let chs: any[] = [];
			if (ch) {
				if (Array.isArray(ch)) {
					chs = ch;
				} else {
					chs = [ch];
				}
			}
			const data: VNodeData = {
				props: {
					...this.$attrs,
				},

				on: {
					...this.$listeners,
					...{
						change: [(val: any) => {
							(this as any).$plaid().fieldValue(fieldName, val);
						}, ...chs],
					},
				},

				scopedSlots: this.$scopedSlots,
			};
			return (
				<comp {...data}>
					{slotTemplates(h, this.$slots)}
				</comp>
			);
		},
	});
};
