<script setup lang="ts">
import { computed, onMounted, onUpdated, PropType, reactive, Ref, ref } from 'vue'
import draggable from 'vuedraggable'
import get from 'lodash/get'

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
  itemTitle: { type: String, default: 'text' },
  itemValue: { type: String, default: 'value' },
  itemIcon: { type: String, default: 'icon' },
  pageField: { type: String, default: 'page' },
  pagesField: { type: String, default: 'pages' },
  pageSizeField: { type: String, default: 'pageSize' },
  totalField: { type: String, default: 'total' },
  itemsField: { type: String, default: 'items' },
  currentField: { type: String, default: 'current' },
  searchField: { type: String, default: 'search' },
  chipColor: String,
  remoteUrl: String,
  page: { type: Number, default: 1 },
  pageSize: { type: Number, default: 20 },
  errorMessages: { type: String || Array<String> || null }
})
const listItems: Ref<Array<any>> = ref([...props.items])
const value = ref()
const cachedSelectedItems: Ref<Array<any>> = ref([...(props.cacheItems ?? [])])
const isLoading = ref(false)
const disabled = ref(false)
const total = ref(0)
const pages = ref(0)
const current = ref(0)
const url = ref(props.remoteUrl)
const pagination = reactive({
  page: props.page,
  pageSize: props.pageSize,
  search: ''
})

const loadData = () => {
  let urlObj
  if (!url.value!.startsWith('http')) {
    urlObj = new URL(url.value!, `${window.location.protocol}//${window.location.host}`)
  } else {
    urlObj = new URL(url.value!)
  }
  urlObj.searchParams.set(props.pageField, pagination.page.toString())
  urlObj.searchParams.set(props.pageSizeField, pagination.pageSize.toString())
  if (pagination.search) {
    urlObj.searchParams.set(props.searchField, pagination.search)
  }
  return fetch(urlObj.toString())
}
const loadRemoteItems = () => {
  if (!url.value) {
    return
  }
  isLoading.value = true

  loadData()
    .then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok ' + response.statusText)
      }
      return response.json()
    })
    .then((r) => {
      total.value = get(r, props.totalField)
      pages.value = get(r, props.pagesField)
      current.value = get(r, props.currentField)
      const items = get(r, props.itemsField)
      if (props.isPaging) {
        listItems.value = items
      } else {
        disabled.value = current.value >= total.value
        if (pagination.page == props.page) {
          listItems.value = items
        } else {
          listItems.value = listItems.value.concat(items || [])
        }
      }
    })
    .finally(() => {
      isLoading.value = false
    })
}

const endIntersect = (isIntersecting: boolean) => {
  if (isIntersecting && !disabled.value && listItems.value.length > 0) {
    pagination.page += 1
    loadRemoteItems()
  }
}

const changeStatus = (e: any) => {
  emit('update:modelValue', e)
  if (!e) {
    return
  }
  if (cachedSelectedItems.value.find((element) => element[props.itemValue] == e[props.itemValue])) {
    return
  }
  cachedSelectedItems.value.push(
    listItems.value.find((element) => element[props.itemValue] == e[props.itemValue])
  )
}

const removeItem = (v: any) => {
  cachedSelectedItems.value = cachedSelectedItems.value.filter(
    (element) => element[props.itemValue] != v[props.itemValue]
  )
  if (v[props.itemValue] == value.value[props.itemValue]) {
    emit('update:modelValue', null)
  }
}
onMounted(() => {
  loadRemoteItems()
})
onUpdated(() => {
  value.value = props.modelValue
})

const reloadSearch = (val: any) => {
  if (!url.value) {
    return
  }
  if (!val) {
    val = ''
  }
  if (value.value && val == value.value[props.itemValue]) {
    return
  }
  pagination.search = val
  pagination.page = 1
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
              :title="element[itemTitle]"
              animation="300"
            >
              <template v-slot:append>
                <v-icon icon="mdi-drag" class="handle mx-2 cursor-grab"></v-icon>
                <v-btn @click="removeItem(element)" variant="text" icon="mdi-delete"></v-btn>
              </template>
            </v-list-item>
            <v-list-item v-else :title="element[itemTitle]" animation="300">
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
      return-object
      v-model="value"
      :items="listItems"
      :loading="isLoading"
      :item-value="itemValue"
      :item-title="itemTitle"
      :clearable="sorting ? false : clearable"
      :hide-details="hideDetails"
      :hide-selected="hideSelected"
      :class="sorting ? 'v-autocomplete-sorting' : ''"
      @update:modelValue="changeStatus"
      :variant="variant"
      :density="density"
      @update:search="reloadSearch"
      :error-messages="props.errorMessages"
    >
      <template v-slot:item="{ item, props }" v-if="hasIcon">
        <v-list-item
          v-bind="props"
          :prepend-avatar="item.raw[itemIcon]"
          :title="item.raw[itemTitle]"
        ></v-list-item>
      </template>
      <template v-slot:chip="{ props, item }" v-if="chipsVisible">
        <v-chip
          v-bind="props"
          :color="chipColor"
          :prepend-avatar="hasIcon ? item.raw[itemIcon] : undefined"
          :text="item.raw[itemTitle]"
        ></v-chip>
      </template>
      <template v-slot:append-item="" v-if="remoteUrl">
        <div class="text-center">
          <v-pagination
            size="20"
            v-if="props.isPaging"
            v-model="pagination.page"
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
                  pagination.page += 1
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
