/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins

// Components
import App from "@/demo/App.vue";

// Composables
import { createApp } from "vue";

const app = createApp(App);

app.mount("#app");
