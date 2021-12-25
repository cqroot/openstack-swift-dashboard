import store from "@/store";

export default function message(message, timeout = "3000") {
  store.commit("setGlobalSnackbarTimeout", timeout);
  store.commit("setGlobalSnackbarMessage", message);
  store.commit("setGlobalSnackbar", true);
}
