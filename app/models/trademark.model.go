package models

import (
	"time"

	"github.com/google/uuid"
)

type Trademark struct {
	ID             uuid.UUID `db:"id" json:"id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	ExpiredAt      time.Time `db:"expired_at" json:"expired_at"`
	RegisterNumber string    `db:"number" json:"number"`
	Name           string    `db:"name" json:"name" validate:"required"`
	Class          int       `db:"class" json:"class" validate:"required"`
	Holder         uuid.UUID `db:"holder_id" json:"holder_id"`
	RegistrationId uuid.UUID `db:"registration_id" json:"registration_id"`
	ApprovedAt     time.Time `db:"approved_at" json:"approved_at"`
	ApprovedBy     uuid.UUID `db:"approved_by" json:"approved_by"`
	File           string    `db:"file" json:"file"`
	Status         int       `db:"status" json:"status"`
}
