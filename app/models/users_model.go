package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Email string `db:"email" json:"email" validate:"required,lte=255"`
	Username string `db:"username" json:"username" validate:"required,lte=255"`
	Password string `db:"password" json:"password" validate:"required,lte=255"`
	Phone int64 `db:"phone" json:"phone" validate:"required,lte=255"`
	Level int16 `db:"level" json:"level" validate:"required,lte=2"`
	Status int16 `db:"status" json:"status" validate:"required,lte=2"`
	Verified int16 `db:"verified" json:"verified" validate:"required,lte=2"`
	UserProfile UserProfile `json:"profile"`
}

type UserProfile struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	FirstName string `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName string `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Address string `db:"address" json:"address"`
	City string `db:"city" json:"city"`
	Province string `db:"province" json:"province"`	
}

func (u UserProfile) Value() (driver.Value, error){
	return json.Marshal(u)
}

func (u *UserProfile) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(j, &u)
}