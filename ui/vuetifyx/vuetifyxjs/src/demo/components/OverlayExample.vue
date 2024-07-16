<script setup lang="ts">
import { ref } from 'vue'
import Overlay from '@/lib/Overlay.vue'

const value = ref()
const iframe = ref()
const ol = ref()

const show = (e: any) => {
  ol.value.showByElement(e)
  value.value = true
}
const showCenter = () => {
  ol.value.showCenter()
  value.value = true
}

window.addEventListener('message', function (e: any) {
  if (e.data.key && e.data.key.indexOf('vue-devtools') >= 0) {
    return
  }
  if (e.data.source && e.data.source.indexOf('vue-devtools') >= 0) {
    return
  }
  ol.value.showByIframe(iframe.value, JSON.parse(e.data))
  value.value = true
})
</script>

<template>
  <v-card :height="600" width="w-100">
    <v-app>
      <v-navigation-drawer permanent>
        <v-list>
          <v-list-item v-for="i in 9" class="mt-2">
            <v-btn @click="show">show{{ i }}</v-btn>
          </v-list-item>
          <v-list-item class="mt-2">
            <v-btn @click="showCenter">show Center</v-btn>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
      <v-main>
        <v-card>
          <div ref="iframe">
            <iframe src="/iframe.html" class="w-100 overflow-y-auto" style="height: 600px"></iframe>
          </div>
        </v-card>
      </v-main>
      <overlay v-model="value" ref="ol">
        <v-card :width="400" :height="400">
          <v-btn @click="value = false"> close</v-btn>
          Hello World
        </v-card>
      </overlay>
    </v-app>
  </v-card>
</template>

<style scoped></style>
