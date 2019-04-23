import 'formdata-polyfill';

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
