<script setup lang="ts">
import { onMounted, reactive, ref, watch } from 'vue'
import get from 'lodash/get'

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: {
    type: Object
  },
  isPaging: Boolean,
  hideDetails: Boolean,
  chips: Boolean,
  level: Number,
  label: String,
  selectOutOfOrder: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  parentField: { type: String },
  levelField: { type: String },
  parentValue: { type: String },
  parentIdField: { type: String },
  itemTitle: { type: String },
  itemValue: { type: String },
  pageField: { type: String },
  pagesField: { type: String },
  pageSizeField: { type: String },
  totalField: { type: String },
  itemsField: { type: String },
  currentField: { type: String },
  searchField: { type: String },
  remoteUrl: String,
  levelStart: { type: Number, default: 0 },
  levelStep: { type: Number, default: 1 },
  page: { type: Number, default: 1 },
  pageSize: { type: Number, default: 20 },
  errorMessage: { type: String, default: '' }
})

const value = ref()
const listItems = ref([])
const isLoading = ref(false)
const loadMoreDisabled = ref(false)
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
  urlObj.searchParams.set(props.pageField!, pagination.page.toString())
  urlObj.searchParams.set(
    props.levelField!,
    (props.level! * props.levelStep! + props.levelStart!).toString()
  )
  urlObj.searchParams.set(props.pageSizeField!, pagination.pageSize.toString())
  if (props.parentValue) {
    urlObj.searchParams.set(props.parentIdField!, props.parentValue)
  }
  if (pagination.search) {
    urlObj.searchParams.set(props.searchField!, pagination.search)
  }
  return fetch(urlObj.toString())
}
onMounted(() => {
  value.value = props.modelValue
})
watch(
  () => props.modelValue,
  () => {
    value.value = props.modelValue
    if (!props.modelValue) {
      listItems.value.splice(0, listItems.value.length)
    }
  }
)

const loadRemoteItems = () => {
  if (!url.value) {
    return
  }
  if (!props.selectOutOfOrder && props.level != 0 && !props.parentValue) {
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
      total.value = get(r, props.totalField!)
      pages.value = get(r, props.pagesField!)
      current.value = get(r, props.currentField!)
      let items = get(r, props.itemsField!)
      items = items ?? []
      if (props.isPaging) {
        listItems.value = items
      } else {
        loadMoreDisabled.value = current.value >= total.value
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
  if (isIntersecting && !loadMoreDisabled.value && listItems.value.length > 0) {
    pagination.page += 1
    loadRemoteItems()
  }
}

const changeStatus = (e: any) => {
  if (!e) {
    emit('update:modelValue', undefined)
    return
  }
  emit('update:modelValue', e)
}

const reloadSearch = (val: any) => {
  if (!val) {
    return
  }
  searchData(val)
}
const searchData = (val: any) => {
  if (!url.value) {
    return
  }
  if (!val) {
    val = ''
  }
  if (value.value && val == value.value[props.itemValue!]) {
    return
  }
  pagination.search = val
  pagination.page = 1
  loadRemoteItems()
}
</script>

<template>
  <div>
    <v-autocomplete
      return-object
      v-model="value"
      :items="listItems"
      :label="label"
      :loading="isLoading"
      :chips="chips"
      :disabled="disabled"
      :item-title="itemTitle"
      :item-value="itemValue"
      :clearable="!chips"
      :hide-details="hideDetails"
      @update:modelValue="changeStatus"
      @click="searchData('')"
      variant="underlined"
      @update:search="reloadSearch"
      :error-messages="errorMessage"
    >
      <template v-slot:append-item v-if="remoteUrl">
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
              :disabled="loadMoreDisabled"
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
