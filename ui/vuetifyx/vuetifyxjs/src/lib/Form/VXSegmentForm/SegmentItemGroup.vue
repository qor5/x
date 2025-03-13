<template>
  <div class="vx-segment-item-wrap">
    <div class="condition-left">
      <div class="connect-decoration" />
      <VXConditionSwitch class="vx-switcher" v-model="condition" type="dropdown" />
    </div>
    <div class="content-right">
      <VXSegmentItem
        class="segment-item-record"
        v-for="(item, idx) in editorModelList"
        :key="getItemKey(item, idx)"
        :modelValue="item"
        @on-select="handleSelectChange(idx, $event)"
        @on-remove="handleRemoveItem(idx)"
      />
      <div class="add-btn">
        <vx-btn
          icon="mdi-plus"
          rounded
          presets="small"
          color="rgb(91,110,113)"
          @click="handleAddItem"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits, PropType, watch } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItem from './SegmentItem.vue'
import type { ConditionItemType } from './type'
import { useCondition, genRecordModel, useItemKeys } from './useUtils'

const props = defineProps({
  modelValue: {
    type: Object as PropType<ConditionItemType>,
    default: () => ({})
  },
  index: {
    type: Number,
    default: 0
  }
})

const emit = defineEmits(['on-remove', 'update:modelValue'])

// 创建一个包装对象来适配 useCondition 函数
const wrappedProps = {
  modelValue: {
    union: 'union' in props.modelValue ? (props.modelValue as any).union : [props.modelValue]
  }
}
const { condition, editorModelList } = useCondition(wrappedProps as any)
const { getItemKey } = useItemKeys()

const handleRemoveItem = (idx: number) => {
  console.log('尝试删除索引:', idx)
  console.log('删除前的列表:', JSON.stringify(editorModelList.value))

  if (editorModelList.value.length > 1) {
    // 确保我们删除的是正确的索引
    const itemToRemove = editorModelList.value[idx]
    console.log('要删除的项目:', JSON.stringify(itemToRemove))

    editorModelList.value = editorModelList.value.filter((_, index) => index !== idx)

    console.log('删除后的列表:', JSON.stringify(editorModelList.value))
  } else {
    emit('on-remove', props.index)
  }
}

const handleSelectChange = (idx: number, value: any) => {
  console.log(idx, value)
  editorModelList.value[idx].tag.builderID = value
}

const handleAddItem = () => {
  const newItem = genRecordModel()
  // 为新项预先生成一个唯一键
  getItemKey(newItem, editorModelList.value.length)
  editorModelList.value.push(newItem)
}
</script>
<style scoped lang="scss">
.vx-segment-item-wrap {
  display: flex;
}
.condition-left {
  display: flex;
  align-items: stretch;
  min-height: 128px;
  position: relative;
  .connect-decoration {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    left: 32px;
    width: 1px;
    height: calc(100% - 100px);
    background: rgb(189, 189, 189);
  }
  &:before {
    content: '';
    position: absolute;
    top: 28px;
    right: 0;
    width: 24px;
    height: 24px;
    border-top-left-radius: 8px;
    border-left: 1px solid rgb(189, 189, 189);
    border-top: 1px solid rgb(189, 189, 189);
  }
  &:after {
    content: '';
    position: absolute;
    bottom: 28px;
    right: 0;
    width: 24px;
    height: 24px;
    border-bottom-left-radius: 8px;
    border-left: 1px solid rgb(189, 189, 189);
    border-bottom: 1px solid rgb(189, 189, 189);
  }
}

.vx-switcher {
  align-self: center;
}

.content-right {
  position: relative;
  margin-left: 8px;
  padding-bottom: 56px;
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.add-btn {
  position: absolute;
  bottom: 14px;
  left: 0;
}
</style>
