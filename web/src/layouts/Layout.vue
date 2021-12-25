<template>
  <v-app>
    <Drawer> </Drawer>

    <AppBar> </AppBar>

    <v-main>
      <div style="margin: 20px">
        <slot> </slot>
      </div>
    </v-main>

    <v-snackbar :timeout="globalSnackbarTimeout" v-model="globalSnackbar">
      {{ globalSnackbarMessage }}
    </v-snackbar>
  </v-app>
</template>

<script>
import { mapState } from "vuex";
import Drawer from "./components/Drawer.vue";
import AppBar from "./components/AppBar.vue";

export default {
  components: {
    Drawer,
    AppBar,
  },
  computed: {
    ...mapState(["globalSnackbarTimeout", "globalSnackbarMessage"]),
    globalSnackbar: {
      get() {
        return this.$store.state.globalSnackbar;
      },
      set(value) {
        this.$store.commit("setGlobalSnackbar", value);
      },
    },
  },
};
</script>

<style lang="scss" scoped></style>
