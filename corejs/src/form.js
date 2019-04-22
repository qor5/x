import "formdata-polyfill";

function newFormWithStates(states) {
	var f = new FormData();
	if (!states) {
		return f;
	}
	mergeStatesIntoForm(f, states);
	return f;
}

function mergeStatesIntoForm(form, states) {
	if (!states) {
		return;
	}
	for (let k in states) {
		form.delete(k);
		for (let v of states[k]) {
			form.append(k, v);
		}
	}
}

export { newFormWithStates, mergeStatesIntoForm };
