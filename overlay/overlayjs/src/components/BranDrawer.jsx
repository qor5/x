import Drawer from "ant-design-vue/lib/vc-drawer/src/Drawer";
import drawerProps from "ant-design-vue/lib/vc-drawer/src/IDrawerPropTypes";
import "ant-design-vue/lib/vc-drawer/assets/index.less";
import { getOptionProps } from 'ant-design-vue/lib/_util/props-util';

const levelMoveFactory = (level) => (e) => {
	const target = e.target;
	// console.log("levelMoveFactory", e, level)
	if (level == "all") {
		return 0
	}

	let mylevels = level
	if (!Array.isArray(level)) {
		mylevels = [level]
	}


	for (let i = 0; i < mylevels.length; i++) {
		const levelClass = mylevels[i].slice(1)
		const max = mylevels.length - i;
		if (target.className.indexOf(levelClass) >= 0) {
			let r = []
			for (let j = 0; j < max; j++) {
				r.push(100 * (max - j))
			}
			if (r.length == 0) {
				return r[0]
			}
			// console.log("return", r)
			return r
		}
	}

	return 0
}

export default {
	name: "BranDrawer",
	// props: ["width", "title", "placement"],
	props: drawerProps,
	data() {
		return {
			isVisible: this.defaultOpen,
			isFirstEnter: this.firstEnter,
			destroyOnClose: false,
			timeout: null,
		};
	},
	methods: {
		show() {
			this.destroyOnClose = false;
			this.isVisible = true;
		},
		close() {
			this.isVisible = false;
			this.isFirstEnter = false;
			clearTimeout(this.timeout)
			this.timeout = setTimeout(() => {
				this.destroyOnClose = true;
				this.$forceUpdate();
			}, 300);
		},
		onMaskClick(e) {
			this.close(e);
		},
	},
	render() {
		const props = getOptionProps(this)
		const { open, getContainer, handler, level, levelMove, firstEnter, ...rest } = props
		const defaultGetContainer = getContainer || "#app"
		const vcProps = {
			props: {
				open: this.isVisible,
				getContainer: defaultGetContainer,
				handler: false,
				levelMove: levelMoveFactory(level),
				firstEnter: this.isFirstEnter,
				level,
				...rest
			},
			on: {
				maskClick: this.onMaskClick,
				...this.$listeners,
			}
		}

		return (
			<div class="bran-drawer">
				{this.$scopedSlots.trigger && this.$scopedSlots.trigger({ parent: this })}
				{!this.destroyOnClose && <Drawer {...vcProps}>
					{this.$scopedSlots.drawer({ parent: this })}
				</Drawer>}
			</div>
		)
	}
};
