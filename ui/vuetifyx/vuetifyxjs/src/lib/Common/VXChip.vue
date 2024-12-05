<template>
  <div
    class="vx-chip-wrap"
    :class="[
      rootAttrs.class,
      {
        'presets-icon': combinedProps.icon,
        [`presets-${presets}`]: presets,
        [`prepend-icon`]: props.prependIcon || slots.prepend,
        [`append-icon`]: props.appendIcon || slots.append,
        [`presets-round`]: round
      }
    ]"
    :style="rootAttrs.style"
    v-bind="rootAttrs"
  >
    <!-- bugfix: icon props not allowed to use slot at the same time -->
    <v-chip v-if="combinedProps.icon" v-bind="combinedProps" />
    <v-chip v-else v-bind="combinedProps">
      <template v-if="slots.prepend" #prepend>
        <slot name="prepend" />
      </template>

      <slot />

      <template v-if="slots.append" #append>
        <slot name="append" />
      </template>
    </v-chip>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, computed, useSlots, defineOptions } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  prependIcon: {
    type: [String, Object, Function, Array],
    default: ''
  },
  appendIcon: {
    type: [String, Object, Function, Array],
    default: ''
  },
  round: {
    type: Boolean,
    default: false
  },
  presets: {
    type: String,
    default: 'badge'
  }
})

// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const presetsSizeOptions = computed(() => {
  const obj: {
    color: string
    variant: string
    size: string
    prependIcon?: string | object | (string | [string, number])[] | (new () => any)
    appendIcon?: string | object | (string | [string, number])[] | (new () => any)
  } = {
    prependIcon: props.prependIcon,
    appendIcon: props.appendIcon,
    color: 'primary',
    size: 'x-small',
    variant: 'flat'
  }
  return obj
})

const combinedProps = computed(() => ({
  ...presetsSizeOptions.value,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-chip-wrap.presets-round {
  .v-chip {
    border-radius: 999px !important;
  }
}

.vx-chip-wrap.presets-badge {
  &.presets-round {
    &.prepend-icon {
      .v-chip {
        padding-right: 6px;
      }
    }

    &.append-icon {
      .v-chip {
        padding-left: 6px;
      }
    }
  }
  .v-chip {
    &:deep(.v-icon--start),
    &:deep(.v-icon__filter) {
      margin-inline-start: 0;
      // margin-inline-end: 3px;
    }
    &:deep(.v-icon--end),
    &:deep(.v-icon__close) {
      margin-inline-end: 0;
      // margin-inline-start: 3px;
    }
  }
}

.vx-chip-wrap.presets-badge {
  .v-chip {
    font-size: 12px;
    line-height: 20px;
    padding: 0 4px;

    &:deep(.v-chip__content) {
      letter-spacing: 0.293px;
    }

    &:deep(.v-btn__content) {
      letter-spacing: -0.41px;
      font-size: 12px;
      font-weight: 400;
      .v-icon {
        font-size: 16px;
      }
    }
  }

  &:deep(.v-chip__append),
  &:deep(.v-chip__prepend) {
    margin-inline: unset;
    .v-icon {
      font-size: 12px;
      // font-size: 16px;
    }
  }
}
</style>
