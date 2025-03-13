<template>
  <div class="vx-segment-form">
    <VXConditionSwitch v-if="editorModelList.length > 0" v-model="condition" />
    <div class="vx-segment-form-block">
      <div v-if="editorModelList.length > 0" class="content">
        <VXSegmentItemGroup
          v-for="(item, idx) in editorModelList"
          :key="idx"
          :item="item"
          :index="idx"
          :options="options"
          @on-remove="handleRemoveGroup"
        />
      </div>
    </div>
    <vx-btn prepend-icon="mdi-plus" presets="x-small" @click="handleAddRule">Add Rule</vx-btn>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, computed, watch, PropType } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItemGroup from './SegmentItemGroup.vue'
import type { ConditionItemType, OptionsType } from './type'
import { useCondition } from './useUtils'
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

const { condition, getConditionKey } = useCondition()
const emit = defineEmits(['update:modelValue'])
const editorModelList = ref<ConditionItemType[]>([])

// 初始化
if (Object.keys(props.modelValue).length > 0) {
  const key = Object.keys(props.modelValue)[0]
  condition.value = getConditionKey(key, 'internal')
  editorModelList.value = props.modelValue[key] ?? []
}

function handleConditionChange(condition: string) {
  // console.log(savedFormModel.value)
}

function handleAddRule() {
  editorModelList.value.push({})
}

function handleUpdateModelValue() {
  emit('update:modelValue', {
    [getConditionKey(condition.value, 'external')]: editorModelList.value
  })
}

function handleRemoveGroup(idx: number) {
  editorModelList.value.splice(idx, 1)
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
