<template>
  <span class="iconTip-wrap">
    <v-tooltip :text="text">
      <template v-slot:activator="{ props }">
        <v-icon :icon="icon" :size="iconSize" v-bind="props" />
      </template>
    </v-tooltip>
  </span>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';
const props = defineProps({
  text: String,
  icon: {
    type: String,
    default: "mdi-help-circle-outline"
  },
  iconSize: {
    type: String,
    default: "default"
  },

})

</script>

<style lang="scss" scoped>
.v-icon {
  color: rgb(var(--v-theme-grey-lighten-1));
  cursor: pointer;
  font-size: 20px;
}
</style>
