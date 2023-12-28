package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id" json:"id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Email      string    `db:"email" json:"email"`
	FirstName  string    `db:"first_name" json:"first_name"`
	LastName   string    `db:"last_name" json:"last_name"`
	Password   string    `db:"password" json:"-"`
	Roles      int16     `db:"roles" json:"roles"`
	Status     int16     `db:"status" json:"status"`
	Verified   int16     `db:"verified" json:"verified"`
	CardNumber int64     `db:"card_number" json:"card_number"`
	Address    string    `db:"address" json:"address"`
	DoB        time.Time `db:"dob" json:"dob"`
	Sex        int       `db:"sex" json:"sex"`
	Village    int       `db:"village" json:"village"`
	District   int       `db:"district" json:"district"`
	Regency    int       `db:"regency" json:"regency"`
	Province   int       `db:"province" json:"province"`
	PostalCode int       `db:"postal_code" json:"postal_code"`
}

type UserProfilesRequest struct {
	Email      string    `db:"email" json:"email" validate:"required"`
	FirstName  string    `db:"first_name" json:"first_name" validate:"required"`
	LastName   string    `db:"last_name" json:"last_name" validate:"required"`
	Password   string    `db:"password" json:"-"`
	CardNumber int64     `db:"card_number" json:"card_number" validate:"required"`
	Address    string    `db:"address" json:"address" validate:"required"`
	DoB        time.Time `db:"dob" json:"dob" validate:"required"`
	Sex        int       `db:"sex" json:"sex" validate:"required"`
	Village    int       `db:"village" json:"village" validate:"required"`
	District   int       `db:"district" json:"district" validate:"required"`
	Regency    int       `db:"regency" json:"regency" validate:"required"`
	Province   int       `db:"province" json:"province" validate:"required"`
	PostalCode int       `db:"postal_code" json:"postal_code" validate:"required"`
}

type UserRolesRequest struct {
	ID       uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Password string    `db:"password" json:"password" validate:"required"`
	Roles    int16     `db:"roles" json:"roles" validate:"required"`
	Status   int16     `db:"status" json:"status" validate:"required"`
	Verified int16     `db:"verified" json:"verified" validate:"required"`
}

type UserPasswordRequest struct {
	ID          uuid.UUID `db:"id" json:"id"`
	OldPassword string    `db:"old_password" json:"old_password" validate:"required"`
	Password    string    `db:"password" json:"password" validate:"required"`
}
