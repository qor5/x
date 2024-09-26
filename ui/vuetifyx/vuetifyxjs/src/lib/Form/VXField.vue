<template>
  <div class="vx-field-wrap">
    <VXLabel v-if="label" :label-for="name" :tooltip="tips" class="mb-2">{{ label }}</VXLabel>

    <!-- text-area -->
    <template v-if="type === 'textarea'">
      <v-textarea
        ref="vInputRef"
        :id="id"
        :name="name"
        :autofocus="autoFocus"
        :readonly="readonly"
        :rows="2"
        :max-rows="20"
        auto-grow
        variant="outlined"
        density="compact"
        v-model="fieldValue"
        :error-messages="errorFiled"
        :disabled="disabled"
        :placeholder="placeholder"
        v-bind="filteredAttrs"
        @update:modelValue="onUpdateModelValue"
      />
    </template>

    <!-- v-text-file -->
    <template v-else>
      <v-text-field
        :id="id"
        ref="vInputRef"
        :name="name"
        :autofocus="autoFocus"
        :readonly="readonly"
        density="compact"
        variant="outlined"
        v-model="fieldValue"
        :type="type"
        :error-messages="errorFiled"
        :disabled="disabled"
        :placeholder="placeholder"
        v-bind="filteredAttrs"
        @update:modelValue="onUpdateModelValue"
      />
    </template>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, PropType, watch, ref, defineExpose } from 'vue'
import VXLabel from '../Common/VXLabel.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import { forwardRefs } from '@/lib/composables/forwardRefs'
const { filteredAttrs } = useFilteredAttrs()
const vInputRef = ref()

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: [String, Number],
  label: String,
  type: String,
  errorMessages: [String, Array] as PropType<string | string[]>,
  remoteValidation: Boolean,
  disabled: Boolean,
  placeholder: String,
  tips: String,
  readonly: Boolean,
  autoFocus: Boolean,
  name: String,
  id: String //id will passthrough set to input, thus click label will focus on input element
})

const fieldValue = ref(props.modelValue)
const errorFiled = ref(props.errorMessages)

watch(()=> props.modelValue, newVal => fieldValue.value = newVal)

function onUpdateModelValue(value: string | number | Record<string, any>) {
  emit('update:modelValue', value)
}

defineExpose(forwardRefs({}, vInputRef))
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

    &:not(.v-input--error,.v-input--readonly):deep(.v-field__outline) {
      color: rgb(var(--v-theme-grey-lighten-2));
      transition: color 0.3s ease;
    }

    &:not(.v-input--error,.v-input--readonly):deep(.v-field:not(.v-field--focused)):hover .v-field__outline {
      color: rgb(var(--v-theme-primary));
    }

    &:not(.v-input--error,.v-input--readonly):deep(.v-field--focused) .v-field__outline {
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
}
</style>
