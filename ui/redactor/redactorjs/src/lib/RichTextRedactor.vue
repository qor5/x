<script setup>
/*
    Redactor Vue Component
    Version 1.2
    Updated: June 9, 2021

    http://imperavi.com/redactor/

    Copyright 2021, Imperavi Ltd.
    License: MIT
*/
import { computed, onMounted, onUnmounted, ref, watch } from "vue";
import "./redactor.min.css";
import "./redactor.min.js";

const Redactor = window.Redactor;

const redactor = ref();
const app = ref();
const emit = defineEmits(["update:modelValue"]);
const props = defineProps({
  modelValue: {
    default: "",
    type: String,
  },
  placeholder: {
    type: String,
    default: null,
  },
  name: {
    type: String,
    default: null,
  },
  config: {
    default: {
      callbacks: {},
    },
    type: Object,
  },
});
const value = computed(() => {
  return props.modelValue;
});
onMounted(() => {
  props.config.callbacks = {
    changed: function (html) {
      emit("update:modelValue", html);
      return html;
    },
  };
  app.value = Redactor(redactor.value, props.config);
});
onUnmounted(() => {
  // Call destroy on redactor to cleanup event handlers
  Redactor(redactor.value, "destroy");

  // unset instance for garbage collection
  redactor.value = null;
  // this.$parent.redactor = null;
});

watch(value, (newValue, oldValue) => {
  if (app.value?.editor.isFocus() || app.value?.editor.isSourceMode()) {
    return;
  }
  app.value?.source.setCode(newValue);
});
</script>

<template>
  <textarea
    ref="redactor"
    :name="name"
    :placeholder="placeholder"
    :value="value"
  />
</template>
