module.exports = {
  transpileDependencies: ["vuetify"],
  publicPath: "./",
  configureWebpack: {
    devServer: {
      proxy: {
        "/v1": {
          target: "http://127.0.0.1:8088",
          changeOrigin: true,
          ws: true,
        },
      },
    },
  },
};
