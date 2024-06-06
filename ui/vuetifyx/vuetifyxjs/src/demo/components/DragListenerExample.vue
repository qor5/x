<script setup lang="ts">
import { ref } from 'vue'
import DragListener from '@/lib/DragListener.vue'

const iframe = ref()
const list = ref([
  { name: 'test1' },
  { name: 'test2' },
  { name: 'test3' },
  { name: 'test4' },
  { name: 'test5' }
])

const drop = (e: any) => {
  console.log('drop', e)
}
</script>

<template>
  <drag-listener @drop="drop">
    <div v-for="item in list" draggable="true">{{ item.name }}</div>
    <div id="vx-drag-target-area">
      <iframe src="https://demo.qor5.theplant-dev.com/" style="pointer-events: none"></iframe>
    </div>
  </drag-listener>
</template>

<style scoped>
#vx-drag-target-area {
  width: 100px;
  height: 100px;
  background-color: red;
}
</style>
