<template>
  <div class="vx-segment-form">
    <VXConditionSwitch
      v-if="form.list.length > 0"
      v-model="form.condition"
      @change="handleConditionChange"
    />
    <div class="vx-segment-form-block">
      <div v-if="form.list.length > 0" class="content">
        <VXSegmentItemGroup
          v-for="(item, idx) in form.list"
          :key="getItemKey(item, idx)"
          :modelValue="item"
          :index="idx"
          @on-remove="handleRemoveGroup"
          @on-data-change="handleUpdateModelValue"
        />
      </div>
    </div>
    <vx-btn prepend-icon="mdi-plus" presets="x-small" @click="handleAddRule">Add Rule</vx-btn>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, PropType, provide, watch } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItemGroup from './SegmentItemGroup.vue'
import type { ConditionItemType, OptionsType } from './type'
import { useCondition, genRecordModel, useItemKeys, convertModel } from './useUtils'

const props = defineProps({
  modelValue: {
    type: Object as PropType<{ [key: string]: ConditionItemType[] }>,
    default: () => ({})
  },
  options: {
    type: Array as PropType<OptionsType[]>,
    default: () => []
  }
})

provide('segmentOptions', props.options)

const { getItemKey } = useItemKeys()
const emit = defineEmits(['update:modelValue'])
const form = ref<any>({
  condition: 'And',
  list: []
})

watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) {
      form.value = convertModel(newVal)
    }
  },
  { immediate: true }
)

function handleConditionChange(condition: string) {
  emitDataChange()
}

function handleAddRule() {
  const newItem = {
    condition: 'Or',
    list: [genRecordModel()]
  }
  getItemKey(newItem, form.value.list.length)
  form.value.list.push(newItem)
  emitDataChange()
}

function handleUpdateModelValue({ idx, value }: { idx: number; value: any }) {
  form.value.list[idx] = value
  emitDataChange()
}

function handleRemoveGroup(idx: number) {
  form.value.list.splice(idx, 1)
  emitDataChange()
}

function emitDataChange() {
  const isValid =
    form.value.list.length > 0 &&
    form.value.list.every(
      (group: any) =>
        group.list &&
        group.list.length > 0 &&
        group.list.every((item: any) => item.tag && item.tag.builderID)
    )

  if (!isValid) {
    console.log('Form data is incomplete, not emitting update')
    return
  }

  // Convert condition types to intersect/union format
  const getConditionKey = (condition: string): string => {
    return condition === 'And' ? 'intersect' : 'union'
  }

  // Create the nested structure with intersect/union
  const externalFormat = {
    [getConditionKey(form.value.condition)]: form.value.list.map((group: any) => ({
      [getConditionKey(group.condition)]: group.list.map((item: any) => ({
        tag: {
          builderID: item.tag.builderID,
          params: item.tag.params || {}
        }
      }))
    }))
  }

  emit('update:modelValue', externalFormat)

  console.log('Current form state:', JSON.stringify(form.value, null, 2))
  console.log('Emitted model:', JSON.stringify(externalFormat, null, 2))
}
</script>

<style lang="scss" scoped>
.vx-segment-form {
  .vx-segment-form-block .content {
    position: relative;
    padding: 16px 0 16px 24px;
    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 14px;
      width: 1px;
      background: rgba(189, 189, 189, 1);
      height: 100%;
    }
  }
}
</style>
