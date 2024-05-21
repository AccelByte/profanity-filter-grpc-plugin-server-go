package provider

type ProfanityValidationRequest struct {
	Value string `json:"value"`
}

type ProfanityValidationResponse struct {
	IsProfane bool   `json:"isProfane"`
	Message   string `json:"message,omitempty"`
}
