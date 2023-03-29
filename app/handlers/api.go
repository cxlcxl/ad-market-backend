package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"market/app/response"
	"market/app/validator/v_data"
	"market/app/vars"
	"market/library/curl"
)

type Api struct{}

type ApiLoginResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

func (h *Api) Login(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VApiLogin)
	xcxLoginUrl := vars.YmlConfig.GetString("XCXLogin.Url")
	appId := vars.YmlConfig.GetString("XCXLogin.AppId")
	secret := vars.YmlConfig.GetString("XCXLogin.Secret")
	data := map[string]string{"appid": appId, "secret": secret, "js_code": params.Code, "grant_type": "authorization_code"}
	var res ApiLoginResponse
	queryParams := curl.HttpBuildQuery(data)
	xcxLoginUrl = fmt.Sprintf("%s?%s", xcxLoginUrl, queryParams)
	if err := curl.New(xcxLoginUrl).Request(&res, curl.JsonHeader()); err != nil {
		response.Fail(ctx, "登陆接口调用失败："+err.Error())
	} else {
		fmt.Println(res, res.OpenId)
		if res.ErrCode == 0 {

		}
		response.Success(ctx, res)
	}
}
