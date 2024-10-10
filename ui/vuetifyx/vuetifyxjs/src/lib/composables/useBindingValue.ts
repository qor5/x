import { ref, watch } from 'vue'

export default function useBindingValue<EmitFnType>(props: Record<string, any>, emit: EmitFnType) {
  const bindingValue = ref(props.modelValue)

  function onUpdateModelValue(value: any) {
    ;(emit as any)('update:modelValue', value)
    bindingValue.value = value
  }

  watch(
    () => props.modelValue,
    (newVal) => {
      bindingValue.value = newVal
    }
  )

  return { bindingValue, onUpdateModelValue }
}
