<template>
  <div class="vx-field-wrap">
    <VXLabel :tooltip="tips" class="mb-2">{{ label }}</VXLabel>

    <!-- text-area -->
    <template v-if="type === 'textarea'">
      <v-textarea
        :rows="2"
        :max-rows="20"
        auto-grow
        variant="outlined"
        density="compact"
        :model-value="fieldValue"
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
        density="compact"
        variant="outlined"
        :model-value="fieldValue"
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
import { defineEmits, computed, PropType, ref } from 'vue'
import VXLabel from '../Common/VXLabel.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
const { filteredAttrs } = useFilteredAttrs()

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
})

const fieldValue = computed(()=> props.modelValue)
const errorFiled = ref(props.errorMessages)

function onUpdateModelValue(value: string|number|Record<string,any>) {
  emit('update:modelValue', value)
  errorFiled.value = ''
}


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

    &:deep(.v-input__details) {
      padding:0;
      min-height: 20px;
      align-items:center;
    }

    &:not(.v-input--error):deep(.v-field__outline) {
      color: rgb(var(--v-theme-grey-lighten-2));
      transition: color 0.3s ease;
    }

    &:not(.v-input--error):deep(.v-field:not(.v-field--focused)):hover .v-field__outline {
      color: rgb(var(--v-theme-primary));
    }

    &:not(.v-input--error):deep(.v-field--focused) .v-field__outline {
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
