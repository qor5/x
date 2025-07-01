<template>
  <div class="vx-label-wrap">
    <!-- label-title -->
    <div class="vx-label-title d-flex align-center">
      <!-- prepend slot -->
      <slot name="prepend" />

      <!-- toggle-label has click event -->
      <template v-if="toggleLabel">
        <span class="toggle-label-wrap" @click="onClickToggleLabel">
          <v-icon
            icon="mdi-menu-down"
            :size="toggleIconSize"
            :color="tooltipIconColor || 'black'"
            v-bind="propsWithoutIcon"
            class="mr-1"
            :class="{ isFolded: toggleStatus }"
          />
          <label v-if="hasDefaultSlot" class="text-subtitle-2 text-high-emphasis" :for="labelFor">
            <slot />
          </label>
        </span>
      </template>
      <!-- normal-label only for display -->
      <template v-else>
        <label v-if="hasDefaultSlot" class="text-subtitle-2 text-high-emphasis" :for="labelFor">
          <slot />
        </label>
      </template>

      <!-- requiredSymbol -->
      <span v-if="requiredSymbol" class="required-symbol ml-1 text-error">*</span>

      <v-tooltip v-if="tooltip" :location="tooltipLocation">
        <pre class="tooltip-display">{{ tooltip }}</pre>
        <template v-slot:activator="{ props }">
          <v-icon
            :icon="icon"
            :size="iconSize"
            :color="tooltipIconColor"
            v-bind="props"
            class="ml-1"
          />
        </template>
      </v-tooltip>
    </div>

    <!-- toggle-label can have a slot be controlled by toggle-label -->
    <div v-if="toggleLabel && toggleStatus" class="toggled-label-content-wrap">
      <slot name="toggle-content" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, useSlots, ref, PropType } from 'vue'

const slots = useSlots()
const props = defineProps({
  tooltip: String,
  tooltipLocation: {
    type: String as PropType<'top' | 'bottom' | 'left' | 'right' | 'start' | 'end'>,
    default: 'end'
  },
  toggleLabel: Boolean,
  toggleIconSize: {
    type: String,
    default: 'default'
  },
  tooltipIconColor: String,
  labelFor: String,
  icon: {
    type: String,
    default: 'mdi-information-outline'
  },
  iconSize: {
    type: String,
    default: 'small'
  },
  requiredSymbol: {
    type: Boolean,
    default: false
  }
})
const { icon: _, ...propsWithoutIcon } = props
const hasDefaultSlot = !!slots.default
const toggleStatus = ref(true)

const onClickToggleLabel = () => {
  toggleStatus.value = !toggleStatus.value
}
</script>

<style lang="scss" scoped>
.toggle-label-wrap {
  & > .v-icon {
    transform: rotate(-90deg);
    transition: transform ease 0.3s;
    &.v-icon--size-default {
      font-size: 18px;
    }
    &.isFolded {
      transform: rotate(0deg);
    }
  }
  label {
    cursor: pointer;
  }
}
.vx-label-title {
  .v-icon {
    color: rgb(var(--v-theme-grey-darken-1));
    cursor: pointer;
  }

  .v-icon--size-small {
    font-size: 16px;
  }
}

.tooltip-display {
  max-width: 50vw;
  white-space: pre-wrap;
  word-break: break-all;
}

.required-symbol {
  line-height: 1;
}
</style>
