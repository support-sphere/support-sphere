package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
