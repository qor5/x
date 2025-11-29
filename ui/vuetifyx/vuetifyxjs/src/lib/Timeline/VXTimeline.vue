<template>
  <div
    ref="root"
    class="vx-timeline-wrap"
    :class="{
      'vx-timeline-sinuous': sinuous,
      'vx-timeline-animate': animateOnScroll,
      'vx-timeline-parallax': parallax
    }"
    v-bind="rootAttrs"
  >
    <v-timeline v-bind="combinedProps">
      <template v-if="!isDefaultSlotReallyEmpty" #default>
        <slot />
      </template>
    </v-timeline>
  </div>
</template>

<script setup lang="ts">
import {
  defineEmits,
  computed,
  useSlots,
  defineOptions,
  onMounted,
  onUnmounted,
  ref,
  nextTick
} from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs, rootAttrs } = useFilteredAttrs()
const slots = useSlots()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  sinuous: Boolean,
  animateOnScroll: Boolean,
  parallax: Boolean
})

const root = ref<HTMLElement | null>(null)
let observer: IntersectionObserver | null = null
let parallaxRafId: number | null = null

const handleParallax = () => {
  if (!root.value) return

  const opposites = root.value.querySelectorAll('.v-timeline-item__opposite')
  const bodies = root.value.querySelectorAll('.v-timeline-item__body')
  const windowHeight = window.innerHeight
  const center = windowHeight / 2

  opposites.forEach((el) => {
    const rect = el.getBoundingClientRect()
    const elCenter = rect.top + rect.height / 2
    const dist = elCenter - center
    const factor = 0.1
    const offset = dist * factor
    ;(el as HTMLElement).style.transform = `translateY(${offset}px)`
  })

  bodies.forEach((el) => {
    const rect = el.getBoundingClientRect()
    const elCenter = rect.top + rect.height / 2
    const dist = elCenter - center
    const factor = -0.02
    const offset = dist * factor
    ;(el as HTMLElement).style.transform = `translateY(${offset}px)`
  })

  parallaxRafId = requestAnimationFrame(handleParallax)
}

onMounted(() => {
  console.log('VXTimeline mounted, animateOnScroll:', props.animateOnScroll)
  if (props.animateOnScroll && root.value) {
    observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          // console.log(
          //   'Intersection entry:',
          //   entry.isIntersecting,
          //   'Rect:',
          //   entry.boundingClientRect,
          //   'Ratio:',
          //   entry.intersectionRatio
          // )
          if (entry.isIntersecting) {
            entry.target.classList.add('is-visible')
          } else {
            entry.target.classList.remove('is-visible')
          }
        })
      },
      {
        threshold: 0
      }
    )

    // Wait for layout to settle
    setTimeout(() => {
      if (root.value) {
        console.log('Root rect:', root.value.getBoundingClientRect())
        // Observe the body and opposite elements instead of the item itself,
        // because v-timeline-item might be display: contents
        const items = root.value.querySelectorAll(
          '.v-timeline-item__body, .v-timeline-item__opposite'
        )
        console.log('Timeline content items found:', items.length)
        items.forEach((item) => {
          observer?.observe(item)
        })
      }
    }, 1000)
  }

  if (props.parallax) {
    handleParallax()
  }
})

onUnmounted(() => {
  observer?.disconnect()
  if (parallaxRafId) {
    cancelAnimationFrame(parallaxRafId)
  }
})

const defaultOptions = computed(() => {
  return {
    // Default options if any
  }
})

const isDefaultSlotReallyEmpty = computed(() => {
  /* @ts-ignore */
  return !slots.default || !slots.default().length
})

// bugfix: bind event will auto bind to rootElement, and result in trigger twice
defineOptions({
  inheritAttrs: false
})

const combinedProps = computed(() => ({
  ...defaultOptions.value,
  ...filteredAttrs.value // passthrough the props that defined by vuetify
}))
</script>

<style lang="scss" scoped>
.vx-timeline-wrap {
  &.vx-timeline-animate {
    :deep(.v-timeline-item__body),
    :deep(.v-timeline-item__opposite) {
      opacity: 0;
      transform: translateY(20px);
      transition:
        opacity 0.6s ease-out,
        transform 0.6s ease-out;

      &.is-visible {
        opacity: 1;
        transform: translateY(0);
      }
    }
  }

  &.vx-timeline-parallax {
    :deep(.v-timeline-item__opposite) {
      transition: opacity 0.6s ease-out !important;
      // transform is controlled by JS
      align-self: center;
    }

    :deep(.v-timeline-item__body) {
      transition: opacity 0.6s ease-out !important;
      // transform is controlled by JS
    }
  }

  &.vx-timeline-sinuous {
    :deep(.v-timeline-divider__before),
    :deep(.v-timeline-divider__after) {
      background-color: transparent !important;
      width: 20px;
      background-image: url("data:image/svg+xml,%3Csvg width='12' height='100' viewBox='0 0 12 100' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M6 0 Q 12 25 6 50 T 6 100' fill='none' stroke='%23ccc' stroke-width='2'/%3E%3C/svg%3E");
      background-repeat: repeat-y;
      background-size: 12px 100px;
    }

    :deep(.v-timeline-divider__before) {
      background-position: bottom center;
    }

    :deep(.v-timeline-divider__after) {
      background-position: top center;
    }
  }
}
</style>
