import Vue, { VNode, CreateElement } from 'vue';

import {
	VDatePicker,
	VTextField,
	VMenu,
} from 'vuetify/lib';


export default Vue.extend({
	components: {
		datePicker: VDatePicker,
		vtextfield: VTextField,
		vmenu: VMenu,
	},
	props: {
		value: String,
		visible: Boolean,
	},
	data() {
		return {
			internalVisible: this.visible,
			internalValue: this.value,
		};
	},

	methods: {
		onChange(value: any) {
			this.internalVisible = false;
			this.internalValue = value;
			this.$emit('input', this.internalValue);
		},
		toggle() {
			this.internalVisible = !this.internalVisible;
		},
	},

	render(h: CreateElement): VNode {
		const self = this;
		return (<vmenu class='d-inline-block' props={{ value: this.internalVisible }} scopedSlots={{
			activator: ({ on }: any) => {
				return <vtextfield
					class='d-inline-block'
					style='width: 120px'
					on={on}
					value={self.internalValue}
					hideDetails={true}
					prependInnerIcon='event'
				></vtextfield>;
			},
		}}
			offsetY={true}
			eager={true}
			fixed={true}
			fullWidth={true}
			minWidth='290px'
			closeOnContentClick={false}
			on={
				{
					input: (value: any) => {
						self.toggle();
					},
				}
			}
		>
			<datePicker
				value={this.internalValue}
				on={{ change: this.onChange }}
			/>
		</vmenu >);
	},
});
