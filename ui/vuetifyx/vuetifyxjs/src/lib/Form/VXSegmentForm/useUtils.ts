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

// Convert external API format to internal form format
export const convertModel = (model: any) => {
  // Return default structure if model is empty or invalid
  if (!model || typeof model !== 'object' || Object.keys(model).length === 0) {
    return {
      condition: 'And',
      list: []
    }
  }

  // If model is already in internal format, return it directly
  if (model.condition && Array.isArray(model.list)) {
    return model
  }

  // Handle the condition/groups format
  if (model.condition && Array.isArray(model.groups)) {
    return {
      condition: model.condition,
      list: model.groups.map((group: any) => ({
        condition: group.condition || 'Or',
        list: Array.isArray(group.rules)
          ? group.rules.map((rule: any) => ({
              tag: {
                builderID: rule.builderID || '',
                params: rule.params || {}
              }
            }))
          : []
      }))
    }
  }

  // Handle the intersect/union format
  const intersectKey = 'intersect'
  const unionKey = 'union'

  if (model[intersectKey] || model[unionKey]) {
    const topLevelKey = model[intersectKey] ? intersectKey : unionKey
    const topLevelCondition = topLevelKey === intersectKey ? 'And' : 'Or'
    const groups = model[topLevelKey] || []

    return {
      condition: topLevelCondition,
      list: Array.isArray(groups)
        ? groups.map((group: any) => {
            const groupKey = group[intersectKey] ? intersectKey : unionKey
            const groupCondition = groupKey === intersectKey ? 'And' : 'Or'
            const items = group[groupKey] || []

            return {
              condition: groupCondition,
              list: Array.isArray(items)
                ? items.map((item: any) => ({
                    tag: item.tag || {
                      builderID: '',
                      params: {}
                    }
                  }))
                : []
            }
          })
        : []
    }
  }

  // Default format or compatibility with old format
  return {
    condition: 'And',
    list: [
      {
        condition: 'Or',
        list: [genRecordModel()]
      }
    ]
  }
}

export function genRecordModel() {
  return {
    tag: {
      builderID: '',
      params: {}
    }
  }
}
