package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type Resume struct {
	Id         uuid.UUID `json:"id" db:"id"`
	UserId     uuid.UUID `json:"userId" db:"user_id"`
	Title      string    `json:"title" db:"title"`
	Experience string    `json:"experienceYear" db:"experience_year"`
	Skills     string    `json:"skills" db:"skills"`
	Education  string    `json:"education" db:"education"`
	About      string    `json:"about" db:"about"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}
