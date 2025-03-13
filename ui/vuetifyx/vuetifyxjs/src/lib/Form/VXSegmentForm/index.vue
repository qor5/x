<template>
  <div class="vx-segment-form">
    <VXConditionSwitch v-if="editorModelList.length > 0" v-model="condition" />
    <div class="vx-segment-form-block">
      <div v-if="editorModelList.length > 0" class="content">
        <VXSegmentItemGroup
          v-for="(item, idx) in editorModelList"
          :key="getItemKey(item, idx)"
          :modelValue="item"
          :index="idx"
          @on-remove="handleRemoveGroup"
          @update:modelValue="handleUpdateModelValue"
        />
      </div>
    </div>
    <vx-btn prepend-icon="mdi-plus" presets="x-small" @click="handleAddRule">Add Rule</vx-btn>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, PropType, provide, computed } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItemGroup from './SegmentItemGroup.vue'
import type { ConditionItemType, OptionsType } from './type'
import { useCondition, genRecordModel, useItemKeys } from './useUtils'

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

const { condition, getConditionKey, editorModelList } = useCondition(props)
const { getItemKey } = useItemKeys()
const emit = defineEmits(['update:modelValue'])

function handleConditionChange(condition: string) {
  // console.log(savedFormModel.value)
}

function handleAddRule() {
  const newItem = {
    union: [genRecordModel()]
  }
  // 为新项预先生成一个唯一键
  getItemKey(newItem, editorModelList.value.length)
  editorModelList.value.push(newItem)
  handleUpdateModelValue()
}

function handleUpdateModelValue() {
  emit('update:modelValue', {
    [getConditionKey(condition.value, 'external')]: editorModelList.value
  })
}

function handleRemoveGroup(idx: number) {
  editorModelList.value.splice(idx, 1)
  handleUpdateModelValue()
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
