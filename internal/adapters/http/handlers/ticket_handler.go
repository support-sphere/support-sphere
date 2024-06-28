package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/support-sphere/support-sphere/internal/core/entity"
	"github.com/support-sphere/support-sphere/internal/core/ports"
)

type TicketHandler struct {
	service ports.TicketService
}

func NewTicketHandler(service ports.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entity.User)

	var ticket entity.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		log.Printf("Error decoding ticket request body: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ticket.CreatedBy = user.UserID
	if err := h.service.CreateTicket(&ticket); err != nil {
		log.Printf("Error creating ticket: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func (h *TicketHandler) GetTicketByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "ticketID")
	ticket, err := h.service.GetTicketByID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ticket)
}

func (h *TicketHandler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.service.GetAllTickets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tickets)
}
