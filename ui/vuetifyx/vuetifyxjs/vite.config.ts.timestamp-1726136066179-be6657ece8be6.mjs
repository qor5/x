// vite.config.ts
import Components from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/unplugin-vue-components@0.26.0_@babel+parser@7.25.6_rollup@4.20.0_vue@3.5.3_typescript@5.5.4_/node_modules/unplugin-vue-components/dist/vite.js";
import { Vuetify3Resolver } from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/unplugin-vue-components@0.26.0_@babel+parser@7.25.6_rollup@4.20.0_vue@3.5.3_typescript@5.5.4_/node_modules/unplugin-vue-components/dist/resolvers.js";
import Vue from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/@vitejs+plugin-vue@5.1.2_vite@5.3.3_@types+node@20.14.15_sass@1.77.8__vue@3.5.3_typescript@5.5.4_/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import Vuetify, { transformAssetUrls } from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/vite-plugin-vuetify@2.0.4_vite@5.3.3_@types+node@20.14.15_sass@1.77.8__vue@3.5.3_typescript@5.5.4__vuetify@3.6.14/node_modules/vite-plugin-vuetify/dist/index.mjs";
import { resolve } from "path";
import vueJsx from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/@vitejs+plugin-vue-jsx@4.0.1_vite@5.3.3_@types+node@20.14.15_sass@1.77.8__vue@3.5.3_typescript@5.5.4_/node_modules/@vitejs/plugin-vue-jsx/dist/index.mjs";
import { defineConfig } from "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/node_modules/.pnpm/vite@5.3.3_@types+node@20.14.15_sass@1.77.8/node_modules/vite/dist/node/index.js";
import { fileURLToPath, URL } from "node:url";
var __vite_injected_original_dirname = "/Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs";
var __vite_injected_original_import_meta_url = "file:///Users/danielchan/thePlant/x/ui/vuetifyx/vuetifyxjs/vite.config.ts";
var vite_config_default = defineConfig({
  build: {
    cssCodeSplit: false,
    /**
     * ⚠️ warning: ⚠️
     * if configured lib options, vite will always inline assets in css includes, such as fonts,
     * it will create a huge size css about 7mb!
     */
    // lib: {
    //   entry: resolve(__dirname, 'src/lib/main.ts'),
    //   formats: ['umd'],
    //   name: 'vuetifyx'
    // },
    copyPublicDir: false,
    rollupOptions: {
      input: resolve(__vite_injected_original_dirname, "src/lib/main.ts"),
      external: ["vue"],
      output: {
        format: "umd",
        name: "vuetifyx",
        globals: {
          vue: "Vue",
          vuetify: "Vuetify"
        },
        chunkFileNames: `[name].js`,
        entryFileNames: "vuetifyx.min.js",
        assetFileNames: (assetInfo) => {
          if (assetInfo.name && assetInfo.name.endsWith(".css")) {
            return "assets/vuetifyx.min.css";
          }
          return "assets/[name].[ext]";
        }
      }
    }
  },
  plugins: [
    Vue({
      template: { transformAssetUrls }
    }),
    vueJsx(),
    // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
    Vuetify({
      autoImport: { labs: true },
      styles: {
        configFile: "src/lib/scss/_vuetify.scss"
      }
    }),
    Components({
      dts: true,
      dirs: ["src/demo/components", "src/lib"],
      resolvers: [Vuetify3Resolver()],
      include: [/\.vue$/]
    })
    // ViteFonts({
    //   google: {
    //     families: [{
    //       name: 'Roboto',
    //       styles: 'wght@100;300;400;500;700;900'
    //     }]
    //   }
    // })
  ],
  define: { "process.env": {} },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", __vite_injected_original_import_meta_url))
    },
    extensions: [".js", ".json", ".jsx", ".mjs", ".ts", ".tsx", ".vue"]
  },
  server: {
    port: 3e3
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMvZGFuaWVsY2hhbi90aGVQbGFudC94L3VpL3Z1ZXRpZnl4L3Z1ZXRpZnl4anNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIi9Vc2Vycy9kYW5pZWxjaGFuL3RoZVBsYW50L3gvdWkvdnVldGlmeXgvdnVldGlmeXhqcy92aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vVXNlcnMvZGFuaWVsY2hhbi90aGVQbGFudC94L3VpL3Z1ZXRpZnl4L3Z1ZXRpZnl4anMvdml0ZS5jb25maWcudHNcIjsvLyBQbHVnaW5zXG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tICd1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlJ1xuaW1wb3J0IHsgVnVldGlmeTNSZXNvbHZlciB9IGZyb20gJ3VucGx1Z2luLXZ1ZS1jb21wb25lbnRzL3Jlc29sdmVycydcbmltcG9ydCBWdWUgZnJvbSAnQHZpdGVqcy9wbHVnaW4tdnVlJ1xuaW1wb3J0IFZ1ZXRpZnksIHsgdHJhbnNmb3JtQXNzZXRVcmxzIH0gZnJvbSAndml0ZS1wbHVnaW4tdnVldGlmeSdcbi8vIGltcG9ydCBWaXRlRm9udHMgZnJvbSAndW5wbHVnaW4tZm9udHMvdml0ZSdcbmltcG9ydCB7IHJlc29sdmUgfSBmcm9tICdwYXRoJ1xuaW1wb3J0IHZ1ZUpzeCBmcm9tICdAdml0ZWpzL3BsdWdpbi12dWUtanN4J1xuXG4vLyBVdGlsaXRpZXNcbmltcG9ydCB7IGRlZmluZUNvbmZpZyB9IGZyb20gJ3ZpdGUnXG5pbXBvcnQgeyBmaWxlVVJMVG9QYXRoLCBVUkwgfSBmcm9tICdub2RlOnVybCdcblxuLy8gaHR0cHM6Ly92aXRlanMuZGV2L2NvbmZpZy9cbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZyh7XG4gIGJ1aWxkOiB7XG4gICAgY3NzQ29kZVNwbGl0OmZhbHNlLFxuICAgIC8qKlxuICAgICAqIFx1MjZBMFx1RkUwRiB3YXJuaW5nOiBcdTI2QTBcdUZFMEZcbiAgICAgKiBpZiBjb25maWd1cmVkIGxpYiBvcHRpb25zLCB2aXRlIHdpbGwgYWx3YXlzIGlubGluZSBhc3NldHMgaW4gY3NzIGluY2x1ZGVzLCBzdWNoIGFzIGZvbnRzLFxuICAgICAqIGl0IHdpbGwgY3JlYXRlIGEgaHVnZSBzaXplIGNzcyBhYm91dCA3bWIhXG4gICAgICovXG4gICAgLy8gbGliOiB7XG4gICAgLy8gICBlbnRyeTogcmVzb2x2ZShfX2Rpcm5hbWUsICdzcmMvbGliL21haW4udHMnKSxcbiAgICAvLyAgIGZvcm1hdHM6IFsndW1kJ10sXG4gICAgLy8gICBuYW1lOiAndnVldGlmeXgnXG4gICAgLy8gfSxcbiAgICBjb3B5UHVibGljRGlyOiBmYWxzZSxcbiAgICByb2xsdXBPcHRpb25zOiB7XG4gICAgICBpbnB1dDogcmVzb2x2ZShfX2Rpcm5hbWUsICdzcmMvbGliL21haW4udHMnKSxcbiAgICAgIGV4dGVybmFsOiBbJ3Z1ZSddLFxuICAgICAgb3V0cHV0OiB7XG4gICAgICAgIGZvcm1hdDogJ3VtZCcsXG4gICAgICAgIG5hbWU6ICd2dWV0aWZ5eCcsXG5cbiAgICAgICAgZ2xvYmFsczoge1xuICAgICAgICAgIHZ1ZTogJ1Z1ZScsXG4gICAgICAgICAgdnVldGlmeTogJ1Z1ZXRpZnknXG4gICAgICAgIH0sXG4gICAgICAgIGNodW5rRmlsZU5hbWVzOiBgW25hbWVdLmpzYCxcbiAgICAgICAgZW50cnlGaWxlTmFtZXM6ICd2dWV0aWZ5eC5taW4uanMnLFxuICAgICAgICBhc3NldEZpbGVOYW1lczogYXNzZXRJbmZvID0+IHtcbiAgICAgICAgICBpZiAoYXNzZXRJbmZvLm5hbWUgJiYgYXNzZXRJbmZvLm5hbWUuZW5kc1dpdGgoJy5jc3MnKSkge1xuICAgICAgICAgICAgcmV0dXJuICdhc3NldHMvdnVldGlmeXgubWluLmNzcyc7XG4gICAgICAgICAgfVxuICAgICAgICAgIHJldHVybiAnYXNzZXRzL1tuYW1lXS5bZXh0XSc7XG4gICAgICAgIH1cbiAgICAgIH1cbiAgICB9XG4gIH0sXG5cbiAgcGx1Z2luczogW1xuICAgIFZ1ZSh7XG4gICAgICB0ZW1wbGF0ZTogeyB0cmFuc2Zvcm1Bc3NldFVybHMgfVxuICAgIH0pLFxuICAgIHZ1ZUpzeCgpLFxuICAgIC8vIGh0dHBzOi8vZ2l0aHViLmNvbS92dWV0aWZ5anMvdnVldGlmeS1sb2FkZXIvdHJlZS9tYXN0ZXIvcGFja2FnZXMvdml0ZS1wbHVnaW4jcmVhZG1lXG4gICAgVnVldGlmeSh7XG4gICAgICBhdXRvSW1wb3J0OiB7IGxhYnM6IHRydWUgfSxcbiAgICAgIHN0eWxlczoge1xuICAgICAgICBjb25maWdGaWxlOiAnc3JjL2xpYi9zY3NzL192dWV0aWZ5LnNjc3MnXG4gICAgICB9XG4gICAgfSksXG4gICAgQ29tcG9uZW50cyh7XG4gICAgICBkdHM6IHRydWUsXG4gICAgICBkaXJzOiBbJ3NyYy9kZW1vL2NvbXBvbmVudHMnLCAnc3JjL2xpYiddLFxuICAgICAgcmVzb2x2ZXJzOiBbVnVldGlmeTNSZXNvbHZlcigpXSxcbiAgICAgIGluY2x1ZGU6IFsvXFwudnVlJC9dXG4gICAgfSksXG4gICAgLy8gVml0ZUZvbnRzKHtcbiAgICAvLyAgIGdvb2dsZToge1xuICAgIC8vICAgICBmYW1pbGllczogW3tcbiAgICAvLyAgICAgICBuYW1lOiAnUm9ib3RvJyxcbiAgICAvLyAgICAgICBzdHlsZXM6ICd3Z2h0QDEwMDszMDA7NDAwOzUwMDs3MDA7OTAwJ1xuICAgIC8vICAgICB9XVxuICAgIC8vICAgfVxuICAgIC8vIH0pXG4gIF0sXG4gIGRlZmluZTogeyAncHJvY2Vzcy5lbnYnOiB7fSB9LFxuICByZXNvbHZlOiB7XG4gICAgYWxpYXM6IHtcbiAgICAgICdAJzogZmlsZVVSTFRvUGF0aChuZXcgVVJMKCcuL3NyYycsIGltcG9ydC5tZXRhLnVybCkpXG4gICAgfSxcbiAgICBleHRlbnNpb25zOiBbJy5qcycsICcuanNvbicsICcuanN4JywgJy5tanMnLCAnLnRzJywgJy50c3gnLCAnLnZ1ZSddXG4gIH0sXG4gIHNlcnZlcjoge1xuICAgIHBvcnQ6IDMwMDBcbiAgfVxufSlcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFDQSxPQUFPLGdCQUFnQjtBQUN2QixTQUFTLHdCQUF3QjtBQUNqQyxPQUFPLFNBQVM7QUFDaEIsT0FBTyxXQUFXLDBCQUEwQjtBQUU1QyxTQUFTLGVBQWU7QUFDeEIsT0FBTyxZQUFZO0FBR25CLFNBQVMsb0JBQW9CO0FBQzdCLFNBQVMsZUFBZSxXQUFXO0FBWG5DLElBQU0sbUNBQW1DO0FBQXFLLElBQU0sMkNBQTJDO0FBYy9QLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQzFCLE9BQU87QUFBQSxJQUNMLGNBQWE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLElBV2IsZUFBZTtBQUFBLElBQ2YsZUFBZTtBQUFBLE1BQ2IsT0FBTyxRQUFRLGtDQUFXLGlCQUFpQjtBQUFBLE1BQzNDLFVBQVUsQ0FBQyxLQUFLO0FBQUEsTUFDaEIsUUFBUTtBQUFBLFFBQ04sUUFBUTtBQUFBLFFBQ1IsTUFBTTtBQUFBLFFBRU4sU0FBUztBQUFBLFVBQ1AsS0FBSztBQUFBLFVBQ0wsU0FBUztBQUFBLFFBQ1g7QUFBQSxRQUNBLGdCQUFnQjtBQUFBLFFBQ2hCLGdCQUFnQjtBQUFBLFFBQ2hCLGdCQUFnQixlQUFhO0FBQzNCLGNBQUksVUFBVSxRQUFRLFVBQVUsS0FBSyxTQUFTLE1BQU0sR0FBRztBQUNyRCxtQkFBTztBQUFBLFVBQ1Q7QUFDQSxpQkFBTztBQUFBLFFBQ1Q7QUFBQSxNQUNGO0FBQUEsSUFDRjtBQUFBLEVBQ0Y7QUFBQSxFQUVBLFNBQVM7QUFBQSxJQUNQLElBQUk7QUFBQSxNQUNGLFVBQVUsRUFBRSxtQkFBbUI7QUFBQSxJQUNqQyxDQUFDO0FBQUEsSUFDRCxPQUFPO0FBQUE7QUFBQSxJQUVQLFFBQVE7QUFBQSxNQUNOLFlBQVksRUFBRSxNQUFNLEtBQUs7QUFBQSxNQUN6QixRQUFRO0FBQUEsUUFDTixZQUFZO0FBQUEsTUFDZDtBQUFBLElBQ0YsQ0FBQztBQUFBLElBQ0QsV0FBVztBQUFBLE1BQ1QsS0FBSztBQUFBLE1BQ0wsTUFBTSxDQUFDLHVCQUF1QixTQUFTO0FBQUEsTUFDdkMsV0FBVyxDQUFDLGlCQUFpQixDQUFDO0FBQUEsTUFDOUIsU0FBUyxDQUFDLFFBQVE7QUFBQSxJQUNwQixDQUFDO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLEVBU0g7QUFBQSxFQUNBLFFBQVEsRUFBRSxlQUFlLENBQUMsRUFBRTtBQUFBLEVBQzVCLFNBQVM7QUFBQSxJQUNQLE9BQU87QUFBQSxNQUNMLEtBQUssY0FBYyxJQUFJLElBQUksU0FBUyx3Q0FBZSxDQUFDO0FBQUEsSUFDdEQ7QUFBQSxJQUNBLFlBQVksQ0FBQyxPQUFPLFNBQVMsUUFBUSxRQUFRLE9BQU8sUUFBUSxNQUFNO0FBQUEsRUFDcEU7QUFBQSxFQUNBLFFBQVE7QUFBQSxJQUNOLE1BQU07QUFBQSxFQUNSO0FBQ0YsQ0FBQzsiLAogICJuYW1lcyI6IFtdCn0K
