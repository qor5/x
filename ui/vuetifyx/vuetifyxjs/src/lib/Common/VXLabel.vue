<template>
  <span class="vx-label-wrap d-inline-flex align-center">
    <label v-if="hasDefaultSlot" class="text-subtitle-2 text-high-emphasis"><slot /></label>
    <v-tooltip v-if="tooltip" ><pre>{{ tooltip }}</pre>
      <template v-slot:activator="{ props }">
        <v-icon :icon="icon" :size="iconSize" :color="tooltipIconColor" v-bind="props" class="ml-1"/>
      </template>
    </v-tooltip>
  </span>
</template>

<script setup lang="ts">
import { useSlots } from 'vue';
import { defineProps } from 'vue';
const slots = useSlots();
const props = defineProps({
  tooltip: String,
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

const hasDefaultSlot = !!slots.default;

</script>

<style lang="scss" scoped>
.vx-label-wrap {
  .v-icon {
    color: rgb(var(--v-theme-grey-darken-1));
    cursor: pointer;
  }

  .v-icon--size-small {
    font-size: 16px;
  }

}
</style>
