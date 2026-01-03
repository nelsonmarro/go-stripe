package models

import "github.com/stripe/stripe-go/v84"

type VTState struct {
	Form                 VTForm                `json:"form"`
	Errors               VTForm                `json:"errors"`
	FormSubmitted        bool                  `json:"form_submitted"`
	IsProcessing         bool                  `json:"is_processing"`
	PaymentIntentSuccess bool                  `json:"payment_intent_success"`
	PaymentIntent        *stripe.PaymentIntent `json:"payment_intent"`
	PaymentMethod        string                `json:"payment_method"`
	PaymentAmount        string                `json:"payment_amount"`
	PaymentCurrency      string                `json:"payment_currency"`
}
