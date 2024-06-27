package ticket

type Repository interface {
	Create(ticket *Ticket) error
	Update(ticket *Ticket) error
	Delete(id string) error
	GetByID(id string) (*Ticket, error)
	GetAll() ([]*Ticket, error)
}
