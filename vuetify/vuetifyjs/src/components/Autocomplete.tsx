import { VAutocomplete } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData, Component } from 'vue';
import { Core, SelectedItems, slotTemplates } from './Helpers';

export default Vue.extend({
	mixins: [Core, SelectedItems],
	props: {
		itemsEventFuncId: Object,
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

	mounted() {
		this.core.setFormValue(this.$props.fieldName, this.$attrs.value);
	},

	watch: {
		searchKeyword(val: string) {
			// console.log('searchKeyword', val);
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
			fieldName,
			itemsEventFuncId,
			multiple,
		} = self.$props;

		let onSearchInput = {};
		let hideSelected = false;
		if (itemsEventFuncId) {
			onSearchInput = {
				'update:search-input': (val: string) => {
					self.searchKeyword = val;
				},
			};
			hideSelected = true;
		}

		const data: VNodeData = {
			props: {
				...{
					// solo: true,
					multiple,
					chips: true,
					deletableChips: multiple,
					clearable: true,
					hideSelected,
				},
				...self.$attrs,
				...{
					items: self._items,
					loading: self.isLoading,
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
		const comp: Component = VAutocomplete;
		return (
			<comp {...data}>
				{slotTemplates(h, this.$slots)}
			</comp>
		);
	},
});

