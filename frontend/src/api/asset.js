import request from "@/utils/request";

export function assetList(params) {
  return request({
    url: "/asset/list",
    method: "get",
    params,
  });
}

export function assetDelete(id) {
  return request({
    url: "/asset/delete/"+id,
    method: "post",
  });
}

export const assetUpload = process.env.VUE_APP_BASE_API + "/asset/upload"
export const assetInfo = process.env.VUE_APP_BASE_API + "/v1/asset/"
