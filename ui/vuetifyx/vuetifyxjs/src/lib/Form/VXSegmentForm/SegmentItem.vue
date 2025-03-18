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
        :error-messages="shouldShowError ? errorMessages : ''"
        :hide-details="!shouldShowError"
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
      <template v-for="(fragment, index) in visibleFragments" :key="getItemKey(fragment, index)">
        <span v-if="fragment.type === 'TEXT'" class="condition-text">{{
          fragment.text || ''
        }}</span>

        <template v-else-if="fragment.type === 'SELECT'">
          <vx-select
            v-if="fragment.multiple"
            v-model="tagParams[fragment.key]"
            item-title="label"
            item-value="value"
            style="min-width: 150px"
            :placeholder="'Select values'"
            :items="fragment.options"
            multiple
            hide-details
            @blur="handleFragmentValueChange(fragment.key, tagParams[fragment.key])"
          />

          <vx-select
            v-else
            v-model="tagParams[fragment.key]"
            item-title="label"
            item-value="value"
            style="min-width: 150px"
            :placeholder="'Select a value'"
            :items="fragment.options"
            hide-details
            @update:modelValue="handleFragmentValueChange(fragment.key, tagParams[fragment.key])"
          />
        </template>

        <vx-field
          v-else-if="fragment.type === 'NUMBER_INPUT'"
          type="number"
          v-model="tagParams[fragment.key]"
          style="min-width: 70px"
          hide-details
          @mouseleave="debouncedHandleFragmentValueChange(fragment.key, tagParams[fragment.key])"
        />
        <vx-date-picker
          v-else-if="fragment.type === 'DATE_PICKER'"
          :type="fragment.includeTime ? 'datetimepicker' : 'datepicker'"
          v-model="tagParams[fragment.key]"
          :style="fragment.includeTime ? 'min-width: 220px' : 'min-width:150px'"
          :placeholder="fragment.includeTime ? 'Select a datetime' : 'Select a date'"
          hide-details
          @blur="handleFragmentValueChange(fragment.key, tagParams[fragment.key])"
        />
      </template>
    </div>
    <v-icon class="delete-icon" color="rgb(158, 158, 158)" size="24" @click="handleRemove"
      >mdi-minus-circle-outline</v-icon
    >
  </div>
</template>

<script setup lang="ts">
import { defineEmits, inject, computed, ref, defineProps, PropType, watch, reactive } from 'vue'
import type { OptionsType } from './type'
import isEqual from 'lodash/isEqual' // Import lodash isEqual method
import { useItemKeys } from './useUtils' // 引入useItemKeys
import { debounce } from 'lodash'

// Extended FragmentType interface to include all possible properties
interface ExtendedFragmentType {
  defaultValue: any
  key: string
  multiple?: boolean
  options?: Array<{ label: string; value: string }>
  required: boolean
  skipIf: null | Record<string, any>
  skipUnless: null | Record<string, any>
  type: 'SELECT' | 'DATE_PICKER' | 'NUMBER_INPUT' | 'TEXT'
  validation: null | string
  text?: string
  includeTime?: boolean
}

const segmentNestedOptions = inject<OptionsType[]>('segmentOptions', [])
const selectedOption = ref<string | null>(null)
const { getItemKey } = useItemKeys()

const props = defineProps({
  modelValue: {
    type: Object as PropType<Record<string, any>>,
    default: () => ({})
  },
  validate: {
    type: Boolean,
    default: false
  }
})

// Store parameters for the current tag instance
const tagParams = reactive<Record<string, any>>({})

// Initialize tag params and watch model changes
watch(
  () => props.modelValue,
  (newVal) => {
    if (newVal && newVal.tag) {
      selectedOption.value = newVal.tag.builderID || null

      // Initialize tag parameters from model
      if (newVal.tag.params && Object.keys(newVal.tag.params).length > 0) {
        Object.keys(newVal.tag.params).forEach((key) => {
          tagParams[key] = newVal.tag.params[key]
        })
      }
    }
  },
  { deep: true, immediate: true }
)

const compMap = ref<Record<string, ExtendedFragmentType[]>>({})

// Get current selected builder's fragments
const currentFragments = computed(() => {
  const key = selectedOption.value || ''
  return compMap.value[key] || []
})

// Visible fragments filtered by skipIf and skipUnless conditions
const visibleFragments = computed(() => {
  return currentFragments.value.filter((fragment) => {
    if (fragment.skipIf && skipIf(fragment.skipIf)) return false
    if (fragment.skipUnless && !skipIf(fragment.skipUnless)) return false
    return true
  })

  function skipIf(conditionObj: Record<string, any> | null): boolean {
    if (!conditionObj || Object.keys(conditionObj).length === 0) return false

    return Object.entries(conditionObj).every(([key, condition]) => {
      if (!key.startsWith('$')) {
        key = '$' + key
        condition = { EQ: condition }
      }
      const fieldName = key.substring(1)
      const fieldValue = tagParams[fieldName] === undefined ? null : tagParams[fieldName]

      if (!condition || typeof condition !== 'object' || Object.keys(condition).length === 0)
        return false

      return Object.entries(condition).every(([operator, value]) => {
        switch (operator) {
          case 'IN':
            if (Array.isArray(value)) {
              return value.some((item) => isEqual(fieldValue, item))
            }
            return false
          case 'EQ':
            return isEqual(fieldValue, value)
          default:
            return false
        }
      })
    })
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
        ...item.builders.map((builder) => {
          compMap.value[builder.id] = builder.view.fragments as ExtendedFragmentType[]

          return {
            id: builder.id,
            name: builder.name,
            categoryID: builder.categoryID
          }
        })
      )
    }
    return acc
  }, [])
})

const emit = defineEmits(['on-remove', 'on-select', 'update:modelValue'])

const shouldShowError = computed(() => {
  return props.validate && !selectedOption.value && userInteracted.value
})

// 跟踪用户是否已与该字段交互
const userInteracted = ref(false)

// Handle builder selection
function handleSelectChange(value: string) {
  userInteracted.value = true
  emit('on-select', value)

  // Reset tag parameters
  Object.keys(tagParams).forEach((key) => {
    delete tagParams[key]
  })

  // Initialize parameters with default values from fragments
  if (compMap.value[value]) {
    compMap.value[value].forEach((fragment) => {
      if (fragment.key) {
        tagParams[fragment.key] = fragment.defaultValue
      }
    })
  }

  // Update model
  updateModel()
}

// Handle fragment value changes
function handleFragmentValueChange(key: string, value: any) {
  console.log('handleFragmentValueChange', key, value)
  tagParams[key] = value
  updateModel()
}

// Update the entire model
function updateModel() {
  if (!selectedOption.value) return

  const updatedModel = {
    ...props.modelValue,
    tag: {
      builderID: selectedOption.value,
      params: { ...tagParams }
    }
  }

  emit('update:modelValue', updatedModel)
}

const handleRemove = () => {
  emit('on-remove')
}

function triggerValidation() {
  userInteracted.value = true
  return !!selectedOption.value
}

const errorMessages = computed(() => {
  return !selectedOption.value ? 'Type cannot be empty' : ''
})

// Expose validation method
defineExpose({
  isValid: () => {
    userInteracted.value = true
    return !!selectedOption.value
  }
})

const debouncedHandleFragmentValueChange = debounce((key: string, value: any) => {
  handleFragmentValueChange(key, value)
}, 0)
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
