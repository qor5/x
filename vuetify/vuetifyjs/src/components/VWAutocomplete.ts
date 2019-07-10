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
		const self = this;
		const core = self.core;
		self.fetchEvent = core.debounce((val: string) => {
			self.isLoading = true;
			core.fetchEvent(self.itemsEventFuncId, { value: val })
				.then((r: any) => {
					// console.log('r.data', r.data);
					self._items = r.data;
				})
				// .catch((err: any) => {
				// 	console.log('debounceFetchEvent', err);
				// })
				.finally(() => (self.isLoading = false));
		}, 500);

		self._items = self.$props.items;
	},


	watch: {
		searchKeyword(val: string) {
			console.log('searchKeyword', val);
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
		// console.log('render', this);
		const self = this;
		const {
			multiple,
		} = self.$attrs;

		const {
			fieldName,
			selectedItems,
			itemsEventFuncId,
		} = self.$props;

		let onSearchInput = {};
		let hideSelected = false;
		if (itemsEventFuncId) {
			onSearchInput = {
				'update:searchInput': (val: string) => {
					self.searchKeyword = val;
				},
			};
			hideSelected = true;
		}

		let value = selectedItems;
		if (!multiple) {
			value = selectedItems[0];
		}

		const data: VNodeData = {
			props: {
				...{
					// solo: true,
					chips: true,
					deletableChips: multiple,
					clearable: true,
					hideSelected,
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
						self.core.setFormValue(fieldName, vals);
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

