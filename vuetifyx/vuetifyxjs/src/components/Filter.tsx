import {encodeFilterData, filterData} from './FilterData';
import Vue, {CreateElement, VNode} from 'vue';
import {
	VBtn,
	VCheckbox,
	VChip,
	VExpansionPanel,
	VExpansionPanelContent,
	VExpansionPanelHeader,
	VExpansionPanels,
	VIcon,
	VMenu,
	VRadio,
	VRadioGroup,
	VSelect,
	VSpacer,
	VTextField,
	VToolbar,
	VToolbarTitle,
} from 'vuetify/lib';

import TextDatePicker from './TextDatePicker';
import LinkageSelect from './LinkageSelect.vue';

export function localTimezoneAbbr() {
	const d = new Date().toString();
	return d
		.split('(')[1]
		.split(' ')
		.map((w) => {
			return w.charAt(0);
		})
		.join('');
}


export const DateItem = Vue.extend({
	components: {
		datePicker: TextDatePicker,
		radioGroup: VRadioGroup,
		radio: VRadio,
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
		translations: {
			type: Object,
			// default: () => {
			// 	return {
			// 		inTheLast: 'is in the last',
			// 		equals: 'is equal to',
			// 		between: 'is between',
			// 		isAfter: 'is after',
			// 		isAfterOrOn: 'is on or after',
			// 		isBefore: 'is before',
			// 		isBeforeOrOn: 'is before or on',
			// 		days: 'days',
			// 		months: 'months',
			// 		and: 'and',
			// 	};
			// },
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || 'inTheLast',
			valueIs: this.$props.value.valueIs,
			valueFrom: this.$props.value.valueFrom,
			valueTo: this.$props.value.valueTo,
			inTheLastUnit: this.$props.value.inTheLastUnit,
			inTheLastValue: this.$props.value.inTheLastValue,
			timezone: this.$props.value.timezone,
			datePickerVisible: false,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setModifier(e: string) {
			this.modifier = e;
			this.inputEmit();
			this.datePickerVisible = true;
			this.$forceUpdate();
		},

		setDate(e: any) {
			this.valueIs = e;
			this.inputEmit();
		},

		setDateFrom(e: any) {
			this.valueFrom = e;
			this.inputEmit();
		},

		setDateTo(e: any) {
			this.valueTo = e;
			this.inputEmit();
		},

		setInTheLastValue(e: any) {
			this.inTheLastValue = e;
			this.inputEmit();
		},

		setInTheLastUnit(e: any) {
			this.inTheLastUnit = e;
			this.inputEmit();
		},

		setTimezone(e: any) {
			this.timezone = e;
			this.$emit('input', this.$data);
		},

		getInput(modifier: string): VNode {
			const t = this.$props.translations;

			if (modifier === 'inTheLast') {
				return (
					<div>
						<vicon class='pr-5'>subdirectory_arrow_right</vicon>
						<vtextfield
							class='d-inline-block pr-5'
							style='width: 80px'
							type='number'
							on={{ change: this.setInTheLastValue }}
							value={(this.inTheLastValue || '').toString()}
							hideDetails={true}
						/>
						<vselect
							class='d-inline-block'
							style='width: 120px'
							on={{ change: this.setInTheLastUnit }}
							value={this.inTheLastUnit || 'days'}
							items={
								[
									{ text: t.days, value: 'days' },
									{ text: t.months, value: 'months' },
								]
							}
							hideDetails={true}
						>
						</vselect>
					</div>
				);
			}

			if (modifier === 'between') {
				return (
					<div>
						<vicon class='pr-5'>subdirectory_arrow_right</vicon>
						<datePicker
							value={this.valueFrom}
							on={{ input: this.setDateFrom }}
							key={modifier + 'from'}
							visible={this.datePickerVisible}
						/>{' '}
						<span class='px-3'>{t.and}</span>
						<datePicker
							value={this.valueTo}
							on={{ input: this.setDateTo }}
							key={modifier + 'to'}
						/>
					</div>
				);

			}

			return (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<datePicker
						value={this.valueIs}
						on={{ input: this.setDate }}
						key={modifier}
						visible={this.datePickerVisible}
					/>
				</div >
			);
		},
	},


	render(h: CreateElement): VNode {
		const t = this.$props.translations;
		return (
			<div >
				<div>
					<vselect
						class='d-inline-block'
						style='width: 200px'
						value={this.modifier}
						items={
							[
								{ text: t.inTheLast, value: 'inTheLast' },
								{ text: t.equals, value: 'equals' },
								{ text: t.between, value: 'between' },
								{ text: t.isAfter, value: 'isAfter' },
								{ text: t.isAfterOrOn, value: 'isAfterOrOn' },
								{ text: t.isBefore, value: 'isBefore' },
								{ text: t.isBeforeOrOn, value: 'isBeforeOrOn' },
							]
						}
						on={{ change: this.setModifier }}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div>
					{this.getInput(this.modifier || 'inTheLast')}
				</div>
				<div>

					<radioGroup
						on={{ change: this.setTimezone }}
						value={this.timezone || 'local'}
						row={true}
						label='Timezone'
						hideDetails={true}
					>
						<radio value='local' label={localTimezoneAbbr()}>
						</radio>
						<radio value='utc' label='UTC'></radio>
					</radioGroup>
				</div>
			</div >
		);
	},
});


export const NumberItem = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
		translations: {
			type: Object,
			// default: () => {
			// 	return {
			// 		equals: 'is equal to',
			// 		between: 'between',
			// 		greaterThan: 'is greater than',
			// 		lessThan: 'is less than',
			// 		and: 'and',
			// 	};
			// },
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || 'equals',
			valueIs: this.$props.value.valueIs,
			valueFrom: this.$props.value.valueFrom,
			valueTo: this.$props.value.valueTo,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setModifier(value: any) {
			this.modifier = value;
			this.inputEmit();
		},

		setNumber(value: any) {
			this.valueIs = value;
			this.inputEmit();
		},

		setNumberFrom(value: any) {
			this.valueFrom = value;
			this.inputEmit();
		},

		setNumberTo(value: any) {
			this.valueTo = value;
			this.inputEmit();
		},

		getInput(modifier: string) {
			const t = this.$props.translations;
			if (modifier === 'between') {
				return (
					<div>
						<vicon class='pr-5'>subdirectory_arrow_right</vicon>
						<vtextfield
							class='d-inline-block'
							style='width: 80px'
							type='number'
							on={{ change: this.setNumberFrom }}
							value={(this.valueFrom || '').toString()}
							hideDetails={true}
						/>
						<span class='px-3'>{t.and}</span>
						<vtextfield
							class='d-inline-block'
							style='width: 80px'
							type='number'
							on={{ change: this.setNumberTo }}
							value={(this.valueTo || '').toString()}
							hideDetails={true}
						/>
					</div>
				);
			}

			return (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<vtextfield
						class='d-inline-block'
						style='width: 120px'
						type='number'
						on={{ change: this.setNumber }}
						value={(this.valueIs || '').toString()}
						key={modifier}
						hideDetails={true}
					/>
				</div >
			);
		},
	},


	render() {
		const t = this.$props.translations;
		return (
			<div>
				<div>
					<vselect
						class='d-inline-block'
						style='width: 200px'
						on={{ change: this.setModifier }}
						value={this.modifier}
						items={
							[
								{ text: t.equals, value: 'equals' },
								{ text: t.between, value: 'between' },
								{ text: t.greaterThan, value: 'greaterThan' },
								{ text: t.lessThan, value: 'lessThan' },
							]
						}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div >
					{this.getInput(this.modifier)}
				</div>
			</div>
		);
	},
});


