import Vue, { VNode, CreateElement } from 'vue';

export const Core = Vue.extend({
	inject: ['core'],
	props: {
		fieldName: String,
	},
});

export const SelectedItems = Vue.extend({
	props: {
		selectedItems: {
			type: Array,
			default: () => [],
		},
		multiple: Boolean,
	},
});


interface Slots { [key: string]: VNode[] | undefined; }

export const slotTemplates = (h: CreateElement, slots: Slots): VNode[] => {
	const templates: VNode[] = [];

	for (const name in slots) {
		if (!slots.hasOwnProperty(name)) {
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

