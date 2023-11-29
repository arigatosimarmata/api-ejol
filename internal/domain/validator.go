package domain

type Validator struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}
