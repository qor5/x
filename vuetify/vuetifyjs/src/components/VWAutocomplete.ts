import { VAutocomplete } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData } from 'vue';
import Core from './Core';

export default Vue.extend({
	name: 'vw-autocomplete',
	mixins: [Core],
	props: {
		itemsEventFuncId: Object,
		selectedItems: {
			type: Array,
			default: () => [],
		},
		items: {
			type: Array,
			default: () => [],
		},
	},

	data() {
		return {
			isLoading: false,
			_items: [],
			searchKeyword: '',
		};
	},

	created() {
		this.fetchEvent = this.core.debounce((val: string) => {
			this.isLoading = true;
			this.core.fetchEvent(this.itemsEventFuncId, { value: val })
				.then((r: any) => {
					// console.log('res.data', res.data);
					this._items = r.data;
				})
				// .catch((err: any) => {
				// 	console.log('debounceFetchEvent', err);
				// })
				.finally(() => (this.isLoading = false));
		}, 500);

		this._items = this.$props.items;
	},


	watch: {
		searchKeyword(val: string) {
			// console.log('this.itemsEventFuncId', this.itemsEventFuncId);
			if (val === null) {
				return;
			}
			// console.log('in search', val);
			// if (this._items && this._items.length > 0) { return; }

			this.fetchEvent(val);
		},
	},

	render(h: CreateElement): VNode {
		// console.log('this.$attrs', this.$attrs);
		const self = this;
		const {
			multiple,
		} = self.$attrs;

		const {
			fieldName,
			selectedItems,
			itemsEventFuncId,
		} = self.$props;

		// console.log('itemsEventFuncId', itemsEventFuncId);
		const form = self.core.form;
		// const values = form.getAll(fieldName);
		// console.log('fieldName', fieldName);
		let onSearchInput = {};
		if (itemsEventFuncId) {
			onSearchInput = {
				'update:searchInput': (val: string) => {
					self.searchKeyword = val;
				},
			};
		}

		let value = selectedItems;
		if (!multiple) {
			value = selectedItems[0];
		}

		const data: VNodeData = {
			props: {
				...{
					solo: true,
					chips: multiple,
					deletableChips: multiple,
				},
				...self.$attrs,
				...{
					items: self._items,
					loading: self.isLoading,
					value,
				},
			},

			on: {
				...{
					change: (vals: any) => {
						form.delete(fieldName);
						if (typeof vals === 'string') {
							vals = [vals];
						}
						vals.forEach((v: string) => {
							form.append(fieldName, v);
						});
					},
					click: (e: any) => {
						self.searchKeyword = '';
					},
				},
				...onSearchInput,
			},
		};
		return h(VAutocomplete, data);
	},
});

