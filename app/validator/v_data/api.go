package v_data

type VApiWxPayAction struct {
}

type VApiLogin struct {
	Code string `json:"code" binding:"required"`
}

type VApiOrder struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
}
