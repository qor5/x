<script>
export default {
	name: 'vx-message-listener',
	props: {
		listenFunc: {
			type: Function,
		},
	},
	mounted() {
		if (!window.vxMessageListener){
			window.addEventListener("message", this.listenFunc,false);
			window.vxMessageListener = true;
		}
	},
};
</script>
