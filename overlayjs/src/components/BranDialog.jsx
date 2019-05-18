import Dialog from "ant-design-vue/lib/vc-dialog";
import getDialogPropTypes from "ant-design-vue/lib/vc-dialog/IDialogPropTypes";
import "ant-design-vue/lib/vc-dialog/assets/index.less";
import { getOptionProps } from 'ant-design-vue/lib/_util/props-util';

export default {
	name: "BranDialog",
	props: getDialogPropTypes(),
	data() {
		return {
			isVisible: this.visible,
		};
	},
	methods: {
		show() {
			this.isVisible = true;
		},
		close() {
			this.isVisible = false;
		},
	},
	render() {
		const props = getOptionProps(this)
		const { visible, getContainer, destroyOnClose, animation, ...rest } = props
		const appContainer = () => {
			return document.getElementById("app")
		}
		let newAnimation = animation
		if (this.visible) {
			newAnimation = ""
		}

		const vcProps = {
			props: {
				visible: this.isVisible,
				getContainer: appContainer,
				destroyOnClose: true,
				animation: newAnimation,
				...rest
			},
			on: {
				...this.$listeners,
				close: this.close,
			}
		}

		return (
			<div class="bran-dialog">
				{this.$scopedSlots.trigger && this.$scopedSlots.trigger({ parent: this })}
				{<Dialog {...vcProps}>
					{this.$scopedSlots.dialog({ parent: this })}
				</Dialog>}
			</div>
		)
	}
};
//