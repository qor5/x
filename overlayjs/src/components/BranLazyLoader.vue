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
    const ef = this.loaderFunc;
    if (!ef) {
      return {
        lazyloader: {
          render: function() {
            return null;
          }
        }
      };
    }
    return {
      lazyloader() {
        if (!window.bran) {
          throw new Error("bran js should be loaded first.");
        }
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
