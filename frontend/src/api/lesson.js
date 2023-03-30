import request from "@/utils/request";

export function lessonList(params) {
  return request({
    url: "/lesson/list",
    method: "get",
    params,
  });
}

export function lessonCreate(data) {
  return request({
    url: "/lesson/create",
    method: "post",
    data,
  });
}

export function lessonUpdate(data) {
  return request({
    url: "/lesson/update",
    method: "post",
    data,
  });
}

export function lessonInfo(id) {
  return request({
    url: "/lesson/"+id,
    method: "get",
  });
}
