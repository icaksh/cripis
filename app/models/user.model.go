package models

import (
	"time"

	"github.com/google/uuid"
)

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type User struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Email     string    `db:"email" json:"email" validate:"required,lte=255"`
	FirstName string    `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName  string    `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Password  string    `db:"password" json:"password" validate:"required,lte=255"`
	Phone     int64     `db:"phone" json:"phone" validate:"required,lte=255"`
	Level     int16     `db:"level" json:"level" validate:"required,lte=2"`
	Status    int16     `db:"status" json:"status" validate:"required,lte=2"`
	Verified  int16     `db:"verified" json:"verified" validate:"required,lte=2"`
}
