import { ref, watch } from 'vue'

export default function useBindingValue<EmitFnType>(
  props: Record<string, any>,
  emit: EmitFnType,
  modelValueFormatter?: (value: any) => any
) {
  const bindingValue = ref(
    modelValueFormatter ? modelValueFormatter(props.modelValue) : props.modelValue
  )
  const bindingFocus = ref(false)

  function onUpdateModelValue(value: any) {
    ;(emit as any)('update:modelValue', value)
    bindingValue.value = value
  }

  function onUpdateFocused(value: boolean) {
    ;(emit as any)('update:focused', value)
    bindingFocus.value = value
  }

  watch(
    () => props.modelValue,
    (newVal) => {
      bindingValue.value = modelValueFormatter ? modelValueFormatter(newVal) : newVal
    }
  )

  return { bindingValue, bindingFocus, onUpdateModelValue, onUpdateFocused }
}
