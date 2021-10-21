<template>
	<v-menu v-model="display" :max-width="dialogWidth" :close-on-content-click="false">
		<template v-slot:activator="{ on }">
			<v-text-field
				v-bind="textFieldProps"
				:disabled="disabled"
				:loading="loading"
				:label="label"
				:value="formattedDatetime"
				v-on="on"
				prepend-icon="edit_calendar"
				readonly
			>
				<template v-slot:progress>
					<slot name="progress">
						<v-progress-linear color="primary" indeterminate absolute height="2"></v-progress-linear>
					</slot>
				</template>
			</v-text-field>
		</template>

		<v-card>
			<v-card-text class="px-0 py-0">
				<v-container>
					<v-row>
						<v-col cols="6" class="pa-0">
							<v-date-picker
								v-model="date"
								v-bind="datePickerProps"
								full-width
								no-title
							></v-date-picker>
						</v-col>
						<v-col cols="6" class="pa-0">
							<v-time-picker
								ref="timer"
								class="v-time-picker-custom"
								v-model="time"
								v-bind="timePickerProps"
								full-width
								scrollable
							></v-time-picker>
						</v-col>
					</v-row>
				</v-container>
			</v-card-text>
			<v-card-actions>
				<v-spacer></v-spacer>
				<slot name="actions" :parent="this">
					<v-btn color="grey lighten-1" text @click.native="clearHandler">{{ clearText }}</v-btn>
					<v-btn color="green darken-1" text @click="okHandler">{{ okText }}</v-btn>
				</slot>
			</v-card-actions>
		</v-card>
	</v-menu>
</template>

<script>
import {format, parse} from 'date-fns'

const DEFAULT_DATE = ''
const DEFAULT_TIME = '00:00:00'
const DEFAULT_DATE_FORMAT = 'yyyy-MM-dd'
const DEFAULT_TIME_FORMAT = 'HH:mm:ss'
const DEFAULT_DIALOG_WIDTH = 580
const DEFAULT_CLEAR_TEXT = 'CLEAR'
const DEFAULT_OK_TEXT = 'OK'

export default {
	name: 'v-datetime-picker',
	model: {
		prop: 'datetime',
		event: 'input'
	},
	props: {
		datetime: {
			type: [Date, String],
			default: null
		},
		disabled: {
			type: Boolean
		},
		loading: {
			type: Boolean
		},
		label: {
			type: String,
			default: ''
		},
		dialogWidth: {
			type: Number,
			default: DEFAULT_DIALOG_WIDTH
		},
		dateFormat: {
			type: String,
			default: DEFAULT_DATE_FORMAT
		},
		timeFormat: {
			type: String,
			default: 'HH:mm'
		},
		clearText: {
			type: String,
			default: DEFAULT_CLEAR_TEXT
		},
		okText: {
			type: String,
			default: DEFAULT_OK_TEXT
		},
		textFieldProps: {
			type: Object
		},
		datePickerProps: {
			type: Object
		},
		timePickerProps: {
			type: Object
		}
	},
	data() {
		return {
			display: false,
			date: DEFAULT_DATE,
			time: DEFAULT_TIME
		}
	},
	mounted() {
		this.init()
	},
	computed: {
		dateTimeFormat() {
			return this.dateFormat + ' ' + this.timeFormat
		},
		defaultDateTimeFormat() {
			return DEFAULT_DATE_FORMAT + ' ' + DEFAULT_TIME_FORMAT
		},
		formattedDatetime() {
			return this.selectedDatetime ? format(this.selectedDatetime, this.dateTimeFormat) : ''
		},
		selectedDatetime() {
			if (this.date && this.time) {
				let datetimeString = this.date + ' ' + this.time
				if (this.time.length === 5) {
					datetimeString += ':00'
				}
				return parse(datetimeString, this.defaultDateTimeFormat, new Date())
			} else {
				return null
			}
		},
		dateSelected() {
			return !this.date
		}
	},
	methods: {
		init() {
			if (!this.datetime) {
				return
			}

			let initDateTime
			if (this.datetime instanceof Date) {
				initDateTime = this.datetime
			} else if (typeof this.datetime === 'string' || this.datetime instanceof String) {
				// see https://stackoverflow.com/a/9436948
				initDateTime = parse(this.datetime, this.dateTimeFormat, new Date())
			}

			this.date = format(initDateTime, DEFAULT_DATE_FORMAT)
			this.time = format(initDateTime, DEFAULT_TIME_FORMAT)
		},
		okHandler() {
			this.resetPicker()
			this.$emit('input', this.selectedDatetime)
		},
		clearHandler() {
			this.resetPicker()
			this.date = DEFAULT_DATE
			this.time = DEFAULT_TIME
			this.$emit('input', null)
		},
		resetPicker() {
			this.display = false
			if (this.$refs.timer) {
				this.$refs.timer.selectingHour = true
			}
		}
	},
	watch: {
		datetime: function () {
			this.init()
		}
	}
}
</script>

<style lang="scss" scoped>
.v-time-picker-title {
	height: 50px !important;
	color: black !important;
}
</style>
