import { type App } from "vue";
import Cropper from "@/lib/Cropper.vue";

declare const window: any;
window.__goplaidVueComponentRegisters =
  window.__goplaidVueComponentRegisters || [];
window.__goplaidVueComponentRegisters.push((app: App, vueOptions: any): any => {
  app.component("vue-cropper", Cropper);
});
