package models

type Credentials struct {
	Username string `db:"username" json:"username" validate:"required,lte=255"`
	Password string `db:"password" json:"password" validate:"required,lte=255"`
}

type SignUp struct {
	FirstName string `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName string `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Email string `db:"email" json:"email" validate:"required,lte=255"`
	Username string `db:"username" json:"username" validate:"required,lte=255"`
	Password string `db:"password" json:"password" validate:"required,lte=255"`
}