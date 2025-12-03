package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.back/configs"
	authdto "go.back/internal/dto/auth"
	"go.back/internal/entity"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(request authdto.Register) (string, error) {
	var id string

	query := fmt.Sprintf("INSERT INTO %s (id, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", configs.UsersTable)

	row := u.db.QueryRow(query, uuid.New().String(), request.Username, request.Email, request.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (u *UserRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", configs.UsersTable)
	err := u.db.Get(&user, query, email)

	return user, err
}

func (u *UserRepository) GetUserById(id string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", configs.UsersTable)
	err := u.db.Get(&user, query, id)

	return user, err
}
