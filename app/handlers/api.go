package handlers

import (
	"fmt"
	"market/app/response"
	servicexcx "market/app/service/xcx"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"market/library/curl"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct{}

type ApiLoginResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

type ApiSchemeResponse struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	OpenLink string `json:"openlink"`
}

type ApiUrlLinkResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UrlLink string `json:"url_link"`
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

func (h *Api) GetUrlLink(ctx *gin.Context) {
	token, err := servicexcx.GetAccessToken()
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	xcxSchemeUrl := vars.YmlConfig.GetString("XCXLogin.UrlLink")
	data := map[string]string{"path": "pages/friend/friend"}
	var res ApiUrlLinkResponse
	c, err := curl.New(xcxSchemeUrl + token).Post().JsonData(data)
	if err != nil {
		response.Fail(ctx, "请求失败："+err.Error())
		return
	}
	if err := c.Request(&res, curl.JsonHeader()); err != nil {
		response.Fail(ctx, "登陆接口调用失败："+err.Error())
	} else {
		fmt.Println(res, res.UrlLink)
		if res.ErrCode == 0 {
			response.Success(ctx, res.UrlLink)
		} else {
			response.Fail(ctx, "获取失败："+res.ErrMsg)
		}
	}
}

func (h *Api) XcxSdk(ctx *gin.Context) {
	nonce := utils.GenerateSecret(16)
	ticket, err := servicexcx.GetJsApiTicket()
	if err != nil {
		response.Fail(ctx, "SDK 获取失败，请刷新重试"+err.Error())
		return
	}
	timestamp := time.Now().Unix()
	h5Url := vars.YmlConfig.GetString("XCXLogin.H5Domain")
	s := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", nonce, ticket, timestamp, h5Url)
	response.Success(ctx, gin.H{
		"app_id":    vars.YmlConfig.GetString("XCXLogin.AppId"),
		"timestamp": timestamp,
		"nonce":     nonce,
		"signature": utils.Sha1(s),
	})
}
