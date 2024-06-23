package repository

import (
	"database/sql"

	"github.com/support-sphere/support-sphere/internal/core/domain"
	"github.com/support-sphere/support-sphere/internal/core/ports"

	_ "github.com/lib/pq"
)

type TicketRepositoryImpl struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) ports.TicketRepository {
	return &TicketRepositoryImpl{db: db}
}

func (r *TicketRepositoryImpl) Save(ticket domain.Ticket) (string, error) {
	query := `INSERT INTO tickets (id, title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, ticket.ID, ticket.Title, ticket.Description, ticket.Status, ticket.CreatedAt, ticket.UpdatedAt)
	if err != nil {
		return "", err
	}
	return ticket.ID, nil
}
