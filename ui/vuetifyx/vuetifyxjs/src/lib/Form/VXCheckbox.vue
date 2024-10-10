<template>
  <div class="vx-checkbox-wrap">
    <VXLabel :tooltip="tips">{{ title }}</VXLabel>
    <v-checkbox
      v-model="model"
      class="ms-n2"
      color="primary"
      :label="labelDisplay"
      :true-icon="trueIcon"
      :false-icon="falseIcon"
      :readonly="readonly"
      :hide-details="hideDetails"
      v-bind="filteredAttrs"
      :class="{ checked: model, readonly }"
    />
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
    readonly?: boolean
    tips?: string
    label?: string
    value?: boolean
    trueLabel?: string
    falseLabel?: string
    trueIcon?: string
    falseIcon?: string
    trueIconColor?: string
    falseIconColor?: string
    title?: string
    hideDetails?: boolean
  }>(),
  {
    hideDetails: false,
    title: '',
    label: '',
    readonly: false,
    trueLabel: '',
    falseLabel: '',
    trueIconColor: '',
    falseIconColor: ''
  }
)

const labelDisplay = computed(() => {
  const label = model.value ? props.trueLabel : props.falseLabel
  return label || props.label
})

const isRGBorHexColor = (colorStr: string) => /rgb|#/.test(colorStr)

const vIconStyle = computed(() => {
  const trueIconColor = isRGBorHexColor(props.trueIconColor)
      ? props.trueIconColor
      : `rgb(var(--v-theme-${props.trueIconColor}))`

  const falseIconColor = isRGBorHexColor(props.falseIconColor)
      ? props.falseIconColor
      : `rgb(var(--v-theme-${props.falseIconColor}))`


  if(model.value) {
    return trueIconColor || 'grey-darken-1'
  } else {
    return falseIconColor || 'grey-darken-1'
  }
})
</script>

<style lang="scss" scoped>

.v-input {

  &.readonly {
    &:deep(.v-selection-control) {
      pointer-events: none;
    }
  }

  &:deep(.v-label) {
    color: rgb(var(--v-theme-grey-darken-3))
  }

  &:deep(.v-icon) {
    color: v-bind(vIconStyle);
  }


}
</style>
