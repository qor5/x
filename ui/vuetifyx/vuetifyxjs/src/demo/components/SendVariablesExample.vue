<script setup lang="ts">
import Seo from '@/lib/SendVariables.vue'
import { ref } from 'vue'

const seo = ref()
const value = ref('1eiuh1ihaubfiad')
</script>

<template>
  <seo ref="seo">
    <v-chip @click="seo.addTags('Test')" :label="true" :outlined="true" variant="outlined">
      <v-icon icon="mdi-plus"></v-icon>
      Test
    </v-chip>
    <v-text-field
      counter
      v-model="value"
      label="Title"
      variant="underlined"
      @focus="seo.tagInputsFocus($refs.setting_title)"
      ref="setting_title"
    >
    </v-text-field>
    <v-textarea
      counter
      v-model="value"
      label="textarea"
      variant="underlined"
      @focus="seo.tagInputsFocus($refs.textarea)"
      ref="textarea"
    >
    </v-textarea>
  </seo>
</template>

<style scoped></style>
