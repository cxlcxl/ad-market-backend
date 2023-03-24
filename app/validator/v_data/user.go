package v_data

import "market/app/vars"

type VUserList struct {
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile"`
	State    uint8  `json:"state" form:"state" binding:"numeric"`
	Pagination
}

type VUserCreate struct {
	Username string `json:"username" form:"username" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	State    uint8  `json:"state" form:"state" binding:"numeric"`
	Pass     string `json:"pass" form:"pass"`
}

type VUserUpdate struct {
	Id int64 `json:"id" form:"id" binding:"required,numeric"`
	VUserCreate
}

type VSelfUpdate struct {
	Username string `json:"username" form:"username" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`

	User *vars.LoginUser
}

type VLogin struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Pass   string `json:"pass" binding:"required,pass"`
}

type VResetPass struct {
	OldPass          string `json:"old_pass" binding:"required,pass"`
	Pass             string `json:"pass" binding:"required,pass"`
	ConfirmationPass string `json:"confirmation_pass" binding:"required,pass"`
	User             *vars.LoginUser
}