export const StringItem = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
		translations: {
			type: Object,
			// default: () => {
			// 	return {
			// 		equals: 'is equal to',
			// 		contains: 'contains',
			// 	};
			// },
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || 'contains',
			valueIs: this.$props.value.valueIs,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setModifier(value: any) {
			this.modifier = value;
			this.inputEmit();
		},

		setValue(value: any) {
			this.valueIs = value;
			this.inputEmit();
		},

		getInput(modifier: string) {
			return (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<vtextfield
						class='d-inline-block'
						style='width: 120px'
						type='text'
						onChange={this.setValue}
						value={this.valueIs}
						hideDetails={true}
					/>
				</div>
			);
		},
	},


	render() {
		const t = this.$props.translations;
		return (
			<div>
				<div>
					<vselect
						class='d-inline-block'
						style='width: 200px'
						on={{ change: this.setModifier }}
						value={this.modifier}
						items={
							[
								{ text: t.equals, value: 'equals' },
								{ text: t.contains, value: 'contains' },
							]
						}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div >
					{this.getInput(this.modifier)}
				</div>
			</div>
		);
	},
});


export const SelectItem = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
	},

	data() {
		return {
			valueIs: this.$props.value.valueIs || (this.$props.value.options && this.$props.value.options.length > 0 && this.$props.value.options[0].value),
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setValue(value: any) {
			this.valueIs = value;
			this.inputEmit();
		},
	},

	render() {
		return (
			<div>
				<vselect
					class='d-inline-block'
					style='width: 200px'
					hideDetails={true}
					on={{ change: this.setValue }}
					value={this.valueIs}
					items={this.value.options}
				>
				</vselect>
			</div>
		);
	},
});

