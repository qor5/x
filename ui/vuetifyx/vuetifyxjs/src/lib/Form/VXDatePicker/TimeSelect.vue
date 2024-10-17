<template>
  <div class="vx-time-select-wrap">
    <div>
      <v-menu height="300">
        <template #activator="{ props }">
          <vx-field
            v-model="hourValue"
            class="time-select"
            type="number"
            width="70"
            maxlength="2"
            :min="0"
            :max="23"
            hide-details
            v-bind="props"
            @update:modelValue="onChooseValue('hour', $event)"
          >
            <template #prepend-inner>
              <div class="displayValue">{{ padZero(hourValue) }}</div>
            </template>
          </vx-field>
        </template>

        <v-list>
          <v-list-item
            :active="hourValue === item - 1"
            v-for="(item, index) in 24"
            :key="index"
            :value="index"
            color="primary"
          >
            <v-list-item-title @click="onChooseValue('hour', item - 1)">{{
              padZero(item - 1)
            }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </div>

    <span class="separate mx-2">:</span>
    <v-menu height="300">
      <template #activator="{ props }">
        <vx-field
          v-model="minuteValue"
          class="time-select"
          type="number"
          width="70"
          maxlength="2"
          :min="0"
          :max="59"
          hide-details
          v-bind="props"
          @update:modelValue="onChooseValue('minute', $event)"
        >
          <template #prepend-inner>
            <div class="displayValue">{{ padZero(minuteValue) }}</div>
          </template>
        </vx-field>
      </template>

      <v-list>
        <v-list-item
          v-for="(item, index) in 60"
          :key="index"
          :value="index"
          :active="minuteValue === item - 1"
          color="primary"
        >
          <v-list-item-title @click="onChooseValue('minute', item - 1)">{{
            padZero(item - 1)
          }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <span class="separate mx-2">:</span>

    <v-menu height="300">
      <template #activator="{ props }">
        <vx-field
          v-model="secondValue"
          class="time-select"
          type="number"
          width="70"
          maxlength="2"
          :min="0"
          :max="59"
          hide-details
          v-bind="props"
          @update:modelValue="onChooseValue('second', $event)"
        >
          <template #prepend-inner>
            <div class="displayValue">{{ padZero(secondValue) }}</div>
          </template>
        </vx-field>
      </template>

      <v-list>
        <v-list-item
          v-for="(item, index) in 60"
          :key="index"
          :value="index"
          :active="secondValue === item - 1"
          color="primary"
        >
          <v-list-item-title @click="onChooseValue('second', item - 1)">{{
            padZero(item - 1)
          }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, watch, defineEmits, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: '00:00:00'
  }
})

const emit = defineEmits(['update:modelValue'])

const hourValue = ref(0)
const minuteValue = ref(0)
const secondValue = ref(0)
// console.log('inner value', props.modelValue)

watch(
  () => props.modelValue,
  (newVal) => {
    if (!newVal) return
    const [hour, minute, second] = newVal.split(':')
    console.log('timeSelectValue', hour, minute, second)

    hourValue.value = +hour
    minuteValue.value = +minute
    secondValue.value = +second
  },
  { immediate: true }
)

function padZero(num: number): string {
  return num < 10 ? `0${num}` : `${num}`
}

const payloadValue = computed(
  () => `${padZero(hourValue.value)}:${padZero(minuteValue.value)}:${padZero(secondValue.value)}`
)

function onChooseValue(type: 'hour' | 'minute' | 'second', value: number) {
  console.log(type, value)
  if (type === 'hour') {
    hourValue.value = value
  }

  if (type === 'minute') {
    minuteValue.value = value
  }

  if (type === 'second') {
    secondValue.value = value
  }

  emit('update:modelValue', payloadValue.value)
}
</script>

<style lang="scss" scoped>
.vx-time-select-wrap {
  display: flex;
  align-items: center;
  .separate {
    font-weight: 700;
    line-height: 1;
  }

  .time-select {
    &:deep(.v-field__prepend-inner) {
      cursor: pointer;
      position: absolute;
      width: 45px;
      height: 40px;
      display: flex;
      align-items: center;
      justify-content: center;
    }
    &:deep(.v-btn--disabled) {
      opacity: 0.3;
    }
    &:deep(.v-btn--disabled) .v-btn__overlay {
      background-color: transparent;
    }
    &:deep(.v-field__field) {
      height: 40px;
    }
    &:deep(input) {
      display: none;
    }
  }
}
</style>
