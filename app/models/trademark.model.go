package models

import (
	"time"

	"github.com/google/uuid"
)

type TrademarkHolder struct {
	Id         uuid.UUID `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	RegisterId uuid.UUID `db:"register_id" json:"register_id"`
	FullName   string    `db:"full_name" json:"full_name" validate:"required"`
	Address    string    `db:"address" json:"address" validate:"required"`
	Village    int       `db:"village" json:"village" validate:"required"`
	District   int       `db:"district" json:"district" validate:"required"`
	Regency    int       `db:"regency" json:"regency" validate:"required"`
	Province   int       `db:"province" json:"province" validate:"required"`
}

type TrademarkOwner struct {
	FullName string `db:"full_name" json:"full_name"`
	Address  string `db:"address" json:"address"`
}

type Trademark struct {
	ID             uuid.UUID       `db:"id" json:"id"`
	CreatedAt      time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time       `db:"updated_at" json:"updated_at"`
	ExpiredAt      time.Time       `db:"expired_at json:"expired_at"`
	RegisterNumber string          `db:"number" json:"number"`
	Name           string          `db:"name" json:"name" validate:"required"`
	Class          int             `db:"class" json:"class" validate:"required"`
	Holder         TrademarkHolder `db:"holder" json:"holder"`
	Owner          TrademarkOwner  `db:"owner" json:"owner"`
	RegistrationId uuid.UUID       `db:"registration_id" json:"registration_id"`
	ApprovedAt     time.Time       `db:"approved_at" json:"approved_at"`
	ApprovedBy     uuid.UUID       `db:"approved_by" json:"approved_by"`
	File           string          `db:"file" json:"file"`
	Status         int             `db:"status" json:"status"`
}
