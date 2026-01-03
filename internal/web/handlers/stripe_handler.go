// Package handlers defines the handlers for the API endpoints.
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nelsonmarro/go-stripe/config"
	"github.com/nelsonmarro/go-stripe/internal/web/models"
	"github.com/nelsonmarro/go-stripe/internal/web/services"
	"github.com/starfederation/datastar-go/datastar"
)

type StripeHandler struct {
	config      *config.Config
	cardService *services.CardService
	errorLogger *log.Logger
}

func NewStripeHandler(
	cardService *services.CardService,
	config *config.Config,
	errorLogger *log.Logger,
) *StripeHandler {
	return &StripeHandler{
		errorLogger: errorLogger,
		config:      config,
		cardService: cardService,
	}
}

func (sh *StripeHandler) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		State models.VTState `json:"state"`
	}

	if err := datastar.ReadSignals(r, &input); err != nil {
		sh.errorLogger.Println("Failed to read signals:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vtState := input.State
	form := vtState.Form

	// 2. Upgrade connection to SSE
	sse := datastar.NewSSE(w, r)

	amount, err := strconv.Atoi(form.Amount)
	if err != nil {
		sh.errorLogger.Println(err)
		_ = sse.ConsoleError(fmt.Errorf("invalid amount format: %w", err))
		return
	}

	// 4. Perform Stripe Logic
	// Assuming currency is USD for this example
	card := models.Card{
		Secret:   sh.config.Stripe.Secret,
		Key:      sh.config.Stripe.Key,
		Currency: "usd",
	}

	// set processing state
	vtState.IsProcessing = true
	_ = sse.MarshalAndPatchSignals(map[string]any{
		"state": vtState,
	})

	pi, _, err := sh.cardService.Charge(card, amount)

	// 5. Send events back to the client based on the outcome
	if err != nil {
		vtState.PaymentIntent = nil
		vtState.PaymentIntentSuccess = false
		vtState.IsProcessing = false
	} else {
		// On success: send the payment intent to the client to proceed with confirmation
		vtState.PaymentIntent = pi
		vtState.PaymentIntentSuccess = true
		// Keep IsProcessing true until the frontend completes the confirmation
	}

	// Patch the nested state signal
	_ = sse.MarshalAndPatchSignals(map[string]any{
		"state": vtState,
	})
}
