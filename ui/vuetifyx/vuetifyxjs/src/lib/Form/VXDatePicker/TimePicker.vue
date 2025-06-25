<template>
  <div class="vx-timepicker-wrap">
    <vx-label
      v-if="label"
      class="mb-2"
      :tooltip="tips"
      :label-for="name"
      :required-symbol="required"
      >{{ label }}</vx-label
    >

    <vx-field
      v-model="inputValue"
      :placeholder="placeholder"
      :focused="isFocus"
      ref="inputRef"
      @blur="onInputBlur"
      @mouseover="isHovering = true"
      @mouseout="isHovering = false"
      @keydown.enter="onInputBlur(inputValue, true)"
      v-bind="filteredAttrs"
      :style="minWidth"
    >
      <!-- time icon -->
      <template v-if="!hideAppendInner" #append-inner
        ><v-icon
          :icon="showClearIcon ? 'mdi-close-circle' : 'mdi-clock-outline'"
          @click.stop="onClickAppendInner"
      /></template>

      <!-- drop down -->
      <v-overlay
        v-model="showMenu"
        open-on-click
        :close-on-content-click="false"
        activator="parent"
        :scrim="false"
        :open-delay="0"
        :close-delay="0"
        min-width="220"
        scroll-strategy="close"
        location-strategy="connected"
      >
        <div class="elevation-5 d-inline-block bg-background rounded-lg overflow-hidden">
          <time-select
            class="time-select-wrap pa-4"
            :format-str="formatStr"
            v-model="timeValue"
            v-bind="timePickerProps"
            @update:modelValue="emitTimeValue"
          />
          <div v-if="needConfirm" class="d-flex justify-end pa-2">
            <v-btn color="primary" variant="text" @click="onConfirm">确定</v-btn>
          </div>
        </div>
      </v-overlay>

      <input readonly class="input-cover" :value="inputValue" :placeholder="placeholder" />
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, PropType, watchEffect, watch, ref } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import TimeSelect from './TimeSelect.vue'
import dayjs from 'dayjs'
const { filteredAttrs } = useFilteredAttrs()

const props = defineProps({
  modelValue: [String, Number],
  tips: String,
  name: String,
  required: Boolean,
  label: String,
  placeholder: String,
  clearable: Boolean,
  needConfirm: Boolean,
  format: {
    type: String,
    default: 'HH:mm:ss'
  },
  hideAppendInner: Boolean,
  timePickerProps: Object
})

const inputValue = ref('')
const inputRef = ref()
const timeValue = ref('00:00:00')
const isHovering = ref(false)
const showMenu = ref(false)
const pendingTimeValue = ref('')
const isFocus = computed(() => showMenu.value)
const emit = defineEmits(['update:modelValue', 'blur', 'click:confirm'])

const formatStr = computed(() => props.format || 'HH:mm:ss')

const showClearIcon = computed(
  () => (isHovering.value || showMenu.value) && inputValue.value && props.clearable
)

const minWidth = computed(() => ({
  minWidth: '120px'
}))

watch(
  () => showMenu.value,
  (oldVal, newVal) => {
    // this state is when finished select time and dropdown is closed
    if (!oldVal && newVal && !props.needConfirm) {
      emitTimeValue(timeValue.value, { extraEmitEvents: ['blur'] })
    }
  }
)

watchEffect(() => {
  convertValueForInput({ value: props.modelValue, shouldEmit: false })
})

function onInputBlur(obj: FocusEvent | string, closeMenu: boolean = false) {
  if (closeMenu) {
    showMenu.value = false
    inputRef.value.blur()
  }

  let value

  if (obj instanceof FocusEvent) {
    const target = obj.target as HTMLInputElement
    value = target.value
  } else {
    value = obj
  }

  // the first time select time will trigger blur event
  if (!value) return

  convertValueForInput({ value, shouldEmit: true, extraEmitEvents: ['blur'] })
}

function convertValueForInput({
  value,
  shouldEmit,
  extraEmitEvents
}: {
  value: string | number | undefined
  shouldEmit?: boolean
  extraEmitEvents?: string[]
}) {
  //case: no init value
  if (!value) {
    inputValue.value = ''
    timeValue.value = '00:00:00'
  } else {
    // Ensure value is in HH:mm:ss format internally
    let fullTimeValue = value
    if (typeof value === 'string' && value.includes(':')) {
      const parts = value.split(':')
      if (parts.length === 2) {
        // If only HH:mm is provided, add :00 for seconds
        fullTimeValue = `${parts[0]}:${parts[1]}:00`
      }
    }

    inputValue.value = props.format
      ? dayjs(`2000-01-01 ${value}`).format(props.format)
      : String(value)
    timeValue.value =
      typeof fullTimeValue === 'string' ? fullTimeValue : dayjs(fullTimeValue).format('HH:mm:ss')
  }

  shouldEmit && emitTimeValue(timeValue.value, { extraEmitEvents })
}

function emitTimeValue(value: string, { extraEmitEvents = [] as string[] } = {}) {
  if (props.needConfirm) {
    pendingTimeValue.value = value
    return
  }

  // Format the value according to the user's specified format
  const formattedValue = props.format ? dayjs(`2000-01-01 ${value}`).format(props.format) : value

  emit('update:modelValue', formattedValue)

  if (extraEmitEvents.includes('blur')) {
    emit('blur', formattedValue)
  }
}

function onClickAppendInner() {
  if (showClearIcon.value) {
    emit('update:modelValue', '')
    showMenu.value = false
  } else {
    showMenu.value = true
  }
}

function onConfirm() {
  const value = pendingTimeValue.value || timeValue.value
  // Format the value according to the user's specified format
  const formattedValue = props.format ? dayjs(`2000-01-01 ${value}`).format(props.format) : value

  emit('update:modelValue', formattedValue)
  emit('click:confirm', { value: ref([formattedValue]), next: Promise.resolve() })
  emit('blur', formattedValue)
  showMenu.value = false
}
</script>

<style lang="scss" scoped>
.v-menu {
  &:deep(.v-overlay__content) {
    border-radius: 8px !important;
  }
}

.vx-timepicker-wrap {
  &:deep(.v-input) {
    .input-cover {
      position: absolute;
      width: 100%;
      height: 100%;
      z-index: 1;
      pointer-events: none;
    }

    .v-field {
      cursor: pointer;
    }

    & input:not(.input-cover) {
      display: none;
    }
  }
}

.time-select-wrap {
  justify-content: center;
  width: 100%;
}
</style>
