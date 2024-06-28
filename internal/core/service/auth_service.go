package service

import (
	"errors"

	"github.com/support-sphere/support-sphere/internal/core/entity"
	"github.com/support-sphere/support-sphere/internal/core/ports"
	"github.com/support-sphere/support-sphere/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *userService) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
