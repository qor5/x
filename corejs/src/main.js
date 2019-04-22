import Vue from "vue";
import { newFormWithStates } from "./form";
// import debounce from "lodash/debounce";
// import wrap from "lodash/wrap";

import { mergeStatesIntoForm } from "./form.js";
import "whatwg-fetch";
import querystring from "query-string";

Vue.config.productionTip = false;

// const dialogElement = document.createElement("div");

function closeDialog() {}

export function fetchEvent(eventFuncId, eventJSON, form, alert) {
	var eventData = JSON.stringify({
		eventFuncId: eventFuncId,
		event: eventJSON
	});

	var search = window.location.search;
	var pstate = eventFuncId.pushState;
	if (pstate) {
		var newSearch = "";
		if (Object.keys(pstate).length > 0) {
			var orig = querystring.parse(window.location.search);
			newSearch = querystring.stringify({ ...orig, ...pstate });
			if (newSearch.length > 0) {
				newSearch = `?${newSearch}`;
			}
		}
		window.history.pushState(
			pstate,
			"",
			window.location.pathname + newSearch
		);
		search = newSearch;
	}

	form.set("__event_data__", eventData);
	return fetch("__execute_event__/" + eventFuncId.id + search, {
		method: "POST",
		// headers: {
		// 	'Content-Type': 'multipart/form-data'
		// },
		body: form
	})
		.then(response => {
			return response.json();
		})
		.then(r => {
			var alertServ = alert || window.__uibuilderAlert;
			if (r.alert && alertServ) {
				alertServ.addAlert(
					r.alert.message,
					r.alert.type,
					r.alert.timeout
				);
			}

			if (r.states) {
				mergeStatesIntoForm(form, r.states);
			}

			// if (r.dialog) {

			// }

			if (r.closeDialog) {
				closeDialog();
			}

			if (r.redirectURL) {
				window.location.replace(r.redirectURL);
			}

			if (r.schema) {
				var body = document.querySelector("#app");
				body.innerHTML = r.schema;
				var styles = document.querySelector("#main_styles");
				styles.text = r.styles;
				var scripts = document.querySelector("#main_scripts");
				scripts.text = r.text

				newVue();
			}
			return r;
		});
}

export function jsonEvent(evt) {
	var v = {};

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

	if (typeof evt == "string" || typeof evt == "number") {
		v.value = evt.toString(); // For Radio, Pager
	}

	return v;
}

function newVue() {
	new Vue({
		el: "#app",
		data: {
			form: null
		},
		created: function() {
			this.form = newFormWithStates(window.__serverSideData__.states);
		},
		methods: {
			fetchEvent1: function(eventFuncId, event) {
				return fetchEvent(
					eventFuncId,
					jsonEvent(event),
					this.form,
					null
				);
			},
			click: function(eventFuncId, evt) {
				this.fetchEvent1(eventFuncId, evt);
			},
			change: function(eventFuncId, fieldName, evt) {
				this.fetchEvent1(eventFuncId, evt);
			}
		}
	});
}

newVue();
