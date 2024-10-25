<script setup lang="ts">
import { FilterItem } from '@/lib/Filter/Model'
import FilterButton from '@/lib/Filter/components/FilterButton.vue'
import { inject, ref, Ref, watch } from 'vue'
import cloneDeep from 'lodash/cloneDeep'

const props = defineProps<{
  modelValue: FilterItem
  isFoldedItem: boolean
  itemComp: any
  translations: any
  compTranslations: any
  internalValue: any
  index: number
  indexKey: string
}>()
const value = ref(cloneDeep(props.modelValue))
const menu = ref(false)
const emit = defineEmits(['update:modelValue', 'change', 'clear'])
const openMenu = inject<(val: string) => void>('openMenu')
const currentOpenMenu = inject<Ref<string>>('currentOpenMenu', ref(''))

watch(
  () => menu.value,
  (isOpen) => {
    if (isOpen) {
      if (openMenu) {
        openMenu(props.indexKey)
      }
      value.value = cloneDeep(props.modelValue)
    }
  }
)
watch(currentOpenMenu, (key) => {
  if (key != props.indexKey) {
    menu.value = false
  }
})

const clickDone = () => {
  menu.value = false
  if (
    !value.value.valueIs &&
    (!value.value.valuesAre || value.value.valuesAre.length == 0) &&
    !value.value.valueFrom &&
    !value.value.valueTo
  ) {
    value.value.selected = false
    Object.assign(props.modelValue, value.value)
    emit('update:modelValue', props.modelValue)
    emit('change', null)
    return
  }
  value.value.selected = true
  Object.assign(props.modelValue, value.value)
  emit('update:modelValue', props.modelValue)
  emit('change', null)
}

const clear = (e: any) => {
  if (!value.value.selected) {
    return
  }
  value.value.valueIs = ''
  value.value.valuesAre = []
  value.value.valueFrom = ''
  value.value.valueTo = ''
  value.value.selected = false
  Object.assign(props.modelValue, value.value)
  emit('update:modelValue', props.modelValue)
  emit('clear', e)
}
</script>

<template>
  <v-menu :close-on-content-click="false" class="rounded-lg" v-model="menu">
    <template v-slot:activator="{ props: menuProps }">
      <filter-button
        :op="modelValue"
        :is-folded-item="isFoldedItem"
        :slotProps="menuProps"
        @clear="clear"
      />
    </template>
    <v-card class="pa-3 bg-white">
      <div>{{ modelValue.translations?.filterBy }}</div>
      <component v-model="value" :is="itemComp" :translations="compTranslations"></component>
      <div>
        <v-btn class="mt-5 float-right" color="primary" rounded @click="clickDone"
          >{{ translations.apply }}
        </v-btn>
      </div>
    </v-card>
  </v-menu>
</template>

<style scoped></style>
