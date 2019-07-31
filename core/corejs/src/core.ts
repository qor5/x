
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import Vue, { VueConstructor } from 'vue';

import {
	mergeStatesIntoForm,
	setPushState,
	EventData,
	setFormValue,
	jsonEvent,
} from './utils';

// Vue.config.productionTip = true;
declare var window: any;

interface EventFuncID {
	id: string;
	params?: string[];
	pushState?: any;
}

interface PortalUpdate {
	name: string;
	schema: string;
	afterLoaded?: string;
}

interface EventResponse {
	states?: any;
	schema?: any;
	data?: any;
	redirectURL?: string;
	reload: boolean;
	reloadPortals?: string[];
	updatePortals?: PortalUpdate[];
}

export class Core {
	public debounce = debounce;

	private debounceFetchEventThenRefresh = debounce(this.fetchEventThenRefresh, 800);
	private form: FormData;
	private rootChangeCurrent: any;
	private changeCurrent: any;


	constructor(form: FormData, rootChangeCurrent: any, changeCurrent: any) {
		this.form = form;
		this.rootChangeCurrent = rootChangeCurrent;
		this.changeCurrent = changeCurrent;
	}

	public fetchEvent(
		eventFuncId: EventFuncID,
		event: EventData,
	): Promise<EventResponse> {
		const eventData = JSON.stringify({
			eventFuncId,
			event,
		});

		const search = setPushState(
			eventFuncId.pushState,
			window.location.search,
			window.location.pathname,
			window.history.pushState,
		);

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

	public componentByTemplate(template: string, afterLoaded?: any): VueConstructor {
		return Vue.extend({
			provide: { core: this },
			template: '<div>' + template + '</div>', // to make only one root.
			methods: this.newVueMethods(),
			mounted() {
				this.$nextTick(() => {
					if (afterLoaded) {
						afterLoaded(this);
					}
				});
			},
			data() {
				return {
					boolean1: false,
					boolean2: false,
					boolean3: false,
					boolean4: false,
					boolean5: false,
				};
			},
		});
	}

	public setFormValue(fieldName: string, val: any) {
		setFormValue(this.form, fieldName, val);
	}

	// public getFormValue(fieldName: string): string {
	// 	return getFormValue(this.form, fieldName);
	// }

	// public getFormValueAsArray(fieldName: string): string[] {
	// 	return getFormValueAsArray(this.form, fieldName);
	// }

	private fetchEventThenRefresh(eventFuncId: EventFuncID, event: EventData) {
		this.fetchEvent(eventFuncId, event)
			.then((r: EventResponse) => {
				if (r.reloadPortals && r.reloadPortals.length > 0) {
					for (const portalName of r.reloadPortals) {
						const portal = window.branLazyPortals[portalName];
						if (portal) {
							portal.reload();
						}
					}
					return r;
				}

				if (r.updatePortals && r.updatePortals.length > 0) {
					for (const pu of r.updatePortals) {
						const portal = window.branLazyPortals[pu.name];
						if (portal) {
							let afterLoaded;
							if (pu.afterLoaded) {
								afterLoaded = new Function('comp', pu.afterLoaded);
							}
							portal.changeCurrent(this.componentByTemplate(pu.schema, afterLoaded));
						}
					}
					return r;
				}

				if (r.schema && r.reload) {
					this.rootChangeCurrent(this.componentByTemplate(r.schema));
					return r;
				}

				if (r.schema) {
					this.changeCurrent(this.componentByTemplate(r.schema));
					return r;
				}

				return r;
			});
	}

	private newVueMethods(): any {
		const self = this;
		return {
			onclick(eventFuncId: EventFuncID, evt: any) {
				self.fetchEventThenRefresh(eventFuncId, jsonEvent(evt));
			},
			oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
				self.controlsOnInput(eventFuncId, fieldName, evt);
			},
		};
	}

	private controlsOnInput(
		eventFuncId?: EventFuncID,
		fieldName?: string,
		evt?: any,
	) {
		if (fieldName) {
			this.form.set(fieldName, evt.target.value);
		}
		if (eventFuncId) {
			this.debounceFetchEventThenRefresh(eventFuncId, jsonEvent(evt));
		}
	}

}
