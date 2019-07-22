
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import querystring from 'query-string';
import { mergeStatesIntoForm } from './form';

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
	data?: any;
	redirectURL?: string;
	reload: boolean;
}

declare var window: any;

export class Core {
	public debounce = debounce;

	private debounceFetchEventThenDo = debounce(this.fetchEventThenDo, 800);
	private form: FormData;
	private rootChangeCurrent: any;
	private changeCurrent: any;


	constructor(form: FormData, rootChangeCurrent: any, changeCurrent: any) {
		this.form = form;
		this.rootChangeCurrent = rootChangeCurrent;
		this.changeCurrent = changeCurrent;
	}

	public newMethods(): any {
		const self = this;
		return {
			onclick(eventFuncId: EventFuncID, evt: any) {
				self.fetchEventThenDo(eventFuncId, self.jsonEvent(evt));
			},
			oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
				self.controlsOnInput(eventFuncId, fieldName, evt);
			},
		};
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

	public fetchEventThenDo(eventFuncId: EventFuncID, event: EventData) {
		this.fetchEvent(eventFuncId, event)
			.then((r: EventResponse) => {
				if (r.schema && r.reload) {
					this.rootChangeCurrent(this.componentByTemplate(r.schema));
				} else if (r.schema) {
					this.changeCurrent(this.componentByTemplate(r.schema));
				}
				return r;
			});
	}

	public componentByTemplate(template: string, afterLoaded?: () => void): any {
		return {
			provide: { core: this },
			template: '<div>' + template + '</div>', // to make only one root.
			methods: this.newMethods(),
			mounted() {
				this.$nextTick(() => {
					if (afterLoaded) {
						afterLoaded();
					}
				});
			},
		};
	}

	public setFormValue(fieldName: string, val: any) {
		if (!fieldName || fieldName.length === 0) {
			return;
		}
		this.form.delete(fieldName);
		if (!val) {
			return;
		}
		// console.log('val', val, 'Array.isArray(val)', Array.isArray(val));
		if (Array.isArray(val)) {
			val.forEach((v) => {
				this.form.append(fieldName, v);
			});
			return;
		}
		this.form.set(fieldName, val);
	}

	public getFormValue(fieldName: string): string {
		const val = this.form.get(fieldName);
		if (typeof val === 'string') {
			return val;
		}
		return '';
	}

	public getFormValueAsArray(fieldName: string): string[] {
		const vals = this.form.getAll(fieldName);
		const r: string[] = [];
		for (const v of vals) {
			if (typeof v === 'string') {
				r.push(v);
			}
		}
		return r;
	}


	private controlsOnInput(
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
			this.debounceFetchEventThenDo(eventFuncId, this.jsonEvent(evt));
		}
	}

}
