<template>
  <div class="vx-range-picker-wrap">
    <vx-field
      ref="inputFieldParent"
      class="vx-range-picker-field"
      :class="{ isFocus }"
      :label="label"
      :required-symbol="required"
      :tooltip="tips"
      :label-for="name"
      v-bind="filteredAttrs"
      v-model:focused="isFocus"
    >
      <template #append-inner>
        <v-icon
          icon="mdi-calendar-range-outline ml-auto"
          size="x-small"
          @click="onClickEditDate(0)"
        />
      </template>

      <div class="vx-range-picker-group d-flex flex-1-1">
        <vx-field
          :class="{ current: current === 0 }"
          v-model:focused="isStartInputFocus"
          ref="startDateInput"
          placeholder="Start at"
          variant="flat"
          class="flex-1-1"
          hide-details
          @click="onClickEditDate(0)"
        />
        <div class="separator" />
        <vx-field
          :class="{ current: current === 1 }"
          v-model:focused="isEndInputFocus"
          ref="endDateInput"
          placeholder="End at"
          variant="flat"
          class="flex-1-1"
          hide-details
          @click="onClickEditDate(1)"
        />
      </div>

      <!-- drop down -->
      <v-overlay
        :model-value="showMenu"
        persistent
        target="parent"
        :scrim="false"
        :open-delay="0"
        :close-delay="0"
        no-click-animation
        min-width="292"
        location-strategy="connected"
        @click:outside="closeEditData()"
      >
        <date-picker-base
          class="elevation-2 d-inline-block bg-background rounded-lg overflow-hidden"
          :model-value="datePickerValue"
          :use-time-select="type === 'datetimepicker'"
          @update:modelValue="onDatePickerValueChange"
        />
      </v-overlay>
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, computed, useSlots, PropType } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import useBindingValue from '@/lib/composables/useBindingValue'
import datePickerBase from './DatePickerBase.vue'
const { filteredAttrs } = useFilteredAttrs()

const inputFieldParent = ref()
const startDateInput = ref()
const endDateInput = ref()
const current = ref()
const isStartInputFocus = ref(false)
const isEndInputFocus = ref(false)
const datePickerValue = ref()
const showMenu = ref(false)

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: [String, Number] as PropType<string | string[]>,
  label: String,
  type: String,
  tips: String,
  id: String,
  name: String,
  required: Boolean,
  passwordVisibleToggle: [Boolean, undefined] as PropType<boolean | undefined>,
  passwordVisibleDefault: Boolean
})
const { bindingValue, onUpdateModelValue } = useBindingValue(props, emit)

const isFocus = computed(() => isStartInputFocus.value || isEndInputFocus.value)

function onClickEditDate(index: number) {
  showMenu.value = true
  current.value = index
}

function closeEditData() {
  if (isFocus.value) return
  showMenu.value = false
  current.value = null
}

const combinedProps = computed(() => ({
  density: 'compact',
  variant: 'outlined',
  modelValue: bindingValue.value,
  id: props.id,
  name: props.name,
  'onUpdate:modelValue': onUpdateModelValue,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))

function onDatePickerValueChange() {}
</script>

<style lang="scss" scoped>
.vx-range-picker-field {
  .current > :deep(.v-input):not(.v-input--error, .v-input--readonly) {
    .v-field {
      &::after {
        height: 3px !important;
        transition: all ease 0.3s;
        background: #3e63dd;
        width: calc(100% - 24px);
      }
    }
  }

  & > :deep(.v-input) {
    & .vx-range-picker-group + input {
      display: none;
    }

    &:not(.v-input--error, .v-input--readonly) {
      & > .v-input__control > .v-field {
        &:hover > .v-field__outline,
        & > .v-field__outline {
          color: rgb(var(--v-theme-grey-lighten-2)) !important;
        }
      }
    }

    & > .v-input__control > .v-field {
      padding-inline-start: 0;

      & > .v-field__field > .v-field__input {
        padding: 0;
      }

      .v-field {
        position: relative;
        &::after {
          transition: all ease 0.3s;
          position: absolute;
          height: 0;
          content: '';
          bottom: -2px;
          left: 12px;
        }
      }
    }
  }

  // &.isFocus > :deep(.v-input) {
  //   & > .v-input__control > .v-field--focused {
  //     .v-field--focused {
  //       &::after {
  //         height: 3px;
  //         transition: all ease 0.3s;
  //         background: #3e63dd;
  //         width: calc(100% - 24px);
  //       }
  //     }
  //   }
  // }
}

.separator {
  display: flex;
  justify-content: center;
  align-items: center;
  &::before {
    display: block;
    content: '';
    height: 1px;
    background: rgb(var(--v-theme-grey));
    width: 16px;
  }
}
</style>
