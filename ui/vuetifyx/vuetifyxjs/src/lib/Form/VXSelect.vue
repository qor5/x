<template>
  <div class="vx-select-wrap">
    <span class="text-subtitle-2 text-high-emphasis section-filed-label mb-2 d-sm-inline-block">
      {{ label }}
    </span>
    <v-autocomplete
      v-if="type === 'autocomplete'"
      :model-value="selectValue"
      :items="items"
      :item-title="itemTitle"
      :item-value="itemValue"
      :multiple="multiple"
      :chips="chips"
      :clearable="clearable"
      :placeholder="placeholder"
      v-bind="attrs"
      class="vx-type-autocomplete"
      variant="outlined"
      density="compact"
      color="primary"
      @update:model-value="onUpdateModelValue"
    />
    <v-select
      v-else
      :model-value="modelValue"
      :disabled="disabled"
      :items="items"
      v-bind="attrs"
      :placeholder="placeholder"
      class="vx-type-select"
      variant="outlined"
      density="compact"
      color="primary"
      @update:model-value="onUpdateModelValue"
    />
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch, onMounted } from "vue"
const props = defineProps({
  modelValue: null,
  type: String,
  label: String,
  errorMessages: String,
  disabled: Boolean,
  attrs: Object,
  placeholder: String,
  items: Array,
  itemTitle: String,
  itemValue: String,
  multiple: Boolean,
  chips: Boolean,
  clearable: Boolean,
})

 onMounted(()=>{
  console.log(selectValue.value, props.items, props.itemTitle, props.itemValue)
 })

const selectValue = ref(props.modelValue)

watch(() => props.modelValue, (newValue) => {
  console.log("watch", )
  selectValue.value = newValue
})

const emit = defineEmits(["update:modelValue"])

function onUpdateModelValue(value: any) {
  emit("update:modelValue", value)
  selectValue.value = value
}

</script>

<style lang="scss" scoped>
.vx-select-wrap {
  .v-input {
    &:deep(.v-autocomplete__selection) {
      margin-inline-end: 4px;
      .v-chip {
        color: rgb(var(--v-theme-primary))
      }
    }

    &:deep(.v-field) {
      --v-theme-overlay-multiplier: var(--v-theme-background-overlay-multiplier);
      background-color: rgb(var(--v-theme-background));

      .v-field__clearable .mdi-close-circle{
        font-size: 18px;
        color: rgb(var(--v-theme-grey-darken-3));
        --v-medium-emphasis-opacity:1;
      }

      .v-field__append-inner .mdi-menu-down {
        font-size: 16px;
      }
    }

    &:deep(.v-field__outline) {
      --v-field-border-width: 1px;
      --v-field-border-opacity:1;
      color: rgb(var(--v-theme-grey-lighten-2));
      transition: color .3s ease;
    }

    &:deep(.v-field:not(.v-field--focused)):hover .v-field__outline{
      color: rgb(var(--v-theme-primary));
    }

    &:deep(.v-field--focused) .v-field__outline {
      color: rgb(var(--v-theme-primary));
    }

    &:deep(input) {
      color: rgb(var(--v-theme-grey-darken-3));
    }

    &.v-input--density-compact:deep(input) {
      &::placeholder {
        font-size: 16px;
        color: rgb(var(--v-theme-grey));
        opacity: 1;
      }
    }
  }
}
</style>
