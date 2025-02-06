<template>
  <div
    class="vx-btn-wrap"
    :class="[rootAttrs.class, `presets-${presets}`, { 'presets-icon': combinedProps.icon }]"
    :style="rootAttrs.style"
    v-bind="rootAttrs"
  >
    <!-- bugfix: icon props not allowed to use slot at the same time -->
    <v-btn v-if="combinedProps.icon" v-bind="combinedProps" />
    <v-btn v-else v-bind="combinedProps">
      <template v-if="slots.prepend" #prepend>
        <slot name="prepend" />
      </template>

      <template v-if="!isDefaultSlotReallyEmpty" #default>
        <slot />
      </template>

      <template v-if="slots.append" #append>
        <slot name="append" />
      </template>
    </v-btn>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  presets: {
    type: String,
    default: 'default'
  }
})

const isDefaultSlotReallyEmpty = computed(() => {
  return !slots.default || !slots.default({}).length
})

const presetsSizeOptions = computed(() => {
  const obj: { color: string; width: string; size: string } = {
    color: 'primary',
    width: 'auto',
    size: 'default'
  }

  switch (props.presets) {
    case 'x-small':
      obj.size = 'small'
      break

    case 'large':
      obj.size = 'large'
      break
  }

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
.vx-btn-wrap {
  .v-btn {
    height: unset;
  }
  &.presets-default {
    &.presets-icon {
      .v-btn {
        padding: 8px;
      }
    }
    .v-btn {
      padding: 8px 16px;
    }
    &:deep(.v-btn__append) {
      margin-left: 4px;
    }
    &:deep(.v-btn__prepend) {
      margin-right: 4px;
    }
    &:deep(.v-btn__prepend),
    &:deep(.v-btn__append) {
      margin-inline: unset;
      .v-icon {
        font-size: 20px;
      }
    }
    &:deep(.v-btn__content) {
      line-height: 20px;
      .v-icon {
        font-size: 20px;
      }
    }
  }

  &.presets-large {
    &.presets-icon {
      .v-btn {
        padding: 12px;
      }
    }
    .v-btn {
      padding: 12px 24px;
      min-width: initial;
    }
    &:deep(.v-btn__content) {
      letter-spacing: 0.244px;
      font-size: 16px;
      font-weight: 400;
      line-height: 24px;
      .v-icon {
        font-size: 24px;
      }
    }
    &:deep(.v-btn__append) {
      margin-left: 4px;
    }
    &:deep(.v-btn__prepend) {
      margin-right: 4px;
    }
    &:deep(.v-btn__prepend),
    &:deep(.v-btn__append) {
      margin-inline: unset;
      .v-icon {
        font-size: 24px;
      }
    }
  }

  &.presets-default {
    &:deep(.v-btn__content) {
      letter-spacing: 0.091px;
      font-weight: 500;
    }
  }

  &.presets-small {
    &.presets-icon {
      .v-btn {
        padding: 6px;
      }
    }
    .v-btn {
      padding: 6px 12px;
      min-width: initial;
    }
    &:deep(.v-btn__append) {
      margin-left: 4px;
    }
    &:deep(.v-btn__prepend) {
      margin-right: 4px;
    }
    &:deep(.v-btn__append),
    &:deep(.v-btn__prepend) {
      margin-inline: unset;
      .v-icon {
        font-size: 16px;
      }
    }
    &:deep(.v-btn__content) {
      letter-spacing: -0.143px;
      font-size: 12px;
      line-height: 16px;
      font-weight: 400;
      .v-icon {
        font-size: 16px;
      }
    }
  }

  &.presets-x-small {
    line-height: 1;
    &.presets-icon {
      .v-btn {
        padding: 4px;
      }
    }
    .v-btn {
      padding: 4px 8px;
    }
    &:deep(.v-btn__append) {
      margin-left: 4px;
    }
    &:deep(.v-btn__prepend) {
      margin-right: 4px;
    }
    &:deep(.v-btn__append),
    &:deep(.v-btn__prepend) {
      margin-inline: unset;
      .v-icon {
        font-size: 16px;
      }
    }
    &:deep(.v-btn__content) {
      letter-spacing: -0.14px;
      font-size: 12px;
      line-height: 16px;
      font-weight: 400;

      .v-icon {
        font-size: 16px;
      }
    }
  }
}
</style>
