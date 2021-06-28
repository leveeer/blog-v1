import axios from "axios";
import Vue from "vue";
import protobuf from "protobufjs";
import protoRoot from "@/proto/proto";


// 创建 axios 实例
let service;
if (process.env.NODE_ENV === "development") {
  service = axios.create({
    baseURL: "/api", // api 的 base_url
    timeout: 5000, // 请求超时时间
    responseType: "arraybuffer",
    headers: {
      "X-Requested-With": "XMLHttpRequest",
      "Content-Type": "application/x-protobuf"
    }
  });
} else {
  // 生产环境下
  service = axios.create({
    baseURL: "/api",
    timeout: 5000
  });
}


export const protoObj = {
  CsId: protoRoot.lookup("proto.CsId"),
  // 请求体message
  RequestPkg: protoRoot.lookupType("proto.RequestPkg"),
  // 响应体的message
  ResponsePkg: protoRoot.lookupType("proto.ResponsePkg"),
  ResultCode: protoRoot.lookup("proto.ResultCode")
};

// request 拦截器 axios 的一些配置
service.interceptors.request.use(
  config => {
    let data;
    let encode;
    switch (config.method) {
      case "post":
        data = protoObj.RequestPkg.create(config.data);
        encode = protoObj.RequestPkg.encode(data).finish();
        config.data = protobuf.util.newBuffer(encode);
        break;
      case "get":
        break;
      default:
        console.log("unKnown method type");
        break;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// response 拦截器 axios 的一些配置
service.interceptors.response.use(
  function(response) {
    try {
      protoObj.ResponsePkg.verify(response.data);
      const buffer = protobuf.util.newBuffer(response.data);
      const message = protoObj.ResponsePkg.decode(buffer);
      const resp = protoObj.ResponsePkg.toObject(message, {
        enums: Number,  // enums as string names
        longs: Number,  // longs as strings (requires long.js)
        bytes: Number,  // bytes as base64 encoded strings
        defaults: false, // includes default values
        arrays: false,   // populates empty arrays (repeated fields) even if defaults=false
        objects: false,  // populates empty objects (map fields) even if defaults=false
        oneofs: true    // includes virtual oneof fields set to the present field's name
      });
      // console.log(resp);
      if (resp.code === 1) {
        Vue.prototype.$toast({ type: "error", message: "系统异常" });
      }
      return resp;
    } catch (err) {
      console.log(err);
    }
    // return response;
  },
  function(error) {
    return Promise.reject(error);
  }
);

export default service;
