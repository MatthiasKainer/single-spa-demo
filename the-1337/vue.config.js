module.exports = {
  lintOnSave: false,
  configureWebpack: {
    devServer: {
      headers: {
        "Access-Control-Allow-Origin": "*"
      },
      disableHostCheck: true,
      sockPort: 1337,
      sockHost: "localhost",
      https: true,
      port: 1337
    },
    externals: ["vue", "vue-router", /^@vue-mf\/.+/]
  },
  filenameHashing: false
};
