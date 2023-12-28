package models

import "time"

type Credentials struct {
	Email    string `db:"email" json:"email" validate:"required,lte=255"`
	Password string `db:"password" json:"password" validate:"required,lte=255"`
	Remember bool   `db:"remember" json:"remember"`
}

type ResetPasswordRequest struct {
	Email string `db:"email" json:"email" validate:"required,lte=255"`
}

type ResetPassword struct {
	AccessToken string `json:"access_token" validate:"required"`
	Password    string `db:"password" json:"email" validate:"required,lte=255"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type SignUp struct {
	FirstName  string    `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName   string    `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Email      string    `db:"email" json:"email" validate:"required,lte=255"`
	Password   string    `db:"password" json:"password" validate:"required,lte=255"`
	Roles      int16     `db:"roles" json:"roles"`
	Status     int16     `db:"status" json:"status" `
	Verified   int16     `db:"verified" json:"verified"`
	CardNumber int64     `db:"card_number" json:"card_number" validate:"required"`
	Address    string    `db:"address" json:"address" validate:"required"`
	DoB        time.Time `db:"dob" json:"dob" validate:"required"`
	Sex        int       `db:"sex" json:"sex"`
	Village    int       `db:"village" json:"village" validate:"required"`
	District   int       `db:"district" json:"district" validate:"required"`
	Regency    int       `db:"regency" json:"regency" validate:"required"`
	Province   int       `db:"province" json:"province" validate:"required"`
	PostalCode int       `db:"postal_code" json:"postal_code" validate:"required"`
}
