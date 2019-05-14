<template>
<div class='bran-drawer'>
  <slot name="trigger" :parent='this'></slot>
  <a-drawer
        placement="right"
        :closable="false"
        :width="width"
        :title="title"
        @close="onClose"
        :visible="visible"
  >
    <slot name="drawer"></slot>
  </a-drawer>
</div>
</template>

<script>
import Drawer from 'ant-design-vue/es/drawer'
import 'ant-design-vue/lib/drawer/style/index.css'

export default {
  name: 'bran-drawer',
  components: {
    'a-drawer': Drawer
  },
  props: ['width', 'title'],
  data() {
    return {
      visible: false,
    }
  },
  methods: {
    showDrawer() {
      this.visible = true
    },
    onClose() {
      this.visible = false
    },
  },
}
</script>
