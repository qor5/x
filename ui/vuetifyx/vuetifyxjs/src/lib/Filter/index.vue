<template>
  <div class="vx-filter-wrap">
    <div class="filter-item">
      <template v-for="(item, key) in options" :key="key">
        <vx-field v-if="item === 'search'" />
      </template>
    </div>
    <div class="filter-result"></div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue'
const props = defineProps({
  modelValue: {
    type: Object,
    default: {}
  },
  options: {
    type: Object,
    default: []
  }
})
</script>

<style lang="scss" scoped></style>
