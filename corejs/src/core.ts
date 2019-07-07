
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import querystring from 'query-string';
import { newFormWithStates, mergeStatesIntoForm } from './form';
import { VNode, VNodeDirective } from 'vue';


// Vue.config.productionTip = true;

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

interface SetupFuncParams {
	el: HTMLElement;
	binding: VNodeDirective;
	vnode: VNode;
	fieldName?: string;
}

interface SetupFuncs {
	[name: string]: (params: SetupFuncParams) => void;
}

declare var window: any;

export class Core {
	public form: FormData;
	public setupFuncs: SetupFuncs = {};

	private debounceFetchEvent = debounce(this.fetchEventThenReload, 800);
	private methods: any = {};

	constructor() {
		const ssd = window.__serverSideData__;
		const states = (ssd && ssd.states) || {};
		this.form = newFormWithStates(states);
		const self = this;

		this.methods = {
			onclick(eventFuncId: EventFuncID, evt: any) {
				self.fetchEventThenReload(this, eventFuncId, self.jsonEvent(evt));
			},
			oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
				self.controlsOnInput(this, eventFuncId, fieldName, evt);
			},
		};
	}

	public extendMethods(newMethods: any) {
		this.methods = { ...this.methods, ...newMethods };
	}

	public extendSetupFuncs(newFuncs: any) {
		this.setupFuncs = { ...this.setupFuncs, ...newFuncs };
	}

	public callSetupFunc(setupFunc: string, el: HTMLElement, binding: VNodeDirective, vnode: VNode) {
		const f = this.setupFuncs[setupFunc];
		if (!f) {
			return;
		}
		f({
			el,
			binding,
			vnode,
			fieldName: binding.value!.fieldName,
		});
	}

	public fetchEvent(
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

		this.form.set('__event_data__', eventData);
		return fetch(`?__execute_event__=${eventFuncId.id}${search}`, {
			method: 'POST',
			// headers: {
			// 	'Content-Type': 'multipart/form-data'
			// },
			body: this.form,
		}).then((r) => {
			return r.json();
		}).then((r: EventResponse) => {
			if (r.states) {
				mergeStatesIntoForm(this.form, r.states);
			}

			if (r.redirectURL) {
				window.location.replace(r.redirectURL);
			}
			return r;
		});
	}

	public jsonEvent(evt: any) {
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

	public fetchEventThenReload(comp: any, eventFuncId: EventFuncID, event: EventData) {
		this.fetchEvent(eventFuncId, event)
			.then((r: EventResponse) => {
				if (r.schema) {
					this.reload(comp, r);
				}
				return r;
			});
	}

	public componentByTemplate(template: string, afterLoaded?: () => void): any {
		return {
			template: '<div>' + template + '</div>', // to make only one root.
			methods: this.methods,
			mounted() {
				this.$nextTick(() => {
					if (afterLoaded) {
						afterLoaded();
					}
				});
			},
		};
	}


	private controlsOnInput(
		comp: any,
		eventFuncId?: EventFuncID,
		fieldName?: string,
		evt?: any,
	) {
		// console.log("comp", comp, "fieldName", fieldName, "event", evt)
		// // console.log("root", comp.$root)
		// // console.log("comp", comp.$el)
		// // console.log("comp.$props", comp.$props)
		// // console.log("comp.$options", comp.$options)
		// console.log("toFormFunc", toFormFunc)

		if (fieldName) {
			this.form.set(fieldName, evt.target.value);
		}
		if (eventFuncId) {
			this.debounceFetchEvent(comp, eventFuncId, this.jsonEvent(evt));
		}
	}

	private reload(comp: any, r: EventResponse) {
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
		comp.$root.changeCurrent(this.componentByTemplate(r.schema));
	}


}
