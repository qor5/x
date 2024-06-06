<template>
  <div class="icon" :class="[`icon--${name}`, `icon--${size}`]">
    <component :is="icon"/>
  </div>
</template>

<script setup lang="ts">
import {defineAsyncComponent} from "vue";

const props = defineProps({
  name: {type: String},
  size: {
    default: "normal",
  },
  modifier: {
    default: null,
  },
  fixAlign: {
    default: true,
  },
});
const icon = defineAsyncComponent(
    () => import(`@/lib/icons/${props.name}.svg`),
);
</script>

<style lang="scss" scoped>
.icon {
  position: relative;
  display: inline-block;
  vertical-align: middle;
  width: 0.8rem;
  height: 0.8rem;
  margin: 0 0.3rem;
  top: 0.05rem;
  fill: currentColor;

  & svg {
    display: inline-block;
    vertical-align: top;
    width: 100%;
    height: 100%;
  }

  &:first-child {
    margin-left: 0;
  }

  &:last-child {
    margin-right: 0;
  }
}
</style>
