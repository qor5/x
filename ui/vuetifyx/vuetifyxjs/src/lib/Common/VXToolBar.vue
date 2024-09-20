<template>
  <div class="vx-toolbar-wrap">
    <div class="vx-toolbar vx-toolbar-default">
      <slot>{{ text }}</slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
const props = defineProps({
  text: String,
  type: {
    type: String,
    default: 'default'
  }
})
</script>

<style lang="scss" scoped>
.vx-toolbar {
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: 24px; /* 150% */
  padding: 8px 16px;
  border-radius: 8px;
}

.vx-toolbar-default {
  background-color: rgb(var(--v-theme-grey-lighten-4));
}
</style>
