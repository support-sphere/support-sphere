package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	handlerTicket "github.com/support-sphere/support-sphere/internal/adapters/ticket/handler"
	repoTicket "github.com/support-sphere/support-sphere/internal/adapters/ticket/repository"
	"github.com/support-sphere/support-sphere/internal/core/ticket/services"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	apiV1 := r.PathPrefix("/v1").Subrouter()
	registerV1Routes(apiV1, db)
}

func registerV1Routes(r *mux.Router, db *sql.DB) {
	ticketRepo := repoTicket.NewPostgresTicketRepository(db)
	ticketService := services.NewTicketService(ticketRepo)
	ticketHandler := handlerTicket.NewTicketHandler(ticketService)

	r.HandleFunc("/tickets", ticketHandler.CreateTicket).Methods("POST")
}
