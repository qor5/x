<template>
  <div
    class="vx-btn-group-wrap"
    :class="[rootAttrs.class]"
    :style="rootAttrs.style"
    v-bind="rootAttrs"
  >
    <!-- bugfix: icon props not allowed to use slot at the same time -->
    <v-btn-group v-bind="combinedProps">
      <template v-if="slots.default">
        <slot />
      </template>
    </v-btn-group>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import { useColor } from '../composables/useColor'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  dividerColor: {
    type: String,
    default: 'rgba(0, 0, 0, 0.12)'
  },
  dividerWidth: {
    type: [Number, String],
    default: '1'
  }
})

const computedDividerColor = computed(() => {
  return useColor(props.dividerColor).color
})

const computedDividerWidth = computed(() => {
  return `${props.dividerWidth}px`
})

const presetsSizeOptions = computed(() => {
  const obj: {} = {}

  return obj
})

// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const combinedProps = computed(() => ({
  ...presetsSizeOptions.value,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-btn-group-wrap {
  line-height: 1;
  .v-btn-group {
    height: auto;

    &:deep(.vx-btn-wrap:not(:last-child)) .v-btn {
      border-inline-end: none;
    }

    &:deep(.vx-btn-wrap:not(:first-child)) .v-btn {
      border-inline-start: none;
    }

    &:deep(.vx-btn-wrap:first-child),
    &:deep(.vx-btn-wrap:first-child) .v-btn {
      border-start-start-radius: inherit;
      border-end-start-radius: inherit;
    }

    &:deep(.vx-btn-wrap:last-child),
    &:deep(.vx-btn-wrap:last-child) .v-btn {
      border-start-end-radius: inherit;
      border-end-end-radius: inherit;
    }
  }

  .v-btn-group--divided:deep(.vx-btn-wrap:not(:last-child)) .v-btn {
    border-inline-end-width: v-bind(computedDividerWidth);
    border-inline-end-style: solid;
    // border-inline-end-color: rgba(var(--v-border-color), var(--v-border-opacity));
    border-inline-end-color: v-bind(computedDividerColor);
  }

  .v-btn-group--divided:deep(.vx-btn-wrap:first-child) .v-btn {
    border-start-start-radius: inherit;
    border-end-start-radius: inherit;
  }
}
</style>
