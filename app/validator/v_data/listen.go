package v_data

type VListenList struct {
	Title string `json:"title" form:"title"`
	State uint8  `json:"state" form:"state"`
	Pagination
}

type VListenCreate struct {
	Title    string       `json:"title" binding:"required"`
	OrderBy  int          `json:"order_by" binding:"required,numeric"`
	SubTitle string       `json:"sub_title" binding:"required"`
	Lists    []ListenList `json:"lists" binding:"required"`
}

type VListenUpdate struct {
	Id       int64        `json:"id" binding:"required,numeric"`
	Title    string       `json:"title" binding:"required"`
	OrderBy  int          `json:"order_by" binding:"required,numeric"`
	State    uint8        `json:"state" binding:"numeric"`
	SubTitle string       `json:"sub_title" binding:"required"`
	Lists    []ListenList `json:"lists" binding:"required"`
}

type ListenList struct {
	Title   string `json:"title" binding:"required"`
	OrderBy int    `json:"order_by" binding:"required,numeric"`
}
