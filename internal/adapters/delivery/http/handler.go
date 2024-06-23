package http

import (
	"encoding/json"
	"net/http"

	"github.com/support-sphere/support-sphere/internal/core/domain"
	"github.com/support-sphere/support-sphere/internal/core/services"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket domain.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTicket(ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}
