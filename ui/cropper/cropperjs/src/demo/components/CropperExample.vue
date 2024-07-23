<script setup lang="ts">
import Cropper from "@/lib/Cropper.vue";
import { ref } from "vue";

const value = ref();
const show = ref();
const src =
  "https://cdn.pixabay.com/photo/2023/11/23/20/40/ocean-8408693_1280.jpg";
</script>

<template>
  <p>{{ value }}</p>
  <button @click="show = !show">click me</button>
  <div v-if="show" class="box">
    <cropper
      v-model="value"
      :src="src"
      :view-mode="3"
      :auto-crop-area="1.0"
    ></cropper>
  </div>
</template>

<style scoped>
.box {
  height: 400px;
  width: 400px;
}
</style>
