<template>
  <div class="vx-segment-item-wrap" :class="{ readonly: readonly }">
    <div class="condition-left">
      <div class="connect-decoration" />
      <VXConditionSwitch
        class="vx-switcher"
        v-model="groupForm.condition"
        type="dropdown"
        @change="handleDataChange"
        :disabled="readonly"
      />
    </div>
    <div class="content-right">
      <VXSegmentItem
        class="segment-item-record"
        v-for="(item, idx) in groupForm.list"
        :key="getItemKey(item, idx)"
        :modelValue="item"
        :validate="validate"
        :readonly="readonly"
        ref="segmentItemRefs"
        @on-remove="() => handleRemoveItem(idx)"
        @update:modelValue="(value) => handleUpdateItem(idx, value)"
      />
      <div class="add-btn" v-if="!readonly">
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
import { ref, defineEmits, PropType, watch, computed } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItem from './SegmentItem.vue'
import type { ConditionItemType } from './type'
import { genRecordModel, useItemKeys } from './useUtils'

// Group condition type definition
interface GroupFormType {
  condition: string
  list: Array<Record<string, any>>
}

const props = defineProps({
  modelValue: {
    type: Object as PropType<GroupFormType>,
    default: () => ({
      condition: 'And',
      list: []
    })
  },
  index: {
    type: Number,
    default: 0
  },
  validate: {
    type: Boolean,
    default: false
  },
  readonly: {
    type: Boolean,
    default: false
  }
})

const groupForm = ref<GroupFormType>({
  condition: 'And',
  list: []
})

watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal) {
      // 只在外部数据变化时更新本地状态
      if (groupForm.value.condition !== newVal.condition) {
        groupForm.value.condition = newVal.condition || 'And'
      }

      // 只在外部数据变化时更新本地列表
      if (Array.isArray(newVal.list)) {
        // 如果列表长度不同，直接替换
        groupForm.value.list = newVal.list
      }
    }
  },
  { deep: true, immediate: true }
)

const emit = defineEmits(['on-remove', 'on-data-change'])

const { getItemKey } = useItemKeys()

const segmentItemRefs = ref<any[]>([])

// Validate all items in this group
defineExpose({
  validate: () => {
    if (!segmentItemRefs.value || segmentItemRefs.value.length === 0) return true

    return segmentItemRefs.value.filter((item: any) => !item.isValid()).length === 0
  }
})

// Remove an item from group
const handleRemoveItem = (idx: number) => {
  if (groupForm.value.list.length > 1) {
    // Remove the item
    groupForm.value.list = groupForm.value.list.filter((_, index) => index !== idx)
    emit('on-data-change', { idx: props.index, value: groupForm.value })
  } else {
    // When it's the last item in the group, notify parent to remove the entire group
    emit('on-remove', props.index)
  }
}

// Handle condition type change (And/Or)
const handleDataChange = () => {
  emit('on-data-change', { idx: props.index, value: { ...groupForm.value } })
}

// Update item in the group
const handleUpdateItem = (idx: number, updatedItem: Record<string, any>) => {
  groupForm.value.list[idx] = updatedItem

  emit('on-data-change', { idx: props.index, value: { ...groupForm.value } })
}

// Add new condition item
function handleAddItem() {
  const newItem = genRecordModel()
  getItemKey(newItem, groupForm.value.list.length)

  groupForm.value.list.push(newItem)

  emit('on-data-change', { idx: props.index, value: { ...groupForm.value } })
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
