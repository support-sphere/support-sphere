package services

import (
	"time"

	"github.com/support-sphere/support-sphere/internal/core/ticket/domain"
	"github.com/support-sphere/support-sphere/internal/core/ticket/ports"
)

type TicketService struct {
	repo ports.TicketRepository
}

func NewTicketService(repo ports.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(title, description, status, createdBy string) (*domain.Ticket, error) {
	ticket := &domain.Ticket{
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
	}

	if err := s.repo.CreateTicket(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}
func (s *TicketService) GetTickets(page, pageSize int) ([]domain.Ticket, error) {
	return s.repo.GetTickets(page, pageSize)
}
func (s *TicketService) GetTicketByID(id int) (*domain.Ticket, error) {
	return s.repo.GetTicketByID(id)
}
