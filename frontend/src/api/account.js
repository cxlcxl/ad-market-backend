import request from "@/utils/request";

export function accountUpdate(data) {
  return request({
    url: "/account/update",
    method: "post",
    data,
  });
}

export function accountList(params) {
  return request({
    url: "/account/list",
    method: "get",
    params,
  });
}

export function accountInfo(account_id) {
  return request({
    url: "/account/" + account_id,
    method: "get",
  });
}
