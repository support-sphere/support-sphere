package ports

import (
	"github.com/support-sphere/support-sphere/internal/core/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

type UserService interface {
	Register(user *entity.User) error
	Login(username, password string) (string, error)
}
