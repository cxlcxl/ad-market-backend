package serviceuser

import (
	"errors"
	"market/app/model"
	"market/app/utils"
	"market/app/validator/v_data"
	"market/app/vars"
)

func ResetPass(params *v_data.VResetPass) error {
	if params.OldPass == params.Pass {
		return errors.New("新密码不能与旧密码一样")
	}
	user, err := model.NewUser(vars.DBMysql).FindUserById(params.User.UserId)
	if err != nil {
		return err
	}
	if utils.Password(params.OldPass, user.Secret) != user.Pass {
		return errors.New("旧密码错误")
	}
	password := utils.Password(params.Pass, user.Secret)
	err = model.NewUser(vars.DBMysql).UpdateUser(map[string]interface{}{"pass": password}, params.User.UserId)
	return err
}
