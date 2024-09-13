<template>
  <div class="vx-text-field-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-1 d-sm-inline-block">
      {{ label }}
    </span>
    <v-text-field
      class="section-field"
      density="compact"
      variant="outlined"
      bg-color="background"
      :type="type"
      :error-messages="errorMessages"
      :disabled="disabled"
      v-bind="attrs"
    />
  </div>
</template>

<script setup>
const props = defineProps({
  label: String,
  type: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object
})
</script>

<style lang="less" scoped></style>
