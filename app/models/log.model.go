package models

import (
	"github.com/google/uuid"
	"time"
)

type Log struct {
	Id        int       `db:"id" json:"id"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
	CreatedBy uuid.UUID `db:"user_id" json:"user_id"`
	Email     string    `db:"email" json:"email"`
	UserName  string    `db:"user_name" json:"user_name"`
	Action    string    `db:"action" json:"action"`
}

type LogYear struct {
	Month         int `db:"month" json:"month"`
	Count         int `db:"count" json:"count"`
	ApprovedCount int `db:"approved_count" json:"approved_count"`
	RejectedCount int `db:"rejected_count" json:"rejected_count"`
}

type LogMonth struct {
	Day           int `db:"day" json:"day"`
	Count         int `db:"count" json:"count"`
	ApprovedCount int `db:"approved_count" json:"approved_count"`
	RejectedCount int `db:"rejected_count" json:"rejected_count"`
}

type LogUserYear struct {
	Month int `db:"month" json:"month"`
	Count int `db:"count" json:"count"`
}

type LogUserMonth struct {
	Day   int `db:"day" json:"day"`
	Count int `db:"count" json:"count"`
}
