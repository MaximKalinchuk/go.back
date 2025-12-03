package repository

import (
	"github.com/jmoiron/sqlx"
	authdto "go.back/internal/dto/auth"
	"go.back/internal/entity"
)

type Repository struct {
	User
}

type User interface {
	CreateUser(request authdto.Register) (string, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserById(email string) (entity.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
