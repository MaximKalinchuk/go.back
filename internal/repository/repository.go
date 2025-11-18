package repository

import (
	"github.com/jmoiron/sqlx"
	"go.back/internal/dto"
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(request dto.Register) (string, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
