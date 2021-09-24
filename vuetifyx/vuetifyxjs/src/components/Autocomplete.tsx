import { VAutocomplete } from 'vuetify/lib';

import Vue, { CreateElement, VNode, VNodeData, Component, PropType } from 'vue';
import { Core, SelectedItems, slotTemplates } from './Helpers';

export default Vue.extend({
	mixins: [Core, SelectedItems],
	props: {
		itemsEventFuncId: Object,
		items: {
			type: Array,
			default: () => ([]),
		} as any,
	},

	data() {
		return {
			isLoading: false,
			_items: [],
			searchKeyword: '',
		};
	},

	created() {
		this._items = this.$props.items;
	},

	mounted() {
		(this as any).$plaid().fieldValue(this.$props.fieldName, this.$attrs.value);
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
			this.isLoading = true;

			(this as any).$plaid().
				debounce(500).
				eventFuncID(this.itemsEventFuncId).
				event({value: val}).
				go().
				then((r: any) => {
					this._items = r.data;
				})
				// .catch((err: any) => {
				// 	console.log('debounceFetchEvent', err);
				// })
				.finally(() => {
					this.isLoading = false
				});
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
						(self as any).$plaid().fieldValue(fieldName, vals);
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

