<template>
  <div>
    <v-dialog v-model="dialogVisible" :width="dialogWidth">
      <template v-slot:activator="{ isActive: isActive, props: activatorProps }">
        <v-text-field
          v-bind="activatorProps"
          :disabled="disabled"
          :error-messages="errorMessages"
          :loading="loading"
          :label="label"
          v-model="formattedDatetime"
          :hide-details="hideDetails"
          color="primary"
          variant="underlined"
          readonly
        >
          <template v-slot:prepend>
            <v-icon
              icon="mdi-calendar-edit"
              :color="isActive ? 'primary' : ''"
              size="x-large"
            ></v-icon>
          </template>
          <template v-slot:loader>
            <v-progress-linear
              color="primary"
              indeterminate
              absolute
              height="2"
            ></v-progress-linear>
          </template>
        </v-text-field>
      </template>

      <template v-slot:default="{ isActive }">
        <v-card>
          <v-card-text class="pa-0">
            <v-container class="d-flex justify-space-between align-center ga-4 py-0">
              <v-date-picker
                v-model="dateOfPicker"
                full-width
                no-title
                v-bind="datePickerProps"
                @update:year="onYearOrMonthChange($event, 'year')"
                @update:month="onYearOrMonthChange($event, 'month')"
                @update:modelValue="onYearOrMonthChange($event, 'modelValue')"
              ></v-date-picker>
              <input type="time" class="text-h2 timer" v-model="timeOfPicker" />
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              v-if="clearText"
              color="grey lighten-1"
              variant="text"
              @click.native="clearHandler(isActive)"
              >{{ clearText }}
            </v-btn>
            <v-tooltip :text="okTips" :disabled="!valueChangedWithoutSaved">
              <template v-slot:activator="{ props }">
                <span v-bind="props">
                  <v-btn
                    color="green darken-1"
                    variant="text"
                    @click="okHandler(isActive)"
                    :disabled="valueChangedWithoutSaved"
                    >{{ okText }}
                  </v-btn>
                </span>
              </template>
            </v-tooltip>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>
<script lang="ts" setup>
import { format, parse } from 'date-fns'
import dayjs from 'dayjs'
import { computed, nextTick, onMounted, Ref, ref, watch } from 'vue'
import { useLocale } from 'vuetify'
const DEFAULT_TIME = '00:00:00'
const DEFAULT_DATE_FORMAT = 'yyyy-MM-dd'
const DEFAULT_TIME_FORMAT = 'HH:mm:ss'
const emit = defineEmits(['update:modelValue', 'input'])
const { t } = useLocale()

const okTips = t('$vuetify.datePicker.okTips')
const props = defineProps({
  modelValue: {
    type: String
  },
  disabled: {
    type: Boolean
  },
  errorMessages: {
    type: Array<string>
  },
  loading: {
    type: Boolean
  },
  label: {
    type: String,
    default: ''
  },
  dialogWidth: {
    type: String,
    default: 'auto'
  },
  dateFormat: {
    type: String,
    default: 'yyyy-MM-dd'
  },
  timeFormat: {
    type: String,
    default: 'HH:mm'
  },
  clearText: {
    type: String,
    default: 'CLEAR'
  },
  okText: {
    type: String,
    default: 'OK'
  },
  textFieldProps: {
    type: Object
  },
  datePickerProps: {
    type: Object
  },
  timePickerProps: {
    type: Object
  },
  hideDetails: {
    type: Boolean
  }
})
const dialogVisible = ref(false)
const date = ref()
const dateOfPicker = ref()
const time = ref(DEFAULT_TIME)
const timer = ref()
const timeOfPicker = ref()
const tempYearAndMonth = ref(['-1', '-1'])

const dateTimeFormat = computed(() => {
  return props.dateFormat + ' ' + props.timeFormat
})
const defaultDateTimeFormat = computed(() => {
  return DEFAULT_DATE_FORMAT + ' ' + DEFAULT_TIME_FORMAT
})
const selectedDatetime = computed(() => {
  if (date.value && time.value) {
    let datetimeString = format(date.value, DEFAULT_DATE_FORMAT) + ' ' + time.value
    if (time.value.length === 5) {
      datetimeString += ':00'
    }
    return parse(datetimeString, defaultDateTimeFormat.value, new Date())
  } else {
    return null
  }
})
const formattedDatetime = computed(() => {
  return selectedDatetime.value ? format(<Date>selectedDatetime.value, dateTimeFormat.value) : ''
})
const dateSelected = () => {
  return !date.value
}

const onYearOrMonthChange = (value: number, type: 'year' | 'month' | 'modelValue') => {
  if (type === 'modelValue') {
    tempYearAndMonth.value = ['-1', '-1']
    console.log('modelValue', value)
    return
  }

  const map = {
    year: 0,
    month: 1
  }
  const curYearAndMonth = getYearAndMonthStr(date.value)
  tempYearAndMonth.value[map[type]] = value.toString()
}

// must choose a date when year or month changed
const valueChangedWithoutSaved = computed(() => {
  if (tempYearAndMonth.value.join('-') === '-1--1') return false

  const curYearAndMonth = getYearAndMonthStr(date.value)

  return curYearAndMonth !== Object.values(tempYearAndMonth.value).join('-')
})

const getYearAndMonthStr = (dateTime: number) => {
  return dayjs(date.value).format('YYYY-MM')
}

const init = () => {
  if (!props.modelValue) {
    return
  }
  // see https://stackoverflow.com/a/9436948
  let initDateTime = parse(props.modelValue, dateTimeFormat.value, new Date())
  date.value = initDateTime
  time.value = format(initDateTime, DEFAULT_TIME_FORMAT)
  dateOfPicker.value = date.value
  timeOfPicker.value = time.value

  tempYearAndMonth.value = getYearAndMonthStr(date.value).split('-')
}

watch(dialogVisible, (newVal) => {
  if (newVal) {
    dateOfPicker.value = date.value
    timeOfPicker.value = time.value
    tempYearAndMonth.value = getYearAndMonthStr(date.value).split('-')
  }
})

const okHandler = (isActive: Ref) => {
  date.value = dateOfPicker.value
  resetPicker(isActive)
  if (!date.value) {
    date.value = new Date()
  }
  emit('update:modelValue', formattedDatetime.value)
}
const clearHandler = (isActive: Ref) => {
  resetPicker(isActive)
  date.value = null
  emit('update:modelValue', null)
}

const resetPicker = (isActive: Ref) => {
  time.value = timeOfPicker.value
  isActive.value = false
  tempYearAndMonth.value = ['-1', '-1']
  if (timer.value) {
    timer.value.selectingHour = true
  }
}
onMounted(() => {
  nextTick(() => {
    init()
  })
})
</script>

<style scoped></style>
