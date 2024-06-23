package ports

import "github.com/support-sphere/support-sphere/internal/core/ticket/domain"

type TicketRepository interface {
	CreateTicket(ticket *domain.Ticket) error
}
