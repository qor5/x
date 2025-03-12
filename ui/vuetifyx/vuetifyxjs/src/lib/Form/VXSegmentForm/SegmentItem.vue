<template>
  <div class="vx-segment-item-wrap">
    <div class="condition-group">
      <vx-select
        style="min-width: 100px"
        placeholder="Select"
        hide-details
        :items="['Number', 'String']"
      ></vx-select>

      <span class="condition-text">of</span>

      <vx-select
        style="min-width: 100px"
        placeholder="Select"
        :items="['View Products']"
        hide-details
      ></vx-select>

      <span class="condition-text">from</span>

      <vx-select
        style="min-width: 100px"
        placeholder="Select"
        :items="['0']"
        hide-details
      ></vx-select>

      <span class="condition-text">to</span>
      <vx-select
        style="min-width: 100px"
        placeholder="Select"
        :items="['20']"
        hide-details
      ></vx-select>

      <span class="condition-text">in</span>

      <vx-select
        style="min-width: 100px"
        placeholder="Select"
        :items="['Last 7 Days']"
        hide-details
      ></vx-select>
    </div>
    <v-icon class="delete-icon" color="rgb(158, 158, 158)" size="24" @click="handleRemove"
      >mdi-minus-circle-outline</v-icon
    >
  </div>
</template>

<script setup lang="ts">
import { defineEmits } from 'vue'

const emit = defineEmits(['on-remove'])

const handleRemove = () => {
  emit('on-remove')
}
</script>

<style scoped lang="scss">
.vx-segment-item-wrap {
  position: relative;
  background: rgb(250, 250, 250);
  border-radius: 4px;
  border: 1px solid rgb(224, 224, 224);
  padding: 8px;
  margin-right: 25px;
}
.condition-group {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}
.delete-icon {
  position: absolute;
  top: 50%;
  right: -30px;
  transform: translateY(-50%);
  cursor: pointer;
}

.condition-text {
  color: rgb(158, 158, 158);
  line-height: 40px;
}
</style>
