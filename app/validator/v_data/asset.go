package v_data

type VAssetList struct {
	Name string `json:"name" form:"name"`
	Pagination
}

type VAssetCreate struct {
	Name string `json:"name" form:"name"`
	Pagination
}
