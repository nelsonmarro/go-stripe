// Package handlers contains HTTP handler functions for the web application.
package handlers

import (
	"log"
	"net/http"

	"github.com/nelsonmarro/go-stripe/config"
	"github.com/nelsonmarro/go-stripe/internal/web/models"
	"github.com/nelsonmarro/go-stripe/internal/web/render"
	"github.com/nelsonmarro/go-stripe/templates/components/receipt"
	"github.com/nelsonmarro/go-stripe/templates/layout"
	"github.com/nelsonmarro/go-stripe/templates/views/virtual_terminal"
	"github.com/starfederation/datastar-go/datastar"
)

type VirtualTerminalHandler struct {
	config      *config.Config
	errorLogger *log.Logger
}

func NewVirtualTerminalHandler(
	config *config.Config,
	errorLogger *log.Logger,
) *VirtualTerminalHandler {
	return &VirtualTerminalHandler{
		errorLogger: errorLogger,
		config:      config,
	}
}

func (h *VirtualTerminalHandler) GetVirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Datastar-Request") == "true" {
		err := render.PatchSPA(w, r, "Virtual Terminal", virtual_terminal.Content())
		if err != nil {
			h.errorLogger.Println(err)
		}
		return
	}

	vTermPage := virtual_terminal.Entry(h.config.Stripe.Key)
	err := vTermPage.Render(r.Context(), w)
	if err != nil {
		h.errorLogger.Println(err)
	}
}

func (h *VirtualTerminalHandler) GetPaymentForm(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	err := sse.PatchElementTempl(layout.ContentWrapper(virtual_terminal.Content()), datastar.WithSelector("#content"))
	if err != nil {
		h.errorLogger.Println(err)
	}
}

func (h *VirtualTerminalHandler) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	var input struct {
		State models.VTState `json:"state"`
	}

	if err := datastar.ReadSignals(r, &input); err != nil {
		h.errorLogger.Println("Failed to read signals:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vtState := input.State
	form := vtState.Form

	vm := receipt.NewVM(
		vtState.PaymentIntent,
		form.CardHolder,
		form.Email,
		vtState.PaymentMethod,
		vtState.PaymentAmount,
		vtState.PaymentCurrency,
	)

	// 3. Initialize Datastar SSE
	sse := datastar.NewSSE(w, r)

	receiptComponent := receipt.Receipt(vm)

	err := sse.PatchElementTempl(receiptComponent,
		datastar.WithSelector("#payment-view"))
	if err != nil {
		log.Println("Failed to patch element template:", err)
		http.Redirect(w, r, "/virtual-terminal", http.StatusSeeOther)
		return
	}
}
