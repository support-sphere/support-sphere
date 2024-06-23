// cmd/server/main.go
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/support-sphere/support-sphere/internal/infrastructure/di"
	pb "github.com/support-sphere/support-sphere/proto"
	"google.golang.org/grpc"
)

func main() {
	// Initialize dependencies
	ticketServer, err := di.InitializeTicketServer()
	if err != nil {
		log.Fatalf("failed to initialize ticket server: %v", err)
	}

	// gRPC server
	grpcLis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTicketServiceServer(grpcServer, ticketServer)

	// HTTP server
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: di.NewRouter(ticketServer.HTTPHandler),
	}

	// Channel to listen for errors
	errs := make(chan error, 2)
	go func() {
		log.Printf("HTTP server listening on %s", httpServer.Addr)
		errs <- httpServer.ListenAndServe()
	}()
	go func() {
		log.Printf("gRPC server listening on %s", grpcLis.Addr())
		errs <- grpcServer.Serve(grpcLis)
	}()

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		grpcServer.GracefulStop()
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("HTTP server shutdown error: %v", err)
		}

		log.Println("Shutting down servers...")
		os.Exit(0)
	}()

	log.Printf("Exiting: %v", <-errs)
}
