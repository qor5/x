import * as constants from './Constants'

function pushKeyVal(segs: any, key: string, mod: string, val: any) {
	const modWithDot = mod ? `.${mod}` : '';
	segs.push([`${key}${modWithDot}`, val.toString()]);
}

function pushDatetimeRangeItem(segs: any, op: any) {
	// Now we only have 'between' modifier, but consider extendability, we keep the modifier system for now.
	const mod = op.modifier || constants.ModifierBetween;

	if (mod === constants.ModifierBetween) {
		if (op.valueFrom) {
			pushKeyVal(
				segs,
				op.key,
				'gte',
				op.valueFrom,
			);
		}
		if (op.valueTo) {
			pushKeyVal(segs, op.key, 'lt', op.valueTo);
		}
		return;
	}
}

function pushDateRangeItem(segs: any, op: any) {
	const mod = op.modifier || constants.ModifierBetween;

	if (mod === constants.ModifierBetween) {
		if (op.valueFrom) {
			pushKeyVal(
				segs,
				op.key,
				'gte',
				op.valueFrom,
			);
		}
		if (op.valueTo) {
			pushKeyVal(segs, op.key, 'lte', op.valueTo);
		}
		return;
	}
}

function pushDateItem(segs: any, op: any) {
	if (!op.valueIs) {
		return
	}

	pushKeyVal(segs, op.key, '', op.valueIs);
}

function pushNumberItem(segs: any, op: any) {
	const mod = op.modifier || 'equals';

	if (mod === 'equals') {
		const floatValue = parseFloat(op.valueIs);
		if (!isNaN(floatValue)) {
			pushKeyVal(segs, op.key, '', floatValue);
		}
		return;
	}

	if (mod === 'between') {
		const floatFrom = parseFloat(op.valueFrom);
		const floatTo = parseFloat(op.valueTo);
		if (!isNaN(floatFrom)) {
			pushKeyVal(segs, op.key, 'gte', floatFrom);
		}
		if (!isNaN(floatTo)) {
			pushKeyVal(segs, op.key, 'lte', floatTo);
		}
		return;
	}

	if (mod === 'greaterThan') {
		const floatValue = parseFloat(op.valueIs);
		if (!isNaN(floatValue)) {
			pushKeyVal(segs, op.key, 'gt', floatValue);
		}
		return;
	}

	if (mod === 'lessThan') {
		const floatValue = parseFloat(op.valueIs);
		if (!isNaN(floatValue)) {
			pushKeyVal(segs, op.key, 'lt', floatValue);
		}
		return;
	}
}

function pushStringItem(segs: any, op: any) {
	const mod = op.modifier || 'equals';
	if (mod === 'equals' && op.valueIs) {
		pushKeyVal(segs, op.key, '', op.valueIs);
		return;
	}

	if (mod === 'contains' && op.valueIs) {
		pushKeyVal(segs, op.key, 'ilike', op.valueIs);
		return;
	}
}

function pushSelectItem(segs: any, op: any) {
	const mod = op.modifier || 'equals';
	if (mod === 'equals' && op.valueIs) {
		pushKeyVal(segs, op.key, '', op.valueIs);
		return;
	}
}

function pushMultipleSelectItem(segs: any, op: any) {
	const mod = op.modifier || 'in';
	if (mod === 'in' && op.valuesAre && op.valuesAre.length > 0) {
		pushKeyVal(segs, op.key, 'in', op.valuesAre);
		return;
	}
	if (mod === 'notIn' && op.valuesAre && op.valuesAre.length > 0) {
		pushKeyVal(segs, op.key, 'notIn', op.valuesAre);
		return;
	}
}

function pushLinkageSelectItem(segs: any, op: any) {
	const mod = op.modifier || 'equals';
	if (mod === 'equals' && op.valuesAre && op.valuesAre.length > 0) {
		pushKeyVal(segs, op.key, '', op.valuesAre);
		return;
	}
}

export function filterData(data: any): any {
	if (!data) {
		return [];
	}

	const r: any = [];
	data.filter((op: any) => op.selected)
		.map((op: any) => {
			if (op.itemType === 'DatetimeRangeItem') {
				pushDatetimeRangeItem(r, op);
			}
			if (op.itemType === 'DateRangeItem') {
				pushDateRangeItem(r, op);
			}
			if (op.itemType === 'DateItem') {
				pushDateItem(r, op);
			}
			if (op.itemType === 'NumberItem') {
				pushNumberItem(r, op);
			}
			if (op.itemType === 'StringItem') {
				pushStringItem(r, op);
			}
			if (op.itemType === 'SelectItem') {
				pushSelectItem(r, op);
			}
			if (op.itemType === 'MultipleSelectItem') {
				pushMultipleSelectItem(r, op);
			}
			if (op.itemType === 'LinkageSelectItem') {
				pushLinkageSelectItem(r, op);
			}
			return op;
		});
	return r;
}

export function encodeFilterData(data: any) {
	return filterData(data).map((e: any) => `${encodeURIComponent(e[0])}=${encodeURIComponent(e[1])}`).join('&');
}
