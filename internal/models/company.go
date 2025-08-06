package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type Company struct {
	Id         uuid.UUID `json:"id" db:"id"`
	FullName   string    `json:"fullName" db:"full_name"`
	Email      string    `json:"email" db:"email"`
	Website    string    `json:"website" db:"website"`
	OwnerId    uuid.UUID `json:"ownerId" db:"owner_id"`
	UserRating float32   `json:"userRating" db:"user_rating"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}
