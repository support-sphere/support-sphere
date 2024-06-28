package ports

import (
	"github.com/support-sphere/support-sphere/internal/core/entity"
)

type TicketRepository interface {
	Create(ticket *entity.Ticket) error
	GetByID(id string) (*entity.Ticket, error)
	GetAll() ([]*entity.Ticket, error)
}
type TicketService interface {
	CreateTicket(ticket *entity.Ticket) error
	GetTicketByID(id string) (*entity.Ticket, error)
	GetAllTickets() ([]*entity.Ticket, error)
}
