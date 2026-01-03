// Package services contains business logic for handling card-related operations.
package services

import (
	"github.com/nelsonmarro/go-stripe/internal/web/models"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/paymentintent"
)

type CardService struct{}

func NewCardService() *CardService {
	return &CardService{}
}

func (cs *CardService) Charge(card models.Card, amount int) (*stripe.PaymentIntent, string, error) {
	return cs.CreatePaymentIntent(card, amount)
}

func (cs *CardService) CreatePaymentIntent(card models.Card, amount int) (*stripe.PaymentIntent, string, error) {
	stripe.Key = card.Secret

	// Create a PaymentIntent with the specified amount and currency.
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount)),
		Currency: stripe.String(card.Currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		msg := ""
		if stripeErr, ok := err.(*stripe.Error); ok {
			msg = cardErrorMessage(stripeErr.Code)
		}
		return nil, msg, err
	}

	return pi, "", nil
}

func cardErrorMessage(code stripe.ErrorCode) string {
	var msg string

	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "The card was declined."

	case stripe.ErrorCodeExpiredCard:
		msg = "The card has expired."

	case stripe.ErrorCodeIncorrectCVC:
		msg = "The CVC code is incorrect."

	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount exceeds the maximum limit."

	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is below the minimum limit."

	case stripe.ErrorCodeProcessingError:
		msg = "An error occurred while processing the card."

	default:
		msg = "An error occurred while processing the card."
	}
	return msg
}
