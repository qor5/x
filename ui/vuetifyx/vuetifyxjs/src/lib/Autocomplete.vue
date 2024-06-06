<script setup lang="ts">
import { computed, onMounted, onUpdated, PropType, Ref, ref } from 'vue'
import draggable from 'vuedraggable'

enum Variant {
  Filled = 'filled',
  Underlined = 'underlined',
  Outlined = 'outlined',
  Plain = 'plain',
  Solo = 'solo',
  SoloInverted = 'solo-inverted',
  SoloFilled = 'solo-filled'
}

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: { type: String },
  variant: { type: String as PropType<Variant>, default: 'underlined' },
  density: { type: String as PropType<null | 'default' | 'comfortable' | 'compact'> },
  items: { type: Array<any>, default: [] },
  cacheItems: { type: Array<any>, default: [] },
  isPaging: Boolean,
  hasIcon: Boolean,
  hideSelected: Boolean,
  hideDetails: Boolean,
  clearable: Boolean,
  chips: Boolean,
  sorting: Boolean,
  itemText: { type: String, default: 'text' },
  itemValue: { type: String, default: 'value' },
  itemIcon: { type: String, default: 'icon' },
  pageKey: { type: String, default: 'page' },
  pagesKey: { type: String, default: 'pages' },
  pageSizeKey: { type: String, default: 'pageSize' },
  totalKey: { type: String, default: 'total' },
  itemsKey: { type: String, default: 'items' },
  currentKey: { type: String, default: 'current' },
  searchKey: { type: String, default: 'search' },
  chipColor: String,
  loadData: Function,
  remote: {
    type: Object,
    default: {
      page: 0,
      pageSize: 0,
      search: ''
    }
  }
})
const listItems: Ref<Array<any>> = ref([...props.items])
const value = ref()
const cachedSelectedItems: Ref<Array<any>> = ref([...(props.cacheItems ?? [])])
const isLoading = ref(false)
const disabled = ref(false)
const total = ref(0)
const pages = ref(0)
const current = ref(0)

// key like `$.data.total` if just `total` this function will be  change key to `data.total`
// and return object multilevel value
const getObjMultiValue = (d: Object, key: string) => {
  const keys = key.split('.')
  if (keys.length === 0) {
    return d
  }
  if (keys[0] === '$') {
    keys.shift()
  } else {
    keys.unshift('data')
  }
  return getObjectValue(d, keys)
}
const getObjectValue = (d: any, keys: Array<string>): any => {
  if (typeof d !== 'object' || keys.length == 0) {
    return d
  }
  const newKey = keys[0]
  keys.shift()
  return getObjectValue(d[newKey], keys)
}
const loadRemoteItems = () => {
  if (!props.loadData) {
    return
  }

  isLoading.value = true
  props
    .loadData()
    .then((r: any) => {
      total.value = getObjMultiValue(r, props.totalKey)
      pages.value = getObjMultiValue(r, props.pagesKey)
      current.value = getObjMultiValue(r, props.currentKey)
      const items = getObjMultiValue(r, props.itemsKey)
      if (props.isPaging) {
        listItems.value = items
      } else {
        disabled.value = current.value >= total.value
        listItems.value = listItems.value.concat(items || [])
      }
    })
    .finally(() => {
      isLoading.value = false
    })
}

const endIntersect = (isIntersecting: boolean) => {
  if (isIntersecting && !disabled.value) {
    props.remote[props.pageKey] += 1
    loadRemoteItems()
  }
}

const changeStatus = (e: any) => {
  if (cachedSelectedItems.value.find((element) => element[props.itemValue] == e)) {
    return
  }
  cachedSelectedItems.value.push(listItems.value.find((element) => element[props.itemValue] == e))
  emit('update:modelValue', value.value)
}

const removeItem = (v: any) => {
  value.value = ''
  cachedSelectedItems.value = cachedSelectedItems.value.filter(
    (element) => element[props.itemValue] != v[props.itemValue]
  )
  emit('update:modelValue', value.value)
}
onMounted(() => {
  loadRemoteItems()
})
onUpdated(() => {
  value.value = props.modelValue
})

const reloadSearch = (val: any) => {
  if (!props.loadData) {
    return
  }
  if (val == props.remote[props.searchKey] || !val) {
    return
  }
  if (val == value.value[props.itemText]) {
    return
  }
  props.remote[props.pageKey] = 1
  props.remote[props.searchKey] = val
  loadRemoteItems()
}

const chipsVisible = computed(() => {
  return props.chips && props.hasIcon && !props.sorting
})
</script>

<template>
  <div>
    <v-card v-if="sorting && cachedSelectedItems.length > 0">
      <v-list>
        <draggable
          animation="300"
          handle=".handle"
          v-model="cachedSelectedItems"
          :item-key="itemValue"
        >
          <template #item="{ element }">
            <v-list-item
              v-if="hasIcon"
              :prepend-avatar="element[itemIcon]"
              :title="element[itemText]"
              animation="300"
            >
              <template v-slot:append>
                <v-icon icon="mdi-drag" class="handle mx-2 cursor-grab"></v-icon>
                <v-btn @click="removeItem(element)" variant="text" icon="mdi-delete"></v-btn>
              </template>
            </v-list-item>
            <v-list-item v-else :title="element[itemText]" animation="300">
              <template v-slot:append>
                <v-icon icon="mdi-drag" class="handle mx-2 cursor-grab"></v-icon>
                <v-btn @click="removeItem(element)" variant="text" icon="mdi-delete"></v-btn>
              </template>
            </v-list-item>
          </template>
        </draggable>
      </v-list>
    </v-card>
    <v-autocomplete
      v-model="value"
      :items="listItems"
      :loading="isLoading"
      :item-value="itemValue"
      :item-title="itemText"
      :clearable="sorting ? false : clearable"
      :hide-details="hideDetails"
      :hide-selected="hideSelected"
      :class="sorting ? 'v-autocomplete-sorting' : ''"
      @update:modelValue="changeStatus"
      :variant="variant"
      :density="density"
      @update:search="reloadSearch"
    >
      <template v-slot:item="{ item, props }" v-if="hasIcon">
        <v-list-item
          v-bind="props"
          :prepend-avatar="item.raw[itemIcon]"
          :title="item.raw[itemText]"
        ></v-list-item>
      </template>
      <template v-slot:chip="{ props, item }" v-if="chipsVisible">
        <v-chip
          v-bind="props"
          :color="chipColor"
          :prepend-avatar="hasIcon ? item.raw[itemIcon] : undefined"
          :text="item.raw[itemText]"
        ></v-chip>
      </template>
      <template v-slot:append-item="" v-if="loadData">
        <div class="text-center">
          <v-pagination
            v-if="props.isPaging"
            v-model="remote[pageKey]"
            rounded="circle"
            :length="pages"
            total-visible="5"
            @update:modelValue="loadRemoteItems()"
          ></v-pagination>
          <div v-else>
            <v-btn
              class="ma-2"
              color="primary"
              :disabled="disabled"
              :loading="isLoading"
              v-intersect="endIntersect"
              @click="
                () => {
                  remote[pageKey] += 1
                  loadRemoteItems()
                }
              "
              >Load more
            </v-btn>
            <v-divider vertical></v-divider>
            <span> {{ current }}/{{ total }} </span>
          </div>
        </div>
      </template>
    </v-autocomplete>
  </div>
</template>

<style scoped></style>
