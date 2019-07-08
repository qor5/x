import Vue from 'vue';

export default Vue.extend({
	inject: ['core'],
	props: {
		fieldName: String,
	},
});
