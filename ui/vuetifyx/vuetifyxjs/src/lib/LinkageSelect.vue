<template>
  <div>
    <v-row v-if="row">
      <v-col v-for="(v, i) in linkageSelectItems" :key="i">
        <v-autocomplete
          :label="labels[i]"
          :items="levelItems[i]"
          item-title="Name"
          item-value="ID"
          v-model="selectedIDs[i]"
          @update:modelValue="selectItem($event, i)"
          :clearable="!chips"
          :error-messages="errorMessages?.[i]"
          :chips="chips"
          :disabled="disabled"
          :hide-details="hideDetails"
          variant="underlined"
        >
        </v-autocomplete>
      </v-col>
    </v-row>
    <div v-else>
      <v-autocomplete
        v-for="(v, i) in linkageSelectItems"
        :label="labels[i]"
        :items="levelItems[i]"
        item-value="ID"
        item-title="Name"
        v-model="selectedIDs[i]"
        variant="underlined"
        @update:modelValue="selectItem($event, i)"
        :clearable="!chips"
        :error-messages="errorMessages?.[i]"
        :chips="chips"
        :disabled="disabled"
        :hide-details="hideDetails"
      >
      </v-autocomplete>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, Ref, watch } from 'vue'

interface LinkageSelectItem {
  ID: string
  Name: string
  ChildrenIDs: string[]
}

const props = defineProps<{
  modelValue: string[] | undefined
  items: LinkageSelectItem[][]
  labels: string[]
  errorMessages?: string[]
  disabled?: boolean
  selectOutOfOrder?: boolean
  chips?: boolean
  row?: boolean
  hideDetails?: boolean
}>()

const emit = defineEmits(['update:modelValue'])
const linkageSelectItems = ref([...props.items])

const levelItems: Ref<any[]> = ref([])

const selectedIDs = computed(() => {
  let ids: Array<string> = [...(props.modelValue ?? [])]
  validateAndResetSelectedIDs(ids)
  return ids
})

onMounted(() => {
  for (let i = 0; i < props.labels.length; i++) {
    levelItems.value.push(getLevelItems(i))
  }
  watch(selectedIDs, () => {
    for (let i = 0; i < props.labels.length; i++) {
      levelItems.value[i] = getLevelItems(i)
    }
  })
})
linkageSelectItems.value.forEach((v: any) => {
  v.forEach((item: LinkageSelectItem) => {
    if (!item.Name) {
      item.Name = item.ID
    }
  })
})

const validateAndResetSelectedIDs = (ids: Array<string>) => {
  linkageSelectItems.value.forEach((v: any, i: number) => {
    if (!ids[i]) {
      ids[i] = ''
    }
  })
  ids.forEach((v, i) => {
    if (!v) {
      ids[i] = ''
      return
    }

    var exists = false
    for (var item of linkageSelectItems.value[i]) {
      if (item.ID === v) {
        exists = true
        break
      }
    }
    if (!exists) {
      ids[i] = ''
      return
    }

    if (i === 0) {
      return
    }
    var pID = ids[i - 1]
    if (!pID) {
      if (!props.selectOutOfOrder) {
        ids[i] = ''
      }
      return
    } else {
      for (const item of linkageSelectItems.value[i - 1]) {
        if (item.ID === pID) {
          for (var id of item.ChildrenIDs) {
            if (id === v) {
              return
            }
          }
        }
      }
    }

    ids[i] = ''
    return
  })
}

const getLevelItems = (level: number): LinkageSelectItem[] => {
  if (level === 0) {
    return linkageSelectItems.value[level]
  }
  let items: LinkageSelectItem[] = []
  if (selectedIDs.value[level - 1]) {
    let idM: any = {}
    for (const item of linkageSelectItems.value[level - 1]) {
      if (item.ID === selectedIDs.value[level - 1]) {
        for (let id of item.ChildrenIDs) {
          idM[id] = true
        }
        break
      }
    }
    for (const item of linkageSelectItems.value[level]) {
      if (idM[item.ID]) {
        items.push(item)
      }
    }
    return items
  }

  if (props.selectOutOfOrder) {
    for (let i = level - 2; i >= 0; i--) {
      if (selectedIDs.value[i]) {
        items = findNextItems(selectedIDs.value[i], i)
        for (let j = i + 1; j < level; j++) {
          let newItems: Array<LinkageSelectItem> = []
          for (const item of items) {
            newItems = newItems.concat(findNextItems(item.ID, j))
          }
          items = newItems
        }

        return items
      }
    }
    return props.items[level]
  }
  return []
}

const selectItem = (v: string, level: number) => {
  let updateSelectIds = [...selectedIDs.value]
  if (v) {
    for (var i = level + 1; i < updateSelectIds.length; i++) {
      if (updateSelectIds[i]) {
        var items = getLevelItems(i)
        if (!items || items.length === 0) {
          updateSelectIds[i] = ''
          continue
        }
        var found = false
        for (var item of items) {
          if (item.ID === updateSelectIds[i]) {
            found = true
            break
          }
        }
        if (!found) {
          updateSelectIds[i] = ''
        }
      }
    }
  } else {
    updateSelectIds[level] = ''
    if (!props.selectOutOfOrder) {
      for (let i = level + 1; i < updateSelectIds.length; i++) {
        updateSelectIds[i] = ''
      }
    }
  }
  if (props.labels.length > level + 1) {
    levelItems.value[level + 1] = getLevelItems(level + 1)
  }
  if (updateSelectIds.every((x) => !x)) {
    emit('update:modelValue', [])
    return
  }
  emit('update:modelValue', updateSelectIds)
}

const findNextItems = (selectedID: any, level: number): LinkageSelectItem[] => {
  if (level + 1 >= linkageSelectItems.value.length) {
    return []
  }
  var childrenIDs: string[] = []
  for (const item of linkageSelectItems.value[level]) {
    if (item.ID === selectedID) {
      childrenIDs = item.ChildrenIDs
      break
    }
  }
  if (childrenIDs.length == 0) {
    return []
  }
  var items = []
  for (const item of linkageSelectItems.value[level + 1]) {
    if (childrenIDs.includes(item.ID)) {
      items.push(item)
    }
  }
  return items
}
</script>
