import { wsURL } from "../utils/constant";
import { protoObj } from "./https";
import protobuf from "protobufjs";
import context from "../main.js";

let websocket = null;
let heartBeat = null;

export function sendMessage(message) {
  console.log("发送给服务器消息：", message);
  let data = protoObj.RequestPkg.create(message);
  let encode = protoObj.RequestPkg.encode(data).finish();
  message = protobuf.util.newBuffer(encode);
  websocket.send(message);
}

function create() {
  webSocket();
}


function webSocket() {
  websocket = new WebSocket(wsURL);
  websocket.binaryType = "arraybuffer";
  websocket.onerror = function(event) {
    console.log(event);
    alert("失败");
  };
// 连接成功建立的回调方法
  websocket.onopen = function(event) {
    // // 发送心跳消息(废弃)  心跳包由服务端发送
    // heartBeat = setInterval(function() {
    //   const beatMessage = {
    //     type: 6,
    //     data: "ping"
    //   };
    //   sendMessage({ csBeatMessage: beatMessage });
    // }, 60 * 1000);
  };
// 接收到消息的回调方法
  websocket.onmessage = function(event) {
    protoObj.ResponsePkg.verify(event.data);
    const buffer = protobuf.util.newBuffer(event.data);
    const message = protoObj.ResponsePkg.decode(buffer);
    const data = protoObj.ResponsePkg.toObject(message, {
      enums: Number,  // enums as string names
      longs: Number,  // longs as strings (requires long.js)
      bytes: Number,  // bytes as base64 encoded strings
      defaults: false, // includes default values
      arrays: false,   // populates empty arrays (repeated fields) even if defaults=false
      objects: false,  // populates empty objects (map fields) even if defaults=false
      oneofs: true    // includes virtual oneof fields set to the present field's name
    });
    console.log("收到服务器的消息：", data);
    switch (data.scChat.type) {
      case 1:
        // 在线人数
        context.$store.commit("updateOnline", data.scChat.scChatOnline.online);
        break;
      // case 2:
      //   // 历史记录
      //   that.chatRecordList = data.data.chatRecordList;
      //   that.chatRecordList.forEach(item => {
      //     if (item.type === 5) {
      //       that.voiceList.push(item.id);
      //     }
      //   });
      //   that.ipAddr = data.data.ipAddr;
      //   that.ipSource = data.data.ipSource;
      //   break;
      case 3:
        // 文字消息
        context.$store.commit("updateChat", data.scChat.scChatMessage);
        break;
      // case 4:
      //   // 撤回
      //   if (data.data.isVoice) {
      //     that.voiceList.splice(that.voiceList.indexOf(data.data.id), 1);
      //   }
      //   for (let i = 0; i < that.chatRecordList.length; i++) {
      //     if (that.chatRecordList[i].id === data.data.id) {
      //       that.chatRecordList.splice(i, 1);
      //       i--;
      //     }
      //   }
      //   break;
      // case 5:
      //   // 语音消息
      //   that.voiceList.push(data.data.id);
      //   that.chatRecordList.push(data.data);
      //   if (!that.isShow) {
      //     that.unreadCount++;
      //   }
      //   break;
    }
  };
//连接关闭的回调方法
  websocket.onclose = function() {
    clearInterval(heartBeat);
  };
}


export default create;