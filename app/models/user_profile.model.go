package models

import (
	"github.com/google/uuid"
	"time"
)

type UserProfile struct {
	ID         uuid.UUID `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	CardNumber int64     `db:"card_number" json:"card_number" validate:"required"`
	Address    string    `db:"address" json:"address" validate:"required"`
	DoB        time.Time `db:"dob" json:"dob" validate:"required"`
	Sex        int       `db:"sex" json:"sex"`
	Village    int       `db:"village" json:"village" validate:"required"`
	District   int       `db:"district" json:"district" validate:"required"`
	Regency    int       `db:"regency" json:"regency" validate:"required"`
	Province   int       `db:"province" json:"province" validate:"required"`
	PostalCode int       `db:"postal_code" json:"postal_code" validate:"required"`
}
