<template>
  <div class="vx-select-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-2 d-sm-inline-block">
      {{ label }}
    </span>
    <v-autocomplete
      v-if="type === 'autocomplete'"
      variant="outlined"
      density="compact"
      :items="items",
      :item-title="itemTitle"
      :item-value="itemValue"
      :multiple="multiple"
      :chips="chips"
      :clearable="clearable"
    />
    <v-select
      v-else
      :model-value="modelValue"
      variant="outlined"
      density="compact"
      :disabled="disabled"
      :items="items"
      v-bind="attrs"
    />
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch } from "vue"
const props = defineProps({
  modelValue: [String, Number],
  type: String,
  label: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object,
  placeholder: String,
  items: Array,
  itemTitle: String,
  itemValue: String,
  multiple: Boolean,
  chips: Boolean,
  clearable: Boolean
})

const selectValue = ref(props.modelValue)

watch(() => props.modelValue, (newValue) => {
  selectValue.value = newValue
})

const emit = defineEmits(["update:modelValue"])

function onUpdateModelValue(value: any) {
  emit("update:modelValue", value)
  selectValue.value = value
}

</script>

<style lang="scss" scoped>
.vx-select-wrap {
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
