import request from "@/utils/request";

export function listenList(params) {
  return request({
    url: "/listen/list",
    method: "get",
    params,
  });
}

export function listenCreate(data) {
  return request({
    url: "/listen/create",
    method: "post",
    data,
  });
}

export function listenUpdate(data) {
  return request({
    url: "/listen/update",
    method: "post",
    data,
  });
}

export function listenInfo(id) {
  return request({
    url: "/listen/"+id,
    method: "get",
  });
}
