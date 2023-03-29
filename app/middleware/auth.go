package middleware

import (
	"encoding/base64"
	"market/app/response"
	"market/app/service/jwt"
	"market/app/utils"
	"market/app/vars"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Headers struct {
	Authorization string `header:"Authorization"`
}

func CheckUserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := Headers{}
		if err := ctx.ShouldBindHeader(&headers); err != nil {
			response.Fail(ctx, "Token 信息读取失败")
			return
		}
		if headers.Authorization == "" {
			response.Fail(ctx, "No token information detected!")
			return
		}
		token := strings.Split(headers.Authorization, " ")
		if len(token) != 2 || len(token[1]) < 10 {
			response.Fail(ctx, "Token 信息有误")
			return
		}
		user, err := jwt.ParseUserToken(token[1])
		if err != nil {
			response.TokenExpired(ctx)
			return
		}
		loginUser := &vars.LoginUser{
			UserId:   user.Id,
			Username: user.Username,
			Mobile:   user.Mobile,
		}
		// 用户登陆信息，在控制器可以直接 get 获取
		ctx.Set(vars.LoginUserKey, loginUser)
		ctx.Next()
	}
}

func CheckApiSecret() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := Headers{}
		if err := ctx.ShouldBindHeader(&headers); err != nil {
			response.Fail(ctx, "请求失败，请重试")
			return
		}
		if headers.Authorization == "" {
			response.Fail(ctx, "请求失败，请重试!")
			return
		}

		if !checkSecret(headers.Authorization) {
			response.Fail(ctx, "请求有误，请重试")
			return
		}
		ctx.Next()
	}
}

func checkSecret(s string) bool {
	appSecret := vars.YmlConfig.GetString("AppSecret")
	decodeString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	ds, err := utils.AesDecrypt(decodeString, []byte(appSecret))
	if err != nil {
		return false
	}
	i, err := strconv.ParseInt(string(ds), 0, 64)
	if err != nil {
		return false
	}
	n := time.Now().Unix()
	if i-n <= 0 {
		return false
	}
	return true
}
