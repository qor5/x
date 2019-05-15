<template>
  <div class="bran-drawer">
    <slot name="trigger" :parent="this"></slot>
    <a-drawer
      placement="right"
      :closable="false"
      :width="width"
      :title="title"
      @close="close"
      :visible="isVisible"
      :getContainer="getContainer"
	  :destroyOnClose="true"
    >
      <slot name="drawer" :parent="this"></slot>
    </a-drawer>
  </div>
</template>

<script>
import Drawer from "ant-design-vue/es/drawer";
import "ant-design-vue/lib/drawer/style/index.css";

export default {
  name: "bran-drawer",
  components: {
    "a-drawer": Drawer
  },
  props: ["width", "title", "visible"],
  data() {
    return {
      isVisible: this.visible === "true"
    };
  },
  methods: {
    show() {
      this.isVisible = true;
    },
    close() {
      this.isVisible = false;
    },
    getContainer() {
      return document.getElementById("app");
    }
  }
};
</script>
