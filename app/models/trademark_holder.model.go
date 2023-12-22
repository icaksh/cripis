package models

import (
	"github.com/google/uuid"
	"time"
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
