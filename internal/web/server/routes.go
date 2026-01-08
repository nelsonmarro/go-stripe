package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nelsonmarro/go-stripe/internal/web/handlers"
	"github.com/nelsonmarro/go-stripe/internal/web/services"
)

func (s *Server) getRoutes() http.Handler {
	mux := chi.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	vTermHandler := s.createVirtualTermHandler()
	cardPageHandler := s.createBuyPageHandler()
	stripeHandler := s.createStripeHandler()

	// UI Routes
	mux.Get("/virtual-terminal", vTermHandler.GetVirtualTerminal)
	mux.Get("/buy-page", cardPageHandler.GetPage)

	// Backend Routes
	mux.Post("/payment-succeeded", vTermHandler.PaymentSucceeded)
	mux.Post("/payment-intent", stripeHandler.GetPaymentIntent)

	return mux
}

func (s *Server) createStripeHandler() *handlers.StripeHandler {
	cardService := services.NewCardService()
	stripeHandler := handlers.NewStripeHandler(cardService, s.Config, s.ErrorLog)
	return stripeHandler
}

func (s *Server) createVirtualTermHandler() *handlers.VirtualTerminalHandler {
	vTermHandler := handlers.NewVirtualTerminalHandler(s.Config, s.ErrorLog)
	return vTermHandler
}

func (s *Server) createBuyPageHandler() *handlers.BuyPageHandler {
	buyPageHandler := handlers.NewBuyPageHandler(
		services.NewCardService(),
		s.Config,
		s.ErrorLog,
	)
	return buyPageHandler
}
