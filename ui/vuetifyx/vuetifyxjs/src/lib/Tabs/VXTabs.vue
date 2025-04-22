<template>
  <div class="vx-tabs-wrap">
    <v-tabs
      v-if="!pill"
      class="vx-tabs"
      :class="[
        {
          'underline-border-contain': underlineBorder === 'contain',
          'underline-border-full': underlineBorder === 'full'
        }
      ]"
      :model-value="bindingValue"
      :v-bind="filteredAttrs"
      color="primary"
      @update:model-value="onUpdateModelValue"
    >
      <slot></slot>
    </v-tabs>

    <v-tabs
      v-else
      class="vx-tabs pill-style"
      :model-value="bindingValue"
      :v-bind="filteredAttrs"
      :ripple="false"
      :hide-slider="true"
      color="#212121"
      @update:model-value="onUpdateModelValue"
    >
      <slot></slot>
    </v-tabs>
  </div>
</template>

<script setup lang="ts">
import { defineProps, PropType } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import useBindingValue from '@/lib/composables/useBindingValue'
const { filteredAttrs } = useFilteredAttrs()

const props = defineProps({
  modelValue: [String, Number, Array] as PropType<string | string[]>,
  underlineBorder: {
    type: String,
    default: ''
  },
  pill: {
    type: Boolean,
    default: false
  }
})
const emit = defineEmits(['update:modelValue'])
const { bindingValue, onUpdateModelValue } = useBindingValue(props, emit)
</script>

<style lang="scss" scoped>
.vx-tabs-wrap {
  &:deep(.v-slide-group__content) {
    border-block-end-width: 0;
  }
}

.vx-tabs-wrap .vx-tabs.v-tabs--horizontal {
  &.underline-border-contain {
    &:deep(.v-btn) {
      border-block-end-width: thin;
      border-block-end-style: solid;
      border-block-end-color: rgba(var(--v-border-color), var(--v-border-opacity)) !important;
    }
  }

  &.underline-border-full {
    &:deep(.v-slide-group__content) {
      flex: 1;
      border-block-end-width: thin;
      border-block-end-style: solid;
      border-block-end-color: rgba(var(--v-border-color), var(--v-border-opacity)) !important;
    }
  }
}

.vx-tabs-wrap .vx-tabs.pill-style {
  &:deep(.v-tab) {
    border-radius: 4px;
    padding: 0 8px;
    margin: 0;
    font-size: 12px;
    font-weight: 400;
    background-color: #eee;
    color: #757575;
    height: 24px;
    min-height: 24px;
    line-height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex: 1;

    &::before,
    &::after {
      display: none !important;
      opacity: 0 !important;
      background-color: transparent !important;
    }

    .v-ripple__container {
      display: none !important;
      opacity: 0 !important;
    }

    &.v-tab--selected {
      background-color: #fff;
      color: #212121;
    }
  }

  &:deep(.v-slide-group__content) {
    gap: 4px;
    background-color: #eee;
    border-radius: 4px;
    padding: 4px;
    height: 32px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  &:deep(.v-tabs__bar) {
    height: auto;
  }
}

.vx-tabs-wrap .vx-tabs {
  &:deep(.v-slide-group__container) {
    .v-tab.v-tab.v-btn {
      min-width: auto;
    }
  }
}
</style>
