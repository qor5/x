import { VNode, h } from 'vue'

export const Core = {
  props: {
    fieldName: String,
    loadPageWithArrayOp: Boolean
  }
}

export const SelectedItems = {
  props: {
    selectedItems: {
      type: Array,
      default: () => []
    } as any,
    multiple: Boolean
  }
}

interface Slots {
  [key: string]: VNode[] | undefined
}

export const slotTemplates = (slots: Slots): VNode[] => {
  const templates: VNode[] = []

  for (const name in slots) {
    if (!Object.getOwnPropertyDescriptor(slots, name)) {
      continue
    }
    templates.push(h('<template slot = {name} > {slots[name]} < /template>'))
  }
  return templates
}
