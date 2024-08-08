<script setup lang="ts">
import Cropper from "cropperjs";
import "cropperjs/dist/cropper.css";
import { type Ref, ref } from "vue";

const emit = defineEmits(["update:modelValue"]);
const cropper: Ref<Cropper | undefined> = ref();
const container = ref();

const img = ref();

const props = defineProps({
  // Library props
  containerStyle: Object,
  src: { type: String, default: "" },
  alt: String,
  imgStyle: Object,

  // CropperJS props
  viewMode: Number,
  dragMode: String,
  initialAspectRatio: Number,
  aspectRatio: Number,
  modelValue: Object,
  preview: [String, Array],
  responsive: {
    type: Boolean,
    default: true,
  },
  restore: {
    type: Boolean,
    default: true,
  },
  checkCrossOrigin: {
    type: Boolean,
    default: true,
  },
  checkOrientation: {
    type: Boolean,
    default: true,
  },
  crossorigin: String,
  modal: {
    type: Boolean,
    default: true,
  },
  guides: {
    type: Boolean,
    default: true,
  },
  center: {
    type: Boolean,
    default: true,
  },
  highlight: {
    type: Boolean,
    default: true,
  },
  background: {
    type: Boolean,
    default: true,
  },
  autoCrop: {
    type: Boolean,
    default: true,
  },
  autoCropArea: Number,
  movable: {
    type: Boolean,
    default: true,
  },
  rotatable: {
    type: Boolean,
    default: true,
  },
  scalable: {
    type: Boolean,
    default: true,
  },
  zoomable: {
    type: Boolean,
    default: false,
  },
  zoomOnTouch: {
    type: Boolean,
    default: false,
  },
  zoomOnWheel: {
    type: Boolean,
    default: false,
  },
  wheelZoomRatio: Number,
  cropBoxMovable: {
    type: Boolean,
    default: true,
  },
  cropBoxResizable: {
    type: Boolean,
    default: true,
  },
  toggleDragModeOnDblclick: {
    type: Boolean,
    default: true,
  },

  // Size limitation
  minCanvasWidth: Number,
  minCanvasHeight: Number,
  minCropBoxWidth: Number,
  minCropBoxHeight: Number,
  minContainerWidth: Number,
  minContainerHeight: Number,

  // callbacks
  // ready: Function,
  // cropstart: Function,
  // cropmove: Function,
  // cropend: Function,
  // crop: Function,
  // zoom: Function
});

const reset = () => {
  return cropper.value?.reset();
};

// Clear the crop box
const clear = () => {
  return cropper.value?.clear();
};

// Init crop box manually
const initCrop = () => {
  return cropper.value?.crop();
};

/**
 * Replace the image's src and rebuild the cropper
 * @param {string} url - The new URL.
 * @param {boolean} [onlyColorChanged] - Indicate if the new image only changed color.
 * @returns {Object} this
 */
const replace = (url: string, onlyColorChanged: boolean) => {
  return cropper.value?.replace(url, onlyColorChanged);
};

// Enable (unfreeze) the cropper
const enable = () => {
  return cropper.value?.enable();
};

// Disable (freeze) the cropper
const disable = () => {
  return cropper.value?.disable();
};

// Destroy the cropper and remove the instance from the image
const destroy = () => {
  return cropper.value?.destroy();
};

/**
 * Move the canvas with relative offsets
 * @param {number} offsetX - The relative offset distance on the x-axis.
 * @param {number} offsetY - The relative offset distance on the y-axis.
 * @returns {Object} this
 */
const move = (offsetX: number, offsetY: number) => {
  return cropper.value?.move(offsetX, offsetY);
};

/**
 * Move the canvas to an absolute point
 * @param {number} x - The x-axis coordinate.
 * @param {number} [y=x] - The y-axis coordinate.
 * @returns {Object} this
 */
const moveTo = (x: number, y = x) => {
  return cropper.value?.moveTo(x, y);
};

/**
 * Zoom the canvas with a relative ratio
 * @param {number} ratio - The target ratio.
 * @param {Event} _originalEvent - The original event if any.
 * @returns {Object} this
 */
const relativeZoom = (ratio: number) => {
  return cropper.value?.zoom(ratio);
};

/**
 * Zoom the canvas to an absolute ratio
 * @param {number} ratio - The target ratio.
 * @param pivot
 * @returns {Object} this
 */
const zoomTo = (ratio: number, pivot?: { x: number; y: number }) => {
  return cropper.value?.zoomTo(ratio, pivot);
};

/**
 * Rotate the canvas with a relative degree
 * @param {number} degree - The rotate degree.
 * @returns {Object} this
 */
const rotate = (degree: number) => {
  return cropper.value?.rotate(degree);
};

/**
 * Rotate the canvas to an absolute degree
 * @param {number} degree - The rotate degree.
 * @returns {Object} this
 */
