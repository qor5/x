<template>
  <div class="vx-breadcrumbs-wrap" v-bind="rootAttrs">
    <v-breadcrumbs v-bind="combinedProps" class="pa-0">
      <template v-if="!isDefaultSlotReallyEmpty" #default>
        <slot />
      </template>

      <template v-if="slots.divider" #divider>
        <slot name="divider" />
      </template>

      <template v-if="slots.prepend" #prepend>
        <slot name="prepend" />
      </template>

      <template v-if="slots.title" v-slot:title="{ item, index }">
        <slot name="title" :item="item" :index="index" />
      </template>
    </v-breadcrumbs>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions, PropType } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({})

const defaultOptions = computed(() => {
  return {
    divider: '»'
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
.vx-breadcrumbs-wrap {
  &:deep(.v-breadcrumbs-item) {
    padding-left: 0;
    padding-right: 0;
  }

  &:deep(.v-breadcrumbs-item),
  &:deep(.v-breadcrumbs-divider) {
    color: rgb(var(--v-theme-grey-darken-3));
  }

  &:deep(.v-breadcrumbs-item--disabled):has(.v-breadcrumbs-item--link) {
    --v-disabled-opacity: 1;
    .v-breadcrumbs-item--link {
      color: rgb(var(--v-theme-grey-darken-3));
    }
  }

  &:deep(.v-breadcrumbs-item--link) {
    color: rgb(var(--v-theme-primary));
  }
}
</style>
