<template>
  <div class="vx-timeline-wrap" :class="{ 'vx-timeline-sinuous': sinuous }" v-bind="rootAttrs">
    <v-timeline v-bind="combinedProps">
      <template v-if="!isDefaultSlotReallyEmpty" #default>
        <slot />
      </template>
    </v-timeline>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  sinuous: Boolean
})

const defaultOptions = computed(() => {
  return {
    // Default options if any
  }
})

const isDefaultSlotReallyEmpty = computed(() => {
  /* @ts-ignore */
  return !slots.default || !slots.default().length
})

// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const combinedProps = computed(() => ({
  ...defaultOptions.value,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-timeline-wrap {
  &.vx-timeline-sinuous {
    :deep(.v-timeline-divider__before),
    :deep(.v-timeline-divider__after) {
      background-color: transparent !important;
      width: 20px;
      background-image: url("data:image/svg+xml,%3Csvg width='12' height='100' viewBox='0 0 12 100' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M6 0 Q 12 25 6 50 T 6 100' fill='none' stroke='%23ccc' stroke-width='2'/%3E%3C/svg%3E");
      background-repeat: repeat-y;
      background-size: 12px 100px;
    }

    :deep(.v-timeline-divider__before) {
      background-position: bottom center;
    }

    :deep(.v-timeline-divider__after) {
      background-position: top center;
    }
  }
}
</style>
