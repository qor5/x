<template>
  <div>
    <v-row v-if="row">
      <v-col v-for="(v, i) in labels" :key="i">
        <autocomplete
          :chips="chips"
          :level="i"
          :label="v"
          :hide-details="hideDetails"
          :is-paging="isPaging"
          :parent-field="parentField"
          :parent-value="parentIDValue[i]"
          :parent-id-field="props.parentIdField"
          :remote-url="remoteUrl"
          :item-title="itemTitle"
          :item-value="itemValue"
          :page-field="pageField"
          :pages-field="pagesField"
          :page-size-field="pageSizeField"
          :total-field="totalField"
          :items-field="itemsField"
          :current-field="currentField"
          :search-field="searchField"
          :page="page"
          :pageSize="pageSize"
          :level-field="levelField"
          :level-start="levelStart"
          :level-step="levelStep"
          v-model="value[i]"
          :disabled="disabled"
          :error-message="errorMessage(i)"
          :select-out-of-order="selectOutOfOrder"
          @update:model-value="changeStatus($event, i)"
        ></autocomplete>
      </v-col>
    </v-row>
    <div v-else>
      <autocomplete
        v-for="(v, i) in labels"
        :key="i"
        :chips="chips"
        :level="i"
        :label="v"
        :hide-details="hideDetails"
        :is-paging="isPaging"
        :parent-field="parentField"
        :parent-value="parentIDValue[i]"
        :parent-id-field="props.parentIdField"
        :remote-url="remoteUrl"
        :item-title="itemTitle"
        :item-value="itemValue"
        :page-field="pageField"
        :pages-field="pagesField"
        :page-size-field="pageSizeField"
        :total-field="totalField"
        :items-field="itemsField"
        :current-field="currentField"
        :search-field="searchField"
        :page="page"
        :pageSize="pageSize"
        :level-field="levelField"
        :level-start="levelStart"
        :level-step="levelStep"
        v-model="value[i]"
        :disabled="disabled"
        :error-message="errorMessage(i)"
        :select-out-of-order="selectOutOfOrder"
        @update:model-value="changeStatus($event, i)"
      ></autocomplete>
    </div>
  </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, reactive, ref, computed } from 'vue'
import Autocomplete from '@/lib/LinkageSelectRemote/components/LinkSelectAutoComplete.vue'

import get from 'lodash/get'

const props = defineProps({
  modelValue: { type: Array<Object>, default: [] },
  labels: { type: Array<string>, default: [] },
  selectOutOfOrder: { type: Boolean, default: false },
  row: { type: Boolean, default: false },
  errorMessages: { type: Array<string>, default: [] },
  disabled: { type: Boolean, default: false },
  chips: { type: Boolean, default: false },
  hideDetails: { type: Boolean, default: false },

  isPaging: { type: Boolean, default: false },
  itemTitle: { type: String, default: 'Name' },
  itemValue: { type: String, default: 'ID' },
  pageField: { type: String, default: 'page' },
  pagesField: { type: String, default: 'pages' },
  pageSizeField: { type: String, default: 'pageSize' },
  totalField: { type: String, default: 'total' },
  itemsField: { type: String, default: 'data' },
  currentField: { type: String, default: 'current' },
  searchField: { type: String, default: 'search' },
  remoteUrl: String,
  page: { type: Number, default: 1 },
  pageSize: { type: Number, default: 20 },

  parentField: { type: String, default: 'parent' },
  parentIdField: { type: String, default: 'parentID' },
  levelField: { type: String, default: 'level' },
  levelStart: { type: Number, default: 0 },
  levelStep: { type: Number, default: 1 }
})
const resolvedErrorMessages = computed(() => props.errorMessages || [])

const emit = defineEmits(['update:modelValue'])

const value = ref([...(props.modelValue ?? [])])

const parentIDValue = reactive(
  props.labels.reduce(
    (obj, _, index) => {
      obj[index + 1] = ''
      return obj
    },
    {} as { [key: number]: string }
  )
)
const changeStatus = (val: any, level: number) => {
  let newVal = val
  for (let i = props.labels.length - 1; i >= 0; i--) {
    if (i > level && !newVal) {
      //@ts-ignore
      value.value[i] = undefined
    }
    if (props.selectOutOfOrder && i < level) {
      let parent = get(newVal, props.parentField)
      if (!newVal || !parent) {
        continue
      }
      value.value[i] = parent
      newVal = parent
      //@ts-ignore
      parentIDValue[i + 1] = parentValue(i + 1)
    }
    if (!value.value[i]) {
      //@ts-ignore
      value.value[i] = undefined
    }
  }
  //@ts-ignore
  if (!val || parentIDValue[level + 1] != val[props.itemValue]) {
    for (let i = level + 1; i < props.labels.length; i++) {
      //@ts-ignore
      parentIDValue[i] = undefined
      //@ts-ignore
      value.value[i] = undefined
    }
  }
  for (let i = level + 1; i < props.labels.length; i++) {
    //@ts-ignore
    parentIDValue[i] = parentValue(i)
  }
  emit('update:modelValue', value.value)
}
const parentValue = (level: number): string => {
  if (props.selectOutOfOrder) {
    for (let i = level - 1; i >= 0; i--) {
      const v = value.value[i]
      if (v) {
        //@ts-ignore
        return v[props.itemValue]
      }
    }
    return ''
  }
  const val = value.value[level - 1]

  //@ts-ignore
  return level - 1 >= 0 && val ? val[props.itemValue] : ''
}

const errorMessage = (level: number): string => {
  if (level > resolvedErrorMessages.value.length - 1) {
    return ''
  }
  return resolvedErrorMessages.value[level]
}

onMounted(() => {
  nextTick(() => {
    value.value = [...(props.modelValue ?? [])]
    if (!props.modelValue) {
      return
    }
    for (let i = 0; i < props.modelValue.length; i++) {
      let item = props.modelValue[i]
      if (!item) {
        return
      }
      //@ts-ignore
      parentIDValue[i + 1] = item[props.itemValue]
    }
  })
})
</script>
