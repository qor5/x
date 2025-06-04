<template>
  <div class="vx-time-select-wrap">
    <vx-field
      v-show="showArea.hour"
      ref="inputFieldHour"
      v-model="hourValue"
      class="time-select"
      type="number"
      width="48"
      control-variant="hidden"
      maxlength="2"
      :disabled="disableHour"
      :min="0"
      :max="23"
      hide-details
      @update:modelValue="onChooseValue('hour', $event)"
      @click="!disableHour && (showHourMenu = true)"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(hourValue) }}</div>
      </template>

      <v-menu v-model="showHourMenu" height="300" target="parent">
        <v-list ref="hourListRef">
          <v-list-item
            :active="hourValue === item - 1"
            v-for="(item, index) in 24"
            :key="index"
            :value="index"
            color="primary"
            @click="onChooseValue('hour', item - 1)"
          >
            <v-list-item-title>{{ padZero(item - 1) }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </vx-field>

    <span v-if="showArea.minute" class="separate mx-2">:</span>

    <vx-field
      v-show="showArea.minute"
      ref="inputFieldMinute"
      v-model="minuteValue"
      class="time-select minute-field"
      type="number"
      width="48"
      :disabled="disableMinute"
      maxlength="2"
      :min="0"
      control-variant="hidden"
      :max="59"
      hide-details
      @update:modelValue="onChooseValue('minute', $event)"
      @click="!disableMinute && (showMinuteMenu = true)"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(minuteValue) }}</div>
      </template>

      <v-menu height="300" v-model="showMinuteMenu" target="parent">
        <v-list ref="minuteListRef">
          <v-list-item
            v-for="(item, index) in 60"
            :key="index"
            :value="index"
            :active="minuteValue === item - 1"
            color="primary"
            @click="onChooseValue('minute', item - 1)"
          >
            <v-list-item-title>{{ padZero(item - 1) }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </vx-field>

    <span v-if="showArea.second" class="separate mx-2">:</span>

    <vx-field
      v-show="showArea.second"
      ref="inputFieldSecond"
      v-model="secondValue"
      class="time-select second-field"
      type="number"
      width="48"
      :disabled="disableSecond"
      maxlength="2"
      :min="0"
      control-variant="hidden"
      :max="59"
      hide-details
      @update:modelValue="onChooseValue('second', $event)"
      @click="!disableSecond && (showSecondMenu = true)"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(secondValue) }}</div>
      </template>

      <v-menu height="300" v-model="showSecondMenu" target="parent">
        <v-list ref="secondListRef">
          <v-list-item
            v-for="(item, index) in 60"
            :key="index"
            :value="index"
            :active="secondValue === item - 1"
            color="primary"
            @click="onChooseValue('second', item - 1)"
          >
            <v-list-item-title>{{ padZero(item - 1) }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </vx-field>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, watch, defineEmits, computed, nextTick } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: '00:00:00'
  },
  formatStr: String,
  disableSecond: Boolean,
  disableMinute: Boolean,
  disableHour: Boolean
})

const emit = defineEmits(['update:modelValue'])

const showHourMenu = ref(false)
const showMinuteMenu = ref(false)
const showSecondMenu = ref(false)
const inputFieldHour = ref()
const inputFieldMinute = ref()
const inputFieldSecond = ref()
const hourListRef = ref()
const minuteListRef = ref()
const secondListRef = ref()
const hourValue = ref(0)
const minuteValue = ref(0)
const secondValue = ref(0)

watch(
  () => props.modelValue,
  (newVal) => {
    if (!newVal) return
    const [hour, minute, second] = newVal.split(':')
    // console.log('timeSelectValue', hour, minute, second)

    hourValue.value = +hour
    minuteValue.value = +minute
    secondValue.value = +second
  },
  { immediate: true }
)

const payloadValue = computed(
  () => `${padZero(hourValue.value)}:${padZero(minuteValue.value)}:${padZero(secondValue.value)}`
)

const showArea = computed(() => {
  // Analyze if current formatStr contains hour, minute, second
  const formatStr = props.formatStr

  if (!formatStr) {
    return {
      hour: false,
      minute: false,
      second: false
    }
  }

  return {
    hour: formatStr.includes('H') || formatStr.includes('h'),
    minute: formatStr.includes('m'),
    second: formatStr.includes('s')
  }
})

function padZero(num: number): string {
  return num < 10 ? `0${num}` : `${num}`
}

// Fast positioning to selected item function
function scrollToActiveItem(listRef: any, activeValue: number) {
  setTimeout(() => {
    // Try different ways to get the DOM element
    let listElement = null

    if (listRef?.value?.$el) {
      listElement = listRef.value.$el
    } else if (listRef?.value) {
      listElement = listRef.value
    }

    if (!listElement) {
      console.warn('Could not find list element')
      return
    }

    const allItems = listElement.querySelectorAll('.v-list-item')

    // Get the correct item by index (since activeValue corresponds to the index)
    const targetItem = allItems[activeValue]

    if (targetItem) {
      // Direct positioning to center, no animation
      targetItem.scrollIntoView({
        behavior: 'auto',
        block: 'center',
        inline: 'nearest'
      })
    } else {
      console.warn('Could not find target item at index:', activeValue)
    }
  }, 200)
}

// Watch menu display status and auto scroll to selected item
watch(showHourMenu, (newVal) => {
  if (newVal) {
    scrollToActiveItem(hourListRef, hourValue.value)
  } else {
    inputFieldHour.value?.blur()
  }
})

watch(showMinuteMenu, (newVal) => {
  if (newVal) {
    scrollToActiveItem(minuteListRef, minuteValue.value)
  } else {
    inputFieldMinute.value?.blur()
  }
})

watch(showSecondMenu, (newVal) => {
  if (newVal) {
    scrollToActiveItem(secondListRef, secondValue.value)
  } else {
    inputFieldSecond.value?.blur()
  }
})

function onChooseValue(type: 'hour' | 'minute' | 'second', value: number) {
  // console.log(type, value)
  if (type === 'hour') {
    inputFieldHour.value?.blur()
    hourValue.value = value
  }

  if (type === 'minute') {
    inputFieldMinute.value?.blur()
    minuteValue.value = value
  }

  if (type === 'second') {
    inputFieldSecond.value?.blur()
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
    &:deep(*) {
      cursor: pointer;
    }
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
