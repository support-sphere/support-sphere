package ticket

import (
	"database/sql"
	"fmt"

	"github.com/support-sphere/support-sphere/internal/core/ticket/domain"
	"github.com/support-sphere/support-sphere/internal/core/ticket/ports"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) ports.TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) CreateTicket(ticket *domain.Ticket) error {
	query := `INSERT INTO tickets (title, description, status, created_at, created_by) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, ticket.Title, ticket.Description, ticket.Status, ticket.CreatedAt, ticket.CreatedBy).Scan(&ticket.ID)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %v", err)
	}
	return nil
}
func (r *TicketRepository) GetTickets(page, pageSize int) ([]domain.Ticket, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT * FROM tickets ORDER BY id LIMIT %d OFFSET %d", pageSize, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []domain.Ticket
	for rows.Next() {
		var t domain.Ticket
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.CreatedBy)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	return tickets, nil
}

func (r *TicketRepository) GetTicketByID(id int) (*domain.Ticket, error) {
	query := "SELECT * FROM tickets WHERE id = $1"

	var t domain.Ticket
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.CreatedBy)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
