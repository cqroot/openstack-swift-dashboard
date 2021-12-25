<template>
  <Layout>
    <router-view class="page_content"></router-view>
  </Layout>
</template>

<script>
import Layout from "@/layouts/Layout.vue";

export default {
  name: "App",
  components: {
    Layout,
  },
};
</script>

<style lang="scss" scoped>
#app {
  background-color: #f4f5f5;
  overflow: auto;
}
</style>
