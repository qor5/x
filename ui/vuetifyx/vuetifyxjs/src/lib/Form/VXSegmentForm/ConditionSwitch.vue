<template>
  <div class="vx-condition-switch-wrap">
    <template v-if="type === 'switch'">
      <div class="vx-condition-btn-group">
        <div class="active-background" :style="activeBackgroundStyle"></div>
        <div
          v-for="item in items"
          :key="item"
          @click="handleClick(item)"
          :class="{ active: modelValue === item }"
        >
          {{ item }}
        </div>
      </div>
    </template>
    <template v-else>
      <div class="vx-condition-select-wrap">
        <select v-model="props.modelValue" class="vx-condition-select" @change="handleChange">
          <option v-for="item in items" :key="item" :value="item">{{ item }}</option>
        </select>
        <div class="select-display">
          <div class="select-value">{{ props.modelValue }}</div>
          <div class="select-arrow">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24">
              <path d="M7 10l5 5 5-5z" />
            </svg>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: 'And'
  },
  type: {
    type: String,
    enum: ['switch', 'dropdown'],
    default: 'switch'
  }
})
const items = ['And', 'Or']

const emit = defineEmits(['update:modelValue', 'change'])

const handleClick = (item: string) => {
  emit('update:modelValue', item)
  emit('change', item)
}

// 计算高亮背景的位置
const activeBackgroundStyle = computed(() => {
  const index = items.findIndex((item) => item.toLowerCase() === props.modelValue.toLowerCase())
  return {
    transform: `translateX(calc(${index * 100}% + ${8 * index}px))`
  }
})

const handleChange = (event: Event) => {
  emit('update:modelValue', (event.target as HTMLSelectElement).value)
}
</script>

<style lang="scss" scoped>
.vx-condition-switch-wrap {
  .vx-condition-btn-group {
    background-color: rgb(238, 238, 238);
    width: 96px;
    height: 32px;
    display: flex;
    justify-content: space-between;
    border-radius: 4px;
    padding: 4px;
    position: relative;

    .active-background {
      position: absolute;
      width: 40px;
      height: 24px;
      background-color: #fff;
      border-radius: 4px;
      transition: transform 0.3s ease;
      z-index: 1;
    }

    div {
      color: rgba(117, 117, 117, 1);
      width: 40px;
      text-align: center;
      line-height: 24px;
      height: 24px;
      border-radius: 4px;
      font-size: 12px;
      cursor: pointer;
      position: relative;
      z-index: 2;
      transition: color 0.3s ease;

      &.active {
        color: rgba(33, 33, 33, 1);
      }
    }
  }

  .vx-condition-select-wrap {
    position: relative;
    width: 56px;

    .vx-condition-select {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 24px;
      opacity: 0; // 使select不可见
      z-index: 2; // 确保select在最上层可交互
      cursor: pointer;
    }

    .select-display {
      position: relative;
      width: 100%;
      height: 24px;
      background: rgba(238, 238, 238, 1);
      border-radius: 4px;
      padding: 0 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      pointer-events: none; // 防止干扰select的点击
    }

    .select-value {
      font-size: 12px;
      color: rgba(66, 66, 66, 1);
    }

    .select-arrow {
      margin-left: 4px; // 固定间隔4px
      display: flex;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
