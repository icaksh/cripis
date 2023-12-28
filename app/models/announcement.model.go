package models

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Announcement struct {
	Id          int            `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	CreatedBy   uuid.UUID      `db:"created_by" json:"created_by"`
	Title       string         `db:"title" json:"title" validate:"required"`
	Description string         `db:"description" json:"description" validate:"required"`
	Image       sql.NullString `db:"image" json:"image"`
}

type AnnouncementCreation struct {
	Title       string         `db:"title" json:"title" validate:"required"`
	Description string         `db:"description" json:"description" validate:"required"`
	Image       sql.NullString `db:"image" json:"image"`
}
