import request from "@/utils/request";

export function sysConfigs(params) {
  return request({
    url: "/config/list",
    method: "get",
    params,
  });
}

export function confCreate(data) {
  return request({
    url: "/config/create",
    method: "post",
    data,
  });
}

export function confInfo(id) {
  return request({
    url: "/config/" + id,
    method: "get",
  });
}

export function confUpdate(data) {
  return request({
    url: "/config/update",
    method: "post",
    data,
  });
}
