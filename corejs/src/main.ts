import Vue from 'vue';
import { newFormWithStates, mergeStatesIntoForm } from './form';
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import querystring from 'query-string';

Vue.config.productionTip = true;

interface EventFuncID {
	id: string;
	params?: string[];
	pushState?: any;
}

interface EventData {
	value?: string;
	checked?: boolean;
}

interface EventResponse {
	states?: any;
	schema?: any;
	redirectURL?: string;
	styles?: string;
	scripts?: string;
}

declare var window: any;

const ssd = window.__serverSideData__;
const states = (ssd && ssd.states) || {};
export const form = newFormWithStates(states);

export function fetchEvent(
	eventFuncId: EventFuncID,
	event: EventData,
): Promise<Response> {
	const eventData = JSON.stringify({
		eventFuncId,
		event,
	});

	let search = window.location.search;
	const pstate = eventFuncId.pushState;
	if (pstate) {
		let newSearch = '';
		if (Object.keys(pstate).length > 0) {
			const orig = querystring.parse(window.location.search);
			newSearch = querystring.stringify({ ...orig, ...pstate });
			if (newSearch.length > 0) {
				newSearch = `?${newSearch}`;
			}
		}
		window.history.pushState(
			pstate,
			'',
			window.location.pathname + newSearch,
		);
		search = newSearch;
	}

	form.set('__event_data__', eventData);
	return fetch('__execute_event__/' + eventFuncId.id + search, {
		method: 'POST',
		// headers: {
		// 	'Content-Type': 'multipart/form-data'
		// },
		body: form,
	});
}

function fetchEventAndProcessDefault(eventFuncId: EventFuncID, event: EventData) {

	fetchEvent(eventFuncId, event)
		.then((r) => {
			return r.json();
		})
		.then((r: EventResponse) => {
			if (r.states) {
				mergeStatesIntoForm(form, r.states);
			}

			if (r.redirectURL) {
				window.location.replace(r.redirectURL);
			}

			if (r.schema) {
				reload(r);
			}
			return r;
		});
}

function reload(r: EventResponse) {
	const app = document.querySelector('#app');
	if (!app) {
		return;
	}

	app.innerHTML = r.schema;
	if (r.styles) {
		let style = document.querySelector('#main_styles');
		if (style && style.parentNode) {
			style.parentNode.removeChild(style);
		}
		style = document.createElement('style');
		style.setAttribute('type', 'text/css');
		style.setAttribute('id', 'main_styles');
		style.appendChild(document.createTextNode(r.styles));
		document.body.insertBefore(style, app);
	}

	if (r.scripts) {
		let script = document.querySelector('#main_scripts');
		if (script && script.parentNode) {
			script.parentNode.removeChild(script);
		}
		script = document.createElement('script');
		script.setAttribute('id', 'main_scripts');
		script.appendChild(document.createTextNode(r.scripts));
		document.body.insertBefore(script, app.nextSibling);
	}

	newVue();
}

function jsonEvent(evt: any) {
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

const debounceFetchEvent = debounce(fetchEventAndProcessDefault, 800);

function controlsOnInput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
	// console.log("evt", evt)
	if (fieldName) {
		form.set(fieldName, evt.target.value);
	}
	if (eventFuncId) {
		debounceFetchEvent(eventFuncId, jsonEvent(evt));
	}
}

function newVue() {
	for (const registerComp of (window.__branVueComponentRegisters || [])) {
		registerComp(Vue);
	}
	const vm = new Vue({
		el: '#app',
		data: {},
		methods: {
			onclick(eventFuncId: EventFuncID, evt: any) {
				fetchEventAndProcessDefault(eventFuncId, jsonEvent(evt));
			},
			oninput: controlsOnInput,
		},
	});
	window.vueComps = [];

}

newVue();
