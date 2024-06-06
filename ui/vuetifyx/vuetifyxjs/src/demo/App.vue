<template>
  <v-app>
    <v-main>
      <v-container>
        <h1 class="text-h2 mb-4">select-many-example</h1>
        <select-many-example></select-many-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">datepicker-example</h1>
        <datepicker-example></datepicker-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">datetimepicker-example</h1>
        <datetimepicker-example></datetimepicker-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">linkage-select-example</h1>
        <linkage-select-example></linkage-select-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">filter-example</h1>
        <filter-example></filter-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">auto-complete-example</h1>
        <auto-complete-example></auto-complete-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">text-datepicker-example</h1>
        <text-datepicker-example></text-datepicker-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">default-navigator-example</h1>
        <default-navigator-example></default-navigator-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">send-variables-example</h1>
        <send-variables-example></send-variables-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">scroll-iframe-example</h1>
        <scroll-iframe-example></scroll-iframe-example>
        <v-divider class="mb-6"></v-divider>
        <h1 class="text-h2 mb-4">drag-listener-example</h1>
        <drag-listener-example></drag-listener-example>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import SelectManyExample from '@/demo/components/SelectManyExample.vue'
import DatepickerExample from '@/demo/components/DatepickerExample.vue'
import LinkageSelectExample from '@/demo/components/LinkageSelectExample.vue'
import DatetimepickerExample from '@/demo/components/DatetimepickerExample.vue'
import FilterExample from '@/demo/components/FilterExample.vue'
import AutoCompleteExample from '@/demo/components/AutoCompleteExample.vue'
import TextDatepickerExample from '@/demo/components/TextDatepickerExample.vue'
import DefaultNavigatorExample from '@/demo/components/DefaultNavigatorExample.vue'
import SendVariablesExample from '@/demo/components/SendVariablesExample.vue'
import ScrollIframeExample from '@/demo/components/ScrollIframeExample.vue'
import DragListenerExample from '@/demo/components/DragListenerExample.vue'
</script>
