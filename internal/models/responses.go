package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

// create table if not exists responses (
// 	id uuid primary key,
// 	resume_id uuid references resume(id),
// 	vacnacy_id uuid references vacancy(id),
// 	status text check (status in ('new', 'viewed','rejected','invited')),
// 	created_at timestamp,
// 	updated_at timestamp
// );

type Responses struct {
	Id        uuid.UUID `json:"id" db:"id"`
	ResumeId  uuid.UUID `json:"resumeId" db:"resume_id"`
	VacancyId uuid.UUID `json:"vacancyId" db:"vacancy_id"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
