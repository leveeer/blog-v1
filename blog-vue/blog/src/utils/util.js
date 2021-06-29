import { protoObj } from "../api/https";
import protobuf from "protobufjs";

export function getResultCode(resultCode) {
  return protoObj.ResultCode.values[resultCode];
}

function getReqValue(reqString) {
  return protoObj.CsId.values[reqString];
}

function getReqString(reqID) {
  return protoObj.CsId.valuesById[reqID];
}


