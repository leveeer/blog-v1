import axios from "axios";
import Vue from "vue";

// 创建 axios 实例
let service;
if (process.env.NODE_ENV === "development") {
  service = axios.create({
    baseURL: "/api", // api 的 base_url
    timeout: 5000 // 请求超时时间
  });
} else {
  // 生产环境下
  service = axios.create({
    baseURL: "/api",
    timeout: 5000,
  });
}

// request 拦截器 axios 的一些配置
service.interceptors.request.use(
  config => {
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// response 拦截器 axios 的一些配置
service.interceptors.response.use(
  function(response) {
    switch (response.data.code) {
      case 50000:
        Vue.prototype.$toast({ type: "error", message: "系统异常" });
    }
    return response;
  },
  function(error) {
    return Promise.reject(error);
  }
);

export default service;
