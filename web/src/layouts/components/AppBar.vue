<template>
  <v-app-bar
    app
    flat
    color="white"
    style="border-bottom: 1px solid #d2d2d2 !important"
  >
    <v-app-bar-nav-icon @click="setDrawer(!drawer)"></v-app-bar-nav-icon>

    <v-toolbar-title
      class="hidden-sm-and-down font-weight-light"
      v-text="$route.name"
    />

    <v-spacer></v-spacer>

    <v-menu offset-y>
      <template v-slot:activator="{ on, attrs }">
        <v-btn class="mr-2" elevation="0" v-bind="attrs" v-on="on">
          {{ target.Name }}
        </v-btn>
      </template>
      <v-list>
        <v-list-item link v-for="(target, index) in targets" :key="index">
          <v-list-item-title @click="setTarget(target)">{{
            target.Name
          }}</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <v-btn icon @click="setNight(!night)">
      <v-icon v-if="night">mdi-weather-night</v-icon>
      <v-icon v-else>mdi-weather-sunny</v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script>
import { mapState, mapMutations } from "vuex";
import axios from "axios";
import message from "@/utils/message";

export default {
  data() {
    return {
      hasError: false,
      errorMessage: "",
      targets: [],
    };
  },
  mounted() {
    axios
      .get("/v1/target")
      .then((response) => {
        this.targets = response.data;
        if (this.targets.length != 0) {
          this.setTarget(this.targets[0]);
        }
      })
      .catch((error) => {
        message(error);
      });
  },
  computed: {
    ...mapState(["drawer", "night", "target"]),
  },
  methods: {
    ...mapMutations(["setDrawer", "setNight", "setTarget"]),
  },
};
</script>
