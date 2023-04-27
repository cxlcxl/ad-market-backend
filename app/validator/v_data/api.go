package v_data

type VApiWxPayAction struct {
	Id           string        `json:"id"`
	ResourceType string        `json:"resource_type"`
	EventType    string        `json:"event_type"`
	Summary      string        `json:"summary"`
	Resource     *VPayResource `json:"resource"`
}

type VPayResource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type VApiLogin struct {
	Code string `json:"code" binding:"required"`
}

type VApiOrder struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
}

type VApiOrderQuery struct {
	Sn string `json:"sn" binding:"required"`
}
