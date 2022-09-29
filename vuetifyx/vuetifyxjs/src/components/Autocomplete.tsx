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
		eventName:  String,
		isPaging: Boolean,
		hasIcon: Boolean,
		cacheItems: Boolean,
		hideSelected: Boolean,
		hideDetails: Boolean,
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
				pages: 0,
				page: 0,
				disabled: false,
			},
		};
	},

	methods: {
		loadRemoteItems() {
			if (!this.remoteUrl || !this.eventName) {
				return;
			}

			this.isLoading = true;
			(this as any).$plaid().url(this.remoteUrl).eventFunc(this.eventName).query("keyword", this.searchKeyword).query("page", this.remote.page).go().then((r: any) => {
				this.remote.current = r.data.current;
				this.remote.total = r.data.total;
				this.remote.pages = r.data.pages;
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
				this.remote.page += 1;
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
			if (!this.remoteUrl || !this.eventName) {
				return;
			}

			if (val === null) {
				this.searchKeyword = '';
				return;
			}

			this.remote.page = 1;
			if (!this.isPaging) {
				this.listItems  = this.cachedSelectedItems
			}

			this.loadRemoteItems();
		},
	},

	render(h: CreateElement): VNode {
		const {
			remoteUrl,
			multiple,
			hideDetails,
		} = this.$props;

		let {
			hideSelected,
			cacheItems,
		} = this.$props

		const slots: VNode[] = slotTemplates(h, this.$slots);
		if (remoteUrl) {
			hideSelected = true;
			cacheItems = false;
			if (this.isPaging){
				const loadmoreNodeData: VNodeData = {
					props: {
						circle: true,
						length: this.remote.pages,
						value: this.remote.page,
						totalVisible: 5,
					},
					on: {
						"input": (v: number) => {
							this.remote.page = v;
							this.loadRemoteItems();
						},
					},
				}

				slots.push(
					<template slot="append-item">
						<div class="text-center">
							<vpagination {...loadmoreNodeData}></vpagination>
						</div>
					</template>
				)
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
							this.remote.page += 1;
							this.loadRemoteItems();
						},
					},
					directives: [{
						name: "intersect",
						value: this.endIntersect,
					}],
				}

				slots.push(
					<template slot="append-item">
						<div class="text-center">
							<v-btn {...loadmoreNodeData}>Load more</v-btn> <v-divider vertical></v-divider> <span>{this.remote.current}/{this.remote.total}</span>
						</div>
					</template>
				)
			}

			if (this.hasIcon){
				this.$scopedSlots["item"] = (props: any)  => {
					const nodes: VNode[] = [];
					nodes.push(
						<v-list-item-avatar tile>
							<img src={props.item.icon}/>
						</v-list-item-avatar>
					)
					nodes.push(
						<v-list-item-content>
							<v-list-item-title v-html={props.item.text}>{props.item.text}</v-list-item-title>
						</v-list-item-content>
					)
					return nodes;
				}

				this.$scopedSlots["selection"] = (props: any)  => {
					const nodes: VNode[] = [];
					const nodeData: VNodeData = {
						props: {
							...props.attrs,
							close: true,
						},
						on: {
							"click:close": () => {
								this.value.splice(this.value.indexOf(props.item.value as never), 1)
								this.$emit("change", this.value);
							},
						},

					}
					nodes.push(
					<v-chip {...nodeData}>
						<v-avatar left>
							<v-img src={props.item.icon}></v-img>
						</v-avatar>
						{props.item.text}
					</v-chip>
					)
					return nodes;
				}
			}
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
					hideDetails,
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
			scopedSlots: {
				...this.$scopedSlots,
			}
		};
		return (
			<vautocomplete {...data}>
				{slots}
			</vautocomplete>
		);
	},
});

