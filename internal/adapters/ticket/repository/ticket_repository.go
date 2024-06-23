package ticket

import (
	"database/sql"
	"fmt"

	"github.com/support-sphere/support-sphere/internal/core/ticket/domain"
	"github.com/support-sphere/support-sphere/internal/core/ticket/ports"
)

type PostgresTicketRepository struct {
	db *sql.DB
}

func NewPostgresTicketRepository(db *sql.DB) ports.TicketRepository {
	return &PostgresTicketRepository{db: db}
}

func (r *PostgresTicketRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `INSERT INTO tickets (title, description, status, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, ticket.Title, ticket.Description, ticket.Status, ticket.CreatedAt, ticket.CreatedBy).Scan(&ticket.ID)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %v", err)
	}
	return nil
}
