import Vue, { CreateElement, VNode, VNodeData } from 'vue';
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
	VDatePicker
} from 'vuetify/lib';
import VAutocomplete from './Autocomplete';
import * as constants from './Constants';
import DateTimePicker from './DateTimePicker.vue';
import DatePicker from './DatePicker.vue';
import { encodeFilterData, filterData } from './FilterData';
import LinkageSelect from './LinkageSelect.vue';



export const DatetimeRangeItem = Vue.extend({
	components: {
		datetimePicker: DateTimePicker,
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
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || constants.ModifierBetween,
			valueIs: this.$props.value.valueIs,
			valueFrom: this.$props.value.valueFrom,
			valueTo: this.$props.value.valueTo,
			inTheLastUnit: this.$props.value.inTheLastUnit,
			inTheLastValue: this.$props.value.inTheLastValue,
			datePickerVisible: false,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', {...this.$props.value, ...this.$data});
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
	},

	render(h: CreateElement): VNode {
		const t = this.$props.translations;
		const modifier = constants.ModifierBetween;

		return (
			<div>
				<datetimePicker
					value={this.valueFrom}
					on={{input: this.setDateFrom}}
					key={modifier + 'from'}
					visible={this.datePickerVisible}
					hideDetails={true}
				/>
				<span>{t.to}</span>
				<datetimePicker
					value={this.valueTo}
					on={{input: this.setDateTo}}
					key={modifier + 'to'}
					hideDetails={true}
				/>
			</div>
		);
	},
});

export const DateRangeItem = Vue.extend({
	components: {
		datePicker: DatePicker,
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
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier || constants.ModifierBetween,
			valueIs: this.$props.value.valueIs,
			valueFrom: this.$props.value.valueFrom,
			valueTo: this.$props.value.valueTo,
			inTheLastUnit: this.$props.value.inTheLastUnit,
			inTheLastValue: this.$props.value.inTheLastValue,
			datePickerVisible: false,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', {...this.$props.value, ...this.$data});
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
	},

	render(h: CreateElement): VNode {
		const t = this.$props.translations;
		const modifier = constants.ModifierBetween;

		return (
			<div>
				<datePicker
					value={this.valueFrom}
					on={{input: this.setDateFrom}}
					key={modifier + 'from'}
					visible={this.datePickerVisible}
					hideDetails={true}
				/>
				<span>{t.to}</span>
				<datePicker
					value={this.valueTo}
					on={{input: this.setDateTo}}
					key={modifier + 'to'}
					hideDetails={true}
				/>
			</div>
		);
	},
});

export const DateItem = Vue.extend({
	components: {
		datePicker: VDatePicker,
	},
	props: {
		value: Object,
	},

	data() {
		return {
			valueIs: this.$props.value.valueIs,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', {...this.$props.value, ...this.$data});
		},

		setValue(value: any) {
			this.valueIs = value;
			this.inputEmit();
		},
	},

	render() {
		return (
			<div class="pt-3">
				<datePicker
					value={this.valueIs}
					on={{change: this.setValue}}
					noTitle={true}
					hideDetails={true}
				/>
			</div>
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
			this.$emit('input', {...this.$props.value, ...this.$data});
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
							on={{change: this.setNumberFrom}}
							value={(this.valueFrom || '').toString()}
							hideDetails={true}
						/>
						<span class='px-3'>{t.and}</span>
						<vtextfield
							class='d-inline-block'
							style='width: 80px'
							type='number'
							on={{change: this.setNumberTo}}
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
						on={{change: this.setNumber}}
						value={(this.valueIs || '').toString()}
						key={modifier}
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
						on={{change: this.setModifier}}
						value={this.modifier}
						items={
							[
								{text: t.equals, value: 'equals'},
								{text: t.between, value: 'between'},
								{text: t.greaterThan, value: 'greaterThan'},
								{text: t.lessThan, value: 'lessThan'},
							]
						}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div>
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
			this.$emit('input', {...this.$props.value, ...this.$data});
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
						on={{change: this.setModifier}}
						value={this.modifier}
						items={
							[
								{text: t.equals, value: 'equals'},
								{text: t.contains, value: 'contains'},
							]
						}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div>
					{this.getInput(this.modifier)}
				</div>
			</div>
		);
	},
});


