package services

import (
	"time"

	"github.com/support-sphere/support-sphere/internal/core/domain"
	"github.com/support-sphere/support-sphere/internal/core/ports"
)

type TicketService struct {
	repo ports.TicketRepository
}

func NewTicketService(repo ports.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ticket domain.Ticket) (string, error) {
	// Business logic for creating a ticket
	ticket.ID = "1"
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()
	return s.repo.Save(ticket)
}
