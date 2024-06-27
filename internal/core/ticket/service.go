package ticket

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTicket(title, description, createdBy string, priority Priority) (*Ticket, error) {
	ticket := &Ticket{
		ID:          "generated_id", // Implement logic to generate ID
		Title:       title,
		Description: description,
		Status:      Open,
		Priority:    priority,
		CreatedAt:   time.Now(),
		CreatedBy:   createdBy,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *Service) GetTicketByID(id string) (*Ticket, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetAllTickets() ([]*Ticket, error) {
	return s.repo.GetAll()
}
