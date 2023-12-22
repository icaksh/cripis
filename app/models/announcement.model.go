package models

import (
	"github.com/google/uuid"
	"time"
)

type Announcement struct {
	Id          int       `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy   uuid.UUID `db:"created_by" json:"created_by"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Image       string    `db:"image" json:"image"`
}
