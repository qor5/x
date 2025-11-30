<template>
  <div
    class="vx-tilted-images"
    @mousemove="handleMouseMove"
    @mouseleave="handleMouseLeave"
    ref="container"
    :style="containerStyle"
  >
    <div class="vx-tilted-images__content" :style="contentStyle">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

const props = defineProps({
  initialRotateX: {
    type: Number,
    default: -19.84
  },
  initialRotateY: {
    type: Number,
    default: -7.09
  },
  initialTranslateX: {
    type: Number,
    default: -8.86
  },
  initialTranslateY: {
    type: Number,
    default: 24.8
  }
})

const container = ref<HTMLElement | null>(null)
const rotateX = ref(props.initialRotateX)
const rotateY = ref(props.initialRotateY)
const translateX = ref(props.initialTranslateX)
const translateY = ref(props.initialTranslateY)

const containerStyle = {
  perspective: '1000px',
  overflow: 'visible',
  padding: '50px'
}

const contentStyle = computed(() => {
  return {
    position: 'relative' as const,
    width: '100%',
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    transformStyle: 'preserve-3d' as const,
    transform: `translateX(${translateX.value}px) translateY(${translateY.value}px) rotateX(${rotateX.value}deg) rotateY(${rotateY.value}deg)`,
    transition: 'transform 0.1s ease-out'
  }
})

const handleMouseMove = (e: MouseEvent) => {
  if (!container.value) return

  const rect = container.value.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top

  const centerX = rect.width / 2
  const centerY = rect.height / 2

  const percentX = (x - centerX) / centerX
  const percentY = (y - centerY) / centerY

  // Adjust rotation based on mouse position relative to center
  // Max rotation change +/- 10 degrees
  rotateY.value = props.initialRotateY + percentX * 10
  rotateX.value = props.initialRotateX - percentY * 10 // Invert Y axis for natural feel
}

const handleMouseLeave = () => {
  // Reset to initial values
  rotateX.value = props.initialRotateX
  rotateY.value = props.initialRotateY
  translateX.value = props.initialTranslateX
  translateY.value = props.initialTranslateY
}
</script>

<style scoped>
.vx-tilted-images {
  width: 100%;
}
</style>
