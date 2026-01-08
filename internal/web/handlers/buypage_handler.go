package handlers

import (
	"log"
	"net/http"

	"github.com/nelsonmarro/go-stripe/config"
	"github.com/nelsonmarro/go-stripe/internal/web/services"
	"github.com/nelsonmarro/go-stripe/templates/pages/buypage"
)

type BuyPageHandler struct {
	config      *config.Config
	cardService *services.CardService
	errorLogger *log.Logger
}

func NewBuyPageHandler(
	cardService *services.CardService,
	config *config.Config,
	errorLogger *log.Logger,
) *BuyPageHandler {
	return &BuyPageHandler{
		errorLogger: errorLogger,
		config:      config,
		cardService: cardService,
	}
}

func (h *BuyPageHandler) GetPage(w http.ResponseWriter, r *http.Request) {
	buyPage := buypage.BuyPage(h.config.Stripe.Key)
	err := buyPage.Render(r.Context(), w)
	if err != nil {
		h.errorLogger.Println(err)
	}
}

func (h *BuyPageHandler) ChargeOnce(w http.ResponseWriter, r *http.Request) {
}