export const MultipleSelectItem = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
		translations: {
			type: Object,
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || 'in',
			valuesAre: this.$props.value.valuesAre || [],
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setModifier(e: string) {
			this.modifier = e;
			this.inputEmit();
		},

		setValue(value: any) {
			this.inputEmit();
		},
	},

	render() {
		const t = this.$props.translations;
		return (
			<div>
				<div>
					<vselect
						class='d-inline-block'
						style='width: 200px'
						value={this.modifier}
						items={
							[
								{ text: t.in, value: 'in' },
								{ text: t.notIn, value: 'notIn' },
							]
						}
						on={{ change: this.setModifier }}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div style='max-height: 160px; overflow-y: scroll;'>
                    {this.value.options.map((opt: SelectOption) => {
                        return (
                        <v-checkbox
                            v-model={this.valuesAre}
                            on={{ change: this.setValue }}
                            label={opt.text}
                            value={opt.value}
                            hideDetails={true}
                            dense={true}
                        >
                        </v-checkbox>
                        );
                    })}
				</div>
			</div>
		);
	},
});

export const LinkageSelectItem = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
        linkageSelect: LinkageSelect,
	},
	props: {
		value: Object,
	},

	data() {
		return {
			valuesAre: this.$props.value.valuesAre || [],
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', { ...this.$props.value, ...this.$data });
		},

		setValue(value: any) {
			this.inputEmit();
		},
	},

	render() {
		return (
			<div>
                <vx-linkageselect 
                    items={this.value.linkageSelectData.items} 
                    labels={this.value.linkageSelectData.labels} 
                    selectOutOfOrder={this.value.linkageSelectData.selectOutOfOrder} 
                    v-model={this.valuesAre} 
                    on={{ input: this.setValue }}
                    row={true}
                    hideDetails={true}
                ></vx-linkageselect>
			</div>
		);
	},
});

/*
data = [
  {
    key: "created",
    label: "Created",
    itemType: "DateItem",
    selected: false,
    modifier: "between",
    valueFrom: new Date(),
    valueTo: new Date(),
  },
  {
    key: "updated",
    label: "Updated",
    itemType: "DateItem",
    selected: true,
    modifier: "inTheLast",
    inTheLastValue: 4,
    inTheLastUnit: "months",
  },
  {
    key: "amount",
    label: "Amount",
    itemType: "NumberItem",
    selected: true,
    modifier: "lessThan",
    value: 100,
  },
  {
    key: "name",
    label: "Name",
    itemType: "StringItem",
    value: "aaa",
  },
  {
    key: "refund",
    label: "Refund",
    itemType: "SelectItem",
    selected: true,
    options: [
      { key: "true", label: "is fully refunded" },
      { key: "false", label: "is not refunded" },
      { key: "partial", label: "is partially refunded" },
    ],
    value: "partial",
  },
]
*/

interface SelectOption {
	text: string;
	value: string;
}

interface FilterItem {
	key: string;
	label: string;
	itemType: string;
	modifier: string;
	valueIs: string;
    valuesAre: string[];
	selected?: boolean;
	valueFrom?: string;
	valueTo?: string;
	inTheLastValue?: string;
	inTheLastUnit?: string;
	timezone?: string;
	options?: SelectOption[];
}

function getSelectedIndexes(value: FilterItem[]): number[] {
	return (value).map((op: FilterItem, i: number) => {
		if (op.selected) {
			return i;
		}
		return -1;
	}).filter((i: number) => i !== -1);
}

function initInternalValue(items: FilterItem[]): FilterItem[] {
	for (const item of items) {
		if (item.itemType === 'SelectItem') {
			if (!item.valueIs && item.options) {
				item.valueIs = item.options[0].value;
			}
		}
	}
	return items;
}

