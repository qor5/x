<template>
  <div class="vx-segment-item-wrap">
    <div class="condition-group">
      <vx-select
        v-model="selectedOption"
        :items="optionsForSelect"
        style="min-width: 150px"
        item-title="name"
        item-value="id"
        placeholder="Select a type"
        hide-details
        @update:modelValue="handleSelectChange"
      >
        <template #item="{ props, item }">
          <template v-if="item.raw.category">
            <v-list-subheader class="text-primary font-weight-medium">{{
              item.raw.category
            }}</v-list-subheader>
          </template>
          <template v-else>
            <v-list-item v-bind="props" class="pl-8">
              <template #title>
                <span>{{ item.raw.name }}</span>
              </template>
            </v-list-item>
          </template>
        </template>
      </vx-select>

      <!-- cascade select -->

      <!-- <span class="condition-text">of</span> -->
    </div>
    <v-icon class="delete-icon" color="rgb(158, 158, 158)" size="24" @click="handleRemove"
      >mdi-minus-circle-outline</v-icon
    >
  </div>
</template>

<script setup lang="ts">
import { defineEmits, inject, computed, ref, defineProps, PropType } from 'vue'
import type { OptionsType } from './type'

const segmentNestedOptions = inject<OptionsType[]>('segmentOptions', [])
const selectedOption = ref<OptionsType | null>(null)

const props = defineProps({
  modelValue: {
    type: Object as PropType<Record<string, any>>,
    default: () => ({})
  }
})

const optionsForSelect = computed(() => {
  return segmentNestedOptions.reduce<
    Array<{
      id: string
      category?: string
      name?: string
      description?: string
      categoryID?: string
    }>
  >((acc, item) => {
    if ('builders' in item) {
      acc.push({
        id: item.id,
        category: item.name,
        description: item.description
      })

      acc.push(
        ...item.builders.map((builder) => ({
          id: builder.id,
          name: builder.name,
          categoryID: builder.categoryID
        }))
      )
    }
    return acc
  }, [])
})

const emit = defineEmits(['on-remove', 'on-select'])

function handleSelectChange(value: OptionsType) {
  emit('on-select', value)
}

const handleRemove = () => {
  emit('on-remove')
}
</script>

<style scoped lang="scss">
.vx-segment-item-wrap {
  position: relative;
  background: rgb(250, 250, 250);
  border-radius: 4px;
  border: 1px solid rgb(224, 224, 224);
  padding: 8px;
  margin-right: 25px;
}
.condition-group {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.delete-icon {
  position: absolute;
  top: 50%;
  right: -30px;
  transform: translateY(-50%);
  cursor: pointer;
}

.condition-text {
  color: rgb(158, 158, 158);
  line-height: 40px;
}
</style>
