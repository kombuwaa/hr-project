package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type Users struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	FullName     string    `json:"fullName" db:"full_name"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	UserRole     string    `json:"userRole" db:"user_role"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}
