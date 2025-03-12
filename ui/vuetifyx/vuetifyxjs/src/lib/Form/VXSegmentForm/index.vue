<template>
  <div class="vx-segment-form">
    <VXConditionSwitch
      v-if="editorModelList.length > 0"
      v-model="condition"
      @change="handleConditionChange"
    />
    <div class="vx-segment-form-block">
      <div v-if="editorModelList.length > 0" class="content">
        <VXSegmentItemGroup
          v-for="(item, idx) in editorModelList"
          :key="idx"
          :item="item"
          :index="idx"
          @on-remove="handleRemoveGroup"
        />
      </div>
    </div>
    <vx-btn prepend-icon="mdi-plus" presets="x-small" @click="handleAddRule">Add Rule</vx-btn>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, computed } from 'vue'
import VXConditionSwitch from './ConditionSwitch.vue'
import VXSegmentItemGroup from './SegmentItemGroup.vue'

type ConditionType = 'intersect' | 'union'

type TagType = {
  tag: {
    builderID: string
    params: Record<string, any>
  }
}

type ConditionItemType =
  | TagType
  | {
      [key in ConditionType]?: ConditionItemType[]
    }

type SavedFormType = {
  [key in ConditionType]?: ConditionItemType[]
}

const props = defineProps({
  modelValue: {
    type: Object,
    default: () => ({})
  },
  options: {
    type: Object,
    default: () => ({})
  }
})

const condition = ref('And')

const editorModelList = ref<ConditionItemType[]>([])

const getConditionKey = (condition: string) => {
  return condition === 'And' ? 'intersect' : 'union'
}

const level1Condition = computed(() => {
  return getConditionKey(condition.value)
})

const currentGroupModel = computed(() => {
  return form.value[level1Condition.value]
})

const form = computed(() => {
  return {
    [level1Condition.value]: [
      // {
      //   intersect: [
      //     {
      //       tag: {
      //         builderID: 'user_gender',
      //         params: {
      //           operator: 'EQ',
      //           value: 'FEMALE'
      //         }
      //       }
      //     },
      //     {
      //       tag: {
      //         builderID: 'user_age',
      //         params: {
      //           max: 35,
      //           min: 25,
      //           operator: 'BETWEEN'
      //         }
      //       }
      //     },
      //     {
      //       tag: {
      //         builderID: 'user_city',
      //         params: {
      //           operator: 'IN',
      //           values: ['TOKYO', 'OSAKA']
      //         }
      //       }
      //     },
      //     {
      //       union: [
      //         {
      //           tag: {
      //             builderID: 'user_signup_source',
      //             params: {
      //               operator: 'EQ',
      //               value: 'WEBSITE'
      //             }
      //           }
      //         },
      //         {
      //           tag: {
      //             builderID: 'user_signup_source',
      //             params: {
      //               operator: 'EQ',
      //               value: 'MOBILE_APP'
      //             }
      //           }
      //         }
      //       ]
      //     }
      //   ]
      // },
      // {
      //   tag: {
      //     builderID: 'event_purchase',
      //     params: {
      //       accumulation: 'DAYS',
      //       countOperator: 'GTE',
      //       countValue: 2,
      //       timeRange: '30D'
      //     }
      //   }
      // }
    ]
  }
})

const handleConditionChange = (condition: string) => {
  console.log(form.value)
}

const handleAddRule = () => {
  editorModelList.value.push({})
}

const handleRemoveGroup = (idx: number) => {
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
