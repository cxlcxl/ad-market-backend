package v_data

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

type VLogin struct {
	Mobile string `json:"mobile" binding:"required,email"`
	Pass   string `json:"pass" binding:"required,pass"`
}
