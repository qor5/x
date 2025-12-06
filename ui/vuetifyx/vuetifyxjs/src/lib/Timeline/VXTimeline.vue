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
const animationTimeouts = new Map<Element, number>()

const handleParallax = () => {
  if (!root.value) return

  const opposites = root.value.querySelectorAll('.v-timeline-item__opposite')
  const bodies = root.value.querySelectorAll('.v-timeline-item__body')
  const windowHeight = window.innerHeight
  const center = windowHeight / 2

  const applyParallax = (el: Element, factor: number) => {
    // If animateOnScroll is enabled, only apply parallax if the element is visible
    if (props.animateOnScroll && !el.classList.contains('is-visible')) {
      return
    }

    const rect = el.getBoundingClientRect()
    const elCenter = rect.top + rect.height / 2
    const dist = elCenter - center
    const offset = dist * factor
    ;(el as HTMLElement).style.transform = `translateY(${offset}px)`
  }

  opposites.forEach((el) => applyParallax(el, 0.1))
  bodies.forEach((el) => applyParallax(el, -0.02))

  parallaxRafId = requestAnimationFrame(handleParallax)
}

onMounted(() => {
  if (props.animateOnScroll && root.value) {
    observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            entry.target.classList.add('is-visible')

            const existingId = animationTimeouts.get(entry.target)
            if (existingId) {
              clearTimeout(existingId)
              animationTimeouts.delete(entry.target)
            }

            // Disable transform transition after animation completes to allow crisp parallax
            const id = window.setTimeout(() => {
              entry.target.classList.add('animation-done')
              animationTimeouts.delete(entry.target)
            }, 600)
            animationTimeouts.set(entry.target, id)
          } else {
            entry.target.classList.remove('is-visible')
            entry.target.classList.remove('animation-done')

            const existingId = animationTimeouts.get(entry.target)
            if (existingId) {
              clearTimeout(existingId)
              animationTimeouts.delete(entry.target)
            }
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
        // Observe the body and opposite elements instead of the item itself,
        // because v-timeline-item might be display: contents
        const items = root.value.querySelectorAll(
          '.v-timeline-item__body, .v-timeline-item__opposite'
        )
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
  animationTimeouts.forEach((id) => clearTimeout(id))
  animationTimeouts.clear()
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

      &.animation-done {
        transition: opacity 0.6s ease-out !important;
      }
    }
  }

  &.vx-timeline-parallax {
    :deep(.v-timeline-item__opposite) {
      align-self: center;
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
