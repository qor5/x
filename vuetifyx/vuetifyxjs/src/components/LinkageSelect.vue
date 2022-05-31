<template>
	<div>
		<v-row v-if="row">
			<v-col v-for="(v, i) in items" :key="v.ID">
				<v-autocomplete
					:key="v.ID"
					:label="labels[i]"
					:items="levelItems(i)"
					item-text="Name"
					item-value="ID"
					v-model="selectedIDs[i]"
					@change="selectItem($event, i)"
					:clearable="chips ? false : true"
					:error-messages="errorMessages[i]"
					:chips="chips"
					:disabled="disabled"
					:hide-details="hideDetails"
				>
				</v-autocomplete>
			</v-col>
		</v-row>
		<v-autocomplete
			v-else
			v-for="(v, i) in items"
			:key="v.ID"
			:label="labels[i]"
			:items="levelItems(i)"
			item-text="Name"
			item-value="ID"
			v-model="selectedIDs[i]"
			@change="selectItem($event, i)"
			:clearable="chips ? false : true"
			:error-messages="errorMessages[i]"
			:chips="chips"
			:disabled="disabled"
			:hide-details="hideDetails"
		>
		</v-autocomplete>
	</div>
</template>

<script>
export default {
	name: 'vx-linkageselect',
	props: {
		value: {
			type: Array,
			default: () => [],
		},
		// [{ID, Name, ChildrenIDs}]
		items: {
			type: Array,
			default: () => [],
		},
		labels: {
			type: Array,
			default: () => [],
		},
		errorMessages: {
			type: Array,
			default: () => [],
		},
		disabled: {
			type: Boolean,
			default: false,
		},
		selectOutOfOrder: {
			type: Boolean,
			default: false,
		},
		chips: {
			type: Boolean,
			default: false,
		},
		row: {
			type: Boolean,
			default: false,
		},
		hideDetails: {
			type: Boolean,
			default: false,
		},
	},
	data() {
		return {
			selectedIDs: [],
		};
	},
	computed: {
		levelItems: function() {
			return function(level) {
				if (level === 0) {
					return this.items[level];
				}
				var items = [];
				if (this.selectedIDs[level - 1]) {
					var idM = {};
					for (const item of this.items[level - 1]) {
						if (item.ID === this.selectedIDs[level - 1]) {
							for (var id of item.ChildrenIDs) {
								idM[id] = true;
							}
							break;
						}
					}
					for (const item of this.items[level]) {
						if (idM[item.ID]) {
							items.push(item);
						}
					}
					return items;
				}

				if (this.selectOutOfOrder) {
					for (let i = level - 2; i >= 0; i--) {
						if (this.selectedIDs[i]) {
							items = this.findNextItems(this.selectedIDs[i], i);
							for (let j = i + 1; j < level; j++) {
								var newItems = [];
								for (const item of items) {
									newItems = newItems.concat(
										this.findNextItems(item.ID, j));
								}
								items = newItems;
							}

							return items;
						}
					}
					return this.items[level];
				}
				return [];
			};
		},
	},
	methods: {
		setValue() {
			this.$emit('input', this.selectedIDs);
		},
		validateAndResetSelectedIDs() {
			this.items.forEach((v, i) => {
				if (!this.selectedIDs[i]) {
					this.selectedIDs[i] = '';
				}
			});
			this.selectedIDs.forEach((v, i) => {
				if (!v) {
					this.selectedIDs[i] = '';
					return;
				}

				var exists = false;
				for (var item of this.items[i]) {
					if (item.ID === v) {
						exists = true;
						break;
					}
				}
				if (!exists) {
					this.selectedIDs[i] = '';
					return;
				}

				if (i === 0) {
					return;
				}
				var pID = this.selectedIDs[i - 1];
				if (!pID) {
					if (!this.selectOutOfOrder) {
						this.selectedIDs[i] = '';
					}
					return;
				} else {
					for (const item of this.items[i - 1]) {
						if (item.ID === pID) {
							for (var id of item.ChildrenIDs) {
								if (id === v) {
									return;
								}
							}
						}
					}
				}

				this.selectedIDs[i] = '';
				return;
			});
		},
		selectItem(v, level) {
			if (v) {
				for (var i = level + 1; i < this.selectedIDs.length; i++) {
					if (this.selectedIDs[i]) {
						var levelItems = this.levelItems[i];
						if (!levelItems || levelItems.length === 0) {
							this.selectedIDs[i] = '';
							continue;
						}
						var found = false;
						for (var item of levelItems) {
							if (item.ID === this.selectedIDs[i]) {
								found = true;
								break;
							}
						}
						if (!found) {
							this.selectedIDs[i] = '';
						}
					}
				}
			} else {
				this.selectedIDs[level] = '';
				if (!this.selectOutOfOrder) {
					for (let i = level + 1; i < this.selectedIDs.length; i++) {
						this.selectedIDs[i] = '';
					}
				}
			}
			this.setValue();
		},
		findNextItems(selectedID, level) {
			if (level + 1 >= this.items.length) {
				return [];
			}
			var childrenIDs = [];
			for (const item of this.items[level]) {
				if (item.ID === selectedID) {
					childrenIDs = item.ChildrenIDs;
					break;
				}
			}
			if (childrenIDs.length == 0) {
				return [];
			}
			var items = [];
			for (const item of this.items[level + 1]) {
				if (childrenIDs.includes(item.ID)) {
					items.push(item);
				}
			}
			return items;
		},
	},
	mounted() {
		this.items.forEach(v => {
			v.forEach(item => {
				if (!item.Name) {
					item.Name = item.ID;
				}
			});
		});
		this.selectedIDs = [...this.value];
		this.validateAndResetSelectedIDs();
		this.$nextTick(() => {
			this.setValue();
		});
	},
};
</script>
