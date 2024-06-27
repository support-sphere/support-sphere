package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/support-sphere/support-sphere/internal/app"
	"github.com/support-sphere/support-sphere/internal/core/ticket"
)

type TicketHandler struct {
	app *app.App
}

func NewTicketHandler(app *app.App) *TicketHandler {
	return &TicketHandler{app: app}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string          `json:"title"`
		Description string          `json:"description"`
		Priority    ticket.Priority `json:"priority"`
		CreatedBy   string          `json:"created_by"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ticket, err := h.app.TicketService().CreateTicket(input.Title, input.Description, input.CreatedBy, input.Priority)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Return the created ticket as JSON response
	json.NewEncoder(w).Encode(ticket)
}