const rotateTo = (degree: number) => {
  return cropper.value?.rotateTo(degree);
};

/**
 * Scale the image on the x-axis.
 * @param {number} scaleX - The scale ratio on the x-axis.
 * @returns {Object} this
 */
const scaleX = (scaleX: number) => {
  return cropper.value?.scaleX(scaleX);
};

/**
 * Scale the image on the y-axis.
 * @param {number} scaleY - The scale ratio on the y-axis.
 * @returns {Object} this
 */
const scaleY = (scaleY: number) => {
  return cropper.value?.scaleY(scaleY);
};

/**
 * Scale the image
 * @param {number} scaleX - The scale ratio on the x-axis.
 * @param {number} [scaleY=scaleX] - The scale ratio on the y-axis.
 * @returns {Object} this
 */
const scale = (scaleX: number, scaleY = scaleX) => {
  return cropper.value?.scale(scaleX, scaleY);
};

/**
 * Get the cropped area position and size data (base on the original image)
 * @param {boolean} [rounded=false] - Indicate if round the data values or not.
 * @returns {Object} The result cropped data.
 */
const getData = (rounded = false) => {
  return cropper.value?.getData(rounded);
};

/**
 * Set the cropped area position and size with new data
 * @param {Object} data - The new data.
 * @returns {Object} this
 */
const setData = (data: Cropper.SetDataOptions) => {
  return cropper.value?.setData(data);
};

/**
 * Get the container size data.
 * @returns {Object} The result container data.
 */
const getContainerData = () => {
  return cropper.value?.getContainerData();
};

/**
 * Get the image position and size data.
 * @returns {Object} The result image data.
 */
const getImageData = () => {
  return cropper.value?.getImageData();
};

/**
 * Get the canvas position and size data.
 * @returns {Object} The result canvas data.
 */
const getCanvasData = () => {
  return cropper.value?.getCanvasData();
};

/**
 * Set the canvas position and size with new data.
 * @param {Object} data - The new canvas data.
 * @returns {Object} this
 */
const setCanvasData = (data: Cropper.SetCanvasDataOptions) => {
  return cropper.value?.setCanvasData(data);
};

/**
 * Get the crop box position and size data.
 * @returns {Object} The result crop box data.
 */
const getCropBoxData = () => {
  return cropper.value?.getCropBoxData();
};

/**
 * Set the crop box position and size with new data.
 * @param {Object} data - The new crop box data.
 * @returns {Object} this
 */
const setCropBoxData = (data: Cropper.SetCropBoxDataOptions) => {
  return cropper.value?.setCropBoxData(data);
};

/**
 * Get a canvas drawn the cropped image.
 * @param {Object} [options={}] - The config options.
 * @returns {HTMLCanvasElement} - The result canvas.
 */
const getCroppedCanvas = (options = {}) => {
  return cropper.value?.getCroppedCanvas(options);
};

/**
 * Change the aspect ratio of the crop box.
 * @param {number} aspectRatio - The new aspect ratio.
 * @returns {Object} this
 */
const setAspectRatio = (aspectRatio: number) => {
  return cropper.value?.setAspectRatio(aspectRatio);
};

/**
 * Change the drag mode.
 * @param {string} mode - The new drag mode.
 * @returns {Object} this
 */
const setDragMode = (mode: Cropper.DragMode) => {
  return cropper.value?.setDragMode(mode);
};
const loaded = () => {
  container.value.style.height = `${img.value.height}px`;
  container.value.style.width = `${img.value.width}px`;

  const { containerStyle, modelValue, src, alt, imgStyle, ...data } = props;
  const propsObj = <any>{};

  for (const [key, value] of Object.entries(data)) {
    propsObj[key] = value;
  }

  propsObj.cropend = (evt: any) => {
    emit("update:modelValue", getData());
  };

  propsObj.zoom = (evt: any) => {
    emit("update:modelValue", getData());
  };

  propsObj.data = modelValue;

  cropper.value = new Cropper(img.value, propsObj);
};

const crossorigin: any = ref(props.crossorigin || undefined);

defineExpose({
  reset,
  clear,
  initCrop,
  replace,
  enable,
  disable,
  destroy,
  move,
  moveTo,
  relativeZoom,
  zoomTo,
  rotate,
  rotateTo,
  scaleX,
  scaleY,
  scale,
  getData,
  setData,
  getContainerData,
  getImageData,
  getCanvasData,
  setCanvasData,
  getCropBoxData,
  setCropBoxData,
  getCroppedCanvas,
  setAspectRatio,
  setDragMode,
});
</script>

<template>
  <div ref="container">
    <img
      ref="img"
      @load="loaded"
      :src="props.src"
      :alt="props.alt || 'image'"
      :crossorigin="crossorigin"
      :style="{ 'max-width': '100%', ...props.imgStyle }"
    />
  </div>
</template>

<style scoped></style>
