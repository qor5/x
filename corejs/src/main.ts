import { newFormWithStates, mergeStatesIntoForm } from './form';
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import querystring from 'query-string';
import Vue, { VNode, VNodeDirective } from 'vue';

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
			search = newSearch;
			if (newSearch.length > 0) {
				search = `&${newSearch}`;
				newSearch = `?${newSearch}`;
			}
		}
		window.history.pushState(
			pstate,
			'',
			window.location.pathname + newSearch,
		);
	}

	form.set('__event_data__', eventData);
	return fetch(`?__execute_event__=${eventFuncId.id}${search}`, {
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

function controlsOnInput(
	comp: any,
	eventFuncId?: EventFuncID,
	fieldName?: string,
	evt?: any,
	toFormFunc?: any,
) {
	console.log("comp", comp, "fieldName", fieldName, "event", evt)
	// console.log("root", comp.$root)
	// console.log("comp", comp.$el)
	// console.log("comp.$props", comp.$props)
	// console.log("comp.$options", comp.$options)
	console.log("toFormFunc", toFormFunc)

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
	oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any, toFormFunc?: any) {
		controlsOnInput(this, eventFuncId, fieldName, evt, toFormFunc);
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

function componentByTemplate(template: string, afterLoaded?: () => void): any {
	return {
		template: '<div>' + template + '</div>', // to make only one root.
		methods,
		mounted() {
			this.$nextTick(() => {
				if (afterLoaded) {
					afterLoaded();
				}
			});
		},
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

Vue.component('BranLazyLoader', {
	name: 'BranLazyLoader',
	props: ['loaderFunc', 'visible', 'afterLoaded'],
	template: `
		<div class="bran-lazy-loader" v-if="visible">
			<component :is="lazyloader"></component>
		</div>
	`,
	data() {
		const ef = this.loaderFunc;
		const afterLoaded = this.afterLoaded;
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
						return componentByTemplate(r.schema, afterLoaded);
					});
			},
		};
	},
});

// Vue.mixin({
// 	props: ['bran-field-name'],
// 	created: function () {
// 		// var myOption = this.$options
// 		// console.log("props", JSON.stringify(this.$props))
// 		console.log("this.$el", this.$el, this)
// 	}
// })

Vue.directive('bran', {
	// When the bound element is inserted into the DOM...
	inserted: function (el: HTMLElement, binding: VNodeDirective, vnode: VNode) {
		console.log("el", el)
		console.log("binding", binding)
		console.log("vnode", vnode)
		console.log("vnode.componentInstance", vnode.componentInstance)
		console.log("vnode.context", vnode.context)
		console.log("vnode.data", vnode.data!)
		console.log("vnode.directives", vnode.data!.directives)
		if (vnode.componentInstance) {
			vnode.componentInstance.$on("change", (v: any) => {
				console.log("change value", v, "fieldName", binding.value.FieldName)
			})
		}
	}
})


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
