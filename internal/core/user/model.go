package user

import "github.com/gofrs/uuid"

type User struct {
	UserID       uuid.UUID `db:"user_id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Email        string    `db:"email"`
	FullName     string    `db:"full_name"`
	Role         string    `db:"role"`
}
