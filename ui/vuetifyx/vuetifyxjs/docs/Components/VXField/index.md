# vx-field 输入框

## Usage

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-field :model-value="dataNoBinding" label="input(autofocus)" autofocus/>
    </v-col>
    <v-col cols="6">
      <vx-field type="textarea" v-model="dataTextArea" label="textarea" />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6" sm="6">
      <vx-field v-model="dataWithTips" tips="this is tips" label="input with tooltip" />
    </v-col>
     <v-col cols="6" sm="6">
      <vx-field  type="textarea" v-model="dataWithTips" tips="this is tips" label="textarea with tooltip" />
    </v-col>
  </v-row>

    <v-row>
    <v-col cols="6">
      <vx-field
        placeholder="enter any value"
        readonly
        label="input(readonly)"
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        type="textarea"
        placeholder="enter any value"
        model-value="textareaaasdfasdf"
        readonly
        label="textarea(readonly)"
      />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <vx-field
        v-model="dataTextErrorMessages"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        label="input with error messages"
      />
    </v-col>

     <v-col cols="6">
      <vx-field
        v-model="dataTextErrorMessages"
        type="textarea"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        label="textarea with error messages"
      />
    </v-col>
  </v-row>

    <v-row>
    <v-col cols="6">
      <vx-field
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        disabled
        label="input disabled"
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        type="textarea"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        disabled
        label="textarea disabled"
      />
    </v-col>
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


## Type

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-field :model-value="dataNoBinding" label="Input (use :model-value)" />
      <div class="text-caption">Value Display:{{ dataNoBinding }}</div>
    </v-col>
    <v-col cols="6" sm="6">
      <vx-field v-model="vModelBinding" label="Input(use v-model)" />
      <div class="text-caption">Value Display:{{ vModelBinding }}</div>
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6" sm="6">
      <vx-field type="password" :model-value="dataNoBinding" label="Password" />
    </v-col>
  </v-row>

  <v-row>
    <v-col cols="6">
      <vx-field type="textarea" v-model="dataTextArea" label="Textarea" />
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const dataNoBinding = ref('hello world')
const vModelBinding = ref('hello world')
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


## Slot

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-label class="mb-2">search bar</vx-label>
      <vx-field :model-value="dataNoBinding" Placeholder="Search" clearable @click:clear="dataNoBinding = ''" width="320">
        <template #append-inner><v-icon icon="mdi-magnify" @click="onSearch" /></template>
      </vx-field>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const dataNoBinding = ref('hello world')

function onSearch() {
  alert("search content is" + dataNoBinding.value)

}
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::