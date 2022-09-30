<template>
	<span>{{datetime}}</span>
</template>

<script>
import {format} from 'date-fns'

const DEFAULT_FORMAT = 'yyyy-MM-dd HH:mm:ss'

export default {
	name: 'v-datetime-formatter',
	props: {
		// unix timestamp
		value: {
			type: Number,
			default: null
		},
		format: {
			type: String,
			default: DEFAULT_FORMAT
		}
	},
	data() {
		return {}
	},
	computed: {
		datetime() {
			if (!this.value || this.value <= 0) {
				return ''
			}
			return format(new Date(this.value * 1000), this.format)
		},
	}
}
</script>
