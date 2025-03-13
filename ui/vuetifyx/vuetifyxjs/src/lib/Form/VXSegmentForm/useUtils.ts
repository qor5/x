import type { ConditionItemType } from './type'
import { ref, shallowRef } from 'vue'

// 使用 WeakMap 替代 Map，提高性能并避免内存泄漏
// WeakMap 允许键是对象，当对象不再被引用时，相应的键值对会被自动垃圾回收
const itemKeysMap = shallowRef(new WeakMap())

export function useItemKeys() {
  // 为每个项生成唯一键
  function getItemKey(item: any, index: number): string {
    if (!itemKeysMap.value.has(item)) {
      itemKeysMap.value.set(item, `item-${Date.now()}-${index}`)
    }
    return itemKeysMap.value.get(item)
  }

  return {
    getItemKey
  }
}

export function useCondition(props: { modelValue: { [key: string]: ConditionItemType[] } }) {
  const condition = ref('And')
  const editorModelList = ref<ConditionItemType[]>([])

  if (Object.keys(props.modelValue).length > 0) {
    const key = Object.keys(props.modelValue)[0]
    condition.value = getConditionKey(key, 'internal')
    editorModelList.value = props.modelValue[key] ?? []
  }

  function getConditionKey(condition: string, type: 'external'): 'intersect' | 'union'
  function getConditionKey(condition: string, type: 'internal'): 'And' | 'Or'
  function getConditionKey(
    condition: string,
    type: 'external' | 'internal' = 'external'
  ): 'intersect' | 'union' | 'And' | 'Or' {
    if (type === 'external') {
      return condition === 'And' ? 'intersect' : 'union'
    }

    return condition === 'intersect' ? 'And' : 'Or'
  }

  return {
    condition,
    editorModelList,
    getConditionKey
  }
}

export function genRecordModel() {
  return {
    tag: {
      builderID: '',
      params: {},
      values: []
    }
  }
}
