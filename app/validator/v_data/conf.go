package v_data

type VConfigList struct {
	Key   string `form:"_k"`
	Desc  string `form:"_desc"`
	State uint8  `form:"state"`
	Pagination
}

type VConfigCreate struct {
	Key    string `json:"key" binding:"required"`
	Desc   string `json:"desc" binding:"required"`
	Val    string `json:"val" binding:"required"`
	Bak1   string `json:"bak1"`
	Bak2   string `json:"bak2"`
	Remark string `json:"remark"`
}

type VConfigUpdate struct {
	Id    int64 `json:"id" binding:"required"`
	State uint8 `json:"state" binding:"numeric"`
	VConfigCreate
}
