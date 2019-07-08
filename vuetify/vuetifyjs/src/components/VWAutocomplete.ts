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
			console.log('in search', val);
			if (this._items && this._items.length > 0) { return; }

			this.isLoading = true;

			// Lazily load input items
			fetch('https://api.coinmarketcap.com/v2/listings/')
				.then((res) => res.json())
				.then((res) => {
					console.log('res.data', res.data);
					this._items = res.data;
				})
				.catch((err) => {
					console.log(err);
				})
				.finally(() => (this.isLoading = false));
		},
	},
	render(h: CreateElement): VNode {
		const self = this;
		const props = self.$props;
		const fieldName = props.fieldName;
		const form = self.core.form;
		const values = form.getAll(fieldName);
		const data: VNodeData = {
			props: {
				solo: true,
				items: self._items,
				multiple: true,
				chips: true,
				deletableChips: true,
				// value: values,
				// searchInput: 'abc',
				// noFilter: true,
				loading: self.isLoading,
				itemText: 'name',
				itemValue: 'symbol',
				noDataText: 'Search your things',
				hideNoData: true,
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

