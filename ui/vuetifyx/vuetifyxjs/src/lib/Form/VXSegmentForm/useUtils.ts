import type { ConditionItemType } from './type'

import { ref } from 'vue'

export function useCondition(initCondition: 'And' | 'Or' = 'And') {
  const condition = ref(initCondition)

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
    condition: condition,
    getConditionKey
  }
}
