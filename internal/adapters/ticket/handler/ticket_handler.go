package http

import (
	"encoding/json"
	"net/http"

	"github.com/support-sphere/support-sphere/internal/core/ticket/services"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

// CreateTicket godoc
// @Summary Create a ticket
// @Description Create a new ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body CreateTicketRequest true "Create ticket"
// @Success 200 {object} TicketResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/tickets [post]
func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedBy   string `json:"created_by"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ticket, err := h.service.CreateTicket(request.Title, request.Description, request.Status, request.CreatedBy)
	if err != nil {
		http.Error(w, "failed to create ticket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}
