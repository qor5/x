import { computed, useAttrs } from 'vue'

interface FilteredAttrs {
  attrs: Record<string, any>
  filteredAttrs: Record<string, any>
  rootAttrs: Record<string, any>
}

/**
 * class and id are supposed to bind to root element,
 * other props should bind to child components
 */
export function useFilteredAttrs(): FilteredAttrs {
  const attrs = useAttrs()

  // filter class and id, should not binding to components props
  const filteredAttrs = computed(() => {
    const { class: _class, id: _id, style: _style, ...rest } = attrs
    return rest
  })

  const rootAttrs = computed(() => {
    const { class: _class, id: _id, style: _style } = attrs
    return { class: _class, id: _id, style: _style }
  })

  return {
    attrs,
    filteredAttrs,
    rootAttrs
  }
}
