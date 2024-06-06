<script setup lang="ts">
import Editor from "@/lib/Editor.vue";
import { ref } from "vue";

const value = ref("<p>I’m running Tiptap with Vue.js. 🎉</p>");
</script>

<template>
  <p>{{ value }}</p>
  <editor v-model="value"></editor>
</template>

<style scoped></style>
