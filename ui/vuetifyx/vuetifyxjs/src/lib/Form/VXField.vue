<template>
  <div class="vx-field-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-2 d-sm-inline-block">
      {{ label }}
    </span>
    <v-text-field class="section-field" density="compact" variant="outlined" :value="modelValue"
      :type="type" :error-messages="errorMessages" :disabled="disabled" :placeholder="placeholder" v-bind="attrs"
      @update:modelValue="onUpdateModelValue" />
  </div>
</template>

<script setup lang="ts">
import { defineEmits } from "vue"
defineProps({
  modelValue: [String, Number],
  label: String,
  type: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object,
  placeholder: String
})

const emit = defineEmits(["update:modelValue"])

function onUpdateModelValue(value: any) {
  emit("update:modelValue", value)
}

</script>

<style lang="scss" scoped>
// @media (hover: hover) {
//   .vx-field-wrap .v-input:deep(.v-field:not(.v-field--focused)):hover .v-field__outline {

//       --v-field-border-opacity:1
//   }
// }

.vx-field-wrap {
  margin-bottom: 2px;

  .v-input {
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
        font-size: 14px;
        color: rgb(var(--v-theme-grey));
        opacity: 1;
      }
    }
  }
}
</style>
