<template>
<div class="icon" :class="[`icon--${name}`, `icon--${size}`]">
	<component :is="icon" />
</div>
</template>

<script>
import add_col_after from "./icons/add_col_after.svg";
import add_col_before from "./icons/add_col_before.svg";
import add_row_after from "./icons/add_row_after.svg";
import add_row_before from "./icons/add_row_before.svg";
import bold from "./icons/bold.svg";
import checklist from "./icons/checklist.svg";
import code from "./icons/code.svg";
import code_block from "./icons/code_block.svg";
import combine_cells from "./icons/combine_cells.svg";
import delete_col from "./icons/delete_col.svg";
import delete_row from "./icons/delete_row.svg";
import delete_table from "./icons/delete_table.svg";
import github from "./icons/github.svg";
import hr from "./icons/hr.svg";
import image from "./icons/image.svg";
import italic from "./icons/italic.svg";
import link from "./icons/link.svg";
import mention from "./icons/mention.svg";
import ol from "./icons/ol.svg";
import paragraph from "./icons/paragraph.svg";
import quote from "./icons/quote.svg";
import redo from "./icons/redo.svg";
import remove from "./icons/remove.svg";
import strike from "./icons/strike.svg";
import table from "./icons/table.svg";
import ul from "./icons/ul.svg";
import underline from "./icons/underline.svg";
import undo from "./icons/undo.svg";

const icons = {
	add_col_after,
	add_col_before,
	add_row_after,
	add_row_before,
	bold,
	checklist,
	code,
	code_block,
	combine_cells,
	delete_col,
	delete_row,
	delete_table,
	github,
	hr,
	image,
	italic,
	link,
	mention,
	ol,
	paragraph,
	quote,
	redo,
	remove,
	strike,
	table,
	ul,
	underline,
	undo
};

var comps = {};
for (let [key, value] of Object.entries(icons)) {
	comps["icon_" + key] = value;
}

export default {
	components: comps,
	props: {
		name: {},
		size: {
			default: "normal"
		},
		modifier: {
			default: null
		},
		fixAlign: {
			default: true
		}
	},
	computed: {
		icon: function () {
			return "icon_" + this.name;
		}
	}
};
</script>

<style lang="scss" scoped>
.icon {
	position: relative;
	display: inline-block;
	vertical-align: middle;
	width: 0.8rem;
	height: 0.8rem;
	margin: 0 0.3rem;
	top: 0.05rem;
	fill: currentColor;
	& svg {
		display: inline-block;
		vertical-align: top;
		width: 100%;
		height: 100%;
	}
	&:first-child {
		margin-left: 0;
	}
	&:last-child {
		margin-right: 0;
	}
}
</style>
