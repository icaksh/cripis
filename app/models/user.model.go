package models

import (
	"time"

	"github.com/google/uuid"
)

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Email      string    `db:"email" json:"email" validate:"required,lte=255"`
	FirstName  string    `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName   string    `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Password   string    `db:"password" json:"password" validate:"required,lte=255"`
	Phone      int64     `db:"phone" json:"phone" validate:"required,lte=255"`
	Roles      int16     `db:"roles" json:"roles" validate:"required,lte=2"`
	Status     int16     `db:"status" json:"status" validate:"required,lte=2"`
	Verified   int16     `db:"verified" json:"verified" validate:"required,lte=2"`
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
