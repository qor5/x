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
    default: 0
  },
  initialRotateY: {
    type: Number,
    default: 0
  },
  initialTranslateX: {
    type: Number,
    default: 0
  },
  initialTranslateY: {
    type: Number,
    default: 0
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
  padding: '50px',
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center'
}

const contentStyle = computed(() => {
  return {
    position: 'relative' as const,
    display: 'grid',
    justifyItems: 'center',
    alignItems: 'center',
    transformOrigin: 'center center',
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

<style lang="scss" scoped>
.vx-tilted-images {
  width: 100%;
}

.vx-tilted-images__content {
  :deep(*) {
    grid-area: 1 / 1;
  }

  @for $i from 1 through 10 {
    :deep(*:nth-child(#{$i})) {
      transform: translateZ(#{$i * 20}px);
    }
  }
}
</style>

<style lang="scss">
.timeline-with-images .v-timeline-item:nth-child(odd) .v-timeline-item__body {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;
}
.timeline-with-images .v-timeline-item:nth-child(odd) .v-timeline-item__opposite {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  text-align: right;
}
.timeline-with-images .v-timeline-item:nth-child(even) .v-timeline-item__body {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  text-align: right;
}
.timeline-with-images .v-timeline-item:nth-child(even) .v-timeline-item__opposite {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-align: left;
}
</style>
