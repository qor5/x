# vx-field 输入框

带样式预设的表单输入项

## API

### Props

| Name          | Introduction                                                                    | Type                                                                                                                        | Default Value |
| ------------- | ------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- | ------------- |
| label         | 输入框标题                                                                      | `String`                                                                                                                    | -             |
| tips          | 输入框标题 tooltip 提示                                                         | `String`                                                                                                                    | -             |
| name          | 输入框 label 的 for 属性，一般用来定位到输入框的元素，需要配合 id 使用          | `String`                                                                                                                    | -             |
| id            | 直接作用到输入框原生元素的 id 属， 可辅助label元素定位性                        | `String`                                                                                                                    | -             |
| placeholder   | 占位提示                                                                        | `String`                                                                                                                    | -             |
| type          | [输入框类型](./#输入框类型)                                                     | `String`                                                                                                                    | `text`        |
| errorMessages | 下方常显的错误信息                                                              | `String`                                                                                                                    | `text`        |
| disabled      | 是否禁用                                                                        | `Boolean`                                                                                                                   | `text`        |
| readonly      | 是否只读                                                                        | `Boolean`                                                                                                                   | `text`        |
| autofocus     | 光标自动聚焦到输入框                                                            | `Boolean`                                                                                                                   | `text`        |
| required      | 输入框标题样式上是否显示必填标星，仍然需要配合rules属性才能做到必填提示，见示例 | `Boolean`                                                                                                                   | `text`        |
| rules         | 输入框标题样式上是否显示必填标星，仍然需要配合rules属性才能做到必填提示         | [ValidationRule](https://github.com/vuetifyjs/vuetify/blob/master/packages/vuetify/src/composables/validation.ts#L16-L20)[] | `[]`          |

## 输入框类型

### Text

> 所有的 [v-text-field](https://vuetifyjs.com/en/api/v-text-field/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-field v-model="dataWithBinding" label="input(autofocus)" autofocus clearable />
      <div class="text-caption mt-2">v-model binding value: {{ dataWithBinding }}</div>
    </v-col>

    <v-col cols="6" sm="6">
      <vx-field
        model-value="data with tips"
        tips="this is tips"
        label="input with tooltip"
        clearable
      />
    </v-col>

    <v-col cols="6" sm="6">
      <vx-field
        model-value="data with tips"
        name="abc"
        id="abc"
        label="click label and focus on input"
        clearable
      />
      <div class="text-caption">use attr (name + id) to achieve this</div>
    </v-col>

    <v-col cols="6">
      <vx-field
        placeholder="enter any value"
        required
        :rules="[(value) => !!value || 'This input field is requried']"
        label="input (required)"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-field placeholder="enter any value" readonly label="input(readonly)" />
    </v-col>

    <v-col cols="6">
      <vx-field
        model-value="This is a description"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        label="input with error messages"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        disabled
        label="input disabled"
      />
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'
const dataWithBinding = ref('hello world')
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::

### Textarea

- 使用时需要指定 type 为 `textarea`

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6">
      <vx-field type="textarea" v-model="dataTextArea" label="textarea" clearable />
    </v-col>

    <v-col cols="6" sm="6">
      <vx-field
        type="textarea"
        v-model="dataWithTips"
        tips="this is tips"
        label="textarea with tooltip"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        type="textarea"
        placeholder="enter any value"
        required
        :rules="[(value) => !!value || 'This textarea field is requried']"
        label="input (required)"
        clearable
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

    <v-col cols="6">
      <vx-field
        v-model="dataTextErrorMessages"
        type="textarea"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        label="textarea with error messages"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        type="textarea"
        error-messages="This is an error message explanation"
        placeholder="enter any value"
        disabled
        label="textarea disabled"
        clearable
      />
    </v-col>

    <v-col cols="6">
      <vx-field
        type="textarea"
        error-messages="This is an error message explanation and rows = 10"
        :rows="10"
        placeholder="enter any value"
        label="textarea disabled"
        clearable
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

> 所有的 [v-textarea](https://vuetifyjs.com/en/api/v-text-field/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

### Password

#### Props

| Name                   | Introduction         | Type      | Default Value |
| ---------------------- | -------------------- | --------- | ------------- |
| passwordVisibleToggle  | 是否开启密码可见按钮 | `Boolean` | `false`       |
| passwordVisibleDefault | 密码是否默认可见钮   | `Boolean` | `false`       |

#### 示例

:::demo

```vue
<template>
  <v-row>
    <v-col cols="4" sm="4">
      <vx-field type="password" :model-value="123456" label="Password" />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        required
        placeholder="enter a password"
        :rules="[(value) => !!value || 'please enter a password']"
        label="Password"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        :model-value="123456"
        password-visible-toggle
        label="Password with visible toogle"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        :model-value="123456"
        password-visible-toggle
        tips="enter password 123456"
        label="Password with title tip"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        :model-value="123456"
        password-visible-toggle
        password-visible-default
        label="Password with visible default"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        :model-value="123456"
        password-visible-toggle
        clearable
        label="Password clearable"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        placeholder="Please enter password"
        disabled
        label="Password disabled"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        placeholder="Please enter password"
        error-messages="This is an error message explanation"
        label="Password with error"
      />
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field
        type="password"
        password-visible-toggle
        clearable
        placeholder="Please enter password"
        error-messages="This is an error message explanation"
        label="Password with error with eye btn"
      />
    </v-col>
  </v-row>
</template>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::

### Number

#### Props

| Name | Introduction | Type | Default Value |
| ---- | ------------ | ---- | ------------- |

#### 示例

:::demo

```vue
<template>
  <v-row>
    <v-col cols="4" sm="4">
      <vx-field type="number" v-model="a" label="Password(string 0)" />
      v-model:{{ a }}
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field type="number" v-model="b" label="Password(string '')" />
      v-model:{{ b }}
    </v-col>

    <v-col cols="4" sm="4">
      <vx-field type="number" v-model="c" label="Password(number 0)" />
      v-model:{{ c }}
    </v-col>
  </v-row>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
const a = ref('0')
const b = ref('')
const c = ref(0)
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::

> 所有的 [v-number-input](https://vuetifyjs.com/en/api/v-number-input/) 原生 props 都可使用 v-bind:[props]="[value]" 实现或覆盖

## Slot

:::demo

```vue
<template>
  <v-row>
    <v-col cols="6" sm="6">
      <vx-label class="mb-2">search bar</vx-label>
      <vx-field
        v-model="dataNoBinding"
        Placeholder="Search"
        clearable
        @click:clear="dataNoBinding = ''"
        width="320"
      >
        <template #append-inner><v-icon icon="mdi-magnify" @click="onSearch" /></template>
      </vx-field>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue'

const dataNoBinding = ref('hello world')

function onSearch() {
  alert('search content is ' + dataNoBinding.value)
}
</script>

<style scoped lang="css">
* {
  word-break: break-word;
}
</style>
```

:::
