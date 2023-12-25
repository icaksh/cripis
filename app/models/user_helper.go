package models

type UserRoles struct {
	ID   int    `db:"id" json:"id"`
	Role string `db:"role" json:"role"`
}