export const Filter = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
		vbtn: VBtn,
		vmenu: VMenu,
		vexpPanels: VExpansionPanels,
		vexpPanel: VExpansionPanel,
		vexpPanelHeader: VExpansionPanelHeader,
		vexpPanelContent: VExpansionPanelContent,
		vtoolbar: VToolbar,
		vtoolbarTitle: VToolbarTitle,
		vspacer: VSpacer,
		vcheckbox: VCheckbox,
		vchip: VChip,
	},

	props: {
		value: { type: Array },
		replaceWindowLocation: Boolean,
		translations: {
			type: Object,
			default: () => {
				return {
					date: {
						inTheLast: 'is in the last',
						equals: 'is equal to',
						between: 'is between',
						isAfter: 'is after',
						isAfterOrOn: 'is on or after',
						isBefore: 'is before',
						isBeforeOrOn: 'is before or on',
						days: 'days',
						months: 'months',
						and: 'and',
					},
					number: {
						equals: 'is equal to',
						between: 'between',
						greaterThan: 'is greater than',
						lessThan: 'is less than',
						and: 'and',
					},
					string: {
						equals: 'is equal to',
						contains: 'contains',
					},
                    multipleSelect: {
                        in: 'in',
                        notIn: 'not in',
                    },
					clear: 'Clear',
					filters: 'Filters',
					filter: 'Filter',
					done: 'Done',
				};
			},
		} as any,
	},

	data() {
		return {
			internalValue: initInternalValue(this.$props.value as FilterItem[]),
			visible: false,
			selectedIndexs: getSelectedIndexes(this.$props.value),
		};
	},

	methods: {
		clickDone(e: any) {
			this.$emit('input', this.internalValue); // input event should be the same format as value

			// collect all query keys in the filter, remove them from location search first. then add it by selecting status
			// but keep original search conditions
			let filterKeys = (this.internalValue).map((op: FilterItem, i: number) => {
				return op.key;
			});

			const event = {
				filterKeys: filterKeys,
				filterData: filterData(this.internalValue),
				encodedFilterData: encodeFilterData(this.internalValue),
			};
			this.$emit('change', event);

			this.visible = false;
		},


		clear(e: any) {
			this.internalValue.map((op: any) => {
				op.selected = false;
			});
			this.selectedIndexs = getSelectedIndexes(this.internalValue);
		},

		togglePopup() {
			this.visible = !this.visible;
		},

		filterCount() {
			let count = 0;
			this.internalValue.map((op: any) => {
				if (op.selected) {
					count++;
				}
			});
			if (count === 0) {
				return;
			}
			return <vchip small={true}>{count}</vchip>;
		},

		onPanelExpand(value: any) {
			this.selectedIndexs = value;
			for (const fi of this.internalValue) {
				fi.selected = false;
			}
			for (const i of this.selectedIndexs) {
				this.internalValue[i].selected = true;
			}
		},

		newUpdateFilterItem(i: number): (val: FilterItem) => void {
			const self = this;
			return (val: FilterItem) => {
				self.internalValue[i] = val;
				self.internalValue[i].selected = true;
			};
		},
	},

	render() {

		const itemTypes: any = {
			DateItem,
			NumberItem,
			StringItem,
			SelectItem,
            MultipleSelectItem,
			LinkageSelectItem,
		};

		const t = this.$props.translations;

		const trans: any = {
			DateItem: t.date,
			NumberItem: t.number,
			StringItem: t.string,
			SelectItem: {},
            MultipleSelectItem: t.multipleSelect,
			LinkageSelectItem: {},
		};

		const body = this.internalValue.map((op: FilterItem, i: number) => {
			if (!itemTypes[op.itemType]) {
				throw new Error(`itemType '${op.itemType}' not supported`);
			}

			const itemComp = itemTypes[op.itemType];

			const comp = <itemComp
				translations={trans[op.itemType]}
				value={op}
				on={{ input: this.newUpdateFilterItem(i) }}
			/>;

			return (
				<vexpPanel
					value={op}
					key={op.key}
				>
					<vexpPanelHeader ripple={true}>
						<vcheckbox
							hideDetails={true}
							inputValue={this.selectedIndexs.includes(i)}
							label={op.label} class='ma-0'></vcheckbox>
					</vexpPanelHeader>
					<vexpPanelContent eager={false}>
						{comp}
					</vexpPanelContent>
				</vexpPanel>
			);
		});
		const self = this;
		return (
			<vmenu props={{ value: self.visible }} scopedSlots={{
				activator: ({ on }: any) => {
					return (<vbtn on={on} depressed>
						<vicon>filter_list</vicon>
						<span class='px-2'>{t.filter}</span>
						{this.filterCount()}
					</vbtn>);
				},
			}}
				offsetY={true}
                allowOverflow={true}
				// absolute={true}
				minWidth='400px'
				maxWidth='400px'
				closeOnContentClick={false}
				on={
					{
						input: (value: any) => {
							self.togglePopup();
						},
					}
				}
				zIndex='2'
			>
				<vtoolbar class='pb-1' color='grey lighten-5' flat={true}>
					<vbtn on={{ click: this.clear }} depressed={true}>{t.clear}</vbtn>
					<vspacer />
					<vtoolbarTitle class=''>
						{t.filters}
					</vtoolbarTitle>
					<vspacer />
					<vbtn color='primary' depressed={true} on={{ click: this.clickDone }}>
						{t.done}
					</vbtn>
				</vtoolbar>
				<vexpPanels
					on={{ change: this.onPanelExpand }}
					focusable={true}
					accordion={true}
					multiple={true}
					value={this.selectedIndexs}
				>{body}</vexpPanels>
			</vmenu>
		);
	},
});

