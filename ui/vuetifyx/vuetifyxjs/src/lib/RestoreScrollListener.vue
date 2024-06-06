<script setup lang="ts">
import { onMounted } from 'vue'

let scrollLeft = 0
let scrollTop = 0
const pause = (duration: number) => {
  return new Promise((res) => setTimeout(res, duration))
}

const backoff = (retries: number, fn: Function, delay = 100) => {
  fn().catch((err: any) =>
    retries > 1 ? pause(delay).then(() => backoff(retries - 1, fn, delay * 2)) : Promise.reject(err)
  )
}

const restoreScroll = () => {
  window.scroll({ left: scrollLeft, top: scrollTop, behavior: 'smooth' })
  if (window.scrollX == scrollLeft && window.scrollY == scrollTop) {
    return Promise.resolve()
  }
  return Promise.reject()
}

onMounted(() => {
  window.addEventListener('fetchStart', (event) => {
    scrollLeft = window.scrollX
    scrollTop = window.scrollY
  })

  window.addEventListener('fetchEnd', (event) => {
    backoff(5, restoreScroll, 100)
  })
})
</script>

<template></template>

<style scoped></style>
