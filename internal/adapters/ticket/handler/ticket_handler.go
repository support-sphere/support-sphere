package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetListTicket godoc
// @Summary Create a ticket
// @Description Create a new ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body GetListTicket true "Create ticket"
// @Success 200 {object} TicketResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/tickets/list [get]
func (h *TicketHandler) GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		http.Error(w, "Invalid pageSize", http.StatusBadRequest)
		return
	}

	tickets, err := h.service.GetTickets(page, pageSize)
	if err != nil {
		http.Error(w, "Failed to get tickets", http.StatusInternalServerError)
		return
	}

	meta := map[string]interface{}{
		"page":     page,
		"pageSize": pageSize,
	}

	response := map[string]interface{}{
		"data": tickets,
		"meta": meta,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TicketHandler) GetTicketByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	ticket, err := h.service.GetTicketByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Ticket not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get ticket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ticket)
}
