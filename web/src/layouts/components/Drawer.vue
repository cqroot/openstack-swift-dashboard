<template>
  <v-navigation-drawer
    app
    width="260"
    disable-resize-watcher
    class="app-navigation-menu"
    v-model="drawer"
    :dark="night"
  >
    <v-list-item to="/" style="min-height: 63px">
      <v-list-item-content>
        <v-list-item-title class="text-h5 text-center">
          Swift Dashboard
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-divider></v-divider>

    <v-list dense nav>
      <v-list-item
        v-for="item in drawerRouter"
        :key="item.name"
        link
        :to="item.path"
      >
        <v-list-item-icon>
          <v-icon>{{ item.meta.icon }}</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>{{ item.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
import drawerRouter from "@/router/drawerRouter.js";

export default {
  data() {
    return {
      drawerRouter,
    };
  },
  computed: {
    drawer: {
      get() {
        return this.$store.state.drawer;
      },
      set(val) {
        this.$store.commit("setDrawer", val);
      },
    },
    night: {
      get() {
        return this.$store.state.night;
      },
      set(val) {
        this.$store.commit("setNight", val);
      },
    },
  },
};
</script>
