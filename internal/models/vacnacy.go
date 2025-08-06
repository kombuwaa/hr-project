package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/satori-uuid"
)

type Vacancy struct {
	Id                 uuid.UUID `json:"id" db:"id"`
	CompanyId          uuid.UUID `json:"companyId" db:"company_id"`
	Title              string    `json:"title" db:"title"`
	Description        string    `json:"description" db:"description"`
	Location           string    `json:"location" db:"location"`
	ExperienceRequired string    `json:"experienceRequired" db:"experience_required"`
	SkillsRequired     []string  `json:"skillsRequired" db:"skills_required"`
	SalaryFrom         int       `json:"salaryFrom" db:"salary_from"`
	SalaryTo           int       `json:"salaryTo" db:"salary_to"`
	CreatedAt          time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt          time.Time `json:"updatedAt" db:"updated_at"`
}
