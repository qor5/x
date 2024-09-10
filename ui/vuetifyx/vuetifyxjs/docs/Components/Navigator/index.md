# navigator 导航栏

## 基本用法

:::demo

```vue
<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      class="rounded-lg my-2 ma-1"
      width="320"
      max-width="320"
      permanent
      elevation="2"
      temporary
    >
      <v-toolbar extended color="white" dark>
        <v-toolbar-title>Q O R </v-toolbar-title>
        <v-row align="center" justify="end">
          <v-col cols="auto">
            <v-btn variant="plain" icon>
              <v-icon>mdi-information</v-icon>
              <v-icon>mdi-menu-down</v-icon>
            </v-btn>
          </v-col>
          <v-divider vertical class="my-6"></v-divider>
          <v-col cols="auto">
            <v-app-bar-nav-icon
              icon="mdi-menu"
              class="text-grey-darken-1"
              @click="drawer = !drawer"
            ></v-app-bar-nav-icon
          ></v-col>
        </v-row>
      </v-toolbar>
      <div class="align-center" max-width="320" elevation="0">
        <v-list lines="two">
          <v-list-item v-for="(item, i) in leftMenuItems" :key="i" :value="item" color="primary">
            <template v-slot:prepend>
              <v-icon :icon="item.icon"></v-icon>
            </template>
            <v-list-item-title v-text="item.title" class="text-grey-darken-1"></v-list-item-title>
          </v-list-item>
        </v-list>
      </div>

      <template v-slot:append>
        <v-divider></v-divider>
        <div class="pa-4">
          <v-row align="center">
            <v-col cols="auto">
              <v-avatar color="blue" size="48" class="rounded-lg">
                <span class="text-h5">{{ userProfile.initials }}</span>
              </v-avatar>
              <span class="mx-2 text-sm-body-2 text-grey-darken-1">{{ userProfile.nickname }}</span>
            </v-col>
            <v-col cols="auto" class="d-flex justify-center align-center">
              <v-btn
                density="compact"
                variant="text"
                icon="mdi-chevron-right"
                class="text-grey-darken-1"
              ></v-btn>
            </v-col>
            <v-divider vertical class="my-2"></v-divider>
            <v-col cols="3" class="d-flex align-center justify-center">
              <v-btn
                density="compact"
                variant="text"
                icon="mdi-bell-outline"
                class="text-grey-darken-1"
              ></v-btn>
            </v-col>
          </v-row>
        </div>
      </template>
    </v-navigation-drawer>
    <v-app-bar :elevation="0" class="my-2">
      <v-row>
        <v-col cols="auto">
          <v-breadcrumbs :items="['Foo', 'Bar', 'Fizz']" v-if="drawer"></v-breadcrumbs>
        </v-col>
      </v-row>
    </v-app-bar>

    <v-main class="d-flex justify-center align-center">
      <h1>Main container</h1>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const userProfile = ref({
  initials: 'VJ',
  nickname: 'Peterson Lee',
  age: 25
})
const drawer = ref(true)
const leftMenuItems = ref([
  {
    icon: 'mdi-information-outline',
    title: 'Dashboard'
  },
  {
    icon: 'mdi-information-outline',
    title: 'Identities'
  },
  {
    icon: 'mdi-information-outline',
    title: 'Authentication'
  },
  {
    icon: 'mdi-information-outline',
    title: 'Account Experience'
  }
])
</script>

<style scoped>
/** css hack, make it display normal in demo preview */
.v-application {
  overflow:hidden;
}
.v-navigation-drawer, .v-application__wrap > header{
  position:absolute!important;
}
.v-navigation-drawer {
  /* height: 865px!important; */
  height: 99%!important;

}
</style>

```
:::