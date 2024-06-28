package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/support-sphere/support-sphere/internal/adapters/http/handlers"
)

func NewRouter(userHandler *handlers.UserHandler, ticketHandler *handlers.TicketHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/v1", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
		r.Post("/tickets", ticketHandler.CreateTicket)
		r.Get("/tickets/{ticketID}", ticketHandler.GetTicketByID)
		r.Get("/tickets", ticketHandler.GetAllTickets)

	})
	return r
}
