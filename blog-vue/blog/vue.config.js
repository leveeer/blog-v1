const sourceMap = process.env.NODE_ENV === "development";
module.exports = {
  transpileDependencies: ["vuetify"],
  // 基本路径
  publicPath: "./",
  // 输出文件目录
  outputDir: "dist",
  // eslint-loader 是否在保存的时候检查
  lintOnSave: false,

  devServer: {
    proxy: {
      // 设置代理
      // proxy all requests starting with /api to jsonplaceholder
      "/api": {
        // target: "http://localhost:8088",
        target: "http://47.119.168.227:8088",
        changeOrigin: true,
        ws: true,
        pathRewrite: {
          "^/api": ""
        }
      }
    },
    disableHostCheck: true,
    before: app => {
    }
  },

  // 生产环境是否生成 sourceMap 文件
  productionSourceMap: sourceMap,

  // 第三方插件配置
  pluginOptions: {
    // ...
  }
};
