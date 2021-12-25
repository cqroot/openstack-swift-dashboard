import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    drawer: null,
    night: false,
    globalSnackbar: false,
    globalSnackbarMessage: "",
    globalSnackbarTimeout: "3000",
    target: { Name: "", ID: -1 },
  },
  mutations: {
    setDrawer(state, payload) {
      state.drawer = payload;
    },
    setNight(state, payload) {
      state.night = payload;
    },
    setGlobalSnackbar(state, payload) {
      state.globalSnackbar = payload;
    },
    setGlobalSnackbarMessage(state, payload) {
      state.globalSnackbarMessage = payload;
    },
    setGlobalSnackbarTimeout(state, payload) {
      state.globalSnackbarTimeout = payload;
    },
    setTarget(state, payload) {
      state.target = payload;
    },
  },
  actions: {},
  modules: {},
});
