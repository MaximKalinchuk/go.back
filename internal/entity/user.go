package entity

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    string    `db:"created_at"`
}
