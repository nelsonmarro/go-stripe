package receipt

import "github.com/stripe/stripe-go/v84"

type VM struct {
	CardHolder string
	Email      string
	PI         *stripe.PaymentIntent
	PMethod    string
	PAmount    string
	PCurrency  string
}

func NewVM(
	pi *stripe.PaymentIntent,
	cardHolder,
	email,
	pMethod,
	pAmount,
	pCurrency string,
) *VM {
	return &VM{
		PI:         pi,
		CardHolder: cardHolder,
		Email:      email,
		PMethod:    pMethod,
		PAmount:    pAmount,
		PCurrency:  pCurrency,
	}
}
