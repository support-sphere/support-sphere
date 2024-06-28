package repository

import (
	"database/sql"
	"time"

	"github.com/support-sphere/support-sphere/internal/core/entity"
	"github.com/support-sphere/support-sphere/internal/core/ports"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) ports.TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ticket *entity.Ticket) error {
	query := `
		INSERT INTO tickets (ticket_id, title, description, status, priority, created_at, updated_at, created_by, assigned_to)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.Exec(query, ticket.TicketID, ticket.Title, ticket.Description, ticket.Status, ticket.Priority, time.Now(), time.Now(), ticket.CreatedBy, ticket.AssignedTo)
	return err
}

func (r *TicketRepository) GetByID(id string) (*entity.Ticket, error) {
	query := `SELECT ticket_id, title, description, status, priority, created_at, updated_at, created_by, assigned_to FROM tickets WHERE ticket_id = $1`
	row := r.db.QueryRow(query, id)

	ticket := &entity.Ticket{}
	err := row.Scan(&ticket.TicketID, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Priority, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.CreatedBy, &ticket.AssignedTo)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *TicketRepository) GetAll() ([]*entity.Ticket, error) {
	query := `SELECT ticket_id, title, description, status, priority, created_at, updated_at, created_by, assigned_to FROM tickets`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*entity.Ticket
	for rows.Next() {
		ticket := &entity.Ticket{}
		if err := rows.Scan(&ticket.TicketID, &ticket.Title, &ticket.Description, &ticket.Status, &ticket.Priority, &ticket.CreatedAt, &ticket.UpdatedAt, &ticket.CreatedBy, &ticket.AssignedTo); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
