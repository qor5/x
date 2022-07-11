import Vue, { CreateElement, VNode, VNodeData } from 'vue';
import { VAutocomplete, VPagination } from 'vuetify/lib';
import { Core, SelectedItems, slotTemplates } from './Helpers';

export default Vue.extend({
	mixins: [Core, SelectedItems],
	components: {
		vpagination: VPagination,
		vautocomplete: VAutocomplete,
	},

	props: {
		remoteUrl: String,
		remoteRes:  String,
		cacheItems: Boolean,
		hideSelected: Boolean,
		isPaging: Boolean,
		items: {
			type: Array,
			default: () => ([]),
		},
	},

	data() {
		return {
			listItems: [],
			cachedSelectedItems: [],
			value: [],
			isLoading: false,
			searchKeyword: '',
			remote: {
				total: 0,
				current: 0,
				disabled: false,
			},
		};
	},

	methods: {
		loadRemoteItems() {
			if (!this.remoteUrl || !this.remoteRes) {
				return;
			}

			this.isLoading = true;
			(this as any).$plaid().url(this.remoteUrl).eventFunc("autocomplete-remote-res-event").query("keyword", this.searchKeyword).query("name", this.remoteRes).query("current", this.remote.current).go().then((r: any) => {
				this.remote.current = r.data.current;
				this.remote.total = r.data.total;
				if (this.isPaging) {
					this.listItems = [].concat(this.cachedSelectedItems || [], r.data.items || []);
				}else{
					if (this.remote.current >= this.remote.total) {
						this.remote.disabled = true;
					}else{
						this.remote.disabled = false;
					}
					this.listItems = [].concat(this.listItems || [], r.data.items || []);
				}
			}).finally(() => {
				this.isLoading = false;
			});
		},
		endIntersect(entrie: any, observer: any, isIntersecting: any) {
			if (isIntersecting && !this.remote.disabled) {
				this.loadRemoteItems();
			}
		},
	},

	created() {
		this.listItems =  this.$props.items || this.$props.selectedItems || [];
		this.cachedSelectedItems = this.$props.selectedItems || [];
		this.value = (this.$attrs.value) as any || [];
	},

	mounted() {
		(this as any).$plaid().fieldValue(this.$props.fieldName, this.$attrs.value);
	},

	watch: {
		searchKeyword(val: string) {
			if (val === null) {
				this.searchKeyword = '';
				return;
			}

			if (this.isPaging) {
				this.remote.current = 1;
			}else{
				this.listItems  = this.cachedSelectedItems
				this.remote.current = 0;
			}

			this.loadRemoteItems();
		},
	},

	render(h: CreateElement): VNode {
		const {
			remoteRes,
			multiple,
		} = this.$props;

		let {
			hideSelected,
			cacheItems,
		} = this.$props

		if (remoteRes) {
			hideSelected = true;
			cacheItems = false;
			const loadmoreNode: VNode[] = [];

			if (this.isPaging){
				const loadmoreNodeData: VNodeData = {
					props: {
						circle: true,
						length: this.remote.total,
						value: this.remote.current,
						totalVisible: 3,
					},
					on: {
						"input": (v: number) => {
							this.remote.current = v;
							this.loadRemoteItems();
						},
					},
				}

				loadmoreNode.push(
					<div class="text-center">
						<vpagination {...loadmoreNodeData}></vpagination>
					</div>
				);
			}else{
				const loadmoreNodeData: VNodeData = {
					props: {
						class: "ma-2",
						color: "primary",
						disabled: this.remote.disabled,
						loading: this.isLoading,
					},
					on: {
						"click": () => {
							this.loadRemoteItems();
						},
					},
					directives: [{
						name: "intersect",
						value: this.endIntersect,
					}],
				}

				loadmoreNode.push(
					<div class="text-center">
						<v-btn {...loadmoreNodeData}>Load more</v-btn> <v-divider vertical></v-divider> <span>{this.remote.current}/{this.remote.total}</span>
					</div>
				);
			}
			this.$slots["append-item"] = loadmoreNode;
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
					cacheItems,
				},
				...this.$attrs,
				...{
					items: this.listItems,
					value: this.value,
					loading: this.isLoading,
				},
			},

			on: {
				...{
					change: (vals: any) => {
						const items: any[] = [];
						this.listItems.forEach((item: any) => {
							if (vals.includes(item.value)) {
								items.push(item);
							}
						})
						this.cachedSelectedItems = (items) as [];
						this.value = vals;
						this.$emit("change", vals);
					},
					focus: (e: any) => {
						this.searchKeyword = '';
					},
					'update:search-input': (val: string) => {
						this.searchKeyword = val;
					},
				},
			},
		};
		return (
			<vautocomplete {...data}>
				{slotTemplates(h, this.$slots)}
			</vautocomplete>
		);
	},
});

