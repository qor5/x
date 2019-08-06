import { encodeFilterData, filterData } from './FilterData';
import Vue, { VNode, Component, CreateElement } from 'vue';
import {
	VRadio,
	VRadioGroup,
	VDatePicker,
	VSelect,
	VTextField,
	VContainer,
	VLayout,
	VFlex,
	VMenu,
	VIcon,
} from 'vuetify/lib';

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

const TextDatePicker = Vue.extend({
	components: {
		datePicker: VDatePicker,
		vtextfield: VTextField,
		vmenu: VMenu,
	},
	props: {
		// value: String,
		// onChange: {},
	},
	data() {
		return {
			show: false,
			value: '',
		};
	},

	methods: {
		onChange(value: any) {
			this.show = false;
			this.value = value;
		},
		toggle() {
			this.show = !this.show;
		},
	},
	render(h: CreateElement) {
		const self = this;
		return (<vmenu class='d-inline-block' props={{ value: this.show }} scopedSlots={{
			activator: ({ on }: any) => {
				return <vtextfield
					class='d-inline-block'
					on={on}
					value={self.value}
					hideDetails={true}
					prependInnerIcon='event'
				></vtextfield>;
			},
		}}
			offsetY={true}
			fullWidth={true}
			minWidth='290px'
			closeOnContentClick={false}
			on={
				{
					input: (event: any) => {
						self.toggle();
					},
				}
			}
		>
			<datePicker
				value={this.value}
				on={{ change: this.onChange }}
				inputFormat='MM/DD/YYYY'
				active={true}
			/>
		</vmenu >);
	},
});

