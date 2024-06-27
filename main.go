package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	https "github.com/support-sphere/support-sphere/internal/adapters/delivery/http"
	"github.com/support-sphere/support-sphere/internal/adapters/delivery/http/handlers"
	"github.com/support-sphere/support-sphere/internal/adapters/repositories"
	"github.com/support-sphere/support-sphere/internal/app"
	"github.com/support-sphere/support-sphere/internal/core/ticket"
	"github.com/support-sphere/support-sphere/internal/core/user"
)

func main() {
	db, err := sqlx.Connect("pgx", "postgres://username:password@localhost:5432/yourdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewPostgresUserRepository(db)
	ticketRepo := repositories.NewTicketRepository(db)

	userService := user.NewService(userRepo)
	ticketService := ticket.NewService(ticketRepo)

	app := app.NewApp(*userService, *ticketService)
	userHandler := handlers.NewUserHandler(app)
	ticketHandler := handlers.NewTicketHandler(app)

	router := https.NewRouter(userHandler, ticketHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shut down the server: %v\n", err)
		}
		log.Println("Server stopped")
	}()
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}
