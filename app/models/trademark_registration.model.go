package models

import (
	"github.com/google/uuid"
	"time"
)

type TrademarkRegistration struct {
	ID                 uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	RegistrationNumber string    `db:"registration_number" json:"registration_number"`
	RegisterId         uuid.UUID `db:"register_id" json:"register_id"`
	ApprovalId         uuid.UUID `db:"approval_id" json:"approval_id"`
	ApprovalDate       time.Time `db:"approval_date" json:"approval_date"`
	SMECertificate     string    `db:"sme_certificate" json:"sme_certificate" validate:"required"`
	RegisterSignature  string    `db:"register_signature" json:"register_signature" validate:"required"`
	Status             int16     `db:"status" json:"status"`
	Notes              string    `db:"notes" json:"notes"`
}
