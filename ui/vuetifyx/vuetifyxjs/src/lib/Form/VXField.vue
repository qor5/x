<template>
  <div class="vx-field-wrap">
    <span class="text-subtitle-2 text-high-emphasis mb-2 d-inline-flex align-center">
      {{ label }}<IconTip v-if="tips !== undefined" :text="tips" class="ml-1"/>
    </span>
    <v-text-field density="compact" variant="outlined" :model-value="fiedValue"
      :type="type" :error-messages="errorMessages" :disabled="disabled" :placeholder="placeholder" v-bind="attrs"
      @update:modelValue="onUpdateModelValue" />
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch } from "vue"
import IconTip from "../Common/IconTip.vue"

const emit = defineEmits(["update:modelValue"])
const props = defineProps({
  modelValue: [String, Number],
  label: String,
  type: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object,
  placeholder: String,
  tips: String
})

const fiedValue = ref(props.modelValue)

watch(() => props.modelValue, (newValue) => {
  fiedValue.value = newValue
})

function onUpdateModelValue(value: any) {
  emit("update:modelValue", value)
  fiedValue.value = value
}

</script>

<style lang="scss" scoped>
.vx-field-wrap {
  margin-bottom: 2px;

  .v-input {
    &.v-input--disabled {
      &:deep(.v-field) {
        background-color: rgb(var(--v-theme-grey-lighten-4));
        color: rgb(var(--v-theme-grey))
      }
    }

    &:deep(.v-field) {
      --v-theme-overlay-multiplier: var(--v-theme-background-overlay-multiplier);
      background-color: rgb(var(--v-theme-background));
    }

    &:deep(.v-field__outline) {
      --v-field-border-width: 1px;
      --v-field-border-opacity:1;
      color: rgb(var(--v-theme-grey-lighten-2));
      transition: color .3s ease;
    }

    &:deep(.v-field:not(.v-field--focused)):hover .v-field__outline{
      color: rgb(var(--v-theme-primary));
    }

    &:deep(.v-field--focused) .v-field__outline {
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
