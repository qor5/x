<template>
  <label class="v-label theme--light" v-html="sanitizeHtml(label)"></label>
  <v-card v-if="internalSelectedItems.length > 0" variant="flat" class="mb-2">
    <v-list>
      <draggable
        v-model="internalSelectedItems"
        :item-key="itemValue"
        @change="changeOrder"
        handle=".handle"
      >
        <template #item="{ element }">
          <v-list-item
            :prepend-avatar="element[itemImage]"
            :title="element[itemText]"
            animation="300"
          >
            <template v-slot:append>
              <v-icon icon="mdi-drag" class="handle mx-2 cursor-grab"></v-icon>
              <v-btn @click="removeItem(element[itemValue])" variant="text" icon="mdi-delete">
              </v-btn>
            </template>
          </v-list-item>
        </template>
      </draggable>
    </v-list>
  </v-card>

  <v-autocomplete
    :item-value="itemValue"
    :item-title="itemText"
    :items="internalItems"
    :label="addItemLabel"
    v-model="autocompleteValue"
    auto-select-first
    @update:modelValue="addItem"
    :loading="isLoading"
    :no-filter="noFilter"
    return-object
    variant="underlined"
  >
    <template v-slot:item="{ props, item }">
      <v-list-item
        v-bind="props"
        :prepend-avatar="item.raw[itemImage]"
        :title="item.raw[itemText]"
      ></v-list-item>
    </template>
  </v-autocomplete>
</template>
<script setup lang="ts">
import draggable from 'vuedraggable'
import DOMPurify from 'dompurify'
import { onMounted, Ref, ref } from 'vue'

function sanitizeHtml(html: string) {
  return html
    ? DOMPurify.sanitize(html, {
        ALLOWED_TAGS: ['span', 'p', 'br', 'strong', 'b', 'em', 'i', 'a'],
        ALLOWED_ATTR: ['class', 'href', 'target', 'rel']
      })
    : ''
}

const props = defineProps({
  items: {
    type: Array<any>,
    default: () => []
  },
  selectedItems: {
    type: Array<any>,
    default: () => []
  },
  itemValue: {
    type: String,
    default: 'id'
  },
  itemText: {
    type: String,
    default: 'text'
  },
  itemImage: {
    type: String,
    default: 'image'
  },
  label: {
    type: String,
    default: ''
  },
  addItemLabel: {
    type: String,
    default: ''
  },
  modelValue: {
    type: Array<any>,
    default: []
  }
})
const internalSelectedItems: Ref<any[]> = ref([])
const internalItems: Ref<any[]> = ref([])
const autocompleteValue: Ref<any[]> = ref([])
const isLoading = ref(false)
const noFilter = ref(false)

const emit = defineEmits(['update:modelValue'])

onMounted(() => {
  // internalSelectedItems.value = props.selectedItems
  internalItems.value = props.items
  internalSelectedItems.value = props.modelValue.map((id) => {
    return props.items.find((item) => item[props.itemValue] === id)
  })
})

// methods
const addItem = (event: any) => {
  autocompleteValue.value = []
  if (
    internalSelectedItems.value.find(
      (element) => element[props.itemValue] == event[props.itemValue]
    )
  ) {
    return
  }
  internalSelectedItems.value.push(
    internalItems.value.find((element) => element[props.itemValue] == event[props.itemValue])
  )
  setValue()
}
const changeOrder = (event: any) => {
  setValue()
}
const removeItem = (id: string) => {
  internalSelectedItems.value = internalSelectedItems.value.filter(
    (element) => element[props.itemValue] != id
  )
  setValue()
}
const setValue = () => {
  emit(
    'update:modelValue',
    internalSelectedItems.value.map((i) => {
      return i[props.itemValue]
    })
  )
}
</script>

<style scoped></style>
