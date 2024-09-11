<template>
  <div class="vx-field-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-1 d-sm-inline-block">
      {{ label }}
    </span>
    <v-text-field
      class="section-field"
      density="compact"
      variant="outlined"
      bg-color="background"
      :model-value="modelValue"
      :type="type"
      :error-messages="errorMessages"
      :disabled="disabled"
      v-bind="attrs"
    />
  </div>
</template>

<script setup>
 const props = defineProps({
  modelValue:[String,Number],
  label: String,
  type: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object
 })

</script>

<style lang="less" scoped>

</style>
