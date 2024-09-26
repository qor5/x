<template>
  <div class="vx-dialog-wrap">
    <v-dialog
      scrollable
      width="auto"
      :model-value="dialogVisible"
      :persistent="persistent"
      v-bind="filteredAttrs"
      @update:model-value="onUpdateModelValue"
    >
      <template v-slot:activator="{ isActive, props: activatorProps }">
        <slot name="activator" :props="{ isActive, activatorProps }" />
      </template>

      <template v-slot:default="{ isActive }">
        <v-card>
          <template #title>
            <span>{{ title }}</span>
          </template>

          <template v-slot:prepend v-if="prependIcon.icon">
            <v-icon :color="prependIcon.color" size="small" :icon="prependIcon.icon" />
          </template>

          <template v-slot:append v-if="!hideClose">
            <v-icon color="#757575" size="small" icon="mdi-close" @click="onClose(isActive)" />
          </template>

          <v-card-text
            :class="{ 'mb-6': !hideFooter, 'pb-0': !hideFooter }"
            :style="[contentWidth, contentMaxWidth, contentHeightStyle]"
          >
            <slot :isActive="isActive"
              ><span class="dialog-content-text">{{ text }}</span></slot
            >
          </v-card-text>
          <v-card-actions :class="props.size" v-if="!hideFooter">
            <slot :isActive="isActive" name="action-btn">
              <v-btn
                v-if="!hideCancel"
                color="grey-darken-3"
                :size="props.size === 'default' ? 'small' : 'default'"
                variant="tonal"
                @click="onCancel(isActive)"
                >{{ cancelText }}</v-btn
              >
              <v-btn
                v-if="!hideOk"
                color="primary"
                :size="props.size === 'default' ? 'small' : 'default'"
                :loading="isOkBtnLoading"
                variant="flat"
                @click="onOk(isActive)"
                >{{ okText }}</v-btn
              >
            </slot>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import {
  defineEmits,
  ref,
  watch,
  defineProps,
  computed,
  PropType,
  Ref,
  getCurrentInstance
} from 'vue'
import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import { useHasEventListener } from '@/lib/composables/useEventListener'

const { filteredAttrs } = useFilteredAttrs()
const { hasEventListener } = useHasEventListener()
const emit = defineEmits(['update:modelValue', 'click:ok', 'click:cancel', 'click:close'])
const props = defineProps({
  modelValue: Boolean,
  title: String,
  // subTitle: String,
  type: {
    type: String as PropType<'default' | 'warn' | 'error' | 'info'>,
    default: 'default'
  },
  size: {
    type: String,
    default: 'default'
  },
  text: String,
  hideOk: {
    type: Boolean,
    default: false
  },
  hideCancel: {
    type: Boolean,
    default: false
  },
  hideHeader: {
    type: Boolean,
    default: false
  },
  hideFooter: {
    type: Boolean,
    default: false
  },
  hideClose: {
    type: Boolean,
    default: false
  },
  okText: {
    type: String,
    default: 'OK'
  },
  cancelText: {
    type: String,
    default: 'Cancel'
  },
  persistent: Boolean,
  contentHeight: {
    type: [Number, String],
    default: 'auto'
  },
  width: {
    type: [Number, String],
    default: ''
  },
  maxWidth: {
    type: [Number, String],
    default: 665
  }
})

watch(
  () => props.modelValue,
  (newValue) => {
    dialogVisible.value = newValue
  }
)

const isOkBtnLoading = ref(false)
const dialogVisible = ref(props.modelValue)
const contentMaxWidth = computed(() => {
  return `max-width:${Math.max(+props.width, +props.maxWidth)}px`
})
const contentWidth = computed(() => {
  let contentWidthStyle

  if (props.size === 'default' && props.width === '') {
    contentWidthStyle = 'width:461px'
  } else {
    contentWidthStyle = `width:${props.width}px`
  }

  return contentWidthStyle
})
const contentHeightStyle = computed(() => `height:${props.contentHeight}px`)
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

function onUpdateModelValue(value: any) {
  emit('update:modelValue', value)
  dialogVisible.value = value
}

const instance = getCurrentInstance()

function onOk(isActive: Ref<boolean>) {
  if (hasEventListener('click:ok')) {
    emit('click:ok', { isActive, isLoading: isOkBtnLoading })
  } else {
    isActive.value = false
  }
}

function onCancel(isActive: Ref<boolean>) {
  if (hasEventListener('click:cancel')) {
    emit('click:cancel', { isActive })
  } else {
    isActive.value = false
  }
}

function onClose(isActive: Ref<boolean>) {
  if (hasEventListener('click-close')) {
    emit('click:close', isActive)
  } else {
    isActive.value = false
  }
}
</script>

<style lang="scss" scoped>
.dialog-content-text {
  font-size: 14px;
  font-weight: 400;
  line-height: 20px;
  color: rgb(var(--v-theme-grey-darken-2));
}

.v-card-actions {
  padding: 0 24px 24px;
  &.default {
    .v-btn.v-btn--size-small {
      min-width: initial;
      padding: 0 12px;
      font-size: 12px;
      font-weight: 400;
      &:deep(.v-btn__content) {
        letter-spacing: 0.04px;
      }
    }
  }

  .v-btn.v-btn--size-default {
    min-width: initial;
    padding: 0 16px;
    font-weight: 500;
    &:deep(.v-btn__content) {
      letter-spacing: 0.25px;
    }
  }

  .v-btn ~ .v-btn:not(.v-btn-toggle .v-btn) {
    margin-inline-start: 10px;
  }
}
</style>
