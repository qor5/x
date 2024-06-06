// @snippet_begin(TipTapRegisterVueComponent)
import { type App } from "vue";
import TipTapEditor from "@/lib/Editor.vue";

declare const window: any;
window.__goplaidVueComponentRegisters =
  window.__goplaidVueComponentRegisters || [];
window.__goplaidVueComponentRegisters.push((app: App, vueOptions: any): any => {
  app.component("tiptap-editor", TipTapEditor);
});
// @snippet_end
