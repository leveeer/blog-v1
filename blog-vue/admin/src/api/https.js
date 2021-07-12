import axios from "axios";
import Vue from "vue";
import protobuf from "protobufjs";
import protoRoot from "@/proto/proto";
import {getResultCode} from "../utils/util";
import {resultMap, tokenPrefix} from "../utils/constant";
import {refreshToken} from "./api";
import router from "../router";


// 创建 axios 实例
let service;
if (process.env.NODE_ENV === "development") {
    service = axios.create({
        baseURL: "/api", // api 的 base_url
        timeout: 5000, // 请求超时时间
        responseType: "arraybuffer",
        // headers: {
        //     "Authorization": tokenPrefix + sessionStorage.getItem("token"),
        // }
    });
} else {
  // 生产环境下
  service = axios.create({
      baseURL: "/api", // api 的 base_url
      timeout: 2000, // 请求超时时间
      responseType: "arraybuffer",
      // headers: {
      //     "Authorization": tokenPrefix + sessionStorage.getItem("token"),
      // }
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
        config.headers.Authorization = tokenPrefix + sessionStorage.getItem("token");
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
let isRefreshing = false;
//重试队列
let requests = [];
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
            if (resp.code === 401) {
                const config = response.config;
                if (!isRefreshing) {
                    isRefreshing = true;
                    return refreshToken().then(data => {
                        console.log(data);
                        // 刷新token成功，将最新的token更新到header中，同时保存在localStorage中
                        const token = data.loginResponse.token;
                        sessionStorage.setItem("token", token);
                        // store.commit("refreshToken", token);
                        // 获取当前失败的请求
                        // 重置一下配置
                        config.headers.Authorization = tokenPrefix + token;
                        // token 刷新后将数组的方法重新执行
                        requests.forEach(cb => cb(token));
                        // requests = []; // 重新请求完清空
                        // 重试当前请求并返回promise
                        return service(config)
                    }).catch(err => {
                        // sessionStorage.removeItem("token");
                        console.error('refresh_token error =>', err);
                        router.push({path: "/login"}).then(r => {
                        });
                        return Promise.reject(err)
                    }).finally(() => {
                        isRefreshing = false
                    })
                } else {
                    router.push({path: "/login"}).then(r => {});
                    // 返回未执行 resolve 的 Promise
                    return new Promise(resolve => {
                        // 用函数形式将 resolve 存入，等待刷新后再执行
                        requests.push(token => {
                            config.headers.Authorization = tokenPrefix + token;
                            resolve(service(config))
                        })
                    })
                }
            }
            if (resp.code === getResultCode(resultMap.Forbidden)) {
                Vue.prototype.$message({type: "error", message: resp.message});
            }
            return resp;
        } catch (err) {
            return Promise.reject(error)
        }
    },

    function (error) {
        return Promise.reject(error);
    }
);

export default service;
