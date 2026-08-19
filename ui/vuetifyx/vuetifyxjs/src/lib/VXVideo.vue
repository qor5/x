<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch, useAttrs, computed } from 'vue'
import Hls from 'hls.js'

defineOptions({
  inheritAttrs: false
})

const props = defineProps({
  src: {
    type: String,
    required: true
  },
  poster: {
    type: String,
    default: ''
  },
  width: {
    type: [String, Number],
    default: '100%'
  },
  height: {
    type: [String, Number],
    default: 'auto'
  },
  rounded: {
    type: [String, Boolean],
    default: false
  }
})

const attrs = useAttrs()
const containerClass = computed(() => {
  const classes: Record<string, boolean> = {
    rounded: props.rounded === true || props.rounded === '',
    [`rounded-${props.rounded}`]: typeof props.rounded === 'string' && props.rounded !== ''
  }

  const cls = attrs.class
  if (typeof cls === 'string' && cls.includes('rounded-xl')) {
    classes['rounded-xl'] = true
  }
  return classes
})

const videoRef = ref<HTMLVideoElement | null>(null)
let hls: Hls | null = null

const initVideo = () => {
  const video = videoRef.value
  if (!video) return
  if (!props.src) return

  if (hls) {
    hls.destroy()
    hls = null
  }

  // 1. Check if browser supports HLS natively (Safari)
  if (video.canPlayType('application/vnd.apple.mpegurl')) {
    video.src = props.src
  }
  // 2. If not, use Hls.js library (Chrome, Firefox, etc.)
  else if (Hls.isSupported()) {
    hls = new Hls({
      xhrSetup: function (xhr, url) {
        xhr.withCredentials = true // Crucial: Sends cookies with the request
      }
    })
    hls.loadSource(props.src)
    hls.attachMedia(video)
  }
}

onMounted(() => {
  initVideo()
})

watch(
  () => props.src,
  () => {
    initVideo()
  }
)

onBeforeUnmount(() => {
  if (hls) {
    hls.destroy()
    hls = null
  }
})

const showOverlay = ref(true)

const playVideo = () => {
  const video = videoRef.value
  if (!video) return
  video.play()
  showOverlay.value = false
}

const onPlay = () => {
  showOverlay.value = false
}

const onPause = () => {
  // Optional: check if we want to show overlay again on pause
  // showOverlay.value = true
}

// Watch poster prop to reset overlay if poster changes
watch(
  () => props.poster,
  () => {
    showOverlay.value = true
  }
)
</script>

<template>
  <div class="vx-video-container" :style="{ width: width, height: height }" :class="containerClass">
    <video
      ref="videoRef"
      v-bind="$attrs"
      class="vx-video-element"
      :poster="poster"
      @play="onPlay"
      @pause="onPause"
    ></video>

    <div
      v-if="showOverlay"
      class="vx-video-overlay"
      :style="{ backgroundImage: `url(${poster})` }"
      @click="playVideo"
    >
      <div class="vx-video-play-button">
        <svg viewBox="0 0 24 24" width="64" height="64" class="play-icon">
          <path fill="currentColor" d="M8 5v14l11-7z" />
        </svg>
      </div>
    </div>
  </div>
</template>

<style scoped>
.vx-video-container {
  position: relative;
  overflow: hidden;
  background-color: #000;
}

.vx-video-element {
  width: 100%;
  height: 100%;
  display: block;
}

.vx-video-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-size: cover;
  background-position: center;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1;
}

.vx-video-play-button {
  background-color: rgba(0, 0, 0, 0.6);
  border-radius: 50%;
  padding: 16px;
  transition:
    transform 0.2s,
    background-color 0.2s;
  color: white;
  display: flex;
}

.vx-video-overlay:hover .vx-video-play-button {
  transform: scale(1.1);
  background-color: rgba(0, 0, 0, 0.8);
}
</style>
