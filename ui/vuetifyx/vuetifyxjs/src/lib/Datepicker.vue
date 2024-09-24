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
            <div class="d-flex flex-column align-center px-6">
              <v-date-picker
                v-model="dateOfPicker"
                v-bind="datePickerProps"
                :year="displayedYear"
                :month="displayedMonth"
                @update:year="onYearOrMonthChange($event, 'year')"
                @update:month="onYearOrMonthChange($event, 'month')"
                @update:modelValue="onYearOrMonthChange($event, 'modelValue')"
              />
            </div>
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
                    :disabled="valueChangedWithoutSaved"
                    @click="okHandler(isActive)"
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
import { computed, nextTick, onMounted, Ref, ref, watch } from 'vue'
import { useLocale } from 'vuetify'
const { t } = useLocale()
const okTips = t('$vuetify.datePicker.okTips')
import { useVDatePickerTimeChange } from '@/lib/composables/useVDatePicker'

const props = defineProps({
  modelValue: {
    type: String,
    default: null
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
  clearText: {
    type: String,
    default: 'CLEAR'
  },
  okText: {
    type: String,
    default: 'OK'
  },
  datePickerProps: {
    type: Object
  },
  hideDetails: {
    type: Boolean
  }
})

const dialogVisible = ref(false)
const date = ref()
const dateOfPicker = ref()
const {
  displayedMonth,
  displayedYear,
  setDisplayedYearAndMonth,
  onYearOrMonthChange,
  valueChangedWithoutSaved
} = useVDatePickerTimeChange(dateOfPicker)

const dateTimeFormat = computed(() => {
  return props.dateFormat
})
const formattedDatetime = computed(() => {
  return date.value ? format(<Date>date.value, dateTimeFormat.value) : ''
})
const init = () => {
  if (!props.modelValue) {
    return
  }
  // see https://stackoverflow.com/a/9436948
  date.value = parse(props.modelValue, dateTimeFormat.value, new Date())
  dateOfPicker.value = date.value
  setDisplayedYearAndMonth(date.value)
}

watch(dialogVisible, (newVal) => {
  if (newVal) {
    dateOfPicker.value = date.value
    setDisplayedYearAndMonth(date.value)
  }
})

const emit = defineEmits(['update:modelValue'])

const okHandler = (isActive: Ref) => {
  date.value = dateOfPicker.value
  isActive.value = false
  if (!date.value) {
    date.value = new Date()
  }
  emit('update:modelValue', formattedDatetime.value)
}
const clearHandler = (isActive: Ref) => {
  isActive.value = false
  date.value = null
  emit('update:modelValue', null)
}

onMounted(() => {
  nextTick(() => {
    init()
  })
})
</script>
