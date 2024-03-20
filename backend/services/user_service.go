package services

import (
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
	user := models.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	return s.repo.CreateUser(user)
}
