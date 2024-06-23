package http

import (
	"github.com/go-chi/chi"
)

func NewRouter(ticketHandler TicketHandler) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/tickets", ticketHandler.CreateTicket)
	return router
}
