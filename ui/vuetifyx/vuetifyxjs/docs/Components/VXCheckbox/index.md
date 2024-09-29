# vx-checkbox

## Usage with v-model

:::demo

```vue
<template>
 <div class="d-flex flex-column ga-2">
  <vx-checkbox
    v-model="checked"
    label="Checkbox"
    :readonly="false"
  />
  <vx-checkbox
    v-model="checked"
    label="ReadonlyCheckbox"
    :readonly="true"
  />
  <p>Checkbox with model-value: {{checked}}</p>
 </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const checked = ref(false)
</script>
```

:::

## Usage with value (readonly)

:::demo

```vue
<template>
 <div class="d-flex flex-column ga-2">
  <vx-checkbox
    label="ReadonlyCheckboxWithValue"
    :value="true"
    :readonly="true"
  />
   <vx-checkbox
    label="ReadonlyCheckboxWithValue"
    :value="false"
    :readonly="true"
  />
 </div>
</template>

<script setup lang="ts">
</script>
```

:::
