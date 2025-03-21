<template>
  <div class="vx-select-wrap">
    <VXLabel v-if="label" :tooltip="tips" class="mb-2" :required-symbol="required">{{
      label
    }}</VXLabel>
    <v-autocomplete
      v-if="type === 'autocomplete'"
      :closable-chips="closableChips"
      :hide-no-data="hideNoData"
      v-model="selectValue"
      :items="items"
      :item-title="itemTitle"
      :item-value="itemValue"
      :multiple="multiple"
      :chips="chips"
      :clearable="clearable"
      :placeholder="placeholder"
      :disabled="disabled"
      :error-messages="errorMessages"
      :hide-details="hideDetails"
      v-bind="filteredAttrs"
      class="vx-type-autocomplete"
      variant="outlined"
      density="compact"
      color="primary"
      @update:model-value="onUpdateModelValue"
    >
      <template
        v-if="hasPrependInnerSlot"
        #prepend-inner="{ isActive, isFocused, controlRef, focus, blur }"
      >
        <slot
          name="prepend-inner"
          :isActive="isActive"
          :isFocused="isFocused"
          :controlRef="controlRef"
          :focus="focus"
          :blur="blur"
          :selectedItems="selectedItems"
        />
      </template>

      <template v-if="hasItemSlot" #item="{ props, index, item }">
        <slot name="item" :props="props" :index="index" :item="item" />
      </template>
    </v-autocomplete>
    <v-select
      v-else
      :closable-chips="closableChips"
      v-model="selectValue"
      :hide-no-data="hideNoData"
      :items="items"
      :item-title="itemTitle"
      :item-value="itemValue"
      :multiple="multiple"
      :chips="chips"
      :clearable="clearable"
      :placeholder="placeholder"
      :disabled="disabled"
      :error-messages="errorMessages"
      :hideDetails="hideDetails"
      v-bind="filteredAttrs"
      class="vx-type-select"
      variant="outlined"
      density="compact"
      color="primary"
      @update:model-value="onUpdateModelValue"
    >
      <template
        v-if="hasPrependInnerSlot"
        #prepend-inner="{ isActive, isFocused, controlRef, focus, blur }"
      >
        <slot
          name="prepend-inner"
          :isActive="isActive"
          :isFocused="isFocused"
          :controlRef="controlRef"
          :focus="focus"
          :blur="blur"
          :selectedItems="selectedItems"
        />
      </template>

      <template v-if="hasItemSlot" #item="{ props, index, item }">
        <slot name="item" :props="props" :index="index" :item="item" />
      </template>
    </v-select>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch, PropType, useSlots, Slots, computed } from 'vue'
import VXLabel from '../Common/VXLabel.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
const { filteredAttrs } = useFilteredAttrs()

const emit = defineEmits(['update:modelValue'])
const slots: Slots = useSlots()
const hasPrependInnerSlot = slots['prepend-inner'] !== undefined
const hasItemSlot = slots['item'] !== undefined
const props = defineProps({
  modelValue: null,
  type: String,
  label: String,
  errorMessages: [String, Array] as PropType<string | string[]>,
  remoteValidation: Boolean,
  disabled: Boolean,
  hideDetails: Boolean,
  hideNoData: Boolean,
  placeholder: String,
  items: Array,
  itemTitle: String,
  itemValue: String,
  multiple: Boolean,
  chips: Boolean,
  closableChips: Boolean,
  clearable: Boolean,
  tips: String,
  required: Boolean
})

const selectValue = ref(props.modelValue)

const selectedItems = computed(() => {
  return props.items?.filter((item: any) => {
    if (props.multiple) {
      return selectValue.value?.includes(props.itemValue ? item[props.itemValue] : item)
    }

    return Array.isArray(selectValue.value)
      ? selectValue.value[0] === (props.itemValue ? item[props.itemValue] : item)
      : selectValue.value === (props.itemValue ? item[props.itemValue] : item)
  })
})

watch(
  () => props.modelValue,
  (newVal) => (selectValue.value = newVal)
)

function onUpdateModelValue(value: any) {
  emit('update:modelValue', value)
}
</script>

<style lang="scss" scoped>
.vx-select-wrap {
  // margin-bottom: 2px;

  .v-input {
    &.v-input--disabled {
      &:deep(.v-field) {
        background-color: rgb(var(--v-theme-grey-lighten-4));
        color: rgb(var(--v-theme-grey));
      }
    }

    &:deep(.v-autocomplete__selection),
    &:deep(.v-select__selection) {
      margin-inline-end: 4px;
      .v-chip {
        color: rgb(var(--v-theme-primary));
      }
    }

    &:deep(.v-field) {
      --v-theme-overlay-multiplier: var(--v-theme-background-overlay-multiplier);
      background-color: rgb(var(--v-theme-background));

      .v-field__clearable .mdi-close-circle {
        font-size: 18px;
        color: rgb(var(--v-theme-grey-darken-3));
        --v-medium-emphasis-opacity: 1;
      }

      .v-field__append-inner .mdi-menu-down {
        font-size: 16px;
      }

      .v-chip__close {
        .mdi-close-circle {
          font-size: 16px;
          &::before {
            content: '\F0156';
          }
        }
      }
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

    &:deep(.v-input__details) {
      padding: 0;
      min-height: 20px;
      align-items: center;
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
