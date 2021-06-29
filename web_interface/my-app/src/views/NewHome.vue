<template>
  <div class="newhome">
    <Counter msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
// @ is an alias to /src
import Counter from "@/components/Counter.vue";

export default {
  name: "NewHome",
  components: {
    Counter,
  },
};
</script>
