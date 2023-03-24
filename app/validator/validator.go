package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"market/app/response"
	"market/app/vars"
	"reflect"
	"regexp"
)

const (
	appCodeRule = `^[a-z0-9A-Z\_]{1,50}$`
	ticketRule  = `^[a-z0-9]+$`
	mobileRule  = `^(1)[\d]{10}$`
)

type BsValidator struct{}

var (
	customValid = map[string]validator.Func{
		"pass":   pass,
		"mobile": mobile,
	}
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for s, f := range customValid {
			_ = v.RegisterValidation(s, f)
		}
	}
}

func pass(fl validator.FieldLevel) bool {
	//if fl.Field().String() == "invalid" {
	//	return false
	//}
	return true
}

func mobile(fl validator.FieldLevel) bool {
	if match, err := regexp.MatchString(mobileRule, fl.Field().String()); err != nil {
		return false
	} else {
		return match
	}
}

func emptyValidator(_ *gin.Context, _ interface{}) error {
	return nil
}

// ctx 上下文
// v   要绑定的数据
// h   绑定完成后调用的方法
// f   自定义扩展验证规则
func bindData(ctx *gin.Context, v interface{}, h func(*gin.Context, interface{}), fs ...func(*gin.Context, interface{}) error) {
	//if err := ctx.ShouldBindBodyWith(v, binding.JSON); err != nil {
	if err := ctx.ShouldBind(v); err != nil {
		response.Fail(ctx, "验证失败："+Translate(err))
		return
	}
	for _, f := range fs {
		if err := f(ctx, v); err != nil {
			response.Fail(ctx, "验证失败："+err.Error())
			return
		}
	}

	h(ctx, v)
}

func bindRouteData(ctx *gin.Context, key string, h func(c *gin.Context, t string)) {
	h(ctx, ctx.Param(key))
}

func fillUser(ctx *gin.Context, p interface{}) error {
	if _, ok := reflect.TypeOf(p).Elem().FieldByName("User"); !ok {
		return errors.New("用户信息绑定失败，请检查是否包含 User 结构体")
	} else {
		u, _ := ctx.Get(vars.LoginUserKey)
		reflect.ValueOf(p).Elem().FieldByName("User").Set(reflect.ValueOf(u))
	}
	return nil
}
