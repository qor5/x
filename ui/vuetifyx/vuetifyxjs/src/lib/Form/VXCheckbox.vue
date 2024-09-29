<template>
  <div class="vx-checkbox-wrap">
    <v-checkbox
      v-if="!readonly" 
      :label="label"
      v-model="model"
      :true-icon="trueIcon"
      :false-icon="falseIcon"
      v-bind="filteredAttrs"
      class="ms-n2"
    />
    <div v-if="readonly" class="d-flex flex-column ga-2 pb-4">
      <VXLabel  :tooltip="tips">{{ label }}</VXLabel>
      <div class="d-flex align-center ga-2">
        <v-icon :icon="readonlyIcon" :color="readonlyColor"></v-icon>
        <span class="v-label">{{ readonlyLabel }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import VXLabel from '../Common/VXLabel.vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
const { filteredAttrs } = useFilteredAttrs()

const model = defineModel<boolean | undefined>({ default: undefined })

const props = withDefaults(
  defineProps<{
    readonly?: boolean | undefined
    tips?: string | undefined,
    label?: string | undefined,
    value?: boolean | undefined,
    trueLabel?: string | undefined,
    falseLabel?: string | undefined,
    trueIcon?: string | undefined,
    falseIcon?: string | undefined,
    trueColor?: string | undefined,
    falseColor?: string | undefined,
  }>(),
  {
    readonly: false,
    trueLabel: "YES",
    falseLabel: "NO",
    trueColor: "primary",
    falseColor: "grey-darken-1",
  }
)

const readonlyValue = computed(() => {
  return model.value !== undefined ? model.value : props.value !== undefined ? props.value : false;
});

const readonlyColor = computed(() => {
  return readonlyValue.value ? props.trueColor : props.falseColor
});

const readonlyLabel = computed(() => {
  return readonlyValue.value ? props.trueLabel : props.falseLabel
});

const readonlyIcon = computed(() => {
  const icon = readonlyValue.value ? props.trueIcon : props.falseIcon
  return icon ?? "mdi-circle-outline"
});

</script>

<style lang="scss" scoped>
.vx-checkbox-wrap {
  margin-bottom: 2px;
}
</style>