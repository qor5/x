import Vue, { VNode, VNodeData, CreateElement } from 'vue';

import {
	VPagination,
} from 'vuetify/lib';


export default Vue.extend({
	inject: ['core'],
	components: {
		vpagination: VPagination,
	},

	render(h: CreateElement): VNode {
		const self = this;
		const data: VNodeData = {
			props: {
				...self.$attrs,
			},

			on: {
				...{
					input: (val: any) => {
						self.core.loadPage({ page: `${val}` });
					},
				},
				...this.$listeners,
			},
		};
		return <vpagination {...data}></vpagination>;
	},

});
