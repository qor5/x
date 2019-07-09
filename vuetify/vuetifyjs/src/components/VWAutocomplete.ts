import { VAutocomplete } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData } from 'vue';
import Core from './Core';

export default Vue.extend({
	name: 'vw-autocomplete',
	mixins: [Core],
	data: () => ({
		isLoading: false,
		_items: [],
		model: null,
		searchKeyword: '',
	}),
	watch: {
		searchKeyword(val: string) {
			if (val === null) {
				return;
			}
			// console.log('in search', val);
			if (this._items && this._items.length > 0) { return; }

			this.isLoading = true;

			// Lazily load input items
			fetch('https://api.coinmarketcap.com/v2/listings/')
				.then((res) => res.json())
				.then((res) => {
					// console.log('res.data', res.data);
					this._items = res.data;
				})
				.catch((err) => {
					// console.log(err);
				})
				.finally(() => (this.isLoading = false));
		},
	},
	render(h: CreateElement): VNode {
		console.log('this.$attrs', this.$attrs);
		const self = this;
		const {
			fieldName,
		} = self.$attrs;
		const form = self.core.form;
		// const values = form.getAll(fieldName);
		const data: VNodeData = {
			props: {
				...{
					solo: true,
					chips: true,
					deletableChips: true,
				},
				...self.$attrs,
				...{
					items: self._items,
					loading: self.isLoading,
				},
			},

			on: {
				'change': (vals: any) => {
					form.delete(fieldName);
					if (typeof vals === 'string') {
						vals = [vals];
					}
					vals.forEach((v: string) => {
						form.append(fieldName, v);
					});
				},
				'update:searchInput': (val: string) => {
					self.searchKeyword = val;
				},
				'click': (e: any) => {
					self.searchKeyword = '';
				},
			},
		};
		return h(VAutocomplete, data);
	},
});

