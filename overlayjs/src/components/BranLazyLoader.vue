<template>
  <div class="bran-lazy-loader" v-if="visible">
    <component :is="lazyloader"></component>
  </div>
</template>

<script>
export default {
  name: "bran-lazy-loader",
  props: ["loaderFunc", "visible"],
  data: function() {
    const efdata = this.loaderFunc;
    return {
      lazyloader() {
        if (!window.bran) {
          throw new Error("bran js should be loaded first.");
        }

        const ef = JSON.parse(efdata);

        return window.bran
          .fetchEvent(ef, {})
          .then(r => {
            return r.json();
          })
          .then(json => {
            return {
              template: json.schema
            };
          });
      }
    };
  }
};
</script>
