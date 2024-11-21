<template>
  <v-tabs
    class="vx-tabs-wrap"
    :class="[
      {
        'underline-border-contain': underlineBorder === 'contain',
        'underline-border-full': underlineBorder === 'full'
      }
    ]"
    :model-value="bindingValue"
    :v-bind="attrs"
    color="primary"
    @update:model-value="onUpdateModelValue"
  >
    <slot></slot>
  </v-tabs>
</template>

<script setup lang="ts">
import { defineProps, PropType, computed } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import useBindingValue from '@/lib/composables/useBindingValue'
const { attrs } = useFilteredAttrs()
const props = defineProps({
  modelValue: [String, Number, Array] as PropType<string | string[]>,
  underlineBorder: {
    type: String,
    default: ''
  },
  widthContain: Boolean
})
const emit = defineEmits(['update:modelValue'])
const { bindingValue, onUpdateModelValue } = useBindingValue(props, emit)

// defineOptions({
//   inheritAttrs: false
// })
</script>

<style lang="scss" scoped>
.vx-tabs-wrap.v-tabs--horizontal {
  &.underline-border-contain {
    &:deep(.v-btn) {
      border-block-end-width: thin;
      border-block-end-style: solid;
      border-block-end-color: rgba(var(--v-border-color), var(--v-border-opacity)) !important;
    }
  }

  &.underline-border-full {
    &:deep(.v-slide-group__content) {
      border-block-end-width: thin;
      border-block-end-style: solid;
      border-block-end-color: rgba(var(--v-border-color), var(--v-border-opacity)) !important;
    }
  }
}

.vx-tabs-wrap {
  &:deep(.v-slide-group__container) {
    .v-tab.v-tab.v-btn {
      min-width: auto;
    }
  }
}
</style>
