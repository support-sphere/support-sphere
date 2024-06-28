package repository

import (
	"database/sql"

	"github.com/gofrs/uuid"
	"github.com/support-sphere/support-sphere/internal/core/entity"
	"github.com/support-sphere/support-sphere/internal/core/ports"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	uuid, _ := uuid.NewV4()

	query := `
        INSERT INTO users (
            user_id, username, password, email, first_name, last_name, role, 
            created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := r.db.Exec(query, uuid, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.Role, user.CreatedAt, user.CreatedBy, user.UpdatedAt, user.UpdatedBy, user.DeletedAt, user.DeletedBy)
	return err
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	query := `
        SELECT user_id, username, password, email, first_name, last_name, role, 
            created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
        FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Role, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy, &user.DeletedAt, &user.DeletedBy)
	if err != nil {
		return nil, err
	}
	return user, nil
}
