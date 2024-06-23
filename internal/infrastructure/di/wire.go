// internal/infrastructure/di/wire.go
//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/streadway/amqp"
	"github.com/support-sphere/support-sphere/internal/adapters/delivery/grpc"
	httpHandler "github.com/support-sphere/support-sphere/internal/adapters/delivery/http"
	"github.com/support-sphere/support-sphere/internal/adapters/event/rabbitmq"
	"github.com/support-sphere/support-sphere/internal/adapters/repository"
	"github.com/support-sphere/support-sphere/internal/core/services"
	"github.com/support-sphere/support-sphere/internal/infrastructure/database"
	"github.com/support-sphere/support-sphere/internal/infrastructure/messaging"
)

func InitializeTicketServer() (*grpc.TicketServer, error) {
	wire.Build(
		database.NewPostgresDB,
		repository.NewTicketRepository,
		services.NewTicketService,
		grpc.NewTicketServer,
		messaging.NewRabbitMQConnection,
		rabbitmq.NewTicketPublisher,
		httpHandler.NewTicketHandler,
		httpHandler.NewRouter,
	)
	return &grpc.TicketServer{}, nil
}

func NewPostgresDB() (*sql.DB, error) {
	return database.NewPostgresDB("your-database-connection-string")
}

func NewRabbitMQConnection() (*amqp.Connection, error) {
	return messaging.NewRabbitMQConnection("your-rabbitmq-url")
}
