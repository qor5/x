<script setup>
import RichTextRedactor from "@/lib/RichTextRedactor.vue";
import {ref} from "vue";

const value = ref("")


</script>


<template>
  <input v-model="value">
  <rich-text-redactor v-model="value" :config="{ plugins:[]}"></rich-text-redactor>
</template>

