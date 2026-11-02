<template>
  <div class="vx-dialog-wrap">
    <v-dialog
      scrollable
      width="auto"
      :model-value="dialogVisible"
      :no-click-animation="noClickAnimation"
      :persistent="persistent"
      v-bind="filteredAttrs"
      @update:model-value="onUpdateModelValue"
    >
      <template v-slot:activator="{ isActive, props: activatorProps }">
        <slot name="activator" :props="{ isActive, activatorProps }" />
      </template>

      <template v-slot:default="{ isActive }">
        <v-card ref="dialogMain">
          <template #title v-if="!resolvedHideHeader">
            <span>{{ title }}</span>
          </template>

          <template v-slot:prepend v-if="prependIcon.icon && !resolvedHideHeader">
            <v-icon :color="prependIcon.color" size="small" :icon="prependIcon.icon" />
          </template>

          <template v-slot:append v-if="!resolvedHideClose">
            <v-icon color="#757575" size="small" icon="mdi-close" @click="onClose(isActive)" />
          </template>
          <v-card-text
            :class="{ 'mb-6': !resolvedHideFooter, 'pb-0': !resolvedHideFooter }"
            :style="[
              contentWidth,
              contentMaxWidth,
              contentHeightStyle,
              { padding: props.contentPadding }
            ]"
          >
            <slot :isActive="isActive"
              ><span class="dialog-content-text">{{ text }}</span></slot
            >
          </v-card-text>
          <v-card-actions :class="props.size" v-if="!resolvedHideFooter">
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
                :disabled="disableOk"
                color="primary"
                :size="props.size === 'default' ? 'small' : 'default'"
                :loading="isOkBtnLoading || loadingOk"
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
import { defineEmits, ref, watch, defineProps, computed, PropType, Ref, effectScope } from 'vue'

import { useFilteredAttrs } from '@/lib/composables/useFilteredAttrs'
import { useHasEventListener } from '@/lib/composables/useEventListener'
import { onClickOutside } from '@vueuse/core'
const { filteredAttrs } = useFilteredAttrs()
const { hasEventListener } = useHasEventListener()
const scope = effectScope()
const emit = defineEmits([
  'update:modelValue',
  'click:ok',
  'click:cancel',
  'click:close',
  'click:outside',
  'update:hideFooter',
  'update:hideHeader',
  'update:hideClose'
])
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
  disableOk: {
    type: Boolean,
    default: false
  },
  loadingOk: {
    type: Boolean,
    default: false
  },
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
  noClickAnimation: Boolean,
  contentPadding: {
    type: String,
    default: ''
  },
  contentOnlyMode: {
    type: Boolean,
    default: false
  },
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

watch(
  () => props.contentOnlyMode,
  (newValue) => {
    if (newValue) {
      emit('update:hideFooter', true)
      emit('update:hideHeader', true)
      emit('update:hideClose', true)
    }
  },
  {
    immediate: true
  }
)

const dialogMain = ref(null)
const isOkBtnLoading = ref(false)
const dialogVisible = ref(props.modelValue)
const contentMaxWidth = computed(() => {
  return `max-width:${Math.max(+props.width, +props.maxWidth)}px`
})
const contentWidth = computed(() => {
  let contentWidthStyle

  if (props.size === 'default' && props.width === '') {
    contentWidthStyle = 'width:461px'
  } else if (+props.width > 0) {
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

const resolvedHideFooter = computed(() => props.hideFooter || props.contentOnlyMode)
const resolvedHideHeader = computed(() => props.hideHeader || props.contentOnlyMode)
const resolvedHideClose = computed(() => props.hideClose || props.contentOnlyMode)

scope.run(() => {
  onClickOutside(dialogMain, (event) => {
    // console.log('outside',event)
    emit('click:outside', event)
  })
})

function onUpdateModelValue(value: any) {
  emit('update:modelValue', value)
  dialogVisible.value = value
}

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
