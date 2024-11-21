<template>
  <div class="vx-paginator-wrap" v-bind="rootAttrs">
    <v-pagination
      v-model="bindingValue"
      total-visible="5"
      show-first-last-page
      density="compact"
      active-color="primary"
      v-bind="combinedProps"
    />
  </div>
</template>

<script setup lang="ts">
import { defineProps, PropType, computed } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import useBindingValue from '@/lib/composables/useBindingValue'
const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const props = defineProps({
  modelValue: [String, Number, Array] as PropType<string | string[]>
})
const emit = defineEmits(['update:modelValue'])
const { bindingValue, onUpdateModelValue } = useBindingValue(props, emit)
// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const combinedProps = computed(() => ({
  modelValue: bindingValue.value,
  'onUpdate:modelValue': onUpdateModelValue,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-paginator-wrap {
  &:deep(.v-pagination) {
    .v-pagination__item {
      flex-shrink: 0.05;
      margin-left: 0.15em;
      margin-right: 0.15em;

      .v-btn {
        --v-btn-size: unset;
        font-weight: 500;
        width: auto;

        &.v-btn--size-small {
          --v-btn-size: 13.875px;
        }
        &.v-btn--size-default {
          --v-btn-size: 16px;
        }
        &.v-btn--density-compact {
          .v-btn__content {
            padding: 0 10px;
            min-width: 28px;
          }
        }
      }
    }
  }
}
</style>