export const DateItem = Vue.extend({
	components: {
		datePicker: TextDatePicker,
		radioGroup: VRadioGroup,
		radio: VRadio,
		vcontainer: VContainer,
		vselect: VSelect,
		vtextfield: VTextField,
		vlayout: VLayout,
		vflex: VFlex,
		vicon: VIcon,
	},
	props: {
		data: Object,
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
				};
			},
		},
	},

	data() {
		return {
			// data: this.$props.data,
		};
	},

	methods: {
		onSelectChange(e: string) {
			this.data.modifier = e;
			this.$forceUpdate();
		},

		setDate(e: any) {
			this.data.value = e;
			this.$forceUpdate();
		},

		setDateFrom(e: any) {
			this.data.valueFrom = e;
			this.$forceUpdate();
		},

		setDateTo(e: any) {
			this.data.valueTo = e;
			this.$forceUpdate();
		},

		setInTheLastValue(e: any) {
			this.data.inTheLastValue = e;
			this.$forceUpdate();
		},

		setInTheLastUnit(e: any) {
			this.data.inTheLastUnit = e;
			this.$forceUpdate();
		},

		setTimezone(e: any) {
			this.data.timezone = e;
			this.$forceUpdate();
		},

		getInput(modifier: string) {
			const t = this.$props.translations;
			const inTheLast = (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<vtextfield
						class='d-inline-block pr-5'
						type='number'
						on={{ change: this.setInTheLastValue }}
						value={(this.data.inTheLastValue || '').toString()}
						hideDetails={true}
					/>
					<vselect
						class='d-inline-block'
						on={{ change: this.setInTheLastUnit }}
						value={this.data.inTheLastUnit || 'days'}
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

			const between = (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<datePicker
						date={this.data.valueFrom}
						onChange={this.setDateFrom}
						inputFormat='MM/DD/YYYY'
						active={true}
					/>{' '}
					<span class='pr-5'>and</span>
					<datePicker
						date={this.data.valueTo}
						onChange={this.setDateTo}
						inputFormat='MM/DD/YYYY'
						active={true}
					/>
				</div>
			);

			const inputs: any = {
				inTheLast,
				between,
			};

			if (inputs[modifier]) {
				return inputs[modifier];
			}

			return (
				<div>
					<vicon class='pr-5'>subdirectory_arrow_right</vicon>
					<datePicker
						date={this.data.value}
						onChange={this.setDate}
						inputFormat='MM/DD/YYYY'
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
						value={this.data.modifier}
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
						on={{ change: this.onSelectChange }}
						hideDetails={true}
					>
					</vselect>
				</div>
				<div>
					{this.getInput(this.data.modifier || 'inTheLast')}
				</div>
				<div>

					<radioGroup
						on={{ change: this.setTimezone }}
						value={this.data.timezone || 'local'}
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

/*
const NumberItem = Vue.extend({
	public static propTypes = {
		data: PropTypes.object.isRequired,
	};

	public state = {
		data: this.props.data,
	};

	public onSelectChange = (e) => {
		this.state.data.modifier = e.target.value;
		this.forceUpdate();
	};

	public setNumber = (e) => {
		this.state.data.value = e.target.value;
		this.forceUpdate();
	};

	public setNumberFrom = (e) => {
		this.state.data.valueFrom = e.target.value;
		this.forceUpdate();
	};

	public setNumberTo = (e) => {
		this.state.data.valueTo = e.target.value;
		this.forceUpdate();
	};

	public getInput(modifier); {
		const between = (
			<div className={classNames(styles.between, styles.number)}>
				<input
					type='number'
					onChange={this.setNumberFrom}
					value={(this.state.data.valueFrom || '').toString()}
				/>
				<span>and</span>
				<input
					type='number'
					onChange={this.setNumberTo}
					value={(this.state.data.valueTo || '').toString()}
				/>
			</div>
		);

		const inputs = {
			between,
		};

		if(inputs[modifier]) {
			return inputs[modifier];
		}

	return(
		<div className = { styles.number } >
				<input
					type='number'
					onChange={this.setNumber}
					value={(this.state.data.value || '').toString()}
				/>;
		</div > ;
	)
}

	public render(); {
	const t = this.props.translations;
	return (
		<div className={styles.control}>
			<div className={styles.modifierContainer}>
				<select
					onChange={this.onSelectChange}
					value={this.state.data.modifier}
				>
					<option value='equals'>{t.equals}</option>
					<option value='between'>{t.between}</option>
					<option value='greaterThan'>{t.greaterThan}</option>
					<option value='lessThan'>{t.lessThan}</option>
				</select>
			</div>
			<div className={styles.inputContainer}>
				{this.getInput(this.state.data.modifier)}
			</div>
		</div>
	);
}
})

const StringItem = Vue.extend({
	public static propTypes = {
		data: PropTypes.object.isRequired,
	};

	public state = {
		data: this.props.data,
	};

	public onSelectChange = (e) => {
		this.state.data.modifier = e.target.value;
		this.forceUpdate();
	}

	public setValue = (e) => {
		this.state.data.value = e.target.value;
		this.forceUpdate();
	}

	public getInput(modifier) {
		return (
			<div className={styles.string}>
				<input
					type='text'
					onChange={this.setValue}
					value={this.state.data.value}
				/>
			</div>
		);
	}

	public render() {
		return (
			<div className={styles.control}>
				<div className={styles.modifierContainer}>
					<select
						onChange={this.onSelectChange}
						value={this.state.data.modifier}
					>
						<option value='equals'>is equal to</option>
						<option value='contains'>contains</option>
					</select>
				</div>
				<div className={styles.inputContainer}>
					{this.getInput(this.state.data.modifier)}
				</div>
			</div>
		);
	},
});

const SelectItem = Vue.extend({
	public static propTypes = {
		data: PropTypes.object.isRequired,
	};

	public state = {
		data: this.defaultData(this.props.data),
	};

	public defaultData = (data) => {
		if (!data.value) {
			data.value = data.options[0].key;
		}
		return data;
	}

	public setValue = (e) => {
		this.state.data.value = e.target.value;
		this.forceUpdate();
	}

	public render() {
		const ops = this.props.data.options.map((op) => {
			return (
				<option value={op.key} key={op.key}>
					{op.label}
				</option>
			);
		});
		return (
			<div className={classNames(styles.control, styles.selectContainer)}>
				<select onChange={this.setValue} value={this.state.data.value}>
					{ops}
				</select>
			</div>
		);
	},
};

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

/*
export const Filter = Vue.extend({
	public static propTypes = {
		data: PropTypes.any.isRequired,
		onDone: PropTypes.func.isRequired,
	};

	public static defaultProps = {
		translations: {
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
		},
	};

	public state = {
		data: this.props.data,
		hidden: true,
	};

	public itemTypes = {
		DateItem,
		NumberItem,
		StringItem,
		SelectItem,
	};

	public clickDone = (e) => {
		this.props.onDone(
			filterData(this.state.data),
			encodeFilterData(this.state.data),
		);
		this.setState({ hidden: true });
	}

	public clear = (e) => {
		this.state.data.map((op) => {
			op.selected = false;
		});
		this.forceUpdate();
	}

	public update = () => {
		this.forceUpdate();
	}

	public togglePopup = (e) => {
		this.setState({ hidden: !this.state.hidden });
	}

	public filterCount = () => {
		let count = 0;
		this.state.data.map((op) => {
			if (op.selected) {
				count++;
			}
		});
		if (count == 0) {
			return;
		}

		return <span className={styles.filterCount}>{count}</span>;
	}

	public render() {
		const t = this.props.translations;

		const trans = {
			DateItem: t.date,
			NumberItem: t.number,
			StringItem: t.string,
			SelectItem: {},
		};

		const body = this.state.data.map((op) => {
			if (!this.itemTypes[op.itemType]) {
				throw new Error(`itemType '${op.itemType}' not supported`);
			}

			const comp = React.createElement(this.itemTypes[op.itemType], {
				translations: trans[op.itemType],
				data: op,
			});

			return (
				<CheckboxToogle
					data={op}
					key={op.key}
					filterUpdate={this.update}
				>
					{comp}
				</CheckboxToogle>
			);
		});

		return (
			<div className={classNames(styles.filter, this.props.className)}>
				<Button onMouseDown={this.togglePopup}>
					<Icon
						glyph={filterIcon}
						className={styles.filterIcon}
					/>
					Filter{this.filterCount()}
				</Button>
				<Popup hidden={this.state.hidden}>
					<div className={styles.header}>
						<Button onMouseDown={this.clear}>{t.clear}</Button>
						<h2>
							<span>{t.filters}</span>
						</h2>
						<Button primary onMouseDown={this.clickDone}>
							{t.done}
						</Button>
					</div>
					<div>{body}</div>
				</Popup>
			</div>
		);
	},
});
*/
