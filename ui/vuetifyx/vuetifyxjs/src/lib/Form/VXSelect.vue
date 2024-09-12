<template>
  <div class="vx-select-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-2 d-sm-inline-block">
      {{ label }}
    </span>
    <v-text-field class="section-field" density="compact" variant="outlined" :value="modelValue"
      :type="type" :error-messages="errorMessages" :disabled="disabled" :placeholder="placeholder" v-bind="attrs"
      @update:modelValue="onUpdateModelValue" />
  </div>
</template>

<script setup lang="ts">
import { defineEmits } from "vue"
defineProps({
  modelValue: [String, Number],
  label: String,
  type: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object,
  placeholder: String
})

const emit = defineEmits(["update:modelValue"])

function onUpdateModelValue(value: any) {
  emit("update:modelValue", value)
}

</script>

<style lang="scss" scoped>
.vx-select-wrap {

}
</style>
