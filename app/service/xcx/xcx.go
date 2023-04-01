package servicexcx

import (
	"errors"
	"fmt"
	"market/app/vars"
	"market/library/curl"
	"time"
)

var xcxAccessTokenCacheKey = "wxxcx:access_token"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessToken() (token string, err error) {
	if vars.DBRedis.HasKey(xcxAccessTokenCacheKey) {
		token = vars.DBRedis.GetString(xcxAccessTokenCacheKey)
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
			return
		}
		vars.DBRedis.SetString(xcxAccessTokenCacheKey, res.AccessToken, time.Second*(7140))
		return res.AccessToken, nil
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