export const SelectItem = Vue.extend({
	components: {
		vautocomplete: VAutocomplete,
		vtextfield: VTextField,
		vicon: VIcon,
	},
	props: {
		value: Object,
	},

	data() {
		return {
			valueIs: this.$props.value.valueIs,
		};
	},

	methods: {
		inputEmit() {
			this.$emit('input', {...this.$props.value, ...this.$data});
		},

		setValue(value: any) {
			this.valueIs = value;
			this.inputEmit();
		},
	},

	render() {
		const data: VNodeData = {
			props: {
				items: this.value.options,
				...this.$props.value.autocompleteDataSource,
			},
			attrs: {
				value: this.valueIs,
				hideDetails: true,
				class:'d-inline-block',
				style:'width: 200px',
			},
			on: {
				change: this.setValue,
			},
		};

		return (
			<div>
				<vautocomplete {...data}>
				</vautocomplete>
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
			this.$emit('input', {...this.$props.value, ...this.$data});
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
								{text: t.in, value: 'in'},
								{text: t.notIn, value: 'notIn'},
							]
						}
						on={{change: this.setModifier}}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div style='max-height: 160px; overflow-y: scroll;'>
					{this.value.options.map((opt: SelectOption) => {
						return (
							<v-checkbox
								v-model={this.valuesAre}
								on={{change: this.setValue}}
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
			this.$emit('input', {...this.$props.value, ...this.$data});
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
					on={{input: this.setValue}}
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
    itemType: "DatetimeRangeItem",
    selected: false,
    modifier: "between",
    valueFrom: new Date(),
    valueTo: new Date(),
  },
  {
    key: "updated",
    label: "Updated",
    itemType: "DatetimeRangeItem",
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

interface IndependentTranslations {
	filterBy: string;
}

interface FilterItem {
	key: string;
	label: string;
	folded: boolean;
	itemType: string;
	modifier: string;
	valueIs: string;
	valuesAre: string[];
	selected?: boolean;
	valueFrom?: string;
	valueTo?: string;
	inTheLastValue?: string;
	inTheLastUnit?: string;
	options?: SelectOption[];
	translations?: IndependentTranslations;
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
	return items;
}

export const Filter = Vue.extend({
	components: {
		vselect: VSelect,
		vautocomplete: VAutocomplete,
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
		value: {type: Array},
		replaceWindowLocation: Boolean,
		translations: {
			type: Object,
			default: () => {
				return {
					date: {
						to: 'to',
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
					clear: 'Clear Filters',
					add: 'Add Filters',
					apply: 'Apply',
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
			const filterKeys = (this.internalValue).map((op: FilterItem, i: number) => {
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


		clearAll(e: any) {
			this.internalValue.map((op: any) => {
				op.selected = false;
			});
			this.selectedIndexs = getSelectedIndexes(this.internalValue);
			this.clickDone(e);
		},

		clear(e: any, op: FilterItem) {
			if (!op.selected) {
				return
			}

			op.selected = false
			this.selectedIndexs = getSelectedIndexes(this.internalValue);
			this.clickDone(e);
			e.stopPropagation()
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
			return (val: FilterItem) => {
				this.internalValue[i] = val;
				this.internalValue[i].selected = true;
			};
		},

		filterButton(op: FilterItem, on: any, isFoldedItem: boolean) {
			let showValue = '';
			if (op.selected) {
				switch (op.itemType) {
					case 'DatetimeRangeItem':
					case 'DateRangeItem': {
						const mod = op.modifier || constants.ModifierBetween;

						if (mod === constants.ModifierBetween) {
							if (op.valueFrom) {
								if (op.valueTo) {
									showValue = `${op.valueFrom} - ${op.valueTo}`
								} else {
									showValue = ` >= ${op.valueFrom}`
								}
							} else {
								if (op.valueTo) {
									showValue = ` < ${op.valueTo}`
								}
							}
						}
						break
					}
					case 'DateItem': {
						showValue = op.valueIs;
						break
					}
					case 'NumberItem': {
						const mod = op.modifier || 'equals';

						if (mod === 'equals') {
							const floatValue = parseFloat(op.valueIs);
							if (!isNaN(floatValue)) {
								showValue += floatValue
							}
						}

						if (mod === 'between') {
							const floatFrom = parseFloat(op.valueFrom || '');
							const floatTo = parseFloat(op.valueTo || '');
							const fromValid = !isNaN(floatFrom)
							const toValid = !isNaN(floatTo)
							if (fromValid) {
								if (toValid) {
									showValue = `${op.valueFrom} - ${op.valueTo}`
								} else {
									showValue = ` >= ${op.valueFrom}`
								}
							} else {
								if (toValid) {
									showValue = ` <= ${op.valueTo}`
								}
							}
						}

						if (mod === 'greaterThan') {
							const floatValue = parseFloat(op.valueIs);
							if (!isNaN(floatValue)) {
								showValue += ' > ' +  op.valueFrom
							}
						}

						if (mod === 'lessThan') {
							const floatValue = parseFloat(op.valueIs);
							if (!isNaN(floatValue)) {
								showValue += ' < ' +  op.valueTo
							}
						}
						break
					}
					case 'StringItem': {
						const mod = op.modifier || 'equals';
						if (mod === 'equals' && op.valueIs) {
							showValue = op.valueIs
						}

						if (mod === 'contains' && op.valueIs) {
							showValue = ' ~ ' + op.valueIs
						}
						break
					}
					case 'SelectItem': {
						const mod = op.modifier || 'equals';
						if (mod === 'equals' && op.valueIs) {
							showValue = op.options!.find(o => o.value === op.valueIs)!.text
						}
						break
					}
					case 'MultipleSelectItem': {
						const mod = op.modifier || 'in';
						const textsAre = op.options!.filter(o => op.valuesAre.includes(o.value)).map(o => o.text)
						if (mod === 'in' && op.valuesAre && op.valuesAre.length > 0) {
							showValue = ' in ' + '[ ' + textsAre.join(', ') + ' ]'
						}
						if (mod === 'notIn' && op.valuesAre && op.valuesAre.length > 0) {
							showValue = ' not in ' + '[ ' + textsAre.join(', ') + ' ]'
						}
						break
					}
					case 'LinkageSelectItem': {
						const mod = op.modifier || 'equals';
						const textsAre = op.options!.filter(o => op.valuesAre.includes(o.value)).map(o => o.text)
						if (mod === 'equals' && op.valuesAre && op.valuesAre.length > 0) {
							showValue = textsAre.join(', ')
						}
						break
					}
					default:
						throw new Error(`itemType '${op.itemType}' not supported`);
				}
			}

			const showValueCopy = showValue
			showValue = ""
			let showLen = 0
			for (let i = 0; i < showValueCopy.length; i++) {
				showValue += showValueCopy.charAt(i)
				if (showValueCopy.charCodeAt(i) > 127) {
					showLen += 2
				} else {
					showLen++
				}
				if (showLen > 66) {
					showValue += '...'
					break
				}
			}

			const body = (
				<span>
					<vicon
						left={true}
						on={{click: (e: any) => this.clear(e, op)}}
					>
						{op.selected ? 'cancel' : 'add_circle'}
					</vicon>
					{/*`overflow hidden` cases vertical align issue*/}
					{/*<span class={'d-inline-block text-truncate'} style={'max-width: 500px;'}>*/}
					<span>
						{op.label}
						{op.selected ?
							<span> | <span class={'primary--text'}>{showValue}</span></span>
							: null
						}
					</span>
				</span>
			)

			if (isFoldedItem) {
				return (
					<div on={on} class={'my-1 pa-1'} style={'cursor: pointer; user-select: none;'}>{body}</div>
				)
			}
			return (
				<vchip
					on={on}
					outlined={true}
					class={`mr-2 my-1 ${op.selected ? '' : 'grey--text text--darken-1'}`}
					style={{borderStyle: op.selected ? 'solid' : 'dashed'}}
				>{body}</vchip>
			)
		},

		filtersGetFunc(f: (item: FilterItem) => boolean, isFoldedItem: boolean) {
			return (itemTypes: any, trans: any) => {
				return this.internalValue.map((op: FilterItem, i: number) => {
					if (!f(op)) {
						return null
					}

					if (!itemTypes[op.itemType]) {
						throw new Error(`itemType '${op.itemType}' not supported`);
					}

					const itemComp = itemTypes[op.itemType];

					const comp = <itemComp
						translations={trans[op.itemType]}
						value={op}
						on={{input: this.newUpdateFilterItem(i)}}
					/>;

					return (
						<vmenu
							scopedSlots={{
								activator: ({on}: any) => {
									return this.filterButton(op, on, isFoldedItem)
								}
							}}
							offsetY={true}
							closeOnContentClick={false}
							rounded={'md'}
						>
							<div class={'pa-3 white'}>
								<div>{op.translations?.filterBy}</div>
								{comp}
								<vbtn
									class={'mt-3'}
									color='primary'
									depressed={true}
									on={{click: this.clickDone}}
								>
									{this.$props.translations.apply}
								</vbtn>
							</div>
						</vmenu>
					);
				}).filter(v => {
					return v != null
				});
			}
		},
	},

	render() {

		const itemTypes: any = {
			DatetimeRangeItem,
			DateRangeItem,
			DateItem,
			NumberItem,
			StringItem,
			SelectItem,
			MultipleSelectItem,
			LinkageSelectItem,
		};

		const t = this.$props.translations;
		const trans: any = {
			DatetimeRangeItem: t.date,
			DateRangeItem: t.date,
			DateItem: t.date,
			NumberItem: t.number,
			StringItem: t.string,
			SelectItem: {},
			MultipleSelectItem: t.multipleSelect,
			LinkageSelectItem: {},
		};

		const fixedFilters = this.filtersGetFunc(item => !item.folded, false)(itemTypes, trans)
		const otherSelectedFilters = this.filtersGetFunc(item => item.folded && !!item.selected, false)(itemTypes, trans)
		const foldedFilters = this.filtersGetFunc(item => item.folded && !item.selected, true)(itemTypes, trans)

		return (
            <div class={'d-flex flex-grow-1'}>
				<div>
					{fixedFilters}
					{otherSelectedFilters}
				</div>
				<vspacer/>
				<vbtn
					on={{click: this.clearAll}}
					plain={true}
					small={true}
					disabled={this.internalValue.findIndex(item => item.selected) < 0}
					class={'my-1'}
				>
					<vicon small={true}>close</vicon> {t.clear}
				</vbtn>
				{foldedFilters.length > 0 ?
					<vmenu
						scopedSlots={{
							activator: ({on}: any) => {
								return <vbtn
									on={on}
									plain={true}
									small={true}
									color={'primary'}
									class={'my-1'}
								>
									<vicon small={true}>filter_alt</vicon> {t.add}
								</vbtn>
							}
						}}
						offsetY={true}
						closeOnContentClick={false}
						rounded={'md'}
					>
						<div class={'pa-2 white'}>{foldedFilters}</div>
					</vmenu>
					: null
				}

            </div>
		);
	},
});
