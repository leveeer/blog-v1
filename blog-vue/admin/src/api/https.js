import axios from "axios";
import Vue from "vue";
import protobuf from "protobufjs";
import protoRoot from "@/proto/proto";
import {getResultCode} from "../utils/util";
import {resultMap, tokenPrefix} from "../utils/constant";
import store from "../store";


// 创建 axios 实例
let service;
if (process.env.NODE_ENV === "development") {
    service = axios.create({
        baseURL: "/api", // api 的 base_url
        timeout: 5000, // 请求超时时间
        responseType: "arraybuffer",
        headers: {
            "Authorization": tokenPrefix + store.state.token,
        }
    });
} else {
  // 生产环境下
  service = axios.create({
      baseURL: "/api", // api 的 base_url
      timeout: 2000, // 请求超时时间
      responseType: "arraybuffer",
      headers: {
          "Authorization": tokenPrefix + store.state.token,
      }
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
                if (config.headers['Content-Type'] === 'multipart/form-data') return config;
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
    function (response) {
      try {
          protoObj.ResponsePkg.verify(response.data);
          const message = protoObj.ResponsePkg.decode(new Uint8Array(response.data));
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
          // if (resp.code === getResultCode(resultMap.Fail)) {
          //     Vue.prototype.$message({type: "error", message: "系统异常"});
          // } else
          if (resp.code === getResultCode(resultMap.Forbidden)) {
              Vue.prototype.$message({type: "error", message: resp.message});
          }
          return resp;
      } catch (err) {
        console.log(err);
      }
      // return response;
    },
    function (error) {
      return Promise.reject(error);
    }
);

export default service;
