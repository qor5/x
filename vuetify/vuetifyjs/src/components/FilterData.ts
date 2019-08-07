import moment from 'moment';

function pushKeyVal(segs: any, key: string, mod: string, val: any) {
	const modWithDot = mod ? `.${mod}` : '';
	segs.push([`${key}${modWithDot}`, val.toString()]);
}

function convertUTC(op: any, date: any) {
	const s = moment(date).format('YYYY-MM-DD');
	if (op.timezone === 'utc') {
		return moment.utc(s);
	}
	return moment(s);
}

function pushDateItem(segs: any, op: any) {
	const mod = op.modifier || 'inTheLast';

	if (mod === 'inTheLast') {
		let m = moment().startOf('day');
		if (op.timezone === 'utc') {
			m = m.utc().startOf('day');
		}

		pushKeyVal(segs, op.key, 'lt', m.add(1, 'days').unix());
		const unit = op.inTheLastUnit || 'days';

		pushKeyVal(
			segs,
			op.key,
			'gte',
			m.subtract(parseInt(op.inTheLastValue || '1', 10), unit).unix(),
		);
		return;
	}

	if (mod === 'equals' && op.valueIs) {
		pushKeyVal(segs, op.key, '', convertUTC(op, op.valueIs).unix());
		return;
	}

	if (mod === 'between') {
		if (op.valueFrom) {
			pushKeyVal(
				segs,
				op.key,
				'gte',
				convertUTC(op, op.valueFrom).unix(),
			);
		}
		if (op.valueTo) {
			pushKeyVal(segs, op.key, 'lt', convertUTC(op, op.valueTo).add(1, 'days').unix());
		}
		return;
	}

	if (mod === 'isAfter' && op.valueIs) {
		pushKeyVal(
			segs,
			op.key,
			'gt',
			convertUTC(op, op.valueIs)
				.add(1, 'days')
				.subtract(1, 'seconds')
				.unix(),
		);
		return;
	}

	if (mod === 'isAfterOrOn') {
		pushKeyVal(segs, op.key, 'gte', convertUTC(op, op.valueIs).unix());
		return;
	}

	if (mod === 'isBefore') {
		pushKeyVal(segs, op.key, 'lt', convertUTC(op, op.valueIs).unix());
		return;
	}

	if (mod === 'isBeforeOrOn') {
		pushKeyVal(
			segs,
			op.key,
			'lte',
			convertUTC(op, op.valueIs)
				.add(1, 'days')
				.subtract(1, 'seconds')
				.unix(),
		);
		return;
	}
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

export function filterData(data: any): any {
	if (!data) {
		return [];
	}

	const r: any = [];
	data.filter((op: any) => op.selected)
		.map((op: any) => {
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
			return op;
		});
	return r;
}

export function encodeFilterData(data: any) {
	return filterData(data).map((e: any) => `${encodeURIComponent(e[0])}=${encodeURIComponent(e[1])}`).join('&');
}
