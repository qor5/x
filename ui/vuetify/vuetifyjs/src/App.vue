<template>
  <v-app>
    <v-main>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Button</h3>
        <Button></Button>
      </v-container>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Table</h3>
        <Table></Table>
      </v-container>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Autocomplete</h3>
        <Autocomplete></Autocomplete>
      </v-container>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Checkbox</h3>
        <Checkbox></Checkbox>
      </v-container>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Text Field</h3>
        <TextField></TextField>
      </v-container>
      <v-container fluid>
        <h3 class="v-heading mb-2 text-h3">Card</h3>
        <Card></Card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>
import Button from "@/components/Button.vue";
import Table from "@/components/Table.vue";
import Autocomplete from "@/components/Autocomplete.vue";
import Checkbox from "@/components/Checkbox.vue";
import TextField from "@/components/TextField.vue";
import Card from "@/components/Card.vue";
</script>
