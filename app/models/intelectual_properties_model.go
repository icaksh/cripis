package models

import (
	"time"

	"github.com/google/uuid"
)

type IntelectualProperties struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Number string `db:"number" json:"number"`
	Name string `db:"name" json:"name"`
	Status int16 `db:"status" json:"status"`
	File string `db:"file" json:"file"`
}

type IntelectualPropertiesType struct {
	ID int16 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Owner struct {
	FirstName string `db:"first_name" json:"first_name" validate:"required,lte=255"`
	LastName string `db:"last_name" json:"last_name" validate:"required,lte=255"`
	Address string `db:"address" json:"address"`
	City string `db:"city" json:"city"`
	Province string `db:"province" json:"province"`	
}

type Merk struct {
	ID uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
    Name string `db:"name" json:"name"`
    Holder uuid.UUID `db:"holder" json:"holder"`
    Owner Owner `db:"owner" json:"owner"`
    ApprovedAt time.Time `db:"approved_at" json:"approved_at"`
	ApprovedBy uuid.UUID `db:"approved_by" json:"approved_by"`
    File string `db:"file" json:"file"`
}