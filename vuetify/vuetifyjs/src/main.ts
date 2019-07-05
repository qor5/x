import Vue from 'vue';
import Vuetify from 'vuetify';
import './main.styl';

declare var window: any;

if (!window.Vue) {
	window.Vue = Vue;
}

(window.__branVueComponentRegisters = window.__branVueComponentRegisters || []).push((v: any) => {
	v.use(Vuetify);
});
