package service

import (
	"github.com/support-sphere/support-sphere/internal/core/entity"
	"github.com/support-sphere/support-sphere/internal/core/ports"
)

type TicketService struct {
	repo ports.TicketRepository
}

func NewTicketService(repo ports.TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) CreateTicket(ticket *entity.Ticket) error {
	return s.repo.Create(ticket)
}

func (s *TicketService) GetTicketByID(id string) (*entity.Ticket, error) {
	return s.repo.GetByID(id)
}

func (s *TicketService) GetAllTickets() ([]*entity.Ticket, error) {
	return s.repo.GetAll()
}
