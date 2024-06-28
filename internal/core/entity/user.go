package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy string    `json:"deleted_by"`
}
