import axios from "axios";
import message from "@/utils/message";

const instance = axios.create({
  baseURL: "/v1",
  timeout: 1000,
});

instance.interceptors.request.use(
  (config) => {
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

instance.interceptors.response.use(
  (res) => {
    return res;
  },
  (err) => {
    message(err);
    return Promise.reject(err);
  }
);

export default instance;
