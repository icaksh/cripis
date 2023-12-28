package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Trademark struct {
	ID                uuid.UUID      `db:"id" json:"id"`
	CreatedAt         time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updated_at"`
	ExpiredAt         sql.NullTime   `db:"expired_at" json:"expired_at"`
	CreatedBy         uuid.UUID      `db:"created_by" json:"created_by"`
	RegisterNumber    string         `db:"registration_number" json:"registration_number"`
	TrademarkName     string         `db:"trademark_name" json:"trademark_name"`
	Class             int            `db:"trademark_class" json:"trademark_class"`
	OwnerName         string         `db:"owner_name" json:"owner_name"`
	Address           string         `db:"address" json:"address"`
	Village           int            `db:"village" json:"village"`
	District          int            `db:"district" json:"district"`
	Regency           int            `db:"regency" json:"regency"`
	Province          int            `db:"province" json:"province"`
	Image             string         `db:"image" json:"image"`
	SMECertificate    string         `db:"sme_certificate" json:"sme_certificate"`
	RegisterSignature string         `db:"register_signature" json:"register_signature"`
	ApprovedAt        sql.NullTime   `db:"approved_at" json:"approved_at"`
	ApprovedBy        uuid.UUID      `db:"approved_by" json:"approved_by"`
	File              sql.NullString `db:"file" json:"file"`
	Status            int            `db:"status" json:"status"`
	Notes             sql.NullString `db:"notes" json:"notes"`
}

type TrademarkSearch struct {
	ID                uuid.UUID      `db:"id" json:"-"`
	CreatedAt         time.Time      `db:"created_at" json:"-"`
	UpdatedAt         time.Time      `db:"updated_at" json:"-"`
	ExpiredAt         sql.NullTime   `db:"expired_at" json:"expired_at"`
	CreatedBy         uuid.UUID      `db:"created_by" json:"-"`
	RegisterNumber    string         `db:"registration_number" json:"registration_number"`
	TrademarkName     string         `db:"trademark_name" json:"trademark_name"`
	Class             int            `db:"trademark_class" json:"trademark_class"`
	OwnerName         string         `db:"owner_name" json:"owner_name"`
	Address           string         `db:"address" json:"address"`
	Village           int            `db:"village" json:"village"`
	District          int            `db:"district" json:"district"`
	Regency           int            `db:"regency" json:"regency"`
	Province          int            `db:"province" json:"province"`
	Image             string         `db:"image" json:"image"`
	SMECertificate    string         `db:"sme_certificate" json:"-"`
	RegisterSignature string         `db:"register_signature" json:"-"`
	ApprovedAt        sql.NullTime   `db:"approved_at" json:"approved_at"`
	ApprovedBy        uuid.UUID      `db:"approved_by" json:"-"`
	File              sql.NullString `db:"file" json:"-"`
	Status            int            `db:"status" json:"status"`
	Notes             sql.NullString `db:"notes" json:"-"`
}

type TrademarkRegistrationRequest struct {
	TrademarkName     string `form:"trademark_name" validate:"required"`
	Class             int    `form:"trademark_class" validate:"required"`
	OwnerName         string `form:"owner_name" validate:"required"`
	Address           string `form:"address" validate:"required"`
	Village           int    `form:"village" validate:"required"`
	District          int    `form:"district" validate:"required"`
	Regency           int    `form:"regency" validate:"required"`
	Province          int    `form:"province" validate:"required"`
	SMECertificate    string `db:"sme_certificate" json:"sme_certificate"`
	RegisterSignature string `db:"register_signature" json:"register_signature"`
	Notes             string `db:"notes" json:"notes"`
}

type TrademarkEditRequest struct {
	ID                uuid.UUID `form:"id" validate:"required"`
	TrademarkName     string    `form:"trademark_name" validate:"required"`
	Class             int       `form:"trademark_class" validate:"required"`
	OwnerName         string    `form:"owner_name" validate:"required"`
	Address           string    `form:"address" validate:"required"`
	Village           int       `form:"village" validate:"required"`
	District          int       `form:"district" validate:"required"`
	Regency           int       `form:"regency" validate:"required"`
	Province          int       `form:"province" validate:"required"`
	Image             string    `form:"image" validate:"required"`
	SMECertificate    string    `form:"sme_certificate" json:"sme_certificate"`
	RegisterSignature string    `form:"register_signature" validate:"required"`
}

type TrademarkStatusRequest struct {
	ID     string `json:"id" validate:"required"`
	Status int    `json:"status" validate:"required"`
	Notes  string `json:"notes"`
}
