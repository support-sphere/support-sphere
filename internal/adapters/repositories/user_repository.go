package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/support-sphere/support-sphere/internal/core/user"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *user.User) error {
	query := `INSERT INTO users (user_id, username, password_hash, email, full_name, role) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, user, user.Username, user.PasswordHash, user.Email, user.FullName, user.Role)
	return err
}

func (r *PostgresUserRepository) FindByUsername(username string) (*user.User, error) {
	var u user.User
	query := `SELECT user_id, username, password_hash, email, full_name, role FROM users WHERE username = $1`
	err := r.db.Get(&u, query, username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
