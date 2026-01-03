package models

type VTForm struct {
	Amount     string `json:"amount"`
	CardHolder string `json:"cardHolder"`
	Email      string `json:"email"`
}
