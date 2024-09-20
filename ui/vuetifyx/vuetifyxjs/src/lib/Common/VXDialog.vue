<template>
  <div class="vx-dialog-wrap">
    <v-dialog
      scrollable
      width="auto"
      :model-value="dialogVisible"
      v-bind="filteredAttrs"
      @update:model-value="onUpdateModelValue"
    >
      <template v-slot:default="{ isActive }">
        <v-card :title="title">
          <template v-slot:prepend v-if="prependIcon.icon">
            <v-icon :color="prependIcon.color" size="small" :icon="prependIcon.icon" />
          </template>

          <template v-slot:append v-if="showClose">
            <v-icon color="#757575" size="small" icon="mdi-close" @click="isActive.value = false" />
          </template>

          <v-card-text :style="[contentWidth]">
            <slot>{{ text }}</slot>
          </v-card-text>
          <v-card-actions class="custom-card-cation" v-if="showOK || showCancel">
            <v-btn
              v-if="showCancel"
              color="grey-darken-3"
              size="default"
              variant="tonal"
              @click="isActive.value = false"
              >{{ cancelText }}</v-btn
            >
            <v-btn v-if="showOK" color="primary" variant="flat" @click="isActive.value = false">{{
              okText
            }}</v-btn>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import { defineEmits, ref, watch, defineProps, computed, PropType } from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'

const { filteredAttrs } = useFilteredAttrs()
const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: Boolean,
  type: {
    type: String as PropType<'default' | 'warn' | 'error' | 'info'>,
    default: 'default'
  },
  text: String,
  showOK: {
    type: Boolean,
    default: true
  },
  showCancel: {
    type: Boolean,
    default: true
  },
  showClose: {
    type: Boolean,
    default: true
  },
  okText: {
    type: String,
    default: 'OK'
  },
  cancelText: {
    type: String,
    default: 'Cancel'
  },
  maxWidth: {
    type: [Number, String],
    default: 665
  },
  title: String
})

const dialogVisible = ref(props.modelValue)
const contentWidth = computed(() => `max-width:${props.maxWidth}px`)
const prependIcon = computed(() => {
  const vCardTitleIconMap = {
    default: {
      icon: '',
      color: ''
    },
    info: {
      color: 'primary',
      icon: 'mdi-alert-circle'
    },
    warn: {
      color: 'warning',
      icon: 'mdi-alert-circle'
    },
    success: {
      color: 'success',
      icon: 'mdi-check-circle'
    },
    error: {
      color: 'error',
      icon: 'mdi-alert-circle'
    }
  }

  return vCardTitleIconMap[props.type]
})

watch(
  () => props.modelValue,
  (newValue) => {
    dialogVisible.value = newValue
  }
)

function onUpdateModelValue(value: any) {
  emit('update:modelValue', value)
  dialogVisible.value = value
}
</script>

<style lang="scss" scoped>
.custom-card-cation {
  padding: 24px;
  padding-top: 0;

  .v-btn.v-btn--size-default {
    min-width: initial;
    padding-left: 16px;
    padding-right: 16px;
    font-weight: 500;
    letter-spacing: 1;
    &:deep(.v-btn__content) {
      letter-spacing: 0.25px;
    }
  }
}
</style>
