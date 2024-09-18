<template>
  <div class="vx-label-wrap">
    <!-- label-title -->
    <div class="vx-label-title d-flex align-center">
      <!-- toggle-label has click event -->
      <template v-if="toggleLabel">
        <span class="toggle-label-wrap" @click="onClickToggleLabel">
          <v-icon icon="mdi-menu-down" :size="toggleIconSize" :color="tooltipIconColor || 'black'" v-bind="propsWithoutIcon"
            class="mr-1" :class="{isFolded: toggleStatus}"/>
          <label v-if="hasDefaultSlot" class="text-subtitle-2 text-high-emphasis">
            <slot />
          </label>
        </span>
      </template>
      <!-- normal-label only for display -->
      <template v-else>
        <label v-if="hasDefaultSlot" class="text-subtitle-2 text-high-emphasis">
          <slot />
        </label>
      </template>

      <v-tooltip v-if="tooltip">
        <pre>{{ tooltip }}</pre>
        <template v-slot:activator="{ props }">
          <v-icon :icon="icon" :size="iconSize" :color="tooltipIconColor" v-bind="props" class="ml-1" />
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
import { defineProps, useSlots, ref } from 'vue';

const slots = useSlots();
const props = defineProps({
  tooltip: String,
  toggleLabel: Boolean,
  toggleIconSize: {
    type: String,
    default: "default"
  },
  tooltipIconColor: String,
  icon: {
    type: String,
    default: "mdi-information-outline"
  },
  iconSize: {
    type: String,
    default: "small"
  }
})
const { icon: _, ...propsWithoutIcon } = props
const hasDefaultSlot = !!slots.default;
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
  label {cursor: pointer;}
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
</style>
