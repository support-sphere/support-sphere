package ports

import "github.com/support-sphere/support-sphere/internal/core/domain"

type TicketRepository interface {
	Save(ticket domain.Ticket) (string, error)
}
