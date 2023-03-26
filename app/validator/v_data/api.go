package v_data

type VApiLogin struct {
	Code string `json:"code" binding:"required"`
}
