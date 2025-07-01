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
      @click="onTextFieldClick('hour')"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(hourValue) }}</div>
      </template>

      <v-menu v-model="showHourMenu" height="300" target="parent">
        <!-- @click.stop avoid close parent menu when click on list item -->
        <v-list ref="hourListRef" @click.stop>
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
      @click="onTextFieldClick('minute')"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(minuteValue) }}</div>
      </template>

      <v-menu height="300" v-model="showMinuteMenu" target="parent">
        <!-- @click.stop avoid close parent menu when click on list item -->
        <v-list ref="minuteListRef" @click.stop>
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
      @click="onTextFieldClick('second')"
    >
      <template #prepend-inner>
        <div class="displayValue">{{ padZero(secondValue) }}</div>
      </template>

      <v-menu height="300" v-model="showSecondMenu" target="parent">
        <!-- @click.stop avoid close parent menu when click on list item -->
        <v-list ref="secondListRef" @click.stop>
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
import { onClickOutside } from '@vueuse/core'
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

function onTextFieldClick(type: 'hour' | 'minute' | 'second') {
  const map = {
    hour: {
      disable: props.disableHour,
      menu: showHourMenu
    },
    minute: {
      disable: props.disableMinute,
      menu: showMinuteMenu
    },
    second: {
      disable: props.disableSecond,
      menu: showSecondMenu
    }
  }
  if (map[type].disable) return
  map[type].menu.value = true
}
// fix: https://theplanttokyo.atlassian.net/browse/QOR5-1395
// fix issue when click on text field, the nested sub-menu will be closed immediately
;[
  {
    refValue: hourListRef,
    showMenu: showHourMenu,
    inputField: inputFieldHour
  },
  {
    refValue: minuteListRef,
    showMenu: showMinuteMenu,
    inputField: inputFieldMinute
  },
  {
    refValue: secondListRef,
    showMenu: showSecondMenu,
    inputField: inputFieldSecond
  }
].forEach(({ refValue, showMenu, inputField }) => {
  onClickOutside(
    refValue,
    (ev: Event) => {
      ev.stopPropagation()
      inputField.value?.blur()
      showMenu.value = false
      if (!inputField.value?.$el.contains(ev.target as Node)) {
        inputFieldHour.value?.$el.contains(ev.target as Node) && onTextFieldClick('hour')
        inputFieldMinute.value?.$el.contains(ev.target as Node) && onTextFieldClick('minute')
        inputFieldSecond.value?.$el.contains(ev.target as Node) && onTextFieldClick('second')
      }
    },
    { capture: true }
  )
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
  }
})

watch(showMinuteMenu, (newVal) => {
  if (newVal) {
    scrollToActiveItem(minuteListRef, minuteValue.value)
  }
})

watch(showSecondMenu, (newVal) => {
  if (newVal) {
    scrollToActiveItem(secondListRef, secondValue.value)
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

  nextTick(() => {
    if (type === 'hour') showHourMenu.value = false
    if (type === 'minute') showMinuteMenu.value = false
    if (type === 'second') showSecondMenu.value = false
  })
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
