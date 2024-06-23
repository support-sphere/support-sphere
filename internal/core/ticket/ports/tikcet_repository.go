package ports

import "github.com/support-sphere/support-sphere/internal/core/ticket/domain"

type TicketRepository interface {
	CreateTicket(ticket *domain.Ticket) error
	GetTickets(page, pageSize int) ([]domain.Ticket, error)
	GetTicketByID(id int) (*domain.Ticket, error)
	GetTotalTickets() (int, error)
}
