<template>
  <div>
    <v-dialog :width="dialogWidth">
      <template v-slot:activator="{ isActive: isActive, props: activatorProps }">
        <v-text-field
          v-bind="activatorProps"
          :disabled="disabled"
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
          <v-card-text class="px-0 py-0">
            <v-container class="d-flex justify-space-between align-center">
              <v-date-picker v-model="date" full-width no-title></v-date-picker>
              <input type="time" class="text-h2 timer" v-model="time" />
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="grey lighten-1" variant="text" @click.native="clearHandler(isActive)"
              >{{ clearText }}
            </v-btn>
            <v-btn color="green darken-1" variant="text" @click="okHandler(isActive)"
              >{{ okText }}
            </v-btn>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>
<script lang="ts" setup>
import { format, parse } from 'date-fns'

import { computed, nextTick, onMounted, Ref, ref, watch } from 'vue'

const DEFAULT_TIME = '00:00:00'
const DEFAULT_DATE_FORMAT = 'yyyy-MM-dd'
const DEFAULT_TIME_FORMAT = 'HH:mm:ss'
const emit = defineEmits(['update:modelValue', 'input'])

const props = defineProps({
  modelValue: {
    type: String
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
    default: 620
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
const date = ref()
const time = ref(DEFAULT_TIME)
const timer = ref()

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
const init = () => {
  if (!props.modelValue) {
    return
  }
  // see https://stackoverflow.com/a/9436948
  let initDateTime = parse(props.modelValue, dateTimeFormat.value, new Date())
  date.value = initDateTime
  time.value = format(initDateTime, DEFAULT_TIME_FORMAT)
}

const okHandler = (isActive: Ref) => {
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
  isActive.value = false
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
