package v_data

type VAccountList struct {
	AccountName string `form:"account_name,optional"`
	Mobile      string `form:"mobile,optional"`
	State       uint8  `form:"state,optional"`
	Page        int64  `form:"page"`
	PageSize    int64  `form:"page_size"`
}

type VAccountUpdate struct {
	Id          int64  `json:"id" binding:"required"`
	AccountName string `json:"account_name" binding:"required"`
	Mobile      string `json:"mobile" binding:"required"`
	Remark      string `json:"remark"`
	State       uint8  `json:"state"`
}
