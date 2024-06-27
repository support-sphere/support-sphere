package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Register(user *User) error {
	return s.repository.Create(user)
}

func (s *Service) Login(username, password string) (*User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials")
	}

	return user, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
