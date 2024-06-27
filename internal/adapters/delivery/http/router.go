package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/support-sphere/support-sphere/internal/adapters/delivery/http/handlers"
	"github.com/support-sphere/support-sphere/pkg/middleware"
)

func NewRouter(userHandler *handlers.UserHandler, ticketHandler *handlers.TicketHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Register)
		r.Use(middleware.AuthMiddleware)
		r.Get("/", userHandler.Login)
	})

	r.Route("/tickets", func(r chi.Router) {
		r.Post("/", ticketHandler.CreateTicket)
		r.Route("/{ticketID}", func(r chi.Router) {
		})
	})

	return r
}
