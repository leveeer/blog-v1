import { protoObj } from "../api/https";

export function getResultCode(resultCode) {
  return protoObj.ResultCode.values[resultCode];
}

export function getReqValue(reqString) {
  return protoObj.CsId.values[reqString];
}

export function getReqString(reqID) {
  return protoObj.CsId.valuesById[reqID];
}

export function dateFormat(value) {
  const date = new Date(value * 1000);
  const y = date.getFullYear();
  let MM = date.getMonth() + 1;
  MM = MM < 10 ? ("0" + MM) : MM;
  let d = date.getDate();
  d = d < 10 ? ("0" + d) : d;
  return y + "-" + MM + "-" + d;
}



