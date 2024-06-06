<script setup lang="ts">
import Cropper from "@/lib/Cropper.vue";
import { ref } from "vue";
const value = ref();
const src =
  "https://cdn.pixabay.com/photo/2023/11/23/20/40/ocean-8408693_1280.jpg";
</script>

<template>
  <p>{{ value }}</p>
  <cropper v-model="value" :src="src"></cropper>
</template>

<style scoped></style>
