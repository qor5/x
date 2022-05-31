import Vue, {CreateElement, VNode} from 'vue';

export const Core = Vue.extend({
	props: {
		fieldName: String,
		loadPageWithArrayOp: Boolean,
	},
});

export const SelectedItems = Vue.extend({
	props: {
		selectedItems: {
			type: Array,
			default: () => [],
		} as any,
		multiple: Boolean,
	},
});


interface Slots { [key: string]: VNode[] | undefined; }

export const slotTemplates = (h: CreateElement, slots: Slots): VNode[] => {
	const templates: VNode[] = [];

	for (const name in slots) {
		if (!Object.getOwnPropertyDescriptor(slots, name)) {
			continue;
		}
		templates.push(
			<template slot={name}>
				{slots[name]}
			</template>,
		);
	}
	return templates;
};


// export const selectValue = (core: any, props: Record<string, any>): any => {
// 	const {
// 		selectedItems,
// 		multiple,
// 		fieldName,
// 	} = props;

// 	if (multiple) {
// 		const formValues: string[] = core.getFormValueAsArray(fieldName);
// 		return selectedItems || formValues;
// 	}

// 	let value = core.getFormValue(fieldName);
// 	if (selectedItems && selectedItems.length) {
// 		value = selectedItems[0];
// 	}
// 	return value;
// };
