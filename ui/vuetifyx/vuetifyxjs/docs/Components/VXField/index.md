# vx-field 输入框

## 基本用法

```html
<vx-field model-value="Hello World" label="field1" />
```

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-field :model-value="dataNoBinding" label="field1" />
    </v-col>
    <v-col cols="6" class="pt-12"><b> no binding:</b> {{ dataNoBinding }} </v-col>

    <v-col cols="6" sm="6">
      <vx-field v-model="dataWithBinding" label="field2" />
    </v-col>

    <v-col cols="6" class="pt-12"><b> with binding:</b> {{ dataWithBinding }} </v-col>

    <v-col cols="6" sm="6">
      <vx-field v-model="dataWithTips" tips="this is tips" label="field with tooltip" />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <vx-field
        v-model="dataTextErrorMessages"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        label="label with error messages"
      />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <vx-field type="textarea" v-model="dataTextArea" label="textarea" />
    </v-col>
    <v-col cols="6" class="pt-12"><b> with binding:</b> {{ dataTextArea }} </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const dataNoBinding = ref('hello world')
const dataWithBinding = ref('hello world')
const dataWithTips = ref('data with tips')
const dataTextArea = ref('textarea data')
const dataTextErrorMessages = ref('This is a description')
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::
