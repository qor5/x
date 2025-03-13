<template>
  <div class="vx-segment-item-wrap">
    <div class="condition-left">
      <div class="connect-decoration" />
      <VXConditionSwitch
        class="vx-switcher"
        v-model="groupForm.condition"
        type="dropdown"
        @change="handleDataChange"
      />
    </div>
    <div class="content-right">
      <VXSegmentItem
        class="segment-item-record"
        v-for="(item, idx) in groupForm.list"
        :key="getItemKey(item, idx)"
        :modelValue="item"
        @on-select="handleSelectChange(item, $event)"
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
import { ref, defineEmits, PropType, watch, computed } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItem from './SegmentItem.vue'
import type { ConditionItemType } from './type'
import { useCondition, genRecordModel, useItemKeys } from './useUtils'

const props = defineProps({
  modelValue: {
    type: Object as any,
    default: () => ({})
  },
  index: {
    type: Number,
    default: 0
  }
})

const groupForm = ref({
  condition: 'And',
  list: []
})

watch(
  () => props.modelValue,
  (newVal) => {
    groupForm.value = newVal
  },
  { deep: true, immediate: true }
)

const emit = defineEmits(['on-remove', 'on-data-change'])

const wrappedProps = {
  modelValue: {
    union: 'union' in props.modelValue ? (props.modelValue as any).union : [props.modelValue]
  }
}

const { getItemKey } = useItemKeys()

const handleRemoveItem = (idx: number) => {
  if (groupForm.value.list.length > 1) {
    groupForm.value.list = groupForm.value.list.filter((_, index) => index !== idx)
    emit('on-data-change', { idx: props.index, value: groupForm.value })
  } else {
    emit('on-remove', props.index)
  }
}

const handleDataChange = (value: any) => {
  emit('on-data-change', { idx: props.index, value: groupForm.value })
}

const handleSelectChange = (item: any, value: any) => {
  // emit('on-data-change', getFormData())
  item.tag.builderID = value
  emit('on-data-change', { idx: props.index, value: groupForm.value })
}

function handleAddItem() {
  const newItem = genRecordModel()
  getItemKey(newItem, groupForm.value.list.length)
  // @ts-ignore :TODO: fix this
  groupForm.value.list.push(newItem)
  emit('on-data-change', { idx: props.index, value: groupForm.value })
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
