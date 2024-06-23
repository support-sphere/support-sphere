package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/support-sphere/support-sphere/internal/infrastructure/delivery"
	"github.com/support-sphere/support-sphere/internal/infrastructure/delivery/routes"
	"github.com/support-sphere/support-sphere/internal/infrastructure/persistance"
)

// @title Helpdesk System Ticket API
// @version 1.0
// @description This is a helpdesk system ticket API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /v1

func main() {
	db, err := persistance.NewPostgresDB("localhost", "5432", "root", "root", "support_sphere")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(r, db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: delivery.LoggerInfoCORS(r),
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	log.Println("Server is running on :8080")

	// Create a channel to listen for the shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server: %v", err)
	}

	log.Println("Server shutdown complete")
}
