package services

import (
	"time"

	"github.com/K1ose/bdsms/backend/models"
	"github.com/K1ose/bdsms/backend/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(username, email, passwordHash string) error {

	createdAt := time.Now()

	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
	}
	return s.repo.CreateUser(user)
}
