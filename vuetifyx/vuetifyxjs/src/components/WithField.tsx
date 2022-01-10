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
			const self = this;
			const {
				fieldName,
			} = self.$props;

			const ch = self.$listeners["change"];

			let chs: Function[] = [];
			if (ch) {
				if (Array.isArray(ch)) {
					chs = ch;
				} else {
					chs = [ch];
				}
			}
			const data: VNodeData = {
				props: {
					...self.$attrs,
				},

				on: {
					...this.$listeners,
					...{
						change: [(val: any) => {
							(self as any).$plaid().fieldValue(fieldName, val);
						}, ...chs],
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
