import { encodeFilterData, filterData } from './FilterData';
import Vue, { VNode, CreateElement } from 'vue';
import {
	VRadio,
	VRadioGroup,
	VDatePicker,
	VSelect,
	VTextField,
	VMenu,
	VIcon,
	VBtn,
	VCard,
	VCardText,
	VCardActions,
	VExpansionPanels,
	VExpansionPanel,
	VExpansionPanelHeader,
	VExpansionPanelContent,
	VToolbar,
	VToolbarTitle,
	VSpacer,
} from 'vuetify/lib';

import TextDatePicker from './TextDatePicker';

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
			default: () => {
				return {
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
				};
			},
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier,
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
		setModifier(e: string) {
			this.modifier = e;
			this.$emit('input', this.$data);
			this.datePickerVisible = true;
			this.$forceUpdate();
		},

		setDate(e: any) {
			this.valueIs = e;
			this.$emit('input', this.$data);
		},

		setDateFrom(e: any) {
			this.valueFrom = e;
			this.$emit('input', this.$data);
		},

		setDateTo(e: any) {
			this.valueTo = e;
			this.$emit('input', this.$data);
		},

		setInTheLastValue(e: any) {
			this.inTheLastValue = e;
			this.$emit('input', this.$data);
		},

		setInTheLastUnit(e: any) {
			this.inTheLastUnit = e;
			this.$emit('input', this.$data);
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
							visible={this.datePickerVisible}
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
			default: () => {
				return {
					equals: 'is equal to',
					between: 'between',
					greaterThan: 'is greater than',
					lessThan: 'is less than',
					and: 'and',
				};
			},
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier,
			valueIs: this.$props.value.valueIs,
			valueFrom: this.$props.value.valueFrom,
			valueTo: this.$props.value.valueTo,
		};
	},

	methods: {

		setModifier(value: any) {
			this.modifier = value;
			this.$emit('input', this.$data);
		},

		setNumber(value: any) {
			this.valueIs = value;
			this.$emit('input', this.$data);
		},

		setNumberFrom(value: any) {
			this.valueFrom = value;
			this.$emit('input', this.$data);
		},

		setNumberTo(value: any) {
			this.valueTo = value;
			this.$emit('input', this.$data);
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
			default: () => {
				return {
					equals: 'is equal to',
					contains: 'contains',
				};
			},
		},
	},

	data() {
		return {
			modifier: this.$props.value.modifier,
			valueIs: this.$props.value.valueIs,
		};
	},

	methods: {

		setModifier(value: any) {
			this.modifier = value;
			this.$emit('input', this.$data);
		},

		setValue(value: any) {
			this.valueIs = value;
			this.$emit('input', this.$data);
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
			valueIs: this.$props.value.valueIs || this.$props.value.options[0].value,
		};
	},

	methods: {
		setValue(value: any) {
			this.valueIs = value;
			this.$emit('input', this.$data);

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

/*
const CheckboxToogle = Vue.extend({
	public static propTypes = {
		data: PropTypes.object.isRequired,
		filterUpdate: PropTypes.func,
	};

	public state = {
		data: this.props.data,
	};

	public onCheckboxChange = (e) => {
		this.state.data.selected = e.target.checked;
		const update = this.props.filterUpdate || this.forceUpdate;
		update();
	}

		public render() {
		const classes = classNames({
			[styles.hidden]: !this.state.data.selected,
		});

		return (
			<div className={styles.checkboxToggle}>
				<label>
					<input
						type='checkbox'
						value={this.state.data.value}
						onChange={this.onCheckboxChange}
						checked={this.state.data.selected ? 'checked' : ''}
					/>
					<span>{this.state.data.label}</span>
				</label>
				<div className={classes}>{this.props.children}</div>
			</div>
		);
	},
});

*/
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
	selected?: boolean;
	valueFrom?: string;
	valueTo?: string;
	inTheLastValue?: string;
	inTheLastUnit?: string;
	timezone?: string;
	options?: SelectOption[];
}

export const Filter = Vue.extend({
	components: {
		vselect: VSelect,
		vtextfield: VTextField,
		vicon: VIcon,
		// dateItem: DateItem,
		// numberItem: NumberItem,
		// stringItem: StringItem,
		// selectItem: SelectItem,
		vbtn: VBtn,
		vmenu: VMenu,
		vexpPanels: VExpansionPanels,
		vexpPanel: VExpansionPanel,
		vexpPanelHeader: VExpansionPanelHeader,
		vexpPanelContent: VExpansionPanelContent,
		vcard: VCard,
		vcardText: VCardText,
		vtoolbar: VToolbar,
		vtoolbarTitle: VToolbarTitle,
		vspacer: VSpacer,
	},

	props: {
		value: { type: Array },
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
					},
					number: {
						equals: 'is equal to',
						between: 'between',
						greaterThan: 'is greater than',
						lessThan: 'is less than',
					},
					string: {
						equals: 'is equal to',
						contains: 'contains',
					},
					clear: 'Clear',
					filters: 'Filters',
					done: 'Done',
				};
			},
		},
	},

	data() {
		return {
			internalValue: this.$props.value as FilterItem[],
			visible: false,
		};
	},

	methods: {

		clickDone(e: any) {
			this.$emit('input', {
				filterData: filterData(this.internalValue),
				encodeFilterData: encodeFilterData(this.internalValue),
			});
			this.visible = false;
		},

		clear(e: any) {
			this.internalValue.map((op: any) => {
				op.selected = false;
			});
			this.$forceUpdate();
		},

		update() {
			this.$forceUpdate();
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

			return <span>{count}</span>;
		},
	},


	render() {

		const itemTypes: any = {
			DateItem,
			NumberItem,
			StringItem,
			SelectItem,
		};

		const t = this.$props.translations;

		const trans: any = {
			DateItem: t.date,
			NumberItem: t.number,
			StringItem: t.string,
			SelectItem: {},
		};

		const body = this.internalValue.map((op: FilterItem) => {
			if (!itemTypes[op.itemType]) {
				throw new Error(`itemType '${op.itemType}' not supported`);
			}

			const itemComp = itemTypes[op.itemType];

			const comp = <itemComp
				translations={trans[op.itemType]}
				value={op}
			/>;

			return (
				<vexpPanel
					value={op}
					key={op.key}
					filterUpdate={this.update}
				>
					<vexpPanelHeader class='subtitle-2'>
						{op.label}
					</vexpPanelHeader>
					<vexpPanelContent>
						{comp}
					</vexpPanelContent>
				</vexpPanel>
			);
		});
		const self = this;
		return (
			<vmenu props={{ value: self.visible }} scopedSlots={{
				activator: ({ on }: any) => {
					return (<vbtn on={on}>
						<vicon>filter_list</vicon>
						Filter
				{this.filterCount()}
					</vbtn>);
				},
			}}
				offsetY={true}
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
				<vtoolbar>
					<vbtn on={{ click: this.clear }}>{t.clear}</vbtn>
					<vtoolbarTitle>
						{t.filters}
					</vtoolbarTitle>
					<vspacer />
					<vbtn color='primary' on={{ click: this.clickDone }}>
						{t.done}
					</vbtn>
				</vtoolbar>
				<vexpPanels multiple={true}>{body}</vexpPanels>
			</vmenu>
		);
	},
});

