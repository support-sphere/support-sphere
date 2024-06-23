package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	svcGrpc "github.com/support-sphere/support-sphere/internal/adapters/ticket/handler/grpc"
	pb "github.com/support-sphere/support-sphere/internal/adapters/ticket/handler/grpc/pb"
	"github.com/support-sphere/support-sphere/internal/infrastructure/delivery"
	"github.com/support-sphere/support-sphere/internal/infrastructure/delivery/routes"
	"github.com/support-sphere/support-sphere/internal/infrastructure/persistance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Start the HTTP server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50052") // Use a different port for gRPC server
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	ticketServer := svcGrpc.NewTicketServer(db)
	pb.RegisterTicketServiceServer(grpcServer, ticketServer)
	reflection.Register(grpcServer)

	log.Println("Starting gRPC server...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}

	log.Println("Server is running on :8080")

	// Create a channel to listen for the shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the HTTP server
	log.Println("Shutting down HTTP server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down HTTP server: %v", err)
	}

	// Shutdown the gRPC server
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()

	log.Println("Server shutdown complete")
}
