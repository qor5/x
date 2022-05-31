import {VAutocomplete} from 'vuetify/lib';

import Vue, {Component, CreateElement, VNode, VNodeData} from 'vue';
import {Core, SelectedItems, slotTemplates} from './Helpers';

export default Vue.extend({
	mixins: [Core, SelectedItems],
	props: {
		itemsEventFuncId: Object,
		items: {
			type: Array,
			default: () => ([]),
		},
	},

	data() {
		return {
			isLoading: false,
			cached_items: [],
			searchKeyword: '',
		};
	},

	created() {
		this.cached_items = this.$props.items;
	},

	mounted() {
		(this as any).$plaid().fieldValue(this.$props.fieldName, this.$attrs.value);
	},

	watch: {
		searchKeyword(val: string) {
			// console.log('searchKeyword', val);
			// console.log('this.itemsEventFuncId', this.itemsEventFuncId);
			if (val === null) {
				this.searchKeyword = '';
				return;
			}
			// console.log('in search', val);
			// if (this._items && this._items.length > 0) { return; }
			this.isLoading = true;

			(this as any).$plaid().eventFuncID(this.itemsEventFuncId).query("keyword", val).go().then((r: any) => {
				const v = [].concat((this as any).selectedItems || [], r.data || []);
				// console.log("after concat", v);
				this.cached_items = v
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

		const {
			fieldName,
			itemsEventFuncId,
			multiple,
		} = this.$props;

		let onSearchInput = {};
		let hideSelected = false;
		if (itemsEventFuncId) {
			onSearchInput = {
				'update:search-input': (val: string) => {
					this.searchKeyword = val;
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
				...this.$attrs,
				...{
					items: this.cached_items,
					loading: this.isLoading,
				},
			},

			on: {
				...{
					change: (vals: any) => {
						this.$emit("change", vals);
					},
					focus: (e: any) => {
						this.searchKeyword = '';
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

