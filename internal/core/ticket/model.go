package ticket

import "time"

type Status string

const (
	Open       Status = "open"
	InProgress Status = "in_progress"
	Closed     Status = "closed"
	OnHold     Status = "on_hold"
)

type Priority string

const (
	Low    Priority = "low"
	Medium Priority = "medium"
	High   Priority = "high"
	Urgent Priority = "urgent"
)

type Ticket struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Priority    Priority  `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	AssignedTo  string    `json:"assigned_to"`
}
