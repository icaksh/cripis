package models

import (
	"github.com/google/uuid"
	"time"
)

type TrademarkRegistration struct {
	ID                uuid.UUID `db:"trademark_id" json:"trademark_id" validate:"required,uuid"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy         uuid.UUID `db:"register_id" json:"register_id"`
	ApprovalId        uuid.UUID `db:"approval_id" json:"approval_id"`
	ApprovalDate      time.Time `db:"approval_date" json:"approval_date"`
	SMECertificate    string    `db:"sme_certificate" json:"sme_certificate"`
	RegisterSignature string    `db:"register_signature" json:"register_signature" validate:"required"`
	Status            int16     `db:"status" json:"status"`
	Notes             string    `db:"notes" json:"notes"`
}

type TrademarkRegistrationRequest struct {
	TrademarkName string `form:"trademark_name" validate:"required"`
	Class         int    `form:"trademark_class" validate:"required"`
	OwnerName     string `form:"owner_name" validate:"required"`
	Address       string `form:"address" validate:"required"`
	Village       int    `form:"village" validate:"required"`
	District      int    `form:"district" validate:"required"`
	Regency       int    `form:"regency" validate:"required"`
	Province      int    `form:"province" validate:"required"`
}
