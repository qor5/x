import type { ConditionItemType } from './type'
import { ref, shallowRef } from 'vue'

const itemKeysMap = shallowRef(new WeakMap())

export function useItemKeys() {
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
  return {}
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

export const convertModel = (model: ConditionItemType) => {
  return {
    condition: 'And',
    list: [
      {
        condition: 'Or',
        list: [
          {
            tag: {
              builderID: '',
              params: {},
              values: []
            }
          }
        ]
      }
    ]
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
