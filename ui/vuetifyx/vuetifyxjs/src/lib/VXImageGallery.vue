<template>
  <v-sheet class="mx-auto" elevation="0" max-width="100%">
    <v-slide-group v-model="model" class="pa-4" show-arrows>
      <v-slide-group-item v-for="(item, i) in items" :key="i" v-slot="{ isSelected, toggle }">
        <v-card
          :color="isSelected ? 'primary' : 'grey-lighten-1'"
          class="ma-4"
          :height="height"
          :width="width || 300"
          @click="toggle"
        >
          <div class="d-flex fill-height align-center justify-center">
            <v-img :src="item.src" :lazy-src="item.lazySrc" cover height="100%" width="100%">
              <template v-slot:placeholder>
                <div class="d-flex align-center justify-center fill-height">
                  <v-progress-circular color="grey-lighten-4" indeterminate></v-progress-circular>
                </div>
              </template>
            </v-img>
          </div>
        </v-card>
      </v-slide-group-item>
    </v-slide-group>
  </v-sheet>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface GalleryItem {
  src: string
  lazySrc?: string
  title?: string
}

const props = defineProps<{
  items: GalleryItem[]
  height?: number | string
  width?: number | string
}>()

const model = ref(null)
</script>
