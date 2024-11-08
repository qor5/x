<template>
  <div class="vx-avatar-wrap">
    <div
      class="vx-avatar"
      :alt="name"
      :title="name"
      :style="{
        borderRadius: computedSize.border,
        fontSize: computedSize.fontSize
      }"
    >
      <span v-if="!imgLoad">{{ displayName }}</span>
    </div>
    <!-- /** only used to get load status, don't use this img to display */ -->
    <img v-show="false" :src="img" :title="name" @load="onImgLoad" />
  </div>
</template>

<script lang="ts" setup>
import { defineProps, PropType, computed, ref } from 'vue'

const sizeMap = {
  'x-small': 16,

  small: 32,
  default: 40,
  large: 64,
  'x-large': 96
}

const breakPointShapeMap = new Map([
  [160, { rounded: 16, font: 60 }],
  [128, { rounded: 16, font: 48 }],
  [96, { rounded: 8, font: 40 }],
  [80, { rounded: 8, font: 28 }],
  [64, { rounded: 8, font: 20 }],
  [48, { rounded: 8, font: 18 }],
  [40, { rounded: 4, font: 16 }], //default
  [32, { rounded: 4, font: 14 }],
  [24, { rounded: 4, font: 12 }],
  [16, { rounded: 4, font: 12 }]
])

const props = defineProps({
  size: {
    type: [String, Number] as PropType<string | number>,
    default: 40
  },
  name: {
    type: String,
    default: ''
  },
  img: {
    type: String,
    default: ''
  }
})

const propsSize2Num = (str: string | number): number | undefined => {
  if (typeof str === 'number' || !isNaN(Number(str))) return +str
  else if (typeof str === 'string' && str in sizeMap) return sizeMap[str as keyof typeof sizeMap]
}

const imgLoad = ref(false)

const displayName = computed(() => props.name.split('')[0].toLocaleUpperCase())
const size = computed(() => `${propsSize2Num(props.size)}px`)
const imgUrl = computed(() => `url(${props.img})`)

const computedSize = computed(() => {
  const sizeNum = propsSize2Num(props.size)
  let rounded = 4
  let font = 16

  if (!sizeNum)
    return {
      border: `${rounded}px`,
      fontSize: `${font}px`
    }

  for (const [key, value] of breakPointShapeMap) {
    if (sizeNum >= key) {
      rounded = value.rounded
      font = value.font
      break
    }
  }

  return {
    border: `${rounded}px`,
    fontSize: `${font}px`
  }
})

const onImgLoad = () => {
  imgLoad.value = true
}
</script>

<style lang="scss" scoped>
.vx-avatar {
  overflow: hidden;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  color: rgba(var(--v-theme-primary), 1);
  background-color: rgba(var(--v-theme-primary-lighten-2), 1);
  background-size: cover;
  width: v-bind(size);
  height: v-bind(size);
  background-image: v-bind(imgUrl);
}
</style>
