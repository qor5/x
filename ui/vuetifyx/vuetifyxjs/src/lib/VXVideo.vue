<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import Hls from 'hls.js'

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
  }
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
    hls = new Hls()
    hls.loadSource(props.src)
    hls.attachMedia(video)
  }
}

onMounted(() => {
  initVideo()
})

watch(() => props.src, () => {
  initVideo()
})

onBeforeUnmount(() => {
  if (hls) hls.destroy()
})
</script>

<template>
  <v-card class="rounded-xl" elevation="4">
    <video
      ref="videoRef"
      controls
      :style="{ width: width, height: height, borderRadius: 'inherit' }"
      :poster="poster"
    ></video>
  </v-card>
</template>
