# v-checkbox 复选框

## 基本用法

:::demo

```vue
<template>
  <div>
    <v-row>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox v-model="ex4" color="red" label="red" value="red" hide-details></v-checkbox>
        <v-checkbox
          v-model="ex4"
          color="red-darken-3"
          label="red-darken-3"
          value="red-darken-3"
          hide-details
        ></v-checkbox>
      </v-col>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox
          v-model="ex4"
          color="indigo"
          label="indigo"
          value="indigo"
          hide-details
        ></v-checkbox>
        <v-checkbox
          v-model="ex4"
          color="indigo-darken-3"
          label="indigo-darken-3"
          value="indigo-darken-3"
          hide-details
        ></v-checkbox>
      </v-col>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox
          v-model="ex4"
          color="orange"
          label="orange"
          value="orange"
          hide-details
        ></v-checkbox>
        <v-checkbox
          v-model="ex4"
          color="orange-darken-3"
          label="orange-darken-3"
          value="orange-darken-3"
          hide-details
        ></v-checkbox>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox
          v-model="ex4"
          color="primary"
          label="primary"
          value="primary"
          hide-details
        ></v-checkbox>
        <v-checkbox
          v-model="ex4"
          color="secondary"
          label="secondary"
          value="secondary"
          hide-details
        ></v-checkbox>
      </v-col>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox
          v-model="ex4"
          color="success"
          label="success"
          value="success"
          hide-details
        ></v-checkbox>
        <v-checkbox v-model="ex4" color="info" label="info" value="info" hide-details></v-checkbox>
      </v-col>
      <v-col cols="12" md="4" sm="4">
        <v-checkbox
          v-model="ex4"
          color="warning"
          label="warning"
          value="warning"
          hide-details
        ></v-checkbox>
        <v-checkbox
          v-model="ex4"
          color="error"
          label="error"
          value="error"
          hide-details
        ></v-checkbox>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      ex4: [
        'red',
        'indigo',
        'orange',
        'primary',
        'secondary',
        'success',
        'info',
        'warning',
        'error',
        'red darken-3',
        'indigo darken-3',
        'orange darken-3'
      ]
    }
  }
}
</script>

```
:::