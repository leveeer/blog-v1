module.exports = {
  devServer: {
    proxy: {
      "/api": {
        target: "http://localhost:8088",
        changeOrigin: true,
        pathRewrite: {
          "^/api": ""
        }
      }
    },
    disableHostCheck: true
  },
  chainWebpack: config => {
    config.resolve.alias.set("@", resolve("src"));
  },
  lintOnSave: false //关闭eslint检查
};

const path = require("path");
function resolve(dir) {
  return path.join(__dirname, dir);
}
