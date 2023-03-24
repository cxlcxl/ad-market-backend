package handlers

import (
	"github.com/gin-gonic/gin"
	"market/app/model"
	"market/app/response"
	"market/app/service/jwt"
	serviceuser "market/app/service/user"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
	"strconv"
	"time"
)

type User struct{}

func (l *User) UserList(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VUserList)
	offset := utils.GetPages(params.Page, params.PageSize)
	users, total, err := model.NewUser(vars.DBMysql).List(params.Username, params.Mobile, params.State, offset, params.PageSize)
	if err != nil {
		response.Fail(ctx, "查询错误: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"total": total, "list": users})
}

func (l *User) Login(ctx *gin.Context, p interface{}) {
	var params = p.(*v_data.VLogin)
	user, err := model.NewUser(vars.DBMysql).FindUserByMobile(params.Mobile)
	if err != nil {
		response.Fail(ctx, "登录失败: "+err.Error())
		return
	}
	if user.State != 1 {
		response.Fail(ctx, "账号已失效不可登陆")
		return
	}
	if user.Pass != utils.Password(params.Pass, user.Secret) {
		response.Fail(ctx, "密码错误")
		return
	}
	token, err := jwt.CreateUserToken(user.Id, user.Username, user.Mobile)
	if err != nil {
		response.Fail(ctx, "登录失败: "+err.Error())
		return
	}
	if _, err = jwt.ParseUserToken(token); err != nil {
		response.Fail(ctx, "TOKEN 生成失败: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{"token": token})
}

func (l User) Profile(ctx *gin.Context) {
	user, exists := ctx.Get(vars.LoginUserKey)
	if !exists {
		response.Fail(ctx, "用户信息获取失败")
		return
	}
	response.Success(ctx, gin.H{
		"user_id":  user.(*vars.LoginUser).UserId,
		"username": user.(*vars.LoginUser).Username,
		"mobile":   user.(*vars.LoginUser).Mobile,
	})
}

func (l User) Logout(ctx *gin.Context) {
	response.Success(ctx, nil)
}

func (l *User) UserInfo(ctx *gin.Context, v string) {
	id, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		response.Fail(ctx, "参数错误")
		return
	}
	user, err := model.NewUser(vars.DBMysql).FindUserById(id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	response.Success(ctx, user)
}

func (l *User) UserCreate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VUserCreate)
	s := utils.GenerateSecret(0)
	user := &model.User{
		Username: params.Username,
		Mobile:   params.Mobile,
		State:    1,
		Secret:   s,
		Pass:     utils.Password(params.Pass, s),
		Timestamp: model.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err := model.NewUser(vars.DBMysql).CreateUser(user)
	if err != nil {
		response.Fail(ctx, "创建失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *User) UserUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VUserUpdate)
	user, err := model.NewUser(vars.DBMysql).FindUserById(params.Id)
	if err != nil {
		response.Fail(ctx, "请求错误："+err.Error())
		return
	}
	d := map[string]interface{}{
		"username":   params.Username,
		"mobile":     params.Mobile,
		"state":      params.State,
		"updated_at": time.Now(),
	}
	if params.Pass != "" {
		d["pass"] = utils.Password(params.Pass, user.Secret)
	}
	err = model.NewUser(vars.DBMysql).UpdateUser(d, params.Id)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l *User) VSelfUpdate(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VSelfUpdate)
	d := map[string]interface{}{
		"username":   params.Username,
		"mobile":     params.Mobile,
		"updated_at": time.Now(),
	}
	err := model.NewUser(vars.DBMysql).UpdateUser(d, params.User.UserId)
	if err != nil {
		response.Fail(ctx, "修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (l User) ResetPass(ctx *gin.Context, p interface{}) {
	params := p.(*v_data.VResetPass)
	if err := serviceuser.ResetPass(params); err != nil {
		response.Fail(ctx, "密码修改失败："+err.Error())
		return
	}
	response.Success(ctx, nil)
}
