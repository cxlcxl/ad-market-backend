package servicexcx

import (
	"errors"
	"fmt"
	"market/app/vars"
	"market/library/curl"
	"time"
)

var xcxAccessTokenCacheKey = "wxxcx:access_token"
var xcxJsApiTicketCacheKey = "wxxcx:jsapi_ticket"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessToken() (token string, err error) {
	if vars.DBRedis.HasKey(xcxAccessTokenCacheKey) {
		token = vars.DBRedis.GetString(xcxAccessTokenCacheKey)
		fmt.Println("token: ", token)
		return
	} else {
		tokenUrl := vars.YmlConfig.GetString("XCXLogin.AccessToken")
		d := map[string]string{
			"grant_type": "client_credential",
			"appid":      vars.YmlConfig.GetString("XCXLogin.AppId"),
			"secret":     vars.YmlConfig.GetString("XCXLogin.Secret"),
		}
		var res TokenResponse
		queryParams := curl.HttpBuildQuery(d)
		tokenUrl = fmt.Sprintf("%s?%s", tokenUrl, queryParams)
		if err = curl.New(tokenUrl).Request(&res, curl.JsonHeader()); err != nil {
			fmt.Println("get token: ", err)
			return
		}
		vars.DBRedis.SetString(xcxAccessTokenCacheKey, res.AccessToken, time.Second*(7140))
		return res.AccessToken, nil
	}
}

type TicketResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func GetJsApiTicket() (ticket string, err error) {
	if vars.DBRedis.HasKey(xcxJsApiTicketCacheKey) {
		ticket = vars.DBRedis.GetString(xcxJsApiTicketCacheKey)
		fmt.Println("ticket: ", ticket)
		return
	} else {
		token := ""
		token, err = GetAccessToken()
		if err != nil {
			return
		}
		ticketUrl := vars.YmlConfig.GetString("XCXLogin.JsApiTicket")
		d := map[string]string{"type": "jsapi", "access_token": token}
		var res TicketResponse
		queryParams := curl.HttpBuildQuery(d)
		ticketUrl = fmt.Sprintf("%s?%s", ticketUrl, queryParams)
		if err = curl.New(ticketUrl).Request(&res, curl.JsonHeader()); err != nil {
			fmt.Println("get ticket: ", err)
			return
		}
		vars.DBRedis.SetString(xcxJsApiTicketCacheKey, res.Ticket, time.Second*(7140))
		return res.Ticket, nil
	}
}

type MobileResponse struct {
	ErrCode   int        `json:"errcode"`
	ErrMsg    string     `json:"errmsg"`
	PhoneInfo *PhoneInfo `json:"phone_info"`
}

type PhoneInfo struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

func GetUserPhoneNumber(code string) (mobile *PhoneInfo, err error) {
	token, err := GetAccessToken()
	if err != nil {
		return
	}
	mobileUrl := vars.YmlConfig.GetString("XCXLogin.UserMobile")
	d := map[string]string{"access_token": token, "code": code}
	var res MobileResponse
	queryParams := curl.HttpBuildQuery(d)
	mobileUrl = fmt.Sprintf("%s?%s", mobileUrl, queryParams)
	if err = curl.New(mobileUrl).Request(&res, curl.JsonHeader()); err != nil {
		return
	}
	if res.ErrCode != 0 {
		return nil, errors.New("请求失败：" + res.ErrMsg)
	}

	return res.PhoneInfo, nil
}
