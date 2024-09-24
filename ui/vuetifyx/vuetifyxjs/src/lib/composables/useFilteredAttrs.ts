import { computed, useAttrs } from 'vue'

interface FilteredAttrs {
  filteredAttrs: Record<string, any>
}

/**
 * class and id are supposed to bind to root element,
 * other props should bind to child components
 * */
export function useFilteredAttrs(): FilteredAttrs {
  const attrs = useAttrs()

  // filter class and id, should not binding to components props
  const filteredAttrs: Record<string, any> = computed(() => {
    const { class: _class, id: _id, style: _style, ...rest } = attrs
    return rest
  })

  return {
    filteredAttrs
  }
}
