import 'formdata-polyfill';
import querystring from 'query-string';

export function newFormWithStates(states: any): FormData {
	const f = new FormData();
	if (!states) {
		return f;
	}
	mergeStatesIntoForm(f, states);
	return f;
}

export function mergeStatesIntoForm(form: FormData, states: any) {
	if (!states) {
		return;
	}
	for (const k of Object.keys(states)) {
		form.delete(k);
		for (const v of states[k]) {
			form.append(k, v);
		}
	}
}

interface StatePusher {
	pushState(data: any, title: string, url?: string | null): void;
}

export function setPushState(
	eventFuncId: any,
	url: string,
	pusher: StatePusher,
	popstate: boolean | undefined,
): any {
	let pstate = eventFuncId.pushState;

	if (typeof pstate === 'string') {
		pstate = querystring.parse(pstate);
	}

	const orig = querystring.parseUrl(url);
	let eventQuery = { __execute_event__: eventFuncId.id };

	if (pstate) {

		Object.keys(pstate).forEach((key) => {
			const v = pstate[key];
			if (!Array.isArray(v)) {
				pstate[key] = [v];
			}
		});

		let pushStateQuery = '';
		if (Object.keys(pstate).length > 0) {
			pstate = { ...orig.query, ...pstate };
			pushStateQuery = querystring.stringify(pstate);
			eventQuery = { ...eventQuery, ...pstate };
			if (pushStateQuery.length > 0) {
				pushStateQuery = `?${pushStateQuery}`;
			}
		}

		if (popstate !== true) {
			const newUrl = orig.url + pushStateQuery;
			const pushedState = { ...pstate, ...{ url: newUrl } };
			pusher.pushState(
				pushedState,
				'',
				newUrl,
			);
		}
	}

	eventFuncId.pushState = pstate;

	return {
		newEventFuncId: eventFuncId,
		eventURL: `${orig.url}?${querystring.stringify(eventQuery)}`,
	};
}


export interface EventData {
	value?: string;
	checked?: boolean;
}

export function jsonEvent(evt: any) {
	const v: EventData = {};

	if (evt && evt.target) {
		// For Checkbox
		if (evt.target.checked) {
			v.checked = evt.target.checked;
		}

		// For Input
		if (evt.target.value !== undefined) {
			v.value = evt.target.value;
		}
		return v;
	}

	// For List
	if (evt.key) {
		v.value = evt.key;
		return v;
	}

	if (typeof evt === 'string' || typeof evt === 'number') {
		v.value = evt.toString(); // For Radio, Pager
	}

	return v;
}


export function setFormValue(form: FormData, fieldName: string, val: any) {
	if (!fieldName || fieldName.length === 0) {
		return;
	}
	form.delete(fieldName);
	if (!val) {
		return;
	}
	// console.log('val', val, 'Array.isArray(val)', Array.isArray(val));
	if (Array.isArray(val)) {
		val.forEach((v) => {
			form.append(fieldName, v);
		});
		return;
	}
	form.set(fieldName, val);
}

// export function getFormValue(form: FormData, fieldName: string): string {
// 	const val = form.get(fieldName);
// 	if (typeof val === 'string') {
// 		return val;
// 	}
// 	return '';
// }

// export function getFormValueAsArray(form: FormData, fieldName: string): string[] {
// 	const vals = form.getAll(fieldName);
// 	const r: string[] = [];
// 	for (const v of vals) {
// 		if (typeof v === 'string') {
// 			r.push(v);
// 		}
// 	}
// 	return r;
// }
