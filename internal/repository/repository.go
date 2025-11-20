package repository

import (
	"github.com/jmoiron/sqlx"
	"go.back/internal/dto"
	"go.back/internal/entity"
)

type Repository struct {
	User
}

type User interface {
	CreateUser(request dto.Register) (string, error)
	GetUser(email string) (entity.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
