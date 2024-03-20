package repositories

import (
	"database/sql"

	"github.com/K1ose/bdsms/backend/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
}

type postgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db}
}

func (r *postgresUserRepository) CreateUser(user models.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password_hash, created_at) VALUES ($1, $2, $3, $4)", user.Username, user.Email, user.PasswordHash, user.CreatedAt)
	return err
}
