<template>
  <div class="vx-field-wrap" :class="rootAttrs.class" :style="rootAttrs.style">
    <VXLabel
      v-if="label"
      :label-for="name"
      :tooltip="tips"
      :required-symbol="required"
      class="mb-2"
      >{{ label }}</VXLabel
    >

    <!-- text-area -->
    <template v-if="type === 'textarea'">
      <v-textarea
        ref="vInputRef"
        v-model:focused="vInputFocus"
        :rows="2"
        :max-rows="20"
        v-bind="combinedProps"
        auto-grow
      />
    </template>

    <!-- password -->
    <template v-else-if="type === 'password'">
      <v-text-field
        ref="vInputRef"
        v-model:focused="vInputFocus"
        class="password-field"
        :type="passwordFieldType"
        v-bind="combinedProps"
      >
        <template #append-inner>
          <slot v-if="hasAppendInnerSlot" name="append-inner" />
          <v-icon
            v-else-if="passwordVisibleToggle"
            :icon="!passwordVisible ? 'mdi-eye-off' : 'mdi-eye'"
            size="xsmall"
            @click="passwordVisible = !passwordVisible"
          />
        </template>
        <slot></slot>
      </v-text-field>
    </template>

    <!-- number -->
    <template v-else-if="type === 'number'">
      <v-number-input
        ref="vInputRef"
        class="number-field"
        control-variant="stacked"
        v-model:focused="vInputFocus"
        :on-update:model-value="onUpdateModelValue"
        inset
        v-bind="combinedProps"
      >
        <template
          v-if="hasPrependInnerSlot"
          #prepend-inner="{ isActive, isFocused, controlRef, focus, blur }"
        >
          <slot name="prepend-inner" :props="{ isActive, isFocused, controlRef, focus, blur }" />
        </template>
      </v-number-input>
      <!-- slot for v-menu and so on -->
      <slot></slot>
    </template>

    <!-- v-text-file -->
    <template v-else>
      <v-text-field ref="vInputRef" v-model:focused="vInputFocus" v-bind="combinedProps">
        <template #append-inner>
          <slot name="append-inner" />
        </template>

        <template #prepend-inner>
          <slot name="prepend-inner" />
        </template>

        <slot></slot>
      </v-text-field>
    </template>
  </div>
</template>

<script setup lang="ts">
import {
  defineEmits,
  ref,
  defineExpose,
  computed,
  useSlots,
  PropType,
  defineOptions,
  Slots
} from 'vue'
import VXLabel from '../Common/VXLabel.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import useBindingValue from '@/lib/composables/useBindingValue'
import { forwardRefs } from '@/lib/composables/forwardRefs'
const { filteredAttrs, rootAttrs } = useFilteredAttrs()

const vInputRef = ref()
const vInputFocus = ref(false)
const slots: Slots = useSlots()
const hasAppendInnerSlot = slots['append-inner'] !== undefined
const hasPrependInnerSlot = slots['prepend-inner'] !== undefined
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: [String, Number, Array] as PropType<string | string[]>,
  label: String,
  id: String,
  type: String,
  tips: String,
  name: String,
  required: Boolean,
  passwordVisibleToggle: [Boolean, undefined] as PropType<boolean | undefined>,
  passwordVisibleDefault: Boolean
})
const passwordVisible = ref(props.passwordVisibleDefault)
const { bindingValue, onUpdateModelValue } = useBindingValue(props, emit, modelValueFormatter)

// fix: when modelValue is a string, the number input will not work
function modelValueFormatter(value: any) {
  if (props.type === 'number' && typeof value !== 'number') {
    if (['', null].includes(value)) {
      return undefined
    }
    return Number(value)
  }
  return value
}

const passwordFieldType = computed(() => {
  if (props.passwordVisibleToggle === undefined) return 'password'

  return passwordVisible.value ? 'text' : 'password'
})

// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const combinedProps = computed(() => ({
  density: 'compact',
  variant: 'outlined',
  modelValue: bindingValue.value,
  id: props.id,
  name: props.name,
  'onUpdate:modelValue': onUpdateModelValue,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))

defineExpose(
  forwardRefs(
    {
      blur() {
        vInputFocus.value = false
      }
    },
    vInputRef
  )
)
</script>

<style lang="scss" scoped>
.vx-field-wrap {
  margin-bottom: 2px;

  .v-input {
    &.v-input--disabled {
      &:deep(.v-field) {
        background-color: rgb(var(--v-theme-grey-lighten-4));
        color: rgb(var(--v-theme-grey));
      }
    }

    &:deep(.v-field) {
      --v-theme-overlay-multiplier: var(--v-theme-background-overlay-multiplier);
      background-color: rgb(var(--v-theme-background));
    }

    &:deep(.v-field__outline) {
      --v-field-border-width: 1px;
      --v-field-border-opacity: 1;
      transition: color 0.3s ease;
    }

    &:not(.v-input--error):deep(.v-field__outline) {
      color: rgb(var(--v-theme-grey-lighten-2));
    }

    &.v-input--error:deep(.v-field__clearable),
    &.v-input--error:deep(.v-field__append-inner) {
      .v-icon {
        color: rgb(var(--v-theme-grey-darken-3));
      }
    }

    &:deep(.v-input__details > .v-messages) {
      order: 1;
    }

    &:deep(.v-counter) {
      order: 0;
      margin-right: 8px;
      white-space: nowrap;
      color: rgb(var(--v-theme-grey-darken-1));
      letter-spacing: 0;
      word-spacing: -3px;
    }

    &:deep(.v-input__details),
    &:deep(.v-messages__message) {
      padding: 0;
      min-height: 20px;
      line-height: 20px;
      align-items: flex-start;
    }

    &:not(.v-input--error, .v-input--readonly):deep(.v-field__outline) {
      color: rgb(var(--v-theme-grey-lighten-2));
      transition: color 0.3s ease;
    }

    &:not(.v-input--error, .v-input--readonly):deep(.v-field:not(.v-field--focused)):hover
      .v-field__outline {
      color: rgb(var(--v-theme-primary));
    }

    &:not(.v-input--error, .v-input--readonly):deep(.v-field--focused) .v-field__outline {
      color: rgb(var(--v-theme-primary));
    }

    &:deep(input) {
      color: rgb(var(--v-theme-grey-darken-3));
    }

    &.v-input--density-compact:deep(input) {
      &::placeholder {
        font-size: 16px;
        color: rgb(var(--v-theme-grey));
        opacity: 1;
      }
    }
  }

  &:deep(.v-field__clearable) {
    i {
      font-size: 16px;
      color: rgb(var(--v-theme-grey-darken-3));
      --v-medium-emphasis-opacity: 1;
    }
  }

  &:deep(.v-field__append-inner) i {
    font-size: 16px;
    color: rgb(var(--v-theme-grey-darken-3));
  }

  .number-field {
    &:deep(.v-number-input__control) {
      .v-btn {
        --v-btn-size: 13.75px;
        --v-btn-height: 12px;
        font-size: var(--v-btn-size);
      }
      .v-btn--variant-elevated {
        box-shadow: none;
      }
      .v-divider {
        display: none;
      }
    }
  }
}
</style>
