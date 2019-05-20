import { newFormWithStates, mergeStatesIntoForm } from './form';
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import querystring from 'query-string';
import Vue from 'vue';

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
const app = document.getElementById('app');
if (!app) {
	throw new Error('#app required');
}

function fetchEvent(
	eventFuncId: EventFuncID,
	event: EventData,
): Promise<EventResponse> {
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
	}).then((r) => {
		return r.json();
	}).then((r: EventResponse) => {
		if (r.states) {
			mergeStatesIntoForm(form, r.states);
		}

		if (r.redirectURL) {
			window.location.replace(r.redirectURL);
		}
		return r;
	});
}


const debounceFetchEvent = debounce(fetchEventAndProcessDefault, 800);

function controlsOnInput(comp: any, eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
	// console.log("evt", evt)
	if (fieldName) {
		form.set(fieldName, evt.target.value);
	}
	if (eventFuncId) {
		debounceFetchEvent(comp, eventFuncId, jsonEvent(evt));
	}
}

const methods = {
	onclick(eventFuncId: EventFuncID, evt: any) {
		fetchEventAndProcessDefault(this, eventFuncId, jsonEvent(evt));
	},
	oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
		controlsOnInput(this, eventFuncId, fieldName, evt);
	},
};

function fetchEventAndProcessDefault(comp: any, eventFuncId: EventFuncID, event: EventData) {
	fetchEvent(eventFuncId, event)
		.then((r: EventResponse) => {
			if (r.schema) {
				reload(comp, r);
			}
			return r;
		});
}

function componentByTemplate(template: string): any {
	return {
		template: '<div>' + template + '</div>', // to make only one root.
		methods,
	};
}

function reload(comp: any, r: EventResponse) {
	// app.innerHTML = r.schema;
	// if (r.styles) {
	// 	let style = document.querySelector('#main_styles');
	// 	if (style && style.parentNode) {
	// 		style.parentNode.removeChild(style);
	// 	}
	// 	style = document.createElement('style');
	// 	style.setAttribute('type', 'text/css');
	// 	style.setAttribute('id', 'main_styles');
	// 	style.appendChild(document.createTextNode(r.styles));
	// 	document.body.insertBefore(style, app);
	// }

	// if (r.scripts) {
	// 	let script = document.querySelector('#main_scripts');
	// 	if (script && script.parentNode) {
	// 		script.parentNode.removeChild(script);
	// 	}
	// 	script = document.createElement('script');
	// 	script.setAttribute('id', 'main_scripts');
	// 	script.appendChild(document.createTextNode(r.scripts));
	// 	document.body.insertBefore(script, app.nextSibling);
	// }
	comp.$root.changeCurrent(componentByTemplate(r.schema));
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


for (const registerComp of (window.__branVueComponentRegisters || [])) {
	registerComp(Vue);
}

Vue.component('bran-lazy-loader', {
	name: 'BranLazyLoader',
	props: ['loaderFunc', 'visible'],
	template: `
		<div class="bran-lazy-loader" v-if="visible">
			<component :is="lazyloader"></component>
		</div>
	`,
	data() {
		const ef = this.loaderFunc;
		if (!ef) {
			return {
				lazyloader: {
					render() {
						return null;
					},
				},
			};
		}
		return {
			lazyloader(): any {
				return fetchEvent(ef, {})
					.then((r) => {
						return componentByTemplate(r.schema);
					});
			},
		};
	},
});

const vm = new Vue({
	data: {
		current: componentByTemplate(app.innerHTML),
	},
	template: `
	<div id="app">
		<component :is="current"></component>
	</div>
`,
	methods: {
		changeCurrent(newView: any) {
			this.current = newView;
		},
	},
});

vm.$mount('#app');
