package v_data

type VLessonList struct {
	Title string `json:"title" form:"title"`
	State uint8  `json:"state" form:"state"`
	Pagination
}

type VLessonCreate struct {
	Title    string       `json:"title" binding:"required"`
	OrderBy  int          `json:"order_by" binding:"required,numeric"`
	Amt      int          `json:"amt" binding:"required,numeric"`
	SubTitle string       `json:"sub_title" binding:"required"`
	ImgUrl   string       `json:"img_url" binding:"required"`
	Lists    []LessonList `json:"lists" binding:"required"`
}

type VLessonUpdate struct {
	Id       int64        `json:"id" binding:"required,numeric"`
	Title    string       `json:"title" binding:"required"`
	OrderBy  int          `json:"order_by" binding:"required,numeric"`
	State    uint8        `json:"state" binding:"numeric"`
	SubTitle string       `json:"sub_title" binding:"required"`
	ImgUrl   string       `json:"img_url" binding:"required"`
	Amt      int          `json:"amt" binding:"required,numeric"`
	Lists    []LessonList `json:"lists" binding:"required"`
}

type LessonList struct {
	Title   string `json:"title" binding:"required"`
	OrderBy int    `json:"order_by" binding:"required,numeric"`
}
