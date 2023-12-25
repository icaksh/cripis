package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Trademark struct {
	ID             uuid.UUID      `db:"id" json:"id"`
	CreatedAt      time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at" json:"updated_at"`
	ExpiredAt      sql.NullTime   `db:"expired_at" json:"expired_at"`
	CreatedBy      uuid.UUID      `db:"created_by" json:"created_by"`
	RegisterNumber string         `db:"registration_number" json:"registration_number"`
	TrademarkName  string         `db:"trademark_name" json:"trademark_name" validate:"required"`
	Class          int            `db:"trademark_class" json:"trademark_class" validate:"required"`
	OwnerName      string         `db:"owner_name" json:"owner_name" validate:"required"`
	Address        string         `db:"address" json:"address" validate:"required"`
	Village        int            `db:"village" json:"village" validate:"required"`
	District       int            `db:"district" json:"district" validate:"required"`
	Regency        int            `db:"regency" json:"regency" validate:"required"`
	Province       int            `db:"province" json:"province" validate:"required"`
	Image          string         `db:"image" json:"image" validate:"required"`
	ApprovedAt     sql.NullTime   `db:"approved_at" json:"approved_at"`
	ApprovedBy     uuid.UUID      `db:"approved_by" json:"approved_by"`
	File           sql.NullString `db:"file" json:"file"`
	Status         int            `db:"status" json:"status"`
}
