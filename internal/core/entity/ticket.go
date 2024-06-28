package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type TicketStatus string
type TicketPriority string

const (
	StatusOpen       TicketStatus = "open"
	StatusInProgress TicketStatus = "in_progress"
	StatusClosed     TicketStatus = "closed"
	StatusOnHold     TicketStatus = "on_hold"
)

const (
	PriorityLow    TicketPriority = "low"
	PriorityMedium TicketPriority = "medium"
	PriorityHigh   TicketPriority = "high"
	PriorityUrgent TicketPriority = "urgent"
)

type Ticket struct {
	TicketID    string         `json:"ticket_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Status      TicketStatus   `json:"status"`
	Priority    TicketPriority `json:"priority"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CreatedBy   uuid.UUID      `json:"created_by"`
	AssignedTo  uuid.UUID      `json:"assigned_to"`
}
