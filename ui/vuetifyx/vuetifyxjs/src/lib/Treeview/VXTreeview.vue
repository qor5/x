<template>
  <div
    class="vx-treeview-wrap"
    :class="[rootAttrs.class, `presets-${props.density}`]"
    :style="rootAttrs.style"
    v-bind="rootAttrs"
  >
    <v-treeview v-bind="combinedProps">
      <template v-if="slots.prepend" v-slot:prepend="{ item, isOpen, isActive, isSelected, select }"
        ><slot
          name="prepend"
          :item="item"
          :isOpen="isOpen"
          :isActive="isActive"
          :isSelected="isSelected"
          :select="select"
        />
      </template>
      <template v-if="!isDefaultSlotReallyEmpty" #default><slot /></template>
      <template v-if="slots.divider" v-slot:divider="{ props }"
        ><slot name="divider" :props="props"
      /></template>
    </v-treeview>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions, PropType } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  color: {
    type: String,
    default: 'primary'
  },
  density: {
    type: String,
    default: 'compact'
  }
})

const defaultOptions = computed(() => {
  return {
    'base-color': 'grey-darken-3',
    activatable: true,
    'open-on-click': true,
    slim: true
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
  density: props.density,
  color: props.color,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-treeview-wrap {
  &.presets-compact {
    &:deep(.v-list-item) {
      min-height: 32px;
    }
  }

  &:deep(.v-list-item__spacer) {
    width: 8px !important;
  }

  &:deep(.v-list-item-title) {
    font-weight: 400;
  }

  &:deep(.v-list-item) {
    &[aria-selected='true'] {
      .v-list-item__prepend .v-icon {
        color: v-bind('props.color') !important;
      }
    }

    &:hover:not([aria-selected='true']) {
      color: var(--v-theme-on-surface) !important;
      .v-list-item__prepend .v-icon {
        --v-medium-emphasis-opacity: 1;
        color: var(--v-theme-on-surface) !important;
      }
      .v-list-item-action .v-btn__content {
        color: var(--v-theme-grey-darken-3) !important;
      }
    }

    .v-list-item__overlay {
      border-radius: 4px;
    }
  }
  &:deep(.v-list-item-action) {
    // margin-right: 10px;
    .v-btn__content {
      color: rgb(var(--v-theme-grey-darken-1));
      .mdi-menu-right {
        &::before {
          content: '\F0142';
        }
      }
      .mdi-menu-down {
        &::before {
          content: '\F0140';
        }
      }
    }
  }
}
</style>
