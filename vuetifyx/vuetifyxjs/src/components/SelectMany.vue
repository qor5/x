<template>
	<div>
		<label class="v-label theme--light" v-html="label"></label>
		<v-card v-if='internalSelectedItems.length > 0'>
			<v-list>
				<vx-draggable
					v-model='internalSelectedItems'
					animation='300'
					handle=".handle"
					@change="changeOrder">
					<div v-for='(item, index) in internalSelectedItems' :key='item.id'>
						<v-list-item>
							<v-list-item-avatar tile>
								<v-img :src="item[itemImage]"></v-img>
							</v-list-item-avatar>
							<v-list-item-content>
								<v-list-item-title>{{ item[itemText] }}
								</v-list-item-title>
							</v-list-item-content>
							<v-list-item-icon>
								<v-btn :icon="true" @click="removeItem(item[itemValue])">
									<v-icon>delete</v-icon>
								</v-btn>
								<v-icon class="handle">reorder</v-icon>
							</v-list-item-icon>
						</v-list-item>
						<v-divider
							v-if='index < internalSelectedItems.length - 1'
							:key='index'></v-divider>
					</div>
				</vx-draggable>
			</v-list>
		</v-card>

		<v-autocomplete
			:item-value="itemValue"
			:item-text="itemText"
			:items="internalItems"
			:label="addItemLabel"
			:value="autocompleteValue"
			auto-select-first
			@change="addItem"
			@focus="focus"
			:loading="isLoading"
			:no-filter="noFilter"
			v-debounce:update:search-input="500"
			@update:search-input:debounced="search"
		>
			<template v-slot:item="data">
				<template>
					<v-list-item-avatar tile>
						<img :src="data.item[itemImage]">
					</v-list-item-avatar>
					<v-list-item-content>
						<v-list-item-title
							v-html="data.item[itemText]"></v-list-item-title>
					</v-list-item-content>
				</template>
			</template>
		</v-autocomplete>
	</div>
</template>

<script>
export default {
	name: 'vx-selectmany',
	props: {
		items: {
			type: Array,
			default: () => [],
		},
		selectedItems: {
			type: Array,
			default: () => [],
		},
		searchItemsFunc: {
			type: Function,
		},
		itemValue: {
			type: String,
			default: 'id',
		},
		itemText: {
			type: String,
			default: 'text',
		},
		itemImage: {
			type: String,
			default: 'image',
		},
		label: {
			type: String,
			default: '',
		},
		addItemLabel: {
			type: String,
			default: '',
		},
	},
	data() {
		return {
			internalSelectedItems: [],
			internalItems: [],
			autocompleteValue: [],
			searchKeyword: '',
			isLoading: false,
			noFilter: false,
		};
	},
	mounted() {
		this.internalSelectedItems = this.selectedItems;
		this.internalItems = this.items;
		if (this.searchItemsFunc) {
			this.noFilter = true;
		}
		this.$nextTick(() => {
			this.setValue();
		});
	},
	watch: {
		searchKeyword(val) {
			// console.log('searchKeyword', val);
			if (val === null) {
				return;
			}
			this.isLoading = true;
			this.searchItemsFunc(val).then(r => {
				this.internalItems = r.data || [];
			}).finally(() => {
				this.isLoading = false;
			});
		},
	},
	methods: {
		addItem(event) {
			this.autocompleteValue = [];
			if (this.internalSelectedItems.find(
				element => element[this.itemValue] == event)) {
				return;
			}
			this.internalSelectedItems.push(
				this.internalItems.find(element => element[this.itemValue] == event));
			this.setValue();
		},
		changeOrder(event) {
			this.setValue();
		},
		removeItem(id) {
			this.internalSelectedItems = this.internalSelectedItems.filter(
				element => element[this.itemValue] != id);
			this.setValue();
		},
		search(val) {
			if (!this.searchItemsFunc) {
				return;
			}
			this.searchKeyword = val;
		},
		focus() {
			this.search('');
		},
		setValue() {
			this.$emit('input', this.internalSelectedItems.map((i) => {
				return i[this.itemValue];
			}));
		},
	},
};
</script>

<style scoped>
.handle {
	margin-left: 8px;
	cursor: move;
}
</style>
