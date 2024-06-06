import RichTextRedactor from "@/lib/RichTextRedactor.vue";

window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || []
window.__goplaidVueComponentRegisters.push((app, vueOptions) => {
    app.component("redactor", RichTextRedactor)
})
